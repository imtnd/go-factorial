package factorial_test

import (
	"github.com/imtnd/go-factorial/pkg/factorial"
	"testing"
)

func TestFactorialOf10is3628800(t *testing.T) {
	actual := factorial.GetFactorial(10)
	expect := 3628800
	if actual != expect {
		t.Errorf("Factorial of 10 is %#v. but actual value is %#v", expect, actual)
	}
}

func TestFactorialOf0is1(t *testing.T) {
	actual := factorial.GetFactorial(0)
	expect := 1
	if actual != expect {
		t.Errorf("Factorial of 0 is %#v. but actual value is %#v", expect, actual)
	}
}
