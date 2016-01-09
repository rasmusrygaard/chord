package chord

type Ring struct {
	Known Node
}

func (r Ring) Lookup(id ID) (Node, error) {
	r.Known.FindSuccessor(id)
}
