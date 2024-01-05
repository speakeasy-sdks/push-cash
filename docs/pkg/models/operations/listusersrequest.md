# ListUsersRequest


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `CreatedAtAfter`                                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                                     | :heavy_minus_sign:                                                                             | Return users created after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.  |
| `CreatedAtBefore`                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                                     | :heavy_minus_sign:                                                                             | Return users created before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp. |