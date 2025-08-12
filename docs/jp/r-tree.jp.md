# R-tree（R 木）とは

R-tree（R 木）は、**多次元空間データ**（主に 2 次元・3 次元の矩形領域）を効率的に管理するための**空間インデックス**データ構造です。各ノードが矩形領域（Minimum Bounding Rectangle, MBR）を持ち、階層的に空間を分割して管理します。

**このデータ構造の最大の利点は、地理情報システム（GIS）、ゲーム、CAD、画像処理などで空間検索・範囲検索を高速に実行できる点です。**

# R-tree の構成要素

## 基本要素

- **MBR（Minimum Bounding Rectangle）**: 最小外接矩形。ノードが管理する空間範囲を表現
- **ノード（Node）**: 空間データまたは子ノードへの参照を保持する単位
- **エントリ（Entry）**: ノード内の各要素。MBR とポインタ（または実データ）のペア
- **リーフノード（Leaf Node）**: 実際のデータオブジェクトを格納するノード
- **内部ノード（Internal Node）**: 子ノードへの参照を持つノード
- **ルートノード（Root Node）**: 木構造の最上位ノード
- **空間オブジェクト**: 点、線分、矩形、多角形などの幾何学的データ

## 空間的特性

- **階層分割**: 空間を階層的に矩形領域で分割
- **重複許可**: 兄弟ノードの MBR が重複することを許可
- **最小化原則**: MBR の面積・周囲長・重複を最小化
- **バランス**: B-tree ライクな高さバランス

# R-tree の種類

## R-tree（基本型）

- 最初に提案された基本的な R-tree
- 挿入時の分割アルゴリズムが単純

## R\*-tree

- R-tree の改良版
- より効率的な分割・再挿入アルゴリズム
- 用途: 高性能が要求される空間データベース

## R+-tree

- ノードの MBR の重複を禁止
- 検索性能は向上するが、更新コストが高い
- 用途: 読み取り重視のアプリケーション

## Hilbert R-tree

- ヒルベルト曲線を利用した空間順序
- より良いクラスタリング性能
- 用途: 大量の空間データ処理

# R-tree の特徴

- **空間検索**: 範囲検索、最近傍検索、交差検索が効率的
- **動的構造**: 空間オブジェクトの動的な追加・削除が可能
- **多次元対応**: 2 次元以上の空間データに対応
- **重複許可**: 兄弟ノードの MBR が重複可能（検索効率とのトレードオフ）
- **バランス保持**: 挿入・削除時に木の高さバランスを維持

# R-tree の構成要素の図解

## 基本的な R-tree の例（2 次元空間）

```
空間分割の例:
┌─────────────────────────────────┐
│               R1               │
│  ┌─────┐         ┌─────────┐   │
│  │ R2  │    ┌────┤   R3    │   │
│  │ ●A  │    │ ●E │    ●F   │   │
│  │  ●B │    │    │         │   │
│  └─────┘    └────┴─────────┘   │
│                                │
│  ┌─────────────┐               │
│  │     R4      │               │
│  │  ●C    ●D   │               │
│  └─────────────┘               │
└─────────────────────────────────┘

木構造表現:
        Root(R1)
       /    |    \
   R2(A,B) R3(E,F) R4(C,D)
```

### 各要素の詳細説明

**ノード数 = 4（ルート + 3 リーフ）**
**空間オブジェクト数 = 6（A, B, C, D, E, F）**

## 各要素の説明

### MBR（Minimum Bounding Rectangle）

- **R1, R2, R3, R4**: 各ノードの最小外接矩形
- 座標で表現: `{x_min, y_min, x_max, y_max}`
- 子ノードまたは空間オブジェクトを完全に包含

### ノード（Node）

```go
type RTreeNode struct {
    MBR      Rectangle     // ノードの最小外接矩形
    IsLeaf   bool         // リーフノードかどうか
    Entries  []Entry     // エントリの配列
    Parent   *RTreeNode  // 親ノードへの参照
}
```

