# B-tree（B 木）とは

B-tree（B 木）は、**自己平衡型の多分木データ構造**で、データベースやファイルシステムなどの大容量データの効率的な検索、挿入、削除を可能にするデータ構造です。1970 年代に Rudolf Bayer と Edward McCreight によって開発されました。

**このデータ構造の最大の利点は、ディスク I/O を最小化し、大量データに対して安定した O(log n)の性能を保証できる点です。**

# B-tree の構成要素

## 基本要素

- **ノード（Node）**: キーと子ノードへのポインタを格納する単位
- **キー（Key）**: 各ノードに格納されるデータの値（ソート済み）
- **ポインタ（Pointer）**: 子ノードへの参照
- **次数（Order/Degree）**: ノードが持てる子ノードの最大数（通常 m で表現）
- **ルート（Root）**: 木の最上位ノード
- **葉ノード（Leaf Node）**: 子を持たないノード
- **内部ノード（Internal Node）**: 子を持つノード

## B-tree の性質

- **最小次数（t）**: 各ノード（ルート以外）は最低 t-1 個のキーを持つ
- **最大次数（2t-1）**: 各ノードは最大 2t-1 個のキーを持つ
- **平衡性**: すべての葉ノードが同じレベルに存在
- **順序性**: 各ノード内のキーは昇順でソートされている
- **分岐条件**: 内部ノードのキー数を k とすると、子ノード数は k+1 個

# B-tree の種類

## Standard B-tree

- 基本的な B-tree
- 内部ノードにもデータを格納
- 用途: 一般的なデータベースインデックス

## B+-tree (B Plus Tree)

- すべてのデータを葉ノードにのみ格納
- 内部ノードはインデックスとして機能
- 用途: データベースの主要インデックス、ファイルシステム

## B\*-tree (B Star Tree)

- ノードの使用率を向上させた変種
- 分割前に隣接ノードとの再分散を試行
- 用途: 高いストレージ効率が必要な場合

## 2-3 木 (2-3 Tree)

- 次数 3 の B-tree（t=2）
- 各ノードが 2 または 3 個の子を持つ
- 用途: 教育目的、理論的解析

## 2-3-4 木 (2-3-4 Tree)

- 次数 4 の B-tree（t=2）
- 赤黒木と等価な構造
- 用途: メモリ内データ構造

# B-tree の特徴

- **ディスク I/O 効率**: ノードサイズをディスクブロックサイズに合わせて最適化
- **平衡保証**: 自動的にバランスを維持し、最悪計算量を保証
- **範囲検索**: 順序性により効率的な範囲クエリが可能
- **高い分岐度**: 木の高さを低く保ち、検索効率を向上
- **動的構造**: 挿入・削除時に自動的に構造を調整

# B-tree の構成要素の図解

## 基本的な B-tree（次数 5、t=3）の例

```
                    [10, 20]
                   /    |    \
                  /     |     \
             [5, 7]  [13, 17]  [25, 30]
            /  |  \   /  |  \   /  |  \
           1   6  8  11 15 18  22 27 35
```

### 各要素の詳細説明

**次数 m = 5（最大子ノード数）**
**最小次数 t = 3**
**各ノード：最低 2 個、最大 4 個のキー**

## B-tree の構造詳細

### ノードの内部構造

```
ノード構造（最大4キー、5ポインタの例）:
┌─────────────────────────────────────────┐
│ P0 │ K1 │ P1 │ K2 │ P2 │ K3 │ P3 │ K4 │ P4 │
└─────────────────────────────────────────┘

where:
- Ki: i番目のキー
- Pi: i番目の子ノードへのポインタ
- P0の子ノード: K1より小さいキーを持つ
- Piの子ノード: Ki < キー < Ki+1の範囲
```

### キーとポインタの関係

```
ルートノード [10, 20] の場合:

      P0     P1     P2
       |      |      |
       v      v      v
   (<10)   (10-20)  (>20)

左の子: 10未満のキー
中央の子: 10以上20未満のキー
右の子: 20以上のキー
```

## B-tree 操作の例

### 挿入操作の例

