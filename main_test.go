package main

import (
	"reflect"
	"testing"
)

// TestGreet tests that our greeting matches the person's name.
func TestGreet(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			arg:  "Brian",
			want: "Hello World, Brian!",
		}, {
			arg:  "Mary",
			want: "Hello World, Mary!",
		},
	}
	for _, test := range tests {
		if got := Greet(test.arg); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Test failed: got %+v, want %+v", got, test.want)
		}
	}
}
