// +build gofuzz

package iprange

func Fuzz(data []byte) int {
	_, err := ParseList(string(data))
	if err != nil {
		return 0
	}
	return 1
}
