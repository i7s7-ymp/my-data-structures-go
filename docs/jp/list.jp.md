# 連結リスト（Linked List）とは

連結リストは、ノードと呼ばれる要素がポインタで繋がれたデータ構造で、各ノードはデータと次のノードへの参照（ポインタ）を持つ。

配列のように連続したメモリ領域に格納されているのではなく、飛び飛びのメモリ領域を利用する。メモリ再確保とデータコピーなしに動的にサイズを変更できる。また、ポインタが既知なノードに対して、要素の挿入と削除を高速に実行できる。この動的な性質と変更の容易さが連結リストを利用する際の利点となる。

## 特徴

- **動的サイズ**: 実行時にサイズを自由に変更可能
- **挿入・削除が効率的**: 任意の位置での挿入・削除が O(1) で実行可能（ポインタが分かっている場合）。
- **非連続メモリ**: 要素が連続したメモリ領域に格納されない
- **ポインタベース**: 各ノードが次のノードへのポインタを持つ
- **メモリ効率**: 必要な分だけメモリを確保（オーバーヘッドあり）
- **順次アクセス**: 先頭から順番にアクセスする必要がある

## 注意点

1. **メモリオーバーヘッド**: 各ノードが次のノードへのポインタを持つため、ポインタ用のメモリ領域が必要。要素数が増えるとこのオーバーヘッドが無視できなくなる。
2. **キャッシュ効率**: メモリが非連続のため、キャッシュ効率が低い。
3. **ランダムアクセスの非効率性**: インデックスによるアクセスが遅い。ポインタではなくインデックス指定で要素を確認したい場合は、リストの先頭から順番に辿るため、O(n)の時間がかかる。

## 種類

- **単方向連結リスト**: 各ノードが次のノードへのポインタのみを持ち、一方向にのみ走査が可能。
- **双方向連結リスト**: 各ノードが前後のノードへのポインタを持ち、双方向への走査が可能。
- **循環連結リスト**: 最後のノードが最初のノードのポインタを持ち、リスト全体がリング状の構造を形成する。

# Go 言語での実装

## 組み込み関数を使う場合（双方向連結リスト、container/list）

Go 言語には双方向連結リストの組み込み実装として container/list パッケージが存在する。ただし、単方向連結リストや循環連結リストの専用組み込み構造は存在しません。

以下の特徴・注意点がある。

- **型安全性**: interface{} を使用するため、型安全性が低い。Go 1.18 以降のジェネリクスは未対応。
- **要素の管理**: \*list.Element を返すため、要素の管理が必要。削除された要素を再利用してはいけない。
- **パフォーマンス**: 挿入・削除: O(1)（要素への参照がある場合）。検索: O(n)。

```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    // 新しいリストを作成
    l := list.New()

    // 要素の追加
    e1 := l.PushBack(1)    // 末尾に追加
    e2 := l.PushBack(2)    // 末尾に追加
    e3 := l.PushFront(0)   // 先頭に追加

    // 現在のリスト: 0 -> 1 -> 2

    // 特定要素の前後に挿入
    l.InsertBefore(1.5, e2)  // e2の前に挿入
    l.InsertAfter(0.5, e3)   // e3の後に挿入

    // リストの走査
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    // 要素の削除
    l.Remove(e1)

    fmt.Printf("Length: %d\n", l.Len())
}
```

主要なメソッド

```go
// 作成
l := list.New()

// 追加
l.PushBack(value)           // 末尾に追加
l.PushFront(value)          // 先頭に追加
l.InsertBefore(v, element)  // 指定要素の前に挿入
l.InsertAfter(v, element)   // 指定要素の後に挿入

// 削除
l.Remove(element)           // 指定要素を削除

// アクセス
l.Front()                   // 先頭要素を取得
l.Back()                    // 末尾要素を取得
l.Len()                     // 長さを取得

// 移動
l.MoveToFront(element)      // 要素を先頭に移動
l.MoveToBack(element)       // 要素を末尾に移動
l.MoveBefore(e, mark)       // eをmarkの前に移動
l.MoveAfter(e, mark)        // eをmarkの後に移動
```

## 自前で端方向連結リストを実装する場合

### 単方向連結リスト

ノード要素とリスト全体を管理する二つの構造体を定義する。

