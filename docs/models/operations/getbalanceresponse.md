# GetBalanceResponse


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AccountBalance`                                                | [*shared.AccountBalance](../../models/shared/accountbalance.md) | :heavy_minus_sign:                                              | Account balance retrieved successfully                          |
| `ContentType`                                                   | *string*                                                        | :heavy_check_mark:                                              | N/A                                                             |
| `StatusCode`                                                    | *int*                                                           | :heavy_check_mark:                                              | N/A                                                             |
| `RawResponse`                                                   | [*http.Response](https://pkg.go.dev/net/http#Response)          | :heavy_minus_sign:                                              | N/A                                                             |
| `Error`                                                         | [*shared.Error](../../models/shared/error.md)                   | :heavy_minus_sign:                                              | Error                                                           |