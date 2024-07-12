package main

import "fmt"

func main() {
	var a int8
	a = 32
	fmt.Println(a)
	var b int16
	b = 1218
	fmt.Println(b)
	var c int32
	c = 54644
	fmt.Println(c)
	var d int64
	d = 8465498
	fmt.Println(d)
	var e int
	e = 5456464
	fmt.Println(e)
	var k float32
	k = 0.55
	fmt.Println(k)
	var j float64
	j = 0.656541
	fmt.Println(j)
	var q bool
	q = true
	fmt.Println(q)
	var w bool
	w = false
	fmt.Println(w)
	var z rune
	z = 99
	fmt.Print(z)
	var o byte
	o = 127
	fmt.Println(o)
	y := "My Name is Neka"
	fmt.Println(y)
	var i complex64
	i = 25 + 6
	fmt.Println(i)
	var h complex128
	h = 46 + 55
	fmt.Println(h)
	var hi uint
	hi = 220
	fmt.Println(hi)
	var ho uint8
	ho = 255
	fmt.Println(ho)
	var he uint16
	he = 12334
	fmt.Println(he)
	var hr uint32
	hr = 54965
	fmt.Println(hr)
	var hg uint64
	hg = 49654321
	fmt.Println(hg)

	var m int
	m = 200
	var n int
	n = 50

	s := m + n
	p := m * n
	r := m / n
	u := m - n
	fmt.Println("ответ на данныйе примеры \n ", s, p, r, u)
}
