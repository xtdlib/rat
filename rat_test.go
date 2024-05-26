package rat

import (
	"fmt"
	"testing"
)

func TestBasics(t *testing.T) {
	a := Parse("2")
	b := Parse("4")
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
	a := Parse("2")
	b := Parse("4")
	a.Minus(b)

}

func TestQuo(t *testing.T) {
	a := Parse("2")
	b := Parse("4")
	if Quo(a, b).String() != "0.5" {
		t.Fatal("Quo")
	}
}

func TestPowInt(t *testing.T) {
	a := Parse("2")
	exp := a.PowInt(10).String()
	if exp != "1024" {
		t.Fatal(exp)
	}

	b := Parse("1/2")
	if a.Add(b).String() != "2.5" {
		t.Fatal()
	}
}

func TestString(t *testing.T) {
	{
		a := Parse("0.5/1")
		if a.String() != "0.5" {
			t.Fatal(a.String())
		}
	}

	{
		a := Parse("1/0.5")
		if a.String() != "2" {
			t.Fatal(a.String())
		}
	}
}

func TestCopy(t *testing.T) {
	a := Parse("1/2")
	b := a.Clone()
	if a.Add(Parse("1")).String() != "1.5" {
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
	a := Parse("1/2")
	b := Parse("1/3")
	if a.LessThan(b) {
		t.Fatal()
	}
}

func TestImu(t *testing.T) {
	d := Parse("1")

	exp := Parse("4")
	got := d.MinusInt(1).AddInt(2).MulInt(6).QuoInt(3) // 4

	if d.String() != "1" {
		t.Fatal(d.String())
	}

	if got.Equal(exp) != true {
		t.Fatal("fatal")
	}
}

func ExampleAdd() {
	a := Parse("0.1")
	fmt.Println(a.Add(Parse("0.2")))                      // 0.3
	fmt.Println(a.Add(Parse("0.2")).Equal(BigRat(3, 10))) // true
	fmt.Println(a)                                        // 0.1
	// Output: 0.3
	// true
	// 0.1
}
