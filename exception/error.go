package exception

func PanicNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}