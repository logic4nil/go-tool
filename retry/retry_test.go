package retry_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/logic4nil/go-tool/retry"
)

func TestNewFixedIntervelRetryer(*testing.T) {
	r := retry.NewFixedIntervelRetryer(2, 1000)

	r.On(func() error {
		fmt.Println("test")
		return errors.New("test")
	})
}
