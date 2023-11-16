# GetBalanceResponse


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AccountBalance`                                                       | [*shared.AccountBalance](../../../pkg/models/shared/accountbalance.md) | :heavy_minus_sign:                                                     | Account balance retrieved successfully                                 |
| `ContentType`                                                          | *string*                                                               | :heavy_check_mark:                                                     | HTTP response content type for this operation                          |
| `StatusCode`                                                           | *int*                                                                  | :heavy_check_mark:                                                     | HTTP response status code for this operation                           |
| `RawResponse`                                                          | [*http.Response](https://pkg.go.dev/net/http#Response)                 | :heavy_check_mark:                                                     | Raw HTTP response; suitable for custom response parsing                |
| `Error`                                                                | [*shared.Error](../../../pkg/models/shared/error.md)                   | :heavy_minus_sign:                                                     | Error                                                                  |