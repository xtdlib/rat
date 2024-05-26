package rat

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var DefaultPrecision = 4

type rat struct {
	bigrat    *big.Rat
	precision int
}

func (r *rat) Add(b *rat) *rat {
	tmpbigrat := new(big.Rat)
	tmpbigrat.Add(r.bigrat, b.bigrat)
	return &rat{
		bigrat:    tmpbigrat,
		precision: r.precision,
	}
}

func (r *rat) AddInt(v int) *rat {
	return r.Add(Int(v))
}

func (r *rat) Neg() *rat {
	tmpbigrat := new(big.Rat)
	tmpbigrat.Neg(r.bigrat)
	return &rat{
		bigrat:    tmpbigrat,
		precision: r.precision,
	}
}

func (r *rat) Mul(b *rat) *rat {
	tmpbigrat := new(big.Rat)
	tmpbigrat.Mul(r.bigrat, b.bigrat)
	return &rat{
		bigrat:    tmpbigrat,
		precision: r.precision,
	}
}

func (r *rat) MulInt(v int) *rat {
	return r.Mul(Int(v))
}

func (r *rat) Quo(b *rat) *rat {
	tmpbigrat := new(big.Rat)
	tmpbigrat.Quo(r.bigrat, b.bigrat)
	return &rat{
		bigrat:    tmpbigrat,
		precision: r.precision,
	}
}

func (r *rat) QuoInt(v int) *rat {
	return r.Quo(Int(v))
}

func (r *rat) Minus(b *rat) *rat {
	tmpbigrat := new(big.Rat)
	tmpbigrat.Neg(b.bigrat)
	tmpbigrat.Add(tmpbigrat, r.bigrat)
	return &rat{
		bigrat:    tmpbigrat,
		precision: r.precision,
	}
}

func (r *rat) MinusInt(v int) *rat {
	return r.Minus(Int(v))
}

func (r *rat) PowInt(n int) *rat {
	nr := Clone(r)
	ret := Int(1)

	for i := 0; i < n; i++ {
		ret = ret.Mul(nr)
	}

	return ret
}

func (r *rat) String() string {
	n, exact := r.bigrat.FloatPrec()
	if exact {
		return r.bigrat.FloatString(min(r.precision, n))
	}
	return r.bigrat.FloatString(r.precision)
}

func (r *rat) Clone() *rat {
	tmpr := new(big.Rat)
	tmpr.Set(r.bigrat)
	return &rat{
		bigrat:    tmpr,
		precision: r.precision,
	}
}

func (r *rat) LessThan(b *rat) bool {
	if r.bigrat.Cmp(b.bigrat) == -1 {
		return true
	}
	return false
}

func (r *rat) GreaterThan(b *rat) bool {
	if r.bigrat.Cmp(b.bigrat) == 1 {
		return true
	}
	return false
}

func (r *rat) Equal(b *rat) bool {
	if r.bigrat.Cmp(b.bigrat) == 0 {
		return true
	}
	return false
}

func Quo(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Quo(a.bigrat, b.bigrat)
	return &rat{
		bigrat:    c,
		precision: a.precision,
	}
}

func Mul(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Mul(a.bigrat, b.bigrat)
	return &rat{
		bigrat:    c,
		precision: a.precision,
	}
}

func Add(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Add(a.bigrat, b.bigrat)
	return &rat{
		bigrat:    c,
		precision: a.precision,
	}
}

func Neg(a *rat) *rat {
	c := new(big.Rat)
	c.Neg(a.bigrat)
	return &rat{
		bigrat:    c,
		precision: a.precision,
	}
}

func Minus(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Add(a.bigrat, Neg(b).bigrat)
	return &rat{
		bigrat:    c,
		precision: a.precision,
	}
}

func Zero() *rat {
	tmpr := big.NewRat(0, 1)
	return &rat{
		bigrat:    tmpr,
		precision: DefaultPrecision,
	}
}

func Clone(r *rat) *rat {
	// nil check?
	tmpr := new(big.Rat)
	tmpr.Set(r.bigrat)
	return &rat{
		bigrat:    tmpr,
		precision: DefaultPrecision,
	}
}

func BigRat(a, b int64) *rat {
	return &rat{
		bigrat:    big.NewRat(a, b),
		precision: DefaultPrecision,
	}
}

func Int(a int) *rat {
	return &rat{
		bigrat:    big.NewRat(int64(a), 1),
		precision: DefaultPrecision,
	}
}

func Int64(a int64) *rat {
	return &rat{
		bigrat:    big.NewRat(a, 1),
		precision: DefaultPrecision,
	}
}

func Parse(v string) *rat {
	// NOTE: not work for "0.5/1"
	// nr := new(big.Rat)
	// _, ok := nr.SetString(v)
	// if !ok {
	// 	panic("rat: invalid rat string " + v)
	// }
	//
	// return &rat{
	// 	bigrat:    nr,
	// 	precision: DefaultPrecision,
	// }

	if strings.Contains(v, "/") {
		arr := strings.Split(v, "/")
		if len(arr) != 2 {
			panic("rat: invalid rat string " + v)
		}
		return Parse(arr[0]).Quo(Parse(arr[1]))
	}

	if strings.Contains(v, ".") {
		tmpr := new(big.Rat)
		_, ok := tmpr.SetString(v)
		if !ok {
			panic("rat: invalid rat string: " + v)
		}
		return &rat{
			bigrat:    tmpr,
			precision: DefaultPrecision,
		}
	}

	tmpr := new(big.Rat)
	_, ok := tmpr.SetString(v)
	if !ok {
		panic("rat: invalid rat string: " + v)
	}
	return &rat{
		bigrat:    tmpr,
		precision: DefaultPrecision,
	}
}

func mustParseInt(v string) int64 {
	iv, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("rat: %v", err))
	}
	return iv
}

func pow10(n int) int64 {
	var v int64 = 1
	for i := 0; i < n; i++ {
		v *= 10
	}
	return v
}
