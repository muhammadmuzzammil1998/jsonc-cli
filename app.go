// MIT License

// Copyright (c) 2020 Muhammad Muzzammil

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"muzzammil.xyz/go-check"
	"muzzammil.xyz/jsonc"
)

type flagVars struct {
	version  bool
	validate bool
	stdout   bool
	asis     bool
	inFile   string
	outFile  string
}

var (
	version string
	flags   flagVars
)

func init() {
	flag.BoolVar(&flags.version, "version", false, "Prints version")
	flag.BoolVar(&flags.validate, "v", false, "Checks if a jsonc file is valid or not (use with -i)")
	flag.BoolVar(&flags.stdout, "p", false, "Just prints the output instead of saving it")
	flag.BoolVar(&flags.asis, "asis", false, "Prints the JSONC as-is")

	flag.StringVar(&flags.inFile, "i", "", "Source file")
	flag.StringVar(&flags.outFile, "o", "", "Destination file")

	flag.Parse()

	version = "1.0"
}

func main() {
	if flags.version {
		fmt.Println(" 🎉 jsonc-cli")
		fmt.Println("  🤖 Version:", version)
		fmt.Println("  🌐 Source: https://muzzammil.xyz/jsonc-cli")
		return
	}

	if strings.Trim(flags.inFile, " ") == "" {
		log.Fatalln(" ❌ No source provided. Try using -i [filename]")
	}

	if _, err := os.Stat(flags.inFile); os.IsNotExist(err) {
		if !strings.HasSuffix(flags.inFile, ".jsonc") {
			flags.inFile += ".jsonc"
		}
		if _, err := os.Stat(flags.inFile); os.IsNotExist(err) {
			fmt.Println(flags.inFile)
			log.Fatalln(" ❌ Source provided does not exist")
		}
	}

	if flags.validate {
		jc, _, err := jsonc.ReadFromFile(flags.inFile)
		check.Error(err, log.Fatalln)
		if jsonc.Valid(jc) {
			fmt.Println(" 👍 Valid - the jsonc file", flags.inFile, "is valid")
		} else {
			fmt.Println(" 👎 Invalid - the jsonc file", flags.inFile, "is invalid")
		}
		return
	}

	if flags.asis {
		print(true)
		return
	}

	if flags.stdout {
		print(false)
		return
	}

	if strings.Trim(flags.outFile, " ") == "" {
		s := strings.Split(flags.inFile, ".")
		flags.outFile = fmt.Sprintf("%s.json", strings.Join(s[:len(s)-1], "."))
		fmt.Println(" 🤖 No destination provided. Using source filename", flags.outFile)
	}

	if !strings.HasSuffix(flags.outFile, ".json") {
		flags.outFile += ".json"
	}

	if _, err := os.Stat(flags.outFile); !os.IsNotExist(err) {
		if !overwrite(" ❓ File %s exists. Overwrite? (y/n): ") {
			return
		}
	}

	if flags.outFile == flags.inFile {
		if !overwrite(" ❓ Destination file and source file are the same. This will overwrite the source file (%s). Do you still wish to continue? (y/n): ") {
			return
		}
	}

	jc, j, err := jsonc.ReadFromFile(flags.inFile)
	check.Error(err, log.Fatalln)

	err = ioutil.WriteFile(flags.outFile, j, 0644)
	check.Error(err, log.Fatalln)

	fmt.Print(" ✅ Coversion successful! ")

	if !jsonc.Valid(jc) {
		fmt.Print("However, the jsonc file seems to be invalid. ")
	}
	fmt.Println("JSON written to", flags.outFile)
}

func overwrite(s string) bool {
	fmt.Printf(s, flags.outFile)
	reader := bufio.NewReader(os.Stdin)
	r, _, _ := reader.ReadLine()

	if string(r) != "y" {
		log.Println(" 🚧 Operation cancelled by user")
		return false
	}

	return true
}

func print(asis bool) {
	jc, j, err := jsonc.ReadFromFile(flags.inFile)
	check.Error(err, log.Fatalln)
	if asis {
		fmt.Println(strings.Replace(string(jc), "\r", "", -1))
		return
	}
	fmt.Println(strings.Replace(string(j), "\r", "", -1))
}
