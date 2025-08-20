# Set（集合）とは

Set（集合）は、**重複のない要素**の集まりを管理するデータ構造。同じ要素を複数回含むことはできない。要素の順序は一般的に保証されないが、要素の存在確認、追加、削除を効率的に行うことができる。数学の集合論に基づいており、以下のような数学的な集合演算を行うことができる。

- **和集合**
- **積集合**
- **差集合**
- **対象差**

## 特徴

- **重複排除**: 同一の要素を複数格納することはできない。
- **順序不同**: 一般に要素が格納される順序は保証されず、インデックスによるアクセスも行えない。実装次第では挿入順やソート順を保持することができる。
- **可変性**: 要素の追加・削除が可能で、動的に内容を変更できる。
- **集合演算**: 和集合、積集合、差集合などの演算を効率的に実行できる。

## 操作の分類

- **基本操作**: 追加、削除、検索、サイズ取得
- **集合演算**: 和集合、積集合、差集合、対称差集合
- **関係演算**: 部分集合、上位集合、等価性の判定
- **変換操作**: 配列やスライスへの変換

## 高速に実行できる操作

- 追加(Add/Insert)
- 削除
- 存在確認（特定の要素が集合内に存在するか判定する）

# Set の種類

## Hash Set（ハッシュ集合）

内部的にハッシュテーブルを用いて要素を格納する。要素を追加するときは、そのハッシュ値を計算してバケットに格納する。ハッシュの特性により、以下の操作を平均的に O(1)で実行できる。

- 要素の追加
- 要素の削除
- 要素の存在確認

デメリットとして、要素の順序は保証されず、イテレーション時の順序も予測不能になる。

メモリ使用量はハッシュテーブル同様にオーバーヘッドがあるため中程度である。

## Tree Set（木構造集合）

内部的に平衡二分探索木を用いて要素を管理する。要素が常にソートされた状態で保持されるため、要素の順序が維持されるのが特徴。これにより特定の範囲に収まる全要素の取得などの範囲検索を効率的に行える。

デメリットとして、木構造に対する処理になることから、以下の処理に O(log n)の計算量がかかる。

メモリ使用量は、Hash Set よりも大きくなる。各要素をノードとして保持するため、ポインタなどのオーバーヘッドが発生するためである。

## Linked Hash Set（連結ハッシュ集合）

Hash Set の高速アクセス性能と連結リストの要素の追加順序を保持する機能を両立させた実装。ハッシュテーブルの要素間を双方向連結リストで連結することで挿入順序を記憶する。

追加・削除・存在確認の処理を平均 O(1)で実行でき、予測可能な挿入順でイテレーションされる。

メモリ使用量は、連結リストを維持ためのポインタが必要なため、Hash Set よりも大きくなる。

## Bit Set（ビット集合）

要素の有無をビット（0/1）で管理する集合データ構造。「小さな整数値」の集合に特化しており、各要素に対応するビットを 1 にすることで存在を表現する。

追加・削除・存在確認を平均 O(1)で実行可能。集合演算もビット演算で高速に実行できる。

注意点として、要素の順序は保持されず、管理できる数字の範囲をあらかじめ決めておく必要がある。

メモリ効率は非常に高く、1 要素が高くても 1bit にしかならない。

```go
type BitSet struct {
    bits []uint64
    size int
}

func NewBitSet(size int) *BitSet {
    return &BitSet{
        bits: make([]uint64, (size+63)/64),
        size: size,
    }
}

func (bs *BitSet) Set(i int)   { bs.bits[i/64] |= 1 << (i % 64) }
func (bs *BitSet) Clear(i int) { bs.bits[i/64] &^= 1 << (i % 64) }
func (bs *BitSet) Get(i int) bool { return bs.bits[i/64]&(1<<(i%64)) != 0 }
```

## Immutable Set（不変集合）

保存された要素の内容を変更できない構造。要素の追加・削除を行う場合、元の Set 構造とは異なる新しい Set 構造を返す。変更不可なので、スレッドセーフになる。

関数型プログラミングや並行処理に適している。

メモリ効率は実装による。

？: スレッドセーフとは？

