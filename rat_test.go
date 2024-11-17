package rat

import (
	"testing"
)

func TestParcentage(t *testing.T) {
	a := parse("3%")
	t.Log(a.String())
}

func TestBasics(t *testing.T) {
	a := parse("2")
	b := parse("4")
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
	a := parse("2")
	b := parse("4")
	a.Minus(b).IsEqual(Rat("-2"))
}

func TestQuo(t *testing.T) {
	a := parse("2")
	b := parse("4")
	if Quo(a, b).String() != "0.5" {
		t.Fatal("Quo")
	}
}

func TestPowInt(t *testing.T) {
	a := parse("2")              // a = 2
	if a.String() != "2" {
		t.Fatalf("expected 2, got %s", a.String())
	}

	if ! a.PowInt(0).IsEqual(Rat("1")) {
		t.Fatalf("expected 0, got %s", a.String())
	}

	a.Mul(Rat("2"))

	exp := a.PowInt(2).String() // a^10 = 1024
	if exp != "4" {
		t.Fatalf("expected 1024, got %s", exp)
	}

}

func TestString(t *testing.T) {
	{
		a := parse("0.5/1")
		if a.String() != "0.5" {
			t.Fatal(a.String())
		}
	}

	{
		a := parse("1/0.5")
		if a.String() != "2" {
			t.Fatal(a.String())
		}
	}

	{
		a := parse("1/3")
		a.Precision = 8
		if a.String() != "0.33333333" {
			t.Fatal(a.String())
		}
	}
}

func TestCopy(t *testing.T) {
	a := parse("1/2")
	b := a.Clone()
	if a.Add(parse("1")).String() != "1.5" {
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
	a := parse("1/2")
	b := parse("1/3")
	if a.IsLessThan(b) {
		t.Fatal()
	}
}

func TestImu(t *testing.T) {
	d := parse("1")

	exp := parse("4")
	got := d.MinusInt(1).AddInt(2).MulInt(6).QuoInt(3) // 4

	if d.String() != "4" {
		t.Fatal(d.String())
	}

	if got.IsEqual(exp) != true {
		t.Fatal("fatal")
	}
}

// func ExampleAdd() {
// 	a := Parse("0.1")
// 	fmt.Println(a.Add(Parse("0.2")))                      // 0.3
// 	fmt.Println(a.Add(Parse("0.2")).Equal(BigRat(3, 10))) // true
// 	fmt.Println(a)                                        // 0.1
// 	// Output: 0.3
// 	// true
// 	// 0.1
// }
