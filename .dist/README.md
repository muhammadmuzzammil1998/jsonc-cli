# go-dist

Distribution automation script for programs written in Go

## Working with `build.sh`

[![asciicast](https://asciinema.org/a/381907.svg)](https://asciinema.org/a/381907)

## What the script does

It builds executable binaries for different operating systems using the `GOOS` feature. The script also builds debian packages for both architectures.
Currently supports these systems:

```
1. darwin/386        8. linux/amd64    15. linux/mips64    22. openbsd/arm
2. darwin/amd64      9. linux/arm      16. linux/mips64le  23. plan9/386
3. dragonfly/amd64  10. linux/arm64    17. netbsd/386      24. plan9/amd64
4. freebsd/386      11. linux/ppc64    18. netbsd/amd64    25. solaris/amd64
5. freebsd/amd64    12. linux/ppc64le  19. netbsd/arm      26. windows/386
6. freebsd/arm      13. linux/mips     20. openbsd/386     27. windows/amd64
7. linux/386        14. linux/mipsle   21. openbsd/amd64
```

## Usage

First of all, navigate to your `go` source directory where you run `go build` command to build the binary. Then follow these -

### Making `.dist` directory and navigating to it

```bash
$ mkdir .dist
$ cd .dist
```

### `git` it

```bash
$ git clone https://github.com/muhammadmuzzammil1998/go-dist .
```

Don't forget to clone in the same directory `(.)`

### Making `build.sh` executable

```bash
$ chmod 777 ./build.sh
```

### Running `build.sh`

```bash
$ ./build.sh
```

## Couple of things to note

1. You should be in the main go directory of your project. That is, the place you run `go test`, `go build` etc. commands.
2. This repository (go-dist) should be a direct sub-directory to the main go directory. Structure should look like this:

```
- ./
-- main.go
-- other_files.Go
-- main_test.Go
-- README.md
-- .dist/
-- -- build.sh

```
