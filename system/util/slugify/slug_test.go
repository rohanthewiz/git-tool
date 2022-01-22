package slugify

import (
	"testing"
)

const expectedStringReturn = "Return string is not slugified as expected"

func TestSlugify(t *testing.T) {
	s := "test->àèâ<-test"
	slug := Marshal(s)
	expected := "test-aea-test"
	if slug != expected {
		t.Fatal(expectedStringReturn, expected, slug)
	}
}

func TestLowerOption(t *testing.T) {
	s := "Test->àèâ<-Test"
	slug := Marshal(s, true)
	expected := "test-aea-test"
	if slug != expected {
		t.Error(expectedStringReturn, expected, slug)
	}
	slug = Marshal(s, false)
	expected = "Test-aea-Test"
	if slug != expected {
		t.Error(expectedStringReturn, expected, slug)
	}
	slug = Marshal(s)
	expected = "Test-aea-Test"
	if slug != expected {
		t.Error(expectedStringReturn, expected, slug)
	}
}
