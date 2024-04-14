# hackerone

HackerOne API Client in Go.

This project is a fork of [liamg/hackerone](https://github.com/liamg/hackerone/). It adds the functionality to search through the Structured
Scope objects in the API. This addition allows you to automate the process of bug bounty target scanning, and large-scale automated security
testing. See the [sam-bee/security-hackerone-target-retrieval](https://github.com/sam-bee/security-hackerone-target-retrieval) project for real world usage.

## Usage Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/liamg/hackerone"
)

func main() {
	h1 := hackerone.New("your-username-here", "your-api-key-here")
	reports, _, _ := h1.Hackers.GetReports(context.TODO(), nil)
	for _, report := range reports {
		fmt.Println(report.Id)
	}
}
```
