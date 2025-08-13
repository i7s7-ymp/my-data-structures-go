# B+-tree（B+木）とは

B+-tree（B+木）は、**B-tree の改良版**として開発された自己平衡型の多分木データ構造で、データベースやファイルシステムにおける大容量データの効率的な検索、挿入、削除、特に**範囲検索**を可能にするデータ構造です。標準的な B-tree とは異なり、**すべてのデータを葉ノードにのみ格納**し、内部ノードは純粋にインデックスとして機能します。

**このデータ構造の最大の利点は、B-tree の特性を保ちながら、葉ノード間の連結により範囲検索やシーケンシャルアクセスを極めて効率的に行える点です。**

# B+-tree の構成要素

## 基本要素

- **内部ノード（Internal Node）**: キーのみを格納し、子ノードへのポインタを持つインデックス専用ノード
- **葉ノード（Leaf Node）**: 実際のデータ（キーと値のペア）を格納するノード
- **キー（Key）**: 検索に使用される値（内部ノードと葉ノードの両方に存在）
- **値（Value）**: 実際のデータ（葉ノードにのみ存在）
- **ポインタ（Pointer）**: 子ノードまたは次の葉ノードへの参照
- **次数（Order）**: ノードが持てる子ノードの最大数
- **葉ノード連結**: 葉ノード間の双方向または単方向リンク

## B+-tree の性質

- **データ格納**: すべてのデータ（キーと値）は葉ノードにのみ格納
- **内部ノード**: キーのみを持ち、検索のガイドとして機能
- **葉ノード連結**: 全ての葉ノードが連結リストで接続
- **平衡性**: すべての葉ノードが同じレベルに存在
- **順序性**: キーは各ノード内で昇順にソート
- **重複キー**: 内部ノードのキーは葉ノードにも複製される

# B+-tree の種類

## Standard B+-tree

- 基本的な B+-tree
- 葉ノードにデータを格納、内部ノードはインデックス
- 用途: データベースインデックス、ファイルシステム

## Clustered B+-tree

- プライマリキーに基づくクラスタ化インデックス
- データページそのものが B+-tree 構造
- 用途: データベースのプライマリインデックス

## Non-clustered B+-tree

- セカンダリインデックス
- 葉ノードにはデータの位置情報を格納
- 用途: データベースのセカンダリインデックス

## Prefix B+-tree

- キーの共通プレフィックスを圧縮
- メモリ効率を向上
- 用途: 文字列キーの効率的管理

## Bulk-loaded B+-tree

- ソート済みデータからの一括構築
- 構築効率を最適化
- 用途: 大量データの初期インデックス作成

# B+-tree の特徴

- **範囲検索最適化**: 葉ノードの連結により効率的な範囲クエリ
- **シーケンシャルアクセス**: 順次アクセスが非常に高速
- **内部ノード効率**: データを持たないため、より多くのキーを格納可能
- **キャッシュ効率**: 内部ノードのサイズが小さく、メモリ効率が良い
- **ディスク I/O 最適化**: ページサイズに最適化された構造

# B+-tree の構成要素の図解

## 基本的な B+-tree（次数 4）の例

```
内部ノード（インデックスのみ）:
              [20, 40]
             /    |    \
            /     |     \
       [10, 15] [25, 30] [45, 50]
       /  |  \   /  |  \   /  |  \

葉ノード（データ格納）:
[5,10] [15,18] [20,25] [30,35] [40,45] [50,60]
  ↔      ↔       ↔       ↔       ↔       ↔
（葉ノード間の双方向リンク）
```

### B-tree と B+-tree の構造比較

**B-tree（標準）:**

```
           [20, 40]
          /    |    \
    [10,15]  [25,30] [45,50]
   /  |  \   /  |  \  /  |  \
  5   12  18 22  28 35 42 48 55
```

**B+-tree:**

```
内部ノード: [20, 40]          ← インデックスのみ
           /    |    \
     [10,15]  [25,30] [45,50]  ← インデックスのみ
      /||\      /||\     /||\

葉ノード: [5,10,15,18] [20,22,25,28] [30,35,40,42] [45,48,50,55]
          ↔             ↔              ↔              ↔
```

## ノードの内部構造詳細

### 内部ノードの構造

```
内部ノード（キーのみ、値なし）:
┌─────────────────────────────────────┐
│ P0 │ K1 │ P1 │ K2 │ P2 │ K3 │ P3 │
└─────────────────────────────────────┘

where:
- Ki: i番目のキー（インデックス用）
- Pi: i番目の子ノードへのポインタ
- 値（データ）は格納されない
```

### 葉ノードの構造

```
葉ノード（キーと値の両方）:
┌───────────────────────────────────────────────┐
│ K1,V1 │ K2,V2 │ K3,V3 │ K4,V4 │ Next Pointer │
└───────────────────────────────────────────────┘

where:
- Ki: i番目のキー
- Vi: i番目の値（実際のデータ）
- Next Pointer: 次の葉ノードへのポインタ
```

## 検索操作の例

### 単一キー検索（Key=25）

