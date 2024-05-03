package types

import "sync"

type Trigger interface {
	Activate(agg chan<- Event, quit <-chan bool, wg *sync.WaitGroup) error
}

type Channel = ChannelTrigger

type ChannelTrigger struct {
	Source <-chan Event
}

var _ Trigger = ChannelTrigger{}

func (t ChannelTrigger) Activate(agg chan<- Event, quit <-chan bool, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case e := <-t.Source:
				agg <- e
			case <-quit:
				return
			}
		}
	}()
	return nil
}
