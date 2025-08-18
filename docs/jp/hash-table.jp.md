# ハッシュテーブル（Hash Table）とは

ハッシュテーブルは、キーと値のペアを効率的に格納・検索するデータ構造。ハッシュ化を用いてキーを配列のインデックスにマッピングし、高速なデータアクセスを実現できる。

**このデータ構造を使う一番の特徴は、キーによる高速な検索・挿入・削除が可能な点である。**

データベースのインデックス、キャッシュシステム、辞書機能など、高速なデータアクセスが必要な場面で広く利用される。

## ハッシュ化の流れ

1. ハッシュ関数を用いてキーを整数値（ハッシュ値）に変換
2. ハッシュ値を配列（バケット配列）の index として利用し、データを格納

### ハッシュ関数

- キーを配列のインデックスに変換する関数
- 良いハッシュ関数の条件：
  - 高速に計算できる
  - 値が均等に分散される
  - 決定的（同じキーは常に同じ値を返す）
- 完全ハッシュ関数
  - どのキーに対しても、ハッシュ値が異なるハッシュ関数
  - すべての異なるキーがユニークなインデックスにマッピングされ、データの挿入・検索・削除は配列への直接アクセスとなり、計算量は $O(1) となる

### バケット配列

- 実際にデータを格納する配列
- 各要素（バケット）にキー・バリューペアを格納

## 特徴

- **キー・バリュー型**: キーと値のペアでデータを管理
- **ハッシュ関数**: キーを配列のインデックスに変換する関数
- **バケット配列**: キー・値のペアを格納する配列
- **高速アクセス**: キーを用いて、平均的に O(1)で値にアクセスできる
- **動的サイズ**: 必要に応じてサイズを拡張可能
- **順序なし**: 要素の挿入順序は保持されない
- **一意キー**: 同じキーは一度しか存在できない

## 注意点

1. **ハッシュ衝突**: 異なるキーが同じハッシュ値を生成し、同じバケットの index にマッピングされてしまうこと。ハッシュ関数の性質上発生してしまうものなので、対策が必要。
2. **最悪時計算量**: ハッシュ衝突が多発すると、最悪の場合 O(n)まで性能が劣化する。
3. **メモリ使用量**: 負荷率を低く保つため、実際の要素数より多くのメモリを使用する。
4. **順序性なし**: 要素の挿入順序や値の順序は保持されない。

# 連想配列とハッシュテーブル

## 連想配列とは

**連想配列（Associative Array）** は、キーと値のペアを管理する抽象的なデータ型。
**辞書（Dictionary）** や **マップ（Map）** とも呼ばれいる。
Go 言語では、`map`型が連想配列に相当する。

## ハッシュテーブルとの関係

連想配列は概念、ハッシュテーブルは実装手段。

- 連想配列：「キーで値を管理したい」という要求（抽象データ型）
- ハッシュテーブル：「その要求を高速に実現する方法」（具体的実装）

ハッシュテーブル以外にも、以下のようなデータ構造で実装可能。
| 実装方法 | 平均計算量 | 特徴 |
| ------------------ | ---------- | ------------------------------ |
| **ハッシュテーブル** | O(1) | 最も高速、順序なし |
| **平衡二分木** | O(log n) | 順序付き、安定した性能 |
| **配列** | O(n) | シンプル、小規模データ向け |
| **連結リスト** | O(n) | 動的、メモリ効率的 |

## 連想配列の特徴

- **キー・バリューペア**: 任意のキーを指定して、対応する値を効率的に格納・取得・削除できる
- **一意性**: 同じキーは一度だけ存在できる
- **動的**: 実行時にキー・バリューペアを追加・削除可能
- **抽象データ型**: 実装方法に依存しない概念的なデータ構造

# ハッシュ衝突の解決方法

現実では完全ハッシュ関数を作ることは難しく、異なるキーが同じ idnex にマッピング（衝突）されてしまう。
衝突を効率よく解決する方法は主に 2 つある。

- **チェイン法**: 同じインデックスの要素を連結リストで管理
- **オープンアドレス法**: 別の空いているインデックスを探す

## チェイン法（Separate Chaining）

各配列要素に連結リストを持たせ、衝突した要素を同じインデックスのリストに追加する方法。

### 仕組み

