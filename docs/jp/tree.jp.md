# 木構造（Tree）とは

木構造は、ノード（頂点、Vertex）とそれらを接続するエッジ（辺、Edge）で構成される階層的なデータ構造。木構造は、ルート（根）ノードから始まり、各ノードが 0 個以上の子ノードを持つことができる。木構造は、グラフの特殊な形であり、サイクルを持たない連結グラフとして定義される。

**階層的なデータの表現や効率的な検索、挿入、削除操作が可能。**

## 基本要素

```
                    A (ルート)
                   / \
                  /   \
                 B     C
                / \   / \
               D   E F   G (葉)
              /
             H (葉)
```

- **ルート（Root）**: 木構造の最上位に位置するノード。
- **ノード（Node）**: データを保持する単位。親ノードや子ノードを持つ。
- **エッジ（Edge）**: ノード間の接続を表す線。
- **葉（Leaf）**: 子ノードを持たないノード。
- **高さ（Height）**: ルートから最も深い葉までのエッジ数。
- **深さ（Depth）**: ルートから特定のノードまでのエッジ数。
- **親（Parent）**: あるノードに直接接続する上位のノード。
- **子（Child）**: あるノードに直接接続する下位のノード。
- **部分木（Subtree）**: 特定のノードをルートとする木。
  - **B を根とする部分木**: B, D, E, H
  - **C を根とする部分木**: C, F, G
  - **D を根とする部分木**: D, H

## 基本的な特性

- 一つの木構造には必ず一つのルートが存在する
- ルート以外の各ノードは、一つだけ親ノードを持つ
- 各ノードは 0 個以上の子ノードを持つことができる
- サイクル（閉路）は存在しない

# 木構造の種類

## 二分木（Binary Tree）

- 各ノードが最大 2 つの子ノードを持つ木。
- 用途: 数式の表現、ヒープ構造など。

## 完全二分木（Complete Binary Tree）

- すべてのレベルが完全に埋められているか、最下層のみが左詰めで埋められている二分木。

```
      1
     / \
    2   3
   / \ /
  4  5 6    ← 最下層が左詰め
```

## 完全木（Full Binary Tree）

- 各ノードが 0 個または 2 個の子ノードを持つ二分木。

```
      1
     / \
    2   3
   / \
  4   5    ← 各ノードが0個または2個の子
```

## ヒープ（Heap）

- 親ノードが子ノードよりも大きい（または小さい）値を持つ完全二分木。
- 用途: 優先度付きキュー、ソート。

```
配列表現: [10, 15, 20, 17, 25, 30, 22]
インデックス: 0  1  2  3  4  5  6

木構造表現:
          10 (0)
         /      \
      15(1)    20(2)
     /    \    /    \
   17(3) 25(4) 30(5) 22(6)
```

## 二分探索木（Binary Search Tree, BST）

- 木構造の中でも、検索操作に特化した構造
- 左部分木のすべてのノードの値が親ノードより小さく、右部分木のすべてのノードの値が親ノードより大きい。O(log n)で目的の値を見つけられる。
- 現在のノードのキーと探索対象のキーを比較し、小さければ左へ、大きければ右へと辿ることで、探索範囲を半分に狭めながら進めることができる。
- データがランダムに挿入された場合、木は比較的にバランスが取れ、高さは O(logN) となり、各操作も O(logN) で実行できる。しかし、データがソートされた順序（またはそれに近い順序）で挿入されると、木は一方に偏った線形リストのような形状になってしまう。このような木を**縮退木（degenerate tree）**と呼び、この場合、木の高さは N に比例し、探索性能は線形探索と同じ O(N) まで劣化してしまう 。
- 用途: 高速な検索、挿入、削除。

```
          50
         /  \
        30   70
       / \   / \
      20 40 60 80
     /
    10
```

## 平衡二分探索木（Balanced Binary Search Tree）

- 自己平衡二分探索木ともいう
- 高さがバランスされた二分探索木
- データの挿入や削除が行われたとき、木が偏らないように構造を自動的に修正するメカニズムを備えた二分探索木
- 木回転という操作により、木の親子関係を局所的に自動で変更することで、二分探索木の特性を維持したままバランスを調整できる
- 木回転の操作は数個のポインタの付け替えなので、O(1) の定数時間で実行できる
- 用途: 最悪計算量を抑えた効率的な操作
- 例: AVL 木、赤黒木