### エントリ（Entry）

```go
type Entry struct {
    MBR    Rectangle    // エントリの最小外接矩形
    Child  *RTreeNode   // 内部ノードの場合: 子ノード
    Object interface{} // リーフノードの場合: 実際のデータ
}
```

### 空間オブジェクト

- **点**: 座標 (x, y)
- **矩形**: {x_min, y_min, x_max, y_max}
- **多角形**: 頂点座標の配列
- **線分**: 開始点と終了点の座標

## 3 次元 R-tree の例

```
3D空間での分割:
     Z軸
     │
     │  ┌─────────┐
     │ ╱         ╱│
     │╱  Box1   ╱ │
     └─────────┘  │
    ╱│            │
   ╱ └────────────┘
  ╱      Y軸
 X軸

MBR表現: {x_min, y_min, z_min, x_max, y_max, z_max}
```

# R-tree で行える処理

## 基本操作

| 機能       | 説明                         | 計算量       | 戻り値            | 得意/苦手 | 補足                          |
| ---------- | ---------------------------- | ------------ | ----------------- | --------- | ----------------------------- |
| Insert     | 空間オブジェクトを挿入       | O(log n)     | bool（成功/失敗） | ✅        | MBR の更新と分割処理          |
| Delete     | 空間オブジェクトを削除       | O(log n)     | bool（成功/失敗） | ✅        | MBR の縮小と併合処理          |
| Search     | 空間オブジェクトを検索       | O(log n)     | オブジェクト      | ✅        | MBR を辿って効率的に検索      |
| RangeQuery | 範囲内のオブジェクト検索     | O(log n + k) | オブジェクト配列  | ✅        | k: 結果数、空間検索の主要機能 |
| PointQuery | 点が含まれるオブジェクト検索 | O(log n + k) | オブジェクト配列  | ✅        | 特定座標での検索              |
| KNNQuery   | k 最近傍検索                 | O(log n + k) | オブジェクト配列  | ✅        | 距離順でのソート検索          |

## 空間検索操作

| 機能            | 説明         | 計算量       | 戻り値           | 得意/苦手 | 補足                         |
| --------------- | ------------ | ------------ | ---------------- | --------- | ---------------------------- |
| Intersects      | 交差判定検索 | O(log n + k) | オブジェクト配列 | ✅        | 重複する領域の検索           |
| Contains        | 包含判定検索 | O(log n + k) | オブジェクト配列 | ✅        | 完全に含まれるオブジェクト   |
| Within          | 内部判定検索 | O(log n + k) | オブジェクト配列 | ✅        | 指定領域内のオブジェクト     |
| Overlaps        | 重複判定検索 | O(log n + k) | オブジェクト配列 | ✅        | 部分的に重複するオブジェクト |
| Touches         | 接触判定検索 | O(log n + k) | オブジェクト配列 | ⚠️        | 境界での接触判定             |
| NearestNeighbor | 最近傍検索   | O(log n)     | オブジェクト     | ✅        | 最も近いオブジェクト 1 つ    |

## 構造操作・解析

| 機能      | 説明                    | 計算量     | 戻り値                | 得意/苦手 | 補足                       |
| --------- | ----------------------- | ---------- | --------------------- | --------- | -------------------------- |
| GetHeight | 木の高さを取得          | O(log n)   | int（高さ）           | ✅        | ルートからリーフまでの深さ |
| GetMBR    | ノードの MBR を取得     | O(1)       | Rectangle             | ✅        | 各ノードの最小外接矩形     |
| IsEmpty   | 空かチェック            | O(1)       | bool（空/非空）       | ✅        | ルートノードの存在確認     |
| Size      | オブジェクト数を取得    | O(n)       | int（オブジェクト数） | ❌        | 全ノードを走査する必要     |
| Validate  | R-tree 構造の妥当性確認 | O(n)       | bool（妥当/不正）     | ⚠️        | デバッグ・テスト用         |
| Optimize  | 木構造の最適化          | O(n log n) | 最適化済み木          | ❌        | 再構築による最適化         |

