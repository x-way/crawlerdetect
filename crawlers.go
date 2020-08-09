package crawlerdetect

import (
	"regexp"
	"strings"
)

var crawlersRegex *regexp.Regexp

func init() {
	crawlersRegex = regexp.MustCompile("(?i)(" + strings.Join(crawlers, "|") + ")")
}
