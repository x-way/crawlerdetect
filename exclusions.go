package crawlerdetect

import (
	"regexp"
	"strings"
)

var exclusionsRegex *regexp.Regexp

func init() {
	exclusionsRegex = regexp.MustCompile("(?i)(" + strings.Join(exclusions, "|") + ")")
}
