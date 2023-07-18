package helper

// ErrorPanic throw a panic when there's an error
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}