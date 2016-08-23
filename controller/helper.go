package controller

func NewError(msg string) map[string]string {
	return map[string]string{"error": msg}
}
