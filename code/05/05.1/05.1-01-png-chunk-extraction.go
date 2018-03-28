// Simple PNG parser. Can be used to discover and extract hidden chunks.
// Minimal error handling, does not play well with malformed chunks and doesn't
// check chunk CRC32 checksums.

package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// 89 50 4E 47 0D 0A 1A 0A
var PNGHeader = "\x89\x50\x4E\x47\x0D\x0A\x1A\x0A"
var iHDRlength = 13

// uInt32ToInt converts a 4 byte big-endian buffer to int.
func uInt32ToInt(buf []byte) (int, error) {
	if len(buf) == 0 || len(buf) > 4 {
		return 0, errors.New("invalid buffer")
	}
	return int(binary.BigEndian.Uint32(buf)), nil
}

// Each chunk starts with a uint32 length (big endian), then 4 byte name,
// then data and finally the CRC32 of the chunk data.
type Chunk struct {
	Length int    // chunk data length
	CType  string // chunk type
	Data   []byte // chunk data
	Crc32  []byte // CRC32 of chunk data
}

// Populate will read bytes from the reader and populate a chunk.
func (c *Chunk) Populate(r io.Reader) error {

	// Four byte buffer.
	buf := make([]byte, 4)

	// Read first four bytes == chunk length.
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	// Convert bytes to int.
	// c.length = int(binary.BigEndian.Uint32(buf))
	var err error
	c.Length, err = uInt32ToInt(buf)
	if err != nil {
		return errors.New("cannot convert length to int")
	}

	// Read second four bytes == chunk type.
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	c.CType = string(buf)

	// Read chunk data.
	tmp := make([]byte, c.Length)
	if _, err := io.ReadFull(r, tmp); err != nil {
		return err
	}
	c.Data = tmp

	// Read CRC32 hash
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	// We don't really care about checking the hash.
	c.Crc32 = buf

	return nil
}

// -----------

type PNG struct {
	Width             int
	Height            int
	BitDepth          int
	ColorType         int
	CompressionMethod int
	FilterMethod      int
	InterlaceMethod   int
	chunks            []*Chunk // Not exported == won't appear in JSON string.
	NumberOfChunks    int
}

// Parse IHDR chunk.
// https://golang.org/src/image/png/reader.go?#L142 is your friend.
func (png *PNG) parseIHDR(iHDR *Chunk) error {
	if iHDR.Length != iHDRlength {
		errString := fmt.Sprintf("invalid IHDR length: got %d - expected %d",
			iHDR.Length, iHDRlength)
		return errors.New(errString)
	}

	// IHDR: http://www.libpng.org/pub/png/spec/1.2/PNG-Chunks.html#C.IHDR

	// Width:              4 bytes
	// Height:             4 bytes
	// Bit depth:          1 byte
	// Color type:         1 byte
	// Compression method: 1 byte
	// Filter method:      1 byte
	// Interlace method:   1 byte

	tmp := iHDR.Data
	var err error

	png.Width, err = uInt32ToInt(tmp[0:4])
	if err != nil || png.Width <= 0 {
		errString := fmt.Sprintf("invalid width in iHDR - got %x", tmp[0:4])
		return errors.New(errString)
	}

	png.Height, err = uInt32ToInt(tmp[4:8])
	if err != nil || png.Height <= 0 {
		errString := fmt.Sprintf("invalid height in iHDR - got %x", tmp[4:8])
		return errors.New(errString)
	}

	png.BitDepth = int(tmp[8])
	png.ColorType = int(tmp[9])

	// Only compression method 0 is supported
	if int(tmp[10]) != 0 {
		errString := fmt.Sprintf("invalid compression method - expected 0 - got %x",
			tmp[10])
		return errors.New(errString)
	}
	png.CompressionMethod = int(tmp[10])

	// Only filter method 0 is supported
	if int(tmp[11]) != 0 {
		errString := fmt.Sprintf("invalid filter method - expected 0 - got %x",
			tmp[11])
		return errors.New(errString)
	}
	png.FilterMethod = int(tmp[11])

	// Only interlace methods 0 and 1 are supported
	if int(tmp[12]) != 0 {
		errString := fmt.Sprintf("invalid interlace method - expected 0 or 1 - got %x",
			tmp[12])
		return errors.New(errString)
	}
	png.InterlaceMethod = int(tmp[12])

	return nil
}

