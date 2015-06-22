// Package gist6003701 implements functions for converting between camelCase and underscore_separated forms for identifier names.
package gist6003701

import (
	"strings"
	"unicode"
)

// UnderscoreSepToCamelCase converts "string_URL_append" to "StringUrlAppend" form.
func UnderscoreSepToCamelCase(s string) string {
	return strings.Replace(strings.Title(strings.Replace(strings.ToLower(s), "_", " ", -1)), " ", "", -1)
}

func addSegment(inout, seg []rune) []rune {
	if len(seg) == 0 {
		return inout
	}
	if len(inout) != 0 {
		inout = append(inout, '_')
	}
	initialism := strings.ToUpper(string(seg))
	if _, ok := initialisms[initialism]; ok {
		inout = append(inout, []rune(initialism)...)
	} else {
		inout = append(inout, seg...)
	}
	return inout
}

// CamelCaseToUnderscoreSep converts "StringUrlAppend" to "string_url_append" form.
func CamelCaseToUnderscoreSep(s string) string {
	var out []rune
	var seg []rune
	for _, r := range s {
		if !unicode.IsLower(r) {
			out = addSegment(out, seg)
			seg = nil
		}
		seg = append(seg, unicode.ToLower(r))
	}
	out = addSegment(out, seg)
	return string(out)
}

// UnderscoreSepToMixedCaps converts "string_URL_append" to "StringURLAppend" form.
func UnderscoreSepToMixedCaps(in string) string {
	var out string
	ss := strings.Split(in, "_")
	for _, s := range ss {
		initialism := strings.ToUpper(s)
		if _, ok := initialisms[initialism]; ok {
			out += initialism
		} else {
			out += strings.Title(s)
		}
	}
	return out
}

// initialisms is the set of initialisms in Go-style Mixed Caps case.
var initialisms = map[string]struct{}{
	"API":   {},
	"ASCII": {},
	"CPU":   {},
	"CSS":   {},
	"DNS":   {},
	"EOF":   {},
	"GUID":  {},
	"HTML":  {},
	"HTTP":  {},
	"HTTPS": {},
	"ID":    {},
	"IP":    {},
	"JSON":  {},
	"LHS":   {},
	"QPS":   {},
	"RAM":   {},
	"RHS":   {},
	"RPC":   {},
	"SLA":   {},
	"SMTP":  {},
	"SQL":   {},
	"SSH":   {},
	"TCP":   {},
	"TLS":   {},
	"TTL":   {},
	"UDP":   {},
	"UI":    {},
	"UID":   {},
	"UUID":  {},
	"URI":   {},
	"URL":   {},
	"UTF8":  {},
	"VM":    {},
	"XML":   {},
	"XSRF":  {},
	"XSS":   {},
}
