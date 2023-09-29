# User
(*User*)

### Available Operations

* [CreateUser](#createuser) - Create user
* [GetUser](#getuser) - Get a user
* [List](#list) - List users
* [UpdateUser](#updateuser) - Update a user's status

## CreateUser

Create new Push users

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
    res, err := s.User.CreateUser(ctx, operations.CreateUserRequest{
        CreateUserRequest: shared.CreateUserRequest{
            Address: "4290 Kutch Courts",
            Email: "Americo26@gmail.com",
            Kyc: shared.Kyc{
                Method: "Cargo silver Toys",
                Token: "connect projection Plastic",
            },
            Name: "Bespoke",
            Phone: "(712) 600-8072 x0568",
            Tag: "rudely utilize",
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


## List

Retrieves a list of users

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
    res, err := s.User.List(ctx, operations.ListUsersRequest{
        CreatedAtAfter: types.MustTimeFromString("2023-04-22T14:13:10.937Z"),
        CreatedAtBefore: types.MustTimeFromString("2021-04-20T09:15:20.369Z"),
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
    res, err := s.User.UpdateUser(ctx, operations.UpdateUserRequest{
        UpdateUserRequest: &shared.UpdateUserRequest{
            Status: shared.UpdateUserRequestStatusActive,
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

