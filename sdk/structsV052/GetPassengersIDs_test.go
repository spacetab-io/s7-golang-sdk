package structsV052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesGetPassengersIDs = []struct {
	name string
	want []string
}{
	{
		"passed CHD pax type",
		MockPassengerKeys,
	},
}

func Test_GetPassengersIDs(t *testing.T) {
	for _, tt := range casesGetPassengersIDs {
		t.Run(tt.name, func(t *testing.T) {
			got := MockOrderViewResponse.GetPassengersIDs()

			if !assert.Equal(t, got, tt.want, "should be equal") {
				t.Errorf("getPaxByID() got %v, want %v", got, tt.want)
			}
		})
	}
}
