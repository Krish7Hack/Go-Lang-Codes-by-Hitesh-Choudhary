package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("Welcome to our pizza app")
    fmt.Println("Please rate our pizza between 1 and 5")

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n') // Read string until you hit \n

    // Trim newline characters from the input
    input = strings.TrimSpace(input)

    fmt.Println("Thanks for rating,", input)

    numRating, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println("Error parsing rating:", err)
    } else {
        fmt.Println("Added 1 to your rating:", numRating+1)
    }
}
