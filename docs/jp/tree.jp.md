# 木構造（Tree）とは

木構造は、ノード（節点）とエッジ（辺）によって構成される階層的なデータ構造です。親子関係によってデータを組織化し、一つの根（ルート）から始まって枝分かれしていく形状を持ちます。

**このデータ構造を使う一番の利点は、階層的なデータを効率的に表現・検索・操作できる点です。**

ルートノードから始まり、各ノードが子ノードを持つ階層構造で、循環のない連結グラフの特殊な形です。

# 基本的な特徴

- **階層構造**: 親子関係による明確な階層
- **単一ルート**: 一つの根ノードから開始
- **非循環**: サイクルが存在しない
- **連結**: 全てのノードが繋がっている
- **用途**: ファイルシステム、データベースインデックス、構文解析、決定木など

# 行える処理

## 木構造（Tree）の基本操作

| 機能        | 説明                             | 計算量          | 戻り値              | 得意/苦手 | 補足（その他）                           |
| ----------- | -------------------------------- | --------------- | ------------------- | --------- | ---------------------------------------- |
| Insert      | 新しいノードを木に追加           | O(log n)～ O(n) | なし（void）        | ✅        | 木の種類により性能が異なる。             |
| Delete      | 指定されたノードを木から削除     | O(log n)～ O(n) | bool（成功/失敗）   | ⚠️        | 子ノードの処理が必要。実装が複雑。       |
| Search/Find | 指定された値を持つノードを検索   | O(log n)～ O(n) | ノードまたは null   | ✅        | 平衡木では効率的。非平衡木では線形時間。 |
| Traverse    | 木の全ノードを訪問               | O(n)            | なし（void）        | ✅        | 前順・中順・後順・レベル順の走査が可能。 |
| GetRoot     | ルートノードを取得               | O(1)            | ルートノード        | ✅        | 木の入り口となるノード。                 |
| GetParent   | 指定ノードの親ノードを取得       | O(1)            | 親ノードまたは null | ✅        | 親への参照を保持している場合。           |
| GetChildren | 指定ノードの子ノードリストを取得 | O(1)            | 子ノードの配列      | ✅        | 子ノードへの参照リスト。                 |
| GetHeight   | 木またはノードの高さを取得       | O(n)            | int（高さ）         | ⚠️        | 全体を走査する必要がある場合が多い。     |
| GetDepth    | 指定ノードの深さを取得           | O(log n)～ O(n) | int（深さ）         | ✅        | ルートからの距離。                       |
| IsLeaf      | 指定ノードが葉ノードかチェック   | O(1)            | bool（葉/内部）     | ✅        | 子ノードの有無を確認。                   |
| IsEmpty     | 木が空かチェック                 | O(1)            | bool（空/非空）     | ✅        | ルートノードの存在確認。                 |
| Size/Count  | 木内のノード数を取得             | O(1)～ O(n)     | int（ノード数）     | ⚠️        | カウンタ保持の有無により異なる。         |
| Clear       | 木の全ノードを削除               | O(n)            | なし（void）        | ✅        | 全ノードのメモリ解放。                   |

## 木構造の特殊操作

| 機能                | 説明                               | 計算量          | 戻り値                  | 得意/苦手 | 補足（その他）                          |
| ------------------- | ---------------------------------- | --------------- | ----------------------- | --------- | --------------------------------------- |
| PreOrderTraversal   | 前順走査（根 → 左 → 右）           | O(n)            | 訪問順序の配列          | ✅        | 木のコピーや構造の保存に適している。    |
| InOrderTraversal    | 中順走査（左 → 根 → 右）           | O(n)            | 訪問順序の配列          | ✅        | 二分探索木でソート済み順序を取得。      |
| PostOrderTraversal  | 後順走査（左 → 右 → 根）           | O(n)            | 訪問順序の配列          | ✅        | ノードの削除や計算処理に適している。    |
| LevelOrderTraversal | レベル順走査（幅優先）             | O(n)            | 訪問順序の配列          | ✅        | 木を層ごとに処理。                      |
| FindLCA             | 最低共通祖先を検索                 | O(log n)～ O(n) | 共通祖先ノード          | ⚠️        | 前処理により高速化可能。                |
| GetPath             | ルートから指定ノードへのパスを取得 | O(log n)～ O(n) | ノードの配列            | ✅        | 経路の復元。                            |
| IsSubtree           | 部分木かどうかをチェック           | O(n×m)          | bool（部分木/非部分木） | ❌        | 複雑な比較処理が必要。                  |
| Clone/Copy          | 木の完全な複製を作成               | O(n)            | 新しい木                | ✅        | 全ノードを複製。                        |
| Serialize           | 木をシリアル化                     | O(n)            | 文字列またはバイト列    | ✅        | 永続化や通信に使用。                    |
| Deserialize         | シリアル化データから木を復元       | O(n)            | 復元された木            | ✅        | シリアル化の逆操作。                    |
| Balance             | 木のバランスを調整                 | O(n)            | なし（void）            | ❌        | 平衡木への変換。実装が複雑。            |
| Rotate              | ノードの回転操作                   | O(1)            | なし（void）            | ⚠️        | AVL 木や Red-Black 木でのバランス調整。 |