```go
type ImmutableSet[T comparable] struct {
    items map[T]struct{}
}

func NewImmutableSet[T comparable](items ...T) *ImmutableSet[T] {
    s := &ImmutableSet[T]{items: make(map[T]struct{})}
    for _, item := range items {
        s.items[item] = struct{}{}
    }
    return s
}

func (s *ImmutableSet[T]) Add(item T) *ImmutableSet[T] {
    newItems := make(map[T]struct{}, len(s.items)+1)
    for k := range s.items {
        newItems[k] = struct{}{}
    }
    newItems[item] = struct{}{}
    return &ImmutableSet[T]{items: newItems}
}
```

## 実装の比較

| 実装方法                    | 追加/削除/存在確認の計算量 | 順序性   | 範囲クエリ         | メモリ使用量 | 集合演算の計算量 | 用途                                         |
| --------------------------- | -------------------------- | -------- | ------------------ | ------------ | ---------------- | -------------------------------------------- |
| ハッシュテーブル            | O(1)平均 / O(n)最悪        | なし     | 不可               | 中程度       | O(n + m)         | 一般的な重複排除・高速検索                   |
| 平衡二分探索木              | O(log n)                   | ソート順 | 効率的に可能       | やや多い     | O(n + m)         | 順序・範囲検索が必要な場合                   |
| ハッシュテーブル+連結リスト | O(1)平均 / O(n)最悪        | 挿入順   | 不可               | やや多い     | O(n + m)         | 挿入順を保持しつつ高速アクセスしたい場合     |
| ビット（Bit Set）           | O(1)                       | なし     | 限定的（整数範囲） | 非常に少ない | O(max(n,m)/64)   | 小さな整数集合・メモリ効率重視               |
| イミュータブル              | O(1)〜O(log n)（実装依存） | 実装依存 | 実装依存           | 実装依存     | O(n + m)         | 関数型・スレッドセーフ・変更不可が必要な場合 |

# Go 言語での Set 実装

## Go 言語の組み込み関数と Set

Go 言語には **Set 型は組み込まれていない**。
map を使って Set を実装するのが一般的。

### map を使った簡易 Set 実装

```go
// 文字列のSet（最も基本的な実装）
type StringSet map[string]bool

func NewStringSet() StringSet {
    return make(StringSet)
}

func (s StringSet) Add(item string) {
    s[item] = true
}

func (s StringSet) Remove(item string) {
    delete(s, item)
}

func (s StringSet) Contains(item string) bool {
    _, exists := s[item]
    return exists
}

func (s StringSet) Size() int {
    return len(s)
}

// 使用例
set := NewStringSet()
set.Add("apple")
set.Add("banana")
set.Add("apple") // 重複は無視される

fmt.Println(set.Contains("apple"))  // true
fmt.Println(set.Size())             // 2
```

### 空の構造体を使ったメモリ効率的な実装

```go
// メモリ効率を重視した実装
type StringSet map[string]struct{}

func NewStringSet() StringSet {
    return make(StringSet)
}

func (s StringSet) Add(item string) {
    s[item] = struct{}{} // 空の構造体でメモリ節約
}

func (s StringSet) Contains(item string) bool {
    _, exists := s[item]
    return exists
}
```

## 汎用的な Set 実装

### 型パラメータ（Generics）を使った実装

