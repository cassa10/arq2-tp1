package exception

import "fmt"

type CannotMapOrderState struct {
	State string
}

func (e CannotMapOrderState) Error() string {
	return fmt.Sprintf("cannot map order state %s", e.State)
}
