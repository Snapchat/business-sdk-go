package businesssdk

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"

	"github.com/umahmood/soundex"
)

/*
   How to Normalize Identifiers:
   - Normalize email addresses by trimming leading and trailing whitespace and converting
   all characters to lowercase before hashing
   - Normalize phone numbers by including the country code, remove any double 0 in front
   of the country code. if the number itself begins with a 0, this should be removed. Also
   exclude any non-numeric characters such as whitespace, parentheses, '+', or '-'
*/

const pattern1 = `^((\+|00)(\d+)[\-\s])?0?(.+)`
const pattern2 = `[^\d.]+`

var p1regex = regexp.MustCompile(pattern1)
var p2regex = regexp.MustCompile(pattern2)

func Sha256(in string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(in)))
}

func Soundex(in string) string {
	return soundex.Code(in)
}

// NormAndHashStr Returns a tuple with hashed and normalized string if it is valid.
// Nil otherwise and boolean to check if value is set
func NormAndHashStr(in string) (*string, bool) {
	if in == "" || strings.TrimSpace(in) == "" {
		return nil, false
	}
	strHash := Sha256(strings.ToLower(strings.TrimSpace(in)))
	return &strHash, true
}

func NormAndSoundexStr(in string) (*string, bool) {
	if in == "" || strings.TrimSpace(in) == "" {
		return nil, false
	}
	strSdx := Soundex(strings.ToLower(strings.TrimSpace(in)))
	return &strSdx, true
}

// NormAndHashPhoneNum Normalizes and hashes phone number. Returns a tuple with hashed
// normalized phone number if set. Nil otherwise and boolean to check if value is set
func NormAndHashPhoneNum(in string) (*string, bool) {
	num := normalizePhoneNum(in)
	if num == "" {
		return nil, false
	}
	numHash := Sha256(num)
	return &numHash, true
}

func normalizePhoneNum(in string) string {
	if in == "" {
		return in
	}
	results := p1regex.FindStringSubmatch(in)

	if len(results) > 0 {
		countryCode := results[3]
		numstr := results[4]

		if countryCode != "" {
			countryCode = p2regex.ReplaceAllString(countryCode, "")
		}

		if numstr != "" {
			numstr = p2regex.ReplaceAllString(numstr, "")
		}

		if countryCode == "" {
			countryCode = "1"
		}

		if numstr != "" {
			return countryCode + numstr
		}
	}

	return ""
}
