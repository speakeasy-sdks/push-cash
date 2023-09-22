<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	pushcash "push-cash"
	"push-cash/pkg/models/shared"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(shared.Security{
            Bearer: "",
        }),
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
<!-- End SDK Example Usage -->