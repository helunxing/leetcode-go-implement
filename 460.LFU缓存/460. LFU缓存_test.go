package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
// ["LFUCache","put","put","get","put","get","get","put","get","get","get"]
// [[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]

func Test_LFUCache(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	ast.Equal(1, cache.Get(1))

	cache.Put(3, 3)

	ast.Equal(-1, cache.Get(2))

	ast.Equal(3, cache.Get(3))

	cache.Put(4, 4)

	ast.Equal(-1, cache.Get(1))

	ast.Equal(3, cache.Get(3))

	ast.Equal(4, cache.Get(4))

}

func Test_LFUCache2(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(0)

	cache.Put(0, 0)

	ast.Equal(-1, cache.Get(0), "没能正确处理好， cap <= 0 的情况")
}

func Test_LFUCache3(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)

	cache.Put(3, 1)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Put(4, 4)

	ast.Equal(2, cache.Get(2), "更新 2 的值后，其 frequence 没有增加")
}
