package chord

import "github.com/rasmusrygaard/chord/chord"

type FingerTable struct {
	Fingers []Finger
}

type Finger struct {
	Start chord.ID
}