```go
package set

import (
    "fmt"
    "strings"
)

// 比較可能な型のための汎用Set
type Set[T comparable] struct {
    items map[T]struct{}
}

// 新しいSetを作成
func New[T comparable]() *Set[T] {
    return &Set[T]{
        items: make(map[T]struct{}),
    }
}

// スライスからSetを作成
func From[T comparable](items []T) *Set[T] {
    s := New[T]()
    for _, item := range items {
        s.Add(item)
    }
    return s
}

// 要素を追加
func (s *Set[T]) Add(item T) bool {
    if s.Contains(item) {
        return false // 既に存在
    }
    s.items[item] = struct{}{}
    return true // 新規追加
}

// 複数要素を追加
func (s *Set[T]) AddAll(items ...T) int {
    count := 0
    for _, item := range items {
        if s.Add(item) {
            count++
        }
    }
    return count
}

// 要素を削除
func (s *Set[T]) Remove(item T) bool {
    if !s.Contains(item) {
        return false // 存在しない
    }
    delete(s.items, item)
    return true // 削除成功
}

// 複数要素を削除
func (s *Set[T]) RemoveAll(items ...T) int {
    count := 0
    for _, item := range items {
        if s.Remove(item) {
            count++
        }
    }
    return count
}

// 要素の存在確認
func (s *Set[T]) Contains(item T) bool {
    _, exists := s.items[item]
    return exists
}

// サイズを取得
func (s *Set[T]) Size() int {
    return len(s.items)
}

// 空かどうかチェック
func (s *Set[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// 全要素を削除
func (s *Set[T]) Clear() {
    s.items = make(map[T]struct{})
}

// スライスに変換
func (s *Set[T]) ToSlice() []T {
    result := make([]T, 0, len(s.items))
    for item := range s.items {
        result = append(result, item)
    }
    return result
}

// 複製を作成
func (s *Set[T]) Clone() *Set[T] {
    clone := New[T]()
    for item := range s.items {
        clone.Add(item)
    }
    return clone
}

// 和集合
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
    result := s.Clone()
    for item := range other.items {
        result.Add(item)
    }
    return result
}

// 積集合
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
    result := New[T]()

    // より小さい集合を基準にする（効率化）
    smaller, larger := s, other
    if other.Size() < s.Size() {
        smaller, larger = other, s
    }

    for item := range smaller.items {
        if larger.Contains(item) {
            result.Add(item)
        }
    }
    return result
}

// 差集合
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
    result := New[T]()
    for item := range s.items {
        if !other.Contains(item) {
            result.Add(item)
        }
    }
    return result
}

// 対称差集合
func (s *Set[T]) SymmetricDifference(other *Set[T]) *Set[T] {
    result := New[T]()

    // sにあってotherにない要素
    for item := range s.items {
        if !other.Contains(item) {
            result.Add(item)
        }
    }

    // otherにあってsにない要素
    for item := range other.items {
        if !s.Contains(item) {
            result.Add(item)
        }
    }

    return result
}

// 部分集合判定
func (s *Set[T]) IsSubset(other *Set[T]) bool {
    if s.Size() > other.Size() {
        return false
    }

    for item := range s.items {
        if !other.Contains(item) {
            return false
        }
    }
    return true
}

// 上位集合判定
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
    return other.IsSubset(s)
}

// 互いに素判定
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
    // より小さい集合を基準にする
    smaller, larger := s, other
    if other.Size() < s.Size() {
        smaller, larger = other, s
    }

    for item := range smaller.items {
        if larger.Contains(item) {
            return false
        }
    }
    return true
}

// 等価性判定
func (s *Set[T]) Equals(other *Set[T]) bool {
    if s.Size() != other.Size() {
        return false
    }

    for item := range s.items {
        if !other.Contains(item) {
            return false
        }
    }
    return true
}

// In-place操作: 和集合で更新
func (s *Set[T]) UnionWith(other *Set[T]) {
    for item := range other.items {
        s.Add(item)
    }
}

// In-place操作: 積集合で更新
func (s *Set[T]) IntersectWith(other *Set[T]) {
    toRemove := make([]T, 0)
    for item := range s.items {
        if !other.Contains(item) {
            toRemove = append(toRemove, item)
        }
    }

    for _, item := range toRemove {
        s.Remove(item)
    }
}

// In-place操作: 差集合で更新
func (s *Set[T]) DifferenceWith(other *Set[T]) {
    for item := range other.items {
        s.Remove(item)
    }
}

// 条件に合う要素を削除
func (s *Set[T]) RemoveIf(predicate func(T) bool) int {
    toRemove := make([]T, 0)
    for item := range s.items {
        if predicate(item) {
            toRemove = append(toRemove, item)
        }
    }

    for _, item := range toRemove {
        s.Remove(item)
    }

    return len(toRemove)
}

// 各要素に関数を適用
func (s *Set[T]) ForEach(fn func(T)) {
    for item := range s.items {
        fn(item)
    }
}

// 条件に合う要素があるかチェック
func (s *Set[T]) Any(predicate func(T) bool) bool {
    for item := range s.items {
        if predicate(item) {
            return true
        }
    }
    return false
}

// 全要素が条件に合うかチェック
func (s *Set[T]) All(predicate func(T) bool) bool {
    for item := range s.items {
        if !predicate(item) {
            return false
        }
    }
    return true
}

// 条件に合う要素でフィルタ
func (s *Set[T]) Filter(predicate func(T) bool) *Set[T] {
    result := New[T]()
    for item := range s.items {
        if predicate(item) {
            result.Add(item)
        }
    }
    return result
}

// 文字列表現
func (s *Set[T]) String() string {
    if s.IsEmpty() {
        return "{}"
    }

    items := make([]string, 0, s.Size())
    for item := range s.items {
        items = append(items, fmt.Sprintf("%v", item))
    }

    return "{" + strings.Join(items, ", ") + "}"
}

// イテレータパターン（チャネル使用）
func (s *Set[T]) Iterator() <-chan T {
    ch := make(chan T, s.Size())

    go func() {
        defer close(ch)
        for item := range s.items {
            ch <- item
        }
    }()

    return ch
}
```

