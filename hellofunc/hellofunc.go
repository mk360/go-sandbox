package hellofunc

import (
	"errors"
	"fmt"
)

func Hello(name string) error {
	if name == "" {
		return errors.New("no name given")
	} else {
		fmt.Println("DZIEN DOBRY " + name)
		return nil
	}
}