```
1. ルートから開始: [20, 40]
   25 > 20 and 25 < 40 → 中央の子へ

2. 内部ノード: [25, 30]
   25 = 25 → 左の子へ（葉ノード）

3. 葉ノード: [20, 22, 25, 28]
   25を発見 → 対応する値を返す
```

### 範囲検索（Key: 15-35）

```
1. 開始点（15）を検索 → 葉ノード[10, 12, 15, 18]
2. 15から開始して、次のポインタで連結された葉ノードを辿る
3. [10,12,15,18] → [20,22,25,28] → [30,35,40,42]
4. 35に達するまで連続してデータを取得
5. 結果: [15, 18, 20, 22, 25, 28, 30, 35]

利点: ディスクI/Oが最小限（連続した葉ノードのみアクセス）
```

## 挿入操作の例

### キー 27 の挿入

```
挿入前:
葉ノード: [20, 22, 25, 28] （容量満杯と仮定）

挿入処理:
1. 適切な葉ノードを検索
2. ノードが満杯なので分割
3. [20, 22, 25] と [27, 28] に分割
4. 中央キー（25）を親に昇格
5. 新しい葉ノードを連結リストに追加

挿入後:
内部ノード: [25, 30, 40] （25が追加）
葉ノード: [20,22,25] ↔ [27,28] ↔ [30,35]
```

## 削除操作の例

### キー 22 の削除

```
削除前:
葉ノード: [20, 22, 25]

削除処理:
1. 葉ノードから22を削除
2. ノードのキー数が最小値を下回る場合:
   - 隣接ノードから借用
   - または隣接ノードとマージ
3. 内部ノードのインデックスを更新

削除後:
葉ノード: [20, 25] （またはマージ後の新しい構造）
```

# B+-tree で行える処理

## 基本操作

| 機能        | 説明             | 計算量       | 戻り値            | 得意/苦手 | 補足（その他）           |
| ----------- | ---------------- | ------------ | ----------------- | --------- | ------------------------ |
| Search      | キーを検索       | O(log n)     | 値または null     | ✅        | 葉ノードまでの一方向検索 |
| Insert      | キーと値を挿入   | O(log n)     | bool（成功/失敗） | ✅        | 葉ノードのみにデータ挿入 |
| Delete      | キーを削除       | O(log n)     | bool（成功/失敗） | ✅        | 葉ノードからのみ削除     |
| FindMin     | 最小値を検索     | O(log n)     | 最小キー          | ✅        | 最左の葉ノードの最初     |
| FindMax     | 最大値を検索     | O(log n)     | 最大キー          | ✅        | 最右の葉ノードの最後     |
| RangeQuery  | 範囲検索         | O(log n + k) | k 個の結果        | ✅        | B+-tree の最大の利点     |
| Predecessor | 前駆要素を検索   | O(1)         | 前駆キー          | ✅        | 葉ノード内または前の葉   |
| Successor   | 後続要素を検索   | O(1)         | 後続キー          | ✅        | 葉ノード内または次の葉   |
| GetHeight   | 木の高さを取得   | O(1)         | int（高さ）       | ✅        | 平衡性により予測可能     |
| IsEmpty     | 木が空かチェック | O(1)         | bool（空/非空）   | ✅        | ルートノードの存在確認   |

## 範囲・シーケンシャル操作

| 機能           | 説明                 | 計算量       | 戻り値            | 得意/苦手 | 補足（その他）             |
| -------------- | -------------------- | ------------ | ----------------- | --------- | -------------------------- |
| RangeCount     | 範囲内の要素数       | O(log n + k) | int（個数）       | ✅        | 範囲走査で効率的にカウント |
| RangeSum       | 範囲内の合計値       | O(log n + k) | 数値（合計）      | ✅        | 葉ノード連結により効率的   |
| RangeUpdate    | 範囲更新             | O(log n + k) | bool（成功/失敗） | ✅        | 連続する葉ノードを更新     |
| SequentialScan | 順次スキャン         | O(n)         | 全要素            | ✅        | 葉ノード連結の最大活用     |
| FirstN         | 最初の N 要素        | O(log n + N) | N 個の要素        | ✅        | 最左から連続取得           |
| LastN          | 最後の N 要素        | O(log n + N) | N 個の要素        | ✅        | 最右から逆順取得           |
| Skip           | 指定数スキップ後取得 | O(log n + k) | スキップ後の要素  | ✅        | 葉ノード内での高速移動     |
| Paginate       | ページネーション     | O(log n + k) | 1 ページ分の要素  | ✅        | オフセット+リミット処理    |
| PrefixSearch   | プレフィックス検索   | O(log n + k) | 前方一致結果      | ✅        | 文字列キーで特に有効       |

## データベース特化操作

