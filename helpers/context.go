package helpers

import (
	"context"
	"time"
)

//TimeOut context to use some cases like database connection or etc
func TimeOut(t byte) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
}
