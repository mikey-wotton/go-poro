package go_poro

import (
	"fmt"
	league "github.com/mikey-wotton/go-league"
	"golang.org/x/net/html"
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	tagName        = "tagName"
	tagDescription = "tagDescription"
)

const (
	base         = "https://www.leagueofgraphs.com"
	championPath = "/summoner/champions"
	tagURL       = base + championPath
)

const (
	mozillaUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36"
)

func GetSummoner(region league.Region, summoner SummonerName, champion ChampName) (*Summoner, error) {
	if !region.Valid() {
		return nil, fmt.Errorf("unknown region '%s' provided", region)
	}

	valid, err := champion.Valid()
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("unknown champion '%s' provided", champion)
	}

	championURI, err := champion.ToURI()
	if err != nil {
		return nil, err
	}

	url := strings.ToLower(fmt.Sprintf(tagURL+"/%s/%s/%s", championURI, region, summoner.ToURI()))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent",mozillaUserAgent )

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	z := html.NewTokenizer(response.Body)

	tags := make([]*Tag, 0)
	end := false
	for !end {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			end = true
		default:
			t := z.Token()
			if t.Data == "div" {
				for _, a := range t.Attr {
					if a.Key == "class" {
						if strings.Contains(a.Val, "tag requireTooltip") {
							tag, err := getTag(t.Attr)
							if err != nil {
								return nil, err
							}

							tags = append(tags, tag)
						}
						break
					}
				}
			}
		}
	}

	return &Summoner{
		Name:   summoner,
		Region: region,
		Champions: []*Champion{
			{
				Name: champion,
				Tags: tags,
			},
		},
	}, nil
}

func getTag(attributes []html.Attribute) (*Tag, error) {
	var toolTip html.Attribute
	for _, attrib := range attributes {
		if attrib.Key == "tooltip" {
			toolTip = attrib
			break
		}
	}

	if toolTip == (html.Attribute{}) {
		return nil, fmt.Errorf("could not find div attribute tooltip")
	}

	name, description, err := getTagNameAndDesc(toolTip.Val)
	if err != nil {
		return nil, err
	}

	colour, err := getTagColour(toolTip.Val)
	if err != nil {
		return nil, err
	}

	return &Tag{
		Name:        name,
		Description: description,
		Colour:      colour,
	}, nil
}

func getTagNameAndDesc(tooltip string) (string, string, error) {
	r := regexp.MustCompile(`<itemname class='.*?'>(?P<` + tagName + `>.+?)<.*?<span class='tagDescription'>(?P<` + tagDescription + `>.+)`)
	matches := r.FindStringSubmatch(tooltip)

	paramsMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(matches) {
			paramsMap[name] = stripHtmlTags(matches[i])
		}
	}

	if len(paramsMap) != 2 {
		return "", "", fmt.Errorf("expected params map len 2 but got %d", len(paramsMap))
	}

	return paramsMap[tagName], paramsMap[tagDescription], nil
}

func getTagColour(tooltip string) (TagColour, error) {
	switch {
	case strings.Contains(tooltip, string(TagColourGreen)):
		return TagColourGreen, nil
	case strings.Contains(tooltip, string(TagColourYellow)):
		return TagColourYellow, nil
	case strings.Contains(tooltip, string(TagColourRed)):
		return TagColourRed, nil
	default:
		return TagColourUnknown, fmt.Errorf("could not get tag colour from tooltip: %s", tooltip)
	}
}

const (
	htmlTagStart = 60 // Unicode `<`
	htmlTagEnd   = 62 // Unicode `>`
)

// Aggressively strips HTML tags from a string.
// It will only keep anything between `>` and `<`.
func stripHtmlTags(s string) string {
	// Setup a string builder and allocate enough memory for the new string.
	var builder strings.Builder
	builder.Grow(len(s) + utf8.UTFMax)

	in := false // True if we are inside an HTML tag.
	start := 0  // The index of the previous start tag character `<`
	end := 0    // The index of the previous end tag character `>`

	for i, c := range s {
		// If this is the last character and we are not in an HTML tag, save it.
		if (i+1) == len(s) && end >= start {
			builder.WriteString(s[end:])
		}

		// Keep going if the character is not `<` or `>`
		if c != htmlTagStart && c != htmlTagEnd {
			continue
		}

		if c == htmlTagStart {
			// Only update the start if we are not in a tag.
			// This make sure we strip out `<<br>` not just `<br>`
			if !in {
				start = i
			}
			in = true

			// Write the valid string between the close and start of the two tags.
			builder.WriteString(s[end:start])
			continue
		}
		// else c == htmlTagEnd
		in = false
		end = i + 1
	}
	s = builder.String()

	s = strings.ReplaceAll(s, "</span>", "") //poro website has bad formatting and additional span close occasionally
	return s
}
