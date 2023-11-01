// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ChargeVoltage struct {
	// The voltage of the charger measured by the vehicle.
	Voltage *float32 `json:"voltage,omitempty"`
}

func (o *ChargeVoltage) GetVoltage() *float32 {
	if o == nil {
		return nil
	}
	return o.Voltage
}
