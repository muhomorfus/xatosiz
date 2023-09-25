package jsonenc

import "encoding/json"

type Encoder struct {
	v any

	encoded []byte
	err     error
}

func New(v any) *Encoder {
	return &Encoder{v: v}
}

func (e *Encoder) ensureEncoded() {
	if e.encoded == nil && e.err == nil {
		e.encoded, e.err = json.Marshal(e.v)
	}
}

func (e *Encoder) Length() int {
	e.ensureEncoded()
	return len(e.encoded)
}

func (e *Encoder) Encode() ([]byte, error) {
	e.ensureEncoded()
	return e.encoded, e.err
}
