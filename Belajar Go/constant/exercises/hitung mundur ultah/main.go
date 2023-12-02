package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
	//masukkan hari ulang tahunmu
    birthday := time.Date(now.Year(), time.Month(2), 15, 0, 0, 0, 0, time.UTC)
	//masukkan hari ulang tahunmu
    if now.After(birthday) {
        birthday = time.Date(now.Year()+1, time.Month(2), 15, 0, 0, 0, 0, time.UTC)
    }
    daysUntilBirthday := birthday.Sub(now).Hours() / 24

    fmt.Printf("There are %.0f days until your next birthday.\n", daysUntilBirthday)
}