## 注意点

1. **メモリリーク**: ノード削除時の適切なメモリ管理が必要
2. **非平衡時の性能劣化**: 線形構造に近づくと検索性能が著しく低下
3. **循環参照**: 親への参照を持つ場合のメモリリーク対策が必要
4. **並行アクセス**: 複数スレッドからの同時アクセス時の整合性確保

# 実装方法

## 実装方式別の比較

| 実装方式           | メリット                               | デメリット                               | 適用場面           | 推奨実装         |
| ------------------ | -------------------------------------- | ---------------------------------------- | ------------------ | ---------------- |
| **配列ベース**     | - キャッシュ効率が良い<br>- メモリ効率 | - 動的サイズ変更が困難<br>- 再配置コスト | 完全二分木、ヒープ | 固定配列         |
| **ポインタベース** | - 動的サイズ<br>- 柔軟な構造変更       | - キャッシュ効率悪い<br>- メモリ断片化   | 一般的な木構造     | 構造体とポインタ |
| **連想配列ベース** | - 実装が簡単<br>- 動的構造変更         | - メモリオーバーヘッド<br>- 性能劣化     | プロトタイピング   | map[NodeID]Node  |

## 配列ベース実装例（完全二分木）

```go
type ArrayBinaryTree struct {
    data []int
    size int
}

func NewArrayBinaryTree(capacity int) *ArrayBinaryTree {
    return &ArrayBinaryTree{
        data: make([]int, capacity),
        size: 0,
    }
}

// 親ノードのインデックス
func (t *ArrayBinaryTree) parent(i int) int {
    return (i - 1) / 2
}

// 左の子ノードのインデックス
func (t *ArrayBinaryTree) leftChild(i int) int {
    return 2*i + 1
}

// 右の子ノードのインデックス
func (t *ArrayBinaryTree) rightChild(i int) int {
    return 2*i + 2
}

func (t *ArrayBinaryTree) Insert(value int) bool {
    if t.size >= len(t.data) {
        return false // 容量不足
    }

    t.data[t.size] = value
    t.size++
    return true
}

// レベル順走査
func (t *ArrayBinaryTree) LevelOrderTraversal() []int {
    result := make([]int, 0, t.size)
    for i := 0; i < t.size; i++ {
        result = append(result, t.data[i])
    }
    return result
}
```

### 特徴

- **最高のキャッシュ効率**: 連続したメモリレイアウト
- **簡単なインデックス計算**: 親子関係が数式で表現可能
- **メモリ効率**: ポインタが不要
- **高速アクセス**: O(1)でのランダムアクセス

### 注意事項

- **完全二分木限定**: 不完全な木構造には不適
- **固定容量**: 事前のサイズ設計が重要
- **メモリ無駄**: 疎な木では未使用領域が多数発生
- **動的変更困難**: 構造変更時の配列再配置が必要

## ポインタベース実装例（一般的な木）