| 機能       | 説明                       | 計算量     | 戻り値            | 得意/苦手 | 補足（その他）               |
| ---------- | -------------------------- | ---------- | ----------------- | --------- | ---------------------------- |
| BulkLoad   | 一括挿入                   | O(n)       | bool（成功/失敗） | ✅        | ソート済みデータの効率的構築 |
| BulkInsert | 複数要素一括挿入           | O(k log n) | int（挿入数）     | ✅        | バッチ処理で効率化           |
| Merge      | 他の B+-tree とマージ      | O(n + m)   | マージされた木    | ✅        | 葉ノード連結により効率的     |
| Split      | 範囲で分割                 | O(log n)   | 分割された木      | ✅        | 指定キーで木を分割           |
| Clone      | インデックスの複製         | O(n)       | 新しい B+-tree    | ✅        | 構造とデータの完全複製       |
| Snapshot   | 現在状態のスナップショット | O(n)       | スナップショット  | ✅        | MVCC 実装のサポート          |
| Statistics | 統計情報取得               | O(n)       | 統計構造体        | ✅        | ノード使用率、高さなど       |

## 永続化・メンテナンス

| 機能             | 説明             | 計算量     | 戻り値              | 得意/苦手 | 補足（その他）               |
| ---------------- | ---------------- | ---------- | ------------------- | --------- | ---------------------------- |
| Serialize        | ディスクに保存   | O(n)       | bool（成功/失敗）   | ✅        | ページ単位での効率的保存     |
| Deserialize      | ディスクから読込 | O(n)       | bool（成功/失敗）   | ✅        | 遅延読み込み対応             |
| Compact          | ストレージ最適化 | O(n)       | bool（成功/失敗）   | ✅        | ページ使用率の最適化         |
| Validate         | 構造の妥当性検証 | O(n)       | bool（有効/無効）   | ✅        | インデックス整合性確認       |
| Rebalance        | 明示的な再平衡   | O(n log n) | bool（成功/失敗）   | ⚠️        | 通常は自動調整               |
| CheckConsistency | 一貫性チェック   | O(n)       | bool（一貫/不整合） | ✅        | 内部ノードと葉ノードの整合性 |
| GetMetadata      | メタデータ取得   | O(1)       | メタデータ構造体    | ✅        | ノード数、高さ、使用率など   |

## 苦手な処理・制限

| 機能                    | 説明                     | 計算量       | 戻り値   | 得意/苦手 | 補足（その他）               |
| ----------------------- | ------------------------ | ------------ | -------- | --------- | ---------------------------- |
| RandomAccess            | インデックス直接アクセス | O(log n)     | 要素     | ⚠️        | 配列のような直接アクセス不可 |
| ReverseRangeQuery       | 逆順範囲検索             | O(log n + k) | 逆順結果 | ⚠️        | 双方向リンクが必要           |
| ComplexJoin             | 複雑な結合操作           | O(n \* m)    | 結合結果 | ❌        | ネストループ結合になりがち   |
| FrequentStructureChange | 頻繁な構造変更           | -            | -        | ❌        | 分割・マージのオーバーヘッド |
| VariableLengthKeys      | 可変長キー               | -            | -        | ⚠️        | ノードサイズ管理が複雑       |
| MultipleSort            | 複数ソート条件           | -            | -        | ❌        | 単一キー順序のみ             |

# B+-tree の実装方法

## 一般的な実装方法

### ディスクベース実装

- ノードをディスクページ単位で管理
- 葉ノードと内部ノードで異なるページレイアウト
- バッファプール管理による効率的なページアクセス

### メモリベース実装

- ノードをメモリ内で管理
- ポインタベースの実装
- 葉ノード間の連結リスト管理

## Go による B+-tree の実装例