### Tree Set の実装例

```go
package treeset

import (
    "fmt"
    "golang.org/x/exp/constraints"
)

// ノード構造体
type node[T constraints.Ordered] struct {
    value       T
    left, right *node[T]
    height      int
}

// Tree Set構造体（AVL木ベース）
type TreeSet[T constraints.Ordered] struct {
    root *node[T]
    size int
}

// 新しいTree Setを作成
func New[T constraints.Ordered]() *TreeSet[T] {
    return &TreeSet[T]{}
}

// 高さを取得
func (n *node[T]) getHeight() int {
    if n == nil {
        return 0
    }
    return n.height
}

// バランス係数を計算
func (n *node[T]) getBalance() int {
    if n == nil {
        return 0
    }
    return n.left.getHeight() - n.right.getHeight()
}

// 高さを更新
func (n *node[T]) updateHeight() {
    leftHeight := n.left.getHeight()
    rightHeight := n.right.getHeight()

    if leftHeight > rightHeight {
        n.height = leftHeight + 1
    } else {
        n.height = rightHeight + 1
    }
}

// 右回転
func (ts *TreeSet[T]) rotateRight(y *node[T]) *node[T] {
    x := y.left
    t2 := x.right

    x.right = y
    y.left = t2

    y.updateHeight()
    x.updateHeight()

    return x
}

// 左回転
func (ts *TreeSet[T]) rotateLeft(x *node[T]) *node[T] {
    y := x.right
    t2 := y.left

    y.left = x
    x.right = t2

    x.updateHeight()
    y.updateHeight()

    return y
}

// 要素を追加
func (ts *TreeSet[T]) Add(value T) bool {
    oldSize := ts.size
    ts.root = ts.insert(ts.root, value)
    return ts.size > oldSize
}

func (ts *TreeSet[T]) insert(n *node[T], value T) *node[T] {
    // 基本的な二分探索木の挿入
    if n == nil {
        ts.size++
        return &node[T]{
            value:  value,
            height: 1,
        }
    }

    if value < n.value {
        n.left = ts.insert(n.left, value)
    } else if value > n.value {
        n.right = ts.insert(n.right, value)
    } else {
        // 重複は追加しない
        return n
    }

    // 高さを更新
    n.updateHeight()

    // バランスを確認して回転
    balance := n.getBalance()

    // Left Left Case
    if balance > 1 && value < n.left.value {
        return ts.rotateRight(n)
    }

    // Right Right Case
    if balance < -1 && value > n.right.value {
        return ts.rotateLeft(n)
    }

    // Left Right Case
    if balance > 1 && value > n.left.value {
        n.left = ts.rotateLeft(n.left)
        return ts.rotateRight(n)
    }

    // Right Left Case
    if balance < -1 && value < n.right.value {
        n.right = ts.rotateRight(n.right)
        return ts.rotateLeft(n)
    }

    return n
}

// 要素の存在確認
func (ts *TreeSet[T]) Contains(value T) bool {
    return ts.search(ts.root, value)
}

func (ts *TreeSet[T]) search(n *node[T], value T) bool {
    if n == nil {
        return false
    }

    if value == n.value {
        return true
    } else if value < n.value {
        return ts.search(n.left, value)
    } else {
        return ts.search(n.right, value)
    }
}

// 最小値を取得
func (ts *TreeSet[T]) Min() (T, bool) {
    if ts.root == nil {
        var zero T
        return zero, false
    }

    min := ts.findMin(ts.root)
    return min.value, true
}

func (ts *TreeSet[T]) findMin(n *node[T]) *node[T] {
    for n.left != nil {
        n = n.left
    }
    return n
}

// 最大値を取得
func (ts *TreeSet[T]) Max() (T, bool) {
    if ts.root == nil {
        var zero T
        return zero, false
    }

    max := ts.findMax(ts.root)
    return max.value, true
}

func (ts *TreeSet[T]) findMax(n *node[T]) *node[T] {
    for n.right != nil {
        n = n.right
    }
    return n
}

// ソート済みスライスに変換
func (ts *TreeSet[T]) ToSortedSlice() []T {
    result := make([]T, 0, ts.size)
    ts.inorderTraversal(ts.root, &result)
    return result
}

func (ts *TreeSet[T]) inorderTraversal(n *node[T], result *[]T) {
    if n != nil {
        ts.inorderTraversal(n.left, result)
        *result = append(*result, n.value)
        ts.inorderTraversal(n.right, result)
    }
}

// 範囲内の要素を取得
func (ts *TreeSet[T]) Range(min, max T) []T {
    var result []T
    ts.rangeQuery(ts.root, min, max, &result)
    return result
}

func (ts *TreeSet[T]) rangeQuery(n *node[T], min, max T, result *[]T) {
    if n == nil {
        return
    }

    if min < n.value {
        ts.rangeQuery(n.left, min, max, result)
    }

    if min <= n.value && n.value <= max {
        *result = append(*result, n.value)
    }

    if n.value < max {
        ts.rangeQuery(n.right, min, max, result)
    }
}

// サイズを取得
func (ts *TreeSet[T]) Size() int {
    return ts.size
}

// 文字列表現
func (ts *TreeSet[T]) String() string {
    values := ts.ToSortedSlice()
    return fmt.Sprintf("TreeSet%v", values)
}
```