```go
package main

import "fmt"

// ノード構造体（リストの単一要素）
type Node struct {
    Data int
    Next *Node
}

// 単方向連結リスト（リスト全体を管理）
type SinglyLinkedList struct {
    Head *Node
    Size int
}

// 先頭に挿入
func (sll *SinglyLinkedList) InsertAtHead(data int) {
    newNode := &Node{Data: data, Next: sll.Head}
    sll.Head = newNode
    sll.Size++
}

// 末尾に挿入
func (sll *SinglyLinkedList) InsertAtTail(data int) {
    newNode := &Node{Data: data, Next: nil}

    if sll.Head == nil {
        sll.Head = newNode
    } else {
        current := sll.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    sll.Size++
}

// 先頭から削除
func (sll *SinglyLinkedList) DeleteFromHead() bool {
    if sll.Head == nil {
        return false
    }

    sll.Head = sll.Head.Next
    sll.Size--
    return true
}

// 検索
func (sll *SinglyLinkedList) Search(target int) *Node {
    current := sll.Head
    for current != nil {
        if current.Data == target {
            return current
        }
        current = current.Next
    }
    return nil
}

// 表示
func (sll *SinglyLinkedList) Display() {
    current := sll.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}
```

### 双方向連結リスト

```go
- package main

import "fmt"

// 双方向ノード
type DoublyNode struct {
    Data int
    Next *DoublyNode
    Prev *DoublyNode
}

// 双方向連結リスト
type DoublyLinkedList struct {
    Head *DoublyNode
    Tail *DoublyNode
    Size int
}

// 先頭に挿入
func (dll *DoublyLinkedList) InsertAtHead(data int) {
    newNode := &DoublyNode{Data: data, Next: nil, Prev: nil}

    if dll.Head == nil {
        dll.Head = newNode
        dll.Tail = newNode
    } else {
        newNode.Next = dll.Head
        dll.Head.Prev = newNode
        dll.Head = newNode
    }
    dll.Size++
}

// 末尾に挿入
func (dll *DoublyLinkedList) InsertAtTail(data int) {
    newNode := &DoublyNode{Data: data, Next: nil, Prev: nil}

    if dll.Tail == nil {
        dll.Head = newNode
        dll.Tail = newNode
    } else {
        dll.Tail.Next = newNode
        newNode.Prev = dll.Tail
        dll.Tail = newNode
    }
    dll.Size++
}

// 先頭から削除
func (dll *DoublyLinkedList) DeleteFromHead() bool {
    if dll.Head == nil {
        return false
    }

    if dll.Head == dll.Tail {
        dll.Head = nil
        dll.Tail = nil
    } else {
        dll.Head = dll.Head.Next
        dll.Head.Prev = nil
    }
    dll.Size--
    return true
}

// 末尾から削除
func (dll *DoublyLinkedList) DeleteFromTail() bool {
    if dll.Tail == nil {
        return false
    }

    if dll.Head == dll.Tail {
        dll.Head = nil
        dll.Tail = nil
    } else {
        dll.Tail = dll.Tail.Prev
        dll.Tail.Next = nil
    }
    dll.Size--
    return true
}

// 前方向表示
func (dll *DoublyLinkedList) DisplayForward() {
    current := dll.Head
    for current != nil {
        fmt.Printf("%d <-> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}

// 後方向表示
func (dll *DoublyLinkedList) DisplayBackward() {
    current := dll.Tail
    for current != nil {
        fmt.Printf("%d <-> ", current.Data)
        current = current.Prev
    }
    fmt.Println("nil")
}
```

### 循環連結リスト

```go
package main

import "fmt"

// 循環連結リストのノード
type CircularNode struct {
    Data int
    Next *CircularNode
}

// 循環連結リスト
type CircularLinkedList struct {
    Tail *CircularNode  // 末尾ノードを保持（先頭は tail.Next）
    Size int
}

// 挿入
func (cll *CircularLinkedList) Insert(data int) {
    newNode := &CircularNode{Data: data}

    if cll.Tail == nil {
        newNode.Next = newNode  // 自分自身を指す
        cll.Tail = newNode
    } else {
        newNode.Next = cll.Tail.Next  // 現在の先頭を指す
        cll.Tail.Next = newNode       // 末尾の次を新ノードに
        cll.Tail = newNode            // 新ノードを末尾に
    }
    cll.Size++
}

// 先頭から削除
func (cll *CircularLinkedList) DeleteFromHead() bool {
    if cll.Tail == nil {
        return false
    }

    if cll.Tail.Next == cll.Tail {  // 要素が1つの場合
        cll.Tail = nil
    } else {
        head := cll.Tail.Next
        cll.Tail.Next = head.Next
    }
    cll.Size--
    return true
}

// 表示（1周分）
func (cll *CircularLinkedList) Display() {
    if cll.Tail == nil {
        fmt.Println("Empty list")
        return
    }

    current := cll.Tail.Next  // 先頭から開始
    start := current

    for {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
        if current == start {  // 1周した
            break
        }
    }
    fmt.Printf("(back to %d)\n", start.Data)
}
```

# 連結リストの応用例（要確認）

