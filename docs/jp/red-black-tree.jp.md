# Red-Black Tree（赤黒木）とは

Red-Black Tree（赤黒木）は、**自己平衡二分探索木**の一種で、各ノードに色（赤または黒）の属性を持たせることで木の平衡を維持する高度なデータ構造です。AVL 木よりも挿入・削除時の回転操作が少なく、実用的な性能を重視した平衡木として設計されています。

**このデータ構造の最大の利点は、最悪の場合でも O(log n)の検索・挿入・削除性能を保証しながら、AVL 木よりも挿入・削除のオーバーヘッドが少ない点です。**

# Red-Black Tree の構成要素

## 基本要素

- **ノード（Node）**: データと色情報を保持する基本単位
- **色（Color）**: 各ノードが持つ赤（Red）または黒（Black）の属性
- **ルート（Root）**: 常に黒色でなければならない最上位ノード
- **リーフ（Leaf/NIL）**: 黒色の仮想的な空ノード（NIL ノード）
- **親ノード（Parent）**: 上位階層への参照
- **子ノード（Children）**: 左の子と右の子への参照
- **黒色高さ（Black Height）**: ノードから任意のリーフまでの黒色ノード数

## Red-Black Tree の性質（5 つの基本規則）

1. **根性質**: ルートノードは黒色
2. **リーフ性質**: すべてのリーフ（NIL）ノードは黒色
3. **赤色性質**: 赤色ノードの子は必ず黒色（連続する赤色ノード禁止）
4. **黒色高さ性質**: 任意のノードから任意のリーフへの経路上の黒色ノード数は同じ
5. **二分探索木性質**: 左の子 < 親 < 右の子の順序関係を維持

# Red-Black Tree の種類

## 標準 Red-Black Tree

- 基本的な赤黒木実装
- 5 つの性質を厳密に維持
- 用途: 一般的な辞書・マップの実装

## Left-Leaning Red-Black Tree (LLRB)

- Robert Sedgewick 提案の簡略化版
- 赤色リンクを左に制限して実装を簡単化
- 用途: 教育目的、実装の簡単化

## AA Tree

- Arne Andersson 提案の簡略化版
- 赤色ノードを右の子のみに制限
- 用途: 実装の簡単化、理解しやすさ重視

## Persistent Red-Black Tree

- 関数型プログラミング向けの不変版
- 過去の状態を保持可能
- 用途: 関数型言語、アンドゥ機能

# Red-Black Tree の特徴

- **平衡保証**: 最悪の場合でも O(log n)の高さ保証
- **効率的な更新**: AVL 木より少ない回転でバランス維持
- **実用性**: 多くのライブラリで採用される実装のしやすさ
- **安定性**: 挿入・削除の性能が予測可能
- **メモリ効率**: 各ノードに 1 ビットの色情報のみ追加

# Red-Black Tree の構成要素の図解

## 基本的な Red-Black Tree の例

```
数値セット: {7, 3, 11, 1, 5, 9, 13, 6, 8, 12, 14}

Red-Black Tree構造:
                    7(B)
                  /      \
               3(R)      11(R)
              /    \     /     \
           1(B)   5(B) 9(B)   13(B)
             \     /    \      /  \
            NIL  6(R)  8(R)  12(R) 14(R)

(B) = 黒色ノード
(R) = 赤色ノード
NIL = 黒色の仮想リーフノード
```

### 5 つの性質の検証

1. **根性質**: ルート 7 は黒色 ✅
2. **リーフ性質**: すべての NIL は黒色 ✅
3. **赤色性質**: 赤色ノード（3,11,6,8,12,14）の子はすべて黒色 ✅
4. **黒色高さ性質**: どのパスも黒色ノード数は 3 個で一致 ✅
5. **二分探索木性質**: 左 < 親 < 右の関係を維持 ✅

## ノード構造の詳細

```go
type Color int

const (
    RED   Color = 0
    BLACK Color = 1
)

type RBNode struct {
    Key     int     // データ
    Color   Color   // 赤(0)または黒(1)
    Left    *RBNode // 左の子
    Right   *RBNode // 右の子
    Parent  *RBNode // 親ノード
}
```

## 回転操作の図解

### 左回転（Left Rotation）

```
回転前:                回転後:
    x                      y
   / \                    / \
  α   y         →        x   γ
     / \                / \
    β   γ              α   β

x.right = y            y.left = x
y.left = x             x.right = β
```

### 右回転（Right Rotation）

```
回転前:                回転後:
      y                    x
     / \                  / \
    x   γ       →        α   y
   / \                      / \
  α   β                    β   γ

y.left = x             x.right = y
x.right = y            y.left = β
```

## 挿入時の色変更・回転パターン

### ケース 1: 叔父ノードが赤色

