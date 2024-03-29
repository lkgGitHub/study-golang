package hash

/*
771. 宝石与石头 [https://leetcode-cn.com/problems/jewels-and-stones/]

给你一个字符串 jewels 代表石头中宝石的类型，另有一个字符串 stones 代表你拥有的石头。
stones 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。

示例 1：
输入：jewels = "aA", stones = "aAAbbbb"
输出：3

示例 2：
输入：jewels = "z", stones = "ZZ"
输出：0


提示：
1 <= jewels.length, stones.length <= 50
jewels 和 stones 仅由英文字母组成
jewels 中的所有字符都是 唯一的

*/
import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"示例 1", args{"aA", "aAAbbbb"}, 3},
		{"示例 2", args{"z", "ZZ"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func numJewelsInStones(jewels string, stones string) int {
	var i int
	m := make(map[int32]struct{}, len(jewels))
	for _, j := range jewels {
		m[j] = struct{}{}
	}
	for _, s := range stones {
		if _, ok := m[s]; ok {
			i++
		}
	}
	return i
}
