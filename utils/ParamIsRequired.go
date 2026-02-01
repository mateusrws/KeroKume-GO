package utils

import "fmt"

func ParamIsRequired(n, t string) error {
	return fmt.Errorf("param %s (type %s) is required", n, t)
}