```
変更前:           変更後:
    G(B)              G(R)
   /    \            /    \
  P(R)  U(R)   →   P(B)  U(B)
 /                /
N(R)            N(R)

G=祖父, P=親, U=叔父, N=新ノード
親と叔父を黒に、祖父を赤に変更
```

### ケース 2: 叔父ノードが黒色（直線パターン）

```
変更前:           右回転後:
  G(B)               P(B)
 /                  /    \
P(R)        →     N(R)  G(R)
/
N(R)

祖父で右回転し、元の親と祖父の色を交換
```

### ケース 3: 叔父ノードが黒色（ジグザグパターン）

```
変更前:        左回転後:      右回転後:
  G(B)           G(B)           N(B)
 /              /              /    \
P(R)     →    N(R)      →    P(R)  G(R)
  \          /
  N(R)     P(R)

親で左回転してケース2に変換後、祖父で右回転
```

## 削除時の修正パターン

### ケース 1: 兄弟が赤色

```
修正前:          左回転後:
   P(B)             S(B)
  /    \           /    \
N(B)  S(R)   →   P(R)  SR(B)
      /  \       /  \
    SL(B) SR(B) N(B) SL(B)

兄弟の色を黒、親の色を赤に変更し、親で左回転
```

### ケース 2: 兄弟と兄弟の子がすべて黒色

```
修正前:          修正後:
   P(?)             P(?)
  /    \           /    \
N(B)  S(B)   →   N(B)  S(R)
      /  \             /  \
    SL(B) SR(B)      SL(B) SR(B)

兄弟を赤に変更し、問題を親に移動
```

# Red-Black Tree で行える処理

## 基本操作

| 機能     | 説明             | 計算量   | 戻り値              | 得意/苦手 | 補足                      |
| -------- | ---------------- | -------- | ------------------- | --------- | ------------------------- |
| Insert   | ノードを挿入     | O(log n) | bool（成功/失敗）   | ✅        | 最大 3 回の回転で平衡維持 |
| Delete   | ノードを削除     | O(log n) | bool（成功/失敗）   | ✅        | 最大 3 回の回転で平衡維持 |
| Search   | ノードを検索     | O(log n) | ノードまたは nil    | ✅        | 二分探索木の基本操作      |
| FindMin  | 最小値を検索     | O(log n) | 最小値ノード        | ✅        | 左端ノードまで辿る        |
| FindMax  | 最大値を検索     | O(log n) | 最大値ノード        | ✅        | 右端ノードまで辿る        |
| Contains | 要素の存在確認   | O(log n) | bool（存在/非存在） | ✅        | Search と同等             |
| Size     | ノード数を取得   | O(n)     | int（ノード数）     | ❌        | 全ノードを数える必要      |
| Height   | 木の高さを取得   | O(n)     | int（高さ）         | ✅        | 平衡保証により予測可能    |
| IsEmpty  | 木が空かチェック | O(1)     | bool（空/非空）     | ✅        | ルートノードの確認        |
| Clear    | 全ノードを削除   | O(n)     | なし（void）        | ✅        | メモリ解放処理            |

## 走査・列挙操作

| 機能                | 説明             | 計算量       | 戻り値         | 得意/苦手 | 補足                       |
| ------------------- | ---------------- | ------------ | -------------- | --------- | -------------------------- |
| InOrderTraversal    | 中順走査（昇順） | O(n)         | ソート済み配列 | ✅        | 赤黒木の大きな利点         |
| PreOrderTraversal   | 前順走査         | O(n)         | 前順配列       | ✅        | 木の構造保存に有用         |
| PostOrderTraversal  | 後順走査         | O(n)         | 後順配列       | ✅        | 削除順序の決定             |
| LevelOrderTraversal | レベル順走査     | O(n)         | レベル順配列   | ✅        | 幅優先探索                 |
| RangeQuery          | 範囲検索         | O(log n + k) | 範囲内要素配列 | ✅        | k: 結果数、効率的          |
| Predecessor         | 前駆ノード検索   | O(log n)     | 前駆ノード     | ✅        | 中順走査での前要素         |
| Successor           | 後続ノード検索   | O(log n)     | 後続ノード     | ✅        | 中順走査での次要素         |
| KthElement          | k 番目の要素検索 | O(log n)     | k 番目の要素   | ⚠️        | サイズ情報の追加管理が必要 |

## 平衡維持・構造操作

