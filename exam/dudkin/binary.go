package dudkin

type Node struct{
	v int
	l *Node
	r *Node
}

func (n *Node)FindSmallest(p int) int {
	if n == nil {
		return p
	}

	v := n.v
	if v < p {
		p = v
	}

	vl := n.l.FindSmallest(p)
	vr := n.r.FindSmallest(p)

	if vl > vr {
		return vr
	} else {
		return vl
	}
}