## 苦手な処理・制限

| 機能       | 説明                 | 計算量     | 戻り値           | 得意/苦手 | 補足                       |
| ---------- | -------------------- | ---------- | ---------------- | --------- | -------------------------- |
| ExactMatch | 完全一致検索         | O(n)       | オブジェクト     | ❌        | 空間構造上、効率的でない   |
| LinearScan | 線形探索             | O(n)       | オブジェクト配列 | ❌        | 空間インデックスの利点なし |
| TextSearch | テキスト検索         | O(n)       | オブジェクト配列 | ❌        | 空間データに不適           |
| Sort       | ソート               | O(n log n) | ソート済み配列   | ❌        | 空間的順序の定義が困難     |
| Merge      | 2 つの R-tree の結合 | O(n + m)   | 結合済み木       | ❌        | 再構築が必要               |
| Split      | R-tree の分割        | O(n)       | 分割済み木       | ❌        | 空間的分割の複雑さ         |

# R-tree の実装方法

## 一般的な実装方法

### ノード分割アルゴリズム

#### 線形分割（Linear Split）

- 最も単純な分割方法
- 最も離れた 2 つのエントリを基準に分割
- 実装が簡単だが、効率は劣る

#### 二次分割（Quadratic Split）

- より良い分割品質
- 面積増加を最小化する分割
- 計算コストと品質のバランス

#### 指数分割（Exponential Split）

- 最適な分割（全組み合わせを評価）
- 計算コストが高い
- 小規模なノードでのみ実用的

### R\*-tree の改良点

- **再挿入**: ノード分割前に一部エントリを再挿入
- **分割軸選択**: より良い分割軸の選択
- **分割点選択**: 重複最小化とマージン最小化

## Go による基本的な R-tree 実装例

