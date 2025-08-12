# ヒープ（Heap）とは

ヒープは「完全二分木」の一種で、各ノードの値がその子ノードの値以上（最大ヒープ）または以下（最小ヒープ）となるように構成されるデータ構造です。主に**優先度付きキュー**や**ヒープソート**などで利用されます。

**このデータ構造の最大の利点は、最小値（または最大値）の高速な取得・削除・挿入が可能な点です。**

# ヒープの構成要素

- **ノード（Node）**: データを保持する単位。親・子の関係を持つ。
- **ルート（Root）**: 最上位のノード。最小ヒープなら最小値、最大ヒープなら最大値を持つ。
- **エッジ（Edge）**: 親子ノードを結ぶ線。
- **葉（Leaf）**: 子ノードを持たないノード。
- **完全二分木**: 全てのレベルが左詰めで埋まっている二分木。

# ヒープの種類

- **最小ヒープ（Min Heap）**: 親 ≤ 子。ルートが最小値。
- **最大ヒープ（Max Heap）**: 親 ≥ 子。ルートが最大値。
- **二項ヒープ・フィボナッチヒープ**: より高度なヒープ構造（主に理論・特殊用途）。

# ヒープの特徴

- **最小/最大値の取得**: O(1)
- **挿入・削除**: O(log n)
- **完全二分木**なので配列で効率的に実装可能
- **順序走査や範囲検索は苦手**

# ヒープの一般的な実装方法

## 配列による実装

- ノードを配列で管理し、インデックスで親子関係を計算
  - 親: `(i-1)/2`
  - 左子: `2i+1`
  - 右子: `2i+2`
- メモリ効率・キャッシュ効率が高い
- 動的サイズ変更も容易

## Go による最小ヒープの実装例

```go
package heap

type MinHeap struct {
    data []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{data: []int{}}
}

func (h *MinHeap) Insert(val int) {
    h.data = append(h.data, val)
    h.upHeap(len(h.data) - 1)
}

func (h *MinHeap) upHeap(i int) {
    for i > 0 {
        parent := (i - 1) / 2
        if h.data[parent] <= h.data[i] {
            break
        }
        h.data[parent], h.data[i] = h.data[i], h.data[parent]
        i = parent
    }
}

func (h *MinHeap) ExtractMin() (int, bool) {
    if len(h.data) == 0 {
        return 0, false
    }
    min := h.data[0]
    last := h.data[len(h.data)-1]
    h.data = h.data[:len(h.data)-1]
    if len(h.data) > 0 {
        h.data[0] = last
        h.downHeap(0)
    }
    return min, true
}

func (h *MinHeap) downHeap(i int) {
    n := len(h.data)
    for {
        left := 2*i + 1
        right := 2*i + 2
        smallest := i
        if left < n && h.data[left] < h.data[smallest] {
            smallest = left
        }
        if right < n && h.data[right] < h.data[smallest] {
            smallest = right
        }
        if smallest == i {
            break
        }
        h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
        i = smallest
    }
}

func (h *MinHeap) Peek() (int, bool) {
    if len(h.data) == 0 {
        return 0, false
    }
    return h.data[0], true
}
```

# 優先度付きキュー（Priority Queue）とは

優先度付きキューは、各要素に「優先度」を持たせ、常に最も優先度の高い（または低い）要素を高速に取り出せるデータ構造です。  
ヒープは優先度付きキューの実装に最適です。

## 優先度付きキューの一般的な実装

- **ヒープ（配列実装）**: 挿入・削除ともに O(log n)、最小/最大値の取得が O(1)
- **連結リストや配列**: 挿入または削除が O(n) となり非効率
- **Go の標準ライブラリ**: `container/heap` パッケージで汎用ヒープをサポート

## Go による優先度付きキューの実装例

Go では`container/heap`を使うのが一般的です。  
以下は最小ヒープによる優先度付きキューの例です。

```go
package main

import (
    "container/heap"
    "fmt"
)

type Item struct {
    value    string
    priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].priority < pq[j].priority // 最小ヒープ
}
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func main() {
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &Item{value: "task1", priority: 3})
    heap.Push(pq, &Item{value: "task2", priority: 1})
    heap.Push(pq, &Item{value: "task3", priority: 2})

    for pq.Len() > 0 {
        item := heap.Pop(pq).(*Item)
        fmt.Println(item.value, item.priority)
    }
}
```

# まとめ

- ヒープは完全二分木を配列で実装し、最小/最大値の高速な取得・削除・挿入が可能
- 優先度付きキューはヒープで効率的に実装できる
- Go では`container/heap`パッケージを使うことで、柔軟な優先度付きキューを簡単に実装できる
