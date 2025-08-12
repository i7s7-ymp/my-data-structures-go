# 連結リスト（Linked List）とは

連結リストは、ノードと呼ばれる要素がポインタで繋がれたデータ構造です。各ノードはデータと次のノードへの参照（ポインタ）を持ちます。

ポインタが既知なノードに対して、要素の挿入と削除を高速に実行できるため、頻繁な挿入・削除が必要な場面で広く利用されます。

配列のように連続したメモリ領域に格納されているのではなく、飛び飛びのメモリ領域を利用する。メモリ再確保とデータコピーなしに動的にサイズを変更できます。

# 基本的な特徴

- **動的サイズ**: 実行時にサイズを自由に変更可能
- **挿入・削除が効率的**: 任意の位置での挿入・削除が O(1) で実行可能（ポインタが分かっている場合）。
- **非連続メモリ**: 要素が連続したメモリ領域に格納されない
- **ポインタベース**: 各ノードが次のノードへのポインタを持つ
- **メモリ効率**: 必要な分だけメモリを確保（オーバーヘッドあり）
- **順次アクセス**: 先頭から順番にアクセスする必要がある

## 注意点

1. **メモリ効率**: 各ノードがポインタを持つため、配列に比べてメモリ使用量が多い。
2. **キャッシュ効率**: メモリが非連続のため、キャッシュ効率が低い。
3. **ランダムアクセスの非効率性**: インデックスによるアクセスが遅い。

# 連結リストの種類

## 単方向連結リスト

- 各ノードが次のノードへのポインタのみを持つ
- 一方向にのみ移動可能

## 双方向連結リスト

- 各ノードが前後のノードへのポインタを持つ
- 双方向に移動可能

## 循環連結リスト

- 最後のノードが最初のノードを指す
- リング状の構造

# 行える処理

| 機能                   | 説明                                 | 計算量     | 戻り値                         | 得意/苦手 | 補足（その他）                                         |
| ---------------------- | ------------------------------------ | ---------- | ------------------------------ | --------- | ------------------------------------------------------ |
| 先頭への挿入           | 新しいノードを先頭に追加             | O(1)       | なし（void）                   | ✅        | 先頭ノードのポインタのみ更新するため高速。             |
| 先頭からの削除         | 先頭ノードを削除                     | O(1)       | 削除された要素または成功フラグ | ✅        | 先頭ポインタの更新のみで完了。                         |
| 任意位置への挿入       | 指定ノードの後に新しいノードを追加   | O(1)       | なし（void）                   | ✅        | 対象ノードのポインタが既知の場合のみ。                 |
| 任意位置からの削除     | 指定ノードを削除                     | O(1)       | 削除された要素または成功フラグ | ✅        | 対象ノードのポインタが既知の場合のみ。                 |
| 動的サイズ変更         | 実行時にサイズを変更                 | O(1)       | なし（void）                   | ✅        | メモリ再確保やデータコピーが不要。                     |
| 末尾への挿入           | 新しいノードを末尾に追加             | O(n)       | なし（void）                   | ❌        | 末尾まで順次移動が必要。末尾ポインタがあれば O(1)。    |
| 末尾からの削除         | 末尾ノードを削除                     | O(n)       | 削除された要素または成功フラグ | ❌        | 単方向の場合、末尾の前のノードを見つける必要がある。   |
| 位置指定での挿入       | 指定インデックスに新しいノードを追加 | O(n)       | なし（void）                   | ❌        | 対象位置まで順次移動が必要。                           |
| 位置指定での削除       | 指定インデックスのノードを削除       | O(n)       | 削除された要素または成功フラグ | ❌        | 対象位置まで順次移動が必要。                           |
| ランダムアクセス       | インデックスでの要素取得             | O(n)       | 該当要素                       | ❌        | 先頭から順次辿る必要がある。                           |
| 検索                   | 特定の値を持つノードを探索           | O(n)       | 該当ノードまたは null          | ❌        | 線形探索のみ可能。ソートされていても二分探索は困難。   |
| 全要素の走査           | 全ノードを順次処理                   | O(n)       | すべての要素                   | -         | 順次アクセスのため効率的だが、全要素が対象。           |
| サイズ取得             | リストの要素数を取得                 | O(1)       | int（要素数）                  | ✅        | サイズを別途管理している場合。                         |
| 空判定                 | リストが空かどうかを確認             | O(1)       | bool（空なら true）            | ✅        | 先頭ポインタが null かどうかの確認のみ。               |
| 逆方向走査             | 後ろから前への要素処理               | O(n)       | すべての要素                   | ❌        | 単方向連結リストでは逆方向移動不可。双方向なら効率的。 |
| ソート                 | 要素の並び替え                       | O(n log n) | ソート済みリスト               | ❌        | マージソートなどを使用。ランダムアクセス不可で非効率。 |
| 中間要素の直接アクセス | 任意位置の要素に直接アクセス         | O(n)       | 該当要素                       | ❌        | 配列のような直接アクセス不可。                         |
| 複製・コピー           | リスト全体の複製を作成               | O(n)       | 新しいリスト                   | -         | 全ノードを新規作成する必要がある。                     |
| 連結・結合             | 複数のリストを結合                   | O(1)       | 結合されたリスト               | ✅        | 末尾ポインタがあれば効率的。                           |

# Go 言語での実装

## 組み込み双方向連結リスト（container/list）

Go 言語には双方向連結リストの組み込み実装として container/list パッケージが存在します。ただし、単方向連結リストや循環連結リストの専用組み込み構造は存在しません。

以下の特徴・注意点があります。

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

```go
package main

import "fmt"

// ノード構造体
type Node struct {
    Data int
    Next *Node
}

// 単方向連結リスト
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

## 実装方法の選択肢新

| 用途                 | 推奨実装                     | 理由                           |
| -------------------- | ---------------------------- | ------------------------------ |
| 汎用的な双方向リスト | container/list               | - 標準ライブラリで安定している |
| 型安全性が重要       | 自前実装（ジェネリクス使用） | 型安全性を確保できる           |
| 単方向で十分         | 自前単方向実装               | シンプルでメモリ効率が良い     |
| 循環処理が必要       | 自前循環実装                 | 特殊な用途に最適化             |
| 高性能が必要         | 自前実装                     | 用途に特化した最適化が可能     |

# 連結リストの応用例

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

| 操作                 | 連結リスト         | 配列     |
| -------------------- | ------------------ | -------- |
| **ランダムアクセス** | O(n)               | O(1)     |
| **先頭への挿入**     | O(1)               | O(n)     |
| **末尾への挿入**     | O(n)               | O(1)平均 |
| **中間への挿入**     | O(n)               | O(n)     |
| **先頭からの削除**   | O(1)               | O(n)     |
| **末尾からの削除**   | O(n)               | O(1)     |
| **検索**             | O(n)               | O(n)     |
| **メモリ使用量**     | 多い（ポインタ分） | 少ない   |
