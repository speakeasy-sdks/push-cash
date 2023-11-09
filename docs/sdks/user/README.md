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
	pushcash "push-cash/v2"
	"push-cash/v2/pkg/models/shared"
	"push-cash/v2/pkg/models/operations"
)

func main() {
    s := pushcash.New(
        pushcash.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.User.CreateUser(ctx, operations.CreateUserRequest{
        CreateUserRequest: shared.CreateUserRequest{
            Address: "4290 Kutch Courts",
            Email: "Americo26@gmail.com",
            Kyc: shared.Kyc{
                Method: shared.MethodAlloy,
                Token: "string",
            },
            Name: "string",
            Phone: "1-801-391-5854 x37232",
            Tag: "string",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateUserRequest](../../pkg/models/operations/createuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateUserResponse](../../pkg/models/operations/createuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetUser

Retrieve a user by ID

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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetUserRequest](../../pkg/models/operations/getuserrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetUserResponse](../../pkg/models/operations/getuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

Retrieves a list of users

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
    res, err := s.User.List(ctx, operations.ListUsersRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListUsersRequest](../../pkg/models/operations/listusersrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListUsersResponse](../../pkg/models/operations/listusersresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateUser

Updates a user's status to either active or suspended

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
    res, err := s.User.UpdateUser(ctx, operations.UpdateUserRequest{
        UpdateUserRequest: &shared.UpdateUserRequest{
            Status: shared.StatusActive,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateUserRequest](../../pkg/models/operations/updateuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateUserResponse](../../pkg/models/operations/updateuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
