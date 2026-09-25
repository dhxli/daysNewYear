package main

import (
	"fmt"
	"time"
	"github.com/dhxli/daysNewYear/internal/daysleft"
)

func main() {
	days := daysleft.Until(time.Now())
	fmt.Printf("До НГ осталось %d дней\n", days)
}