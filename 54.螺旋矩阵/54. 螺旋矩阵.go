package leetcode

func spiralOrder(matrix [][]int) []int {
	m:=len(matrix)
	if m==0{
		return []int{}
	}
	n:=len(matrix[0])
	if n==0{
		return []int{}
	}
	res:=[]int{}
	x1:=0
	x2:=m-1
	y1:=0
	y2:=n-1
	for x1<=x2 && y1<=y2{
		for j:=y1;j<=y2;j++{
			res=append(res,matrix[x1][j])
		}
		for i:=x1+1;i<=x2;i++{
			res=append(res,matrix[i][y2])
		}
		if x1==x2 || y1==y2{
			break
		}
		for j:=y2-1;j>y1;j--{
			res=append(res,matrix[x2][j])
		}
		for i:=x2;i>x1;i--{
			res=append(res,matrix[i][y1])
		}
		x1++
		x2--
		y1++
		y2--
	}
	return res
}