```
初期状態（t=3、最大4キー）:
           [10, 20, 30]
          /    |    |   \
         5   [15]  25  [35, 40]

キー12を挿入:
1. ルートから検索: 10 < 12 < 20 → 中央の子[15]へ
2. [15]に12を挿入: [12, 15]

結果:
           [10, 20, 30]
          /    |    |   \
         5  [12,15] 25  [35, 40]
```

### 分割操作の例

```
ノード [5, 10, 15, 20] にキー12を挿入する場合:
1. ノードが満杯なので分割が必要
2. 中央値15を親に昇格
3. 左ノード[5, 10, 12] と右ノード[20] に分割

分割前:
     親ノード
        |
   [5,10,15,20] ← 満杯

分割後:
     親ノード[15]
     /          \
[5,10,12]      [20]
```

## B+-tree の構造例

```
B+-tree（データは葉のみ）:
            [10, 20]        ← 内部ノード（インデックス）
           /    |    \
          /     |     \
      [5,7]  [13,17]  [25,30] ← 葉ノード（実データ）
       ↔      ↔       ↔     ← 葉ノード間の連結

特徴:
- 内部ノードはインデックスのみ
- 実データはすべて葉ノードに格納
- 葉ノード間が連結（範囲検索に効率的）
```

# B-tree で行える処理

## 基本操作

| 機能        | 説明           | 計算量       | 戻り値            | 得意/苦手 | 補足（その他）             |
| ----------- | -------------- | ------------ | ----------------- | --------- | -------------------------- |
| Search      | キーを検索     | O(log n)     | ノードまたは null | ✅        | ルートから葉への一方向検索 |
| Insert      | キーを挿入     | O(log n)     | bool（成功/失敗） | ✅        | 分割操作で平衡を維持       |
| Delete      | キーを削除     | O(log n)     | bool（成功/失敗） | ✅        | マージ操作で平衡を維持     |
| FindMin     | 最小値を検索   | O(log n)     | 最小キー          | ✅        | 最左の葉ノードまで辿る     |
| FindMax     | 最大値を検索   | O(log n)     | 最大キー          | ✅        | 最右の葉ノードまで辿る     |
| RangeQuery  | 範囲検索       | O(log n + k) | k 個の結果        | ✅        | B-tree の最大の利点        |
| Predecessor | 前駆要素を検索 | O(log n)     | 前駆キー          | ✅        | 順序性を利用               |
| Successor   | 後続要素を検索 | O(log n)     | 後続キー          | ✅        | 順序性を利用               |
| GetHeight   | 木の高さを取得 | O(1)         | int（高さ）       | ✅        | 平衡性により予測可能       |
| GetSize     | ノード数を取得 | O(n)         | int（ノード数）   | ⚠️        | 全ノードの走査が必要       |

## 範囲・統計操作

| 機能        | 説明           | 計算量       | 戻り値            | 得意/苦手 | 補足（その他）         |
| ----------- | -------------- | ------------ | ----------------- | --------- | ---------------------- |
| RangeCount  | 範囲内の要素数 | O(log n)     | int（個数）       | ✅        | 効率的な範囲クエリ     |
| RangeSum    | 範囲内の合計値 | O(log n + k) | 数値（合計）      | ✅        | セグメント木的使用     |
| RangeUpdate | 範囲更新       | O(log n + k) | bool（成功/失敗） | ⚠️        | 遅延伝播が必要         |
| KthElement  | k 番目の要素   | O(log n)     | k 番目のキー      | ✅        | サイズ情報が必要       |
| Rank        | 要素の順位     | O(log n)     | int（順位）       | ✅        | サイズ情報が必要       |
| Split       | 木を分割       | O(log n)     | 分割された木      | ✅        | B-tree の柔軟性        |
| Merge       | 木を結合       | O(log n)     | 結合された木      | ✅        | 効率的な結合           |
| BulkLoad    | 一括挿入       | O(n)         | bool（成功/失敗） | ✅        | ソート済みデータに最適 |

## 永続化・メンテナンス

