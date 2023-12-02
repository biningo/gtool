package exception

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertError(t *testing.T) {
	{
		defer func() {
			err := recover()
			assert.Nil(t, err)
		}()
		assertError(nil)
	}

	{
		defer func() {
			err := recover()
			assert.NotNil(t, err)
		}()
		assertError(errors.New("error"))
	}
}
