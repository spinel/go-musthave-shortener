package pkg

import "fmt"

//FormatLocalURL used to format base url and code.
func FormatLocalURL(baseURL, code string) string {
	return fmt.Sprintf("%s/%s", baseURL, code)
}
