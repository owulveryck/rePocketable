package pocket

import (
	"context"
	"testing"
	"time"
)

func Test_hasSlept(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()
		healthCheck := 2 * time.Millisecond
		sleptTime := 1 * time.Millisecond
		c := hasSlept(ctx, healthCheck, sleptTime)
		select {
		case <-c:
			return
		case <-ctx.Done():
			t.Errorf("hasSlept() expected signal")
		}
	})
}