| 機能        | 説明             | 計算量 | 戻り値            | 得意/苦手 | 補足（その他）      |
| ----------- | ---------------- | ------ | ----------------- | --------- | ------------------- |
| Serialize   | ディスクに保存   | O(n)   | bool（成功/失敗） | ✅        | ディスク I/O 最適化 |
| Deserialize | ディスクから読込 | O(n)   | bool（成功/失敗） | ✅        | 遅延読み込み可能    |
| Validate    | 構造の妥当性検証 | O(n)   | bool（有効/無効） | ✅        | デバッグ・テスト用  |
| Rebalance   | 明示的な再平衡   | O(n)   | bool（成功/失敗） | ⚠️        | 通常は自動調整      |
| Compact     | ストレージ最適化 | O(n)   | bool（成功/失敗） | ✅        | 断片化解消          |
| Clone       | 木の複製         | O(n)   | 新しい B-tree     | ✅        | メモリ内コピー      |
| Statistics  | 統計情報取得     | O(n)   | 統計構造体        | ✅        | 性能分析用          |

## 苦手な処理・制限

| 機能                 | 説明                     | 計算量   | 戻り値     | 得意/苦手 | 補足（その他）           |
| -------------------- | ------------------------ | -------- | ---------- | --------- | ------------------------ |
| RandomAccess         | インデックス直接アクセス | O(log n) | 要素       | ⚠️        | 配列より遅い             |
| Iteration            | 全要素の反復処理         | O(n)     | イテレータ | ⚠️        | ディスクアクセスが多い   |
| MemoryCompact        | メモリ使用量最小化       | -        | -          | ❌        | 内部断片化は避けられない |
| SingleElement        | 単一要素のみの管理       | -        | -          | ❌        | オーバーヘッドが大きい   |
| FrequentSmallUpdates | 頻繁な小更新             | -        | -          | ❌        | ディスク I/O が頻発      |
| StringKeys           | 可変長文字列キー         | -        | -          | ⚠️        | ノードサイズ管理が複雑   |

# B-tree の実装方法

## 一般的な実装方法

### ディスクベース実装

- ノードをディスクページ単位で管理
- ページサイズ（4KB、8KB など）に合わせてノード設計
- バッファプール管理による効率的な I/O

### メモリベース実装

- ノードをメモリ内で管理
- ポインタベースの実装
- キャッシュ効率を重視した設計

## Go による B-tree の実装例

