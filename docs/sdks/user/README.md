# User

### Available Operations

* [CreateUser](#createuser) - Create user
* [GetUser](#getuser) - Get a user
* [ListUsers](#listusers) - List users
* [UpdateUser](#updateuser) - Update a user's status

## CreateUser

Create new Push users

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
    res, err := s.User.CreateUser(ctx, operations.CreateUserRequest{
        CreateUserRequest: shared.CreateUserRequest{
            Address: "822 Grant Oval",
            Email: "Aiyana.Cummerata0@yahoo.com",
            Kyc: shared.Kyc{
                Method: shared.KYCMethodAlloy,
                Token: "fugit",
            },
            Name: "Marshall Glover",
            Phone: "942-853-5856 x289",
            Tag: "dolorum",
        },
        XIdempotencyKey: "f1bbb856-fb17-11ed-be56-0242ac120002",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateUserRequest](../../models/operations/createuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**


## GetUser

Retrieve a user by ID

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
    res, err := s.User.GetUser(ctx, operations.GetUserRequest{
        ID: "user_28CJjV7P4Go5PNJvfzghiD",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetUserRequest](../../models/operations/getuserrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetUserResponse](../../models/operations/getuserresponse.md), error**


## ListUsers

Retrieves a list of users

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
    res, err := s.User.ListUsers(ctx, operations.ListUsersRequest{
        CreatedAtAfter: types.MustTimeFromString("2022-07-21T01:01:39.776Z"),
        CreatedAtBefore: types.MustTimeFromString("2020-01-25T11:09:22.009Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListUsers200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListUsersRequest](../../models/operations/listusersrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**


## UpdateUser

Updates a user's status to either active or suspended

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
    res, err := s.User.UpdateUser(ctx, operations.UpdateUserRequest{
        UpdateUserRequest: &shared.UpdateUserRequest{
            Status: shared.UpdateUserRequestStatusSuspended,
        },
        ID: "user_28CJjV7P4Go5PNJvfzghiD",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateUserRequest](../../models/operations/updateuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