### Bit Set の実装例

```go
package bitset

import (
    "fmt"
    "math/bits"
)

// Bit Set構造体
type BitSet struct {
    bits []uint64
    size int
}

// 新しいBit Setを作成
func New(size int) *BitSet {
    wordsNeeded := (size + 63) / 64
    return &BitSet{
        bits: make([]uint64, wordsNeeded),
        size: size,
    }
}

// ビットを設定
func (bs *BitSet) Set(index int) {
    if index >= bs.size || index < 0 {
        return
    }

    wordIndex := index / 64
    bitIndex := index % 64
    bs.bits[wordIndex] |= (1 << bitIndex)
}

// ビットをクリア
func (bs *BitSet) Clear(index int) {
    if index >= bs.size || index < 0 {
        return
    }

    wordIndex := index / 64
    bitIndex := index % 64
    bs.bits[wordIndex] &^= (1 << bitIndex)
}

// ビットをフリップ
func (bs *BitSet) Flip(index int) {
    if index >= bs.size || index < 0 {
        return
    }

    wordIndex := index / 64
    bitIndex := index % 64
    bs.bits[wordIndex] ^= (1 << bitIndex)
}

// ビットの状態を取得
func (bs *BitSet) Get(index int) bool {
    if index >= bs.size || index < 0 {
        return false
    }

    wordIndex := index / 64
    bitIndex := index % 64
    return (bs.bits[wordIndex] & (1 << bitIndex)) != 0
}

// 設定されているビット数をカウント
func (bs *BitSet) Count() int {
    count := 0
    for _, word := range bs.bits {
        count += bits.OnesCount64(word)
    }
    return count
}

// 和集合
func (bs *BitSet) Union(other *BitSet) *BitSet {
    maxSize := bs.size
    if other.size > maxSize {
        maxSize = other.size
    }

    result := New(maxSize)
    minWords := len(bs.bits)
    if len(other.bits) < minWords {
        minWords = len(other.bits)
    }

    // 共通部分
    for i := 0; i < minWords; i++ {
        result.bits[i] = bs.bits[i] | other.bits[i]
    }

    // 残り部分をコピー
    if len(bs.bits) > minWords {
        copy(result.bits[minWords:], bs.bits[minWords:])
    } else if len(other.bits) > minWords {
        copy(result.bits[minWords:], other.bits[minWords:])
    }

    return result
}

// 積集合
func (bs *BitSet) Intersection(other *BitSet) *BitSet {
    minSize := bs.size
    if other.size < minSize {
        minSize = other.size
    }

    result := New(minSize)
    minWords := len(bs.bits)
    if len(other.bits) < minWords {
        minWords = len(other.bits)
    }

    for i := 0; i < minWords; i++ {
        result.bits[i] = bs.bits[i] & other.bits[i]
    }

    return result
}

// 文字列表現
func (bs *BitSet) String() string {
    setBits := make([]int, 0)
    for i := 0; i < bs.size; i++ {
        if bs.Get(i) {
            setBits = append(setBits, i)
        }
    }
    return fmt.Sprintf("BitSet%v", setBits)
}
```

