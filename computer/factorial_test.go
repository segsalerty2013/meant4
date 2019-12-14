package computer

import "testing"

func TestFactorial(t *testing.T) {
	t.Log("Testing")

	// test that factorial of 0 is 1
	if calculate(0) != 1 {
		t.Error("Test Fail calculating factorial of 0 to be 1")
	}

	// test that factorial of 10 is 3628800
	if calculate(10) != 3628800 {
		t.Error("Test Fail calculating factorial of 10 to be 3628800")
	}
	t.Log("Test complete")
}
