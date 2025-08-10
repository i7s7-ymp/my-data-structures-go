# キュー（Queue）とは

キューは、FIFO（First In, First Out：先入れ先出し）の原則に従ってデータを管理する線形データ構造です。最初に追加された要素が最初に取り出される特性を持ちます。

**このデータ構造を使う一番の利点は、順序を保ちながら効率的にデータを処理できる点です。** キューは、タスクのスケジューリングやデータのストリーム処理など、順序が重要な場面で広く利用されます。

## 基本的な特徴

- **FIFO 原則**: 最初に入れたものが最初に出る
- **単一アクセスポイント**: 先頭（取り出し）と末尾（追加）のみ操作可能
- **動的サイズ**: 実行時にサイズを変更可能（実装方法による。固定長配列やリングバッファの場合は固定サイズ）
- **順序性**: 要素の挿入順序をそのまま保持
- **制限されたアクセス**: 中間要素への直接アクセス不可

# 行える処理

| 機能         | 説明                                 | 計算量          | 戻り値              | 得意/苦手 | 補足（その他）                                   |
| ------------ | ------------------------------------ | --------------- | ------------------- | --------- | ------------------------------------------------ |
| Enqueue      | キューの末尾に新しい要素を追加       | O(1)            | なし（void）        | ✅        | 順序を保ちながら効率的に追加可能。               |
| Dequeue      | キューの先頭から要素を取り出して削除 | O(1)            | 取り出した要素      | ✅        | 他の要素への影響なし。空の場合はエラー。         |
| Peek/Front   | キューの先頭要素を削除せずに参照     | O(1)            | 先頭の要素          | ✅        | インデックス計算や検索が不要。空の場合はエラー。 |
| IsEmpty      | キューが空かどうかを確認             | O(1)            | bool（空なら true） | ✅        |                                                  |
| Size         | キュー内の要素数を取得               | O(1)            | int（要素数）       | ✅        |                                                  |
| Clear        | キューのすべての要素を削除           | O(n)または O(1) | なし（void）        | -         | 実装による。                                     |
| 中間アクセス | 中間要素にアクセス                   | O(n)            | 該当要素            | ❌        | 中間要素への直接アクセスは非効率。               |
| 検索         | 特定の値を持つ要素を検索             | O(n)            | 該当要素または bool | ❌        | 順次アクセスが必要。                             |

# 注意点

1. **中間要素へのアクセスが非効率**: キューは先頭と末尾以外の要素に直接アクセスできないため、ランダムアクセスが必要な場合には不向きです。
2. **サイズ制限**: 固定サイズのキューを使用する場合、サイズを超えるとエラーやデータの上書きが発生する可能性があります。
3. **空の状態**: キューが空の状態で `Dequeue` や `Peek` を呼び出すとエラーになるため、事前に空かどうかを確認する必要があります。

# デック（Deque）とリングバッファ（Circular Buffer）

## デック（Deque）とは

デック（Deque: Double-Ended Queue）は、両端から要素の追加と削除が可能なデータ構造です。スタックやキューの拡張版とも言えます。

### 特徴

- **両端操作**: 要素の追加と削除が両端で可能。
- **柔軟性**: スタック（LIFO）やキュー（FIFO）としても利用可能。
- **用途**: タスクスケジューリング、キャッシュ管理、文字列処理など。

[デック（Deque）について詳しくはこちら](deque.jp.md)

## リングバッファ（Circular Buffer）とは

リングバッファは、固定サイズの配列を使用してデータを循環的に管理するデータ構造です。配列の末尾が先頭に接続されるように設計されています。

### 特徴

- **固定サイズ**: メモリ使用量が一定。
- **効率的なメモリ利用**: 配列を再利用するため、メモリ割り当てや解放のオーバーヘッドが少ない。
- **用途**: 音声処理、ネットワークバッファ、リアルタイムシステムなど。

[リングバッファ（Circular Buffer）について詳しくはこちら](circular-buffer.jp.md)

# 実装方法

## 配列・スライス・連結リストによる実装の比較

以下は、通常のキューを配列、スライス、連結リストで実装した場合のメリットとデメリットを表にまとめたものです。

