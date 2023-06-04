package example

// Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees. Can you do this in place?

func RotateMatrix(matrix [][]int) [][]int {

	lines := len(matrix)
	columns := len(matrix[0])

	newMatrix := make([][]int, columns)
	for i := range newMatrix {
		newMatrix[i] = make([]int, lines)
	}

	for i := range matrix {
		for j := range matrix[i] {
			newMatrix[j][lines-1-i] = matrix[i][j]
		}

	}

	return newMatrix

}
