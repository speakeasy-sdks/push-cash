# Intent
(*Intent*)

### Available Operations

* [CancelIntent](#cancelintent) - Cancel an intent
* [CreateIntent](#createintent) - Create intent
* [GetIntent](#getintent) - Get an intent
* [List](#list) - List intents

## CancelIntent

Cancels a specific intent identified by its ID

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
    res, err := s.Intent.CancelIntent(ctx, operations.CancelIntentRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Intent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CancelIntentRequest](../../models/operations/cancelintentrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CancelIntentResponse](../../models/operations/cancelintentresponse.md), error**


## CreateIntent

Create a payment intent

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
    res, err := s.Intent.CreateIntent(ctx, operations.CreateIntentRequest{
        CreateIntentRequest: &shared.CreateIntentRequest{
            Amount: 135934,
            Currency: "Seychelles Rupee",
            Direction: shared.DirectionCashIn,
            UserID: "Chair Music failing",
        },
        XIdempotencyKey: "f1bbb856-fb17-11ed-be56-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Intent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateIntentRequest](../../models/operations/createintentrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateIntentResponse](../../models/operations/createintentresponse.md), error**


## GetIntent

Get an intent by ID

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
    res, err := s.Intent.GetIntent(ctx, operations.GetIntentRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Intent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetIntentRequest](../../models/operations/getintentrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetIntentResponse](../../models/operations/getintentresponse.md), error**


## List

Retrieves a list of intents

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
    res, err := s.Intent.List(ctx, operations.ListIntentsRequest{
        CreatedAtAfter: types.MustTimeFromString("2023-04-22T14:13:10.937Z"),
        CreatedAtBefore: types.MustTimeFromString("2021-04-20T09:15:20.369Z"),
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
        Status: []shared.IntentStatus{
            shared.IntentStatusDeclined,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListIntents200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListIntentsRequest](../../models/operations/listintentsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListIntentsResponse](../../models/operations/listintentsresponse.md), error**

