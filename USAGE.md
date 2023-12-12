<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"testsdkcreation"
	"testsdkcreation/pkg/models/shared"
)

func main() {
	s := testsdkcreation.New(
		testsdkcreation.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	var vehicleID string = "36ab27d0-fd9d-4455-823a-ce30af709ffc"

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
<!-- End SDK Example Usage [usage] -->