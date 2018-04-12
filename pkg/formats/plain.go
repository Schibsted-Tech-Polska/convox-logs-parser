package formats

type PlainMessage struct {
	message string
}

func NewPlainMessage(message string) PlainMessage {
	pm := PlainMessage{message: message}
	return pm
}

func (pm PlainMessage) String() string {
	return pm.message
}
