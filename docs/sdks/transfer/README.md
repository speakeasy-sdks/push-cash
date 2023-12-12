# Transfer
(*Transfer*)

### Available Operations

* [CreateTransfer](#createtransfer) - Create a transfer
* [GetTransfer](#gettransfer) - Retrieve a transfer
* [List](#list) - List transfers

## CreateTransfer

Create a transfer

### Example Usage

```go
package main

import(
	"push-cash/v2/pkg/models/shared"
	pushcash "push-cash/v2"
	"context"
	"push-cash/v2/pkg/models/operations"
	"log"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Transfer.CreateTransfer(ctx, operations.CreateTransferRequest{
        CreateTransferRequest: &shared.CreateTransferRequest{
            Amount: 759686,
            Currency: shared.CurrencyUsd,
            Direction: shared.DirectionCashIn,
        },
        XIdempotencyKey: "f1bbb856-fb17-11ed-be56-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateTransferRequest](../../pkg/models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateTransferResponse](../../pkg/models/operations/createtransferresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTransfer

Retrieves a specific transfer by its ID.

### Example Usage

```go
package main

import(
	"push-cash/v2/pkg/models/shared"
	pushcash "push-cash/v2"
	"context"
	"push-cash/v2/pkg/models/operations"
	"log"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Transfer.GetTransfer(ctx, operations.GetTransferRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Transfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetTransferRequest](../../pkg/models/operations/gettransferrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetTransferResponse](../../pkg/models/operations/gettransferresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

Retrieves a list of transfers.

### Example Usage

```go
package main

import(
	"push-cash/v2/pkg/models/shared"
	pushcash "push-cash/v2"
	"context"
	"push-cash/v2/pkg/models/operations"
	"log"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Transfer.List(ctx, operations.ListTransfersRequest{
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
        Status: []shared.TransferStatus{
            shared.TransferStatusFailed,
        },
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListTransfersRequest](../../pkg/models/operations/listtransfersrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListTransfersResponse](../../pkg/models/operations/listtransfersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
