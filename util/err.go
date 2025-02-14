package util

func CheckErr(err error) {
	if err == nil {
		return
	}
	panic(err)
}
