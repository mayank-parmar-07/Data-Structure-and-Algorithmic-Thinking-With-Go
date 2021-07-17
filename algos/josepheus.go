package algos

import "fmt"

func Josepheus(map_set map[int]bool, n, k, curr, left int) {
	if left == 1 {
		for key, val := range map_set {
			if val == false {
				fmt.Println(key+1, " number person is saved")
				return
			}
		}
	} else {
		count := 0
		node_to_check := curr
		for count < k-1 {
			node_to_check = (node_to_check + 1) % n
			if x := map_set[node_to_check]; !x {
				count += 1
			}
		}
		fmt.Println(curr, " kills ", node_to_check)
		map_set[node_to_check] = true
		for true {
			if x := map_set[node_to_check]; !x {
				break
			}
			node_to_check = (node_to_check + 1) % n
		}
		Josepheus(map_set, n, k, node_to_check, left-1)
	}
}
