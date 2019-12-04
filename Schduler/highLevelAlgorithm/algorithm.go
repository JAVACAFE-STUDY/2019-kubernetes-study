schedule(pod): string
	nodes := getAllHealthyNodes()
	viableNodes := []
	for node in nodes:
 		for predicate in predicates:
 			if predicate(node, pod):
 				viableNodes.append(node)
	scoredNodes := PriorityQueue<score, Node[]>
 	priorities := GetPriorityFunctions()
 	for node in viableNodes:
 		score = CalculateCombinedPriority(node, pod, priorities)
 		scoredNodes[score].push(node)
 	bestScore := scoredNodes.top().score
 	selectedNodes := []
 	while scoredNodes.top().score == bestScore:
 		selectedNodes.append(scoredNodes.pop())
 	node := selectAtRandom(selectedNodes)
 	return node.Name