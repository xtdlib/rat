package rat

import (
	"log/slog"
	"math/big"
	"strings"
)

var DefaultPrecision = 7

func Rat[T ~int | ~int32 | ~int64 | ~float64 | ~float32 | ~string](v T) *rat {
	switch v := any(v).(type) {
	case float32:
		return parseFloat64(float64(v))
	case float64:
		return parseFloat64(float64(v))
	case int:
		return parseInt64(int64(v))
	case int32:
		return parseInt64(int64(v))
	case int64:
		return parseInt64(int64(v))
	case string:
		return parse(v)
	}
	return nil
}

type rat struct {
	bigrat    *big.Rat
	Precision int
}

func (r *rat) Add(b *rat) *rat {
	r.bigrat.Add(r.bigrat, b.bigrat)
	return r
}

func (r *rat) AddInt(v int) *rat {
	return r.Add(parseInt(v))
}

func (r *rat) Neg() *rat {
	r.bigrat.Neg(r.bigrat)
	return r
}

func (r *rat) Mul(b *rat) *rat {
	r.bigrat.Mul(r.bigrat, b.bigrat)
	return r
}

func (r *rat) MulInt(v int) *rat {
	return r.Mul(parseInt(v))
}

func (r *rat) Quo(b *rat) *rat {
	r.bigrat.Quo(r.bigrat, b.bigrat)
	return r
}

func (r *rat) QuoInt(v int) *rat {
	return r.Quo(parseInt(v))
}

func (r *rat) Minus(b *rat) *rat {
	r.Add(Neg(b))
	// tmpbigrat := new(big.Rat)
	// tmpbigrat.Neg(b.bigrat)
	// r.bigrat.Add(r.bigrat, tmpbigrat)
	return r
}

func (r *rat) MinusInt(v int) *rat {
	return r.Minus(parseInt(v))
}

func (r *rat) PowInt(n int) *rat {
	v := r.Clone()
	if n == 0 {
		r.Set(parseInt(1))
		return r
	}
	if n > 0 {
		return r.PowInt(n - 1).Mul(v)
	}
	return r
}

func (r *rat) String() string {
	n, exact := r.bigrat.FloatPrec()
	if exact {
		return r.bigrat.FloatString(min(r.Precision, n))
	}
	return r.bigrat.FloatString(r.Precision)
}

func (r *rat) Set(v *rat) *rat {
	r.bigrat.Set(v.bigrat)
	return r
}

func (r *rat) Clone() *rat {
	tmpr := new(big.Rat)
	tmpr.Set(r.bigrat)
	return &rat{
		bigrat:    tmpr,
		Precision: r.Precision,
	}
}

func (r *rat) IsLessThan(b *rat) bool {
	if r.bigrat.Cmp(b.bigrat) == -1 {
		return true
	}
	return false
}

func (r *rat) IsGreaterThan(b *rat) bool {
	if r.bigrat.Cmp(b.bigrat) == 1 {
		return true
	}
	return false
}

func (r *rat) IsEqual(b *rat) bool {
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
		Precision: a.Precision,
	}
}

func Mul(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Mul(a.bigrat, b.bigrat)
	return &rat{
		bigrat:    c,
		Precision: a.Precision,
	}
}

func Add(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Add(a.bigrat, b.bigrat)
	return &rat{
		bigrat:    c,
		Precision: a.Precision,
	}
}

func Neg(a *rat) *rat {
	c := new(big.Rat)
	c.Neg(a.bigrat)
	return &rat{
		bigrat:    c,
		Precision: a.Precision,
	}
}

func Minus(a *rat, b *rat) *rat {
	c := new(big.Rat)
	c.Add(a.bigrat, Neg(b).bigrat)
	return &rat{
		bigrat:    c,
		Precision: a.Precision,
	}
}

func Zero() *rat {
	tmpr := big.NewRat(0, 1)
	return &rat{
		bigrat:    tmpr,
		Precision: DefaultPrecision,
	}
}

func Clone(r *rat) *rat {
	// nil check?
	tmpr := new(big.Rat)
	tmpr.Set(r.bigrat)
	return &rat{
		bigrat:    tmpr,
		Precision: DefaultPrecision,
	}
}

func New(a, b int64) *rat {
	return &rat{
		bigrat:    big.NewRat(a, b),
		Precision: DefaultPrecision,
	}
}

func parseFloat64(a float64) *rat {
	bigrat := new(big.Rat)
	bigrat.SetFloat64(a)
	return &rat{
		bigrat:    bigrat,
		Precision: DefaultPrecision,
	}
}

func parseInt(a int) *rat {
	return &rat{
		bigrat:    big.NewRat(int64(a), 1),
		Precision: DefaultPrecision,
	}
}

func parseInt64(a int64) *rat {
	return &rat{
		bigrat:    big.NewRat(a, 1),
		Precision: DefaultPrecision,
	}
}

func parse(v string) (out *rat) {
	defer func() {
		if out == nil {
			slog.Error("rat: invalid rat string " + v)
		}
	}()

	if strings.HasSuffix(v, "%") {
		defer func() {
			if out == nil {
				return
			}
			out.QuoInt(100)
		}()
		v = v[0 : len(v)-1]
	}

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
			return nil
		}
		x := parse(arr[0])
		if x == nil {
			return nil
		}
		y := parse(arr[1])
		if y == nil {
			return nil
		}
		return x.Quo(y)
	}

	if strings.Contains(v, ".") {
		tmpr := new(big.Rat)
		_, ok := tmpr.SetString(v)
		if !ok {
			return nil
		}
		return &rat{
			bigrat:    tmpr,
			Precision: DefaultPrecision,
		}
	}

	tmpr := new(big.Rat)
	_, ok := tmpr.SetString(v)
	if !ok {
		slog.Error("rat: invalid rat string " + v)
	}
	return &rat{
		bigrat:    tmpr,
		Precision: DefaultPrecision,
	}
}
