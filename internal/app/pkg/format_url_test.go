package pkg

import (
	"fmt"
)

func ExampleFormatLocalURL() {
	baseURL := "http://localhost"
	code := "testcode"
	url := FormatLocalURL(baseURL, code)

	fmt.Println(url)

	// Output:
	// http://localhost/testcode
}