```go
package btree

import (
    "fmt"
    "sort"
)

// B-treeノード構造体
type Node struct {
    Keys     []int    // キーの配列
    Children []*Node  // 子ノードへのポインタ
    IsLeaf   bool     // 葉ノードかどうか
    NumKeys  int      // 現在のキー数
}

// B-tree構造体
type BTree struct {
    Root     *Node    // ルートノード
    MinDegree int     // 最小次数 t
}

// 新しいB-treeを作成
func NewBTree(minDegree int) *BTree {
    return &BTree{
        Root:      nil,
        MinDegree: minDegree,
    }
}

// 新しいノードを作成
func (bt *BTree) newNode(isLeaf bool) *Node {
    maxKeys := 2*bt.MinDegree - 1
    return &Node{
        Keys:     make([]int, maxKeys),
        Children: make([]*Node, 2*bt.MinDegree),
        IsLeaf:   isLeaf,
        NumKeys:  0,
    }
}

// キーを検索
func (bt *BTree) Search(key int) bool {
    if bt.Root == nil {
        return false
    }
    return bt.searchNode(bt.Root, key)
}

func (bt *BTree) searchNode(node *Node, key int) bool {
    i := 0
    // 現在のノード内でキーの位置を見つける
    for i < node.NumKeys && key > node.Keys[i] {
        i++
    }

    // キーが見つかった場合
    if i < node.NumKeys && key == node.Keys[i] {
        return true
    }

    // 葉ノードの場合、キーは存在しない
    if node.IsLeaf {
        return false
    }

    // 子ノードを再帰的に検索
    return bt.searchNode(node.Children[i], key)
}

// キーを挿入
func (bt *BTree) Insert(key int) {
    if bt.Root == nil {
        bt.Root = bt.newNode(true)
        bt.Root.Keys[0] = key
        bt.Root.NumKeys = 1
        return
    }

    // ルートが満杯の場合、分割
    if bt.isNodeFull(bt.Root) {
        oldRoot := bt.Root
        bt.Root = bt.newNode(false)
        bt.Root.Children[0] = oldRoot
        bt.splitChild(bt.Root, 0)
    }

    bt.insertNonFull(bt.Root, key)
}

// ノードが満杯かチェック
func (bt *BTree) isNodeFull(node *Node) bool {
    return node.NumKeys == 2*bt.MinDegree-1
}

// 満杯でないノードに挿入
func (bt *BTree) insertNonFull(node *Node, key int) {
    i := node.NumKeys - 1

    if node.IsLeaf {
        // 葉ノードの場合、適切な位置に挿入
        for i >= 0 && node.Keys[i] > key {
            node.Keys[i+1] = node.Keys[i]
            i--
        }
        node.Keys[i+1] = key
        node.NumKeys++
    } else {
        // 内部ノードの場合、適切な子ノードを見つける
        for i >= 0 && node.Keys[i] > key {
            i--
        }
        i++

        // 子ノードが満杯の場合、分割
        if bt.isNodeFull(node.Children[i]) {
            bt.splitChild(node, i)
            if node.Keys[i] < key {
                i++
            }
        }
        bt.insertNonFull(node.Children[i], key)
    }
}

// 子ノードを分割
func (bt *BTree) splitChild(parent *Node, index int) {
    fullChild := parent.Children[index]
    newChild := bt.newNode(fullChild.IsLeaf)

    t := bt.MinDegree
    newChild.NumKeys = t - 1

    // 右半分のキーを新しいノードに移動
    for j := 0; j < t-1; j++ {
        newChild.Keys[j] = fullChild.Keys[j+t]
    }

    // 内部ノードの場合、子ポインタも移動
    if !fullChild.IsLeaf {
        for j := 0; j < t; j++ {
            newChild.Children[j] = fullChild.Children[j+t]
        }
    }

    fullChild.NumKeys = t - 1

    // 親ノードに新しい子ポインタを挿入
    for j := parent.NumKeys; j >= index+1; j-- {
        parent.Children[j+1] = parent.Children[j]
    }
    parent.Children[index+1] = newChild

    // 親ノードに中央キーを挿入
    for j := parent.NumKeys - 1; j >= index; j-- {
        parent.Keys[j+1] = parent.Keys[j]
    }
    parent.Keys[index] = fullChild.Keys[t-1]
    parent.NumKeys++
}

// キーを削除
func (bt *BTree) Delete(key int) bool {
    if bt.Root == nil {
        return false
    }

    deleted := bt.deleteFromNode(bt.Root, key)

    // ルートが空になった場合、新しいルートを設定
    if bt.Root.NumKeys == 0 {
        if bt.Root.IsLeaf {
            bt.Root = nil
        } else {
            bt.Root = bt.Root.Children[0]
        }
    }

    return deleted
}

func (bt *BTree) deleteFromNode(node *Node, key int) bool {
    i := 0
    for i < node.NumKeys && key > node.Keys[i] {
        i++
    }

    if i < node.NumKeys && key == node.Keys[i] {
        // キーが現在のノードに存在
        if node.IsLeaf {
            // 葉ノードからキーを削除
            for j := i; j < node.NumKeys-1; j++ {
                node.Keys[j] = node.Keys[j+1]
            }
            node.NumKeys--
            return true
        } else {
            // 内部ノードからキーを削除
            return bt.deleteFromInternalNode(node, i)
        }
    } else if !node.IsLeaf {
        // キーが子ノードに存在する可能性
        return bt.deleteFromChild(node, i, key)
    }

    return false
}

// 内部ノードからキーを削除
func (bt *BTree) deleteFromInternalNode(node *Node, index int) bool {
    key := node.Keys[index]
    leftChild := node.Children[index]
    rightChild := node.Children[index+1]

    t := bt.MinDegree

    if leftChild.NumKeys >= t {
        // 左の子から最大値を取得して置換
        predecessor := bt.getPredecessor(leftChild)
        node.Keys[index] = predecessor
        return bt.deleteFromNode(leftChild, predecessor)
    } else if rightChild.NumKeys >= t {
        // 右の子から最小値を取得して置換
        successor := bt.getSuccessor(rightChild)
        node.Keys[index] = successor
        return bt.deleteFromNode(rightChild, successor)
    } else {
        // 両方の子ノードが最小サイズの場合、マージ
        bt.merge(node, index)
        return bt.deleteFromNode(leftChild, key)
    }
}

// 子ノードからキーを削除
func (bt *BTree) deleteFromChild(node *Node, index int, key int) bool {
    child := node.Children[index]
    t := bt.MinDegree

    if child.NumKeys < t {
        // 子ノードが最小サイズの場合、借用またはマージ
        if index > 0 && node.Children[index-1].NumKeys >= t {
            bt.borrowFromPrev(node, index)
        } else if index < node.NumKeys && node.Children[index+1].NumKeys >= t {
            bt.borrowFromNext(node, index)
        } else {
            if index == node.NumKeys {
                index--
            }
            bt.merge(node, index)
            child = node.Children[index]
        }
    }

    return bt.deleteFromNode(child, key)
}

// 前駆ノード（最大値）を取得
func (bt *BTree) getPredecessor(node *Node) int {
    for !node.IsLeaf {
        node = node.Children[node.NumKeys]
    }
    return node.Keys[node.NumKeys-1]
}

// 後続ノード（最小値）を取得
func (bt *BTree) getSuccessor(node *Node) int {
    for !node.IsLeaf {
        node = node.Children[0]
    }
    return node.Keys[0]
}

// 前の兄弟から借用
func (bt *BTree) borrowFromPrev(parent *Node, index int) {
    child := parent.Children[index]
    sibling := parent.Children[index-1]

    // 子ノードのキーを右にシフト
    for i := child.NumKeys - 1; i >= 0; i-- {
        child.Keys[i+1] = child.Keys[i]
    }

    // 内部ノードの場合、子ポインタもシフト
    if !child.IsLeaf {
        for i := child.NumKeys; i >= 0; i-- {
            child.Children[i+1] = child.Children[i]
        }
        child.Children[0] = sibling.Children[sibling.NumKeys]
    }

    // 親のキーを子に移動
    child.Keys[0] = parent.Keys[index-1]

    // 兄弟の最大キーを親に移動
    parent.Keys[index-1] = sibling.Keys[sibling.NumKeys-1]

    child.NumKeys++
    sibling.NumKeys--
}

// 次の兄弟から借用
func (bt *BTree) borrowFromNext(parent *Node, index int) {
    child := parent.Children[index]
    sibling := parent.Children[index+1]

    // 親のキーを子に移動
    child.Keys[child.NumKeys] = parent.Keys[index]

    // 内部ノードの場合、子ポインタも移動
    if !child.IsLeaf {
        child.Children[child.NumKeys+1] = sibling.Children[0]
    }

    // 兄弟の最小キーを親に移動
    parent.Keys[index] = sibling.Keys[0]

    // 兄弟のキーを左にシフト
    for i := 1; i < sibling.NumKeys; i++ {
        sibling.Keys[i-1] = sibling.Keys[i]
    }

    // 内部ノードの場合、子ポインタもシフト
    if !sibling.IsLeaf {
        for i := 1; i <= sibling.NumKeys; i++ {
            sibling.Children[i-1] = sibling.Children[i]
        }
    }

    child.NumKeys++
    sibling.NumKeys--
}

// ノードをマージ
func (bt *BTree) merge(parent *Node, index int) {
    child := parent.Children[index]
    sibling := parent.Children[index+1]
    t := bt.MinDegree

    // 親のキーをchildに移動
    child.Keys[t-1] = parent.Keys[index]

    // siblingのキーをchildに移動
    for i := 0; i < sibling.NumKeys; i++ {
        child.Keys[i+t] = sibling.Keys[i]
    }

    // 内部ノードの場合、子ポインタも移動
    if !child.IsLeaf {
        for i := 0; i <= sibling.NumKeys; i++ {
            child.Children[i+t] = sibling.Children[i]
        }
    }

    // 親のキーを左にシフト
    for i := index + 1; i < parent.NumKeys; i++ {
        parent.Keys[i-1] = parent.Keys[i]
    }

    // 親の子ポインタを左にシフト
    for i := index + 2; i <= parent.NumKeys; i++ {
        parent.Children[i-1] = parent.Children[i]
    }

    child.NumKeys += sibling.NumKeys + 1
    parent.NumKeys--
}

// 範囲検索
func (bt *BTree) RangeSearch(min, max int) []int {
    var result []int
    if bt.Root != nil {
        bt.rangeSearchNode(bt.Root, min, max, &result)
    }
    return result
}

func (bt *BTree) rangeSearchNode(node *Node, min, max int, result *[]int) {
    i := 0
    for i < node.NumKeys {
        if !node.IsLeaf {
            // 子ノードを先に処理
            if node.Keys[i] >= min {
                bt.rangeSearchNode(node.Children[i], min, max, result)
            }
        }

        // 現在のキーが範囲内の場合、結果に追加
        if node.Keys[i] >= min && node.Keys[i] <= max {
            *result = append(*result, node.Keys[i])
        }

        if node.Keys[i] > max {
            break
        }
        i++
    }

    // 最後の子ノードを処理
    if !node.IsLeaf && i < len(node.Children) && node.Children[i] != nil {
        if i == 0 || node.Keys[i-1] <= max {
            bt.rangeSearchNode(node.Children[i], min, max, result)
        }
    }
}

// 中順走査（ソート順）
func (bt *BTree) InOrderTraversal() []int {
    var result []int
    if bt.Root != nil {
        bt.inOrderNode(bt.Root, &result)
    }
    return result
}

func (bt *BTree) inOrderNode(node *Node, result *[]int) {
    i := 0
    for i < node.NumKeys {
        if !node.IsLeaf {
            bt.inOrderNode(node.Children[i], result)
        }
        *result = append(*result, node.Keys[i])
        i++
    }

    if !node.IsLeaf {
        bt.inOrderNode(node.Children[i], result)
    }
}

// 最小値を検索
func (bt *BTree) FindMin() (int, bool) {
    if bt.Root == nil {
        return 0, false
    }

    node := bt.Root
    for !node.IsLeaf {
        node = node.Children[0]
    }
    return node.Keys[0], true
}

// 最大値を検索
func (bt *BTree) FindMax() (int, bool) {
    if bt.Root == nil {
        return 0, false
    }

    node := bt.Root
    for !node.IsLeaf {
        node = node.Children[node.NumKeys]
    }
    return node.Keys[node.NumKeys-1], true
}

// 木の高さを取得
func (bt *BTree) GetHeight() int {
    if bt.Root == nil {
        return 0
    }
    return bt.getNodeHeight(bt.Root)
}

func (bt *BTree) getNodeHeight(node *Node) int {
    if node.IsLeaf {
        return 1
    }
    return 1 + bt.getNodeHeight(node.Children[0])
}

// 統計情報構造体
type Statistics struct {
    NodeCount    int
    KeyCount     int
    Height       int
    MinDegree    int
    AvgKeysPerNode float64
}

// 統計情報を取得
func (bt *BTree) GetStatistics() Statistics {
    stats := Statistics{
        MinDegree: bt.MinDegree,
        Height:    bt.GetHeight(),
    }

    if bt.Root != nil {
        bt.collectStats(bt.Root, &stats)
        if stats.NodeCount > 0 {
            stats.AvgKeysPerNode = float64(stats.KeyCount) / float64(stats.NodeCount)
        }
    }

    return stats
}

func (bt *BTree) collectStats(node *Node, stats *Statistics) {
    stats.NodeCount++
    stats.KeyCount += node.NumKeys

    if !node.IsLeaf {
        for i := 0; i <= node.NumKeys; i++ {
            if node.Children[i] != nil {
                bt.collectStats(node.Children[i], stats)
            }
        }
    }
}

// 木の妥当性を検証
func (bt *BTree) Validate() error {
    if bt.Root == nil {
        return nil
    }
    return bt.validateNode(bt.Root, true)
}

func (bt *BTree) validateNode(node *Node, isRoot bool) error {
    t := bt.MinDegree

    // キー数の検証
    if !isRoot && node.NumKeys < t-1 {
        return fmt.Errorf("node has too few keys: %d, minimum: %d", node.NumKeys, t-1)
    }
    if node.NumKeys > 2*t-1 {
        return fmt.Errorf("node has too many keys: %d, maximum: %d", node.NumKeys, 2*t-1)
    }

    // キーの順序性の検証
    for i := 1; i < node.NumKeys; i++ {
        if node.Keys[i-1] >= node.Keys[i] {
            return fmt.Errorf("keys are not in ascending order")
        }
    }

    // 子ノードの検証
    if !node.IsLeaf {
        for i := 0; i <= node.NumKeys; i++ {
            if node.Children[i] == nil {
                return fmt.Errorf("internal node has nil child")
            }
            if err := bt.validateNode(node.Children[i], false); err != nil {
                return err
            }
        }
    }

    return nil
}

// 文字列表現
func (bt *BTree) String() string {
    if bt.Root == nil {
        return "Empty B-tree"
    }
    return bt.nodeString(bt.Root, 0)
}

func (bt *BTree) nodeString(node *Node, depth int) string {
    indent := ""
    for i := 0; i < depth; i++ {
        indent += "  "
    }

    result := fmt.Sprintf("%sKeys: %v\n", indent, node.Keys[:node.NumKeys])

    if !node.IsLeaf {
        for i := 0; i <= node.NumKeys; i++ {
            if node.Children[i] != nil {
                result += bt.nodeString(node.Children[i], depth+1)
            }
        }
    }

    return result
}
```

