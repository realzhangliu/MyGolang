package Misc

import (
	"fmt"
	"math"
)

var w [5]int = [5]int{5, 5, 3, 4, 3}           //重量数组
var v [5]int = [5]int{400, 500, 200, 300, 350} //价值数组

func CallDynamic() {
	var cap int = 10
	var n int = len(w)
	max := computer(n-1, cap)
	fmt.Println("【", cap, "容量的背包在", n, "个物品里选择能装下的最大价值是", max, "】")
}

func computer(nIndex int, cap int) int {
	//基准条件：如果索引无效或者容量不足，直接返回当前价值0
	if nIndex < 0 || cap <= 0 {
		return 0
	}
	//不放第n个物品所得价值
	res := computer(nIndex-1, cap)
	//放第n个物品所得值(前提是要放的下)
	var v2 int
	if w[nIndex] <= cap {
		//计算放的下的方案值    v[n]是当前物品的价值，computer(n-1, cap-w[n])是计算前一个物品，在减去这个物品容量后的容量下的最大价值方案
		v2 = v[nIndex] + computer(nIndex-1, cap-w[nIndex])

	}
	return int(math.Max(float64(res), float64(v2)))
}
