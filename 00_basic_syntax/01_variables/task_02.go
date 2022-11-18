package main

func main() {
	aa := "t"
	ab := aa
	aa = "s"
	println(aa)
	println(ab)

	// its possible to assing constants to constants
	const x = "hallo"
	const c = x
	println(c)

	// you can't assign variables to constants
	v := "hallo"
	const d = v
	println(d)
}
