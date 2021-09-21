package structs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{

		{rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTests {

		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf(" %#v got %g want %g", tt.shape, got, tt.want)

		}
	}

}
