package hello

import (
	"fmt"

	"github.com/google/uuid"
)

func Greet(audience string) string {
	id := uuid.New()
	return fmt.Sprintf("Hello, %s! Your UUID is %s", audience, id.String())
}