| 機能        | 説明           | 計算量   | 戻り値              | 得意/苦手 | 補足               |
| ----------- | -------------- | -------- | ------------------- | --------- | ------------------ |
| RotateLeft  | 左回転         | O(1)     | 新しい部分木根      | ✅        | 平衡維持の基本操作 |
| RotateRight | 右回転         | O(1)     | 新しい部分木根      | ✅        | 平衡維持の基本操作 |
| FixInsert   | 挿入後の修正   | O(log n) | なし（void）        | ✅        | 最大 3 回の回転    |
| FixDelete   | 削除後の修正   | O(log n) | なし（void）        | ✅        | 最大 3 回の回転    |
| Validate    | 5 つの性質検証 | O(n)     | bool（妥当/不正）   | ✅        | デバッグ・テスト用 |
| BlackHeight | 黒色高さ計算   | O(log n) | int（黒色高さ）     | ✅        | 平衡度合いの確認   |
| IsBalanced  | 平衡状態確認   | O(n)     | bool（平衡/非平衡） | ✅        | 5 つの性質の確認   |
| GetColor    | ノードの色取得 | O(1)     | Color（赤/黒）      | ✅        | NIL ノードは黒色   |

## 集合演算・比較

| 機能         | 説明         | 計算量     | 戻り値              | 得意/苦手 | 補足                   |
| ------------ | ------------ | ---------- | ------------------- | --------- | ---------------------- |
| Union        | 和集合       | O(m + n)   | 新しい赤黒木        | ✅        | 2 つの木をマージ       |
| Intersection | 積集合       | O(m + n)   | 新しい赤黒木        | ✅        | 共通要素の抽出         |
| Difference   | 差集合       | O(m + n)   | 新しい赤黒木        | ✅        | 一方にのみ存在する要素 |
| IsSubset     | 部分集合判定 | O(m log n) | bool（部分集合/否） | ✅        | 全要素の存在確認       |
| IsEqual      | 等価性判定   | O(n)       | bool（等価/非等価） | ✅        | 構造とデータの比較     |
| Clone        | 木の複製     | O(n)       | 新しい赤黒木        | ✅        | 深いコピーの作成       |
| Merge        | 木のマージ   | O(m + n)   | マージ済み木        | ⚠️        | 重複要素の処理が複雑   |

## デバッグ・可視化

| 機能               | 説明              | 計算量 | 戻り値          | 得意/苦手 | 補足                     |
| ------------------ | ----------------- | ------ | --------------- | --------- | ------------------------ |
| Print              | 木構造の出力      | O(n)   | なし（出力）    | ✅        | デバッグに有用           |
| ToDotFormat        | Graphviz 形式変換 | O(n)   | DOT 文字列      | ✅        | 可視化に最適             |
| ToArray            | 配列形式変換      | O(n)   | 配列            | ✅        | 中順走査結果             |
| GetStatistics      | 統計情報取得      | O(n)   | 統計構造体      | ✅        | ノード数、高さ、黒色高さ |
| ValidateProperties | 性質の詳細検証    | O(n)   | 検証結果        | ✅        | 各性質の個別確認         |
| GetMemoryUsage     | メモリ使用量計算  | O(n)   | int（バイト数） | ✅        | メモリ効率の確認         |

## 苦手な処理・制限

| 機能            | 説明                 | 計算量     | 戻り値       | 得意/苦手 | 補足                     |
| --------------- | -------------------- | ---------- | ------------ | --------- | ------------------------ |
| RandomAccess    | インデックスアクセス | O(n)       | i 番目の要素 | ❌        | 順序統計木が必要         |
| BulkOperations  | 一括操作             | O(n log n) | 処理済み木   | ❌        | 個別操作の繰り返し       |
| RangeUpdate     | 範囲更新             | O(n)       | 更新済み木   | ❌        | セグメント木が適している |
| Persistence     | 永続化               | O(n)       | 永続化済み木 | ❌        | 特殊な実装が必要         |
| Parallelization | 並列処理             | -          | -            | ❌        | 木構造の依存関係         |
| Compression     | 圧縮                 | O(n)       | 圧縮済み木   | ❌        | 複雑な圧縮アルゴリズム   |

# Red-Black Tree の実装方法

## 一般的な実装方法

### NIL ノードによる実装

- 仮想的な黒色リーフノードを明示的に表現
- 全ての NULL ポインタを NIL ノードに置き換え
- 実装が明確だが、メモリ使用量が増加

### NULL ポインタベース実装

- NULL ポインタを黒色のリーフとして扱う
- メモリ効率が良いが、実装時の条件分岐が複雑
- 多くの実用実装で採用

### Left-Leaning 実装（LLRB）

- 赤色リンクを左のみに制限
- 実装が大幅に簡略化
- 性能は標準版とほぼ同等

## Go による基本的な Red-Black Tree 実装例

