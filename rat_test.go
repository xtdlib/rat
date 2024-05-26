package rat

import (
	"testing"
)

func TestBasics(t *testing.T) {
	a := String("2")
	b := String("4")
	if Add(a, b).String() != "6" {
		t.Fatal()
	}

	if Minus(a, b).String() != "-2" {
		t.Fatal()
	}

	if Minus(a, b).Add(b).String() != "2" {
		t.Fatal()
	}

}

func TestNeg(t *testing.T) {
	a := String("2")
	b := String("4")
	a.Minus(b)

}

func TestQuo(t *testing.T) {
	a := String("2")
	b := String("4")
	if Quo(a, b).String() != "0.5" {
		t.Fatal("Quo")
	}
}

func TestPowInt(t *testing.T) {
	a := String("2")
	exp := a.PowInt(10).String()
	if exp != "1024" {
		t.Fatal(exp)
	}

	b := String("1/2")
	if a.Add(b).String() != "2.5" {
		t.Fatal()
	}
}

func TestString(t *testing.T) {
	{
		a := String("0.5/1")
		if a.String() != "0.5" {
			t.Fatal(a.String())
		}
	}

	{
		a := String("1/0.5")
		if a.String() != "2" {
			t.Fatal(a.String())
		}
	}
}

func TestCopy(t *testing.T) {
	a := String("1/2")
	b := a.Copy()
	if a.Add(String("1")).String() != "1.5" {
		t.Fatal()
	}
	if b.String() != "0.5" {
		t.Fatal()
	}
}

func TestZero(t *testing.T) {
	z := Zero()
	if z.String() != "0" {
		t.Fatal(z.String())
	}
}

func TestCmp(t *testing.T) {
	a := String("1/2")
	b := String("1/3")
	if a.LessThan(b) {
		t.Fatal()
	}
}

func TestImu(t *testing.T) {
	d := String("1")
	// a := String("1/2")
	
	exp := String("4")
    got := d.MinusInt(1).AddInt(2).MulInt(6).QuoInt(3)

	if d.String() != "1" {
		t.Fatal(d.String())
	}

	if got.Equal(exp) != true {
		t.Log(d.MinusInt(1))
		t.Fatal("fatal")
	}

}
