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
    res, err := s.Intent.CancelIntent(ctx, operations.CancelIntentRequest{
        ID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CancelIntentRequest](../../pkg/models/operations/cancelintentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CancelIntentResponse](../../pkg/models/operations/cancelintentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateIntent

Create a payment intent

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
    res, err := s.Intent.CreateIntent(ctx, operations.CreateIntentRequest{
        CreateIntentRequest: &shared.CreateIntentRequest{
            Amount: 1450,
            Currency: shared.CurrencyUsd,
            Direction: shared.DirectionCashIn,
            UserID: "user_28CJjV7P4Go5PNJvfzghiD",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateIntentRequest](../../pkg/models/operations/createintentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateIntentResponse](../../pkg/models/operations/createintentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetIntent

Get an intent by ID

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
    res, err := s.Intent.GetIntent(ctx, operations.GetIntentRequest{
        ID: "<id>",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetIntentRequest](../../pkg/models/operations/getintentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetIntentResponse](../../pkg/models/operations/getintentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Retrieves a list of intents

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
    res, err := s.Intent.List(ctx, operations.ListIntentsRequest{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListIntentsRequest](../../pkg/models/operations/listintentsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListIntentsResponse](../../pkg/models/operations/listintentsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
