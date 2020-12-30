package builder

import "testing"

func TestBuilder(t *testing.T) {
	thin := Thin{}
	fat := Fat{}

	director := Director{&thin}
	director.CreatePerson()

	directorf := Director{&fat}
	directorf.CreatePerson()
}
