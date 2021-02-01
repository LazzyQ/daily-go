package balance

type Node struct {
	Name    string
	Current int
	Weight  int
}

// 出处: https://tenfy.cn/2018/11/12/smooth-weighted-round-robin/
func SmoothWrr(nodes []*Node) (best *Node) {
	if len(nodes) == 0 {
		return
	}

	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}
		total += node.Weight
		node.Current += node.Weight
		if best == nil || node.Current > best.Current {
			best = node
		}
	}
	if best == nil {
		return
	}
	best.Current -= total
	return
}
