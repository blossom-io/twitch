package link

import "regexp"

const (
	RegexpLink = `\bhttps?://\S+\b` // http or https
)

func ExtractLink(msg string) (link string, found bool) {
	re := regexp.MustCompile(RegexpLink)

	URLs := re.FindAllString(msg, 1)

	if len(URLs) >= 1 {
		return URLs[0], true
	}

	return "", false
}
