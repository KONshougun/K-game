package effect

type TriggerTiming byte

const (
	None TriggerTiming = iota
	Before
	//During
	//...
	After
)
