# グラフ（Graph）とは

グラフは、ノード（頂点、Vertex）とエッジ（辺、Edge）によって構成される非線形データ構造です。ノード間の関係や接続を表現し、ネットワーク、経路、関係性などの複雑なデータを効率的に管理できます。

**このデータ構造を使う一番の利点は、複雑な関係性やネットワーク構造を自然に表現し、最短経路探索やネットワーク分析などの高度なアルゴリズムを効率的に実行できる点です。**

木構造とは異なり、任意のノード間に辺を持つことができ、循環や複数の経路を持つ構造を表現可能です。

# グラフの構成要素

## 基本要素

- **ノード（頂点、Vertex）**: グラフの基本単位となる点
- **エッジ（辺、Edge）**: ノード間の接続を表す線
- **重み（Weight）**: エッジに付与される数値（距離、コスト等）
- **次数（Degree）**: あるノードに接続されているエッジの数

## エッジの種類

- **有向エッジ**: 方向性を持つ辺（A → B）
- **無向エッジ**: 方向性を持たない辺（A ↔ B）
- **重み付きエッジ**: 数値が関連付けられた辺
- **重みなしエッジ**: 接続の有無のみを表す辺

# グラフの種類

## 方向性による分類

### 無向グラフ（Undirected Graph）

- エッジに方向性がない
- A-B 間の接続は双方向
- 用途: ソーシャルネットワーク、道路網

### 有向グラフ（Directed Graph, Digraph）

- エッジに方向性がある
- A→B の接続があっても B→A は存在しない場合がある
- 用途: Web ページのリンク、タスクの依存関係

## 重みによる分類

### 重み付きグラフ（Weighted Graph）

- エッジに重み（数値）が関連付けられている
- 用途: 最短経路問題、コスト最適化

### 重みなしグラフ（Unweighted Graph）

- エッジの重みは一定（通常 1）
- 用途: 単純な接続関係の表現

## 接続性による分類

### 連結グラフ（Connected Graph）

- すべてのノードが他のノードと経路で結ばれている
- 到達不可能なノードが存在しない

### 非連結グラフ（Disconnected Graph）

- 一部のノードが他のノードと接続されていない
- 複数の連結成分に分かれている

## 特殊なグラフ

### 完全グラフ（Complete Graph）

- すべてのノード間にエッジが存在
- n 個のノードを持つ完全グラフは n(n-1)/2 本のエッジを持つ

### 二部グラフ（Bipartite Graph）

- ノードを 2 つのグループに分けることができ、同一グループ内でエッジが存在しない
- 用途: マッチング問題、割り当て問題

### サイクリックグラフ（Cyclic Graph）

- 閉路（サイクル）を含むグラフ

### 非サイクリックグラフ（Acyclic Graph）

- 閉路を含まないグラフ
- 有向非サイクリックグラフ（DAG: Directed Acyclic Graph）は特に重要

# グラフの構成要素の図解

## 基本的なグラフの例

```
        A ――――――5―――――→ B
        │                 │
        │3                │2
        ↓                 ↓
        C ←――――4――――――― D
         ∖               ↗
          ∖1           ↗6
           ∖         ↗
            ↘     ↗
              E
```

## 各要素の説明

### ノード（頂点、Vertex）

- **A, B, C, D, E**: 円や点で表される各要素
- データを格納する基本単位
- 例: 都市、人、Web ページ、タスクなど

### エッジ（辺、Edge）

- **A→B, A→C, B→D, D→C, C→E, E→D**: ノード間を結ぶ線
- ノード間の関係や接続を表現
- 例: 道路、友人関係、リンク、依存関係など

### 重み（Weight）

- **数字 1, 2, 3, 4, 5, 6**: エッジに付与される数値
- 距離、コスト、時間、強度などを表現
- 重みなしグラフでは通常 1 または存在のみ

### エッジの方向性

- **→**: 有向エッジ（方向性あり）
- **―**: 無向エッジ（方向性なし、双方向）

### 次数（Degree）

各ノードの次数は以下の通り：

- **A**: 出次数 2（A→B, A→C）
- **B**: 入次数 1、出次数 1（A→B, B→D）
- **C**: 入次数 2、出次数 1（A→C, D→C, C→E）
- **D**: 入次数 2、出次数 1（B→D, E→D, D→C）
- **E**: 入次数 1、出次数 1（C→E, E→D）

## 具体例での説明

### 道路網の例

```
    東京 ――――100km―――――→ 横浜
     │                    │
     │80km               │50km
     ↓                    ↓
    千葉 ←――――70km――――― 埼玉
      ∖                 ↗
       ∖40km          ↗60km
        ∖            ↗
         ↘         ↗
          茨城
```

- **ノード**: 都市（東京、横浜、千葉、埼玉、茨城）
- **エッジ**: 道路・経路
- **重み**: 距離（km）
- **方向**: 一方通行（→）または双方向（↔）

### ソーシャルネットワークの例