```go
package redblacktree

import (
    "fmt"
    "strings"
)

// 色の定義
type Color int

const (
    RED   Color = 0
    BLACK Color = 1
)

// ノード構造体
type RBNode struct {
    Key    int
    Color  Color
    Left   *RBNode
    Right  *RBNode
    Parent *RBNode
}

// Red-Black Tree構造体
type RedBlackTree struct {
    Root *RBNode
    size int
}

// 新しいRed-Black Treeを作成
func New() *RedBlackTree {
    return &RedBlackTree{
        Root: nil,
        size: 0,
    }
}

// 新しいノードを作成
func newNode(key int, color Color) *RBNode {
    return &RBNode{
        Key:    key,
        Color:  color,
        Left:   nil,
        Right:  nil,
        Parent: nil,
    }
}

// ノードの色を取得（NILノードは黒色）
func (rbt *RedBlackTree) getColor(node *RBNode) Color {
    if node == nil {
        return BLACK
    }
    return node.Color
}

// ノードの色を設定
func (rbt *RedBlackTree) setColor(node *RBNode, color Color) {
    if node != nil {
        node.Color = color
    }
}

// 左回転
func (rbt *RedBlackTree) rotateLeft(x *RBNode) {
    y := x.Right
    x.Right = y.Left

    if y.Left != nil {
        y.Left.Parent = x
    }

    y.Parent = x.Parent

    if x.Parent == nil {
        rbt.Root = y
    } else if x == x.Parent.Left {
        x.Parent.Left = y
    } else {
        x.Parent.Right = y
    }

    y.Left = x
    x.Parent = y
}

// 右回転
func (rbt *RedBlackTree) rotateRight(y *RBNode) {
    x := y.Left
    y.Left = x.Right

    if x.Right != nil {
        x.Right.Parent = y
    }

    x.Parent = y.Parent

    if y.Parent == nil {
        rbt.Root = x
    } else if y == y.Parent.Left {
        y.Parent.Left = x
    } else {
        y.Parent.Right = x
    }

    x.Right = y
    y.Parent = x
}

// 挿入
func (rbt *RedBlackTree) Insert(key int) bool {
    // 標準的な二分探索木の挿入
    newNode := newNode(key, RED)

    if rbt.Root == nil {
        rbt.Root = newNode
        rbt.Root.Color = BLACK
        rbt.size++
        return true
    }

    current := rbt.Root
    var parent *RBNode

    for current != nil {
        parent = current
        if key < current.Key {
            current = current.Left
        } else if key > current.Key {
            current = current.Right
        } else {
            return false // 重複は挿入しない
        }
    }

    newNode.Parent = parent
    if key < parent.Key {
        parent.Left = newNode
    } else {
        parent.Right = newNode
    }

    rbt.size++

    // Red-Black Tree特有の修正
    rbt.fixInsert(newNode)
    return true
}

// 挿入後の修正
func (rbt *RedBlackTree) fixInsert(node *RBNode) {
    for node != rbt.Root && rbt.getColor(node.Parent) == RED {
        if node.Parent == node.Parent.Parent.Left {
            // 父が祖父の左の子の場合
            uncle := node.Parent.Parent.Right

            if rbt.getColor(uncle) == RED {
                // ケース1: 叔父が赤色
                rbt.setColor(node.Parent, BLACK)
                rbt.setColor(uncle, BLACK)
                rbt.setColor(node.Parent.Parent, RED)
                node = node.Parent.Parent
            } else {
                if node == node.Parent.Right {
                    // ケース2: ノードが父の右の子（ジグザグ）
                    node = node.Parent
                    rbt.rotateLeft(node)
                }
                // ケース3: ノードが父の左の子（直線）
                rbt.setColor(node.Parent, BLACK)
                rbt.setColor(node.Parent.Parent, RED)
                rbt.rotateRight(node.Parent.Parent)
            }
        } else {
            // 父が祖父の右の子の場合（対称）
            uncle := node.Parent.Parent.Left

            if rbt.getColor(uncle) == RED {
                // ケース1: 叔父が赤色
                rbt.setColor(node.Parent, BLACK)
                rbt.setColor(uncle, BLACK)
                rbt.setColor(node.Parent.Parent, RED)
                node = node.Parent.Parent
            } else {
                if node == node.Parent.Left {
                    // ケース2: ノードが父の左の子（ジグザグ）
                    node = node.Parent
                    rbt.rotateRight(node)
                }
                // ケース3: ノードが父の右の子（直線）
                rbt.setColor(node.Parent, BLACK)
                rbt.setColor(node.Parent.Parent, RED)
                rbt.rotateLeft(node.Parent.Parent)
            }
        }
    }
    rbt.Root.Color = BLACK
}

// 検索
func (rbt *RedBlackTree) Search(key int) *RBNode {
    current := rbt.Root

    for current != nil {
        if key == current.Key {
            return current
        } else if key < current.Key {
            current = current.Left
        } else {
            current = current.Right
        }
    }

    return nil
}

// 最小値ノードを検索
func (rbt *RedBlackTree) findMin(node *RBNode) *RBNode {
    for node != nil && node.Left != nil {
        node = node.Left
    }
    return node
}

// 削除
func (rbt *RedBlackTree) Delete(key int) bool {
    nodeToDelete := rbt.Search(key)
    if nodeToDelete == nil {
        return false
    }

    rbt.size--

    var y *RBNode // 実際に削除されるノード
    var x *RBNode // 削除されるノードの子

    if nodeToDelete.Left == nil || nodeToDelete.Right == nil {
        y = nodeToDelete
    } else {
        y = rbt.findMin(nodeToDelete.Right) // 後続ノード
    }

    if y.Left != nil {
        x = y.Left
    } else {
        x = y.Right
    }

    var xParent *RBNode
    if x != nil {
        x.Parent = y.Parent
        xParent = x.Parent
    } else {
        xParent = y.Parent
    }

    if y.Parent == nil {
        rbt.Root = x
    } else if y == y.Parent.Left {
        y.Parent.Left = x
    } else {
        y.Parent.Right = x
    }

    if y != nodeToDelete {
        nodeToDelete.Key = y.Key
    }

    if y.Color == BLACK {
        rbt.fixDelete(x, xParent)
    }

    return true
}

// 削除後の修正
func (rbt *RedBlackTree) fixDelete(node *RBNode, parent *RBNode) {
    for node != rbt.Root && rbt.getColor(node) == BLACK {
        if node == parent.Left {
            sibling := parent.Right

            if rbt.getColor(sibling) == RED {
                // ケース1: 兄弟が赤色
                rbt.setColor(sibling, BLACK)
                rbt.setColor(parent, RED)
                rbt.rotateLeft(parent)
                sibling = parent.Right
            }

            if rbt.getColor(sibling.Left) == BLACK &&
               rbt.getColor(sibling.Right) == BLACK {
                // ケース2: 兄弟と兄弟の子がすべて黒色
                rbt.setColor(sibling, RED)
                node = parent
                parent = node.Parent
            } else {
                if rbt.getColor(sibling.Right) == BLACK {
                    // ケース3: 兄弟の右の子が黒色
                    rbt.setColor(sibling.Left, BLACK)
                    rbt.setColor(sibling, RED)
                    rbt.rotateRight(sibling)
                    sibling = parent.Right
                }

                // ケース4: 兄弟の右の子が赤色
                rbt.setColor(sibling, rbt.getColor(parent))
                rbt.setColor(parent, BLACK)
                rbt.setColor(sibling.Right, BLACK)
                rbt.rotateLeft(parent)
                node = rbt.Root
            }
        } else {
            // 対称ケース（nodeが右の子の場合）
            sibling := parent.Left

            if rbt.getColor(sibling) == RED {
                rbt.setColor(sibling, BLACK)
                rbt.setColor(parent, RED)
                rbt.rotateRight(parent)
                sibling = parent.Left
            }

            if rbt.getColor(sibling.Right) == BLACK &&
               rbt.getColor(sibling.Left) == BLACK {
                rbt.setColor(sibling, RED)
                node = parent
                parent = node.Parent
            } else {
                if rbt.getColor(sibling.Left) == BLACK {
                    rbt.setColor(sibling.Right, BLACK)
                    rbt.setColor(sibling, RED)
                    rbt.rotateLeft(sibling)
                    sibling = parent.Left
                }

                rbt.setColor(sibling, rbt.getColor(parent))
                rbt.setColor(parent, BLACK)
                rbt.setColor(sibling.Left, BLACK)
                rbt.rotateRight(parent)
                node = rbt.Root
            }
        }
    }

    rbt.setColor(node, BLACK)
}

// 中順走査
func (rbt *RedBlackTree) InOrderTraversal() []int {
    var result []int
    rbt.inOrderHelper(rbt.Root, &result)
    return result
}

func (rbt *RedBlackTree) inOrderHelper(node *RBNode, result *[]int) {
    if node != nil {
        rbt.inOrderHelper(node.Left, result)
        *result = append(*result, node.Key)
        rbt.inOrderHelper(node.Right, result)
    }
}

// サイズを取得
func (rbt *RedBlackTree) Size() int {
    return rbt.size
}

// 空かどうかチェック
func (rbt *RedBlackTree) IsEmpty() bool {
    return rbt.size == 0
}

// 高さを取得
func (rbt *RedBlackTree) Height() int {
    return rbt.getHeight(rbt.Root)
}

func (rbt *RedBlackTree) getHeight(node *RBNode) int {
    if node == nil {
        return 0
    }

    leftHeight := rbt.getHeight(node.Left)
    rightHeight := rbt.getHeight(node.Right)

    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
}

// 黒色高さを取得
func (rbt *RedBlackTree) BlackHeight() int {
    return rbt.getBlackHeight(rbt.Root)
}

func (rbt *RedBlackTree) getBlackHeight(node *RBNode) int {
    if node == nil {
        return 1 // NILノードは黒色
    }

    leftBlackHeight := rbt.getBlackHeight(node.Left)
    blackIncrement := 0
    if node.Color == BLACK {
        blackIncrement = 1
    }

    return leftBlackHeight + blackIncrement
}

// Red-Black Treeの性質を検証
func (rbt *RedBlackTree) Validate() bool {
    if rbt.Root == nil {
        return true
    }

    // 性質1: ルートは黒色
    if rbt.Root.Color != BLACK {
        return false
    }

    // 性質3と4: 赤色ノードの子が黒色 & 黒色高さの一致
    _, valid := rbt.validateHelper(rbt.Root)
    return valid
}

func (rbt *RedBlackTree) validateHelper(node *RBNode) (int, bool) {
    if node == nil {
        return 1, true // NILノードは黒色
    }

    // 性質3: 赤色ノードの子は黒色
    if node.Color == RED {
        if (node.Left != nil && node.Left.Color == RED) ||
           (node.Right != nil && node.Right.Color == RED) {
            return 0, false
        }
    }

    leftBlackHeight, leftValid := rbt.validateHelper(node.Left)
    rightBlackHeight, rightValid := rbt.validateHelper(node.Right)

    if !leftValid || !rightValid {
        return 0, false
    }

    // 性質4: 黒色高さの一致
    if leftBlackHeight != rightBlackHeight {
        return 0, false
    }

    blackIncrement := 0
    if node.Color == BLACK {
        blackIncrement = 1
    }

    return leftBlackHeight + blackIncrement, true
}

// 木構造を文字列で表現
func (rbt *RedBlackTree) String() string {
    if rbt.Root == nil {
        return "Empty Tree"
    }

    var sb strings.Builder
    rbt.printHelper(rbt.Root, "", true, &sb)
    return sb.String()
}

func (rbt *RedBlackTree) printHelper(node *RBNode, prefix string, isLast bool, sb *strings.Builder) {
    if node != nil {
        sb.WriteString(prefix)

        if isLast {
            sb.WriteString("└── ")
            prefix += "    "
        } else {
            sb.WriteString("├── ")
            prefix += "│   "
        }

        colorStr := "B"
        if node.Color == RED {
            colorStr = "R"
        }

        sb.WriteString(fmt.Sprintf("%d(%s)\n", node.Key, colorStr))

        // 子ノードを出力
        if node.Left != nil || node.Right != nil {
            if node.Right != nil {
                rbt.printHelper(node.Right, prefix, node.Left == nil, sb)
            }
            if node.Left != nil {
                rbt.printHelper(node.Left, prefix, true, sb)
            }
        }
    }
}
```

