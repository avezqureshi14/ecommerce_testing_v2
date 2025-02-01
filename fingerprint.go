package main

import (
	"crypto/sha1"
	"encoding/hex"
	"regexp"
	"strings"
)

var numRe = regexp.MustCompile(`\d+`)

// fingerprint groups the same client error even when line numbers move.
func fingerprint(message, stack string) string {
	msg := strings.ToLower(strings.TrimSpace(message))
	msg = numRe.ReplaceAllString(msg, "#")
	top := ""
	for _, line := range strings.Split(stack, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		top = numRe.ReplaceAllString(strings.ToLower(line), "#")
		break
	}
	sum := sha1.Sum([]byte(msg + "|" + top))
	return hex.EncodeToString(sum[:8])
}
