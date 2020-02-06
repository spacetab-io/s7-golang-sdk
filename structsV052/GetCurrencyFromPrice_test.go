package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var curr = "RUB"

var casesCurrencyFromPrice = []struct {
	name string
	want *string
}{
	{
		"passed ADT pax type",
		&curr,
	},
}

func Test_CurrencyFromPrice(t *testing.T) {
	for _, tt := range casesCurrencyFromPrice {
		t.Run(tt.name, func(t *testing.T) {
			got := MockAirlineOffer.CurrencyFromPrice()

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID() got %v, want %v", got, tt.want)
			}
		})
	}
}
