package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var SH1ID = "SH1"
var SH2ID = "SH2"

var casesAirShoppingGetPaxByID = []struct {
	name       string
	paxID      string
	passengers *AnonymousTravelerList
	want       *Passenger
}{
	{
		"passed nonexistent SH2 pax ID",
		SH2ID,
		&MocksAnonymousTravelerList,
		outputPassengerCHD,
	},
	{
		"correct arguments",
		SH1ID,
		&MocksAnonymousTravelerList,
		outputPassengerADT,
	},
}

func Test_AirShoppingGetPaxByID(t *testing.T) {
	for _, tt := range casesAirShoppingGetPaxByID {
		t.Run(tt.name, func(t *testing.T) {

			got := airShoppingGetPaxByID(tt.paxID, MocksAnonymousTravelerList)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID(%v, %v) got %v, want %v", tt.paxID, tt.passengers, got, tt.want)
			}
		})
	}
}
