package bfk

import "testing"

func dummy() {}

var helloworld = []byte("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")

var a [256]func()

func init() {
	for _, c := range []byte("+-><,.[]") {
		a[c] = dummy
	}
}

var m = map[byte]func(){
	'+': dummy,
	'-': dummy,
	'>': dummy,
	'<': dummy,
	',': dummy,
	'.': dummy,
	'[': dummy,
	']': dummy,
}

var f func()

func BenchmarkArray(b *testing.B) {
	var v func()
	for i := 0; i < b.N; i++ {
		for _, c := range helloworld {
			v = a[c]
		}
	}
	f = v
}

func BenchmarkSwitch(b *testing.B) {
	var v func()
	for i := 0; i < b.N; i++ {
		for _, c := range helloworld {
			switch c {
			case '+':
				v = dummy
			case '-':
				v = dummy
			case '>':
				v = dummy
			case '<':
				v = dummy
			case ',':
				v = dummy
			case '.':
				v = dummy
			case '[':
				v = dummy
			case ']':
				v = dummy
			}
		}
	}
	f = v
}

func BenchmarkMap(b *testing.B) {
	var v func()
	for i := 0; i < b.N; i++ {
		for _, c := range helloworld {
			v = m[c]
		}
	}
	f = v
}
