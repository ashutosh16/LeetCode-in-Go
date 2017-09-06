package Problem0076

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	s string
	t string
}

// ans 是答案
type ans struct {
	one string
}

func Test_Problem0076(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"ADOBECODEBANC",
				"ABC",
			},
			ans{
				"BANC",
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, minWindow(p.s, p.t), "输入:%v", p)
	}
}
