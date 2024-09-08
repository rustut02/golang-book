package weightconv

import "fmt"

type Kilos float64
type Pounds float64

const (
	weightCoefficient = 0.453592
)

func (k Kilos) String() string {
	return fmt.Sprintf("%.2f kg", k)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%.2f lbs", p)
}
