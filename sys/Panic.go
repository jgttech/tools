package sys

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
