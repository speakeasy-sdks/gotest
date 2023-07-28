<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
	"testsdkcreation/pkg/models/operations"
)

func main() {
    s := testsdkcreation.New(
        testsdkcreation.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )
    vehicleID := "36ab27d0-fd9d-4455-823a-ce30af709ffc"

    ctx := context.Background()
    res, err := s.Vehicles.GetLocation(ctx, vehicleID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Location != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->