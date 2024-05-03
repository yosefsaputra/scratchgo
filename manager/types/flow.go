package types

type Flow struct {
	Triggers []Trigger
	Actions  []Action

	AggregateCh chan Event
}

type flowBuilder struct {
	flow Flow
}

func NewFlow() *flowBuilder {
	return &flowBuilder{
		flow: Flow{
			AggregateCh: make(chan Event),
		},
	}
}

func (f *flowBuilder) WithTriggers(t ...Trigger) *flowBuilder {
	f.flow.Triggers = append(f.flow.Triggers, t...)
	return f
}

func (f *flowBuilder) WithActions(a ...Action) *flowBuilder {
	f.flow.Actions = append(f.flow.Actions, a...)
	return f
}

func (f *flowBuilder) Build() Flow {
	return f.flow
}
