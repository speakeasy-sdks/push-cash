# Event

Successful operation


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CreatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | when the event occurred                                                |
| `EventType`                                                            | [EventEventType](../../models/shared/eventeventtype.md)                | :heavy_check_mark:                                                     | the type of event indicates how the status of the resource has changed |
| `ID`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | the identifier for the event                                           |
| `SourceID`                                                             | *string*                                                               | :heavy_check_mark:                                                     | the identifier of the resource who's status has updated                |
| `SourceType`                                                           | [EventSourceType](../../models/shared/eventsourcetype.md)              | :heavy_check_mark:                                                     | the type of the resource which has been updated                        |