package main

import (
	"context"
	"fmt"

	"github.com/sam-bee/security-hackerone-api-client"
)

func main() {
	h1 := hackerone.New("your-username-here", "your-api-key-here")
	reports, _, _ := h1.Hackers.GetReports(context.TODO(), nil)
	for _, report := range reports {
		fmt.Println(report.Id)
	}
}
