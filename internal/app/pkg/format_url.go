package pkg

import "fmt"

func FormatLocalURL(baseURL, code string) string {
	return fmt.Sprintf("%s/%s", baseURL, code)
}