```
    Alice ←――――――――――→ Bob
      │                   │
      │フォロー           │フォロー
      ↓                   ↓
    Carol ←――友達――――→ David
       ∖                 ↗
        ∖いいね        ↗メンション
         ∖            ↗
          ↘         ↗
           Eve
```

- **ノード**: ユーザー（Alice、Bob、Carol、David、Eve）
- **エッジ**: 関係性（フォロー、友達、いいね、メンション）
- **重み**: 関係の強度（省略可能）
- **方向**: フォローは一方向、友達は双方向

## グラフの種類による表現の違い

### 無向グラフ

```
    A ――――― B
    │       │
    │       │
    C ――――― D
```

すべてのエッジが双方向（―で表現）

### 有向グラフ

```
    A ――→ B
    ↑     ↓
    │     │
    C ←―― D
```

すべてのエッジが一方向（→ で表現）

### 重み付きグラフ

```
    A ――5―― B
    │3     2│
    │       │
    C ――4―― D
```

各エッジに重み（数値）が付与

### 完全グラフ（4 ノード）

```
    A ―――――― B
    │∖     ∕│
    │ ∖   ∕ │
    │  ∖ ∕  │
    │   ∕∖  │
    │  ∕  ∖ │
    │ ∕    ∖│
    │∕      ∖
    C ―――――― D
```

すべてのノード間にエッジが存在

## 特殊な構造

### 木構造（グラフの特殊形）

```
      A
     ∕ ∖
    B   C
   ∕   ∕ ∖
  D   E   F
```

- サイクルなし
- 連結
- n 個のノードに n-1 個のエッジ

### サイクルを含むグラフ

```
    A ――→ B
    ↑     ↓
    │     │
    D ←―― C
```

A→B→C→D→A の閉路（サイクル）が存在

### 二部グラフ

```
  グループ1    グループ2
     A ――――――― X
     │ ∖     ∕ │
     │  ∖   ∕  │
     B ――∖∕―――― Y
          ∕∖
     C ――∕  ∖―― Z
```

ノードを 2 つのグループに分割でき、同一グループ内でエッジなし

# 基本的な特徴

- **任意接続**: ノード間に任意の接続関係を表現可能
- **非階層構造**: 木構造と異なり、明確な階層が存在しない場合がある
- **循環許可**: サイクルを含む構造を表現可能
- **複数経路**: 2 点間に複数の経路が存在可能
- **用途**: ネットワーク分析、経路探索、関係性解析、依存関係管理など

# 行える処理

## グラフ（Graph）の基本操作

| 機能         | 説明                   | 計算量      | 戻り値              | 得意/苦手 | 補足（その他）                         |
| ------------ | ---------------------- | ----------- | ------------------- | --------- | -------------------------------------- |
| AddVertex    | 新しいノードを追加     | O(1)        | bool（成功/失敗）   | ✅        | 隣接リスト・隣接行列ともに効率的       |
| AddEdge      | ノード間にエッジを追加 | O(1)        | bool（成功/失敗）   | ✅        | 重み付きエッジにも対応                 |
| RemoveVertex | ノードを削除           | O(V+E)      | bool（成功/失敗）   | ❌        | 関連するすべてのエッジも削除が必要     |
| RemoveEdge   | エッジを削除           | O(1)～ O(V) | bool（成功/失敗）   | ⚠️        | 実装方式により計算量が異なる           |
| HasVertex    | ノードの存在確認       | O(1)        | bool（存在/非存在） | ✅        | ハッシュテーブル使用時                 |
| HasEdge      | エッジの存在確認       | O(1)～ O(V) | bool（存在/非存在） | ⚠️        | 隣接行列なら O(1)、隣接リストなら O(V) |
| GetVertices  | すべてのノードを取得   | O(V)        | ノードの配列        | ✅        | ノード数に比例                         |
| GetEdges     | すべてのエッジを取得   | O(E)        | エッジの配列        | ✅        | エッジ数に比例                         |
| GetNeighbors | 隣接ノードを取得       | O(1)～ O(V) | 隣接ノードの配列    | ✅        | 隣接リストなら効率的                   |
| GetDegree    | ノードの次数を取得     | O(1)～ O(V) | int（次数）         | ⚠️        | 実装方式により異なる                   |
| IsEmpty      | グラフが空かチェック   | O(1)        | bool（空/非空）     | ✅        | ノード数の確認                         |
| Clear        | グラフの全要素を削除   | O(V+E)      | なし（void）        | ✅        | 全ノード・エッジのメモリ解放           |

## グラフ探索・走査アルゴリズム

