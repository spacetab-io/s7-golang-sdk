package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetPriceByPax = []struct {
	name    string
	paxType string
	want    *OfferPrice
}{
	{
		"correct params",
		"ADT",
		MockOfferPrice,
	},
}

func Test_GetPriceByPax(t *testing.T) {
	for _, tt := range casesGetPriceByPax {
		t.Run(tt.name, func(t *testing.T) {
			got := MockResponse.GetPriceByPax(tt.paxType)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID(%v,) got %v, want %v", tt.paxType, got, tt.want)
			}
		})
	}
}
