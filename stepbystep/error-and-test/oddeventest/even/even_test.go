package even

import (
	"testing"
)

func TestEven(t *testing.T)  {
	if !Even(10) {
		t.Log("10 is even")
		t.Fail()
	}

	if Even(9) {
		t.Log("9 is not even")
		t.Fail()
	}
}

func TestOdd(t *testing.T)  {
	if !Odd(9) {
		t.Log("9 is odd")
		t.Fail()
	}

	if Odd(10) {
		t.Log("10 is not odd")
		t.Fail()
	}
}