ハッシュテーブル保存時、衝突が発生したデータを連結リストに保存し。そのハッシュ値の index を持つバケットにはデータが保存された連結リストへのポインタを持たせる

1. キーからハッシュ値を計算し、対応するバケットを選択
2. バケットがからなら、新しいデータを格納
3. バケットにデータが存在していたら、バケットが指す連結リストの末尾に新しいデータを追加

### 特徴

- **実装が簡単**: 連結リストを使用するため理解しやすい
- **動的サイズ**: テーブルサイズを超えても要素を追加可能
- **削除が容易**: 要素の削除が簡単（連結リストからノードを削除するだけ）

### 注意点

- **メモリ増加**: 連結リスト管理のために追加のメモリが必要
- **探索時間の増加**: 一つの連結リストが極端に長くなると、連結リストの検索性能が線形探索と同じになり、遅くなる（最悪 O(n)）
- **キャッシュ効率の低下**: 連結リストをたどる際にメモリアクセスが非連続的になるため、CPU キャッシュ効率が悪くなることがある

### 向いているケース

- データの挿入・削除が頻繁に行われる場合
- 要素数が多く、衝突が発生しやすい場合

## オープンアドレス法（Open Addressing）

衝突が発生した場合、別の空いている位置を探して要素を配置する方法。

### 仕組み

- ハッシュテーブル内の空いている別のバケットを探してデータを格納する。外部のデータ構造を利用せず、全てのデータはハッシュテーブルの配列内に直接保存される
- 空きバケットの探し方（探査法）にいくつか種類がある

### 実装方法

#### 線形探査法（Linear Probing）

- **探査方法**: 衝突したバケットを基準に順番に空きを探す
- **キャッシュ効率**: 連続したメモリアクセスでキャッシュ効率が良い
- **一次クラスタリングの問題**: データが連続したブロック（クラスタ）を作りやすくなり、クラスタが大きくなると新しいデータを挿入する際に何度も衝突が起き、性能が悪化する可能性がある

#### 二次探査法（Quadratic Probing）

- **探査方法**: 衝突したバケットを基準に n\*\*2（1, 4, 9,...）離れたバケットを見ていく
- **一次クラスタリングを緩和**: 線形探査より分散が良い
- **二次クラスタリングの問題**: 特定のパターンで衝突が連鎖する可能性がある

#### 二重ハッシュ法（Double Hashing）

- **探査方法**: 2 種類のハッシュを用意する。最も効率が良い。
  1. 最初のハッシュ関数で格納位置を計算
  2. 衝突したら、2 番目のハッシュ関数で「次の候補を探すための間隔」を計算し、その間隔でバケットを探索
- **均等分散**: 探索の間隔がキーごとに異なるため最も均等に分散されやすく（クラスタリングが発生しづらく）、高いパフォーマンスを維持できる

### 特徴

- **メモリ効率**: ポインタを使わないので、追加のメモリが不要で、メモリ効率が良い
- **キャッシュ効率**: データが配列内に密集しているので、CPU のキャッシュ効率が良い傾向にある

### 注意点

- **性能低下**: テーブルが満杯に近づくと性能が急激に低下
- **削除の複雑さ**: 要素の削除が複雑。単純にデータを削除すると、探査の連鎖が途切れてしまい、その先にあるはずのデータが見つからなくなる可能性がある。利用する場合は「削除済み」を表す特別なマーカーを置く
- **サイズ設計**: ハッシュテーブルのサイズをあらかじめ設計する必要がある

### 向いているケース

- メモリ使用量を厳密に管理したい
- データの追加・削除が少なく、読み取りが中心
- テーブルの負荷率を小さく保てる

# 実装方法

## Go の組み込み関数（map）による実装

Go には`map`という組み込みのハッシュテーブル型があり、最も簡単で効率的な実装方法。

### 特徴

- **簡潔性**: 宣言と操作が非常にシンプル
- **高性能**: Go runtime による最適化済み実装
- **動的サイズ**: 自動的にサイズが拡張される
- **型安全**: コンパイル時に型チェックが行われる
- **ガベージコレクション**: メモリ管理が自動化されている

### 注意点

