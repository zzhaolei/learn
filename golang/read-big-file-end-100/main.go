package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	lineNum  int
	filepath string
)

func init() {
	flag.StringVar(&filepath, "filepath", "", "指定文件")
	flag.IntVar(&lineNum, "line", 0, "指定行数")
}

func fixRingSlice(ring []string, ringIdx int) {
	if ringIdx == 0 {
		return
	}
	for i := len(ring) - 1; i > 0; i-- {
		ringIdx--
		if ringIdx < 0 {
			ringIdx = 0
		}
		ring[i], ring[ringIdx] = ring[ringIdx], ring[i]
	}
}

func main() {
	flag.Parse()
	if filepath == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := make([]string, lineNum)
	ring := 0

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		lines[ring] = line
		ring += 1
		if ring >= lineNum {
			ring = 0
		}
	}

	fixRingSlice(lines, ring)
	for _, line := range lines {
		fmt.Println(line)
	}
}
