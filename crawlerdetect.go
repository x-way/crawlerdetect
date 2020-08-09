// Package crawlerdetect provides a Go version of the https://github.com/JayBizzle/Crawler-Detect PHP library
// It can be used to detect crawlers based on their HTTP User Agent header
//
// Here is a simple example
//
//     uastring := "curl/7.54.0"
//     if crawlerdetect.IsCrawler(uastring) {
//             fmt.Println("Found a crawler")
//     }
//
package crawlerdetect

import "strings"

func excludeString(input string) string {
	return exclusionsRegex.ReplaceAllString(strings.TrimSpace(input), "")
}

// IsCrawler checks for a user agent string if it is a crawler
func IsCrawler(input string) bool {
	return crawlersRegex.MatchString(excludeString(input))
}
