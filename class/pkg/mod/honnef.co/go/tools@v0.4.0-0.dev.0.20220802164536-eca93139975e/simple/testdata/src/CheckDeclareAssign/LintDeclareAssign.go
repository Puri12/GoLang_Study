package pkg

func fn() {
	var x int //@ diag(`should merge variable declaration with assignment on next line`)
	x = 1
	_ = x

	var y interface{} //@ diag(`should merge variable declaration with assignment on next line`)
	y = 1
	_ = y

	if true {
		var x string //@ diag(`should merge variable declaration with assignment on next line`)
		x = ""
		_ = x
	}

	var z []string
	z = append(z, "")
	_ = z

	var f func()
	f = func() { f() }
	_ = f

	var a int
	a = 1
	a = 2
	_ = a

	var b int
	b = 1
	// do stuff
	b = 2
	_ = b

	var c int
	unrelated = 1
	_ = c
}

var unrelated int
