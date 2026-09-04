package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()

    newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())

    diff := newYear.Sub(now)
    
    days := int(diff.Hours() / 24)
    
    fmt.Printf("до НГ осталось %d дней\n", days)
}