## B+-tree の実装例（部分）

```go
package btree

// B+-treeノード構造体
type BPlusNode struct {
    Keys     []int       // キーの配列
    Children []*BPlusNode // 子ノードへのポインタ（内部ノードのみ）
    Values   []interface{} // データ値（葉ノードのみ）
    Next     *BPlusNode   // 次の葉ノードへのポインタ（葉ノードのみ）
    IsLeaf   bool         // 葉ノードかどうか
    NumKeys  int          // 現在のキー数
}

// B+-tree構造体
type BPlusTree struct {
    Root      *BPlusNode
    MinDegree int
    LeftMost  *BPlusNode // 最左の葉ノード（範囲検索用）
}

// 範囲検索（B+-tree特有の効率的実装）
func (bpt *BPlusTree) RangeSearchOptimized(min, max int) []interface{} {
    var result []interface{}

    // 開始点を見つける
    node := bpt.findLeafNode(min)
    if node == nil {
        return result
    }

    // 葉ノードを線形に辿る
    for node != nil {
        for i := 0; i < node.NumKeys; i++ {
            if node.Keys[i] >= min && node.Keys[i] <= max {
                result = append(result, node.Values[i])
            } else if node.Keys[i] > max {
                return result
            }
        }
        node = node.Next
    }

    return result
}

// 葉ノードを見つける
func (bpt *BPlusTree) findLeafNode(key int) *BPlusNode {
    if bpt.Root == nil {
        return nil
    }

    node := bpt.Root
    for !node.IsLeaf {
        i := 0
        for i < node.NumKeys && key >= node.Keys[i] {
            i++
        }
        node = node.Children[i]
    }

    return node
}
```