## Left-Leaning Red-Black Tree (LLRB) の実装例

```go
package llrb

// LLRB特有のより簡単な実装
type LLRBNode struct {
    Key   int
    Value interface{}
    Left  *LLRBNode
    Right *LLRBNode
    Red   bool // 赤色かどうか
}

type LLRB struct {
    root *LLRBNode
    size int
}

// 新しいLLRBを作成
func NewLLRB() *LLRB {
    return &LLRB{}
}

// ノードが赤色かチェック
func isRed(node *LLRBNode) bool {
    if node == nil {
        return false
    }
    return node.Red
}

// 色を反転
func (node *LLRBNode) flipColors() {
    node.Red = !node.Red
    if node.Left != nil {
        node.Left.Red = !node.Left.Red
    }
    if node.Right != nil {
        node.Right.Red = !node.Right.Red
    }
}

// 左回転
func rotateLeft(h *LLRBNode) *LLRBNode {
    x := h.Right
    h.Right = x.Left
    x.Left = h
    x.Red = h.Red
    h.Red = true
    return x
}

// 右回転
func rotateRight(h *LLRBNode) *LLRBNode {
    x := h.Left
    h.Left = x.Right
    x.Right = h
    x.Red = h.Red
    h.Red = true
    return x
}

// 挿入（LLRB簡略版）
func (llrb *LLRB) Insert(key int, value interface{}) {
    llrb.root = llrb.insert(llrb.root, key, value)
    llrb.root.Red = false // ルートは常に黒色
}

func (llrb *LLRB) insert(h *LLRBNode, key int, value interface{}) *LLRBNode {
    if h == nil {
        llrb.size++
        return &LLRBNode{Key: key, Value: value, Red: true}
    }

    if key < h.Key {
        h.Left = llrb.insert(h.Left, key, value)
    } else if key > h.Key {
        h.Right = llrb.insert(h.Right, key, value)
    } else {
        h.Value = value // 更新
        return h
    }

    // LLRB修正: 右が赤で左が黒なら左回転
    if isRed(h.Right) && !isRed(h.Left) {
        h = rotateLeft(h)
    }

    // 連続する左の赤リンクなら右回転
    if isRed(h.Left) && isRed(h.Left.Left) {
        h = rotateRight(h)
    }

    // 両方の子が赤なら色を反転
    if isRed(h.Left) && isRed(h.Right) {
        h.flipColors()
    }

    return h
}

// 検索
func (llrb *LLRB) Search(key int) (interface{}, bool) {
    node := llrb.search(llrb.root, key)
    if node == nil {
        return nil, false
    }
    return node.Value, true
}

func (llrb *LLRB) search(node *LLRBNode, key int) *LLRBNode {
    for node != nil {
        if key == node.Key {
            return node
        } else if key < node.Key {
            node = node.Left
        } else {
            node = node.Right
        }
    }
    return nil
}
```

