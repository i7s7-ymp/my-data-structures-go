# スタック（Stack）とは

スタックは、LIFO（Last In, First Out：後入れ先出し）の原則に従ってデータを管理する線形データ構造。最後に追加された要素が最初に取り出される特性を持つ。例えば、本を積み重ねていき、取るときは最後に積み重ねた一番上の本からとる場合の動作と同じである。

**要素の追加・削除・参照がすべて O(1)の一定時間で実行できる。** この高速性により、関数呼び出しの管理、式の評価、深度優先探索など、多くのアルゴリズムの基盤として活用されている。

## 特徴

- **LIFO 原則**: 最後に入れたものが最初に出る
- **単一アクセスポイント**: トップ（先頭）からのみアクセス可能
- **動的サイズ**: 実行時にサイズを変更可能
- **順序性**: 要素の挿入順序を逆順で取り出す
- **制限されたアクセス**: 中間要素への直接アクセス不可

## 基本的な操作

- **Push**: スタックの最上部（トップ）に要素を追加する
- **Pop**: スタックのトップから要素を取り出し、削除する
- **Peek(または Top)**: スタックのトップにある要素を参照する（削除しない）

# 行える処理

# 注意点

1. **サイズ制限**: 固定サイズのスタックを使用する場合、サイズを超えるとエラーが発生する
2. **スタックオーバーフロー**: 動的サイズのスタックの場合、メモリ不足や再起の深すぎる呼び出しによってスタックオーバーフローが発生する可能性がある。サイズ変更の制限を設ける必要がある。

# 実装方法

スタックは、配列、スライス、または連結リストを使用して実装できます。

## 配列・スライス・連結リストによる実装の比較

| 実装方法       | メリット                                                                                                                     | デメリット                                                                                           |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| **配列**       | - メモリが連続しているためキャッシュ効率が高い<br>- 実装が簡単で高速<br>- 再帰処理の深さ制限や固定サイズのバッファ管理に有用 | - サイズが固定されており、事前に適切なサイズを決める必要がある<br>- サイズ超過時にエラーが発生       |
| **スライス**   | - 動的にサイズを変更可能<br>- 標準ライブラリの`append`を活用できるため実装が簡単                                             | - サイズ変更時にメモリ再割り当てが発生し、オーバーヘッドが増える<br>- メモリの断片化が発生する可能性 |
| **連結リスト** | - サイズが動的に変化するためメモリ効率が良い<br>- 要素の追加・削除が O(1)                                                    | - メモリが非連続のためキャッシュ効率が低い<br>- 各ノードにポインタを持つためメモリ使用量が増える     |

## 配列を使用したスタックの実装

- **固定サイズ**: 配列を使用するため、スタックのサイズは固定。
- **メモリ効率**: 必要な分だけメモリを使用するため、メモリの無駄が少ない。
- **インデックス管理**: スタックのトップをインデックスで管理するため、実装が簡単。

```go
package main

import "fmt"

// スタック構造体
type ArrayStack struct {
    data []int
    top  int
}

// 新しいスタックを作成
func NewArrayStack(size int) *ArrayStack {
    return &ArrayStack{
        data: make([]int, size),
        top:  -1,
    }
}

// Push: 要素を追加
func (s *ArrayStack) Push(value int) bool {
    if s.top == len(s.data)-1 {
        return false // スタックが満杯
    }
    s.top++
    s.data[s.top] = value
    return true
}

// Pop: 要素を取り出す
func (s *ArrayStack) Pop() (int, bool) {
    if s.top == -1 {
        return 0, false // スタックが空
    }
    value := s.data[s.top]
    s.top--
    return value, true
}

// Peek: 末尾要素を参照
func (s *ArrayStack) Peek() (int, bool) {
    if s.top == -1 {
        return 0, false
    }
    return s.data[s.top], true
}

func main() {
    stack := NewArrayStack(3)

    // Push操作
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    // Pop操作
    if value, ok := stack.Pop(); ok {
        fmt.Println("取り出した要素:", value) // 取り出した要素: 30
    }

    fmt.Println("スタックの状態:", stack.data[:stack.top+1]) // スタックの状態: [10 20]
}
```

