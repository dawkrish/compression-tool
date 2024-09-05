# Compression Tool

Inspired by this [coding challenge](https://codingchallenges.fyi/challenges/challenge-huffman)

## Installation

All the code blocks are meant to be run in a terminal

* Install Go
You can download Go from [[here]](https://go.dev/dl/)

* Install the binary
  ```bash
  go install github.com/dawkrish/compression-tool@latest
  ```

This will place the binary in `GOPATH`. You can check your `GOPATH` by typing the command `go env`

* Confirm
  ```bash
  which compression-tool
  ```

  should output something like

  ```bash
    /Users/<username>/go/bin/compression-tool
  ```

## Usage

* There are only 2 flags
  * `-h` for help
  * `-d` for decompressing


* Example
  * `./compression-tool <file-name>` : will compress the content of the file into a new file `<file-name_compressed.txt>`
  * `./compression-tool -d <file-name>` : will decompress the content of the file into a new file `<file-name_decompressed.txt>`

Note : the `-d` flag must come before the `<file-name>`. You can look the reason for it [[here]](https://gobyexample.com/command-line-flags)
> Note that the flag package requires all flags to appear before positional arguments (otherwise the flags will be interpreted as positional arguments)


## Contribution
 The project was simple and small, but effective, if you face any issue, please raise an issue
 