```go
package bplustree

import (
    "fmt"
    "sort"
)

// B+-treeノード構造体（基底）
type Node struct {
    Keys     []int
    IsLeaf   bool
    NumKeys  int
}

// 内部ノード構造体
type InternalNode struct {
    Node
    Children []*Node
}

// 葉ノード構造体
type LeafNode struct {
    Node
    Values []interface{} // 実際のデータ値
    Next   *LeafNode     // 次の葉ノードへのポインタ
    Prev   *LeafNode     // 前の葉ノードへのポインタ（双方向の場合）
}

// B+-tree構造体
type BPlusTree struct {
    Root      *Node
    MinDegree int        // 最小次数
    LeftMost  *LeafNode  // 最左の葉ノード（範囲検索用）
    RightMost *LeafNode  // 最右の葉ノード
}

// 新しいB+-treeを作成
func NewBPlusTree(minDegree int) *BPlusTree {
    return &BPlusTree{
        Root:      nil,
        MinDegree: minDegree,
        LeftMost:  nil,
        RightMost: nil,
    }
}

// 新しい内部ノードを作成
func (bt *BPlusTree) newInternalNode() *InternalNode {
    maxKeys := 2*bt.MinDegree - 1
    return &InternalNode{
        Node: Node{
            Keys:    make([]int, maxKeys),
            IsLeaf:  false,
            NumKeys: 0,
        },
        Children: make([]*Node, 2*bt.MinDegree),
    }
}

// 新しい葉ノードを作成
func (bt *BPlusTree) newLeafNode() *LeafNode {
    maxKeys := 2*bt.MinDegree - 1
    return &LeafNode{
        Node: Node{
            Keys:    make([]int, maxKeys),
            IsLeaf:  true,
            NumKeys: 0,
        },
        Values: make([]interface{}, maxKeys),
        Next:   nil,
        Prev:   nil,
    }
}

// キーを検索
func (bt *BPlusTree) Search(key int) (interface{}, bool) {
    if bt.Root == nil {
        return nil, false
    }

    leaf := bt.findLeafNode(key)
    if leaf == nil {
        return nil, false
    }

    // 葉ノード内でキーを検索
    for i := 0; i < leaf.NumKeys; i++ {
        if leaf.Keys[i] == key {
            return leaf.Values[i], true
        }
        if leaf.Keys[i] > key {
            break
        }
    }

    return nil, false
}

// 葉ノードを検索
func (bt *BPlusTree) findLeafNode(key int) *LeafNode {
    if bt.Root == nil {
        return nil
    }

    current := bt.Root

    // 内部ノードを辿って葉ノードに到達
    for !current.IsLeaf {
        internal := current.(*InternalNode)
        i := 0

        // 適切な子ノードを見つける
        for i < internal.NumKeys && key >= internal.Keys[i] {
            i++
        }

        current = internal.Children[i]
    }

    return current.(*LeafNode)
}

// キーと値を挿入
func (bt *BPlusTree) Insert(key int, value interface{}) error {
    if bt.Root == nil {
        // 空の木の場合、最初の葉ノードを作成
        leaf := bt.newLeafNode()
        leaf.Keys[0] = key
        leaf.Values[0] = value
        leaf.NumKeys = 1

        bt.Root = &leaf.Node
        bt.LeftMost = leaf
        bt.RightMost = leaf
        return nil
    }

    // ルートが満杯の場合、分割
    if bt.isNodeFull(bt.Root) {
        oldRoot := bt.Root
        newRoot := bt.newInternalNode()
        newRoot.Children[0] = oldRoot
        bt.Root = &newRoot.Node

        bt.splitChild(newRoot, 0)
    }

    return bt.insertNonFull(bt.Root, key, value)
}

// ノードが満杯かチェック
func (bt *BPlusTree) isNodeFull(node *Node) bool {
    return node.NumKeys == 2*bt.MinDegree-1
}

// 満杯でないノードに挿入
func (bt *BPlusTree) insertNonFull(node *Node, key int, value interface{}) error {
    if node.IsLeaf {
        // 葉ノードの場合
        leaf := node.(*LeafNode)
        return bt.insertIntoLeaf(leaf, key, value)
    } else {
        // 内部ノードの場合
        internal := node.(*InternalNode)
        i := internal.NumKeys - 1

        // 適切な子ノードを見つける
        for i >= 0 && internal.Keys[i] > key {
            i--
        }
        i++

        // 子ノードが満杯の場合、分割
        if bt.isNodeFull(internal.Children[i]) {
            bt.splitChild(internal, i)
            if internal.Keys[i] <= key {
                i++
            }
        }

        return bt.insertNonFull(internal.Children[i], key, value)
    }
}

// 葉ノードにキーと値を挿入
func (bt *BPlusTree) insertIntoLeaf(leaf *LeafNode, key int, value interface{}) error {
    i := leaf.NumKeys - 1

    // 適切な位置を見つけて挿入
    for i >= 0 && leaf.Keys[i] > key {
        leaf.Keys[i+1] = leaf.Keys[i]
        leaf.Values[i+1] = leaf.Values[i]
        i--
    }

    // 重複キーのチェック
    if i >= 0 && leaf.Keys[i] == key {
        // 重複キーの場合、値を更新
        leaf.Values[i] = value
        return nil
    }

    leaf.Keys[i+1] = key
    leaf.Values[i+1] = value
    leaf.NumKeys++

    return nil
}

// 子ノードを分割
func (bt *BPlusTree) splitChild(parent *InternalNode, index int) {
    fullNode := parent.Children[index]
    t := bt.MinDegree

    if fullNode.IsLeaf {
        // 葉ノードの分割
        fullLeaf := fullNode.(*LeafNode)
        newLeaf := bt.newLeafNode()

        // 右半分のデータを新しい葉ノードに移動
        newLeaf.NumKeys = t - 1
        for j := 0; j < t-1; j++ {
            newLeaf.Keys[j] = fullLeaf.Keys[j+t]
            newLeaf.Values[j] = fullLeaf.Values[j+t]
        }

        fullLeaf.NumKeys = t

        // 葉ノード間のリンクを更新
        newLeaf.Next = fullLeaf.Next
        newLeaf.Prev = fullLeaf
        if fullLeaf.Next != nil {
            fullLeaf.Next.Prev = newLeaf
        } else {
            bt.RightMost = newLeaf
        }
        fullLeaf.Next = newLeaf

        // 親ノードに新しい子ポインタを挿入
        bt.insertChildPointer(parent, index, &newLeaf.Node, newLeaf.Keys[0])

    } else {
        // 内部ノードの分割
        fullInternal := fullNode.(*InternalNode)
        newInternal := bt.newInternalNode()

        newInternal.NumKeys = t - 1

        // 右半分のキーを新しい内部ノードに移動
        for j := 0; j < t-1; j++ {
            newInternal.Keys[j] = fullInternal.Keys[j+t]
        }

        // 右半分の子ポインタを移動
        for j := 0; j < t; j++ {
            newInternal.Children[j] = fullInternal.Children[j+t]
        }

        fullInternal.NumKeys = t - 1

        // 中央キーを親に昇格
        promotedKey := fullInternal.Keys[t-1]

        // 親ノードに新しい子ポインタを挿入
        bt.insertChildPointer(parent, index, &newInternal.Node, promotedKey)
    }
}

// 親ノードに子ポインタとキーを挿入
func (bt *BPlusTree) insertChildPointer(parent *InternalNode, index int, newChild *Node, key int) {
    // 子ポインタをシフト
    for j := parent.NumKeys; j >= index+1; j-- {
        parent.Children[j+1] = parent.Children[j]
    }
    parent.Children[index+1] = newChild

    // キーをシフト
    for j := parent.NumKeys - 1; j >= index; j-- {
        parent.Keys[j+1] = parent.Keys[j]
    }
    parent.Keys[index] = key
    parent.NumKeys++
}

// キーを削除
func (bt *BPlusTree) Delete(key int) bool {
    if bt.Root == nil {
        return false
    }

    leaf := bt.findLeafNode(key)
    if leaf == nil {
        return false
    }

    // 葉ノード内でキーを検索して削除
    found := false
    for i := 0; i < leaf.NumKeys; i++ {
        if leaf.Keys[i] == key {
            // キーを削除
            for j := i; j < leaf.NumKeys-1; j++ {
                leaf.Keys[j] = leaf.Keys[j+1]
                leaf.Values[j] = leaf.Values[j+1]
            }
            leaf.NumKeys--
            found = true
            break
        }
    }

    if !found {
        return false
    }

    // ノードのアンダーフローをチェックして修正
    if leaf.NumKeys < bt.MinDegree-1 && bt.Root != &leaf.Node {
        bt.handleUnderflow(&leaf.Node)
    }

    // ルートが空になった場合の処理
    if bt.Root.NumKeys == 0 && !bt.Root.IsLeaf {
        internal := bt.Root.(*InternalNode)
        bt.Root = internal.Children[0]
    }

    return true
}

// アンダーフローの処理
func (bt *BPlusTree) handleUnderflow(node *Node) {
    // 実装は複雑になるため、基本的なケースのみ示す
    // 実際の実装では、兄弟ノードからの借用やマージが必要

    if node.IsLeaf {
        leaf := node.(*LeafNode)

        // 左の兄弟から借用を試行
        if leaf.Prev != nil && leaf.Prev.NumKeys > bt.MinDegree-1 {
            bt.borrowFromLeftSibling(leaf)
        } else if leaf.Next != nil && leaf.Next.NumKeys > bt.MinDegree-1 {
            // 右の兄弟から借用
            bt.borrowFromRightSibling(leaf)
        } else {
            // マージが必要
            bt.mergeLeafNodes(leaf)
        }
    }
}

// 左の兄弟から借用
func (bt *BPlusTree) borrowFromLeftSibling(leaf *LeafNode) {
    leftSibling := leaf.Prev

    // 現在のノードのデータを右にシフト
    for i := leaf.NumKeys; i > 0; i-- {
        leaf.Keys[i] = leaf.Keys[i-1]
        leaf.Values[i] = leaf.Values[i-1]
    }

    // 左の兄弟から最大要素を移動
    leaf.Keys[0] = leftSibling.Keys[leftSibling.NumKeys-1]
    leaf.Values[0] = leftSibling.Values[leftSibling.NumKeys-1]
    leaf.NumKeys++
    leftSibling.NumKeys--
}

// 右の兄弟から借用
func (bt *BPlusTree) borrowFromRightSibling(leaf *LeafNode) {
    rightSibling := leaf.Next

    // 右の兄弟から最小要素を移動
    leaf.Keys[leaf.NumKeys] = rightSibling.Keys[0]
    leaf.Values[leaf.NumKeys] = rightSibling.Values[0]
    leaf.NumKeys++

    // 右の兄弟のデータを左にシフト
    for i := 0; i < rightSibling.NumKeys-1; i++ {
        rightSibling.Keys[i] = rightSibling.Keys[i+1]
        rightSibling.Values[i] = rightSibling.Values[i+1]
    }
    rightSibling.NumKeys--
}

// 葉ノードをマージ
func (bt *BPlusTree) mergeLeafNodes(leaf *LeafNode) {
    var target *LeafNode

    if leaf.Prev != nil {
        target = leaf.Prev
        // leafの内容をtargetに移動
        for i := 0; i < leaf.NumKeys; i++ {
            target.Keys[target.NumKeys] = leaf.Keys[i]
            target.Values[target.NumKeys] = leaf.Values[i]
            target.NumKeys++
        }

        // リンクを更新
        target.Next = leaf.Next
        if leaf.Next != nil {
            leaf.Next.Prev = target
        } else {
            bt.RightMost = target
        }
    } else if leaf.Next != nil {
        target = leaf.Next
        // leafの内容をtargetの前に挿入
        // 実装は複雑になるため省略
    }
}

// 範囲検索
func (bt *BPlusTree) RangeQuery(minKey, maxKey int) []interface{} {
    var result []interface{}

    if bt.Root == nil {
        return result
    }

    // 開始点を検索
    startLeaf := bt.findLeafNode(minKey)
    if startLeaf == nil {
        return result
    }

    // 開始位置を見つける
    startPos := 0
    for startPos < startLeaf.NumKeys && startLeaf.Keys[startPos] < minKey {
        startPos++
    }

    // 葉ノードを線形に辿って範囲内のデータを収集
    current := startLeaf
    pos := startPos

    for current != nil {
        for pos < current.NumKeys {
            if current.Keys[pos] > maxKey {
                return result
            }
            if current.Keys[pos] >= minKey {
                result = append(result, current.Values[pos])
            }
            pos++
        }
        current = current.Next
        pos = 0
    }

    return result
}

// 最小値を取得
func (bt *BPlusTree) FindMin() (int, interface{}, bool) {
    if bt.LeftMost == nil || bt.LeftMost.NumKeys == 0 {
        return 0, nil, false
    }

    return bt.LeftMost.Keys[0], bt.LeftMost.Values[0], true
}

// 最大値を取得
func (bt *BPlusTree) FindMax() (int, interface{}, bool) {
    if bt.RightMost == nil || bt.RightMost.NumKeys == 0 {
        return 0, nil, false
    }

    lastIndex := bt.RightMost.NumKeys - 1
    return bt.RightMost.Keys[lastIndex], bt.RightMost.Values[lastIndex], true
}

// 順次スキャン（全要素）
func (bt *BPlusTree) SequentialScan() []interface{} {
    var result []interface{}

    current := bt.LeftMost
    for current != nil {
        for i := 0; i < current.NumKeys; i++ {
            result = append(result, current.Values[i])
        }
        current = current.Next
    }

    return result
}

// 前方一致検索
func (bt *BPlusTree) PrefixSearch(prefix string) []interface{} {
    var result []interface{}

    // 文字列キーの場合の実装例
    // 実際の実装では、keyの型を汎用化する必要がある

    current := bt.LeftMost
    for current != nil {
        for i := 0; i < current.NumKeys; i++ {
            // keyを文字列として扱う場合の例
            keyStr := fmt.Sprintf("%d", current.Keys[i])
            if len(keyStr) >= len(prefix) && keyStr[:len(prefix)] == prefix {
                result = append(result, current.Values[i])
            }
        }
        current = current.Next
    }

    return result
}

// 統計情報構造体
type Statistics struct {
    Height           int
    InternalNodes    int
    LeafNodes        int
    TotalKeys        int
    AvgKeysPerNode   float64
    UtilizationRate  float64
    MinDegree        int
}

// 統計情報を取得
func (bt *BPlusTree) GetStatistics() Statistics {
    stats := Statistics{
        MinDegree: bt.MinDegree,
    }

    if bt.Root != nil {
        stats.Height = bt.calculateHeight(bt.Root)
        bt.collectStatistics(bt.Root, &stats)

        totalNodes := stats.InternalNodes + stats.LeafNodes
        if totalNodes > 0 {
            stats.AvgKeysPerNode = float64(stats.TotalKeys) / float64(totalNodes)
            maxPossibleKeys := totalNodes * (2*bt.MinDegree - 1)
            stats.UtilizationRate = float64(stats.TotalKeys) / float64(maxPossibleKeys)
        }
    }

    return stats
}

// 高さを計算
func (bt *BPlusTree) calculateHeight(node *Node) int {
    if node.IsLeaf {
        return 1
    }

    internal := node.(*InternalNode)
    return 1 + bt.calculateHeight(internal.Children[0])
}

// 統計情報を収集
func (bt *BPlusTree) collectStatistics(node *Node, stats *Statistics) {
    stats.TotalKeys += node.NumKeys

    if node.IsLeaf {
        stats.LeafNodes++
    } else {
        stats.InternalNodes++
        internal := node.(*InternalNode)
        for i := 0; i <= internal.NumKeys; i++ {
            if internal.Children[i] != nil {
                bt.collectStatistics(internal.Children[i], stats)
            }
        }
    }
}

// 一括挿入（ソート済みデータ）
func (bt *BPlusTree) BulkLoad(sortedData []struct {
    Key   int
    Value interface{}
}) error {
    if len(sortedData) == 0 {
        return nil
    }

    // 効率的な一括構築アルゴリズム
    // ここでは簡略化して順次挿入
    for _, item := range sortedData {
        if err := bt.Insert(item.Key, item.Value); err != nil {
            return err
        }
    }

    return nil
}

// 木の妥当性を検証
func (bt *BPlusTree) Validate() error {
    if bt.Root == nil {
        return nil
    }

    // B+-tree特有の検証
    if err := bt.validateBPlusTreeProperties(bt.Root); err != nil {
        return err
    }

    // 葉ノード連結の検証
    if err := bt.validateLeafNodeLinks(); err != nil {
        return err
    }

    return nil
}

// B+-tree特有のプロパティを検証
func (bt *BPlusTree) validateBPlusTreeProperties(node *Node) error {
    if node.IsLeaf {
        // 葉ノードの検証
        leaf := node.(*LeafNode)

        // キーの順序性
        for i := 1; i < leaf.NumKeys; i++ {
            if leaf.Keys[i-1] >= leaf.Keys[i] {
                return fmt.Errorf("leaf node keys not in ascending order")
            }
        }

        // 最小キー数の検証（ルート以外）
        if node != bt.Root && leaf.NumKeys < bt.MinDegree-1 {
            return fmt.Errorf("leaf node has too few keys")
        }

    } else {
        // 内部ノードの検証
        internal := node.(*InternalNode)

        // 子ノードの検証
        for i := 0; i <= internal.NumKeys; i++ {
            if internal.Children[i] == nil {
                return fmt.Errorf("internal node has nil child")
            }
            if err := bt.validateBPlusTreeProperties(internal.Children[i]); err != nil {
                return err
            }
        }

        // 内部ノードはデータを持たないことの確認
        // （この実装では構造的に保証されている）
    }

    return nil
}

// 葉ノード連結の検証
func (bt *BPlusTree) validateLeafNodeLinks() error {
    if bt.LeftMost == nil {
        return nil
    }

    current := bt.LeftMost
    prev := (*LeafNode)(nil)

    for current != nil {
        // 双方向リンクの整合性
        if current.Prev != prev {
            return fmt.Errorf("leaf node backward link inconsistency")
        }

        // キーの順序性（ノード間）
        if prev != nil && prev.NumKeys > 0 && current.NumKeys > 0 {
            if prev.Keys[prev.NumKeys-1] >= current.Keys[0] {
                return fmt.Errorf("leaf nodes not in ascending order")
            }
        }

        prev = current
        current = current.Next
    }

    // 最右ノードの確認
    if bt.RightMost != prev {
        return fmt.Errorf("rightmost leaf node pointer inconsistency")
    }

    return nil
}

// 文字列表現
func (bt *BPlusTree) String() string {
    if bt.Root == nil {
        return "Empty B+-tree"
    }
    return bt.nodeString(bt.Root, 0)
}

func (bt *BPlusTree) nodeString(node *Node, depth int) string {
    indent := ""
    for i := 0; i < depth; i++ {
        indent += "  "
    }

    if node.IsLeaf {
        leaf := node.(*LeafNode)
        result := fmt.Sprintf("%sLeaf Keys: %v\n", indent, leaf.Keys[:leaf.NumKeys])
        result += fmt.Sprintf("%sLeaf Values: %v\n", indent, leaf.Values[:leaf.NumKeys])
        return result
    } else {
        internal := node.(*InternalNode)
        result := fmt.Sprintf("%sInternal Keys: %v\n", indent, internal.Keys[:internal.NumKeys])

        for i := 0; i <= internal.NumKeys; i++ {
            if internal.Children[i] != nil {
                result += bt.nodeString(internal.Children[i], depth+1)
            }
        }
        return result
    }
}
```

