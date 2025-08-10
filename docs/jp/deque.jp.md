# デック（Deque）とは

デック（Deque: Double-Ended Queue）は、両端から要素の追加と削除が可能なデータ構造です。スタックやキューの拡張版とも言えます。

**このデータ構造を使う一番の利点は、柔軟性が高く、スタック（LIFO）やキュー（FIFO）としても利用できる点です。**

## 基本的な特徴

- **両端操作**: 要素の追加と削除が両端で可能。
- **柔軟性**: スタックやキューとしても利用可能。
- **順序性**: 要素の挿入順序を保持。
- **動的サイズ**: 実行時にサイズを変更可能（実装方法による）。

# 行える処理

| 機能      | 説明                                 | 計算量 | 戻り値              | 得意/苦手 | 補足（その他）                 |
| --------- | ------------------------------------ | ------ | ------------------- | --------- | ------------------------------ |
| PushFront | デックの前に新しい要素を追加         | O(1)   | なし（void）        | ✅        | 両端での操作が可能。           |
| PushBack  | デックの後ろに新しい要素を追加       | O(1)   | なし（void）        | ✅        | 両端での操作が可能。           |
| PopFront  | デックの前から要素を取り出して削除   | O(1)   | 取り出した要素      | ✅        | 両端での操作が可能。           |
| PopBack   | デックの後ろから要素を取り出して削除 | O(1)   | 取り出した要素      | ✅        | 両端での操作が可能。           |
| PeekFront | デックの前の要素を参照               | O(1)   | 前の要素            | ✅        | インデックス計算や検索が不要。 |
| PeekBack  | デックの後ろの要素を参照             | O(1)   | 後ろの要素          | ✅        | インデックス計算や検索が不要。 |
| IsEmpty   | デックが空かどうかを確認             | O(1)   | bool（空なら true） | ✅        |                                |
| Size      | デック内の要素数を取得               | O(1)   | int（要素数）       | ✅        |                                |

# 注意点

1. **中間要素へのアクセスが非効率**: デックは両端以外の要素に直接アクセスできないため、ランダムアクセスが必要な場合には不向きです。
2. **サイズ制限**: 固定サイズのデックを使用する場合、サイズを超えるとエラーやデータの上書きが発生する可能性があります。
3. **空の状態**: デックが空の状態で `PopFront` や `PopBack` を呼び出すとエラーになるため、事前に空かどうかを確認する必要があります。

# 実装方法

## 実装方法のメリット・デメリット

| 実装方法   | メリット                                     | デメリット                                                   |
| ---------- | -------------------------------------------- | ------------------------------------------------------------ |
| 配列       | - メモリが連続しておりキャッシュ効率が高い   | - サイズ変更不可、事前にサイズを決定する必要がある           |
|            | - 両端の操作が効率的（固定サイズ内での操作） | - サイズ超過時にオーバーフローが発生                         |
| スライス   | - 動的にサイズ変更可能                       | - 両端操作でスライスの再割り当てが発生し、パフォーマンス低下 |
|            | - 標準ライブラリを活用でき実装が簡単         | - 頻繁な再割り当てで効率が悪化                               |
| 連結リスト | - 両端の操作が効率的（O(1)）                 | - メモリが非連続でキャッシュ効率が低い                       |
|            | - サイズが動的に変化しメモリ効率が良い       | - 各ノードにポインタを持つためメモリ使用量が増加             |

## 配列を使用したデックの実装

```go
package main

import "fmt"

// デック構造体
type ArrayDeque struct {
    data []int
    head int
    tail int
    size int
    capacity int
}

// 新しいデックを作成
func NewArrayDeque(capacity int) *ArrayDeque {
    return &ArrayDeque{
        data: make([]int, capacity),
        capacity: capacity,
    }
}

// PushFront: デックの前に要素を追加
func (d *ArrayDeque) PushFront(value int) bool {
    if d.size == d.capacity {
        return false // デックが満杯
    }
    d.head = (d.head - 1 + d.capacity) % d.capacity
    d.data[d.head] = value
    d.size++
    return true
}

// PushBack: デックの後ろに要素を追加
func (d *ArrayDeque) PushBack(value int) bool {
    if d.size == d.capacity {
        return false // デックが満杯
    }
    d.data[d.tail] = value
    d.tail = (d.tail + 1) % d.capacity
    d.size++
    return true
}

// PopFront: デックの前から要素を削除
func (d *ArrayDeque) PopFront() (int, bool) {
    if d.size == 0 {
        return 0, false // デックが空
    }
    value := d.data[d.head]
    d.head = (d.head + 1) % d.capacity
    d.size--
    return value, true
}

// PopBack: デックの後ろから要素を削除
func (d *ArrayDeque) PopBack() (int, bool) {
    if d.size == 0 {
        return 0, false // デックが空
    }
    d.tail = (d.tail - 1 + d.capacity) % d.capacity
    value := d.data[d.tail]
    d.size--
    return value, true
}

func main() {
    deque := NewArrayDeque(5)

    // 要素を追加
    deque.PushFront(10)
    deque.PushBack(20)
    deque.PushFront(5)

    fmt.Println("デックの状態:", deque.data) // デックの状態: [5 10 20]

    // 要素を削除
    if value, ok := deque.PopFront(); ok {
        fmt.Println("前から削除:", value) // 前から削除: 5
    }
    if value, ok := deque.PopBack(); ok {
        fmt.Println("後ろから削除:", value) // 後ろから削除: 20
    }
}
```

