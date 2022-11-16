// Package crawlerdetect provides a Go version of the https://github.com/JayBizzle/Crawler-Detect PHP library
// It can be used to detect crawlers based on their HTTP User Agent header
//
// Here is a simple example
//
//	uastring := "curl/7.54.0"
//	if crawlerdetect.IsCrawler(uastring) {
//	        fmt.Println("Found a crawler")
//	}
//
// To use custom patterns a dedicated instance can be created:
//
//	uastring := "curl/7.54.0"
//	crawlerDetect := crawlerDetect.New()
//	crawlerDetect.SetCrawlers([]string{`curl`})
//	if crawlerDetect.IsCrawler(uastring) {
//	        fmt.Println("Found a crawler")
//	}
package crawlerdetect

import (
	"regexp"
	"strings"
	"sync"
)

var defaultCrawlerDetect *CrawlerDetect

func init() {
	defaultCrawlerDetect = New()
}

// IsCrawler checks for a user agent string if it is a crawler
func IsCrawler(input string) bool {
	return defaultCrawlerDetect.IsCrawler(input)
}

// CrawlerDetect contains the patterns and exclusions for detecting crawlers
type CrawlerDetect struct {
	crawlers        []string
	crawlersRegex   *regexp.Regexp
	exclusions      []string
	exclusionsRegex *regexp.Regexp
	m               sync.RWMutex
}

// New creates a new CrawlerDetect instance with the default patterns
func New() *CrawlerDetect {
	c := new(CrawlerDetect)
	c.SetCrawlers(crawlers)
	c.SetExclusions(exclusions)
	return c
}

// SetCrawlers sets a custom list of crawler patterns
func (c *CrawlerDetect) SetCrawlers(crawlers []string) {
	c.m.Lock()
	c.crawlers = crawlers
	c.compileCrawlers()
	c.m.Unlock()
}

// SetExclusions sets a custom list of exclusion patterns
func (c *CrawlerDetect) SetExclusions(exclusions []string) {
	c.m.Lock()
	c.exclusions = exclusions
	c.compileExclusions()
	c.m.Unlock()
}

func (c *CrawlerDetect) compileCrawlers() {
	c.crawlersRegex = regexp.MustCompile("(?i)(" + strings.Join(c.crawlers, "|") + ")")
}

func (c *CrawlerDetect) compileExclusions() {
	c.exclusionsRegex = regexp.MustCompile("(?i)(" + strings.Join(c.exclusions, "|") + ")")
}

func (c *CrawlerDetect) excludeString(input string) string {
	return c.exclusionsRegex.ReplaceAllString(strings.TrimSpace(input), "")
}

// IsCrawler checks for a user agent string if it is a crawler
func (c *CrawlerDetect) IsCrawler(input string) bool {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.crawlersRegex.MatchString(c.excludeString(input))
}
