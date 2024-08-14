package base

import (
	"math/big"
	"testing"
)

func TestNewInteger(t *testing.T) {
	num := NewInteger(42)
	if !num.IsInteger() {
		t.Errorf("Expected number to be an integer")
	}
	if num.intValue.Cmp(big.NewInt(42)) != 0 {
		t.Errorf("Expected value 42, got %v", num.intValue)
	}
}

func TestNewFloat(t *testing.T) {
	num := NewFloat(42.42)
	if num.IsInteger() {
		t.Errorf("Expected number to be a float")
	}
	if num.floatValue.Cmp(big.NewFloat(42.42)) != 0 {
		t.Errorf("Expected value 42.42, got %v", num.floatValue)
	}
}

func TestCompare(t *testing.T) {
	intNum1 := NewInteger(42)
	intNum2 := NewInteger(43)
	floatNum1 := NewFloat(42.42)
	floatNum2 := NewFloat(43.43)

	if intNum1.Compare(intNum2) != -1 {
		t.Errorf("Expected -1, got %d", intNum1.Compare(intNum2))
	}
	if intNum2.Compare(intNum1) != 1 {
		t.Errorf("Expected 1, got %d", intNum2.Compare(intNum1))
	}
	if intNum1.Compare(NewInteger(42)) != 0 {
		t.Errorf("Expected 0, got %d", intNum1.Compare(NewInteger(42)))
	}
	if floatNum1.Compare(floatNum2) != -1 {
		t.Errorf("Expected -1, got %d", floatNum1.Compare(floatNum2))
	}
	if floatNum2.Compare(floatNum1) != 1 {
		t.Errorf("Expected 1, got %d", floatNum2.Compare(floatNum1))
	}
	if floatNum1.Compare(NewFloat(42.42)) != 0 {
		t.Errorf("Expected 0, got %d", floatNum1.Compare(NewFloat(42.42)))
	}
}

func TestEquals(t *testing.T) {
	intNum1 := NewInteger(42)
	intNum2 := NewInteger(42)
	floatNum1 := NewFloat(42.42)
	floatNum2 := NewFloat(42.42)

	if !intNum1.Equals(intNum2) {
		t.Errorf("Expected true, got false")
	}
	if !floatNum1.Equals(floatNum2) {
		t.Errorf("Expected true, got false")
	}
	if intNum1.Equals(floatNum1) {
		t.Errorf("Expected false, got true")
	}
}

func TestHash(t *testing.T) {
	intNum := NewInteger(42)
	floatNum := NewFloat(42.42)

	if intNum.Hash() == 0 {
		t.Errorf("Expected non-zero hash, got 0")
	}
	if floatNum.Hash() == 0 {
		t.Errorf("Expected non-zero hash, got 0")
	}
}

func TestString(t *testing.T) {
	intNum := NewInteger(42)
	floatNum := NewFloat(42.42)

	if intNum.String() != "42" {
		t.Errorf("Expected '42', got %s", intNum.String())
	}
	if floatNum.String() != "42.42" {
		t.Errorf("Expected '42.42', got %s", floatNum.String())
	}
}

func TestSize(t *testing.T) {
	intNum := NewInteger(42)
	floatNum := NewFloat(42.42)

	if intNum.Size() != len(intNum.intValue.Bytes()) {
		t.Errorf("Expected size %d, got %d", len(intNum.intValue.Bytes()), intNum.Size())
	}
	if floatNum.Size() != len(floatNum.floatValue.Text('g', -1)) {
		t.Errorf("Expected size %d, got %d", len(floatNum.floatValue.Text('g', -1)), floatNum.Size())
	}
}

func TestToAtom(t *testing.T) {
	intNum := NewInteger(42)
	atom := intNum.ToAtom()

	if atom.value != int64(42) {
		t.Errorf("Expected atom value 42, got %v", atom.value)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got none")
		}
	}()
	floatNum := NewFloat(42.42)
	floatNum.ToAtom()
}
