// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Speed struct {
	Speed *float32 `json:"speed,omitempty"`
}

func (o *Speed) GetSpeed() *float32 {
	if o == nil {
		return nil
	}
	return o.Speed
}