## スライスを使用したデックの実装

```go
package main

import "fmt"

type SliceDeque struct {
    data []int
}

// PushFront: デックの前に要素を追加
func (d *SliceDeque) PushFront(value int) {
    d.data = append([]int{value}, d.data...)
}

// PushBack: デックの後ろに要素を追加
func (d *SliceDeque) PushBack(value int) {
    d.data = append(d.data, value)
}

// PopFront: デックの前から要素を削除
func (d *SliceDeque) PopFront() (int, bool) {
    if len(d.data) == 0 {
        return 0, false
    }
    value := d.data[0]
    d.data = d.data[1:]
    return value, true
}

// PopBack: デックの後ろから要素を削除
func (d *SliceDeque) PopBack() (int, bool) {
    if len(d.data) == 0 {
        return 0, false
    }
    value := d.data[len(d.data)-1]
    d.data = d.data[:len(d.data)-1]
    return value, true
}

func main() {
    deque := &SliceDeque{}

    // 要素を追加
    deque.PushFront(10)
    deque.PushBack(20)
    deque.PushFront(5)

    fmt.Println("デックの状態:", deque.data) // デックの状態: [5 10 20]

    // 要素を削除
    if value, ok := deque.PopFront(); ok {
        fmt.Println("前から削除:", value) // 前から削除: 5
    }
    if value, ok := deque.PopBack(); ok {
        fmt.Println("後ろから削除:", value) // 後ろから削除: 20
    }
}
```

## 連結リストを使用したデックの実装

```go
package main

import "fmt"

// ノード構造体
type DequeNode struct {
    value int
    prev  *DequeNode
    next  *DequeNode
}

// デック構造体
type LinkedListDeque struct {
    head *DequeNode
    tail *DequeNode
}

// PushFront: デックの前に要素を追加
func (d *LinkedListDeque) PushFront(value int) {
    newNode := &DequeNode{value: value}
    if d.head != nil {
        newNode.next = d.head
        d.head.prev = newNode
    }
    d.head = newNode
    if d.tail == nil {
        d.tail = newNode
    }
}

// PushBack: デックの後ろに要素を追加
func (d *LinkedListDeque) PushBack(value int) {
    newNode := &DequeNode{value: value}
    if d.tail != nil {
        newNode.prev = d.tail
        d.tail.next = newNode
    }
    d.tail = newNode
    if d.head == nil {
        d.head = newNode
    }
}

// PopFront: デックの前から要素を削除
func (d *LinkedListDeque) PopFront() (int, bool) {
    if d.head == nil {
        return 0, false
    }
    value := d.head.value
    d.head = d.head.next
    if d.head != nil {
        d.head.prev = nil
    } else {
        d.tail = nil
    }
    return value, true
}

// PopBack: デックの後ろから要素を削除
func (d *LinkedListDeque) PopBack() (int, bool) {
    if d.tail == nil {
        return 0, false
    }
    value := d.tail.value
    d.tail = d.tail.prev
    if d.tail != nil {
        d.tail.next = nil
    } else {
        d.head = nil
    }
    return value, true
}

func main() {
    deque := &LinkedListDeque{}

    // 要素を追加
    deque.PushFront(10)
    deque.PushBack(20)
    deque.PushFront(5)

    // 要素を削除
    if value, ok := deque.PopFront(); ok {
        fmt.Println("前から削除:", value) // 前から削除: 5
    }
    if value, ok := deque.PopBack(); ok {
        fmt.Println("後ろから削除:", value) // 後ろから削除: 20
    }
}
```

# デックの応用例

1. **タスクスケジューリング**: 優先度付きタスクの管理。
2. **キャッシュ管理**: LRU（Least Recently Used）キャッシュの実装。
3. **文字列処理**: 回文チェックや文字列のリバース操作。

# 他のデータ構造との関係

- **スタック**: デックの片側のみを使用すればスタックとして利用可能。
- **キュー**: デックの片側のみを使用すればキューとして利用可能。

# 他のデータ構造との比較

| 特徴                    | デック         | スタック       | キュー         |
| ----------------------- | -------------- | -------------- | -------------- |
| **両端操作**            | 可能           | 不可           | 不可           |
| **片端操作**            | 可能           | 可能           | 可能           |
| **順序性**              | 挿入順序を保持 | 挿入順序を保持 | 挿入順序を保持 |
| **ランダムアクセス**    | 不可           | 不可           | 不可           |
| **サイズ変更**          | 動的または固定 | 動的または固定 | 動的または固定 |
| **操作の計算量 (追加)** | O(1)           | O(1)           | O(1)           |
| **操作の計算量 (削除)** | O(1)           | O(1)           | O(1)           |