# B+-tree の応用例

## 1. データベース管理システム

- **MySQL の InnoDB**: プライマリインデックスとしてクラスタ化 B+-tree
- **PostgreSQL**: B+-tree インデックス（btree インデックス）
- **Oracle Database**: B+-tree 構造によるインデックス
- **SQLServer**: クラスタ化・非クラスタ化インデックス

## 2. ファイルシステム

- **NTFS**: ディレクトリ構造の管理
- **XFS**: ファイルシステムメタデータ
- **ZFS**: データセットの階層管理
- **Btrfs**: B+-tree ライクな構造

## 3. NoSQL データベース

- **MongoDB**: インデックス構造
- **CouchDB**: ビューインデックス
- **LevelDB**: SSTable のインデックス部分
- **RocksDB**: MemTable と SSTable の管理

## 4. キーバリューストア

- **BerkeleyDB**: B+-tree ストレージエンジン
- **LMDB**: Lightning Memory-Mapped Database
- **WiredTiger**: MongoDB のストレージエンジン
- **InnoDB**: MySQL のストレージエンジン

## 5. 分散システム

- **Google Bigtable**: タブレット内のデータ管理
- **Apache HBase**: HFile 内のデータ構造
- **Amazon DynamoDB**: 内部インデックス構造
- **Apache Cassandra**: SSTable のインデックス

