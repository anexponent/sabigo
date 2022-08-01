package utils

import (
	"reflect"
	"testing"
)

func TestRandomCharsGen(t *testing.T) {

	got, _ := GenerateToken(64)
	t.Errorf("got %q", reflect.ValueOf(got))
	if len(got) != 64 {
		t.Errorf("got %q", got)
	}
}
