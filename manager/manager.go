package manager

import (
	"fmt"
	"sync"

	"github.com/yosefsaputra/scratchgo/manager/types"
)

func NewManager(flows []types.Flow, quit chan bool) Manager {
	return Manager{
		Flows: flows,
		quit:  quit,
	}
}

type Manager struct {
	Flows []types.Flow
	quit  chan bool
}

func (m *Manager) Start() error {
	var triggerWg sync.WaitGroup
	var actionWg sync.WaitGroup

	for _, flow := range m.Flows {
		for _, trigger := range flow.Triggers {
			err := trigger.Activate(flow.AggregateCh, m.quit, &triggerWg) // todo: timeout
			if err != nil {
				panic(err) // todo: don't panic
			}
		}

		actionWg.Add(1)
		go func(flow types.Flow) {
			defer actionWg.Done()
			for {
				select {
				case <-m.quit:
					return
				case <-flow.AggregateCh:
					for _, action := range flow.Actions {
						err := action.Run() // todo: timeout
						if err != nil {
							fmt.Printf("error running action: %v\n", err)
						}
					}
				}
			}
		}(flow)
	}

	triggerWg.Wait()
	actionWg.Wait()

	return fmt.Errorf("not implemented")
}

func (m *Manager) RegisterFlow(flow types.Flow) {
	m.Flows = append(m.Flows, flow)
}
