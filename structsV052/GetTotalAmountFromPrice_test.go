package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesTotalAmountFromPrice = []struct {
	name     string
	paxType  string
	traveler AnonymousTravelerList
	want     *float64
}{
	{
		"passed ADT pax type",
		"ADT",
		MocksAnonymousTravelerList,
		&MockADTTotalAmount,
	},
	{
		"passed CHD pax type",
		"CHD",
		MocksAnonymousTravelerList,
		&MockCHDTotalAmount,
	},
	{
		"passed inexisted pax type",
		"undefined",
		MocksAnonymousTravelerList,
		nil,
	},
}

func Test_TotalAmountFromPrice(t *testing.T) {
	for _, tt := range casesTotalAmountFromPrice {
		t.Run(tt.name, func(t *testing.T) {
			got := MockAirlineOffer.TotalAmountFromPrice(tt.traveler, tt.paxType)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID() got %v, want %v", got, tt.want)
			}
		})
	}
}