| 機能                        | 説明               | 計算量 | 戻り値                | 得意/苦手 | 補足（その他）                         |
| --------------------------- | ------------------ | ------ | --------------------- | --------- | -------------------------------------- |
| BFS                         | 幅優先探索         | O(V+E) | 訪問順序の配列        | ✅        | 最短経路探索に適している               |
| DFS                         | 深度優先探索       | O(V+E) | 訪問順序の配列        | ✅        | トポロジカルソート、サイクル検出に適用 |
| TopologicalSort             | トポロジカルソート | O(V+E) | ソート済み配列        | ✅        | DAG でのみ適用可能                     |
| StronglyConnectedComponents | 強連結成分の検出   | O(V+E) | 成分ごとのノード配列  | ⚠️        | 有向グラフでのみ適用                   |
| ConnectedComponents         | 連結成分の検出     | O(V+E) | 成分ごとのノード配列  | ✅        | 無向グラフでの島の検出                 |
| HasCycle                    | サイクルの検出     | O(V+E) | bool（サイクル有/無） | ✅        | DFS ベースの実装                       |
| IsBipartite                 | 二部グラフ判定     | O(V+E) | bool（二部/非二部）   | ✅        | BFS または DFS で色分け                |

## 最短経路・最適化アルゴリズム

| 機能                | 説明                           | 計算量       | 戻り値         | 得意/苦手 | 補足（その他）               |
| ------------------- | ------------------------------ | ------------ | -------------- | --------- | ---------------------------- |
| Dijkstra            | 単一始点最短経路               | O((V+E)logV) | 距離・経路配列 | ✅        | 非負重みグラフで最適         |
| BellmanFord         | 単一始点最短経路（負重み対応） | O(VE)        | 距離・経路配列 | ⚠️        | 負の閉路検出可能だが低速     |
| FloydWarshall       | 全点対最短経路                 | O(V³)        | 距離行列       | ❌        | 密なグラフでは非効率         |
| AStar               | A\*探索（ヒューリスティック）  | O(b^d)       | 最短経路       | ✅        | ゲーム・ナビで高性能         |
| MinimumSpanningTree | 最小全域木（Kruskal/Prim）     | O(ElogE)     | 木の辺集合     | ✅        | ネットワーク設計に適用       |
| MaxFlow             | 最大流問題                     | O(VE²)       | 最大流量       | ❌        | 複雑なアルゴリズム実装が必要 |

## 特殊操作・解析

| 機能          | 説明                       | 計算量     | 戻り値              | 得意/苦手 | 補足（その他）                 |
| ------------- | -------------------------- | ---------- | ------------------- | --------- | ------------------------------ |
| Clone         | グラフの完全複製           | O(V+E)     | 新しいグラフ        | ✅        | 全ノード・エッジを複製         |
| Transpose     | グラフの転置（有向グラフ） | O(V+E)     | 転置グラフ          | ✅        | エッジの方向を反転             |
| Complement    | 補グラフの生成             | O(V²)      | 補グラフ            | ❌        | 存在しないエッジを追加         |
| IsIsomorphic  | グラフ同型判定             | O(V!)      | bool（同型/非同型） | ❌        | NP 完全問題、実用的でない      |
| FindCliques   | クリークの検出             | O(3^(V/3)) | クリーク配列        | ❌        | 指数時間、近似アルゴリズム推奨 |
| ColoringGraph | グラフ彩色                 | O(V²)      | 色割り当て          | ⚠️        | 最適彩色は困難、貪欲法で近似   |

# 注意点

1. **メモリ使用量**: 密なグラフでは隣接行列、疎なグラフでは隣接リストが効率的
2. **サイクル処理**: 無限ループを避けるため訪問済みノードの管理が必要
3. **重み付きエッジ**: 負の重みを含む場合、一部のアルゴリズムが適用不可
4. **有向・無向の区別**: アルゴリズムによって適用可能なグラフ種類が限定される
5. **スケーラビリティ**: ノード数・エッジ数の増加に対するアルゴリズム選択が重要

# 実装方法

## 実装方式別の比較

| 実装方式         | メリット                                                             | デメリット                                                             | 適用場面                 | 推奨実装             |
| ---------------- | -------------------------------------------------------------------- | ---------------------------------------------------------------------- | ------------------------ | -------------------- |
| **隣接行列**     | - エッジ存在確認が O(1)<br>- 実装が簡単<br>- 密なグラフで効率的      | - メモリ使用量が O(V²)<br>- 疎なグラフで無駄<br>- ノード追加でリサイズ | 密なグラフ、小規模グラフ | 2 次元配列           |
| **隣接リスト**   | - メモリ効率が良い<br>- 疎なグラフで高効率<br>- 動的サイズ変更が容易 | - エッジ存在確認が O(V)<br>- 実装がやや複雑<br>- キャッシュ効率劣る    | 疎なグラフ、大規模グラフ | map+slice            |
| **エッジリスト** | - シンプルな構造<br>- 全エッジ操作が効率的<br>- メモリ効率良好       | - 隣接ノード取得が O(E)<br>- エッジ存在確認が O(E)<br>- 探索に不適     | エッジ中心の処理         | エッジ構造体の slice |

## 隣接行列実装例

