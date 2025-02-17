package stack

type Priority int

const (
	Low Priority = iota
	Middle
	High
)

type PriorityStack struct {
	low    *Stack
	middle *Stack
	high   *Stack
}

func NewPriorityStack(low, middle, high []any) *PriorityStack {
	return &PriorityStack{
		low:    NewStack(low),
		middle: NewStack(middle),
		high:   NewStack(high),
	}
}

func (ps *PriorityStack) Enqueue(item any, priority Priority) error {
	switch priority {
	case Low:
		return ps.low.Enqueue(item)
	case Middle:
		return ps.middle.Enqueue(item)
	case High:
		return ps.high.Enqueue(item)
	}

	return ErrUnknownPriority
}
