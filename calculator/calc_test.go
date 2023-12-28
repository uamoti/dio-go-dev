package main

import "testing"

func TestPlusInt(t *testing.T) {
	result := Plus(1, 2, 3)
	if result != 6 {
		t.Error("Expected 6, got", result)
	}
}

func TestPlusFloat(t *testing.T) {
	result := Plus(2.5, 2.5)
	if result != 5.0 {
		t.Error("Expected 5.0, got", result)
	}
}

func TestPlusNeg(t *testing.T) {
	result := Plus(-2, 2, 3, -1)
	if result != 2 {
		t.Error("Expected 2, got", result)
	}
}

func TestMinusInt(t *testing.T) {
	result := Minus(1, 2, 3)
	if result != -4 {
		t.Error("Expected -4, got", result)
	}
}

func TestMinusFloat(t *testing.T) {
	result := Minus(3.5, 1, 0.5)
	if result != 2.0 {
		t.Error("Expected 2.0, got", result)
	}
}

func TestMinusNeg(t *testing.T) {
	result := Minus(-1, -2, -3)
	if result != 4 {
		t.Error("Expected 4, got", result)
	}
}

func TestMultInt(t *testing.T) {
	result := Mult(1, 2, 3)
	if result != 6 {
		t.Error("Expected 6, got", result)
	}
}

func TestMultFloat(t *testing.T) {
	result := Mult(2, 3.0, 0.5)
	if result != 3.0 {
		t.Error("Expected 3.0, got", result)
	}
}

func TestMultNeg(t *testing.T) {
	result := Mult(1, 2, -3)
	if result != -6 {
		t.Error("Expected -6, got", result)
	}
}

func TestDivInt(t *testing.T) {
	result := Div(12, 2, 3)
	if result != 2 {
		t.Error("Expected 2, got", result)
	}
}

func TestDivFloat(t *testing.T) {
	result := Div(18.0, 3, 0.5)
	if result != 12.0 {
		t.Error("Expected 12.0, got", result)
	}
}

func TestDivNeg(t *testing.T) {
	result := Div(18, -2, -3)
	if result != 3 {
		t.Error("Expected 3, got", result)
	}
}