| 使用場面                                 | 推奨タイプ           | 理由                                     |
| ---------------------------------------- | -------------------- | ---------------------------------------- |
| **頻繁な挿入・削除操作**                 | 双方向連結リスト     | 任意位置での O(1) 挿入・削除が可能       |
| **サイズが予測困難なデータ**             | 単方向連結リスト     | 動的メモリ割り当てで柔軟にサイズ変更     |
| **メモリ効率を重視する場合**             | 単方向連結リスト     | 必要最小限のメモリ使用                   |
| **双方向トラバーサル**                   | 双方向連結リスト     | 前後両方向への効率的な移動が可能         |
| **スタック・キューの実装**               | 単方向連結リスト     | 先頭での挿入・削除が O(1) で効率的       |
| **アンドゥ・リドゥ機能**                 | 双方向連結リスト     | 履歴の前後移動が効率的                   |
| **音楽プレイヤーのプレイリスト**         | 循環双方向連結リスト | 次の曲・前の曲への移動、リピート機能     |
| **テキストエディタの実装**               | 双方向連結リスト     | カーソル移動、文字挿入・削除が効率的     |
| **ブラウザの履歴管理**                   | 双方向連結リスト     | 戻る・進むボタンの実装が簡単             |
| **LRU キャッシュの実装**                 | 双方向連結リスト     | 最近使用要素の移動が O(1) で可能         |
| **グラフの隣接リスト表現**               | 単方向連結リスト     | 各ノードの隣接ノードを動的に管理         |
| **多項式の表現**                         | 単方向連結リスト     | 次数の異なる項を効率的に管理             |
| **疎行列の表現**                         | 単方向連結リスト     | 非ゼロ要素のみを効率的に格納             |
| **メモリ制約が厳しい環境**               | 単方向連結リスト     | 最小限のメモリオーバーヘッド             |
| **リアルタイムシステム**                 | 避ける               | メモリ断片化とキャッシュミスのリスク     |
| **大量データの順次処理**                 | 単方向連結リスト     | ストリーミング処理に適している           |
| **イテレータパターンの実装**             | 双方向連結リスト     | 前後への柔軟な移動が可能                 |
| **リンクされたデータ構造**               | 単方向連結リスト     | 他のデータ構造の基盤として使用           |
| **フリーメモリ管理**                     | 単方向連結リスト     | 使用可能メモリブロックの管理             |
| **タスクスケジューリング**               | 双方向連結リスト     | タスクの優先度変更、削除が効率的         |
| **データベースの B+ ツリー**             | 双方向連結リスト     | リーフノード間のリンクで範囲検索を効率化 |
| **ネットワークパケットバッファ**         | 単方向連結リスト     | 可変長パケットの効率的な管理             |
| **ガベージコレクションの実装**           | 単方向連結リスト     | 到達可能オブジェクトの管理               |
| **関数型プログラミングのリスト**         | 単方向連結リスト     | イミュータブルなデータ構造として最適     |
| **組み込みシステム（リアルタイム以外）** | 単方向連結リスト     | 動的メモリ管理が必要な場合               |
| **プロトタイピング・開発初期**           | 双方向連結リスト     | 柔軟な操作が開発速度を向上させる         |

# 配列との比較

配列は連続メモリであり、キャッシュの局所性やランダムアクセス、走査に強みがある。
連結リストは非連続メモリであり、要素の挿入・削除におけるポインタ操作の効率性が良い。

| 特性                     | 配列 / Go スライス                                                | 連結リスト                                         |
| ------------------------ | ----------------------------------------------------------------- | -------------------------------------------------- |
| メモリ割り当て           | 連続したブロック（静的配列はスタック、動的配列/スライスはヒープ） | 非連続（ノードは個別にヒープに割り当て）           |
| アクセス時間（ランダム） | O(1)                                                              | O(N)                                               |
| 検索時間                 | O(N)（線形探索）、ソート済みなら O(logN)（二分探索）              | O(N)                                               |
| 挿入・削除（中間）       | O(N)（要素のシフトが必要）                                        | O(1)（対象ノードが既知の場合）、ノード探索に O(N)  |
| 挿入・削除（末尾）       | 償却 O(1)（スライスの append）                                    | O(1)（末尾ポインタを保持する場合）                 |
| メモリオーバーヘッド     | 低い（要素ごとの追加コストなし）                                  | 高い（要素ごとに 1 つまたは 2 つのポインタを格納） |
| キャッシュ局所性         | 非常に良い（高い空間的局所性）                                    | 悪い（ポインタチェインがキャッシュミスを誘発）     |
| 柔軟性（サイズ）         | 固定長（配列）または再確保を伴う動的（スライス）                  | 高度に動的（ノード単位                             |