# Red-Black Tree の応用例

## 1. 標準ライブラリの実装

- **C++ STL map/set**: std::map、std::set の内部実装
- **Java TreeMap/TreeSet**: Java 標準ライブラリの連想配列
- **.NET SortedDictionary**: .NET Framework の順序付き辞書
- **Linux CFS**: 完全公平スケジューラのプロセス管理

## 2. データベース・インデックス

- **メモリ内インデックス**: リレーショナル DB の主キーインデックス
- **範囲クエリ**: Between 句による効率的な範囲検索
- **ソート処理**: Order By 句の効率的な実装
- **重複除去**: Distinct 処理での重複排除

## 3. リアルタイムシステム

- **プロセススケジューラ**: OS のタスク優先度管理
- **イベント管理**: 時刻順イベントキューの実装
- **リソース管理**: 限定リソースの公平な配分
- **ロードバランサ**: サーバー負荷の動的分散

## 4. ゲーム開発

- **ランキングシステム**: プレイヤースコアの順位管理
- **ゲームオブジェクト管理**: 座標によるオブジェクトソート
- **AI 決定木**: ゲーム AI の状態空間探索
- **衝突検出**: 空間分割での効率的な衝突判定

## 5. 金融・取引システム

- **オーダーブック**: 株式取引の買い注文・売り注文管理
- **価格優先順位**: 価格・時刻優先でのマッチング
- **リスク管理**: ポートフォリオの動的リバランス
- **高頻度取引**: ミリ秒単位での取引処理

