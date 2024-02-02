# Balance
(*Balance*)

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
	"push-cash/v3/pkg/models/shared"
	pushcash "push-cash/v3"
	"context"
	"log"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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

**[*operations.GetBalanceResponse](../../pkg/models/operations/getbalanceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTransaction

Retrieves a specific transaction by its ID.

### Example Usage

```go
package main

import(
	"push-cash/v3/pkg/models/shared"
	pushcash "push-cash/v3"
	"context"
	"push-cash/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Balance.GetTransaction(ctx, operations.GetTransactionRequest{
        ID: "<ID>",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTransactionRequest](../../pkg/models/operations/gettransactionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetTransactionResponse](../../pkg/models/operations/gettransactionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Retrieves a list of transactions

### Example Usage

```go
package main

import(
	"push-cash/v3/pkg/models/shared"
	pushcash "push-cash/v3"
	"context"
	"push-cash/v3/pkg/models/operations"
	"log"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Balance.List(ctx, operations.ListTransactionsRequest{
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListTransactionsRequest](../../pkg/models/operations/listtransactionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListTransactionsResponse](../../pkg/models/operations/listtransactionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