```go
type AdjacencyMatrixGraph struct {
    matrix    [][]int  // 隣接行列
    vertices  []string // ノードのラベル
    nodeMap   map[string]int // ラベル→インデックスマッピング
    nodeCount int
    directed  bool
}

func NewAdjacencyMatrixGraph(maxNodes int, directed bool) *AdjacencyMatrixGraph {
    matrix := make([][]int, maxNodes)
    for i := range matrix {
        matrix[i] = make([]int, maxNodes)
    }

    return &AdjacencyMatrixGraph{
        matrix:   matrix,
        vertices: make([]string, 0, maxNodes),
        nodeMap:  make(map[string]int),
        directed: directed,
    }
}

func (g *AdjacencyMatrixGraph) AddVertex(label string) bool {
    if _, exists := g.nodeMap[label]; exists {
        return false // 既に存在
    }

    if g.nodeCount >= len(g.matrix) {
        return false // 容量不足
    }

    g.nodeMap[label] = g.nodeCount
    g.vertices = append(g.vertices, label)
    g.nodeCount++
    return true
}

func (g *AdjacencyMatrixGraph) AddEdge(from, to string, weight int) bool {
    fromIdx, fromExists := g.nodeMap[from]
    toIdx, toExists := g.nodeMap[to]

    if !fromExists || !toExists {
        return false
    }

    g.matrix[fromIdx][toIdx] = weight
    if !g.directed {
        g.matrix[toIdx][fromIdx] = weight
    }
    return true
}

func (g *AdjacencyMatrixGraph) HasEdge(from, to string) bool {
    fromIdx, fromExists := g.nodeMap[from]
    toIdx, toExists := g.nodeMap[to]

    if !fromExists || !toExists {
        return false
    }

    return g.matrix[fromIdx][toIdx] != 0
}

func (g *AdjacencyMatrixGraph) GetNeighbors(vertex string) []string {
    idx, exists := g.nodeMap[vertex]
    if !exists {
        return nil
    }

    var neighbors []string
    for i := 0; i < g.nodeCount; i++ {
        if g.matrix[idx][i] != 0 {
            neighbors = append(neighbors, g.vertices[i])
        }
    }
    return neighbors
}

// 幅優先探索
func (g *AdjacencyMatrixGraph) BFS(start string) []string {
    startIdx, exists := g.nodeMap[start]
    if !exists {
        return nil
    }

    visited := make([]bool, g.nodeCount)
    queue := []int{startIdx}
    result := []string{}

    visited[startIdx] = true

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        result = append(result, g.vertices[current])

        for i := 0; i < g.nodeCount; i++ {
            if g.matrix[current][i] != 0 && !visited[i] {
                visited[i] = true
                queue = append(queue, i)
            }
        }
    }

    return result
}
```

### 特徴

- **高速エッジ確認**: O(1)でエッジの存在確認が可能
- **簡単実装**: 2 次元配列のみで実装可能
- **密グラフ最適**: エッジ数が多い場合にメモリ効率が良い
- **固定サイズ**: 事前にノード数の上限を決定する必要

### 注意事項

- **メモリ消費**: O(V²)のメモリが常に必要
- **疎グラフ非効率**: エッジが少ない場合、多くのメモリが無駄
- **動的拡張困難**: ノード数の動的増加が困難
- **大規模グラフ不適**: ノード数が多いとメモリ不足

## 隣接リスト実装例

