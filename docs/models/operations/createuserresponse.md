# CreateUserResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |
| `User`                                                 | [*shared.User](../../models/shared/user.md)            | :heavy_minus_sign:                                     | User created successfully                              |
| `Error`                                                | [*shared.Error](../../models/shared/error.md)          | :heavy_minus_sign:                                     | Error                                                  |