package helper

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
