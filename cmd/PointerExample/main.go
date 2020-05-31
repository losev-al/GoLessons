package main

type test struct {
	s string
	i int32
}

//go:noinline
func print(t *test) *test {

	println("testWithReturn() t =", t)
	return t
}

//go:noinline
func testWithReturn() *test {
	x := test{}
	println("testWithReturn() &x =", &x)
	return &x
}

//go:noinline
func testWithoutReturn() {
	x := test{}
	println("testWithoutReturn() &x =", &x)
	print(&x)
}

func main() {
	testWithoutReturn()
	testWithoutReturn()
	a := testWithReturn()
	println("main() a =", a)
	a = testWithReturn()
	println("main() a =", a)
}