# B-tree の応用例

## 1. データベース管理システム

- **MySQL の InnoDB**: B+-tree をプライマリインデックスに使用
- **PostgreSQL**: B-tree インデックス（実際は B+-tree）
- **Oracle Database**: B-tree インデックスによる高速検索
- **SQLite**: B+-tree によるテーブルとインデックス管理

## 2. ファイルシステム

- **HFS+**: macOS のファイルシステム（B-tree 構造）
- **ReiserFS**: Linux ファイルシステム（B+-tree）
- **NTFS**: Windows ファイルシステム（B+-tree 構造）
- **XFS**: 高性能ファイルシステム（B+-tree）

## 3. NoSQL データベース

- **MongoDB**: B-tree インデックス
- **CouchDB**: B-tree によるデータ格納
- **RocksDB**: LSM-tree と B-tree のハイブリッド
- **BerkeleyDB**: B-tree ストレージエンジン

## 4. インメモリデータベース

- **Redis**: ソート済みセット（B-tree 変種）
- **SAP HANA**: B-tree インデックス
- **VoltDB**: インメモリ B-tree
- **MemSQL**: 高速 B-tree インデックス

## 5. 分散システム

- **Google Bigtable**: B-tree ライクな構造
- **Amazon DynamoDB**: B-tree ベースインデックス
- **Apache Cassandra**: SSTable での B-tree 使用
- **HBase**: B-tree インデックシング

