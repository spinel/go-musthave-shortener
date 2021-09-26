package pkg

import "fmt"

func FormatLocalUrl(baseUrl, code string) string {
	return fmt.Sprintf("%s/%s", baseUrl, code)
}
