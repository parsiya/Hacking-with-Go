# Hacking with Go
This is my attempt at filling the gap in Go security tooling. When starting to learn Go, I learned from a lot of tutorials but I could find nothing that is geared towards security professionals.

These documents are based on the `Gray/Black Hat Python/C#` series of books. I like their style. Join me as I learn more about Go and attempt to introduce Go to security denizens without fluff and through practical applications.

## Table of Contents

- [01 - Setting up a Go development environment](content/01.md)
- [02 - Basics](content/02.0.md)
    + [02.1 - Packages, functions, variables, basic types, casting and constants](content/02.1.md)
    + [02.2 - for, if, else, switch and defer](content/02.2.md)
    + [02.3 - Pointers, structs, arrays, slices and range](content/02.3.md)
    + [02.4 - Methods and interfaces](content/02.4.md)
    + [02.5 - Printf, Scanf, bufio readers and maps](content/02.5.md)
    + [02.6 - Goroutines and channels](content/02.6.md)
    + [02.7 - Error handling](content/02.7.md)
- [03 - Useful Go packages - WIP](content/03.0.md)
    + [03.1 - flag package](content/03.1.md)
    + [03.2 - log package](content/03.2.md)
- [04- Go networking](content/04.0.md)
    + [04.1 - Basic TCP and UDP clients](content/04.1.md)
    + [04.2 - TCP servers](content/04.2.md)
    + [04.3 - TCP proxy](content/04.3.md)
    + [04.4 - SSH clients](content/04.4.md)

## Code

- [01 - Setting up a Go development environment](code/01)
- [02 - Basics](code/02)
- [03 - Useful Go packages](code/03)
- [04 - Go networking](code/04)

### FAQ

**Why not use Python?**  
Python reigns supreme in security and for good reason. It's a powerful programming language. There are a lot of supporting libraries out there both in security and for general use. However, I think Go has its merits and can occupy a niche.

**Why not use other tutorials?**  
There are a lot of tutorials for Go out there. None are geared towards security professionals. Our needs are different, we want to write quick and dirty scripts that work (hence Python is so successful). Similar guides are available in Python and other programming languages.

**Why not just use Black Hat Go?**  
There's a book named [Black Hat Go][black-hat-go] by No Starch in production. Looking at the author list, I cannot  compete with them in terms of experience and knowledge. That is a proper book with editors and a publisher while I am just some rando learning as I go. It does not take a lot of CPU power to decide the book will be better.

But the book is not out yet. Today is December 6th 2017 and the book is marked for release in August 2018. The book page does not have any released chapters or material. We can assume it's going to be similar to the other `gray|black hat` books. This repository and that book are inevitably going to have a lot of overlap. Think of this as warm up while we wait.

**Rewrite in Rust/Haskell**  
Honestly I will be very much interested in a similar guide for Rust/Haskell geared for security people. Please let me know if you create one.

## Feedback
I am always interested in feedback. There will be errors and there are always better ways to code. Please create an issue here. If this has helped you please let me know, it helps with the grind.

## Other resources
There are tons of Go resources online. I am going to try not to re-hash what has been already created. Hacking with Go is not meant to be self-contained. When in doubt, use one of these resources or just search.

The following links helped me get started:

- GoDoc: [https://godoc.org/][go-doc]
- A Tour of Go: [https://tour.golang.org/][tour-of-go]
- Go by Example: [https://gobyexample.com/][go-by-example]
- Go playground: [https://play.golang.org/][go-playground]
- Effective Go: [https://golang.org/doc/effective_go.html][effective-go]

## License

- Code in this repository is licensed under [GPLv3](LICENSE).
- Non-code content is licensed under [Creative Commons Attribution-NonCommercial 4.0][CC-4] (CC BY-NC 4.0).

<!-- Links -->

[black-hat-go]: https://www.nostarch.com/blackhatgo
[go-doc]: https://godoc.org/
[tour-of-go]: https://tour.golang.org/
[go-by-example]: https://gobyexample.com/
[go-playground]: https://play.golang.org/
[CC-4]: https://creativecommons.org/licenses/by-nc-sa/4.0/
[effective-go]: https://golang.org/doc/effective_go.html
