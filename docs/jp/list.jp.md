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

# 連結リストの応用例

## 単方向リスト

一方向にしかたどることができないが、構造がシンプルでメモリ消費量が少ない。

- **スタックとキューの実装**
  - スタック (LIFO: Last-In, First-Out): データをリストの先頭に追加（push）し、先頭から削除（pop）することで効率的にスタックを実現できる。先頭への操作は常に O(1) 。
  - キュー (FIFO: First-In, First-Out): リストの末尾にデータを追加（enqueue）し、先頭から削除（dequeue）する。末尾ノードへのポインタを別途保持することで、両方の操作を O(1) で実現できる。OS のタスクスケジューリングやプリンタの印刷待ち行列などに使われる。
- **ハッシュテーブルのチェイン法**
  - ハッシュテーブルで異なるキーが同じハッシュ値（インデックス）になってしまう「衝突（collision）」を解決する方法の一つ。同じインデックスにマッピングされたデータを連結リストでつなげて保持する。この場合、一方向に探索できれば十分なため、シンプルな単方向リストがよく用いられる。
- **メモリ管理のフリーリスト**
  - OS やプログラムが動的にメモリを確保・解放する際、解放されたメモリ領域（空きブロック）を管理するために使われる。空きブロック同士を単方向連結リストでつないでおき、新たなメモリ確保要求があった際にリストから適切なサイズのブロックを探して割り当てる。
- **グラフの隣接リスト表現**
  - グラフ構造において、ある頂点（ノード）から接続されている他の頂点をリストアップする方法として使われる。各頂点ごとに連結リストを持ち、その頂点から辺が出ている先の頂点をリストに追加する。頂点によって接続数が異なっても柔軟に対応できる。
- **タスクスケジューリング**
  - OS やアプリケーションが実行すべきタスクを管理する際、多くの場合キュー構造が用いられる。実行待ちのタスクを連結リスト（キューとして実装）で管理し、先頭から順に処理する。タスクは動的に発生・完了するため、挿入と削除が高速な連結リストが適している。高度な操作が必要な場合は、双方向リストでの実装も適している（後述）

## 双方向リスト

双方向にたどることができ、特定のノードからの挿入・削除が容易になる反面、ポインタ 1 つ分の余計なメモリ領域が必要。

- **Web ブラウザの「戻る」「進む」機能**
  - 閲覧履歴を双方向連結リストで管理する。「戻る」ボタンは前のノードへ移動し、「進む」ボタンは次のノードへ移動する操作に対応する。現在のページで新しいリンクを開くと、現在のノード以降の「進む」履歴は削除され、新しいノードが末尾に追加される。
- **テキストエディタやワープロソフトの Undo/Redo 機能**
  - ユーザーの操作履歴（文字入力、削除、書式変更など）を双方向連結リストで保存する。「元に戻す（Undo）」はリストを逆方向にたどり、「やり直す（Redo）」は順方向にたどる操作に対応する。
- **LRU (Least Recently Used) キャッシュアルゴリズムの実装**
  - キャッシュメモリからデータを追い出す際に、「最も最近使われていない」データを削除するアルゴリズム。双方向連結リストとハッシュマップを組み合わせて実装するのが一般的。データがアクセスされると、そのデータをリストの先頭（最も最近使われた位置）に移動させる。キャッシュがいっぱいになったら、リストの末尾（最も最近使われていない位置）のデータを削除する。双方向リストなら、特定のノードをリストの途中から削除して先頭へ移動させる操作が高速（O(1)）で行える。
- **タスクスケジューリング（タスクキャンセル、優先度の動的な変更）**
  - **タスクのキャンセル**: 実行待ちリストの途中にある特定のタスクをキャンセル（削除）する必要がある場合。双方向連結リストなら、削除したいノードへの参照さえあれば、前後のノードのポインタを O(1) でつなぎ直すことができる。単方向連結リストでは、削除対象の前にあるノードを見つけるためにリストの先頭から走査する必要があり、効率が悪くなる。
  - **優先度の動的な変更**: 実行待ちの途中でタスクの優先度が変わり、リスト内で順序を入れ替える必要がある場合も同様に双方向リストが適している。

