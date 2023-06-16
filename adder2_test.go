package adder31plus

import (
	"testing"
)

type DataSet struct {
	val1 string
	val2 string
	ans string
}

func TestAdder(t *testing.T) {
	TestData := []DataSet{
		{"123", "456", "579"},
		{"-123", "-456", "-579"},
		{"999", "1", "1000"},
		{"-999", "-1", "-1000"},
		{"1", "999", "1000"},
		{"-1", "-999", "-1000"},
		{"33", "-44", "-11"},   
		{"-44", "33", "-11"},
		{"0", "-9", "-9"},
		{"-9", "0", "-9"},
		{"9", "-10", "-1"},
		{"-10", "9", "-1"},
		{"-98", "9", "-89"},
		{"9", "-98", "-89"},
		{"123", "-456", "-333"},
		{"-456", "123", "-333"},
		{"123", "-4567", "-4444"},
		{"-4567", "123", "-4444"},
		{"999", "-456", "543"},
		{"-456", "999", "543"},
		{"999999", "-1000000", "-1"},
		{"-1000000", "999999", "-1"},
		{"01","02", "3"},
		{"-001", "02", "1"},
		{"+1","-1","0"},
		{"0", "-0", "0"},
		{"-0", "0", "0"},
	}

	for _, data := range TestData {
		ans := Adder(data.val1, data.val2)
		expected := data.ans
		if ans != expected {
			t.Fatalf("Expected %s, got %s", expected, ans)
		}
	}
}