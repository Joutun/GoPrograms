package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Virhe:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	tempSum := make(map[string]float64)
	humSum := make(map[string]float64)
	presSum := make(map[string]float64)
	count := make(map[string]int)

	var c byte
	// var lineStart = true
	var skipHeader = true

	var fieldIndex int
	var buffer []byte
	var hourKey string

	var temp, hum, pres float64
	var neg bool
	var frac bool
	var fracDiv float64
	var value float64

	for {
		c, err = reader.ReadByte()
		if err != nil {
			break
		}

		if skipHeader {
			if c == '\n' {
				skipHeader = false
			}
			continue
		}

		if c == ',' || c == '\n' {
			// Käsittele kenttä
			if fieldIndex == 0 && len(buffer) >= 13 {
				hourKey = string(buffer[:13])
			}

			if fieldIndex == 1 || fieldIndex == 2 || fieldIndex == 3 {
				if neg {
					value = -value
				}
				if fieldIndex == 1 {
					temp = value
				} else if fieldIndex == 2 {
					hum = value
				} else if fieldIndex == 3 {
					pres = value
				}
			}

			// Reset parsintamuuttujat
			buffer = buffer[:0]
			fieldIndex++
			value = 0
			neg = false
			frac = false
			fracDiv = 1

			if c == '\n' {
				if fieldIndex > 3 {
					tempSum[hourKey] += temp
					humSum[hourKey] += hum
					presSum[hourKey] += pres
					count[hourKey]++
				}
				fieldIndex = 0
				// lineStart = true
			}

			continue
		}

		// lineStart = false

		if fieldIndex <= 3 {
			buffer = append(buffer, c)

			if c == '-' {
				neg = true
				continue
			}

			if c == '.' {
				frac = true
				continue
			}

			if c >= '0' && c <= '9' {
				d := float64(c - '0')
				if !frac {
					value = value*10 + d
				} else {
					fracDiv *= 10
					value += d / fracDiv
				}
			}
		}
	}

	fmt.Println("Tunti | Lämpötila | Kosteus | Ilmanpaine")
	fmt.Println("---------------------------------------")

	for h := range count {
		n := float64(count[h])
		fmt.Printf(
			"%s | %.2f | %.2f | %.2f\n",
			h,
			tempSum[h]/n,
			humSum[h]/n,
			presSum[h]/n,
		)
	}
}

/*

   Tässä esimerkki tiedostosta data.csv.  Kopioi ~ merkien välissä
   olevat rivit tiedostoon data.csv samaan hakemistoon, missä
   esimerkki.go-ohjelma sijaitsee.

~
Päivämäärä,Lämpötila (°C),Kosteus (°C),Ilmanpaine (hPa),RSSI (dBm),Jännite (V),Kiihtyvyys X (g),Kiihtyvyys Y (g),Kiihtyvyys Z (g),Liikelaskuri (liikettä),Lähetysteho (dBm),Mittausjärjestysnumero
2025-06-09 08:04:45,24.64,38.1,995.42,,,,,,,,
2025-06-09 08:09:46,24.65,38.1,995.48,,,,,,,,
2025-06-09 08:14:46,24.65,38.14,995.52,,,,,,,,
2025-06-09 09:04:53,24.71,37.92,995.31,,,,,,,,
2025-06-09 09:09:54,24.72,37.85,995.27,,,,,,,,
2025-06-09 09:14:54,24.73,37.82,995.26,,,,,,,,
2025-06-09 09:19:55,24.74,37.78,995.26,,,,,,,,
2025-06-09 18:24:49,25.65,33.8,992.26,-55,2.549,−0.072,−0.004,1.004,119,4,5699
2025-06-09 18:24:54,25.65,33.79,992.26,-63,2.549,−0.072,0.004,0.996,119,4,5701
2025-06-09 18:24:57,25.65,33.79,992.26,-54,2.549,−0.076,0.004,0.996,119,4,5702
2025-06-09 18:24:59,25.65,33.78,992.26,-58,2.549,−0.072,0.004,0.996,119,4,5703
2025-06-09 18:25:02,25.65,33.81,992.26,-57,2.549,−0.076,0,1,119,4,5704
2025-06-09 18:25:08,25.65,33.77,992.26,-56,2.549,−0.068,0,1,119,4,5706
2025-06-09 18:25:13,25.65,33.78,992.26,-58,2.549,−0.072,0,0.996,119,4,5708
2025-06-09 18:25:16,25.65,33.8,992.25,-58,2.549,−0.068,0.004,0.996,119,4,5709
~
*/
