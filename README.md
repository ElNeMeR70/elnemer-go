# elnemer-go

Sovereign Multi-Agent Swarm & Arabic NLP Engine for Go.
Architected by Eng. Mahmoud Bek ElNemer (@ElNeMeR70).

## Installation
```bash
go get github.com/ElNeMeR70/elnemer-go
```

## Quick Start
```go
package main

import (
	"fmt"
	"github.com/ElNeMeR70/elnemer-go"
)

func main() {
	// Arabic NLP
	clean := elnemer.StripTashkeel("مُحَمَّدٌ")
	fmt.Println("Clean:", clean)
	fmt.Println("Jummal:", elnemer.JummalValue("محمد")) // 92

	// Multi-Agent Swarm
	swarm := elnemer.NewSwarm("ProductionCluster")
	swarm.AddAgent(elnemer.NewAgent("Auditor", "Security"))
	results := swarm.Broadcast("Audit architecture")
	fmt.Printf("%+v\n", results)
}
```

## License
MIT
