# リングバッファ（Circular Buffer）とは

リングバッファは、固定サイズの配列を使用してデータを循環的に管理するデータ構造です。配列の末尾が先頭に接続されるように設計されています。

**このデータ構造を使う一番の利点は、メモリを効率的に使用しながら、一定のサイズ内でデータを循環させることができる点です。**

## 基本的な特徴

- **固定サイズ**: メモリ使用量が一定。
- **効率的なメモリ利用**: 配列を再利用するため、メモリ割り当てや解放のオーバーヘッドが少ない。
- **FIFO 原則**: キューと同様に、先入れ先出しのデータ構造。
- **用途**: 音声処理、ネットワークバッファ、リアルタイムシステムなど。

# 行える処理

| 機能    | 説明                             | 計算量 | 戻り値                | 得意/苦手 | 補足（その他）                                   |
| ------- | -------------------------------- | ------ | --------------------- | --------- | ------------------------------------------------ |
| Enqueue | バッファに新しい要素を追加       | O(1)   | なし（void）          | ✅        | サイズが固定されているため、効率的に追加可能。   |
| Dequeue | バッファから要素を取り出して削除 | O(1)   | 取り出した要素        | ✅        | 他の要素への影響なし。空の場合はエラー。         |
| Peek    | バッファの先頭要素を参照         | O(1)   | 先頭の要素            | ✅        | インデックス計算や検索が不要。空の場合はエラー。 |
| IsEmpty | バッファが空かどうかを確認       | O(1)   | bool（空なら true）   | ✅        |                                                  |
| IsFull  | バッファが満杯かどうかを確認     | O(1)   | bool（満杯なら true） | ✅        |                                                  |

# 注意点

1. **サイズ制限**: リングバッファは固定サイズのため、サイズを超えるとデータが上書きされる可能性があります。
2. **空の状態**: バッファが空の状態で `Dequeue` や `Peek` を呼び出すとエラーになるため、事前に空かどうかを確認する必要があります。
3. **動的サイズ変更が不可**: サイズを変更する場合は新しいバッファを作成する必要があります。

# 実装方法

## 実装方法のメリット・デメリット

| 実装方法   | メリット                                                             | デメリット                                                                                             |
| ---------- | -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| 配列       | - メモリが連続しておりキャッシュ効率が高い<br>- メモリ管理が簡単     | - サイズ変更不可、事前にサイズを決定する必要がある<br>- サイズ超過時にデータが上書きされる可能性がある |
| スライス   | - 動的にサイズ変更可能<br>- 標準ライブラリを活用でき実装が簡単       | - サイズ変更時にメモリ再割り当てが発生し、オーバーヘッドが増加<br>頻繁な再割り当てでパフォーマンス低下 |
| 連結リスト | - サイズが動的に変化しメモリ効率が良い<br>- 循環構造を簡単に実現可能 | - メモリが非連続でキャッシュ効率が低い<br>- 各ノードにポインタを持つためメモリ使用量が増加             |

## 配列を使用したリングバッファの実装

- **固定サイズ**: 配列を使用するため、リングバッファのサイズは固定。
- **メモリ効率**: 必要な分だけメモリを使用するため、メモリの無駄が少ない。
- **循環構造**: 配列の末尾が先頭に接続されるように設計し、head と tail をインデックスで管理。
- **計算量**: Enqueue と Dequeue の操作は O(1) で実行可能。

```go
package main

import "fmt"

// リングバッファ構造体
type ArrayCircularBuffer struct {
    data     []int
    head     int
    tail     int
    capacity int
    size     int
}

// 新しいリングバッファを作成
func NewArrayCircularBuffer(capacity int) *ArrayCircularBuffer {
    return &ArrayCircularBuffer{
        data:     make([]int, capacity),
        capacity: capacity,
    }
}

// Enqueue: 要素を追加
func (cb *ArrayCircularBuffer) Enqueue(value int) bool {
    if cb.size == cb.capacity {
        return false // バッファが満杯
    }
    cb.data[cb.tail] = value
    cb.tail = (cb.tail + 1) % cb.capacity
    cb.size++
    return true
}

// Dequeue: 要素を取り出す
func (cb *ArrayCircularBuffer) Dequeue() (int, bool) {
    if cb.size == 0 {
        return 0, false // バッファが空
    }
    value := cb.data[cb.head]
    cb.head = (cb.head + 1) % cb.capacity
    cb.size--
    return value, true
}

// Peek: 先頭要素を参照
func (cb *ArrayCircularBuffer) Peek() (int, bool) {
    if cb.size == 0 {
        return 0, false
    }
    return cb.data[cb.head], true
}

func main() {
    buffer := NewArrayCircularBuffer(3)

    // Enqueue操作
    buffer.Enqueue(10)
    buffer.Enqueue(20)
    buffer.Enqueue(30)

    // Dequeue操作
    if value, ok := buffer.Dequeue(); ok {
        fmt.Println("取り出した要素:", value) // 取り出した要素: 10
    }

    // Enqueue操作
    buffer.Enqueue(40)

    fmt.Println("バッファの状態:", buffer.data) // バッファの状態: [40 20 30]
}
```

## スライスを使用したリングバッファの実装