## 6. ネットワーク・通信

- **ルーティングテーブル**: IP アドレスの効率的ルーティング
- **QoS 管理**: 帯域幅の優先度管理
- **セッション管理**: 接続セッションの順序管理
- **パケットスケジューリング**: ネットワークパケットの公平な配信

## 7. Web・検索エンジン

- **検索インデックス**: キーワードによる高速検索
- **ページランク**: ウェブページの重要度順位
- **キャッシュ管理**: アクセス頻度による動的キャッシュ
- **負荷分散**: サーバー負荷の動的調整

## 8. メモリ管理・OS

- **仮想メモリ**: ページテーブルの効率的管理
- **ファイルシステム**: inode 管理とディレクトリ索引
- **プロセス管理**: プロセス ID による効率的管理
- **デバイス管理**: デバイスドライバの優先度制御

# Red-Black Tree と他のデータ構造との比較

## 平衡木比較

| 特徴                 | Red-Black Tree   | AVL Tree       | Splay Tree   | B-Tree       | Skip List       |
| -------------------- | ---------------- | -------------- | ------------ | ------------ | --------------- |
| **平衡条件**         | 色による緩い平衡 | 厳密な高さ平衡 | アクセス適応 | ノード分割   | 確率的平衡      |
| **最大高さ**         | 2 log(n+1)       | 1.44 log(n+2)  | O(n)最悪     | O(log n)     | O(log n)高確率  |
| **検索時間**         | O(log n)         | O(log n)       | O(log n)償却 | O(log n)     | O(log n)期待    |
| **挿入時間**         | O(log n)         | O(log n)       | O(log n)償却 | O(log n)     | O(log n)期待    |
| **削除時間**         | O(log n)         | O(log n)       | O(log n)償却 | O(log n)     | O(log n)期待    |
| **回転回数（挿入）** | 最大 3 回        | 最大 2 回      | 任意回数     | 分割操作     | なし            |
| **回転回数（削除）** | 最大 3 回        | 最大 log n 回  | 任意回数     | 分割・結合   | なし            |
| **メモリ使用量**     | ノード+1bit      | ノード+高さ    | ノード       | 大きなノード | ノード+ポインタ |
| **実装複雑度**       | 中程度           | 中程度         | 簡単         | 複雑         | 簡単            |
| **最悪性能保証**     | ✅ 保証          | ✅ 保証        | ❌ O(n)      | ✅ 保証      | ❌ 確率的       |
| **実用性**           | ✅ 非常に高い    | ⚠️ 削除重い    | ⚠️ 予測困難  | ✅ DB 向け   | ⚠️ メモリ多用   |

