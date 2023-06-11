package helper

func String_ptr(str string) *string {
	s := new(string)
	*s = str
	return s
}