### AVL 木

木回転を行う条件は厳格で、「任意のノードにおいて、左右の部分きの高さの差が 1 以下でなければならない」

- 左右の部分木の他 k さのさをバランスファクターという
- 挿入・削除操作によって、あるノードのバランスファクターが ±2 になると、自動で修正される

バランスが取れている状態

```
        20 (BF=0, h=2)
       /              \
   10 (BF=0, h=1)  30 (BF=0, h=1)
   /         \      /         \
5 (BF=0,h=0) 15(BF=0,h=0) 25(BF=0,h=0) 35(BF=0,h=0)

BF = Balance Factor (左の高さ - 右の高さ)
h = height (そのノードの高さ)

全てのノードでバランスファクターが -1, 0, 1 のいずれか
```

バランスが取れていない状態

```
バランス崩壊前（ノード30でBF=2）:
        30 (BF=2)
       /
    20 (BF=1)
   /
 10 (BF=0)

右回転後（バランス回復）:
    20 (BF=0)
   /         \
10 (BF=0)  30 (BF=0)
```

### 赤黒木

AVL 木よりもゆるい条件で、以下の 5 つのルールを維持する

- ルール
  1. 各ノードは「赤」か「黒」のいずれか
  2. ルートノードは「黒」
  3. 全ての葉ノードは「黒」
  4. 「赤」ノードの子ノードは、両方とも「黒」（「赤」ノードは連続しない）
  5. 任意のノードからその子孫の葉ノードまでの全てのパスには、同じ数の「黒」ノードが含まれる
- 最も長いパス（「赤」「黒」が交互に続く）の長さが、最も短いパス（「黒」のみのパス）の長さの 2 倍を超えることがなく、木の高さは O(log n)に保たれる

バランスが取れている状態

```
              30(B)
           /         \
       15(R)         50(R)
      /     \       /     \
   10(B)   20(B) 40(B)   60(B)
  /    \         /  \   /    \
5(R)  12(R)   35(R) 45(R) 55(R) 70(R)

特徴:
- 最長パス（赤黒交互）: 30→50→60→70 = 4エッジ
- 最短パス（黒のみ）: 30→15→10 = 3エッジ
- 最長パス ≤ 2 × 最短パス（赤黒木の性質）
```

### AVL 木と赤黒木の比較

| 特性               | AVL 木                                                           | 赤黒木                                             |
| ------------------ | ---------------------------------------------------------------- | -------------------------------------------------- |
| **検索時間**       | より高速（木の高さが厳密に揃えられるため、木の高さが比較的低い） | やや遅い（ゆるい平衡のため、木の高さが比較的高い） |
| **挿入・削除時間** | やや遅い（回転多い）                                             | より高速（回転少ない）                             |
| **バランス規則**   | 厳格（左右部分木の高さの差が 1 以下）                            | 緩やか（5 つのルールに基づく）                     |
| **メモリ**         | 高さ情報が必要                                                   | 色情報のみ（1 ビット）                             |
| **最大高さ**       | 1.44 × log(n)                                                    | 2 × log(n)                                         |
| **ユースケース**   | 検索重視（読み取り >> 書き込み）                                 | 挿入・削除が頻繁、バランスの取れた性能             |

# 木構造の実装方法

## 一般的な実装方法

### 配列を用いた実装

- 完全二分木に適している。
- ノードのインデックスを利用して親子関係を表現。
  - 親ノードのインデックス: `i`
  - 左子ノードのインデックス: `2i + 1`
  - 右子ノードのインデックス: `2i + 2`

### ポインタを用いた実装

- ノードを構造体で表現し、ポインタで親子関係をリンク。
- 任意の木構造に適している。

## Go を用いた実装例

### 二分探索木の実装例

