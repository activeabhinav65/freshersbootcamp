package main

import "fmt"

type matrix struct {
  mat [2][4] int

  row int
  col int
}
func (ans matrix) rows() int{
  return len(ans.mat)
}
func (ans matrix) cols() int {
  return len(ans.mat[0])
}

func (ans matrix) change() {
    ans.mat[0][1]=4
  }
func (ans matrix) sum(mat2 [2][4] int) {
  for i := 0; i < 2; i++ {
    for j := 0; j < 4; j++ {
      ans.mat[i][j] = ans.mat[i][j] + mat2[i][j]
    }
  }
}

func main() {
    var mat = matrix{[2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}, 2, 4}

    fmt.Println(mat.rows())
    fmt.Println(mat.cols())
    mat.change()
    var mat2=[2][4] int{{3,4,5,7},{5,6,7,8}}
    mat.sum(mat2)

}