## 6. インメモリデータベース

- **SAP HANA**: インメモリカラムストア
- **Redis**: ソート済みセット（Skip List との組み合わせ）
- **VoltDB**: インメモリ B+-tree
- **MemSQL**: インメモリ行ストア

# B+-tree と他のデータ構造との比較

## B-tree vs B+-tree 比較

| 特徴                       | B-tree                   | B+-tree                  | 優劣                     |
| -------------------------- | ------------------------ | ------------------------ | ------------------------ |
| **データ格納**             | 全ノードにデータ         | 葉ノードのみ             | B+-tree 有利（範囲検索） |
| **範囲検索**               | ツリー走査が必要         | 葉ノード連結で効率的     | B+-tree 有利             |
| **シーケンシャルアクセス** | O(n log n)               | O(n)                     | B+-tree 有利             |
| **単一キー検索**           | やや高速                 | やや低速                 | B-tree 有利              |
| **内部ノード効率**         | データでファンアウト減少 | キーのみでファンアウト大 | B+-tree 有利             |
| **メモリ使用量**           | やや少ない               | やや多い（重複キー）     | B-tree 有利              |
| **ディスク I/O**           | やや多い                 | 少ない（範囲検索時）     | B+-tree 有利             |
| **実装複雑度**             | 中程度                   | 高い                     | B-tree 有利              |