```go
type TreeNode struct {
    Value    int
    Children []*TreeNode
    Parent   *TreeNode
}

type Tree struct {
    Root *TreeNode
    size int
}

func NewTree() *Tree {
    return &Tree{}
}

func (t *Tree) Insert(parentValue, newValue int) bool {
    if t.Root == nil {
        t.Root = &TreeNode{Value: newValue}
        t.size++
        return true
    }

    parent := t.search(t.Root, parentValue)
    if parent == nil {
        return false // 親ノードが見つからない
    }

    newNode := &TreeNode{
        Value:  newValue,
        Parent: parent,
    }
    parent.Children = append(parent.Children, newNode)
    t.size++
    return true
}

func (t *Tree) search(node *TreeNode, value int) *TreeNode {
    if node == nil {
        return nil
    }

    if node.Value == value {
        return node
    }

    for _, child := range node.Children {
        if result := t.search(child, value); result != nil {
            return result
        }
    }

    return nil
}

// 前順走査
func (t *Tree) PreOrderTraversal() []int {
    var result []int
    t.preOrder(t.Root, &result)
    return result
}

func (t *Tree) preOrder(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }

    *result = append(*result, node.Value)
    for _, child := range node.Children {
        t.preOrder(child, result)
    }
}

// ノード削除（子ノードも削除）
func (t *Tree) Delete(value int) bool {
    if t.Root == nil {
        return false
    }

    // ルートノードの削除
    if t.Root.Value == value {
        t.deleteSubtree(t.Root)
        t.Root = nil
        return true
    }

    return t.deleteNode(t.Root, value)
}

func (t *Tree) deleteNode(node *TreeNode, value int) bool {
    for i, child := range node.Children {
        if child.Value == value {
            // 子ノードとその部分木を削除
            t.deleteSubtree(child)
            // スライスから削除
            node.Children = append(node.Children[:i], node.Children[i+1:]...)
            return true
        }

        if t.deleteNode(child, value) {
            return true
        }
    }
    return false
}

func (t *Tree) deleteSubtree(node *TreeNode) {
    if node == nil {
        return
    }

    for _, child := range node.Children {
        t.deleteSubtree(child)
    }

    t.size--
    // Goのガベージコレクタが自動的にメモリを回収
}
```

### 特徴

- **柔軟な構造**: 任意の形状の木を表現可能
- **動的サイズ**: 実行時の構造変更が容易
- **直感的実装**: 木の概念を直接表現
- **メモリ効率**: 必要な分だけメモリを使用

### 注意事項

- **メモリ断片化**: ノードが分散配置される
- **キャッシュミス**: 非連続メモリアクセス
- **ポインタ管理**: 親子関係の整合性確保が重要
- **メモリリーク**: 削除時の適切な後処理が必要

## 二分探索木実装例（BST）

```go
type BSTNode struct {
    Value int
    Left  *BSTNode
    Right *BSTNode
}

type BinarySearchTree struct {
    Root *BSTNode
    size int
}

func NewBST() *BinarySearchTree {
    return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insert(value int) {
    bst.Root = bst.insertNode(bst.Root, value)
}

func (bst *BinarySearchTree) insertNode(node *BSTNode, value int) *BSTNode {
    if node == nil {
        bst.size++
        return &BSTNode{Value: value}
    }

    if value < node.Value {
        node.Left = bst.insertNode(node.Left, value)
    } else if value > node.Value {
        node.Right = bst.insertNode(node.Right, value)
    }
    // 重複値は挿入しない

    return node
}

func (bst *BinarySearchTree) Search(value int) bool {
    return bst.searchNode(bst.Root, value)
}

func (bst *BinarySearchTree) searchNode(node *BSTNode, value int) bool {
    if node == nil {
        return false
    }

    if value == node.Value {
        return true
    } else if value < node.Value {
        return bst.searchNode(node.Left, value)
    } else {
        return bst.searchNode(node.Right, value)
    }
}

func (bst *BinarySearchTree) Delete(value int) bool {
    var deleted bool
    bst.Root, deleted = bst.deleteNode(bst.Root, value)
    if deleted {
        bst.size--
    }
    return deleted
}

func (bst *BinarySearchTree) deleteNode(node *BSTNode, value int) (*BSTNode, bool) {
    if node == nil {
        return nil, false
    }

    var deleted bool

    if value < node.Value {
        node.Left, deleted = bst.deleteNode(node.Left, value)
    } else if value > node.Value {
        node.Right, deleted = bst.deleteNode(node.Right, value)
    } else {
        // 削除対象ノードを発見
        deleted = true

        if node.Left == nil {
            return node.Right, deleted
        } else if node.Right == nil {
            return node.Left, deleted
        }

        // 両方の子を持つ場合：右部分木の最小値で置き換え
        minNode := bst.findMin(node.Right)
        node.Value = minNode.Value
        node.Right, _ = bst.deleteNode(node.Right, minNode.Value)
    }

    return node, deleted
}

func (bst *BinarySearchTree) findMin(node *BSTNode) *BSTNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}

// 中順走査（ソート済み順序）
func (bst *BinarySearchTree) InOrderTraversal() []int {
    var result []int
    bst.inOrder(bst.Root, &result)
    return result
}

func (bst *BinarySearchTree) inOrder(node *BSTNode, result *[]int) {
    if node == nil {
        return
    }

    bst.inOrder(node.Left, result)
    *result = append(*result, node.Value)
    bst.inOrder(node.Right, result)
}
```

