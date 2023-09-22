# Transfer

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
	"context"
	"log"
	pushcash "push-cash"
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
    res, err := s.Transfer.CreateTransfer(ctx, operations.CreateTransferRequest{
        CreateTransferRequest: &shared.CreateTransferRequest{
            Amount: 468651,
            Currency: "praesentium",
            Direction: shared.DirectionCashOut,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateTransferResponse](../../models/operations/createtransferresponse.md), error**


## GetTransfer

Retrieves a specific transfer by its ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	pushcash "push-cash"
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
    res, err := s.Transfer.GetTransfer(ctx, operations.GetTransferRequest{
        ID: "097b0074-f154-471b-9e6e-13b99d488e1e",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetTransferRequest](../../models/operations/gettransferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetTransferResponse](../../models/operations/gettransferresponse.md), error**


## List

Retrieves a list of transfers.

### Example Usage

```go
package main

import(
	"context"
	"log"
	pushcash "push-cash"
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
    res, err := s.Transfer.List(ctx, operations.ListTransfersRequest{
        CreatedAtAfter: types.MustTimeFromString("2022-11-01T07:52:08.326Z"),
        CreatedAtBefore: types.MustTimeFromString("2022-03-02T21:33:21.372Z"),
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
        Status: []shared.TransferStatus{
            shared.TransferStatusCreated,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListTransfers200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListTransfersRequest](../../models/operations/listtransfersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListTransfersResponse](../../models/operations/listtransfersresponse.md), error**

