package main

import (
	"fmt" 
	"os"
	"flag"
	"bufio"
	"strings"
)

func main() {
	if false == isPipeline() {
		fmt.Println("This command can only be used with pipelines.")
		return
	}

	delimiter := flag.String("d", " ", "Delimiter: Used to split input. Defaults to whitespace")
	index := flag.Int("i", 0, "Index: Print only the values from this index.")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		str := s.Text()

		flag.Parse()

		out := strings.Split(str, *delimiter)

		if isFlagPassed("i") {
			if *index < 0 || *index >= len(out) {
				fmt.Printf("Only %v item(s) in input. -i should be between 0 and %v\n", len(out), len(out) -1)
				return
			}

			fmt.Println(out[*index])
		} else {
			for _, e := range out {
				fmt.Println(e)
			}
		}
	}
}

func isPipeline() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}

func isFlagPassed(name string) bool {
	r := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            r = true
        }
    })
    return r
}