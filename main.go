package main

import "fmt"

// --------------------= GLOBAL =--------------------

const MaxLogs int = 999

// --------------------= UTILITY =--------------------

func digitsToString(number int) string {
	var lastDigit int
	var result string

	if number == 0 {
		result = ""
	} else {
		lastDigit = number % 10
		result = digitsToString(number/10) + string(byte(lastDigit)+'0')
	}

	return result
}

func IntToStr(number int) string {
	var isNegative bool
	var result string

	if number == 0 {
		result = "0"
	} else {
		if number < 0 {
			isNegative = true
			number = -number
		}

		result = digitsToString(number)

		if isNegative {
			result = "-" + result
		}
	}

	return result
}

func FloatToStr(value float64) string {
	var result, precisionStr string
	var valueInt, precisionInt int
	var isNegative bool

	if value < 0 {
		isNegative = true
		value = -value
	}

	valueInt = int(value)
	precisionInt = int((value - float64(valueInt)) * 100)

	if precisionInt == 0 {
		precisionStr = "00"
	} else {
		if precisionInt < 10 {
			precisionStr = "0" + digitsToString(precisionInt)
		} else {
			precisionStr = digitsToString(precisionInt)
		}
	}

	result = IntToStr(valueInt) + "." + precisionStr

	if isNegative {
		result = "-" + result
	}

	return result
}

// --------------------= DURATION =--------------------

type Duration struct {
	hours   int
	minutes int
	seconds int
}

func NewDuration(hours int, minutes int, seconds int) Duration {
	var duration Duration

	duration.hours = hours
	duration.minutes = minutes
	duration.seconds = seconds

	return duration
}

func DurationToStr(duration Duration) string {
	var result string

	if duration.hours > 0 {
		result = result + IntToStr(duration.hours) + " Jam "
	}

	if duration.minutes > 0 {
		result = result + IntToStr(duration.minutes) + " Menit "
	}

	result = result + IntToStr(duration.seconds) + " Detik"

	return result
}

// --------------------= POWER =--------------------

type Power float64

func PowerToStr(power Power) string {
	return FloatToStr(float64(power)) + " Watt"
}

// --------------------= LOG =--------------------

type Log struct {
	deviceName     string
	deviceLocation string
	power          Power
	duration       Duration
}

type Logs struct {
	list [MaxLogs]Log
	n    int
}

func AddLog(logs *Logs, deviceName string, deviceLocation string, power Power, duration Duration) {
	var log Log

	log.deviceName = deviceName
	log.deviceLocation = deviceLocation
	log.power = power
	log.duration = duration

	logs.list[logs.n] = log

	logs.n = logs.n + 1
}

// --------------------= TABLE =--------------------

func PrintTableHead() {
	fmt.Println("╭─────┬──────────────────────────┬─────────────────────────┬───────────────────────────────────────────────────┬─────────────────╮")
	fmt.Println("│ No. │      Nama Perangkat      │     Lokasi Perangkat    │                  Durasi Pemakaian                 │ Komsumsi Energi │")
}

func PrintTableSeparator() {
	fmt.Println("├─────┼──────────────────────────┼─────────────────────────┼───────────────────────────────────────────────────┼─────────────────┤")
}

func PrintTableFoot() {
	fmt.Println("╰─────┴──────────────────────────┴─────────────────────────┴───────────────────────────────────────────────────┴─────────────────╯")
}

func PrintLogTable(logs Logs) {
	var index int
	var currentLog Log

	PrintTableHead()

	for index = 1; index <= logs.n; index++ {
		PrintTableSeparator()

		currentLog = logs.list[index-1]

		fmt.Printf("│ %-3d │ %-24s │ %-23s │ %-49s │ %-15s │\n",
			index,
			currentLog.deviceName,
			currentLog.deviceLocation,
			DurationToStr(currentLog.duration),
			PowerToStr(currentLog.power),
		)
	}

	PrintTableFoot()
}

// --------------------= MAIN =--------------------

func main() {
	var logs Logs

	AddLog(&logs, "Laptop", "Kamar_Tidur", 123.23, NewDuration(1, 20, 10))
	AddLog(&logs, "Kipas_Angin", "Ruang_Tamu", 45.50, NewDuration(4, 0, 0))
	AddLog(&logs, "Televisi", "Ruang_Keluarga", 110.00, NewDuration(2, 30, 0))

	PrintLogTable(logs)
}
