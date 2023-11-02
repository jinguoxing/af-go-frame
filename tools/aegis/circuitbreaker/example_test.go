package circuitbreaker_test

import (
	"fmt"

	"github.com/jinguoxing/af-go-frame/tools/aegis/circuitbreaker/sre"
)

// This is a example of using a circuit breaker Do() when return nil.
func Example() {
	b := sre.NewBreaker()
	for i := 0; i < 1000; i++ {
		b.MarkSuccess()
	}
	for i := 0; i < 100; i++ {
		b.MarkFailed()
	}

	err := b.Allow()
	fmt.Printf("err=%v", err)
	// Output: err=<nil>
}
