package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
)

var filepath string

func init() {
	flag.StringVar(&filepath, "filepath", "", "指定文件路径")
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

	h := sha256.New()
	buf := make([]byte, 20)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		_, err = h.Write(buf[:n])
		if err != nil {
			log.Printf("write buf to hash err: %s", err)
		}
	}

	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
