# Intent

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
	"push-cash"
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
        ID: "29396fea-7596-4eb1-8faa-a2352c595590",
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
	"push-cash"
	"push-cash/pkg/models/operations"
	"push-cash/pkg/models/shared"
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
            Amount: 438601,
            Currency: shared.CurrencyUsd,
            Direction: shared.DirectionCashOut,
            UserID: "doloribus",
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
	"push-cash"
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
        ID: "f1a3a2fa-9467-4739-a51a-a52c3f5ad019",
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
	"push-cash"
	"push-cash/pkg/models/operations"
	"push-cash/pkg/types"
	"push-cash/pkg/models/shared"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(shared.Security{
            Bearer: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Intent.List(ctx, operations.ListIntentsRequest{
        CreatedAtAfter: types.MustTimeFromString("2020-12-24T08:13:29.299Z"),
        CreatedAtBefore: types.MustTimeFromString("2022-01-11T05:45:42.485Z"),
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
        Status: []shared.IntentStatus{
            shared.IntentStatusTimedout,
            shared.IntentStatusDeclined,
            shared.IntentStatusDeclined,
            shared.IntentStatusChargedback,
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

