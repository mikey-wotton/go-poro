package go_poro

import (
	league "github.com/mikey-wotton/go-league"
	"strings"
)

type Summoner struct {
	Name      SummonerName
	Region    league.Region
	Champions []*Champion
}

type SummonerName string

func (s SummonerName) ToFriendly() string {
	return strings.ReplaceAll(string(s), "%20", " ")
}

func (s SummonerName) ToURI() string {
	return strings.ReplaceAll(string(s), " ", "%20")
}