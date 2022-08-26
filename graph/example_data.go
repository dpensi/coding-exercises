package graph

func GetGraph1() node[string] {
	nodeA := node[string]{Value: "a"}
	nodeB := node[string]{Value: "b"}
	nodeC := node[string]{Value: "c"}
	nodeD := node[string]{Value: "d"}
	nodeE := node[string]{Value: "e"}
	nodeF := node[string]{Value: "f"}
	nodeG := node[string]{Value: "g"}
	nodeH := node[string]{Value: "h"}
	nodeI := node[string]{Value: "i"}
	nodeJ := node[string]{Value: "j"}

	nodeA.AddArc(&nodeB)
	nodeA.AddArc(&nodeF)
	nodeA.AddArc(&nodeG)

	nodeB.AddArc(&nodeA)
	nodeB.AddArc(&nodeF)
	nodeB.AddArc(&nodeE)
	nodeB.AddArc(&nodeD)
	nodeB.AddArc(&nodeC)

	nodeC.AddArc(&nodeB)
	nodeC.AddArc(&nodeD)

	nodeD.AddArc(&nodeB)
	nodeD.AddArc(&nodeC)

	nodeE.AddArc(&nodeF)
	nodeE.AddArc(&nodeB)

	nodeF.AddArc(&nodeA)
	nodeF.AddArc(&nodeB)
	nodeF.AddArc(&nodeE)

	nodeG.AddArc(&nodeA)
	nodeG.AddArc(&nodeH)
	nodeG.AddArc(&nodeI)
	nodeG.AddArc(&nodeJ)

	nodeH.AddArc(&nodeG)
	nodeH.AddArc(&nodeI)

	nodeI.AddArc(&nodeG)
	nodeI.AddArc(&nodeH)

	nodeJ.AddArc(&nodeG)

	return nodeA
}

func GetStatusGraph1() statusNode[string] {
	nodeA := statusNode[string]{Value: "a"}
	nodeB := statusNode[string]{Value: "b"}
	nodeC := statusNode[string]{Value: "c"}
	nodeD := statusNode[string]{Value: "d"}
	nodeE := statusNode[string]{Value: "e"}
	nodeF := statusNode[string]{Value: "f"}
	nodeG := statusNode[string]{Value: "g"}
	nodeH := statusNode[string]{Value: "h"}
	nodeI := statusNode[string]{Value: "i"}
	nodeJ := statusNode[string]{Value: "j"}

	nodeA.AddArc(&nodeB)
	nodeA.AddArc(&nodeF)
	nodeA.AddArc(&nodeG)

	nodeB.AddArc(&nodeA)
	nodeB.AddArc(&nodeF)
	nodeB.AddArc(&nodeE)
	nodeB.AddArc(&nodeD)
	nodeB.AddArc(&nodeC)

	nodeC.AddArc(&nodeB)
	nodeC.AddArc(&nodeD)

	nodeD.AddArc(&nodeB)
	nodeD.AddArc(&nodeC)

	nodeE.AddArc(&nodeF)
	nodeE.AddArc(&nodeB)

	nodeF.AddArc(&nodeA)
	nodeF.AddArc(&nodeB)
	nodeF.AddArc(&nodeE)

	nodeG.AddArc(&nodeA)
	nodeG.AddArc(&nodeH)
	nodeG.AddArc(&nodeI)
	nodeG.AddArc(&nodeJ)

	nodeH.AddArc(&nodeG)
	nodeH.AddArc(&nodeI)

	nodeI.AddArc(&nodeG)
	nodeI.AddArc(&nodeH)

	nodeJ.AddArc(&nodeG)

	return nodeA
}

func GetStatusNode(v string) statusNode[string] {
	/* var globalStatusNodePtr *statusNode[string]
	nodeA := GetStatusGraph1()
	Dfs[string](nodeA,func(n node[string]) {
		if n.Value == v {
			globalStatusNodePtr = &n
		}
	}) */
	return statusNode[string]{}
}
