package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetClassTypePAX = []struct {
	name    string
	paxType string
	want    *ClassOfService
}{
	{
		"passed ADT pax type",
		"ADT",
		&MockClassOfService,
	},
	{
		"passed inexistent pax type",
		"CHD",
		nil,
	},
}

func Test_GetClassTypePAX(t *testing.T) {
	for _, tt := range casesGetClassTypePAX {
		t.Run(tt.name, func(t *testing.T) {
			got := MockOfferPrice.GetClassTypePax(tt.paxType, MocksAnonymousTravelerList)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID(%v) got %v, want %v", tt.paxType, got, tt.want)
			}
		})
	}
}
