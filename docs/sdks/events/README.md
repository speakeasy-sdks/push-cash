# Events
(*.Events*)

### Available Operations

* [GetEvent](#getevent) - Retrieve an event
* [List](#list) - List events

## GetEvent

Retrieves a specific event by its ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	pushcash "push-cash/v2"
	"push-cash/v2/pkg/models/shared"
	"push-cash/v2/pkg/models/operations"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(""),
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetEventRequest](../../models/operations/geteventrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.GetEventResponse](../../models/operations/geteventresponse.md), error**


## List

Retrieves a list of events.

### Example Usage

```go
package main

import(
	"context"
	"log"
	pushcash "push-cash/v2"
	"push-cash/v2/pkg/models/shared"
	"push-cash/v2/pkg/models/operations"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(""),
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListEventsRequest](../../models/operations/listeventsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ListEventsResponse](../../models/operations/listeventsresponse.md), error**