# 配列との比較

配列は連続メモリであり、キャッシュの局所性やランダムアクセス、走査に強みがある。
連結リストは非連続メモリであり、要素の挿入・削除におけるポインタ操作の効率性が良い。

### データ構造の比較: 配列と連結リスト

| 特性                                        | 静的配列 (Static Array)                                   | 動的配列 (Dynamic Array)                            | 単方向連結リスト (Singly Linked List) | 双方向連結リスト (Doubly Linked List)       |
| :------------------------------------------ | :-------------------------------------------------------- | :-------------------------------------------------- | :------------------------------------ | :------------------------------------------ |
| **メモリ割り当て**                          | コンパイル時または実行時に連続ブロックを確保 (サイズ固定) | 実行時にヒープ上に連続ブロックを確保 (自動拡張あり) | 非連続 (ノードごとにヒープに割り当て) | 非連続 (ノードごとにヒープに割り当て)       |
| **サイズ変更**                              | 不可 (固定長)                                             | 可能 (再確保とデータコピーが発生)                   | 容易 (ノード単位で追加・削除)         | 容易 (ノード単位で追加・削除)               |
| **ランダムアクセス**<br/>(インデックス指定) | $O(1)$                                                    | $O(1)$                                              | $O(N)$                                | $O(N)$                                      |
| **検索時間**<br/>(値で探索)                 | $O(N)$<br/>(ソート済みなら $O(\log N)$)                   | $O(N)$<br/>(ソート済みなら $O(\log N)$)             | $O(N)$                                | $O(N)$                                      |
| **挿入・削除 (先頭)**                       | $O(N)$ (全要素シフト)                                     | $O(N)$ (全要素シフト)                               | $O(1)$                                | $O(1)$                                      |
| **挿入・削除 (中間)**                       | $O(N)$ (要素シフト)                                       | $O(N)$ (要素シフト)                                 | $O(N)$ ※注 1                          | $O(N)$ ※注 1                                |
| **挿入・削除 (末尾)**                       | $O(1)$ (アクセスのみ)                                     | 償却 $O(1)$ (挿入時)                                | 挿入: $O(1)$ / 削除: $O(N)$ \*※注 2   | $O(1)$ ※\*注 2                              |
| **メモリオーバーヘッド**                    | 最小 (データ本体のみ)                                     | 低い (データ+管理情報)                              | 高い (データ + 次へのポインタ 1 つ)   | 非常に高い (データ + 前後へのポインタ 2 つ) |
| **キャッシュ局所性**                        | 非常に良い                                                | 非常に良い                                          | 悪い (メモリ上に分散するため)         | 悪い (メモリ上に分散するため)               |

\*注 1: 中間ノードの挿入・削除について (連結リスト)

- 表中の O(N) は、挿入・削除したい位置（または特定のノード）を探すための探索コストを含む
- もし対象ノードへのポインタ（参照）が既知の場合、ノード自体のつなぎ替え操作は双方向リストなら O(1) となる
- 単方向リストの場合、削除操作は「削除対象の一つ前のノード」を見つけるために結局 O(N) の探索が必要になることがある（挿入は対象ノードの直後なら O(1)）

\*\*注 2: 末尾ノードの操作について (連結リスト)

- 計算量は、リストが末尾ノードへのポインタ（tail pointer）を保持していることを前提とする
- 単方向リストの末尾削除: 末尾ノードを削除するには、その一つ前のノードのポインタを null にする必要があるが、単方向リストでは末尾ノードから一つ前のノードへ戻れないため、先頭から N−1 番目のノードまで走査する必要があり、O(N) かかる
- 双方向リストの末尾削除: 末尾ノードから prev ポインタで一つ前のノードに O(1) でアクセスできるため、削除操作全体も O(1) で完了する
