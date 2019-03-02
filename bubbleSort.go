package main

func bubbleSort(Slice []int) []int{
	n := len(Slice)-1

	for i := 0; i <= n; i++{
		for j := 0; j < n-i; j++{
			if Slice[j] > Slice[j+1]{
				Slice[j], Slice[j+1] = Slice[j+1], Slice[j]
			}
		}
	}
	return Slice
}