```go
package rtree

import (
    "math"
)

// 矩形を表現する構造体
type Rectangle struct {
    MinX, MinY, MaxX, MaxY float64
}

// 空間オブジェクトのインターフェース
type SpatialObject interface {
    GetMBR() Rectangle
    GetID() interface{}
}

// R-treeのノード
type RTreeNode struct {
    MBR      Rectangle
    IsLeaf   bool
    Entries  []Entry
    Parent   *RTreeNode
}

// エントリ
type Entry struct {
    MBR    Rectangle
    Child  *RTreeNode
    Object SpatialObject
}

// R-tree構造体
type RTree struct {
    Root     *RTreeNode
    MaxEntries int // ノードの最大エントリ数
    MinEntries int // ノードの最小エントリ数
}

// 新しいR-treeを作成
func NewRTree(maxEntries int) *RTree {
    minEntries := maxEntries / 2
    if minEntries < 2 {
        minEntries = 2
    }

    return &RTree{
        MaxEntries: maxEntries,
        MinEntries: minEntries,
    }
}

// オブジェクトの挿入
func (rt *RTree) Insert(obj SpatialObject) {
    if rt.Root == nil {
        rt.Root = &RTreeNode{
            MBR:     obj.GetMBR(),
            IsLeaf:  true,
            Entries: []Entry{{MBR: obj.GetMBR(), Object: obj}},
        }
        return
    }

    rt.insert(rt.Root, obj)
}

func (rt *RTree) insert(node *RTreeNode, obj SpatialObject) {
    if node.IsLeaf {
        // リーフノードに直接挿入
        entry := Entry{MBR: obj.GetMBR(), Object: obj}
        node.Entries = append(node.Entries, entry)
        node.MBR = expandMBR(node.MBR, obj.GetMBR())

        // ノードが満杯の場合は分割
        if len(node.Entries) > rt.MaxEntries {
            rt.splitNode(node)
        }
    } else {
        // 最適な子ノードを選択
        bestChild := rt.chooseSubtree(node, obj.GetMBR())
        rt.insert(bestChild, obj)

        // MBRを更新
        node.MBR = rt.calculateNodeMBR(node)
    }
}

// 範囲検索
func (rt *RTree) RangeQuery(queryMBR Rectangle) []SpatialObject {
    var results []SpatialObject
    if rt.Root != nil {
        rt.rangeQuery(rt.Root, queryMBR, &results)
    }
    return results
}

func (rt *RTree) rangeQuery(node *RTreeNode, queryMBR Rectangle, results *[]SpatialObject) {
    if !intersects(node.MBR, queryMBR) {
        return
    }

    if node.IsLeaf {
        for _, entry := range node.Entries {
            if intersects(entry.MBR, queryMBR) {
                *results = append(*results, entry.Object)
            }
        }
    } else {
        for _, entry := range node.Entries {
            rt.rangeQuery(entry.Child, queryMBR, results)
        }
    }
}

// 最適な子ノードを選択
func (rt *RTree) chooseSubtree(node *RTreeNode, mbr Rectangle) *RTreeNode {
    var bestChild *RTreeNode
    minExpansion := math.Inf(1)

    for _, entry := range node.Entries {
        expansion := calculateExpansion(entry.MBR, mbr)
        if expansion < minExpansion {
            minExpansion = expansion
            bestChild = entry.Child
        }
    }

    return bestChild
}

// ノード分割（簡略化版）
func (rt *RTree) splitNode(node *RTreeNode) {
    if len(node.Entries) <= rt.MaxEntries {
        return
    }

    // 線形分割アルゴリズム
    seed1, seed2 := rt.pickSeeds(node.Entries)

    newNode := &RTreeNode{
        IsLeaf:  node.IsLeaf,
        Parent:  node.Parent,
        Entries: []Entry{node.Entries[seed2]},
    }

    node.Entries = []Entry{node.Entries[seed1]}

    // 残りのエントリを分配
    for i, entry := range node.Entries {
        if i == seed1 || i == seed2 {
            continue
        }

        if rt.assignToGroup(entry, node, newNode) {
            node.Entries = append(node.Entries, entry)
        } else {
            newNode.Entries = append(newNode.Entries, entry)
        }
    }

    // MBRを再計算
    node.MBR = rt.calculateNodeMBR(node)
    newNode.MBR = rt.calculateNodeMBR(newNode)

    // 親ノードの処理
    if node.Parent == nil {
        // ルートノードの分割
        rt.createNewRoot(node, newNode)
    } else {
        // 親ノードに新しいノードを追加
        rt.addToParent(newNode)
    }
}

// ヘルパー関数群

func intersects(mbr1, mbr2 Rectangle) bool {
    return mbr1.MinX <= mbr2.MaxX && mbr1.MaxX >= mbr2.MinX &&
           mbr1.MinY <= mbr2.MaxY && mbr1.MaxY >= mbr2.MinY
}

func expandMBR(mbr1, mbr2 Rectangle) Rectangle {
    return Rectangle{
        MinX: math.Min(mbr1.MinX, mbr2.MinX),
        MinY: math.Min(mbr1.MinY, mbr2.MinY),
        MaxX: math.Max(mbr1.MaxX, mbr2.MaxX),
        MaxY: math.Max(mbr1.MaxY, mbr2.MaxY),
    }
}

func calculateExpansion(current, new Rectangle) float64 {
    expanded := expandMBR(current, new)
    currentArea := (current.MaxX - current.MinX) * (current.MaxY - current.MinY)
    expandedArea := (expanded.MaxX - expanded.MinX) * (expanded.MaxY - expanded.MinY)
    return expandedArea - currentArea
}

func (rt *RTree) calculateNodeMBR(node *RTreeNode) Rectangle {
    if len(node.Entries) == 0 {
        return Rectangle{}
    }

    mbr := node.Entries[0].MBR
    for i := 1; i < len(node.Entries); i++ {
        mbr = expandMBR(mbr, node.Entries[i].MBR)
    }
    return mbr
}

// 分割用のシード選択（簡略化版）
func (rt *RTree) pickSeeds(entries []Entry) (int, int) {
    maxSeparation := -1.0
    seed1, seed2 := 0, 1

    for i := 0; i < len(entries); i++ {
        for j := i + 1; j < len(entries); j++ {
            separation := rt.calculateSeparation(entries[i].MBR, entries[j].MBR)
            if separation > maxSeparation {
                maxSeparation = separation
                seed1, seed2 = i, j
            }
        }
    }

    return seed1, seed2
}

func (rt *RTree) calculateSeparation(mbr1, mbr2 Rectangle) float64 {
    xSeparation := math.Abs((mbr1.MinX + mbr1.MaxX) - (mbr2.MinX + mbr2.MaxX))
    ySeparation := math.Abs((mbr1.MinY + mbr1.MaxY) - (mbr2.MinY + mbr2.MaxY))
    return xSeparation + ySeparation
}

func (rt *RTree) assignToGroup(entry Entry, group1, group2 *RTreeNode) bool {
    expansion1 := calculateExpansion(group1.MBR, entry.MBR)
    expansion2 := calculateExpansion(group2.MBR, entry.MBR)
    return expansion1 <= expansion2
}

func (rt *RTree) createNewRoot(node1, node2 *RTreeNode) {
    newRoot := &RTreeNode{
        IsLeaf: false,
        Entries: []Entry{
            {MBR: node1.MBR, Child: node1},
            {MBR: node2.MBR, Child: node2},
        },
    }
    newRoot.MBR = expandMBR(node1.MBR, node2.MBR)

    node1.Parent = newRoot
    node2.Parent = newRoot
    rt.Root = newRoot
}

func (rt *RTree) addToParent(newNode *RTreeNode) {
    parent := newNode.Parent
    entry := Entry{MBR: newNode.MBR, Child: newNode}
    parent.Entries = append(parent.Entries, entry)

    if len(parent.Entries) > rt.MaxEntries {
        rt.splitNode(parent)
    }
}
```