## 使用例とベンチマーク

```go
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    // 基本的な使用例
    basicExample()

    // 集合演算の例
    setOperationsExample()

    // パフォーマンス比較
    performanceComparison()
}

func basicExample() {
    fmt.Println("=== 基本的な使用例 ===")

    // Hash Set
    hashSet := New[int]()
    hashSet.AddAll(1, 2, 3, 4, 5)
    fmt.Printf("Hash Set: %s\n", hashSet)
    fmt.Printf("Contains 3: %t\n", hashSet.Contains(3))
    fmt.Printf("Size: %d\n", hashSet.Size())

    // Tree Set
    treeSet := NewTreeSet[int]()
    treeSet.Add(5)
    treeSet.Add(2)
    treeSet.Add(8)
    treeSet.Add(1)
    fmt.Printf("Tree Set: %s\n", treeSet)

    min, _ := treeSet.Min()
    max, _ := treeSet.Max()
    fmt.Printf("Min: %d, Max: %d\n", min, max)
}

func setOperationsExample() {
    fmt.Println("\n=== 集合演算の例 ===")

    setA := From([]int{1, 2, 3, 4, 5})
    setB := From([]int{4, 5, 6, 7, 8})

    fmt.Printf("Set A: %s\n", setA)
    fmt.Printf("Set B: %s\n", setB)

    union := setA.Union(setB)
    fmt.Printf("A ∪ B: %s\n", union)

    intersection := setA.Intersection(setB)
    fmt.Printf("A ∩ B: %s\n", intersection)

    difference := setA.Difference(setB)
    fmt.Printf("A - B: %s\n", difference)

    symmetricDiff := setA.SymmetricDifference(setB)
    fmt.Printf("A △ B: %s\n", symmetricDiff)

    fmt.Printf("A ⊆ B: %t\n", setA.IsSubset(setB))
    fmt.Printf("A ∩ B = ∅: %t\n", setA.IsDisjoint(setB))
}

func performanceComparison() {
    fmt.Println("\n=== パフォーマンス比較 ===")

    const n = 100000

    // Hash Set ベンチマーク
    start := time.Now()
    hashSet := New[int]()
    for i := 0; i < n; i++ {
        hashSet.Add(rand.Intn(n))
    }
    hashSetTime := time.Since(start)

    // Tree Set ベンチマーク
    start = time.Now()
    treeSet := NewTreeSet[int]()
    for i := 0; i < n; i++ {
        treeSet.Add(rand.Intn(n))
    }
    treeSetTime := time.Since(start)

    // Bit Set ベンチマーク
    start = time.Now()
    bitSet := NewBitSet(n)
    for i := 0; i < n; i++ {
        bitSet.Set(rand.Intn(n))
    }
    bitSetTime := time.Since(start)

    fmt.Printf("Hash Set: %v (要素数: %d)\n", hashSetTime, hashSet.Size())
    fmt.Printf("Tree Set: %v (要素数: %d)\n", treeSetTime, treeSet.Size())
    fmt.Printf("Bit Set:  %v (要素数: %d)\n", bitSetTime, bitSet.Count())

    // 検索パフォーマンス
    searchValue := rand.Intn(n)

    start = time.Now()
    for i := 0; i < 10000; i++ {
        hashSet.Contains(searchValue)
    }
    hashSearchTime := time.Since(start)

    start = time.Now()
    for i := 0; i < 10000; i++ {
        treeSet.Contains(searchValue)
    }
    treeSearchTime := time.Since(start)

    start = time.Now()
    for i := 0; i < 10000; i++ {
        bitSet.Get(searchValue)
    }
    bitSearchTime := time.Since(start)

    fmt.Printf("\n検索パフォーマンス (10000回):\n")
    fmt.Printf("Hash Set: %v\n", hashSearchTime)
    fmt.Printf("Tree Set: %v\n", treeSearchTime)
    fmt.Printf("Bit Set:  %v\n", bitSearchTime)
}
```

