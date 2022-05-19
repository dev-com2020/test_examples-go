package main

/*
  #include "hello.c"
*/
import "C"
import (
	"errors"
	"log"
)

func main() {
	err := Hello()
	if err != nil {
		log.Fatal(err)
	}
}

func Hello() error {
	_, err := C.Hello()
	if err != nil {
		return errors.New("blad z funkcji" + err.Error())
	}
	return nil
}
