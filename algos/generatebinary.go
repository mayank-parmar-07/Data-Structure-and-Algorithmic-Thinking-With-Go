package algos

func Generate_str(results *[]string, n, counter int, curr string) {
	if counter == n {
		*results = append(*results, curr)
	} else {
		Generate_str(results, n, counter+1, curr+"0")
		Generate_str(results, n, counter+1, curr+"1")
	}
}