- **順序性なし**: 要素の反復順序は保証されない（Go 1.0 以降はランダム化されている）
- **並行安全性なし**: 複数の goroutine から同時アクセスする場合は`sync.Map`を使用する必要がある
- **削除時の動作**: `delete()`関数使用時、キーが存在しなくてもエラーにならない
- **ゼロ値の扱い**: 存在しないキーアクセス時はゼロ値が返される

```go
package main

import "fmt"

func main() {
    // マップの作成方法

    // 1. make関数を使用
    hashMap := make(map[string]int)

    // 2. マップリテラルを使用
    hashMap2 := map[string]int{
        "apple":  100,
        "banana": 200,
    }

    // 3. 初期容量を指定（パフォーマンス最適化）
    hashMap3 := make(map[string]int, 100)

    // 基本操作
    hashMap["key1"] = 10        // 挿入
    value := hashMap["key1"]    // 取得
    delete(hashMap, "key1")     // 削除

    // 存在確認（二番目の戻り値で判定）
    value, exists := hashMap["key1"]
    if exists {
        fmt.Println("値:", value)
    } else {
        fmt.Println("キーが存在しません")
    }

    // 要素数の取得
    size := len(hashMap)
    fmt.Println("要素数:", size)

    // 全要素の反復処理
    for key, value := range hashMap2 {
        fmt.Printf("%s: %d\n", key, value)
    }
}
```

### 並行安全なマップ（sync.Map）

複数の goroutine から同時アクセスする場合は`sync.Map`を使用する。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map

    // 値の設定
    m.Store("key1", "value1")

    // 値の取得
    if value, ok := m.Load("key1"); ok {
        fmt.Println("値:", value)
    }

    // 値の削除
    m.Delete("key1")

    // 存在しない場合のみ設定
    actual, loaded := m.LoadOrStore("key2", "value2")
    if !loaded {
        fmt.Println("新しく設定:", actual)
    }

    // 全要素の反復処理
    m.Range(func(key, value interface{}) bool {
        fmt.Printf("%v: %v\n", key, value)
        return true // 継続する場合はtrue
    })
}
```

## 自前実装によるハッシュテーブル

### パターン 1: スライスを使用したチェイン法

最も実装しやすく、理解しやすい方法です。

#### 特徴

- **実装の簡単さ**: スライスの append を活用して簡潔に実装可能
- **動的サイズ**: 各バケットが動的に拡張される
- **削除の容易さ**: スライスの要素削除は比較的簡単

#### 注意点

- **メモリ断片化**: スライスの拡張時に新しいメモリ領域が確保される
- **キャッシュ効率**: 連続メモリだがバケット間の局所性は低い
- **削除コスト**: 要素削除時にスライスの再配置が発生

```go
package main

import "fmt"

// キー・バリューペア
type KeyValue struct {
    Key   string
    Value int
}

// ハッシュテーブル構造体
type SliceHashTable struct {
    buckets [][]KeyValue
    size    int
    capacity int
}

// 新しいハッシュテーブルを作成
func NewSliceHashTable(capacity int) *SliceHashTable {
    return &SliceHashTable{
        buckets:  make([][]KeyValue, capacity),
        capacity: capacity,
    }
}

// ハッシュ関数（djb2アルゴリズム）
func (ht *SliceHashTable) hash(key string) int {
    hash := 5381
    for _, char := range key {
        hash = ((hash << 5) + hash) + int(char)
    }
    if hash < 0 {
        hash = -hash
    }
    return hash % ht.capacity
}

// 要素を挿入
func (ht *SliceHashTable) Put(key string, value int) {
    index := ht.hash(key)

    // 既存のキーを検索
    for i, kv := range ht.buckets[index] {
        if kv.Key == key {
            ht.buckets[index][i].Value = value // 値を更新
            return
        }
    }

    // 新しいキー・バリューペアを追加
    ht.buckets[index] = append(ht.buckets[index], KeyValue{Key: key, Value: value})
    ht.size++

    // 負荷率が高い場合はリハッシュ
    if float64(ht.size)/float64(ht.capacity) > 0.75 {
        ht.rehash()
    }
}

// 要素を取得
func (ht *SliceHashTable) Get(key string) (int, bool) {
    index := ht.hash(key)

    for _, kv := range ht.buckets[index] {
        if kv.Key == key {
            return kv.Value, true
        }
    }

    return 0, false
}

