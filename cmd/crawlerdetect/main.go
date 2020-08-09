package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/x-way/crawlerdetect"
)

func main() {
	inverse := flag.Bool("v", false, "show non-crawler lines")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if crawlerdetect.IsCrawler(line) != *inverse {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
