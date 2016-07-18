package goutils

func Colorize(text string, status string) string {
	out := ""
	switch status {
	case "success":
		out = "\033[32;1m" //green
	case "fail":
		out = "\033[31;1m" //red
	case "warn":
		out = "\033[33;1m" //yellow
	case "note":
		out = "\033[34;1m" //blue
	default:
		out = "\033[0m" //no color
	}
	return out + text + "\033[0m"
}