// 要素を削除
func (ht *SliceHashTable) Delete(key string) bool {
    index := ht.hash(key)

    for i, kv := range ht.buckets[index] {
        if kv.Key == key {
            // スライスから要素を削除
            ht.buckets[index] = append(ht.buckets[index][:i], ht.buckets[index][i+1:]...)
            ht.size--
            return true
        }
    }

    return false
}

// リハッシュ（容量を倍にして再配置）
func (ht *SliceHashTable) rehash() {
    oldBuckets := ht.buckets
    ht.capacity *= 2
    ht.buckets = make([][]KeyValue, ht.capacity)
    ht.size = 0

    // 全ての要素を再挿入
    for _, bucket := range oldBuckets {
        for _, kv := range bucket {
            ht.Put(kv.Key, kv.Value)
        }
    }
}

func main() {
    ht := NewSliceHashTable(4)

    ht.Put("apple", 100)
    ht.Put("banana", 200)
    ht.Put("orange", 300)

    if value, exists := ht.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }

    fmt.Printf("delete banana: %t\n", ht.Delete("banana"))
    fmt.Printf("size: %d, capacity: %d\n", ht.size, ht.capacity)
}
```

### パターン 2: 連結リストを使用したチェイン法

メモリ効率を重視した実装方法です。

#### 特徴

- **メモリ効率**: 必要な分だけメモリを使用
- **動的サイズ**: ノード単位での柔軟なサイズ変更
- **削除の効率性**: ポインタ操作のみで削除可能

#### 注意点

- **キャッシュ効率の低下**: メモリが非連続で参照の局所性が低い
- **メモリオーバーヘッド**: 各ノードにポインタが必要
- **実装の複雑さ**: ポインタ操作による複雑性

```go
package main

import "fmt"

// ノード構造体
type Node struct {
    Key   string
    Value int
    Next  *Node
}

// ハッシュテーブル構造体
type LinkedHashTable struct {
    buckets  []*Node
    size     int
    capacity int
}

// 新しいハッシュテーブルを作成
func NewLinkedHashTable(capacity int) *LinkedHashTable {
    return &LinkedHashTable{
        buckets:  make([]*Node, capacity),
        capacity: capacity,
    }
}

// ハッシュ関数
func (ht *LinkedHashTable) hash(key string) int {
    hash := 0
    for _, char := range key {
        hash = hash*31 + int(char)
    }
    if hash < 0 {
        hash = -hash
    }
    return hash % ht.capacity
}

// 要素を挿入
func (ht *LinkedHashTable) Put(key string, value int) {
    index := ht.hash(key)

    // 既存のキーを検索
    current := ht.buckets[index]
    for current != nil {
        if current.Key == key {
            current.Value = value // 値を更新
            return
        }
        current = current.Next
    }

    // 新しいノードを先頭に追加
    newNode := &Node{Key: key, Value: value, Next: ht.buckets[index]}
    ht.buckets[index] = newNode
    ht.size++

    // 負荷率チェック
    if float64(ht.size)/float64(ht.capacity) > 0.75 {
        ht.rehash()
    }
}

// 要素を取得
func (ht *LinkedHashTable) Get(key string) (int, bool) {
    index := ht.hash(key)

    current := ht.buckets[index]
    for current != nil {
        if current.Key == key {
            return current.Value, true
        }
        current = current.Next
    }

    return 0, false
}

// 要素を削除
func (ht *LinkedHashTable) Delete(key string) bool {
    index := ht.hash(key)

    if ht.buckets[index] == nil {
        return false
    }

    // 先頭ノードが対象の場合
    if ht.buckets[index].Key == key {
        ht.buckets[index] = ht.buckets[index].Next
        ht.size--
        return true
    }

    // 中間・末尾ノードを検索
    current := ht.buckets[index]
    for current.Next != nil {
        if current.Next.Key == key {
            current.Next = current.Next.Next
            ht.size--
            return true
        }
        current = current.Next
    }

    return false
}

// リハッシュ
func (ht *LinkedHashTable) rehash() {
    oldBuckets := ht.buckets
    ht.capacity *= 2
    ht.buckets = make([]*Node, ht.capacity)
    ht.size = 0

    // 全ての要素を再挿入
    for _, head := range oldBuckets {
        current := head
        for current != nil {
            ht.Put(current.Key, current.Value)
            current = current.Next
        }
    }
}

