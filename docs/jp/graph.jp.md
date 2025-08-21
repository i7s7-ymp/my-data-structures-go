# グラフ（Graph）とは

グラフは、ノード（頂点、Vertex）とエッジ（辺、Edge）によって構成される非線形データ構造。ノード間の関係や接続を表現し、ネットワーク、経路、関係性などの複雑なデータを効率的に管理できる。

**複雑な関係性やネットワーク構造を自然に表現し、最短経路探索やネットワーク分析などの高度なアルゴリズムを効率的に実行できる。**

木構造とは異なり、任意のノード間に辺を持つことができ、循環や複数の経路を持つ構造を表現することができる。

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

## グラフの種類

- **無向グラフ（Undirected Graph）**: エッジに方向性がない。あるエッジ間の接続は双方向。道路網などが該当する。
- **有向グラフ（Directed Graph, Digraph）**: エッジに方向性がある。Web ページのリンクなどが該当する。
- **重み付きグラフ（Weighted Graph）**: エッジに重み（数値）が関連付けられている。最短経路問題のグラフなどで利用する。
- **連結グラフ（Connected Graph）**: すべてのノードが他のノードと経路で結ばれており、到達不可能なノードが存在しない。
- **非連結グラフ（Disconnected Graph）**: 一部のノードが他のノードと接続されていない。複数の連結成分に分かれている。
- **完全グラフ（Complete Graph）**: すべてのノード間にエッジが存在、n 個のノードを持つ完全グラフは n(n-1)/2 本のエッジを持つ、
- **二部グラフ（Bipartite Graph）**: ノードを 2 つのグループに分けることができ、同一グループ内でエッジが存在しない。マッチング問題、割り当て問題などで利用される。
- **サイクリックグラフ（Cyclic Graph）**: 閉路（サイクル）を含むグラフ。
- **非サイクリックグラフ（Acyclic Graph）**: 閉路を含まないグラフ。有向非サイクリックグラフ（DAG: Directed Acyclic Graph）は特に重要。

## 具体例

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

# 実装方法

グラフを表現する方法は主に二種類ある。

## 隣接行列(Adjacency Matrix)

V 個の頂点を持つグラフを、VxV の二次元配列（行列）で表現する方法。行列の要素 M[i][j] にあった意が存在すれば、頂点 i から j へのエッジが存在し、0 なら存在しないことを表す。

- 長所: 2 つの頂点間のエッジ存在確認の実行時間は O(1) の定数時間となる。
- 短所: 頂点数 V に対して O(V\*\*2) のメモリ領域が必要。エッジ数が少ない疎なグラフ（sparse graph）の場合、行列のほとんどは 0 になり、メモリが無駄になる。

## 隣接リスト（Adjacency List）

頂点ごとに、その頂点に隣接する（エッジで繋がっている）頂点のリストを保持する方法。現実世界の多くのグラフ構造（ソーシャルネットワーク、Web ページのリンクなど）は疎であるため、隣接リストが一般的に用いられる。

- 長所: 必要なメモリ領域が O(V+E)（E はエッジ数）であり、疎なグラフに対して非常に空間効率が良い。
- 短所: 2 頂点間にエッジが存在するかを確認するには、一方の頂点リストを線形探索する必要があり、最悪 O(V)（頂点の字数に比例）の時間がかかる。

## 実装方式別の比較

| 実装方式         | メリット                                                             | デメリット                                                             | 適用場面                 | 推奨実装             |
| ---------------- | -------------------------------------------------------------------- | ---------------------------------------------------------------------- | ------------------------ | -------------------- |
| **隣接行列**     | - エッジ存在確認が O(1)<br>- 実装が簡単<br>- 密なグラフで効率的      | - メモリ使用量が O(V²)<br>- 疎なグラフで無駄<br>- ノード追加でリサイズ | 密なグラフ、小規模グラフ | 2 次元配列           |
| **隣接リスト**   | - メモリ効率が良い<br>- 疎なグラフで高効率<br>- 動的サイズ変更が容易 | - エッジ存在確認が O(V)<br>- 実装がやや複雑<br>- キャッシュ効率劣る    | 疎なグラフ、大規模グラフ | map+slice            |
| **エッジリスト** | - シンプルな構造<br>- 全エッジ操作が効率的<br>- メモリ効率良好       | - 隣接ノード取得が O(E)<br>- エッジ存在確認が O(E)<br>- 探索に不適     | エッジ中心の処理         | エッジ構造体の slice |

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

## 隣接行列

Go 言語では、スライスを用いて実装できる。

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

### 実装

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

## 隣接リスト

Go 言語で隣接リストを実装する際、最も柔軟で一般的な方法は、map とスライスを組み合わせた map[int]int である。
map のキーを頂点とし、値をその頂点に隣接する頂点のリスト（スライス）とすることで、頂点が整数でない場合や、頂点が連続していない場合にも容易に対応できる。

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

### 実装

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

## エッジリスト

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

### 実装

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

探索アルゴリズムはこのグラフ上で特定の目的（経路探索、到達可能性の確認、最短経路の計算など）を達成するために使用される。
グラフの構造（有向/無向、重み付き/重みなし、連結/非連結など）に応じて、適切な探索アルゴリズムを選択する必要がある。

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

# まとめ

グラフは複雑な関係性を表現する最も汎用的なデータ構造で、ネットワーク分析から人工知能まで幅広い分野で活用されている。適切な実装方式と効率的なアルゴリズムの選択により、大規模で複雑な問題の解決が可能になる。
