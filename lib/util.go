package lib

//DoIf runs a passed function if the condition is true
func DoIf(condition bool, f func()) {
	if condition {
		f()
	}
}
