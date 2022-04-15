package hash

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
374. 子域名访问计数
难度：中等
网站域名 "discuss.leetcode.com" 由多个子域名组成。顶级域名为 "com" ，二级域名为 "leetcode.com" ，最低一级为 "discuss.leetcode.com" 。当访问域名 "discuss.leetcode.com" 时，同时也会隐式访问其父域名 "leetcode.com" 以及 "com" 。

计数配对域名 是遵循 "rep d1.d2.d3" 或 "rep d1.d2" 格式的一个域名表示，其中 rep 表示访问域名的次数，d1.d2.d3 为域名本身。

例如，"9001 discuss.leetcode.com" 就是一个 计数配对域名 ，表示 discuss.leetcode.com 被访问了 9001 次。
给你一个 计数配对域名 组成的数组 cpdomains ，解析得到输入中每个子域名对应的 计数配对域名 ，并以数组形式返回。可以按 任意顺序 返回答案。

示例 1：
输入：cpdomains = ["9001 discuss.leetcode.com"]
输出：["9001 leetcode.com","9001 discuss.leetcode.com","9001 com"]
解释：例子中仅包含一个网站域名："discuss.leetcode.com"。
按照前文描述，子域名 "leetcode.com" 和 "com" 都会被访问，所以它们都被访问了 9001 次。

示例 2：
输入：cpdomains = ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
输出：["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
解释：按照前文描述，会访问 "google.mail.com" 900 次，"yahoo.com" 50 次，"intel.mail.com" 1 次，"wiki.org" 5 次。
而对于父域名，会访问 "mail.com" 900 + 1 = 901 次，"com" 900 + 50 + 1 = 951 次，和 "org" 5 次。


提示：
1 <= cpdomain.length <= 100
1 <= cpdomain[i].length <= 100
cpdomain[i] 会遵循 "repi d1i.d2i.d3i" 或 "repi d1i.d2i" 格式
repi 是范围 [1, 104] 内的一个整数
d1i、d2i 和 d3i 由小写英文字母组成
*/
func TestSubdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"示例 1：", args{[]string{"9001 discuss.leetcode.com"}}, []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"}},
		{"示例 2：", args{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}}, []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantMap := make(map[string]struct{}, len(tt.want))
			for _, w := range tt.want {
				wantMap[w] = struct{}{}
			}
			got := subdomainVisits(tt.args.cpdomains)
			for _, g := range got {
				if _, ok := wantMap[g]; !ok {
					t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
				} else {
					delete(wantMap, g)
				}
			}
			if len(wantMap) != 0 {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func subdomainVisits(cpdomains []string) []string {
	domainMap := make(map[string]int, len(cpdomains)*2)
	for _, domain := range cpdomains {
		domainSlice := strings.Split(domain, " ")
		if len(domainSlice) != 2 {
			continue
		}
		count, _ := strconv.Atoi(domainSlice[0])
		ds := strings.Split(domainSlice[1], ".")
		for i := 0; i < len(ds); i++ {
			var d string
			switch i {
			case 0:
				d = domainSlice[1]
			case 1:
				d = strings.TrimPrefix(domainSlice[1], ds[0]+".")
			case 2:
				d = ds[2]
			}

			if v, ok := domainMap[d]; ok {
				domainMap[d] = v + count
			} else {
				domainMap[d] = count
			}
		}
	}
	domains := make([]string, 0, len(domainMap))
	for k, v := range domainMap {
		domains = append(domains, fmt.Sprintf("%d %s", v, k))
	}

	return domains
}
