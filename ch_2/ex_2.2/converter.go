// Converter converts numeric argument to different measurements
package main

import (
	"fmt"
	"os"
	"strconv"

	"temp/tempconv"
)

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		ft := Feet(t)
		m := Meter(t)
		p := Pound(t)
		k := Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", ft, FtToM(ft), m, MToFt(m))
		fmt.Printf("%s = %s, %s = %s\n", p, PToK(p), k, KToP(k))
	}
}

func FtToM(val Feet) Meter { return Meter(val * 0.3048) }

func MToFt(val Meter) Feet { return Feet(val / 0.3048) }

func PToK(val Pound) Kilogram { return Kilogram(val / 2.2046) }

func KToP(val Kilogram) Pound { return Pound(val * 2.2046) }

func (f Feet) String() string { return fmt.Sprintf("%gft", f) }

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func (p Pound) String() string { return fmt.Sprintf("%glbs", p) }

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
