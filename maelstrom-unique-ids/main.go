package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const (
	minChar           = 33
	maxChar           = 126
	machineIDLen      = 6
	incrementRollOver = 1000000
	idFormat          = "%s-%d-%d"
)

type autoIncrementID struct {
	sync.Mutex
	last       int
	lastRolled int64
}

func (a *autoIncrementID) getAndIncrease(now int64) int {
	a.Lock()
	defer a.Unlock()

	if now == a.lastRolled {
		time.Sleep(1 * time.Millisecond)
	}

	num := a.last

	newNum := a.last + 1

	if newNum > incrementRollOver {
		newNum = 0
	}

	a.last = newNum

	return num
}

func main() {
	machineID := generateMachineID()

	aid := &autoIncrementID{}

	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		t := time.Now().UnixMilli()

		id := fmt.Sprintf(idFormat, machineID, t, aid.getAndIncrease(t))

		return n.Reply(msg, map[string]any{
			"type": "generate_ok",
			"id":   id,
		})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func generateMachineID() string {
	machineIDBytes := make([]byte, machineIDLen)

	for i := range machineIDLen {
		machineIDBytes[i] = byte(rand.Intn(maxChar-minChar) + minChar)
	}

	machineID := string(machineIDBytes)

	return machineID
}
