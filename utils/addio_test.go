package utils

import "testing"

func Test1(t *testing.T) {
	expected := int32(10)
	actual := Adder(1, 2, 3, 4)

	if expected != actual {
		t.Error("Test 1 Failed")
	} else {
		t.Log("Passed 1")
	}
}

func Test2(t *testing.T) {
	expected := int32(11)
	actual := Adder(1, 2, 3, 4, 1)

	if expected != actual {
		t.Error("Test 2 Failed")
	} else {
		t.Log("Passed 2")
	}
}

func Test3(t *testing.T) {
	expected := int32(6)
	actual := Adder(1, 2, 3, 5, -4)

	if expected != actual {
		t.Error("Test 3 Failed")
	} else {
		t.Log("Passed 3")
	}
}
