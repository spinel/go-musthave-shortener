package pkg

func CheckNA(s string) string {
	if s == "" {
		s = "N/A"
	}
	return s
}
