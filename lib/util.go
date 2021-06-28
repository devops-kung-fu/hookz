package lib

//DoIf runs a passed function if the condition is true
func DoIf(f func(), condition bool) {
	if condition {
		f()
	}
}
