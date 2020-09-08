package prepare_exam

// -> [1,7,5,6,4] 7
// search element
// O(N)

func Search (a []int, b int) int {
	for i, x := range a {
		if b == x {
			return i
		}
	}

	return -1
}

// search in sorted list
// -> [1,3,5,6,7] 7
// -> O(log N)
// #1 Steps:
// [1,3,5,6,7] start 0, end 4, mid 5 5 < 7
// [1,3,5,6,7] start 3, end 4, mid 6 6 < 7
// [1,3,4,6,7] start 4, end 4, mid 7 == 7 -> 4

// #2 Steps:
// [1,3,5,6,7] mid 5 5 < 7
// [6,7] mid 6 6 < 7
// [7] mid 7 7 == 7 -> true
func BinarySearch (a []int, x int, start, end int) int {
	if end - start <= 0 {
		return -1
	}
	mid := (end + start) / 2
	if a[mid] == x {
		return mid
	}
	if a[mid] < x {
		return BinarySearch(a, x, mid + 1, end)
	}

	return BinarySearch(a, x, start, mid - 1)
}

// dfs, bfs
/*
a - c - d
 \ /
  b
 */
/*
dfs(G, f) -> false
dfs(G, c) -> true

 */

func DFS(G *map[string][]string, node string, x string, visited *map[string]bool) bool {
	(*visited)[node]= true

	if node == x {
		return true
	}

	for _, el := range (*G)[node] {
		if _, ok := (*visited)[el]; !ok {
			neighborFound := DFS(G, el, x, visited)
			if neighborFound {
				return true
			}
		}
	}

	return false
}
