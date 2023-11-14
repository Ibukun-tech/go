package goes

func ThinkingWell() string {
	return "You are still learning and the best way to learn is by doingthe practical of what you did"
}
func Math(value int, value1 int, s string) int {
	switch {
	case s == "+":
		return value + value1
	case s == "-":
		return value - value1
	}
	return 0
}
func Food(s string) string {
	return "This is the best food " + s
}
