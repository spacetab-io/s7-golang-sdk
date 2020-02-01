package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetTaxFromPrice = []struct {
	name     string
	paxType  string
	traveler AnonymousTravelerList
	want     *float64
}{
	{
		"passed ADT pax type",
		"ADT",
		MocksAnonymousTravelerList,
		&MockADTTax,
	},
	{
		"passed CHD pax type",
		"CHD",
		MocksAnonymousTravelerList,
		&MockCHDTax,
	},
	{
		"passed inexisted pax type",
		"undefined",
		MocksAnonymousTravelerList,
		nil,
	},
}

func Test_GetTaxFromPrice(t *testing.T) {
	for _, tt := range casesGetTaxFromPrice {
		t.Run(tt.name, func(t *testing.T) {
			got := MockAirlineOffer.GetTaxFromPrice(tt.traveler, tt.paxType)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID() got %v, want %v", got, tt.want)
			}
		})
	}
}