## 他のデータ構造との比較

| 特徴                  | B+-tree         | Hash Table          | AVL Tree  | Skip List | LSM Tree            |
| --------------------- | --------------- | ------------------- | --------- | --------- | ------------------- |
| **検索性能**          | O(log n)        | O(1)平均            | O(log n)  | O(log n)  | O(log n)            |
| **挿入性能**          | O(log n)        | O(1)平均            | O(log n)  | O(log n)  | O(1)償却            |
| **削除性能**          | O(log n)        | O(1)平均            | O(log n)  | O(log n)  | O(1)償却            |
| **範囲検索**          | ✅ 非常に効率的 | ❌ 不可             | ✅ 効率的 | ✅ 効率的 | ⚠️ やや複雑         |
| **順序保持**          | ✅ あり         | ❌ なし             | ✅ あり   | ✅ あり   | ⚠️ 部分的           |
| **ディスク I/O 効率** | ✅ 最適         | ❌ ランダムアクセス | ❌ 非効率 | ❌ 非効率 | ✅ 最適（書き込み） |
| **メモリ効率**        | ✅ 高い         | ⚠️ 負荷率依存       | ⚠️ 中程度 | ⚠️ 中程度 | ✅ 高い             |
| **実装複雑度**        | ❌ 高い         | ✅ 低い             | ⚠️ 中程度 | ⚠️ 中程度 | ❌ 高い             |
| **大量データ**        | ✅ 最適         | ⚠️ メモリ制約       | ❌ 非効率 | ❌ 非効率 | ✅ 最適             |