```go
type Edge struct {
    To     string
    Weight int
}

type AdjacencyListGraph struct {
    adjList  map[string][]Edge
    directed bool
}

func NewAdjacencyListGraph(directed bool) *AdjacencyListGraph {
    return &AdjacencyListGraph{
        adjList:  make(map[string][]Edge),
        directed: directed,
    }
}

func (g *AdjacencyListGraph) AddVertex(vertex string) bool {
    if _, exists := g.adjList[vertex]; exists {
        return false // 既に存在
    }

    g.adjList[vertex] = make([]Edge, 0)
    return true
}

func (g *AdjacencyListGraph) AddEdge(from, to string, weight int) bool {
    // ノードが存在しない場合は作成
    if _, exists := g.adjList[from]; !exists {
        g.AddVertex(from)
    }
    if _, exists := g.adjList[to]; !exists {
        g.AddVertex(to)
    }

    // エッジを追加
    g.adjList[from] = append(g.adjList[from], Edge{To: to, Weight: weight})

    if !g.directed {
        g.adjList[to] = append(g.adjList[to], Edge{To: from, Weight: weight})
    }

    return true
}

func (g *AdjacencyListGraph) RemoveEdge(from, to string) bool {
    fromList, fromExists := g.adjList[from]
    if !fromExists {
        return false
    }

    // fromからtoへのエッジを削除
    for i, edge := range fromList {
        if edge.To == to {
            g.adjList[from] = append(fromList[:i], fromList[i+1:]...)
            break
        }
    }

    // 無向グラフの場合、逆方向も削除
    if !g.directed {
        toList, toExists := g.adjList[to]
        if toExists {
            for i, edge := range toList {
                if edge.To == from {
                    g.adjList[to] = append(toList[:i], toList[i+1:]...)
                    break
                }
            }
        }
    }

    return true
}

func (g *AdjacencyListGraph) HasEdge(from, to string) bool {
    fromList, exists := g.adjList[from]
    if !exists {
        return false
    }

    for _, edge := range fromList {
        if edge.To == to {
            return true
        }
    }
    return false
}

func (g *AdjacencyListGraph) GetNeighbors(vertex string) []string {
    edges, exists := g.adjList[vertex]
    if !exists {
        return nil
    }

    neighbors := make([]string, len(edges))
    for i, edge := range edges {
        neighbors[i] = edge.To
    }
    return neighbors
}

// 深度優先探索
func (g *AdjacencyListGraph) DFS(start string) []string {
    visited := make(map[string]bool)
    result := []string{}

    g.dfsRecursive(start, visited, &result)
    return result
}

func (g *AdjacencyListGraph) dfsRecursive(vertex string, visited map[string]bool, result *[]string) {
    if visited[vertex] {
        return
    }

    visited[vertex] = true
    *result = append(*result, vertex)

    for _, edge := range g.adjList[vertex] {
        if !visited[edge.To] {
            g.dfsRecursive(edge.To, visited, result)
        }
    }
}

// Dijkstra最短経路アルゴリズム
func (g *AdjacencyListGraph) Dijkstra(start string) map[string]int {
    distances := make(map[string]int)
    visited := make(map[string]bool)

    // 全ノードの距離を無限大で初期化
    for vertex := range g.adjList {
        distances[vertex] = math.MaxInt32
    }
    distances[start] = 0

    // 優先度付きキュー（簡易実装）
    for len(visited) < len(g.adjList) {
        // 未訪問の最小距離ノードを選択
        current := ""
        minDistance := math.MaxInt32

        for vertex, distance := range distances {
            if !visited[vertex] && distance < minDistance {
                current = vertex
                minDistance = distance
            }
        }

        if current == "" {
            break // 到達不可能なノードが残っている
        }

        visited[current] = true

        // 隣接ノードの距離を更新
        for _, edge := range g.adjList[current] {
            if !visited[edge.To] {
                newDistance := distances[current] + edge.Weight
                if newDistance < distances[edge.To] {
                    distances[edge.To] = newDistance
                }
            }
        }
    }

    return distances
}

// トポロジカルソート（DFS版）
func (g *AdjacencyListGraph) TopologicalSort() []string {
    if !g.directed {
        return nil // 有向グラフでのみ適用可能
    }

    visited := make(map[string]bool)
    stack := []string{}

    for vertex := range g.adjList {
        if !visited[vertex] {
            g.topologicalSortDFS(vertex, visited, &stack)
        }
    }

    // スタックを逆順にして返す
    result := make([]string, len(stack))
    for i, vertex := range stack {
        result[len(stack)-1-i] = vertex
    }

    return result
}

func (g *AdjacencyListGraph) topologicalSortDFS(vertex string, visited map[string]bool, stack *[]string) {
    visited[vertex] = true

    for _, edge := range g.adjList[vertex] {
        if !visited[edge.To] {
            g.topologicalSortDFS(edge.To, visited, stack)
        }
    }

    *stack = append(*stack, vertex)
}
```

### 特徴

- **メモリ効率**: 実際に存在するエッジのみメモリを使用
- **動的サイズ**: ノードとエッジの動的追加・削除が容易
- **疎グラフ最適**: エッジ数が少ない場合に威力を発揮
- **隣接ノード高速取得**: 特定ノードの隣接ノードを効率的に取得

### 注意事項

- **エッジ確認コスト**: エッジの存在確認に O(V)時間が必要
- **キャッシュ効率**: 非連続メモリアクセスでキャッシュミス発生
- **実装複雑性**: 隣接行列より実装が複雑
- **メモリ断片化**: 多数の小さなスライスによる断片化

## エッジリスト実装例