## より高度な実装（R\*-tree 要素）

```go
// R*-treeの再挿入処理
func (rt *RTree) reinsert(node *RTreeNode, level int) {
    if len(node.Entries) <= rt.MaxEntries {
        return
    }

    // 中心からの距離でソート
    entries := make([]Entry, len(node.Entries))
    copy(entries, node.Entries)

    center := rt.calculateCenter(node.MBR)
    sort.Slice(entries, func(i, j int) bool {
        dist1 := rt.distanceFromCenter(entries[i].MBR, center)
        dist2 := rt.distanceFromCenter(entries[j].MBR, center)
        return dist1 > dist2
    })

    // 30%のエントリを再挿入
    reinsertCount := len(entries) * 3 / 10
    if reinsertCount < 1 {
        reinsertCount = 1
    }

    // 再挿入するエントリを分離
    toReinsert := entries[:reinsertCount]
    node.Entries = entries[reinsertCount:]

    // MBRを再計算
    node.MBR = rt.calculateNodeMBR(node)

    // エントリを再挿入
    for _, entry := range toReinsert {
        if node.IsLeaf {
            rt.Insert(entry.Object)
        } else {
            // 内部ノードの場合の処理
            rt.reinsertInternal(entry, level)
        }
    }
}

type Point struct {
    X, Y float64
}

func (rt *RTree) calculateCenter(mbr Rectangle) Point {
    return Point{
        X: (mbr.MinX + mbr.MaxX) / 2,
        Y: (mbr.MinY + mbr.MaxY) / 2,
    }
}

func (rt *RTree) distanceFromCenter(mbr Rectangle, center Point) float64 {
    mbrCenter := rt.calculateCenter(mbr)
    dx := mbrCenter.X - center.X
    dy := mbrCenter.Y - center.Y
    return math.Sqrt(dx*dx + dy*dy)
}
```

# R-tree の応用例

## 1. 地理情報システム（GIS）

