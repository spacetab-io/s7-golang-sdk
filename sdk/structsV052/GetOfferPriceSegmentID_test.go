package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var segID = "SEG1"

var casesGetOfferPriceSegmentID = []struct {
	name string
	want *string
}{
	{
		"common case",
		&segID,
	},
}

func Test_GetOfferPriceSegmentID(t *testing.T) {
	for _, tt := range casesGetOfferPriceSegmentID {
		t.Run(tt.name, func(t *testing.T) {
			got := MockOfferPrice.GetOfferPriceSegmentID()

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID() got %v", got)
			}
		})
	}
}
