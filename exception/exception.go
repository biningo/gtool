package exception

func assertError(err error) {
	if err != nil {
		panic(err)
	}
}
