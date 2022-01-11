package main

import "errors"

func main() {

}

func Generate(i int) error {
	if i%2 == 0 {
		return errors.New("oh no")
	}
	return nil
}