## mapstruct{}による実装

最も一般的かつ推奨される方法。
キーを Set の要素とし、値をからの構造体 struct{} とする map を利用する。
HashSet に相当し、各操作の平均計算量は O(1) となる。

```go
// Set: ユニークな要素のコレクション
type Set struct{
    elements mapstruct{}
}

// NewSet: 新しいSetを作成する
func NewSet() *Set{
    return &Set{
        elements: make(mapstructt{}),
    }
}

// Add: 要素をSetに挿入する
func (s *Set) Add(value T) {
    s.elements[value] = struct{}{}
}

// Remove: Setから要素を削除
func (s *Set) Remove(value T) {
    delete(s.elements, value)
}

// Contains: 要素が Set に含まれているかどうか
func (s *Set) Contains(value T) bool {
    _, found := s.elements[value]
    return found
}

```

特徴

- map のキーの一意性により、Set の要素の一意性を実現できる
- map の値として意味のあるデータを持たせる必要はなく、struct{}が最もメモリ効率が良い
  - 空の構造体（struct{}）はフィールドを持たないので、メモリ幅がゼロであり、値を格納するための追加のメモリを一切消費しない

集合演算も map ベースの実装で実現可能。

```go
// 2つの Set の和集合を返す
func (s *Set) Union(other *Set) *Set {
    result := NewSet()
    for key := range s.elements {
        result.Add(key)
    }
    for key := range other.elements {
        result.Add(key)
    }
    return result
}

// 2つの Set の積集合を返す
func (s *Set) Intersection(other *Set) *Set {
    result := NewSet()
    for key := range s.elements {
        if other.Contains(key) {
            result.Add(key)
        }
    }
    return result
}

```

## サードパーティライブラリの活用（github.com/deckarep/golang-set）

より高度な機能やスレッドセーフな実装が必要な場合、実績のあるサードパーティライブラリである github.com/deckarep/golang-set が有効。

Python の Set 実装をモデルとした豊富な API を提供し、ジェネリクスをサポートしている。

パフォーマンスを重視した非スレッドセーフ版と、並行処理に適したスレッドセーフ版の両方を提供している。

```go
package main

import (
    "fmt"
    mapset "github.com/deckarep/golang-set/v2"
)

func main() {
    required := mapset.NewSet[string]("cooking", "english", "math", "biology")
    sciences := mapset.NewSet[string]("biology", "chemistry")

    // 和集合
    // 出力: "cooking", "english", "math", "biology", "chemistry"
    allClasses := reequired.Union(sciences)
    fmt.Println("All classes:", allClassees)

    // 存在確認
    fmt.Println("Is cooking a science?", sciences.Contains("cooking"))

    // 差集合
    // 出力: "cooking", "english", "math"
    nonScience := required.Difference(sciences)
    fmt.Println("Required but not science:", nonScience)

    // 積集合
    // 出力: "biology"
    requiredScience := required.Intersect(sciences)
    fmt.Println("Required science class:", requiredScience)

    // 要素数
    fmt.Println("Number of science classes:", sciences.Cardinality())

}

```

