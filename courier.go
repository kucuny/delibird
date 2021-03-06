package delibird

import (
	"fmt"
	"reflect"
)

// Courier is the interface representing the standardized methods to
// parse shipment tracking html
type Courier interface {
	// Parse html to tracking object
	Parse(invoice string) (Track, *ApiError)
	// Courier code
	Code() string
	// Courier name
	Name() string
}

var courierMap = map[string]Courier{}

// NewCourier creates courier object by courier company code
func NewCourier(name string) (Courier, *ApiError) {
	if value, ok := courierMap[name]; ok {
		courier := reflect.New(reflect.TypeOf(value).Elem()).Interface().(Courier)
		return courier, nil
	}

	return nil, NewApiError(NoCode, fmt.Sprintf("%v is not supported.", name))
}

// RegisterCourier register new courier
func RegisterCourier(name string, courier Courier) {
	courierMap[name] = courier
}