## 用途別最適選択

| 用途                     | 第 1 選択      | 第 2 選択         | 第 3 選択      | 避けるべき     |
| ------------------------ | -------------- | ----------------- | -------------- | -------------- |
| **汎用的な辞書・マップ** | Red-Black Tree | AVL Tree          | Skip List      | Splay Tree     |
| **読み取り重視**         | AVL Tree       | Red-Black Tree    | B-Tree         | Splay Tree     |
| **更新重視**             | Red-Black Tree | Splay Tree        | Skip List      | AVL Tree       |
| **局所性のあるアクセス** | Splay Tree     | Red-Black Tree    | AVL Tree       | B-Tree         |
| **大量データ・外部記憶** | B-Tree         | B+ Tree           | -              | Red-Black Tree |
| **並行処理**             | Skip List      | Lock-free RB Tree | -              | Splay Tree     |
| **組み込みシステム**     | Red-Black Tree | AVL Tree          | -              | Skip List      |
| **標準ライブラリ**       | Red-Black Tree | -                 | -              | Splay Tree     |
| **データベース**         | B-Tree         | B+ Tree           | Red-Black Tree | Splay Tree     |
| **リアルタイム処理**     | Red-Black Tree | AVL Tree          | -              | Splay Tree     |

## パフォーマンス特性比較

| 操作パターン       | Red-Black Tree  | AVL Tree             | Splay Tree      | ハッシュテーブル | 配列（ソート済み） |
| ------------------ | --------------- | -------------------- | --------------- | ---------------- | ------------------ |
| **ランダム検索**   | ✅ O(log n)     | ✅ O(log n)          | ⚠️ O(log n)償却 | ✅ O(1)平均      | ✅ O(log n)        |
| **順次検索**       | ✅ O(log n)     | ✅ O(log n)          | ✅ O(1)償却     | ❌ O(1)順序なし  | ✅ O(1)            |
| **範囲検索**       | ✅ O(log n + k) | ✅ O(log n + k)      | ⚠️ 再構築必要   | ❌ 不適          | ✅ O(log n + k)    |
| **挿入**           | ✅ O(log n)     | ⚠️ O(log n)+重い回転 | ✅ O(log n)償却 | ✅ O(1)平均      | ❌ O(n)            |
| **削除**           | ✅ O(log n)     | ❌ O(log n)+重い回転 | ✅ O(log n)償却 | ✅ O(1)平均      | ❌ O(n)            |
| **メモリ効率**     | ✅ 良好         | ⚠️ 高さ情報          | ✅ 良好         | ⚠️ 負荷率依存    | ✅ 最高            |
| **キャッシュ効率** | ⚠️ ポインタ追跡 | ⚠️ ポインタ追跡      | ⚠️ ポインタ追跡 | ⚠️ ハッシュ関数  | ✅ 局所性良好      |
| **実装の簡単さ**   | ⚠️ 中程度       | ⚠️ 中程度            | ✅ 簡単         | ✅ 簡単          | ✅ 簡単            |
| **最悪性能保証**   | ✅ O(log n)保証 | ✅ O(log n)保証      | ❌ O(n)可能     | ❌ O(n)可能      | ✅ O(log n)保証    |

## 制約条件別選択

| 制約条件                | 推奨選択            | 理由                       |
| ----------------------- | ------------------- | -------------------------- |
| **メモリ < 1MB**        | Red-Black Tree      | 最小限の追加情報           |
| **リアルタイム制約**    | Red-Black Tree      | 予測可能な性能             |
| **頻繁な更新**          | Red-Black Tree      | 効率的な挿入・削除         |
| **読み取り専用**        | AVL Tree            | 厳密な平衡で最高の検索性能 |
| **大量データ（GB 級）** | B-Tree              | ディスク I/O 効率          |
| **局所性アクセス**      | Splay Tree          | アクセスパターン適応       |
| **並行処理**            | Lock-free Skip List | 並行制御の容易さ           |
| **組み込み環境**        | Red-Black Tree      | 安定した性能と小メモリ     |

# まとめ

Red-Black Tree（赤黒木）は、色の概念を導入した自己平衡二分探索木で、最悪の場合でも O(log n)の性能を保証しながら、AVL 木よりも挿入・削除のオーバーヘッドが少ない実用的なデータ構造です。

5 つの基本性質により平衡を維持し、多くの標準ライブラリやシステムソフトウェアで採用されています。Go 言語での実装においても、適切な回転操作と色管理により効率的な赤黒木を構築できます。

特に、汎用的な辞書・マップ、リアルタイムシステム、データベースインデックスなど、安定した性能が要求される場面で優秀な性能を発揮します。実装の複雑さと性能のバランスを考慮すると、最も実用的な平衡木の一