```go
type GraphEdge struct {
    From   string
    To     string
    Weight int
}

type EdgeListGraph struct {
    edges    []GraphEdge
    vertices map[string]bool
    directed bool
}

func NewEdgeListGraph(directed bool) *EdgeListGraph {
    return &EdgeListGraph{
        edges:    make([]GraphEdge, 0),
        vertices: make(map[string]bool),
        directed: directed,
    }
}

func (g *EdgeListGraph) AddVertex(vertex string) bool {
    if g.vertices[vertex] {
        return false
    }
    g.vertices[vertex] = true
    return true
}

func (g *EdgeListGraph) AddEdge(from, to string, weight int) bool {
    // ノードを自動追加
    g.vertices[from] = true
    g.vertices[to] = true

    edge := GraphEdge{From: from, To: to, Weight: weight}
    g.edges = append(g.edges, edge)

    if !g.directed {
        reverseEdge := GraphEdge{From: to, To: from, Weight: weight}
        g.edges = append(g.edges, reverseEdge)
    }

    return true
}

func (g *EdgeListGraph) GetEdges() []GraphEdge {
    return g.edges
}

func (g *EdgeListGraph) GetVertices() []string {
    vertices := make([]string, 0, len(g.vertices))
    for vertex := range g.vertices {
        vertices = append(vertices, vertex)
    }
    return vertices
}

// Kruskal最小全域木アルゴリズム
func (g *EdgeListGraph) KruskalMST() []GraphEdge {
    if g.directed {
        return nil // 無向グラフでのみ適用可能
    }

    // エッジを重みでソート
    edges := make([]GraphEdge, len(g.edges))
    copy(edges, g.edges)

    sort.Slice(edges, func(i, j int) bool {
        return edges[i].Weight < edges[j].Weight
    })

    // Union-Find構造
    parent := make(map[string]string)
    rank := make(map[string]int)

    // 初期化
    for vertex := range g.vertices {
        parent[vertex] = vertex
        rank[vertex] = 0
    }

    // Find操作
    var find func(string) string
    find = func(x string) string {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    // Union操作
    union := func(x, y string) bool {
        rootX := find(x)
        rootY := find(y)

        if rootX == rootY {
            return false // すでに同じ集合
        }

        if rank[rootX] < rank[rootY] {
            parent[rootX] = rootY
        } else if rank[rootX] > rank[rootY] {
            parent[rootY] = rootX
        } else {
            parent[rootY] = rootX
            rank[rootX]++
        }
        return true
    }

    mst := []GraphEdge{}

    for _, edge := range edges {
        if union(edge.From, edge.To) {
            mst = append(mst, edge)
            if len(mst) == len(g.vertices)-1 {
                break // 最小全域木完成
            }
        }
    }

    return mst
}
```

### 特徴

- **シンプル構造**: エッジの配列のみで構成
- **エッジ中心処理**: エッジ操作が中心の場合に効率的
- **ソート対応**: エッジの重みでのソートが容易
- **全域木アルゴリズム**: Kruskal アルゴリズムに最適

### 注意事項

- **隣接ノード取得コスト**: O(E)の時間が必要
- **エッジ存在確認コスト**: 線形探索で O(E)
- **グラフ探索不適**: BFS/DFS には不向き
- **メモリ効率**: 隣接リストより多くのメモリを使用する場合あり

## 実装選択の指針

| 要件                 | 推奨実装     | 理由                           |
| -------------------- | ------------ | ------------------------------ |
| **密なグラフ**       | 隣接行列     | エッジ存在確認が O(1)で効率的  |
| **疎なグラフ**       | 隣接リスト   | メモリ効率が良く、実用的       |
| **動的構造変更**     | 隣接リスト   | ノード・エッジの追加削除が容易 |
| **最短経路探索**     | 隣接リスト   | Dijkstra、BFS に適している     |
| **最小全域木**       | エッジリスト | Kruskal アルゴリズムに最適     |
| **メモリ制約**       | 隣接リスト   | 実際のエッジのみメモリ使用     |
| **高速エッジ確認**   | 隣接行列     | O(1)での存在確認               |
| **大規模グラフ**     | 隣接リスト   | スケーラビリティが高い         |
| **全エッジ操作**     | エッジリスト | エッジ中心の処理に特化         |
| **小規模完全グラフ** | 隣接行列     | シンプルで効率的               |

# グラフの応用例

## 1. ネットワーク・通信

- **コンピュータネットワーク**: ルーター間の接続とデータ経路
- **ソーシャルネットワーク**: ユーザー間の友人関係・フォロー関係
- **通信プロトコル**: パケットルーティング、負荷分散
- **インターネット**: Web ページ間のリンク構造、PageRank アルゴリズム

## 2. 地理・交通システム

- **道路網**: 交差点とその間の道路、カーナビゲーション
- **公共交通**: 駅・停留所間の路線、乗り換え案内
- **航空路線**: 空港間のフライト接続、最適経路探索
- **物流ネットワーク**: 倉庫・配送センター間の配送経路

## 3. プロジェクト管理・スケジューリング

- **タスク依存関係**: プロジェクトのタスク間の前後関係
- **ガントチャート**: 作業の順序関係とクリティカルパス
- **リソース配分**: 限られたリソースの最適割り当て
- **工程管理**: 製造業の工程間の依存関係

## 4. データベース・知識管理

- **データベース設計**: テーブル間の関連性、外部キー関係
- **知識グラフ**: 概念間の関係性、セマンティック Web
- **推薦システム**: ユーザー・商品間の関係性分析
- **検索エンジン**: Web ページの関連性とランキング

## 5. ゲーム・シミュレーション

- **ゲームマップ**: ゲーム内の移動可能エリア
- **状態遷移**: ゲームの状態変化、AI の行動決定
- **パスファインディング**: NPC の移動経路探索
- **戦略シミュレーション**: ユニット間の影響関係

## 6. 生物学・化学

- **分子構造**: 原子間の結合、化学反応経路
- **遺伝子ネットワーク**: 遺伝子間の調節関係
- **タンパク質相互作用**: タンパク質間の結合関係
- **生態系**: 種間の捕食関係、食物連鎖

## 7. 金融・経済

