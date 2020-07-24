package logs

import (
	"fmt"
)

func Logs(value string) error {
	fmt.Printf("%v\n", value)

	return nil
}