```go
package tree

// ノード構造体
type Node struct {
    Key   int
    Left  *Node
    Right *Node
}

// 二分探索木構造体
type BinarySearchTree struct {
    Root *Node
}

// 新しいノードを作成
func NewNode(key int) *Node {
    return &Node{Key: key}
}

// ノードを挿入
func (bst *BinarySearchTree) Insert(key int) {
    bst.Root = insertNode(bst.Root, key)
}

func insertNode(node *Node, key int) *Node {
    if node == nil {
        return NewNode(key)
    }
    if key < node.Key {
        node.Left = insertNode(node.Left, key)
    } else if key > node.Key {
        node.Right = insertNode(node.Right, key)
    }
    return node
}

// ノードを検索
func (bst *BinarySearchTree) Search(key int) *Node {
    return searchNode(bst.Root, key)
}

func searchNode(node *Node, key int) *Node {
    if node == nil || node.Key == key {
        return node
    }
    if key < node.Key {
        return searchNode(node.Left, key)
    }
    return searchNode(node.Right, key)
}

// ノードを削除
func (bst *BinarySearchTree) Delete(key int) {
    bst.Root = deleteNode(bst.Root, key)
}

func deleteNode(node *Node, key int) *Node {
    if node == nil {
        return nil
    }
    if key < node.Key {
        node.Left = deleteNode(node.Left, key)
    } else if key > node.Key {
        node.Right = deleteNode(node.Right, key)
    } else {
        // 子ノードが 0 または 1 の場合
        if node.Left == nil {
            return node.Right
        } else if node.Right == nil {
            return node.Left
        }
        // 子ノードが 2 の場合
        minRight := findMin(node.Right)
        node.Key = minRight.Key
        node.Right = deleteNode(node.Right, minRight.Key)
    }
    return node
}

func findMin(node *Node) *Node {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}
```

### 平衡二分探索木（例: AVL 木）の実装例

平衡二分探索木では、挿入や削除時に高さのバランスを保つための回転操作（右回転、左回転）が必要です。以下は AVL 木の基本的な構造です。

```go
// AVL 木のノード構造体
type AVLNode struct {
    Key    int
    Height int
    Left   *AVLNode
    Right  *AVLNode
}

// AVL 木構造体
type AVLTree struct {
    Root *AVLNode
}

// ノードの高さを取得
func height(node *AVLNode) int {
    if node == nil {
        return 0
    }
    return node.Height
}

// 高さを更新
func updateHeight(node *AVLNode) {
    node.Height = max(height(node.Left), height(node.Right)) + 1
}

// バランス因子を計算
func balanceFactor(node *AVLNode) int {
    if node == nil {
        return 0
    }
    return height(node.Left) - height(node.Right)
}

// 右回転
func rotateRight(y *AVLNode) *AVLNode {
    x := y.Left
    T := x.Right

    x.Right = y
    y.Left = T

    updateHeight(y)
    updateHeight(x)

    return x
}

// 左回転
func rotateLeft(x *AVLNode) *AVLNode {
    y := x.Right
    T := y.Left

    y.Left = x
    x.Right = T

    updateHeight(x)
    updateHeight(y)

    return y
}

// ノードを挿入（バランス調整を含む）
func (tree *AVLTree) Insert(key int) {
    tree.Root = insertAVLNode(tree.Root, key)
}

func insertAVLNode(node *AVLNode, key int) *AVLNode {
    if node == nil {
        return &AVLNode{Key: key, Height: 1}
    }
    if key < node.Key {
        node.Left = insertAVLNode(node.Left, key)
    } else if key > node.Key {
        node.Right = insertAVLNode(node.Right, key)
    } else {
        return node // 重複キーは無視
    }

    updateHeight(node)

    // バランス調整
    balance := balanceFactor(node)
    if balance > 1 && key < node.Left.Key {
        return rotateRight(node)
    }
    if balance < -1 && key > node.Right.Key {
        return rotateLeft(node)
    }
    if balance > 1 && key > node.Left.Key {
        node.Left = rotateLeft(node.Left)
        return rotateRight(node)
    }
    if balance < -1 && key < node.Right.Key {
        node.Right = rotateRight(node.Right)
        return rotateLeft(node)
    }

    return node
}
```

# 木構造の応用例

## 1. ファイルシステム・ディレクトリ構造