- **取引ネットワーク**: 金融機関間の取引関係
- **リスク管理**: 投資商品間の相関関係
- **サプライチェーン**: 企業間の取引関係
- **経済指標**: 各種指標間の影響関係

# グラフ構造と探索アルゴリズム

探索アルゴリズムはこのグラフ上で特定の目的（経路探索、到達可能性の確認、最短経路の計算など）を達成するために使用されます。
グラフの構造（有向/無向、重み付き/重みなし、連結/非連結など）に応じて、適切な探索アルゴリズムを選択する必要があります。

## グラフ構造と探索アルゴリズムの選択指針

| 用途                                                             | 推奨アルゴリズム                             | 適用例                                                 | 計算量                                           |
| ---------------------------------------------------------------- | -------------------------------------------- | ------------------------------------------------------ | ------------------------------------------------ |
| 最短経路探索（重みなしグラフ）、到達可能性の確認                 | 幅優先探索（BFS: Breadth-First Search）      | ソーシャルネットワークでの友人関係の探索               | O(V+E)（V:ノード数、E:エッジ数）                 |
| サイクル検出、連結成分の確認、トポロジカルソート                 | 深度優先探索（DFS: Depth-First Search）      | タスクの依存関係の解析                                 | O(V+E)                                           |
| 単一始点最短経路探索（非負重みグラフ）                           | ダイクストラ法（Dijkstra's Algorithm）       | カーナビゲーションシステムでの最短経路探索             | O((V+E) logV)（ヒープを使用した場合）            |
| 単一始点最短経路探索（負重み対応）                               | ベルマンフォード法（Bellman-Ford Algorithm） | 金融システムでのリスク評価                             | O(VE)                                            |
| 全点対最短経路                                                   | フロイド・ワーシャル法                       | 都市間の最短経路計算                                   | O(V^3)                                           |
| 最小全域木（Minimum Spanning Tree, MST）の構築                   | プリム法、クラスカル法                       | ネットワーク設計での最小コスト接続、電力網の最適設計   | O((V+E) logV)（ヒープを使用した場合）, O(E logE) |
| 有向非巡回グラフ（DAG）の順序付け                                | トポロジカルソート                           | タスクの依存関係解析                                   | O(V+E)                                           |
| 有向グラフの強連結成分を検出                                     | 強連結成分分解（SCC）                        | ソーシャルネットワークでのコミュニティや依存関係の検出 | O(V+E)                                           |
| ゴールノードへの効率的な最短経路探索（ヒューリスティックを使用） | A\*アルゴリズム                              | ゲーム AI での経路探索                                 | O(b^d)（b:分岐数、d:深さ）                       |

# 他のデータ構造との関係

## 木構造との関係

- **木はグラフの特殊形**: サイクルのない連結グラフ
- **全域木**: グラフから木を抽出する概念
- **最小全域木**: 重み付きグラフから最小コストの木を構成
- **探索順序**: DFS/BFS で木構造が形成される

## 配列・リストとの関係

- **隣接行列**: 2 次元配列でグラフを表現
- **隣接リスト**: 配列と連結リストの組み合わせ
- **エッジリスト**: エッジを配列で管理
- **経路**: ノードの配列として表現

## ハッシュテーブルとの関係

- **ノード管理**: ノード ID と情報のマッピング
- **高速アクセス**: ノードの存在確認とアクセス
- **インデックス**: 隣接リストでのノード検索
- **キャッシュ**: 計算結果の一時保存

## キュー・スタックとの関係

- **BFS**: キューを使用した幅優先探索
- **DFS**: スタックまたは再帰を使用した深度優先探索
- **トポロジカルソート**: スタックで順序を管理
- **経路復元**: 親ノードをスタックで管理

# 他のデータ構造との比較

## グラフ実装方式別比較

| 特徴               | 隣接行列       | 隣接リスト     | エッジリスト | 木構造      |
| ------------------ | -------------- | -------------- | ------------ | ----------- |
| **メモリ使用量**   | ❌ O(V²)固定   | ✅ O(V+E)動的  | ⚠️ O(E)固定  | ✅ O(V)最小 |
| **エッジ存在確認** | ✅ O(1)最速    | ❌ O(V)線形    | ❌ O(E)線形  | ⚠️ O(log V) |
| **隣接ノード取得** | ⚠️ O(V)線形    | ✅ O(次数)最適 | ❌ O(E)線形  | ✅ O(1)直接 |
| **エッジ追加**     | ✅ O(1)        | ✅ O(1)        | ✅ O(1)      | ✅ O(1)     |
| **エッジ削除**     | ✅ O(1)        | ⚠️ O(V)        | ❌ O(E)      | ✅ O(1)     |
| **ノード追加**     | ❌ O(V²)再構築 | ✅ O(1)        | ✅ O(1)      | ✅ O(1)     |
| **密グラフ効率**   | ✅ 最適        | ❌ 非効率      | ❌ 非効率    | ❌ 適用不可 |
| **疎グラフ効率**   | ❌ 非効率      | ✅ 最適        | ✅ 良好      | ✅ 最適     |
| **実装複雑さ**     | ✅ シンプル    | ⚠️ 中程度      | ✅ シンプル  | ✅ シンプル |
| **キャッシュ効率** | ✅ 高い        | ❌ 低い        | ⚠️ 中程度    | ✅ 高い     |

