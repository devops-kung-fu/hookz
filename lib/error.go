//Package lib Functionality for the Hookz CLI
package lib

import "log"

//IsError Checks to see if an error exists, and if so
//writes it to the log with the provided prefix
func IsError(err error, prefix string) error {
	if err != nil {
		log.Printf("%v: %v", prefix, err)
	}
	return err
}

//IsErrorBool Checks to see if an error exists, and if so
//returns true after writing the error to the log with the provided prefix
func IsErrorBool(err error, prefix string) (b bool) {
	if err != nil {
		log.Printf("%v: %v", prefix, err)
		b = true
	}
	return
}