- **ディレクトリ階層**: フォルダとファイルの階層構造
- **パス管理**: ルートからファイルまでの一意な経路
- **権限管理**: 親ディレクトリから子への権限継承
- **検索システム**: ディレクトリを辿った効率的なファイル検索

## 2. データベース・インデックス

- **B 木インデックス**: データベースの効率的な検索インデックス
- **B+木**: 範囲検索に最適化されたデータベース構造
- **二分探索木**: メモリ内データの高速検索
- **クエリ最適化**: SQL 文の実行計画の木構造表現

## 3. コンパイラ・言語処理

- **抽象構文木（AST）**: プログラムの構文構造を表現
- **パースツリー**: 文法解析の結果を木構造で表現
- **式の評価**: 数式の演算子優先順位を木で管理
- **スコープ管理**: 変数のスコープ階層を木で表現

## 4. 組織・階層管理

- **組織図**: 企業の部門・役職の階層構造
- **権限管理**: 組織階層に基づく権限の継承
- **レポートライン**: 報告関係の明確な階層
- **人事システム**: 昇進・異動の経路管理

## 5. ネットワーク・通信

- **ルーティングツリー**: ネットワーク経路の最適化
- **スパニングツリー**: ネットワークループの防止
- **マルチキャスト配信**: 効率的なデータ配信経路
- **DNS 階層**: ドメイン名の階層構造

## 6. データ圧縮・符号化

- **ハフマン木**: 頻度に基づく効率的なデータ圧縮
- **エントロピー符号化**: 情報理論に基づく圧縮
- **可変長符号**: 文字頻度に応じた符号割り当て
- **圧縮アルゴリズム**: LZ 系アルゴリズムでの辞書管理

## 7. Web・UI・表示システム

- **DOM（Document Object Model）**: HTML の階層構造
- **XML パース**: マークアップ言語の構造解析
- **UI コンポーネント階層**: GUI の親子関係管理
- **レンダリングツリー**: Web ページの描画構造

## 8. 地理・空間データ

- **空間分割**: QuadTree、OctTree による空間の効率的分割
- **地理情報システム**: 地域の階層的な行政区分
- **ゲームマップ**: 3D ゲームの空間管理
- **衝突検出**: ゲームオブジェクトの効率的な衝突判定

## 9. 優先度・スケジューリング

- **ヒープ**: 優先度付きキューの実装
- **タスクスケジューラ**: OS のプロセス管理
- **イベント処理**: 優先度に基づくイベント管理
- **リソース割り当て**: 限られたリソースの効率的配分

## 10. 検索・情報検索

- **検索エンジン**: キーワード検索の効率化
- **トライ木**: 文字列検索・オートコンプリート
- **辞書検索**: 単語の前方一致検索
- **推薦システム**: ユーザー属性の階層分類

## 11. 生物学・系統学

- **系統樹**: 生物の進化系統関係
- **分類体系**: 界・門・綱・目・科・属・種の階層
- **遺伝子解析**: DNA 配列の類似性分析
- **進化系統学**: 種の分岐点と進化過程

## 12. 数学・アルゴリズム

- **セグメント木**: 範囲クエリの効率的処理
- **BIT（Binary Indexed Tree）**: 累積和の高速計算
- **区間木**: 区間に対する操作の最適化
- **フラクタル**: 自己相似図形の生成

## 13. 金融・リスク管理

- **ポートフォリオ管理**: 投資商品の階層分類
- **リスク評価**: リスクファクターの階層構造
- **オプション価格**: 二項ツリーモデル
- **金融商品分類**: 資産クラスの階層管理

# まとめ

木構造は、階層的なデータを効率的に管理するための基本的なデータ構造です。特に二分探索木や平衡二分探索木は、検索や挿入、削除操作を効率的に行うために広く使用されています。Go を用いた実装では、構造体とポインタを活用することで柔軟な木構造を表現でき、用途や要件に応じてさまざまなバリエーションを選択できます。

木構造はファイルシステムやデータベース、機械学習、ネットワークなど多くの分野で応用されており、データの階層管理や効率的な検索・更新処理を実現します。用途やデータ特性、性能要件に応じて最適な木構造を選択し、適切に実装することが重要です。
