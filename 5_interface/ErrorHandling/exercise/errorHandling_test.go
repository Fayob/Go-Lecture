package main

import "testing"

func TetsParseTime(t *testing.T)  {
	table := []struct {
		time string
		ok bool
	}{
		{"19:00:12", true},
		{"aa:bb:cc", false},
		{"0:59:59", true},
		{"", false},
		{"bad", false},
		{"5:23:", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}