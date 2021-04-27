package lib

import "log"

func IsError(err error, pre string) error {
	if err != nil {
		log.Printf("%v: %v", pre, err)
	}
	return err
}

func IsErrorBool(err error, pre string) (b bool) {
	if err != nil {
		log.Printf("%v: %v", pre, err)
		b = true
	}
	return
}
