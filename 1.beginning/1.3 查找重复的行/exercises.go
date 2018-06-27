package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// 判断是否输入文件
	if len(files) == 0 {
		countLines(counts, os.Stdin)
	} else {
		for i := 0; i < len(files); i++ {
			newCounts := make(map[string]int)
			f, err := os.Open(files[i])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			countLines(newCounts, f)
			f.Close()
			fmt.Printf("filename: %v\n", files[i])
			dup(newCounts)
		}
	}
	dup(counts)

}

func dup(count map[string]int) {
	for line, n := range count {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}

func countLines(counts map[string]int, f *os.File) {
	// 将标准输入的内容来判断是否有重复的行
	input := bufio.NewScanner(f)
	for input.Scan() {
		//counts[input.Text()]++ 等于line := input.Text() ;counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
}