| 実装方法                   | メリット                                           | デメリット                                                                                 |
| -------------------------- | -------------------------------------------------- | ------------------------------------------------------------------------------------------ |
| スライスをキューとして実装 | 実装が簡単                                         | 先頭要素を削除すると全ての要素をシフトするため O(n) で非効率                               |
|                            | 動的サイズで柔軟に対応可能                         | サイズ変更時にメモリ再割り当てが発生し、オーバーヘッドが増加                               |
|                            | 他の操作（スタックなど）にも流用可能で汎用性が高い | スライスの切り取り操作では元のスライスのメモリが解放されず、メモリの断片化が発生する可能性 |
|                            |                                                    | 固定サイズのキューには不向き                                                               |
| 配列                       | メモリが連続しておりキャッシュ効率が高い           | サイズ変更不可、事前にサイズを決定する必要がある                                           |
|                            | メモリ管理が簡単                                   | サイズ超過時にオーバーフローが発生                                                         |
| スライス                   | 動的にサイズ変更可能                               | サイズ変更時にメモリ再割り当てが発生し、オーバーヘッドが増加                               |
|                            | 標準ライブラリを活用でき実装が簡単                 | 頻繁な再割り当てでパフォーマンス低下                                                       |
| 連結リスト                 | サイズが動的に変化しメモリ効率が良い               | メモリが非連続でキャッシュ効率が低い                                                       |
|                            | 要素の挿入・削除が一定時間で可能（O(1)）           | 各ノードにポインタを持つためメモリ使用量が増加                                             |

## スライスをキューのように扱う実装

Go にはキュー専用の組み込み型はありませんが、スライスを直接使用してキューのように扱うことができます。

```go
package main

import "fmt"

func main() {
    // キューとして使用するスライス
    queue := []int{}

    // Enqueue操作
    queue = append(queue, 10) // キューに10を追加
    queue = append(queue, 20) // キューに20を追加
    queue = append(queue, 30) // キューに30を追加

    fmt.Println("キューの状態:", queue) // キューの状態: [10 20 30]

    // Dequeue操作
    if len(queue) > 0 {
        front := queue[0] // 先頭要素を取得
        queue = queue[1:] // 先頭要素を削除
        fmt.Println("取り出した要素:", front) // 取り出した要素: 10
    }

    // 現在のキューの状態
    fmt.Println("キューの状態:", queue) // キューの状態: [20 30]
}
```

## 配列を使用したキューの実装

以下は、配列を使用してキューを実装する例です。

```go
package main

import "fmt"

// キュー構造体
type ArrayQueue struct {
    data     []int
    head     int
    tail     int
    capacity int
    size     int
}

// 新しいキューを作成
func NewArrayQueue(capacity int) *ArrayQueue {
    return &ArrayQueue{
        data:     make([]int, capacity),
        capacity: capacity,
    }
}

// Enqueue: キューに要素を追加
func (q *ArrayQueue) Enqueue(value int) bool {
    if q.size == q.capacity {
        return false // キューが満杯
    }
    q.data[q.tail] = value
    q.tail = (q.tail + 1) % q.capacity
    q.size++
    return true
}

// Dequeue: キューから要素を取り出す
func (q *ArrayQueue) Dequeue() (int, bool) {
    if q.size == 0 {
        return 0, false // キューが空
    }
    value := q.data[q.head]
    q.head = (q.head + 1) % q.capacity
    q.size--
    return value, true
}

// Peek: キューの先頭要素を参照
func (q *ArrayQueue) Peek() (int, bool) {
    if q.size == 0 {
        return 0, false
    }
    return q.data[q.head], true
}

func main() {
    queue := NewArrayQueue(5)

    // Enqueue操作
    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    // Peek操作
    if front, ok := queue.Peek(); ok {
        fmt.Println("先頭要素:", front) // 先頭要素: 10
    }

    // Dequeue操作
    for i := 0; i < 3; i++ {
        if value, ok := queue.Dequeue(); ok {
            fmt.Println("取り出した要素:", value)
        }
    }
}
```

## スライスを使用したキューの実装

```go
package main

import "fmt"

// キュー構造体
type Queue struct {
    data []int
}

// Enqueue: キューに要素を追加
func (q *Queue) Enqueue(value int) {
    q.data = append(q.data, value)
}

// Dequeue: キューから要素を取り出す
func (q *Queue) Dequeue() (int, bool) {
    if len(q.data) == 0 {
        return 0, false // キューが空の場合
    }
    value := q.data[0]
    q.data = q.data[1:]
    return value, true
}

// Peek: キューの先頭要素を参照
func (q *Queue) Peek() (int, bool) {
    if len(q.data) == 0 {
        return 0, false // キューが空の場合
    }
    return q.data[0], true
}

// IsEmpty: キューが空かどうかを確認
func (q *Queue) IsEmpty() bool {
    return len(q.data) == 0
}

// Size: キューの要素数を取得
func (q *Queue) Size() int {
    return len(q.data)
}

func main() {
    queue := &Queue{}

    // Enqueue操作
    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    // Peek操作
    if front, ok := queue.Peek(); ok {
        fmt.Println("先頭要素:", front) // 先頭要素: 10
    }

    // Dequeue操作
    for !queue.IsEmpty() {
        if value, ok := queue.Dequeue(); ok {
            fmt.Println("取り出した要素:", value)
        }
    }

    // キューが空か確認
    fmt.Println("キューが空:", queue.IsEmpty()) // キューが空: true
}
```