- **スライスを使用**: リングバッファのデータを保持するためにスライスを使用。
- **動的サイズ**: スライスの動的サイズ変更機能を活用し、柔軟にサイズを変更可能。
- **循環構造**: スライスのインデックスをモジュロ演算（%）で管理し、循環的な動作を実現。
- **計算量**: サイズ変更時にメモリ再割り当てが発生する場合、計算量が O(n) になる可能性がある。

```go
package main

import "fmt"

// リングバッファ構造体
type SliceCircularBuffer struct {
    data []int
    head int
    tail int
    size int
}

// 新しいリングバッファを作成
func NewSliceCircularBuffer(capacity int) *SliceCircularBuffer {
    return &SliceCircularBuffer{
        data: make([]int, capacity),
    }
}

// Enqueue: 要素を追加
func (cb *SliceCircularBuffer) Enqueue(value int) bool {
    if cb.size == len(cb.data) {
        return false // バッファが満杯
    }
    cb.data[cb.tail] = value
    cb.tail = (cb.tail + 1) % len(cb.data)
    cb.size++
    return true
}

// Dequeue: 要素を取り出す
func (cb *SliceCircularBuffer) Dequeue() (int, bool) {
    if cb.size == 0 {
        return 0, false // バッファが空
    }
    value := cb.data[cb.head]
    cb.head = (cb.head + 1) % len(cb.data)
    cb.size--
    return value, true
}

func main() {
    buffer := NewSliceCircularBuffer(3)

    // Enqueue操作
    buffer.Enqueue(10)
    buffer.Enqueue(20)
    buffer.Enqueue(30)

    // Dequeue操作
    if value, ok := buffer.Dequeue(); ok {
        fmt.Println("取り出した要素:", value) // 取り出した要素: 10
    }

    // Enqueue操作
    buffer.Enqueue(40)

    fmt.Println("バッファの状態:", buffer.data) // バッファの状態: [40 20 30]
}
```

## 連結リストを使用したリングバッファの実装

- **動的サイズ**: 連結リストを使用することで、リングバッファのサイズを動的に変更可能。
- **循環構造**: リストの末尾ノードを先頭ノードに接続することで、循環構造を実現。
- **計算量**: Enqueue と Dequeue の操作は O(1) で実行可能。
- **メモリ効率**: 必要な分だけメモリを使用するため、メモリの無駄が少ないが、各ノードにポインタを持つためメモリ使用量が増加。

```go
package main

import "fmt"

// ノード構造体
type CircularNode struct {
    value int
    next  *CircularNode
}

// リングバッファ構造体
type LinkedListCircularBuffer struct {
    head *CircularNode
    tail *CircularNode
    size int
    capacity int
}

// 新しいリングバッファを作成
func NewLinkedListCircularBuffer(capacity int) *LinkedListCircularBuffer {
    return &LinkedListCircularBuffer{
        capacity: capacity,
    }
}

// Enqueue: 要素を追加
func (cb *LinkedListCircularBuffer) Enqueue(value int) bool {
    if cb.size == cb.capacity {
        return false // バッファが満杯
    }
    newNode := &CircularNode{value: value}
    if cb.tail != nil {
        cb.tail.next = newNode
    }
    cb.tail = newNode
    if cb.head == nil {
        cb.head = newNode
    }
    cb.tail.next = cb.head // 循環構造を維持
    cb.size++
    return true
}

// Dequeue: 要素を取り出す
func (cb *LinkedListCircularBuffer) Dequeue() (int, bool) {
    if cb.size == 0 {
        return 0, false // バッファが空
    }
    value := cb.head.value
    cb.head = cb.head.next
    cb.tail.next = cb.head // 循環構造を維持
    cb.size--
    return value, true
}

func main() {
    buffer := NewLinkedListCircularBuffer(3)

    // Enqueue操作
    buffer.Enqueue(10)
    buffer.Enqueue(20)
    buffer.Enqueue(30)

    // Dequeue操作
    if value, ok := buffer.Dequeue(); ok {
        fmt.Println("取り出した要素:", value) // 取り出した要素: 10
    }

    // Enqueue操作
    buffer.Enqueue(40)

    fmt.Println("バッファの状態: 循環構造のため直接表示不可")
}
```

# リングバッファの応用例

1. **音声処理**: 音声データのバッファリング。
2. **ネットワークバッファ**: パケットの順序を保ちながら処理。
3. **リアルタイムシステム**: 一定サイズのデータを効率的に管理。

# 他のデータ構造との関係

- **キュー**: リングバッファは固定サイズのキューとして利用可能。
- **デック**: 両端操作をサポートするリングバッファも実装可能。

# 他のデータ構造との比較

| 特徴                    | リングバッファ | キュー         | デック         |
| ----------------------- | -------------- | -------------- | -------------- |
| **固定サイズ**          | あり           | なし（動的）   | なし（動的）   |
| **両端操作**            | 制限あり       | 不可           | 可能           |
| **順序性**              | FIFO           | FIFO           | 挿入順序を保持 |
| **ランダムアクセス**    | 不可           | 不可           | 不可           |
| **サイズ変更**          | 不可           | 動的または固定 | 動的または固定 |
| **操作の計算量 (追加)** | O(1)           | O(1)           | O(1)           |
| **操作の計算量 (削除)** | O(1)           | O(1)           | O(1)           |