- **地図検索**: 指定領域内の施設・店舗検索
- **ナビゲーション**: 経路探索での空間インデックス
- **都市計画**: 土地利用の空間分析
- **環境監視**: センサーデータの空間集約

## 2. ゲーム開発

- **衝突検出**: キャラクターとオブジェクトの衝突判定
- **視野計算**: プレイヤーの視野内オブジェクト検索
- **空間分割**: 3D ゲーム世界の効率的管理
- **AI 経路探索**: NPC の移動経路計算

## 3. CAD・設計システム

- **図形検索**: 特定領域内の図形要素検索
- **干渉チェック**: 部品間の干渉検出
- **レイヤー管理**: 設計レイヤーの空間インデックス
- **寸法計算**: 最近接要素の検索

## 4. 画像・動画処理

- **オブジェクト検出**: 画像内のオブジェクト領域管理
- **顔認識**: 顔領域の空間インデックス
- **動き追跡**: 動画での物体追跡
- **画像検索**: 類似画像の空間特徴検索

## 5. 科学・工学シミュレーション

- **粒子シミュレーション**: 粒子間の相互作用計算
- **有限要素法**: メッシュ要素の空間管理
- **流体解析**: 計算セルの空間インデックス
- **天体計算**: 天体位置の空間検索

# R-tree と他のデータ構造との比較

## 空間データ構造比較

| 特徴           | R-tree          | Quadtree        | KD-tree         | Grid Index    | B-tree      |
| -------------- | --------------- | --------------- | --------------- | ------------- | ----------- |
| **空間検索**   | ✅ 最適         | ✅ 効率的       | ⚠️ 点データ向け | ⚠️ 固定分割   | ❌ 不適     |
| **範囲検索**   | ✅ O(log n + k) | ✅ O(log n + k) | ⚠️ O(√n + k)    | ✅ O(1 + k)   | ❌ O(n)     |
| **挿入・削除** | ✅ O(log n)     | ✅ O(log n)     | ❌ O(n)         | ✅ O(1)       | ✅ O(log n) |
| **動的更新**   | ✅ 効率的       | ✅ 効率的       | ❌ 再構築必要   | ✅ 容易       | ✅ 効率的   |
| **メモリ効率** | ✅ 良好         | ⚠️ 空間に依存   | ✅ 良好         | ❌ 大量メモリ | ✅ 最高     |
| **矩形データ** | ✅ 最適         | ⚠️ 近似的       | ❌ 不適         | ⚠️ 近似的     | ❌ 不適     |
| **実装複雑度** | ❌ 複雑         | ⚠️ 中程度       | ✅ 比較的簡単   | ✅ 簡単       | ⚠️ 中程度   |

## 用途別最適選択

| 用途                | 第 1 選択  | 第 2 選択 | 第 3 選択  | 避けるべき |
| ------------------- | ---------- | --------- | ---------- | ---------- |
| **GIS・地図データ** | R-tree     | Quadtree  | Grid Index | KD-tree    |
| **ゲーム衝突検出**  | R-tree     | Quadtree  | -          | B-tree     |
| **CAD 設計**        | R-tree     | -         | -          | Grid Index |
| **画像処理**        | R-tree     | Quadtree  | Grid Index | KD-tree    |
| **点データ検索**    | KD-tree    | R-tree    | Quadtree   | Grid Index |
| **固定領域検索**    | Grid Index | Quadtree  | R-tree     | KD-tree    |

# まとめ

R-tree は、多次元空間データの効率的な管理と検索を実現する専門的なデータ構造です。特に矩形領域を持つ空間オブジェクトの管理において優れた性能を発揮し、GIS、ゲーム開発、CAD、画像処理などの分野で広く利用されています。

実装は複雑ですが、空間検索・範囲検索の高速化という明確な利点があり、適切な分割アルゴリズムの選択により、アプリケーションの要件に応じた最適化が可能です。Go 言語での実装においても、構造体とポインタを活用することで効率的な R-tree を構築できます。