### 特徴

- **効率的検索**: 平均 O(log n)での検索・挿入・削除
- **ソート済み出力**: 中順走査でソート済み順序を取得
- **範囲検索**: 特定範囲の値を効率的に取得可能
- **動的操作**: 実行時の要素追加・削除が効率的

### 注意事項

- **非平衡時の劣化**: 偏った挿入で線形構造になるリスク
- **削除の複雑さ**: 特に両方の子を持つノードの削除
- **重複値処理**: 同一値の扱いを明確に定義する必要
- **平衡維持**: AVL 木や Red-Black 木の検討が必要な場合あり

## 実装選択の指針

| 要件                   | 推奨実装       | 理由                           |
| ---------------------- | -------------- | ------------------------------ |
| **検索性能重視**       | 二分探索木     | 平均 O(log n)での効率的検索    |
| **メモリ効率重視**     | 配列ベース     | ポインタオーバーヘッドなし     |
| **柔軟な構造変更**     | ポインタベース | 任意の構造変更が容易           |
| **キャッシュ効率重視** | 配列ベース     | 連続メモリアクセス             |
| **完全二分木**         | 配列ベース     | インデックス計算で高速アクセス |
| **動的サイズ変更**     | ポインタベース | 実行時の柔軟なサイズ調整       |
| **ソート済みデータ**   | 二分探索木     | 中順走査でソート済み順序を取得 |
| **階層データ表現**     | ポインタベース | 自然な階層構造の表現           |
| **高速レンジ検索**     | 二分探索木     | 特定範囲の効率的検索           |
| **プロトタイピング**   | 連想配列ベース | 実装の簡単さ                   |

# 他のデータ構造との比較

## 木構造実装方式別比較

| 特徴               | 配列ベース木       | ポインタベース木          | 二分探索木                |
| ------------------ | ------------------ | ------------------------- | ------------------------- |
| **メモリ使用量**   | ✅ 最効率          | ❌ ポインタオーバーヘッド | ❌ ポインタオーバーヘッド |
| **キャッシュ効率** | ✅ 最適            | ❌ 非連続メモリ           | ❌ 非連続メモリ           |
| **検索性能**       | ❌ O(n)線形        | ❌ O(n)線形               | ✅ O(log n)平均           |
| **挿入性能**       | ⚠️ O(1)制限あり    | ✅ O(1)位置特定後         | ✅ O(log n)平均           |
| **削除性能**       | ❌ O(n)再構築      | ✅ O(1)位置特定後         | ✅ O(log n)平均           |
| **構造柔軟性**     | ❌ 完全二分木のみ  | ✅ 任意構造               | ⚠️ 二分木のみ             |
| **バランス保証**   | ✅ 自動的に完全    | ❌ 手動管理               | ❌ 手動管理               |
| **実装複雑さ**     | ✅ シンプル        | ⚠️ 中程度                 | ❌ 複雑                   |
| **適用場面**       | ヒープ、完全二分木 | 一般的な階層データ        | 検索・ソート              |

## 全データ構造比較

