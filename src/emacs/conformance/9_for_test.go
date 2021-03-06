package conformance

func testFor(n int) int {
	x := 0
	for i := 0; i < n; i++ {
		x++
	}
	return x
}

func testWhile(n int) int {
	x, i := 0, 0
	for i < n {
		x, i = x+1, i+1
	}
	return x
}

func testForBreak(n int) int {
	x := 0
	for {
		x += n
		break
	}
	return x
}

func testForContinue(n int) int {
	x := 0
	for i := 0; i < n*2; i++ {
		if x == n {
			continue
		}
		x++
	}
	return x
}

func testNestedFor(n int) int {
	x := 0
	for i := 0; i < n; i++ {
		for i := 0; i < 4; i++ {
			x++
		}
		x -= 3
	}
	return x
}

func testNestedForBreak(n int) int {
	x := 0
	for i := 0; i < n; i++ {
		for {
			x--
			break
		}
		x += 2
	}
	return x
}

func testNestedForContinue(n int) int {
	x := 0
	for i := 0; i < n*2; i++ {
		for j := 0; j < 5; j++ {
			if x == n {
				continue
			}
			x++
		}
	}
	return x
}

func testForScopes1(n int) int {
	x := n + 1
	for {
		y := 1
		x -= y
		break
	}
	return x
}

func testForScopes2(n int) int {
	x := n + 1
	for {
		{
			y := 1
			x -= y
			break
		}
	}
	return x
}

func testNestedForScopes1(n int) int {
	x := n + 1
	for i := 0; i < 1; i++ {
		for {
			y := 1
			x -= y
			break
			z := 2 // Unreachable
			panic(z)
		}
	}
	return x
}

func testNestedForScopes2(n int) int {
	x := n + 1
	for i := 0; i < 1; i++ {
		for {
			{
				y := 1
				x -= y
				break
				z := 2 // Unreachable
				panic(z)
			}
		}
	}
	return x
}
