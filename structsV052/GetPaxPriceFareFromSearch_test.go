package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetPaxPriceFareFromSearch = []struct {
	PaxType string
	name    string
	want    float64
}{
	{
		"ADT",
		"passed ADT pax type",
		MockADTFare,
	},
	{
		"CHD",
		"passed CHD pax type",
		MockCHDFare,
	},
}

func Test_GetPaxPriceFareFromSearch(t *testing.T) {
	for _, tt := range casesGetPaxPriceFareFromSearch {
		t.Run(tt.name, func(t *testing.T) {
			got := MockAirlineOffer.GetPaxPriceFareFromSearch(MocksAnonymousTravelerList, tt.PaxType)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID() got %v, want %v", got, tt.want)
			}
		})
	}
}
