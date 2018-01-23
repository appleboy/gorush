# Sample Function

The following is a sample Lambda function that transforms Kinesis Firehose records by doing a ToUpper on the data.

```go

import (
    "fmt"
    "strings"
    "github.com/aws/aws-lambda-go/events"
)

func handleRequest(evnt events.KinesisFirehoseEvent) events.KinesisFirehoseResponse {

    fmt.Printf("InvocationId: %s\n", evnt.InvocationId)
    fmt.Printf("DeliveryStreamArn: %s\n", evnt.DeliveryStreamArn)
    fmt.Printf("Region: %s\n", evnt.Region)

    var response events.KinesisFirehoseResponse

    for _, record := range evnt.Records {
        fmt.Printf("RecordId: %s\n", record.RecordId)
        fmt.Printf("ApproximateArrivalTimestamp: %s\n", record.ApproximateArrivalTimestamp)

        // Transform data: ToUpper the data
        var transformedRecord kinesisfhevents.FirehoseResponseRecord
        transformedRecord.RecordId = record.RecordId
        transformedRecord.Result = kinesisfhevents.TransformedStateOk
        transformedRecord.Data = strings.ToUpper(string(record.Data))

        response.Records = append(response.Records, transformedRecord)
    }

    return response
}
```
