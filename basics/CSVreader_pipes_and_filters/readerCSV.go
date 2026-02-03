package main

import (
	"bufio"
	"fmt"
	"os"
)

type SensorData struct {
	Hour        string
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func ReadFile(data string, lines chan byte) {
	var c byte

	file, err := os.Open(data)
	if err != nil {
		fmt.Println("Virhe:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		c, err = reader.ReadByte()
		if err != nil {
			break
		}
		lines <- c
	}
	close(lines)
}

func ParseLines(lines chan byte, results chan SensorData) {
	var skipHeader = true
	var fieldIndex int
	var buffer []byte
	var hourKey string
	var temp, hum, pres float64
	var neg bool
	var frac bool
	var fracDiv float64
	var value float64

	for c := range lines {
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
					results <- SensorData{
						Hour:        hourKey,
						Temperature: temp,
						Humidity:    hum,
						Pressure:    pres,
					}
					fieldIndex = 0
					// lineStart = true
				}

				continue
			}

			// lineStart = false
		}
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
	close(results)
}

func main() {
	lines := make(chan byte)
	results := make(chan SensorData)

	go ReadFile("data.csv", lines)
	go ParseLines(lines, results)

	tempSum := make(map[string]float64)
	humSum := make(map[string]float64)
	presSum := make(map[string]float64)
	count := make(map[string]int)

	for data := range results {
		tempSum[data.Hour] += data.Temperature
		humSum[data.Hour] += data.Humidity
		presSum[data.Hour] += data.Pressure
		count[data.Hour]++
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
