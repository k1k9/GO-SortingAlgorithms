package main

func selectSort(Slice []int) []int{
	// aIndex is index of assumption the smallest number in Slice
	// n is a Slice length
	aIndex, n := 0, len(Slice)

	for index,_:= range Slice{

		// Assumption the smallest number
		aNum := Slice[aIndex]

		// Find smaller number than aNum and if find, save as aNum
		for j := aIndex; j < n; j++{
			if aNum > Slice[j]{
				aNum, aIndex = Slice[j], j
			}
		}

		// Swap the smallest number position with number on index (index)
		Slice[aIndex] = Slice[index]
		Slice[index] = aNum

		// Decrese search range of the smallest number
		aIndex = index+1 
	}
	return Slice
}