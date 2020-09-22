package main

import "fmt"

type customErr struct {
	ErrorString string
	ReturnValue int
	Fatal       bool
}

// implement error interface
func (c *customErr) Error() string {
	return fmt.Sprintf("Custom error: %v (exit %v, fatal == %v)", c.ErrorString, c.ReturnValue, c.Fatal)
}

func foo(e error) {
	fmt.Printf("%T satisfies the error interface\n", e)
}

func main() {
	c := customErr{
		ErrorString: "Something went wrong",
		ReturnValue: 255,
		Fatal:       true,
	}

	foo(&c)
	fmt.Println(c.Error())
}
