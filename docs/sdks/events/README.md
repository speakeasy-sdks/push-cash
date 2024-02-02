# Events
(*Events*)

### Available Operations

* [GetEvent](#getevent) - Retrieve an event
* [List](#list) - List events

## GetEvent

Retrieves a specific event by its ID.

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
    res, err := s.Events.GetEvent(ctx, operations.GetEventRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Event != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetEventRequest](../../pkg/models/operations/geteventrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetEventResponse](../../pkg/models/operations/geteventresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Retrieves a list of events.

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
    res, err := s.Events.List(ctx, operations.ListEventsRequest{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListEventsRequest](../../pkg/models/operations/listeventsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListEventsResponse](../../pkg/models/operations/listeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
