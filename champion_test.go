package go_poro

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChampName_ToURI(t *testing.T) {
	tests := map[string]struct{
		champName string

		expValid bool
		expErr error
	}{
		"Name with multiple spaces ok: Nunu and Willump to nunuandWillump": {
			champName: "Aurelion Sol",
			expValid: true,
			expErr: nil,
		},
		"Name with a ' ok: Kai'Sa to kaisa": {
			champName: "Kai'Sa",
			expValid: true,
			expErr: nil,
		},
		"Name with a space ok: Aurelion Sol to aureliansol": {
			champName: "aurelion Sol",
			expValid: true,
			expErr: nil,
		},
		"Simplest: Aatrox to aatrox": {
			champName: "Aatrox",
			expValid: true,
			expErr: nil,
		},
		"Unknown champ name": {
			champName: "bobby",
			expValid: false,
			expErr: nil,
		},
		"Mispelt champ name Zoi is false": {
			champName: "zoi",
			expValid: false,
			expErr: nil,
		},
		"Empty champName returns false without error": {
			champName: "",
			expValid: false,
			expErr: nil,
		},
	}


	for desc, test := range tests {
		t.Run(desc, func(t *testing.T) {
			valid, err := ChampName(test.champName).Valid()

			assert.Equal(t, test.expValid, valid)
			assert.Equal(t, test.expErr, err)
		})
	}
}