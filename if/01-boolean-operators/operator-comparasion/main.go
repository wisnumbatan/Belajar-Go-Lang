package main

import "fmt"

func main() {
    speed := 100 // Mendefinisikan variabel 'speed' dengan nilai 100 (mil per jam)

    // Memeriksa apakah kecepatan lebih besar atau sama dengan 80
    fast := speed >= 80

    // Memeriksa apakah kecepatan kurang dari 20
    slow := speed < 20

    // Mencetak tipe data dari variabel 'fast' (harusnya boolean)
    fmt.Printf("fast's type is %T\n", fast)

    // Mencetak apakah sedang bergerak cepat
    fmt.Printf("going fast? %t\n", fast)

    // Mencetak apakah sedang bergerak lambat
    fmt.Printf("going slow? %t\n", slow)

    // Mencetak apakah kecepatan adalah 100 mph
    fmt.Printf("is it 100 mph? %t\n", speed == 100)

    // Mencetak apakah kecepatan bukan 100 mph
    fmt.Printf("is it not 100 mph? %t\n", speed != 100)
}
