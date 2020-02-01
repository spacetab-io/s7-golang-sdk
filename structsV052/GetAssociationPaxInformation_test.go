package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetAssociationPaxInformation = []struct {
	name    string
	paxType string
	want    *Associations
}{
	{
		"passed ADT pax type",
		"ADT",
		MockAssociation,
	},
	{
		"passed inexistent pax type",
		"CHD",
		nil,
	},
}

func Test_GetAssociationPaxInformation(t *testing.T) {
	for _, tt := range casesGetAssociationPaxInformation {
		t.Run(tt.name, func(t *testing.T) {
			got := MockOfferPrice.GetAssociationPaxInformation(tt.paxType, MocksAnonymousTravelerList)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID(%v) got %v, want %v", tt.paxType, got, tt.want)
			}
		})
	}
}
