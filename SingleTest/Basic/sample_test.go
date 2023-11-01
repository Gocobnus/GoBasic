package Basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	var flagTest = []struct {
		name   string
		input  string
		sep    string
		output []string
	}{
		{"base case", "a,b,c", ",", []string{"a", "b", "c"}},
		{"wrong case", "a:b:c", ",", []string{"a:b:c"}},
		{"multi case", "abcd", "bc", []string{"a", "d"}},
		{"chinese case", "吃葡萄不吐葡萄皮", "葡", []string{"吃", "萄不吐", "萄皮"}},
	}
	for _, tt := range flagTest {
		//assert := assert.New(t) 可以直接实例化一次，后续assert不用传入testing.T
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, Split(tt.input, tt.sep))
		})

	}
}
