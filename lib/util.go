package lib

func PrintIf(f func(), condition bool) {
	if condition {
		f()
	}
}
