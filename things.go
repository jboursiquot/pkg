package things

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Thing is a thing.
type Thing struct {
	Name string
}

func (t Thing) String() string {
	return fmt.Sprintf("Hello, I'm %s", t.Name)
}

// DoHardThings will do hard things for a while or until you cancel the context.
func DoHardThings(ctx context.Context) <-chan struct{} {
	log.Println("things.DoHardThings >> called...")

	// to simulate hard work for 2 seconds
	d := 2 * time.Second
	completeAfter := time.After(d)

	ch := make(chan struct{})
	thingCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	for {
		select {
		case <-thingCtx.Done():
			err := thingCtx.Err()
			if err == context.Canceled {
				log.Println("things.DoHardThings >> context was canceled")
			} else if err == context.DeadlineExceeded {
				log.Println("things.DoHardThings >> context deadline was exceeded")
			}
			return ch
		case <-completeAfter:
			log.Printf("things.DoHardThings >> did hard things for %v", d)
			return ch
		default:
			log.Println("things.DoHardThings >> doing hard things...")
			time.Sleep(1 * time.Second)
		}
	}
}
