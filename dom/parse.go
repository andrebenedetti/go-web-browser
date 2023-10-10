package dom

func ParseText(body string) string {
	text := ""
	inAngle := false
	for _, c := range body {
		if c == '<' {
			inAngle = true
		} else if c == '>' {
			inAngle = false
		} else if !inAngle {
			text = text + string(c)
		}
	}

	return text
}