## 用途別最適選択

| 用途                         | 第 1 選択  | 第 2 選択 | 第 3 選択 | 避けるべき |
| ---------------------------- | ---------- | --------- | --------- | ---------- |
| **データベースインデックス** | B+-tree    | B-tree    | LSM Tree  | Hash Table |
| **範囲クエリ中心**           | B+-tree    | Skip List | AVL Tree  | Hash Table |
| **大量データ検索**           | B+-tree    | LSM Tree  | B-tree    | AVL Tree   |
| **単純なキー検索**           | Hash Table | B+-tree   | AVL Tree  | Skip List  |
| **順次スキャン**             | B+-tree    | Skip List | -         | Hash Table |
| **書き込み重視**             | LSM Tree   | B+-tree   | -         | B-tree     |
| **読み取り重視**             | B+-tree    | B-tree    | AVL Tree  | LSM Tree   |
| **メモリ制約**               | B+-tree    | AVL Tree  | -         | Hash Table |

# パフォーマンス特性

## 操作別パフォーマンス

| データサイズ | 検索時間 | 挿入時間 | 範囲検索（1000 件） | 順次スキャン |
| ------------ | -------- | -------- | ------------------- | ------------ |
| **1K 要素**  | ~1μs     | ~3μs     | ~50μs               | ~100μs       |
| **1M 要素**  | ~5μs     | ~15μs    | ~200μs              | ~10ms        |
| **1B 要素**  | ~20μs    | ~100μs   | ~500μs              | ~10s         |

## 木の高さとパフォーマンス

```
次数が64のB+-tree（一般的なディスクページサイズ4KB想定）:

要素数        木の高さ    検索時のディスクアクセス
1,000         2          2回
1,000,000     3          3回
1,000,000,000 4          4回

利点: データ量が増えても木の高さはゆっくりとしか増加しない
```

# まとめ

B+-tree（B+木）は、B-tree の発展形として開発され、データベースやファイルシステムにおける大量データの効率的管理に特化したデータ構造です。すべてのデータを葉ノードに格納し、葉ノード間の連結により範囲検索とシーケンシャルアクセスを極めて効率的に実現します。

Go での実装では、内部ノードと葉ノードの異なる構造を適切に管理し、葉ノード間の連結を維持することが重要です。特に分割・マージ操作時の連結リスト更新と、B+-tree 特有の重複キー管理に注意が必要です。

B+-tree は、範囲検索が頻繁で大量データを扱うシステム、特にデータベースやファイルシステムにおいて、その真価を発揮するデータ構造です。適切な実装により、安定した O(log n)性能と効率的な範囲処理を両立できます。
