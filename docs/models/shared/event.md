# Event


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CreatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | when the event occurred                                                |
| `EventType`                                                            | [shared.EventType](../../models/shared/eventtype.md)                   | :heavy_check_mark:                                                     | the type of event indicates how the status of the resource has changed |
| `ID`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | the identifier for the event                                           |
| `SourceID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | the identifier of the resource who's status has updated                |
| `SourceType`                                                           | [shared.SourceType](../../models/shared/sourcetype.md)                 | :heavy_check_mark:                                                     | the type of the resource which has been updated                        |