func main() {
    ht := NewLinkedHashTable(4)

    ht.Put("apple", 100)
    ht.Put("banana", 200)

    if value, exists := ht.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }
}
```

### パターン 3: オープンアドレス法（線形探査）

配列の連続性を活かした実装方法です。

#### 特徴

- **キャッシュ効率**: 連続したメモリアクセスで高いキャッシュ効率
- **メモリ効率**: ポインタが不要で省メモリ
- **シンプルな構造**: 配列のみを使用

#### 注意点

- **削除の複雑さ**: 削除マーカーや再配置が必要
- **クラスタリング**: 要素が密集して性能劣化の原因となる
- **負荷率制限**: 高い負荷率で急激に性能が劣化

```go
package main

import "fmt"

// エントリ状態
type EntryState int

const (
    Empty EntryState = iota
    Occupied
    Deleted
)

// エントリ構造体
type Entry struct {
    Key   string
    Value int
    State EntryState
}

// ハッシュテーブル構造体
type OpenAddressHashTable struct {
    entries  []Entry
    size     int
    capacity int
    deleted  int // 削除済みエントリ数
}

// 新しいハッシュテーブルを作成
func NewOpenAddressHashTable(capacity int) *OpenAddressHashTable {
    return &OpenAddressHashTable{
        entries:  make([]Entry, capacity),
        capacity: capacity,
    }
}

// ハッシュ関数
func (ht *OpenAddressHashTable) hash(key string) int {
    hash := 0
    for _, char := range key {
        hash = hash*31 + int(char)
    }
    if hash < 0 {
        hash = -hash
    }
    return hash % ht.capacity
}

// 線形探査で適切なスロットを探す
func (ht *OpenAddressHashTable) findSlot(key string, forInsertion bool) int {
    index := ht.hash(key)
    deletedIndex := -1

    for i := 0; i < ht.capacity; i++ {
        currentIndex := (index + i) % ht.capacity
        entry := &ht.entries[currentIndex]

        if entry.State == Empty {
            if forInsertion && deletedIndex != -1 {
                return deletedIndex // 削除済みスロットを再利用
            }
            return currentIndex
        }

        if entry.State == Deleted && forInsertion && deletedIndex == -1 {
            deletedIndex = currentIndex
        }

        if entry.State == Occupied && entry.Key == key {
            return currentIndex
        }
    }

    if forInsertion && deletedIndex != -1 {
        return deletedIndex
    }

    return -1 // 見つからない
}

// 要素を挿入
func (ht *OpenAddressHashTable) Put(key string, value int) bool {
    // 負荷率チェック（削除済みエントリも考慮）
    if float64(ht.size+ht.deleted)/float64(ht.capacity) > 0.7 {
        ht.rehash()
    }

    index := ht.findSlot(key, true)
    if index == -1 {
        return false
    }

    entry := &ht.entries[index]
    if entry.State == Occupied {
        // 既存キーの値を更新
        entry.Value = value
    } else {
        // 新しいエントリを追加
        if entry.State == Deleted {
            ht.deleted--
        }
        entry.Key = key
        entry.Value = value
        entry.State = Occupied
        ht.size++
    }

    return true
}

// 要素を取得
func (ht *OpenAddressHashTable) Get(key string) (int, bool) {
    index := ht.findSlot(key, false)
    if index == -1 || ht.entries[index].State != Occupied {
        return 0, false
    }

    return ht.entries[index].Value, true
}

// 要素を削除
func (ht *OpenAddressHashTable) Delete(key string) bool {
    index := ht.findSlot(key, false)
    if index == -1 || ht.entries[index].State != Occupied {
        return false
    }

    ht.entries[index].State = Deleted
    ht.size--
    ht.deleted++

    return true
}

// リハッシュ
func (ht *OpenAddressHashTable) rehash() {
    oldEntries := ht.entries
    ht.capacity *= 2
    ht.entries = make([]Entry, ht.capacity)
    ht.size = 0
    ht.deleted = 0

    // 有効なエントリを再挿入
    for _, entry := range oldEntries {
        if entry.State == Occupied {
            ht.Put(entry.Key, entry.Value)
        }
    }
}

