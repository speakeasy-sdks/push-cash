<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	pushcash "push-cash/v2"
	"push-cash/v2/pkg/models/shared"
)

func main() {
	s := pushcash.New(
		pushcash.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Balance.GetBalance(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.AccountBalance != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->