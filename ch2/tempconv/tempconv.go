package tempconv


import "fmt"

type Celsius 		float64
type Fahrenheit 	float64
type LengthForFoot 	float64
type LengthForMeter 	float64
type WeightForPound 	float64
type WeightForKG    	float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%f°C", f)}
func (foot LengthForFoot) String() string {return fmt.Sprintf("%f", foot)}
func (meter LengthForMeter) String() string {return fmt.Sprintf("%f", meter)}
func (pound WeightForPound) String() string {return fmt.Sprintf("%f", pound)}
func (kg WeightForKG) String() string {return fmt.Sprintf("%f", kg)}
