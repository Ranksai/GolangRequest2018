package main

import (
	"fmt"

	"os"

	"github.com/pkg/errors"
)

type Stringer interface {
	String() string
}

type String string

func (s String) String() string {
	return string(s)
}

type Int int

func (i Int) String() string {
	return string(i)
}

type MyError struct {
}

func (m *MyError) Error() error {
	return errors.New("not Stringer")
}

func ToStringer(v interface{}) (Stringer, error) {
	stringer, ok := v.(Stringer)
	if !ok {
		err := new(MyError)
		return nil, err.Error()
	}
	return stringer, nil
}

func main() {
	v := String("test")
	if stringer, err := ToStringer(v); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	} else {
		fmt.Println(stringer.String())
	}
}