## 6. 組み込みシステム

- **組み込みデータベース**: SQLite ライクなシステム
- **IoT データストレージ**: 効率的なデータ管理
- **車載システム**: 高信頼性データ管理
- **産業制御システム**: リアルタイムデータアクセス

# B-tree と他のデータ構造との比較

## 木構造比較

| 特徴                  | B-tree    | B+-tree         | 二分探索木   | AVL 木    | 赤黒木    | ハッシュテーブル |
| --------------------- | --------- | --------------- | ------------ | --------- | --------- | ---------------- |
| **検索性能**          | O(log n)  | O(log n)        | O(log n)平均 | O(log n)  | O(log n)  | O(1)平均         |
| **挿入性能**          | O(log n)  | O(log n)        | O(log n)平均 | O(log n)  | O(log n)  | O(1)平均         |
| **削除性能**          | O(log n)  | O(log n)        | O(log n)平均 | O(log n)  | O(log n)  | O(1)平均         |
| **範囲検索**          | ✅ 効率的 | ✅ 非常に効率的 | ✅ 効率的    | ✅ 効率的 | ✅ 効率的 | ❌ 不可          |
| **ディスク I/O 効率** | ✅ 最適   | ✅ 最適         | ❌ 非効率    | ❌ 非効率 | ❌ 非効率 | ⚠️ 実装依存      |
| **メモリ効率**        | ✅ 高い   | ✅ 高い         | ⚠️ 中程度    | ⚠️ 中程度 | ⚠️ 中程度 | ⚠️ 負荷率依存    |
| **実装複雑度**        | ❌ 複雑   | ❌ 非常に複雑   | ✅ 簡単      | ❌ 複雑   | ⚠️ 中程度 | ✅ 簡単          |
| **順序保持**          | ✅ あり   | ✅ あり         | ✅ あり      | ✅ あり   | ✅ あり   | ❌ なし          |
| **大量データ**        | ✅ 最適   | ✅ 最適         | ❌ 非効率    | ❌ 非効率 | ❌ 非効率 | ✅ 適している    |

