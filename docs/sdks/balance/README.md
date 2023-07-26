# Balance

### Available Operations

* [GetBalance](#getbalance) - Get Push Account balance
* [GetTransaction](#gettransaction) - Get a transaction
* [List](#list) - List transactions

## GetBalance

View Push Account balance

### Example Usage

```go
package main

import(
	"context"
	"log"
	"push-cash"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetBalanceResponse](../../models/operations/getbalanceresponse.md), error**


## GetTransaction

Retrieves a specific transaction by its ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"push-cash"
	"push-cash/pkg/models/shared"
	"push-cash/pkg/models/operations"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Balance.GetTransaction(ctx, operations.GetTransactionRequest{
        ID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetTransactionRequest](../../models/operations/gettransactionrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetTransactionResponse](../../models/operations/gettransactionresponse.md), error**


## List

Retrieves a list of transactions

### Example Usage

```go
package main

import(
	"context"
	"log"
	"push-cash"
	"push-cash/pkg/models/shared"
	"push-cash/pkg/models/operations"
	"push-cash/pkg/types"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Balance.List(ctx, operations.ListTransactionsRequest{
        CreatedAtAfter: types.MustTimeFromString("2022-12-17T05:46:24.151Z"),
        CreatedAtBefore: types.MustTimeFromString("2022-03-03T02:15:00.468Z"),
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
        Status: shared.TransactionStatusAvailable.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListTransactions200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListTransactionsRequest](../../models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListTransactionsResponse](../../models/operations/listtransactionsresponse.md), error**

