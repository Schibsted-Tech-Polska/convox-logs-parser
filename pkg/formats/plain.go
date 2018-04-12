package formats

// PlainMessage is a structure for standard plain or unformatted text messages
type PlainMessage struct {
	message string
}

// NewPlainMessage is a PlainMessage factory
func NewPlainMessage(message string) PlainMessage {
	pm := PlainMessage{message: message}
	return pm
}

func (pm PlainMessage) String() string {
	return pm.message
}
