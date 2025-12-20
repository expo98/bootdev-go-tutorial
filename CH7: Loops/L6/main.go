package main

func countConnections(groupSize int) int {

	var count int
	for i := 1; i < groupSize; i++ {

		for j := i; j < groupSize; j++ {
			count++
		}
	}

	return count
}
