package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var outputPassengerADT = &Passenger{
	ObjectKey: "SH1",
	PTC: PassengerPTC{
		Quantity: "1",
		Value:    "ADT",
	},
}

var outputPassengerCHD = &Passenger{
	ObjectKey: "SH2",
	PTC: PassengerPTC{
		Quantity: "1",
		Value:    "CHD",
	},
}

var cases = []struct {
	name       string
	paxType    string
	passengers *PassengersRefs
	want       *Passenger
}{
	{
		"passed CHD pax type",
		"",
		MockPassengersRefs,
		nil,
	},
	{
		"correct arguments",
		"SH1",
		MockPassengersRefs,
		outputPassengerADT,
	},
}

func Test_GetPaxByID(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := getPaxByID(tt.paxType, tt.passengers)

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID(%v, %v) got %v, want %v", tt.paxType, tt.passengers, got, tt.want)
			}
		})
	}
}