## 連結リストを用いた動的サイズのキューの実装

以下は、連結リストを使用してキューを実装する例です。

```go
package main

import "fmt"

// ノード構造体
type Node struct {
    value int
    next  *Node
}

// キュー構造体
type LinkedListQueue struct {
    head *Node
    tail *Node
    size int
}

// Enqueue: キューに要素を追加
func (q *LinkedListQueue) Enqueue(value int) {
    newNode := &Node{value: value}
    if q.tail != nil {
        q.tail.next = newNode
    }
    q.tail = newNode
    if q.head == nil {
        q.head = newNode
    }
    q.size++
}

// Dequeue: キューから要素を取り出す
func (q *LinkedListQueue) Dequeue() (int, bool) {
    if q.head == nil {
        return 0, false
    }
    value := q.head.value
    q.head = q.head.next
    if q.head == nil {
        q.tail = nil
    }
    q.size--
    return value, true
}

// Peek: キューの先頭要素を参照
func (q *LinkedListQueue) Peek() (int, bool) {
    if q.head == nil {
        return 0, false
    }
    return q.head.value, true
}

// IsEmpty: キューが空かどうかを確認
func (q *LinkedListQueue) IsEmpty() bool {
    return q.size == 0
}

// Size: キューの要素数を取得
func (q *LinkedListQueue) Size() int {
    return q.size
}

func main() {
    queue := &LinkedListQueue{}

    // Enqueue操作
    queue.Enqueue(10)
    queue.Enqueue(20)
    queue.Enqueue(30)

    // Peek操作
    if front, ok := queue.Peek(); ok {
        fmt.Println("先頭要素:", front) // 先頭要素: 10
    }

    // Dequeue操作
    for !queue.IsEmpty() {
        if value, ok := queue.Dequeue(); ok {
            fmt.Println("取り出した要素:", value)
        }
    }

    // キューが空か確認
    fmt.Println("キューが空:", queue.IsEmpty()) // キューが空: true
}
```

# キューの応用例

1. **タスクスケジューリング**: CPU のタスク管理やプロセススケジューリングで使用。
2. **データストリーム処理**: データの順序を保ちながら処理する。
3. **幅優先探索（BFS）**: グラフや木構造の探索アルゴリズム。
4. **プリントジョブ管理**: プリンタのジョブキュー。
5. **ネットワークパケット処理**: パケットの順序を保ちながら処理。

# 他のデータ構造との関係

- **スタック**: LIFO（後入れ先出し）と対照的に、FIFO（先入れ先出し）を採用。
- **配列**: 配列を使用してキューを実装可能。
- **連結リスト**: 動的サイズのキューを実装可能。
- **デック（Deque）**: 両端キューとして、キューの拡張版。

# 他のデータ構造との比較

| 特徴                    | キュー          | スタック         | デック         | リングバッファ | 配列           | 連結リスト     |
| ----------------------- | --------------- | ---------------- | -------------- | -------------- | -------------- | -------------- |
| **順序性**              | FIFO            | LIFO             | 両端操作可能   | FIFO           | 任意           | 任意           |
| **ランダムアクセス**    | 不可            | 不可             | 不可           | 不可           | 可能           | 不可           |
| **サイズ変更**          | 動的または固定  | 動的または固定   | 動的または固定 | 固定           | 固定または動的 | 動的           |
| **操作の計算量 (追加)** | O(1)            | O(1)             | O(1)           | O(1)           | O(1)平均       | O(1)           |
| **操作の計算量 (削除)** | O(1)            | O(1)             | O(1)           | O(1)           | O(n)           | O(1)           |
| **メモリ効率**          | 中程度          | 中程度           | 中程度         | 高い           | 高い           | 中程度         |
| **用途**                | タスク管理、BFS | 関数呼び出し管理 | 柔軟な両端操作 | バッファ管理   | 汎用データ格納 | 動的データ管理 |
