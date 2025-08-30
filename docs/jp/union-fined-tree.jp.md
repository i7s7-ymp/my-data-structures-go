# Union-Find 木（Union-Find Tree）とは

Union-Find 木（または Disjoint Set Union, DSU、素集合データ構造）は、**複数の要素をグループ（集合）に分けて管理し、グループの統合（Union）や所属判定（Find）を高速に行うためのデータ構造**です。主にグラフの連結成分判定やクラスター管理、ネットワークの連結性判定などで利用されます。

**このデータ構造の最大の利点は、集合の合併と所属判定をほぼ定数時間（アッカーマン関数の逆数オーダー）で実現できる点です。**

# Union-Find 木の構成要素

## 基本要素

- **要素（Element）**: 管理対象となる個々のデータ（ノード）
- **親配列（Parent Array）**: 各要素の親を示す配列。根（リーダー）は自分自身を指す
- **ランク配列（Rank Array）**: 各集合の木の高さやサイズを管理し、効率的な合併を実現
- **集合（Set/Group）**: 互いに連結された要素の集まり
- **根（Root/Leader）**: 集合を代表する要素

# Union-Find 木の特徴

- **集合の合併（Union）**: 2 つの集合を 1 つにまとめる
- **所属判定（Find）**: 要素がどの集合に属するか（根を求める）
- **経路圧縮（Path Compression）**: Find 時に親を根に直接つなぎ、木の高さを低減
- **ランク合併（Union by Rank/Size）**: 木の高さやサイズを考慮して合併し、効率化
- **動的構造**: 要素の追加や集合の統合が容易

# Union-Find 木の構成要素の図解

## 基本的な Union-Find 木の例

```
初期状態（各要素が独立した集合）:
Index: 0 1 2 3 4
Parent:0 1 2 3 4

Union(1, 2) 実行後:
Index: 0 1 2 3 4
Parent:0 1 1 3 4

Union(3, 4) 実行後:
Index: 0 1 2 3 4
Parent:0 1 1 3 3

Union(2, 3) 実行後（経路圧縮あり）:
Index: 0 1 2 3 4
Parent:0 1 1 1 3
```

# Union-Find 木の実装方法

## 一般的な実装方法

### 配列ベース実装

- 各要素の親を配列で管理（parent[i]）
- ランクやサイズも配列で管理（rank[i]や size[i]）
- 経路圧縮とランク合併を組み合わせることで高速化

### ポインタベース実装

- ノードを構造体で表現し、親ノードへのポインタを持つ
- 配列実装より柔軟だが、メモリ効率や速度は配列実装が優れる

## Go による Union-Find 木の実装例

```go
package unionfind

// Union-Find構造体
type UnionFind struct {
    parent []int // 各要素の親
    size   []int // 各集合のサイズ
}

// n個の要素で初期化
func NewUnionFind(n int) *UnionFind {
    uf := &UnionFind{
        parent: make([]int, n),
        size:   make([]int, n),
    }
    for i := 0; i < n; i++ {
        uf.parent[i] = i
        uf.size[i] = 1
    }
    return uf
}

// 根を求める（経路圧縮あり）
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

// 2つの集合を統合（ランク合併）
func (uf *UnionFind) Union(x, y int) {
    rx := uf.Find(x)
    ry := uf.Find(y)
    if rx == ry {
        return
    }
    // サイズが大きい方に小さい方をつなぐ
    if uf.size[rx] < uf.size[ry] {
        uf.parent[rx] = ry
        uf.size[ry] += uf.size[rx]
    } else {
        uf.parent[ry] = rx
        uf.size[rx] += uf.size[ry]
    }
}

// 同じ集合か判定
func (uf *UnionFind) Same(x, y int) bool {
    return uf.Find(x) == uf.Find(y)
}

// 集合のサイズを取得
func (uf *UnionFind) Size(x int) int {
    return uf.size[uf.Find(x)]
}

// 集合の代表（根）を取得
func (uf *UnionFind) Leader(x int) int {
    return uf.Find(x)
}

// すべての集合を取得（map[代表]要素リスト）
func (uf *UnionFind) AllGroups() map[int][]int {
    groups := make(map[int][]int)
    for i := range uf.parent {
        leader := uf.Find(i)
        groups[leader] = append(groups[leader], i)
    }
    return groups
}
```

# Union-Find 木の応用例

- **グラフの連結成分判定**: 無向グラフの連結判定やクラスター分割
- **クラスター管理**: ソーシャルネットワークやクラスタリング
- **ネットワーク接続管理**: ネットワークの連結性や障害検出
- **画像処理**: ラベリングや領域分割
- **パーコレーション判定**: 物理シミュレーション
- **クルスカル法**: 最小全域木アルゴリズム
