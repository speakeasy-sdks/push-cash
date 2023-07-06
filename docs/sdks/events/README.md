# Events

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
    res, err := s.Events.GetEvent(ctx, operations.GetEventRequest{
        ID: "c2ddf7cc-78ca-41ba-928f-c816742cb739",
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
	"push-cash"
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
    res, err := s.Events.List(ctx, operations.ListEventsRequest{
        CreatedAtAfter: types.MustTimeFromString("2022-12-25T03:24:03.949Z"),
        CreatedAtBefore: types.MustTimeFromString("2022-05-20T13:30:46.463Z"),
        Cursor: pushcash.String("vjl8vk3l4o8dhsjlzh=="),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListEvents200ApplicationJSONObject != nil {
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

