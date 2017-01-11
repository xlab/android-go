package android

type Unwind []func()

func (u Unwind) Add(cleanup func()) {
	u = append(u, cleanup)
}

func (u Unwind) Unwind() {
	for i := len(u) - 1; i >= 0; i-- {
		u[i]()
	}
}

func (u Unwind) Discard() {
	if len(u) > 0 {
		u = u[:0]
	}
}

func orPanic(err error, finalizers ...func()) {
	if err != nil {
		for _, fn := range finalizers {
			fn()
		}
		panic(err)
	}
}

func s(str string) string {
	if str[len(str)-1] == '\x00' {
		return str
	}
	return str + "\x00"
}