func main() {
    ht := NewOpenAddressHashTable(8)

    ht.Put("apple", 100)
    ht.Put("banana", 200)
    ht.Put("orange", 300)

    if value, exists := ht.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }

    fmt.Printf("delete banana: %t\n", ht.Delete("banana"))
}
```

### パターン 4: オープンアドレス法（二次探査）

線形探査のクラスタリング問題を軽減した実装です。

#### 特徴

- **クラスタリング軽減**: 線形探査より分散が改善される
- **キャッシュ効率**: 連続メモリアクセスの利点を維持
- **探査パターンの改善**: より均等な分散

#### 注意点

- **実装の複雑さ**: 探査関数が複雑になる
- **テーブルサイズ制限**: 特定のサイズでないと全スロットを探査できない
- **削除の問題**: 線形探査と同様の削除問題が存在

```go
package main

import "fmt"

// エントリ構造体（オープンアドレス法用）
type QuadraticEntry struct {
    Key   string
    Value int
    State EntryState
}

// ハッシュテーブル構造体（二次探査）
type QuadraticHashTable struct {
    entries  []QuadraticEntry
    size     int
    capacity int
}

// 新しいハッシュテーブルを作成（容量は2の累乗にする）
func NewQuadraticHashTable(capacity int) *QuadraticHashTable {
    // 容量を2の累乗に調整
    actualCapacity := 1
    for actualCapacity < capacity {
        actualCapacity *= 2
    }

    return &QuadraticHashTable{
        entries:  make([]QuadraticEntry, actualCapacity),
        capacity: actualCapacity,
    }
}

// ハッシュ関数
func (ht *QuadraticHashTable) hash(key string) int {
    hash := 0
    for _, char := range key {
        hash = hash*31 + int(char)
    }
    if hash < 0 {
        hash = -hash
    }
    return hash % ht.capacity
}

// 二次探査でスロットを探す
func (ht *QuadraticHashTable) findSlot(key string, forInsertion bool) int {
    index := ht.hash(key)

    for i := 0; i < ht.capacity; i++ {
        // 二次探査: h(k) + i²
        currentIndex := (index + i*i) % ht.capacity
        entry := &ht.entries[currentIndex]

        if entry.State == Empty || (entry.State == Deleted && forInsertion) {
            return currentIndex
        }

        if entry.State == Occupied && entry.Key == key {
            return currentIndex
        }
    }

    return -1
}

// 要素を挿入
func (ht *QuadraticHashTable) Put(key string, value int) bool {
    if float64(ht.size)/float64(ht.capacity) > 0.5 {
        ht.rehash()
    }

    index := ht.findSlot(key, true)
    if index == -1 {
        return false
    }

    entry := &ht.entries[index]
    if entry.State != Occupied {
        ht.size++
    }

    entry.Key = key
    entry.Value = value
    entry.State = Occupied

    return true
}

// 要素を取得
func (ht *QuadraticHashTable) Get(key string) (int, bool) {
    index := ht.findSlot(key, false)
    if index == -1 || ht.entries[index].State != Occupied {
        return 0, false
    }

    return ht.entries[index].Value, true
}

// リハッシュ
func (ht *QuadraticHashTable) rehash() {
    oldEntries := ht.entries
    ht.capacity *= 2
    ht.entries = make([]QuadraticEntry, ht.capacity)
    ht.size = 0

    for _, entry := range oldEntries {
        if entry.State == Occupied {
            ht.Put(entry.Key, entry.Value)
        }
    }
}