| 特徴               | 二分探索木                | 配列（ソート済み）  | ハッシュテーブル | 連結リスト                | リングバッファ  |
| ------------------ | ------------------------- | ------------------- | ---------------- | ------------------------- | --------------- |
| **検索性能**       | ✅ O(log n)平均           | ✅ O(log n)二分探索 | ✅ O(1)平均      | ❌ O(n)線形               | ❌ O(n)線形     |
| **挿入性能**       | ✅ O(log n)平均           | ❌ O(n)シフト       | ✅ O(1)平均      | ✅ O(1)位置特定後         | ✅ O(1)         |
| **削除性能**       | ✅ O(log n)平均           | ❌ O(n)シフト       | ✅ O(1)平均      | ✅ O(1)位置特定後         | ✅ O(1)         |
| **ソート順序**     | ✅ 中順走査で取得         | ✅ 常にソート済み   | ❌ 順序なし      | ❌ 挿入順序               | ❌ 挿入順序     |
| **メモリ使用量**   | ❌ ポインタオーバーヘッド | ✅ 最効率           | ⚠️ 負荷率に依存  | ❌ ポインタオーバーヘッド | ✅ 固定・効率的 |
| **範囲検索**       | ✅ 効率的                 | ✅ 効率的           | ❌ 全走査必要    | ❌ 全走査必要             | ❌ 不適         |
| **キャッシュ効率** | ❌ 非連続メモリ           | ✅ 最適             | ✅ 良好          | ❌ 非連続メモリ           | ✅ 最適         |
| **動的サイズ**     | ✅ 柔軟                   | ⚠️ 再割当コスト     | ✅ 柔軟          | ✅ 柔軟                   | ❌ 固定サイズ   |
| **実装複雑さ**     | ❌ 複雑                   | ✅ シンプル         | ⚠️ 中程度        | ✅ シンプル               | ✅ シンプル     |
| **階層表現**       | ✅ 自然                   | ❌ 不適             | ❌ 不適          | ⚠️ 可能だが非効率         | ❌ 不適         |

# 使用場面別最適選択

| 使用場面                     | 第 1 選択              | 第 2 選択        | 第 3 選択        | 避けるべき                   |
| ---------------------------- | ---------------------- | ---------------- | ---------------- | ---------------------------- |
| **ファイルシステム**         | ポインタベース木       | 二分探索木       | -                | 配列、リングバッファ         |
| **データベースインデックス** | B+木（専用実装）       | 二分探索木       | ハッシュテーブル | 配列、連結リスト             |
| **構文解析木**               | ポインタベース木       | -                | -                | 配列ベース木                 |
| **決定木**                   | ポインタベース木       | 二分探索木       | -                | リングバッファ               |
| **ヒープ実装**               | 配列ベース木           | -                | -                | ポインタベース木             |
| **優先度付きキュー**         | 配列ベース木（ヒープ） | 二分探索木       | -                | 連結リスト                   |
| **辞書・検索システム**       | 二分探索木             | ハッシュテーブル | -                | 配列、連結リスト             |
| **組織図・階層データ**       | ポインタベース木       | -                | -                | 配列ベース木、リングバッファ |
| **ゲームの状態空間探索**     | ポインタベース木       | -                | -                | 配列ベース木                 |
| **コンパイラの中間表現**     | ポインタベース木       | -                | -                | 配列ベース木                 |
| **XML パーサー**             | ポインタベース木       | -                | -                | 配列ベース木                 |
| **数式処理**                 | ポインタベース木       | -                | -                | 配列ベース木                 |

# パフォーマンス特性

| 操作パターン         | 配列ベース木      | ポインタベース木  | 二分探索木        |
| -------------------- | ----------------- | ----------------- | ----------------- |
| **頻繁な検索**       | ❌ 線形検索       | ❌ 線形検索       | ✅ 対数時間検索   |
| **大量データ処理**   | ✅ キャッシュ効率 | ❌ キャッシュミス | ⚠️ バランス次第   |
| **動的構造変更**     | ❌ 再構築必要     | ✅ 柔軟対応       | ✅ 効率的変更     |
| **メモリ制約環境**   | ✅ 最小メモリ     | ❌ ポインタ分重い | ❌ ポインタ分重い |
| **リアルタイム処理** | ✅ 予測可能       | ⚠️ GC 影響あり    | ⚠️ バランス次第   |
| **並行アクセス**     | ⚠️ 配列操作の同期 | ❌ 複雑な同期制御 | ❌ 複雑な同期制御 |

木構造は階層的なデータの表現と操作に最適化されたデータ構造で、特に検索・ソート・階層管理が重要な場面で威力を発揮します。用途に応じた適切な実装方式の選択により、最適な性能を実現できます。
