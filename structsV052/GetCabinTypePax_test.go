package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetCabinTypePAX = []struct {
	name    string
	paxType string
	want    *CabinFlightSegment
}{
	{
		"passed ADT pax type",
		"ADT",
		&MockCabinFlightSegment,
	},
	{
		"passed inexistent pax type",
		"CHD",
		nil,
	},
}

func Test_GetCabinTypePAX(t *testing.T) {
	for _, tt := range casesGetCabinTypePAX {
		t.Run(tt.name, func(t *testing.T) {
			got := MockOfferPrice.GetCabinTypePax(tt.paxType, MocksAnonymousTravelerList)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID(%v) got %v, want %v", tt.paxType, got, tt.want)
			}
		})
	}
}