## スライスを使用したスタックの実装

スライスを使って実装することが多い。

- **スライスを使用**: スタックのデータを保持するためにスライスを使用
- **O(1)の Push 操作**: 組み込みの `append` により、`Push` を O(1) で実行可能
- **O(1)の Pop 操作**: スライスを再スライスする s = s[:len(s)-1] で `Pop` を O(1) で実現可能
- **エラーハンドリング**: スタックが空の場合に適切なエラー処理を実装

```go
package main

import "fmt"

// スタック構造体
type SliceStack struct {
    data []int
}

// 新しいスタックを作成
func NewSliceStack() *SliceStack {
    return &SliceStack{
        data: []int{},
    }
}

// Push: 要素を追加
func (s *SliceStack) Push(value int) {
    s.data = append(s.data, value)
}

// Pop: 要素を取り出す
func (s *SliceStack) Pop() (int, bool) {
    if len(s.data) == 0 {
        return 0, false // スタックが空
    }
    value := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return value, true
}

// Peek: 末尾要素を参照
func (s *SliceStack) Peek() (int, bool) {
    if len(s.data) == 0 {
        return 0, false
    }
    return s.data[len(s.data)-1], true
}

func main() {
    stack := NewSliceStack()

    // Push操作
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    // Pop操作
    if value, ok := stack.Pop(); ok {
        fmt.Println("取り出した要素:", value) // 取り出した要素: 30
    }

    fmt.Println("スタックの状態:", stack.data) // スタックの状態: [10 20]
}
```

## 連結リストを使用したスタックの実装

- **動的サイズ**: 連結リストを使用することで、スタックのサイズを動的に変更可能。
- **メモリ効率**: 必要な分だけメモリを使用するため、メモリの無駄が少ない。
- **ポインタ管理**: ノード間の接続をポインタで管理するため、実装がやや複雑。

```go
package main

import "fmt"

// ノード構造体
type Node struct {
    value int
    next  *Node
}

// スタック構造体
type LinkedListStack struct {
    top *Node
}

// 新しいスタックを作成
func NewLinkedListStack() *LinkedListStack {
    return &LinkedListStack{}
}

// Push: 要素を追加
func (s *LinkedListStack) Push(value int) {
    newNode := &Node{value: value, next: s.top}
    s.top = newNode
}

// Pop: 要素を取り出す
func (s *LinkedListStack) Pop() (int, bool) {
    if s.top == nil {
        return 0, false // スタックが空
    }
    value := s.top.value
    s.top = s.top.next
    return value, true
}

// Peek: 末尾要素を参照
func (s *LinkedListStack) Peek() (int, bool) {
    if s.top == nil {
        return 0, false
    }
    return s.top.value, true
}

func main() {
    stack := NewLinkedListStack()

    // Push操作
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    // Pop操作
    if value, ok := stack.Pop(); ok {
        fmt.Println("取り出した要素:", value) // 取り出した要素: 30
    }

    fmt.Println("スタックの状態: 連結リストのため直接表示不可")
}
```

# スタックの応用例

- **関数呼び出しスタック**: プログラムが関数を呼び出すたびに、その関数のコンテキスト（ローカル変数、リターンアドレスなど）がコールスタックにプッシュされる。関数が終了すると、そのコンテキストがポップされ、実行が呼び出し元に戻る。再帰関数の実現もこの仕組みに基づいている。
- **アンドゥ・リドゥ機能**: テキストエディタや画像編集ソフトなどで、ユーザーが行った操作をスタックにプッシュしていく。「元に戻す（Undo）」操作は、スタックから最新の操作をポップしてその逆の操作を実行することに対応する。
- **式評価と構文解析**: 中置記法（例: 3 + 4 _ 2）の数式を後置記法（例: 3 4 2 _ +）に変換したり、後置記法の式を評価したりする際にスタックが用いられる。また、コンパイラがソースコードの構文（括弧の対応など）を解析する際にも利用される。
- **バックトラッキングアルゴリズム**: 迷路探索や数独ソルバーのような問題で、ある経路を進む際にその状態をスタックにプッシュする。行き止まりに達した場合、スタックをポップして前の状態に戻り、別の経路を試す（バックトラックする）。
