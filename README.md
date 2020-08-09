# crawlerdetect

[![CircleCI](https://circleci.com/gh/x-way/crawlerdetect.svg?style=svg)](https://circleci.com/gh/x-way/crawlerdetect)
[![Go Report Card](https://goreportcard.com/badge/github.com/x-way/crawlerdetect)](https://goreportcard.com/report/github.com/x-way/crawlerdetect)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/x-way/crawlerdetect)](https://pkg.go.dev/github.com/x-way/crawlerdetect)

## About
**crawlerdetect** is a Go version of PHP class @[CrawlerDetect](https://github.com/JayBizzle/Crawler-Detect). 

It helps to detect bots/crawlers/spiders via the user agent and other HTTP-headers. Currently able to detect 1,000's of bots/spiders/crawlers.

## Installation

`go get github.com/x-way/crawlerdetect`

## Basic Usage
```
import "fmt"
import "github.com/x-way/crawlerdetect"

func main() {
    uastring := "curl/7.54.0"
    if crawlerdetect.IsCrawler(uastring) {
        fmt.Println("Found a crawler")
    }
}
```

## Contributing
The patterns and testcases are synced from the PHP repo.
If you find a bot/spider/crawler user agent that crawlerdetect fails to detect, please submit a pull request with the regex pattern and a testcase to the [upstream PHP repo](https://github.com/JayBizzle/Crawler-Detect).