# Set の応用例

## 1. 重複排除・ユニーク処理

- **データクリーニング**: 重複レコードの除去
- **ID の一意性確保**: ユーザー ID、商品 ID の管理
- **重複ファイル検出**: ファイルハッシュによる重複判定
- **メール配信**: 重複メールアドレスの排除

## 2. メンバーシップ管理

- **ユーザー権限**: 特定機能へのアクセス権管理
- **グループ管理**: ユーザーのグループ所属判定
- **ブラックリスト/ホワイトリスト**: 許可/禁止 IP アドレス管理
- **購読管理**: ニュースレター購読者管理

## 3. グラフ・ネットワーク処理

- **隣接ノード管理**: グラフの隣接リスト
- **訪問済みノード**: グラフ探索での重複訪問防止
- **連結成分**: グラフの連結成分解析
- **最短経路**: ダイクストラ法での訪問済み管理

## 4. キャッシュ・データ管理

- **キーの管理**: キャッシュのキー集合
- **無効化管理**: 無効化対象の管理
- **依存関係**: データ間の依存関係管理
- **変更追跡**: 変更されたオブジェクトの追跡

## 5. セキュリティ・認証

- **セッション管理**: アクティブセッションの追跡
- **トークン管理**: 有効な API トークンの管理
- **ブルートフォース対策**: 攻撃 IP アドレスの記録
- **パスワードポリシー**: 使用済みパスワードの管理

## 6. 検索・フィルタリング

- **検索条件**: 複数条件の組み合わせ
- **タグシステム**: 商品・記事のタグ管理
- **カテゴリフィルタ**: 商品カテゴリの選択状態
- **ファセット検索**: 多面的な検索条件

## 7. スケジューリング・並行処理

- **実行中タスク**: 現在実行中のタスク ID
- **ロック管理**: 取得済みロックの管理
- **リソース管理**: 使用中リソースの追跡
- **ワーカー管理**: アクティブワーカーの管理

## 8. アルゴリズム・データ処理

- **Union-Find**: 素集合データ構造の基盤
- **ブルームフィルタ**: 確率的集合データ構造
- **A\*探索**: オープンリスト・クローズドリスト
- **動的計画法**: メモ化での重複計算回避

## 9. Web 開発・API

- **CORS 設定**: 許可ドメインの管理
- **レート制限**: API アクセス制限の管理
- **フィーチャーフラグ**: 有効機能の管理
- **A/B テスト**: テストグループの管理

## 10. ゲーム開発

- **所持アイテム**: プレイヤーの所持品管理
- **スキル習得**: 習得済みスキルの管理
- **フレンドリスト**: 友達関係の管理
- **達成実績**: アンロック済み実績の管理

# Set と他のデータ構造との比較

## Set(HashSet) vs Array/Slice

Array/Slice では、要素の存在確認に対する計算量が大きく異なる。ソートされていない Array/Slice では線形探索が必要で O(n) 必要。Set（特に HashSet）は平均 O(1) で完了する。

順序性も異なり、Array/Slice は保持し、重複も許容するが、Set はどちらも許容しない。

## Set(HashSet) vs Map

ほとんどの性能は同じであり、キーに対する値をどう扱うかが大きく異なる。
Set は要素の存在確認が主な役割なので、値として mapstruct{} を用いる。Map はキーに対応する値を取得することに関心がある。Set は Map の特殊なケース（値に関心がないケース）とみなすことができる。

# まとめ

Set（集合）は、重複のない要素の集まりを効率的に管理するための基本的なデータ構造で、重複排除、高速な存在確認、集合演算などの機能を提供し、様々なアプリケーションで活用されている。

Go 言語では標準で Set 型が提供されていないが、map を使った実装や、Generics を活用した汎用的な実装が可能。用途に応じて Hash Set、Tree Set、Bit Set などの実装方式を選択し、パフォーマンスとメモリ効率のバランスを考慮することが重要。