// Populate populates the PNG fields (and other fields).
func (png *PNG) Populate() error {
	if err := png.parseIHDR(png.chunks[0]); err != nil {
		return err
	}
	png.NumberOfChunks = len(png.chunks)
	return nil
}

// PrintChunks will return a string containign chunk number, name and the first 20
// bytes of each chunk.
func (png PNG) PrintChunks() string {
	var output string
	for i, c := range png.chunks {
		output += fmt.Sprintf("-----------\n")
		output += fmt.Sprintf("Chunk # %d\n", i)
		output += fmt.Sprintf("Chunk length: %d\n", c.Length)
		output += fmt.Sprintf("Chunk type: %v\n", c.CType)

		limit := 20
		if len(c.Data) < 20 {
			limit = len(c.Data)
		}
		output += fmt.Sprintf("Chunk data (20 bytes): % x\n", c.Data[:limit])
	}
	return output
}

// ExportChunk returns a chunk's data as []byte. Numbering starts from 0.
func (png PNG) ExportChunk(chunkNumber int) ([]byte, error) {
	if chunkNumber >= png.NumberOfChunks {
		errString := fmt.Sprintf("invalid chunk number. Got: %d, "+
			"file has %d chunks. Chunk numbers starts from zero.",
			chunkNumber, len(png.chunks))
		return nil, errors.New(errString)
	}
	return png.chunks[chunkNumber].Data, nil
}

// ----------

var (
	filename  string
	chunkFlag bool
)

func init() {
	// Setup flags.
	flag.StringVar(&filename, "file", "", "input file")
	flag.BoolVar(&chunkFlag, "c", false, "print chunks")
	flag.BoolVar(&chunkFlag, "chunks", false, "print chunks")

	// Parse flags.
	flag.Parse()
}

func main() {

	imgFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	// Read first 8 bytes == PNG header.
	header := make([]byte, 8)
	// Read CRC32 hash
	if _, err := io.ReadFull(imgFile, header); err != nil {
		panic(err)
	}
	if string(header) != PNGHeader {
		fmt.Printf("Wrong PNG header.\nGot %x - Expected %x\n", header, PNGHeader)
		return
	}

	var png PNG

	// Reset err
	err = nil
	for err == nil {
		var c Chunk
		err = (&c).Populate(imgFile)
		// Drop the last empty chunk.
		if c.CType != "" {
			png.chunks = append(png.chunks, &c)
		}
	}

	if err := (&png).Populate(); err != nil {
		fmt.Println("Failed to populate PNG fields.")
		panic(err)
	}

	fmt.Println("PNG info")
	jsoned, err := json.MarshalIndent(png, "", "    ")
	if err != nil {
		fmt.Printf("%+v", png)
	} else {
		fmt.Println(string(jsoned))
	}

	if chunkFlag {
		fmt.Println("\nPrinting chunks\n----------")
		fmt.Println(png.PrintChunks())
		fmt.Println("----------")
	}

	// Reset err
	err = nil
	for {
		fmt.Printf("Enter chunk number or -1 to quit: ")
		var n int
		fmt.Scanln(&n)
		// Quit if user enters -1.
		if n == -1 {
			break
		}
		// Export chunk.
		c, err := png.ExportChunk(n)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		chunkfilename := fmt.Sprintf("%s-chunk-%d", filename, n)
		chunkfile, err := os.Create(chunkfilename)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		defer chunkfile.Close()
		chunkfile.Write(c)
		fmt.Println("Chunk saved to ", chunkfilename)
	}
	fmt.Println("Qutting!")
}
