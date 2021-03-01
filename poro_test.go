package go_poro

import (
	"fmt"
	league "github.com/mikey-wotton/go-league"
	"github.com/mikey-wotton/go-poro/testdata"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSummoner(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, testdata.ElBrayayin)
	}))
	defer ts.Close()

	p := New(Config{
		tagURL:	ts.URL,
	})

	summoner, err := p.GetSummoner("euw", "El Brayayin", "Katarina")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, summoner.Name, SummonerName("El Brayayin"))
	assert.Equal(t, summoner.Region, league.RegionEUW)
	assert.Equal(t, summoner.Champions[0].Name, ChampName("Katarina"))
	assert.Equal(t, len(summoner.Champions[0].Tags), 3)
	assert.Contains(t, summoner.Champions[0].Tags, &Tag{
		Name:        "High KDA",
		Description: "This player had a high kda",
		Colour:      TagColourGreen,
	})
	assert.Contains(t, summoner.Champions[0].Tags, &Tag{
		Name:        "Invader",
		Description: "This player often invades at the beginning of a game (100%).",
		Colour:      TagColourYellow,
	})
	assert.Contains(t, summoner.Champions[0].Tags, &Tag{
		Name:        "Katarina casual",
		Description: "This player has only played 1 games with Katarina this season.",
		Colour:      TagColourRed,
	})
}