## 用途別最適選択

| 用途                         | 第 1 選択        | 第 2 選択 | 第 3 選択        | 避けるべき           |
| ---------------------------- | ---------------- | --------- | ---------------- | -------------------- |
| **データベースインデックス** | B+-tree          | B-tree    | ハッシュテーブル | 二分探索木           |
| **ファイルシステム**         | B+-tree          | B-tree    | -                | 二分探索木、ハッシュ |
| **大量データ範囲検索**       | B+-tree          | B-tree    | AVL 木           | ハッシュテーブル     |
| **メモリ内高速検索**         | 赤黒木           | AVL 木    | ハッシュテーブル | B-tree               |
| **ディスク効率重視**         | B-tree           | B+-tree   | -                | 二分探索木           |
| **単純なキー検索**           | ハッシュテーブル | B-tree    | 赤黒木           | 二分探索木           |
| **順序保持必須**             | B+-tree          | B-tree    | AVL 木           | ハッシュテーブル     |
| **挿入・削除頻繁**           | B-tree           | 赤黒木    | ハッシュテーブル | 二分探索木           |

## パフォーマンス特性

| データサイズ | B-tree 検索時間 | B-tree 挿入時間 | 二分探索木検索 | ハッシュ検索 |
| ------------ | --------------- | --------------- | -------------- | ------------ |
| **1K 要素**  | ~500ns          | ~800ns          | ~600ns         | ~100ns       |
| **1M 要素**  | ~2μs            | ~5μs            | ~10μs          | ~100ns       |
| **1G 要素**  | ~10μs           | ~50μs           | ~100μs         | ~100ns       |

# まとめ

B-tree（B 木）は、大量データの効率的な管理に特化した自己平衡型多分木データ構造です。ディスク I/O の最小化と安定した O(log n)性能を両立し、データベースやファイルシステムなどの基盤技術として広く採用されています。

Go での実装では、ノードの分割・結合操作を適切に管理し、木の平衡性を維持することが重要です。B+-tree などの変種も含め、用途や要件に応じて最適な設計を選択し、メモリ効率と I/O 効率を両立させることが成功の鍵となります。

特に、大量データや永続化が必要なシステムでは、B-tree の導入により大幅な性能改善とスケーラビリティの向上を実現できます。
