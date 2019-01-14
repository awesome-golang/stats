package stats

import (
	"fmt"
	"testing"
)

func ExampleLinearRegression() {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
	}

	r, _ := LinearRegression(data)
	fmt.Println(r)
	// Output: [{1 2.400000000000001} {2 3.1} {3 3.7999999999999994}]
}

func TestLinearRegression(t *testing.T) {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := LinearRegression(data)
	a := 2.3800000000000026
	if r[0].Y != a {
		t.Errorf("%v != %v", r[0].Y, a)
	}
	a = 3.0800000000000014
	if r[1].Y != a {
		t.Errorf("%v != %v", r[1].Y, a)
	}
	a = 3.7800000000000002
	if r[2].Y != a {
		t.Errorf("%v != %v", r[2].Y, a)
	}
	a = 4.479999999999999
	if r[3].Y != a {
		t.Errorf("%v != %v", r[3].Y, a)
	}
	a = 5.179999999999998
	if r[4].Y != a {
		t.Errorf("%v != %v", r[4].Y, a)
	}

	_, err := LinearRegression([]Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestExponentialRegression(t *testing.T) {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := ExponentialRegression(data)
	a, _ := Round(r[0].Y, 3)
	if a != 2.515 {
		t.Errorf("%v != %v", r[0].Y, 2.515)
	}
	a, _ = Round(r[1].Y, 3)
	if a != 3.032 {
		t.Errorf("%v != %v", r[1].Y, 3.032)
	}
	a, _ = Round(r[2].Y, 3)
	if a != 3.655 {
		t.Errorf("%v != %v", r[2].Y, 3.655)
	}
	a, _ = Round(r[3].Y, 3)
	if a != 4.407 {
		t.Errorf("%v != %v", r[3].Y, 4.407)
	}
	a, _ = Round(r[4].Y, 3)
	if a != 5.313 {
		t.Errorf("%v != %v", r[4].Y, 5.313)
	}

	_, err := ExponentialRegression([]Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestExponentialRegressionYCoordErr(t *testing.T) {
	c := []Coordinate{{1, -5}, {4, 25}, {6, 5}}
	_, err := ExponentialRegression(c)
	if err != YCoordErr {
		t.Errorf(err.Error())
	}
}

func TestLogarithmicRegression(t *testing.T) {
	data := []Coordinate{
		{1, 2.3},
		{2, 3.3},
		{3, 3.7},
		{4, 4.3},
		{5, 5.3},
	}

	r, _ := LogarithmicRegression(data)
	a := 2.1520822363811702
	if r[0].Y != a {
		t.Errorf("%v != %v", r[0].Y, a)
	}
	a = 3.3305559222492214
	if r[1].Y != a {
		t.Errorf("%v != %v", r[1].Y, a)
	}
	a = 4.019918836568674
	if r[2].Y != a {
		t.Errorf("%v != %v", r[2].Y, a)
	}
	a = 4.509029608117273
	if r[3].Y != a {
		t.Errorf("%v != %v", r[3].Y, a)
	}
	a = 4.888413396683663
	if r[4].Y != a {
		t.Errorf("%v != %v", r[4].Y, a)
	}

	_, err := LogarithmicRegression([]Coordinate{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
