package main
import (
	"fmt"
	"net"
	"sync"
)

// we've used WaitGroup to the synchronized scanner
// which used different implementation of the goroutines

func main() {
	var wg sync.WaitGroup // creating a synchronized counter
	for i:= 1; i<= 65535; i++ {
	wg.Add(1) // create a goroutine to scan a port, defered call to wg.Done
	go func(j int) {
	defer wg.Done() // decrements the counter whenever a unit of work is done
	address:= fmt.Sprintf("scanme.nmap.org:%d",j)
	conn, err := net.Dial("tcp",address)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Printf("%d open\n",j)
	
	}(i)
	}
wg.Wait() // blocks until the work is done, then our counter returns to zero
}