func main() {
    ht := NewQuadraticHashTable(8)

    ht.Put("apple", 100)
    ht.Put("banana", 200)

    if value, exists := ht.Get("apple"); exists {
        fmt.Printf("apple: %d\n", value)
    }
}
```

## 実装方法別メリット・デメリット比較

| 実装方法                 | メリット                                                                                                                                                   | デメリット                                                                                                                                                      |
| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Go の組み込み map**    | - 最も簡潔で高性能<br>- Go runtime による最適化済み<br>- 自動メモリ管理<br>- 型安全性<br>- 標準ライブラリのサポート<br>- 豊富なドキュメントと事例          | - 内部実装をカスタマイズできない<br>- 並行安全性がない（sync.Map は必要）<br>- 反復順序が保証されない<br>- 特殊な要求（順序保持、カスタムハッシュ等）に対応困難 |
| **sync.Map（並行安全）** | - 並行アクセス安全<br>- ロックフリーな読み取り最適化<br>- Go runtime による最適化<br>- 型安全性（interface{}だが実行時チェック）                           | - 型安全性が弱い（interface{}使用）<br>- 通常の map より重い<br>- 書き込み頻度が高い場合に性能劣化<br>- メモリ使用量が多い                                      |
| **スライスチェイン法**   | - 実装が簡単で理解しやすい<br>- 動的サイズで柔軟性が高い<br>- Go の標準機能（append）を活用<br>- 削除操作が比較的簡単<br>- デバッグが容易                  | - メモリ断片化が発生しやすい<br>- スライス拡張時のコピーコスト<br>- 削除時の要素シフトコスト<br>- キャッシュ効率が中程度                                        |
| **連結リストチェイン法** | - メモリ効率が良い（必要な分のみ使用）<br>- 動的サイズ変更が柔軟<br>- 削除がポインタ操作のみで高速<br>- クラスタリング問題なし<br>- 理論的に無制限の要素数 | - キャッシュ効率が低い（メモリが非連続）<br>- 各ノードのポインタオーバーヘッド<br>- 実装がやや複雑<br>- メモリアクセスパターンが予測困難                        |
| **線形探査法**           | - 最高のキャッシュ効率（連続メモリアクセス）<br>- メモリ使用量が最小（ポインタ不要）<br>- シンプルな実装<br>- 高い局所性による高速アクセス                 | - クラスタリング問題で性能劣化<br>- 削除が複雑（削除マーカー必要）<br>- 負荷率制限が厳しい（0.7 以下推奨）<br>- 最悪時の性能劣化が大きい                        |
| **二次探査法**           | - 線形探査よりクラスタリングが少ない<br>- キャッシュ効率を維持<br>- メモリ使用量が少ない<br>- より均等な分散                                               | - 実装が複雑<br>- テーブルサイズに制限（2 の累乗等）<br>- 削除が複雑<br>- 完全なクラスタリング解決は不可                                                        |
| **二重ハッシュ法**       | - 最も均等な分散<br>- クラスタリング問題をほぼ解決<br>- 理論的に最適な探査パターン<br>- キャッシュ効率を維持                                               | - 実装が最も複雑<br>- 二つのハッシュ関数が必要<br>- 計算オーバーヘッドが大きい<br>- デバッグが困難<br>- ハッシュ関数の品質に依存                                |

### 使用場面別推奨実装

| 使用場面                       | 推奨実装                      | 理由                                 |
| ------------------------------ | ----------------------------- | ------------------------------------ |
| **一般的なアプリケーション**   | Go の組み込み map             | 簡潔、高性能、保守性が高い           |
| **並行アクセスが必要**         | sync.Map                      | 並行安全性が保証される               |
| **学習・教育目的**             | スライスチェイン法            | 理解しやすく、実装が簡単             |
| **メモリ制約が厳しい**         | 線形探査法                    | 最小のメモリ使用量                   |
| **高頻度の削除操作**           | 連結リストチェイン法          | 削除操作が効率的                     |
| **最高のパフォーマンスが必要** | 線形探査法（低負荷率）        | 最高のキャッシュ効率                 |
| **大量データの処理**           | 二次探査法または二重ハッシュ  | クラスタリングを避けて安定した性能   |
| **リアルタイムシステム**       | 組み込み map または線形探査法 | 予測可能な性能特性                   |
| **カスタムハッシュ関数が必要** | 自前実装                      | ハッシュ関数をカスタマイズ可能       |
| **順序保持が必要**             | 別途順序管理 + map            | map と順序管理用の構造体を組み合わせ |

# ハッシュテーブルの応用例

1. **データベースインデックス**: 高速なレコード検索
2. **キャッシュシステム**: Web ブラウザやアプリケーションのキャッシュ
3. **辞書・マップ**: プログラミング言語の連想配列
4. **セット実装**: 重複排除やメンバーシップテスト
5. **ルーティングテーブル**: ネットワークルーティング
6. **シンボルテーブル**: コンパイラでの変数・関数管理
