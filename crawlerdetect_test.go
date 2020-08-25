package crawlerdetect

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestCrawlers(t *testing.T) {
	testFromFile(t, "testdata/crawlers.txt", true)
}

func TestDevices(t *testing.T) {
	testFromFile(t, "testdata/devices.txt", false)
}

func testFromFile(t *testing.T, filename string, want bool) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		if got := IsCrawler(input); got != want {
			t.Errorf("IsCrawler(%v) = %v; want %v", input, got, want)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func TestRegexpUpdate(t *testing.T) {
	crawlerDetect := New()
	// test with -race
	go crawlerDetect.IsCrawler("foo")
	crawlerDetect.SetCrawlers([]string{"bar"})
	crawlerDetect.SetExclusions([]string{"baz"})
}