## 全データ構造比較

| 特徴             | グラフ          | 木構造      | 配列          | ハッシュテーブル | 連結リスト        |
| ---------------- | --------------- | ----------- | ------------- | ---------------- | ----------------- |
| **任意関係表現** | ✅ 最適         | ⚠️ 階層のみ | ❌ 不適       | ⚠️ キー値のみ    | ❌ 線形のみ       |
| **検索性能**     | ⚠️ O(V+E)       | ✅ O(log V) | ✅ O(log n)   | ✅ O(1)平均      | ❌ O(n)           |
| **挿入性能**     | ✅ O(1)         | ✅ O(log V) | ❌ O(n)       | ✅ O(1)平均      | ✅ O(1)           |
| **削除性能**     | ⚠️ O(V+E)       | ✅ O(log V) | ❌ O(n)       | ✅ O(1)平均      | ✅ O(1)           |
| **メモリ効率**   | ⚠️ 実装依存     | ✅ 良好     | ✅ 最高       | ⚠️ 負荷率依存    | ⚠️ ポインタ分重い |
| **関係性分析**   | ✅ 最適         | ⚠️ 親子のみ | ❌ 不適       | ❌ 不適          | ❌ 不適           |
| **経路探索**     | ✅ 最適         | ⚠️ 単一経路 | ❌ 不適       | ❌ 不適          | ❌ 不適           |
| **ソート**       | ⚠️ トポロジカル | ✅ 中順走査 | ✅ 各種ソート | ❌ 不適          | ❌ 非効率         |
| **実装複雑さ**   | ❌ 複雑         | ⚠️ 中程度   | ✅ シンプル   | ⚠️ 中程度        | ✅ シンプル       |
| **並行処理**     | ❌ 困難         | ⚠️ 注意必要 | ✅ 比較的容易 | ⚠️ 同期必要      | ❌ 困難           |

# 使用場面別最適選択

| 使用場面                   | 第 1 選択          | 第 2 選択        | 第 3 選択 | 避けるべき       |
| -------------------------- | ------------------ | ---------------- | --------- | ---------------- |
| **ナビゲーションシステム** | 隣接リストグラフ   | 隣接行列グラフ   | -         | エッジリスト     |
| **ソーシャルネットワーク** | 隣接リストグラフ   | -                | -         | 隣接行列         |
| **Web クローラー**         | 隣接リストグラフ   | エッジリスト     | -         | 隣接行列         |
| **ネットワークトポロジー** | 隣接リストグラフ   | 隣接行列グラフ   | -         | エッジリスト     |
| **タスクスケジューリング** | 隣接リストグラフ   | -                | -         | 配列、連結リスト |
| **最小全域木**             | エッジリストグラフ | 隣接リストグラフ | -         | 隣接行列         |
| **ゲームのマップ**         | 隣接行列グラフ     | 隣接リストグラフ | -         | エッジリスト     |
| **データベース関係**       | 隣接リストグラフ   | -                | -         | 隣接行列         |
| **分子構造解析**           | 隣接リストグラフ   | エッジリスト     | -         | 隣接行列         |
| **推薦システム**           | 隣接リストグラフ   | ハッシュテーブル | -         | 配列             |

# パフォーマンス特性

| 操作パターン           | 隣接行列        | 隣接リスト     | エッジリスト  |
| ---------------------- | --------------- | -------------- | ------------- |
| **頻繁なエッジ確認**   | ✅ O(1)最適     | ❌ O(V)非効率  | ❌ O(E)非効率 |
| **大量ノード少エッジ** | ❌ メモリ無駄   | ✅ 最効率      | ✅ 効率的     |
| **動的構造変更**       | ❌ リサイズ必要 | ✅ 柔軟対応    | ✅ 柔軟対応   |
| **全エッジ走査**       | ❌ O(V²)        | ✅ O(E)        | ✅ O(E)最適   |
| **隣接ノード頻繁取得** | ❌ O(V)         | ✅ O(次数)最適 | ❌ O(E)       |
| **メモリ制約環境**     | ❌ 固定 O(V²)   | ✅ 動的 O(V+E) | ✅ 固定 O(E)  |
| **キャッシュ最適化**   | ✅ 連続メモリ   | ❌ 断片化      | ⚠️ 部分的     |
| **並行アクセス**       | ⚠️ 配列同期     | ❌ 複雑同期    | ⚠️ リスト同期 |

# まとめ

グラフは複雑な関係性を表現する最も汎用的なデータ構造で、ネットワーク分析から人工知能まで幅広い分野で活用されています。適切な実装方式と効率的なアルゴリズムの選択により、大規模で複雑な問題の解決が可能になります。
