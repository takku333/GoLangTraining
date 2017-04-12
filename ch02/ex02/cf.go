// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

// type flags struct {
//         convTemp *bool
//         convLength *bool
//         convWeight *bool
// }

func convTemp(val float64) (tempconv.Celsius, tempconv.Fahrenheit) {
	c := tempconv.Celsius(val)
	f := tempconv.Fahrenheit(val)
	return tempconv.FToC(f), tempconv.CToF(c)
}

func convLength(val float64) (meter, feet) {
	fe := feet(val)
	met := meter(val)
	return feetToMeter(fe), meterToFeet(met)
}

func convWeight(val float64) (kgram, pound) {
	po := pound(val)
	kg := kgram(val)
	return poundToKgram(po), kgramToPound(kg)
}

func meterToFeet(m meter) feet { return feet(m * 3.2808) }

func feetToMeter(f feet) meter { return meter(f / 3.2808) }

func kgramToPound(k kgram) pound { return pound(k / 2.20462262) }

func poundToKgram(p pound) kgram { return kgram(p * 2.20462262) }

func main() {
	// fs := flags{
	//     convTemp: flag.Bool("convTemp", false, "convert Temperature"),
	//     convLength: flag.Bool("convLength", false, "convert length"),
	//     convWeight: flag.Bool("convWeight", false, "convert Weight"),
	// }

	// flag.Parse()
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if input.Text() == "q" {
				break
			}
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ConvUnit: %v\n", err)
				os.Exit(1)
			}
			c, f := convTemp(t)
			fmt.Printf("%s = %s, %s = %s\n",
				c, tempconv.Fahrenheit(t), f, tempconv.Celsius(t))

			met, fe := convLength(t)
			fmt.Printf("%s = %s, %s = %s\n",
				met, feet(t), fe, meter(t))

			kg, po := convWeight(t)
			fmt.Printf("%s = %s, %s = %s\n",
				kg, pound(t), po, kgram(t))
		}
	} else {
		for _, arg := range args {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ConvUnit: %v\n", err)
				os.Exit(1)
			}
			c, f := convTemp(t)
			fmt.Printf("%s = %s, %s = %s\n",
				c, tempconv.Fahrenheit(t), f, tempconv.Celsius(t))

			met, fe := convLength(t)
			fmt.Printf("%s = %s, %s = %s\n",
				met, feet(t), fe, meter(t))

			kg, po := convWeight(t)
			fmt.Printf("%s = %s, %s = %s\n",
				kg, pound(t), po, pound(t))
		}
	}

}

//各flagを確認し、全てfalseだった場合は、全てのflagをtrueに変更する
// func checkFlags (fs flags) {

// }
