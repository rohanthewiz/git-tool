package util

import (
	"crypto/sha1"
	"fmt"
	"git-tool/system/util/slugify"
	"strconv"
	"strings"
	"time"
)

// Split given string and clean up any leading or trailing whitespace
// Don't include empty tokens
func StrSplitAndTrim(instr, separator string) (out []string) {
	arr := strings.Split(instr, separator)
	for _, item := range arr {
		if trimmed := strings.TrimSpace(item); trimmed != "" { // don't add an empty item
			out = append(out, trimmed)
		}
	}
	return
}

func SlugWithRandomString(title string) string {
	slug := slugify.Marshal(strings.ToLower(title))
	timestr := time.Now().Format("2006-0102")
	return slug + "-" + timestr + "-" + Sha1WithUnixTimeShort([]byte(slug))
}

func Slugify(instr string) string {
	return slugify.Marshal(strings.ToLower(instr))
}

func Sha1WithUnixTimeShort(data []byte) string {
	str := fmt.Sprintf("%d.%s", time.Now().UnixNano(), string(data))
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))[4:16]
}

func ParseStrToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return float64(0)
	}
	return f
}

// Effectively "round" a float
func FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

// Join words consistently with only one space between
func JoinWords(words ...string) (out string) {
	for _, word := range words {
		if out != "" {
			out += " "
		}
		out += strings.TrimSpace(word)
	}
	return
}

func TruncateString(str string, limit int, withEllipses bool) string {
	if len(str) > limit {
		str = str[:limit]
		if withEllipses {
			str += "..."
		}
	}
	return str
}

func StringtoBool(str string) (value bool) {
	if strings.ToLower(str) == "true" {
		value = true
	} else {
		value = false
	}
	return
}

func BooltoString(value bool) (str string) {
	if value == false {
		str = "false"
	} else {
		str = "true"
	}
	return
}
