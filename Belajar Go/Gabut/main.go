package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jarak dalam meter: ")
	jarakStr, _ := reader.ReadString('\n')
	jarakStr = strings.TrimSpace(jarakStr)
	jarak, err := strconv.ParseFloat(jarakStr, 64)
	if err != nil {
		fmt.Println("Input tidak valid untuk jarak. Mohon masukkan angka yang valid.")
		return
	}

	fmt.Print("Masukkan waktu dalam detik: ")
	waktuStr, _ := reader.ReadString('\n')
	waktuStr = strings.TrimSpace(waktuStr)
	waktu, err := strconv.ParseFloat(waktuStr, 64)
	if err != nil {
		fmt.Println("Input tidak valid untuk waktu. Mohon masukkan angka yang valid.")
		return
	}

	if waktu == 0 {
		fmt.Println("Waktu tidak boleh nol. Mohon masukkan waktu yang valid.")
		return
	}

	kecepatan := jarak / waktu

	fmt.Printf("Kecepatan angin adalah %.2f m/s.\n", kecepatan)
}
