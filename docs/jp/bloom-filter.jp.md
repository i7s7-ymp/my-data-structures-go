# Bloom Filter（ブルームフィルター）とは

Bloom Filter（ブルームフィルター）は、**確率的データ構造**の一種で、ある要素が集合に含まれているかどうかを効率的に判定するためのデータ構造です。Burton Howard Bloom が 1970 年に提案したこの構造は、メモリ効率を重視し、偽陽性（false positive）は許容するが偽陰性（false negative）は発生しない特徴を持ちます。

**このデータ構造の最大の利点は、非常に少ないメモリで大量の要素の存在確認を高速に行えることと、要素数に関係なく一定の性能を保証する点です。**

# Bloom Filter の構成要素

## 基本要素

- **ビット配列（Bit Array）**: 固定サイズのビット列、すべて 0 で初期化
- **ハッシュ関数群（Hash Functions）**: 複数の独立したハッシュ関数
- **フィルターサイズ（m）**: ビット配列のサイズ（ビット数）
- **ハッシュ関数数（k）**: 使用するハッシュ関数の数
- **要素数（n）**: 挿入される要素の予想数
- **偽陽性率（p）**: 許容する偽陽性の確率

## Bloom Filter の性質

1. **偽陰性なし**: 要素が存在しないのに「存在する」と判定することはない
2. **偽陽性あり**: 要素が存在しないのに「存在する」と判定する可能性がある
3. **削除不可**: 標準的な Bloom Filter では要素の削除はできない
4. **固定サイズ**: ビット配列のサイズは事前に決定される
5. **追加のみ**: 要素の追加のみ可能（削除は不可）

# Bloom Filter の種類

## Standard Bloom Filter

- 基本的な Bloom Filter
- 最もシンプルな実装
- 用途: 一般的な存在確認

## Counting Bloom Filter

- 各ビットの代わりにカウンタを使用
- 要素の削除が可能
- 用途: 動的な集合の管理

## Scalable Bloom Filter

- 容量が不足したら新しいフィルターを追加
- 動的にサイズを拡張可能
- 用途: 要素数が予測困難な場合

## Compressed Bloom Filter

- ビット配列を圧縮して格納
- メモリ使用量をさらに削減
- 用途: 極限のメモリ制約環境

## Cuckoo Filter

- Bloom Filter の改良版
- 削除操作をサポート
- 用途: 削除が必要な場合

## Quotient Filter

- 商（quotient）を利用した確率的フィルター
- 削除操作とマージ操作をサポート
- 用途: 高度な集合操作が必要な場合

# Bloom Filter の特徴

- **メモリ効率**: 要素数に対して非常に少ないメモリ使用量
- **高速処理**: O(k)の一定時間での操作（k はハッシュ関数数）
- **スケーラビリティ**: 要素数が増えても性能は一定
- **確率的**: 100%の正確性ではなく確率的な判定
- **キャッシュ効率**: 連続したメモリアクセスパターン

# Bloom Filter の構成要素の図解

## 基本的な Bloom Filter の構造

```
要素集合: {"apple", "banana", "cherry"}
ビット配列サイズ: m = 16
ハッシュ関数数: k = 3

初期状態（すべて0）:
Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Bits:  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0

"apple"を挿入:
hash1("apple") = 3  → bits[3] = 1
hash2("apple") = 7  → bits[7] = 1
hash3("apple") = 12 → bits[12] = 1

Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Bits:  0  0  0  1  0  0  0  1  0  0  0  0  1  0  0  0

"banana"を挿入:
hash1("banana") = 2  → bits[2] = 1
hash2("banana") = 7  → bits[7] = 1 (既に1)
hash3("banana") = 14 → bits[14] = 1

Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Bits:  0  0  1  1  0  0  0  1  0  0  0  0  1  0  1  0

"cherry"を挿入:
hash1("cherry") = 1  → bits[1] = 1
hash2("cherry") = 8  → bits[8] = 1
hash3("cherry") = 12 → bits[12] = 1 (既に1)

最終状態:
Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Bits:  0  1  1  1  0  0  0  1  1  0  0  0  1  0  1  0
```

## 検索操作の例

```
"apple"の存在確認:
hash1("apple") = 3  → bits[3] = 1 ✓
hash2("apple") = 7  → bits[7] = 1 ✓
hash3("apple") = 12 → bits[12] = 1 ✓
結果: 存在する可能性が高い（True）

"grape"の存在確認:
hash1("grape") = 1  → bits[1] = 1 ✓
hash2("grape") = 5  → bits[5] = 0 ✗
hash3("grape") = 10 → bits[10] = 0 ✗
結果: 存在しない（False）

"orange"の存在確認（偽陽性の例）:
hash1("orange") = 2  → bits[2] = 1 ✓
hash2("orange") = 7  → bits[7] = 1 ✓
hash3("orange") = 14 → bits[14] = 1 ✓
結果: 存在する可能性が高い（False Positive）
```

## 偽陽性率の計算

```
偽陽性率の理論値:
p = (1 - e^(-kn/m))^k

ここで:
- k: ハッシュ関数数
- n: 挿入された要素数
- m: ビット配列のサイズ

最適なハッシュ関数数:
k = (m/n) * ln(2)

例（m=1000, n=100の場合）:
最適k = (1000/100) * ln(2) ≈ 6.93 → k=7
偽陽性率 p ≈ 0.0081 (約0.81%)
```

## Counting Bloom Filter の構造

```
標準Bloom Filter:
Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Bits:  0  1  1  1  0  0  0  1  1  0  0  0  1  0  1  0

Counting Bloom Filter:
Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
Count: 0  1  1  1  0  0  0  2  1  0  0  0  2  0  1  0

特徴:
- 各位置にカウンタ（通常4-8ビット）
- 要素追加時にカウンタをインクリメント
- 要素削除時にカウンタをデクリメント
- カウンタが0になったらその位置は未使用
```

# Bloom Filter で行える処理

## 基本操作

| 機能              | 説明                 | 計算量 | 戻り値                          | 得意/苦手 | 補足                                 |
| ----------------- | -------------------- | ------ | ------------------------------- | --------- | ------------------------------------ |
| Add               | 要素を追加           | O(k)   | なし（void）                    | ✅        | k 個のハッシュ値を計算してビット設定 |
| Contains          | 要素の存在確認       | O(k)   | bool（存在可能性/確実に非存在） | ✅        | Bloom Filter の主要機能              |
| Size              | フィルターサイズ取得 | O(1)   | int（ビット数）                 | ✅        | 固定サイズのため定数時間             |
| IsEmpty           | 空かどうか確認       | O(m)   | bool（空/非空）                 | ❌        | 全ビットをスキャンする必要           |
| Clear             | 全ビットをクリア     | O(m)   | なし（void）                    | ✅        | ビット配列を 0 で初期化              |
| EstimatedCount    | 推定要素数を計算     | O(m)   | int（推定要素数）               | ⚠️        | 数式による近似計算                   |
| FalsePositiveRate | 現在の偽陽性率       | O(m)   | float（偽陽性率）               | ⚠️        | 設定ビット数から推定                 |
| Capacity          | 推奨最大要素数       | O(1)   | int（最大要素数）               | ✅        | 設計時に決定された値                 |
| HashCount         | ハッシュ関数数       | O(1)   | int（関数数）                   | ✅        | 設計パラメータ                       |
| BitCount          | 設定済みビット数     | O(m)   | int（ビット数）                 | ⚠️        | 全ビットをカウント                   |

## 集合演算（特殊）

| 機能         | 説明               | 計算量 | 戻り値                    | 得意/苦手 | 補足                          |
| ------------ | ------------------ | ------ | ------------------------- | --------- | ----------------------------- |
| Union        | 和集合             | O(m)   | 新しい Bloom Filter       | ✅        | ビット単位の OR 演算          |
| Intersection | 積集合             | O(m)   | 新しい Bloom Filter       | ⚠️        | ビット単位の AND 演算（近似） |
| IsSubset     | 部分集合判定       | O(m)   | bool（部分集合可能性/否） | ⚠️        | 近似的な判定のみ              |
| Merge        | フィルターの結合   | O(m)   | なし（自身を更新）        | ✅        | 他のフィルターとの OR 演算    |
| Clone        | フィルターの複製   | O(m)   | 新しい Bloom Filter       | ✅        | ビット配列の完全コピー        |
| Equals       | フィルターの等価性 | O(m)   | bool（等価/非等価）       | ✅        | ビット配列の完全比較          |
| Serialize    | シリアライズ       | O(m)   | バイト配列                | ✅        | ネットワーク転送やストレージ  |
| Deserialize  | デシリアライズ     | O(m)   | Bloom Filter              | ✅        | バイト配列からの復元          |

## 統計・分析操作

| 機能                  | 説明                 | 計算量  | 戻り値            | 得意/苦手 | 補足                           |
| --------------------- | -------------------- | ------- | ----------------- | --------- | ------------------------------ |
| GetOptimalHashCount   | 最適ハッシュ関数数   | O(1)    | int（最適 k 値）  | ✅        | 理論値 k = (m/n)ln(2)          |
| GetOptimalSize        | 最適フィルターサイズ | O(1)    | int（最適 m 値）  | ✅        | 理論値 m = -n\*ln(p)/(ln(2)^2) |
| GetCurrentUtilization | 現在の使用率         | O(m)    | float（使用率%）  | ⚠️        | 設定ビット数/総ビット数        |
| EstimateFPR           | 現在の偽陽性率推定   | O(m)    | float（推定 FPR） | ⚠️        | 現在の状態から計算             |
| GetMemoryUsage        | メモリ使用量         | O(1)    | int（バイト数）   | ✅        | m/8 + オーバーヘッド           |
| Validate              | パラメータ妥当性検証 | O(1)    | bool（妥当/不正） | ✅        | 設計パラメータの整合性確認     |
| GetStatistics         | 統計情報取得         | O(m)    | 統計構造体        | ⚠️        | 包括的な状態情報               |
| Benchmark             | 性能測定             | O(n\*k) | ベンチマーク結果  | ✅        | 実際の性能測定                 |

## 高度な操作（拡張版）

| 機能        | 説明                 | 計算量  | 戻り値                        | 得意/苦手 | 補足                       |
| ----------- | -------------------- | ------- | ----------------------------- | --------- | -------------------------- |
| Remove      | 要素削除             | O(k)    | bool（成功/失敗）             | ❌        | Counting Bloom Filter 必要 |
| AddAll      | 複数要素一括追加     | O(n\*k) | int（追加数）                 | ✅        | バッチ処理で効率化         |
| ContainsAll | 複数要素存在確認     | O(n\*k) | bool（全て存在可能性/否）     | ✅        | 短絡評価で最適化           |
| ContainsAny | いずれか要素存在確認 | O(n\*k) | bool（いずれか存在可能性/否） | ✅        | 短絡評価で最適化           |
| Resize      | フィルターサイズ変更 | O(m)    | 新しい Bloom Filter           | ❌        | 全要素の再挿入が必要       |
| Optimize    | パラメータ最適化     | O(m)    | 最適化済みフィルター          | ⚠️        | 現在の状態に基づく最適化   |
| Compress    | フィルター圧縮       | O(m)    | 圧縮済みフィルター            | ⚠️        | 特殊な圧縮アルゴリズム     |
| Decompress  | フィルター解凍       | O(m)    | 解凍済みフィルター            | ⚠️        | 圧縮の逆操作               |

## 苦手な処理・制限

| 機能            | 説明                 | 計算量 | 戻り値 | 得意/苦手 | 補足                         |
| --------------- | -------------------- | ------ | ------ | --------- | ---------------------------- |
| Delete          | 要素削除             | -      | -      | ❌        | 標準版では不可能             |
| ListElements    | 要素列挙             | -      | -      | ❌        | 要素の実際の値は保存されない |
| Count           | 正確な要素数         | -      | -      | ❌        | 推定のみ可能                 |
| Iteration       | 要素の反復処理       | -      | -      | ❌        | 要素の実体が保存されない     |
| IndexAccess     | インデックスアクセス | -      | -      | ❌        | 順序概念が存在しない         |
| RangeQuery      | 範囲クエリ           | -      | -      | ❌        | 順序関係を保持しない         |
| ExactMembership | 確実な存在確認       | -      | -      | ❌        | 偽陽性が発生する             |
| Update          | 要素の更新           | -      | -      | ❌        | 要素の実体を保存しない       |
| GetElement      | 要素の取得           | -      | -      | ❌        | ハッシュ値のみ保存           |
| Sort            | 要素のソート         | -      | -      | ❌        | 順序概念が存在しない         |

# Bloom Filter の実装方法

## 一般的な実装方法

### ビット配列ベース実装

- 固定サイズのビット配列を使用
- 各ビットは 1 つの位置を表現
- メモリ効率が最高だが、削除は不可能

### カウンタベース実装（Counting Bloom Filter）

- ビットの代わりにカウンタを使用
- 削除操作が可能
- メモリ使用量が増加（通常 4-8 倍）

### ハッシュ関数の選択

- 独立性の高い複数のハッシュ関数
- MD5、SHA-1、MurmurHash、xxHash など
- ダブルハッシュによる擬似的な複数ハッシュ関数生成

## Go による基本的な Bloom Filter 実装例

```go
package bloomfilter

import (
    "crypto/sha1"
    "crypto/sha256"
    "encoding/binary"
    "hash"
    "hash/fnv"
    "math"
)

// Bloom Filter構造体
type BloomFilter struct {
    bitArray    []bool  // ビット配列
    size        uint    // ビット配列のサイズ
    hashCount   uint    // ハッシュ関数の数
    elementCount uint    // 追加された要素数（推定）
}

// 新しいBloom Filterを作成
func New(expectedElements uint, falsePositiveRate float64) *BloomFilter {
    size := optimalSize(expectedElements, falsePositiveRate)
    hashCount := optimalHashCount(size, expectedElements)

    return &BloomFilter{
        bitArray:    make([]bool, size),
        size:        size,
        hashCount:   hashCount,
        elementCount: 0,
    }
}

// 指定パラメータでBloom Filterを作成
func NewWithParameters(size, hashCount uint) *BloomFilter {
    return &BloomFilter{
        bitArray:    make([]bool, size),
        size:        size,
        hashCount:   hashCount,
        elementCount: 0,
    }
}

// 最適なビット配列サイズを計算
func optimalSize(n uint, p float64) uint {
    // m = -n * ln(p) / (ln(2)^2)
    size := -float64(n) * math.Log(p) / (math.Log(2) * math.Log(2))
    return uint(math.Ceil(size))
}

// 最適なハッシュ関数数を計算
func optimalHashCount(m, n uint) uint {
    // k = (m/n) * ln(2)
    k := float64(m) / float64(n) * math.Log(2)
    return uint(math.Round(k))
}

// 複数のハッシュ値を生成（ダブルハッシュ法）
func (bf *BloomFilter) hash(data []byte) []uint {
    hashes := make([]uint, bf.hashCount)

    // 2つの基本ハッシュ値を計算
    h1 := bf.hash1(data)
    h2 := bf.hash2(data)

    // ダブルハッシュ法で複数のハッシュ値を生成
    for i := uint(0); i < bf.hashCount; i++ {
        hash := (h1 + uint(i)*h2) % bf.size
        hashes[i] = hash
    }

    return hashes
}

// 第1ハッシュ関数（FNV-1a）
func (bf *BloomFilter) hash1(data []byte) uint {
    h := fnv.New32a()
    h.Write(data)
    return uint(h.Sum32()) % bf.size
}

// 第2ハッシュ関数（SHA-1の最初の4バイト）
func (bf *BloomFilter) hash2(data []byte) uint {
    h := sha1.Sum(data)
    hash := binary.BigEndian.Uint32(h[:4])
    return uint(hash) % bf.size
}

// 要素を追加
func (bf *BloomFilter) Add(data []byte) {
    hashes := bf.hash(data)
    for _, hash := range hashes {
        bf.bitArray[hash] = true
    }
    bf.elementCount++
}

// 文字列を追加
func (bf *BloomFilter) AddString(str string) {
    bf.Add([]byte(str))
}

// 要素の存在確認
func (bf *BloomFilter) Contains(data []byte) bool {
    hashes := bf.hash(data)
    for _, hash := range hashes {
        if !bf.bitArray[hash] {
            return false // 確実に存在しない
        }
    }
    return true // 存在する可能性が高い
}

// 文字列の存在確認
func (bf *BloomFilter) ContainsString(str string) bool {
    return bf.Contains([]byte(str))
}

// 複数要素の存在確認
func (bf *BloomFilter) ContainsAll(data [][]byte) bool {
    for _, item := range data {
        if !bf.Contains(item) {
            return false
        }
    }
    return true
}

// いずれかの要素が存在するか確認
func (bf *BloomFilter) ContainsAny(data [][]byte) bool {
    for _, item := range data {
        if bf.Contains(item) {
            return true
        }
    }
    return false
}

// フィルターサイズを取得
func (bf *BloomFilter) Size() uint {
    return bf.size
}

// ハッシュ関数数を取得
func (bf *BloomFilter) HashCount() uint {
    return bf.hashCount
}

// 追加された要素数を取得（推定）
func (bf *BloomFilter) Count() uint {
    return bf.elementCount
}

// 現在の偽陽性率を計算
func (bf *BloomFilter) FalsePositiveRate() float64 {
    // p = (1 - e^(-kn/m))^k
    ratio := float64(bf.hashCount) * float64(bf.elementCount) / float64(bf.size)
    return math.Pow(1.0-math.Exp(-ratio), float64(bf.hashCount))
}

// 設定されているビット数をカウント
func (bf *BloomFilter) BitCount() uint {
    count := uint(0)
    for _, bit := range bf.bitArray {
        if bit {
            count++
        }
    }
    return count
}

// フィルターの使用率を計算
func (bf *BloomFilter) Utilization() float64 {
    return float64(bf.BitCount()) / float64(bf.size)
}

// 推定要素数を計算（ビット数から逆算）
func (bf *BloomFilter) EstimatedCount() uint {
    setBits := bf.BitCount()
    if setBits == 0 {
        return 0
    }

    // n ≈ -m * ln(1 - X/m) / k
    ratio := float64(setBits) / float64(bf.size)
    if ratio >= 1.0 {
        return ^uint(0) // オーバーフロー
    }

    estimated := -float64(bf.size) * math.Log(1.0-ratio) / float64(bf.hashCount)
    return uint(math.Round(estimated))
}

// フィルターをクリア
func (bf *BloomFilter) Clear() {
    for i := range bf.bitArray {
        bf.bitArray[i] = false
    }
    bf.elementCount = 0
}

// 空かどうか確認
func (bf *BloomFilter) IsEmpty() bool {
    return bf.BitCount() == 0
}

// 和集合（Union）
func (bf *BloomFilter) Union(other *BloomFilter) (*BloomFilter, error) {
    if bf.size != other.size || bf.hashCount != other.hashCount {
        return nil, fmt.Errorf("incompatible bloom filters")
    }

    result := NewWithParameters(bf.size, bf.hashCount)
    for i := range bf.bitArray {
        result.bitArray[i] = bf.bitArray[i] || other.bitArray[i]
    }

    return result, nil
}

// 積集合（Intersection）
func (bf *BloomFilter) Intersection(other *BloomFilter) (*BloomFilter, error) {
    if bf.size != other.size || bf.hashCount != other.hashCount {
        return nil, fmt.Errorf("incompatible bloom filters")
    }

    result := NewWithParameters(bf.size, bf.hashCount)
    for i := range bf.bitArray {
        result.bitArray[i] = bf.bitArray[i] && other.bitArray[i]
    }

    return result, nil
}

// 部分集合判定（近似）
func (bf *BloomFilter) IsSubset(other *BloomFilter) bool {
    if bf.size != other.size || bf.hashCount != other.hashCount {
        return false
    }

    for i := range bf.bitArray {
        if bf.bitArray[i] && !other.bitArray[i] {
            return false
        }
    }
    return true
}

// 等価性判定
func (bf *BloomFilter) Equals(other *BloomFilter) bool {
    if bf.size != other.size || bf.hashCount != other.hashCount {
        return false
    }

    for i := range bf.bitArray {
        if bf.bitArray[i] != other.bitArray[i] {
            return false
        }
    }
    return true
}

// フィルターを複製
func (bf *BloomFilter) Clone() *BloomFilter {
    clone := NewWithParameters(bf.size, bf.hashCount)
    copy(clone.bitArray, bf.bitArray)
    clone.elementCount = bf.elementCount
    return clone
}

// メモリ使用量を取得（バイト）
func (bf *BloomFilter) MemoryUsage() uint {
    // ビット配列 + 構造体のオーバーヘッド
    return bf.size/8 + 32 // 概算
}

// シリアライズ（バイナリ形式）
func (bf *BloomFilter) Serialize() []byte {
    // ヘッダー: size(4) + hashCount(4) + elementCount(4)
    header := make([]byte, 12)
    binary.BigEndian.PutUint32(header[0:4], uint32(bf.size))
    binary.BigEndian.PutUint32(header[4:8], uint32(bf.hashCount))
    binary.BigEndian.PutUint32(header[8:12], uint32(bf.elementCount))

    // ビット配列を詰め込み
    bitBytes := make([]byte, (bf.size+7)/8)
    for i, bit := range bf.bitArray {
        if bit {
            byteIndex := i / 8
            bitIndex := i % 8
            bitBytes[byteIndex] |= 1 << bitIndex
        }
    }

    return append(header, bitBytes...)
}

// デシリアライズ
func Deserialize(data []byte) (*BloomFilter, error) {
    if len(data) < 12 {
        return nil, fmt.Errorf("invalid data: too short")
    }

    size := uint(binary.BigEndian.Uint32(data[0:4]))
    hashCount := uint(binary.BigEndian.Uint32(data[4:8]))
    elementCount := uint(binary.BigEndian.Uint32(data[8:12]))

    expectedBytes := (size + 7) / 8
    if uint(len(data)) != 12+expectedBytes {
        return nil, fmt.Errorf("invalid data: wrong size")
    }

    bf := NewWithParameters(size, hashCount)
    bf.elementCount = elementCount

    bitBytes := data[12:]
    for i := uint(0); i < size; i++ {
        byteIndex := i / 8
        bitIndex := i % 8
        if byteIndex < uint(len(bitBytes)) {
            bf.bitArray[i] = (bitBytes[byteIndex] & (1 << bitIndex)) != 0
        }
    }

    return bf, nil
}

// 統計情報構造体
type Statistics struct {
    Size                uint
    HashCount          uint
    ElementCount       uint
    BitCount           uint
    Utilization        float64
    FalsePositiveRate  float64
    EstimatedCount     uint
    MemoryUsage        uint
}

// 統計情報を取得
func (bf *BloomFilter) Statistics() Statistics {
    return Statistics{
        Size:              bf.size,
        HashCount:         bf.hashCount,
        ElementCount:      bf.elementCount,
        BitCount:          bf.BitCount(),
        Utilization:       bf.Utilization(),
        FalsePositiveRate: bf.FalsePositiveRate(),
        EstimatedCount:    bf.EstimatedCount(),
        MemoryUsage:       bf.MemoryUsage(),
    }
}

// 文字列表現
func (bf *BloomFilter) String() string {
    stats := bf.Statistics()
    return fmt.Sprintf("BloomFilter{size=%d, hashCount=%d, elements=%d, utilization=%.2f%%, fpr=%.4f}",
        stats.Size, stats.HashCount, stats.ElementCount, stats.Utilization*100, stats.FalsePositiveRate)
}
```

## Counting Bloom Filter の実装例

```go
package bloomfilter

import (
    "fmt"
    "math"
)

// Counting Bloom Filter構造体
type CountingBloomFilter struct {
    counters     []uint8  // カウンタ配列（通常4-8ビット）
    size         uint     // 配列のサイズ
    hashCount    uint     // ハッシュ関数の数
    elementCount uint     // 追加された要素数
    maxCounter   uint8    // カウンタの最大値
}

// 新しいCounting Bloom Filterを作成
func NewCountingBloomFilter(expectedElements uint, falsePositiveRate float64) *CountingBloomFilter {
    size := optimalSize(expectedElements, falsePositiveRate)
    hashCount := optimalHashCount(size, expectedElements)

    return &CountingBloomFilter{
        counters:     make([]uint8, size),
        size:         size,
        hashCount:    hashCount,
        elementCount: 0,
        maxCounter:   255, // uint8の最大値
    }
}

// ハッシュ値を生成（前述と同じ）
func (cbf *CountingBloomFilter) hash(data []byte) []uint {
    // 同じハッシュ関数を使用
    hashes := make([]uint, cbf.hashCount)
    h1 := cbf.hash1(data)
    h2 := cbf.hash2(data)

    for i := uint(0); i < cbf.hashCount; i++ {
        hash := (h1 + uint(i)*h2) % cbf.size
        hashes[i] = hash
    }

    return hashes
}

func (cbf *CountingBloomFilter) hash1(data []byte) uint {
    h := fnv.New32a()
    h.Write(data)
    return uint(h.Sum32()) % cbf.size
}

func (cbf *CountingBloomFilter) hash2(data []byte) uint {
    h := sha1.Sum(data)
    hash := binary.BigEndian.Uint32(h[:4])
    return uint(hash) % cbf.size
}

// 要素を追加
func (cbf *CountingBloomFilter) Add(data []byte) error {
    hashes := cbf.hash(data)

    // オーバーフローチェック
    for _, hash := range hashes {
        if cbf.counters[hash] >= cbf.maxCounter {
            return fmt.Errorf("counter overflow at position %d", hash)
        }
    }

    // カウンタをインクリメント
    for _, hash := range hashes {
        cbf.counters[hash]++
    }

    cbf.elementCount++
    return nil
}

// 要素を削除
func (cbf *CountingBloomFilter) Remove(data []byte) error {
    hashes := cbf.hash(data)

    // 削除可能性チェック
    for _, hash := range hashes {
        if cbf.counters[hash] == 0 {
            return fmt.Errorf("cannot remove: element not present")
        }
    }

    // カウンタをデクリメント
    for _, hash := range hashes {
        cbf.counters[hash]--
    }

    if cbf.elementCount > 0 {
        cbf.elementCount--
    }

    return nil
}

// 要素の存在確認
func (cbf *CountingBloomFilter) Contains(data []byte) bool {
    hashes := cbf.hash(data)
    for _, hash := range hashes {
        if cbf.counters[hash] == 0 {
            return false
        }
    }
    return true
}

// フィルターをクリア
func (cbf *CountingBloomFilter) Clear() {
    for i := range cbf.counters {
        cbf.counters[i] = 0
    }
    cbf.elementCount = 0
}

// 非ゼロカウンタ数を取得
func (cbf *CountingBloomFilter) NonZeroCount() uint {
    count := uint(0)
    for _, counter := range cbf.counters {
        if counter > 0 {
            count++
        }
    }
    return count
}

// 使用率を計算
func (cbf *CountingBloomFilter) Utilization() float64 {
    return float64(cbf.NonZeroCount()) / float64(cbf.size)
}

// 偽陽性率を計算
func (cbf *CountingBloomFilter) FalsePositiveRate() float64 {
    ratio := float64(cbf.hashCount) * float64(cbf.elementCount) / float64(cbf.size)
    return math.Pow(1.0-math.Exp(-ratio), float64(cbf.hashCount))
}

// メモリ使用量（カウンタは通常のビットより多い）
func (cbf *CountingBloomFilter) MemoryUsage() uint {
    return cbf.size * uint(unsafe.Sizeof(uint8(0))) + 32
}
```

## Scalable Bloom Filter の実装例

```go
package bloomfilter

// Scalable Bloom Filter構造体
type ScalableBloomFilter struct {
    filters    []*BloomFilter
    capacity   uint
    fpRate     float64
    growth     uint  // 容量の成長率
    maxFilters uint  // 最大フィルター数
}

// 新しいScalable Bloom Filterを作成
func NewScalableBloomFilter(initialCapacity uint, falsePositiveRate float64) *ScalableBloomFilter {
    filter := New(initialCapacity, falsePositiveRate)

    return &ScalableBloomFilter{
        filters:    []*BloomFilter{filter},
        capacity:   initialCapacity,
        fpRate:     falsePositiveRate,
        growth:     2, // 2倍ずつ成長
        maxFilters: 10, // 最大10個のフィルター
    }
}

// 要素を追加
func (sbf *ScalableBloomFilter) Add(data []byte) error {
    currentFilter := sbf.filters[len(sbf.filters)-1]

    // 現在のフィルターが満杯の場合、新しいフィルターを追加
    if currentFilter.Count() >= sbf.capacity {
        if len(sbf.filters) >= int(sbf.maxFilters) {
            return fmt.Errorf("maximum number of filters reached")
        }

        newCapacity := sbf.capacity * sbf.growth
        newFpRate := sbf.fpRate / float64(len(sbf.filters)+1) // FPRを分散
        newFilter := New(newCapacity, newFpRate)

        sbf.filters = append(sbf.filters, newFilter)
        sbf.capacity = newCapacity
        currentFilter = newFilter
    }

    currentFilter.Add(data)
    return nil
}

// 要素の存在確認
func (sbf *ScalableBloomFilter) Contains(data []byte) bool {
    for _, filter := range sbf.filters {
        if filter.Contains(data) {
            return true
        }
    }
    return false
}

// 全体の偽陽性率を計算
func (sbf *ScalableBloomFilter) FalsePositiveRate() float64 {
    // 複数フィルターでの累積FPR
    fpRate := 0.0
    for _, filter := range sbf.filters {
        fpRate += filter.FalsePositiveRate()
    }
    return fpRate
}

// 総要素数を取得
func (sbf *ScalableBloomFilter) TotalCount() uint {
    total := uint(0)
    for _, filter := range sbf.filters {
        total += filter.Count()
    }
    return total
}

// フィルター数を取得
func (sbf *ScalableBloomFilter) FilterCount() int {
    return len(sbf.filters)
}
```

# Bloom Filter の応用例

## 1. データベース・ストレージシステム

- **キーの存在確認**: 高価なディスクアクセス前の事前チェック
- **LSM-Tree**: LevelDB や Cassandra での SSTable の事前フィルタリング
- **分散データベース**: ノード間でのデータ所在確認
- **キャッシュシステム**: キャッシュミス削減のための事前判定

## 2. ネットワーク・CDN

- **URL 重複チェック**: Web Crawler での訪問済み URL 管理
- **CDN キャッシュ**: エッジサーバーでのコンテンツ存在確認
- **DDoS 対策**: 悪意ある IP アドレスの高速ブラックリスト
- **P2P ネットワーク**: ファイル共有での重複コンテンツ検出

## 3. ビッグデータ・アナリティクス

- **重複データ除去**: 大量データの前処理での重複排除
- **データストリーム**: リアルタイムストリームでの重複イベント検出
- **ログ解析**: 大量ログファイルでの特定パターン存在確認
- **ETL パイプライン**: データ変換時の効率的なフィルタリング

## 4. セキュリティ・脅威検出

- **マルウェア検出**: 既知の悪意あるファイルハッシュの高速チェック
- **フィッシング対策**: 危険 URL の高速ブラックリスト
- **スパムフィルタ**: スパムメールのシグネチャ検出
- **侵入検知**: ネットワークトラフィックの異常パターン検出

## 5. 推薦システム・機械学習

- **協調フィルタリング**: ユーザー間の共通アイテム存在確認
- **特徴選択**: 機械学習での特徴量の事前フィルタリング
- **データサンプリング**: 大量データからの効率的なサンプル抽出
- **A/B テスト**: ユーザーセグメントの高速判定

## 6. 検索エンジン・情報検索

- **重複ページ除去**: Web 検索での重複コンテンツフィルタリング
- **クエリ最適化**: 検索クエリの事前フィルタリング
- **インデックス最適化**: 検索インデックスの効率化
- **関連性判定**: 文書間の関連性事前チェック

## 7. ブロックチェーン・暗号通貨

- **トランザクション重複**: ブロックチェーンでの重複トランザクション検出
- **ウォレット管理**: アドレスの効率的な管理
- **マイニングプール**: 重複作業の回避
- **分散台帳**: ノード間でのデータ同期最適化

## 8. IoT・センサーデータ

- **センサーデータ重複**: IoT デバイスからの重複データフィルタリング
- **エッジコンピューティング**: 限られたリソースでの効率的処理
- **データ圧縮**: センサーデータの効率的な前処理
- **異常検知**: 既知の正常パターンとの高速比較

## 9. ゲーム・エンターテイメント

- **プレイヤー管理**: オンラインゲームでの重複アカウント検出
- **コンテンツ配信**: ゲームアセットの効率的な配信
- **チート検出**: 既知のチートパターンの高速検出
- **マッチメイキング**: プレイヤーマッチングの最適化

## 10. 金融・フィンテック

- **取引重複検出**: 金融取引での重複処理防止
- **リスク管理**: ブラックリスト顧客の高速スクリーニング
- **不正検知**: 既知の不正パターンとの高速照合
- **KYC/AML**: 顧客識別での効率的なスクリーニング

# Bloom Filter と他のデータ構造との比較

## 確率的データ構造比較

| 特徴               | Bloom Filter    | Cuckoo Filter | Skip List | Hash Table    | Set（Tree-based） |
| ------------------ | --------------- | ------------- | --------- | ------------- | ----------------- |
| **存在確認時間**   | O(k)            | O(1)          | O(log n)  | O(1)平均      | O(log n)          |
| **挿入時間**       | O(k)            | O(1)償却      | O(log n)  | O(1)平均      | O(log n)          |
| **削除時間**       | ❌ 不可         | O(1)償却      | O(log n)  | O(1)平均      | O(log n)          |
| **メモリ効率**     | ✅ 最高         | ✅ 高い       | ⚠️ 中程度 | ⚠️ 負荷率依存 | ❌ 低い           |
| **偽陽性**         | ✅ あり         | ✅ あり       | ❌ なし   | ❌ なし       | ❌ なし           |
| **偽陰性**         | ❌ なし         | ❌ なし       | ❌ なし   | ❌ なし       | ❌ なし           |
| **要素削除**       | ❌ 不可         | ✅ 可能       | ✅ 可能   | ✅ 可能       | ✅ 可能           |
| **要素列挙**       | ❌ 不可         | ❌ 不可       | ✅ 可能   | ✅ 可能       | ✅ 可能           |
| **動的サイズ**     | ❌ 固定         | ⚠️ 制限あり   | ✅ 動的   | ✅ 動的       | ✅ 動的           |
| **キャッシュ効率** | ✅ 高い         | ✅ 高い       | ❌ 低い   | ⚠️ 中程度     | ❌ 低い           |
| **実装複雑度**     | ✅ 簡単         | ⚠️ 中程度     | ⚠️ 中程度 | ✅ 簡単       | ⚠️ 中程度         |
| **並行処理**       | ✅ 読み取り安全 | ⚠️ 注意必要   | ❌ 困難   | ⚠️ 同期必要   | ❌ 困難           |

## 用途別最適選択

| 用途                             | 第 1 選択     | 第 2 選択             | 第 3 選択     | 避けるべき      |
| -------------------------------- | ------------- | --------------------- | ------------- | --------------- |
| **高速存在確認（読み取り専用）** | Bloom Filter  | Cuckoo Filter         | Hash Table    | Set（Tree）     |
| **削除が必要な確率的フィルタ**   | Cuckoo Filter | Counting Bloom Filter | Hash Table    | Bloom Filter    |
| **極限メモリ制約**               | Bloom Filter  | Cuckoo Filter         | -             | Hash Table、Set |
| **完全な正確性が必要**           | Hash Table    | Set（Tree）           | Skip List     | Bloom Filter    |
| **範囲クエリが必要**             | Set（Tree）   | Skip List             | -             | Bloom Filter    |
| **大量データの事前フィルタ**     | Bloom Filter  | Cuckoo Filter         | -             | Hash Table      |
| **分散システム**                 | Bloom Filter  | Cuckoo Filter         | Hash Table    | Set（Tree）     |
| **リアルタイムストリーム**       | Bloom Filter  | Cuckoo Filter         | Hash Table    | Set（Tree）     |
| **データベースインデックス**     | Bloom Filter  | Hash Table            | Set（Tree）   | Skip List       |
| **キャッシュシステム**           | Bloom Filter  | Hash Table            | Cuckoo Filter | Set（Tree）     |

## メモリ使用量比較（1M 要素の場合）

| データ構造                | メモリ使用量 | 偽陽性率 | 削除可能 | 備考                 |
| ------------------------- | ------------ | -------- | -------- | -------------------- |
| **Bloom Filter（1%FPR）** | ~1.2MB       | 1%       | ❌       | 最小メモリ           |
| **Cuckoo Filter**         | ~2.0MB       | <3%      | ✅       | 削除可能             |
| **Hash Table**            | ~24MB        | 0%       | ✅       | 完全正確             |
| **TreeSet**               | ~32MB        | 0%       | ✅       | ソート済み           |
| **Counting Bloom Filter** | ~4.8MB       | 1%       | ✅       | 削除可能、メモリ多用 |

## 性能特性比較

| 操作パターン                | Bloom Filter | Hash Table    | TreeSet     | 配列（ソート済み） |
| --------------------------- | ------------ | ------------- | ----------- | ------------------ |
| **ランダム検索（1M 要素）** | ~100ns       | ~150ns        | ~500ns      | ~20μs              |
| **大量挿入（1M 要素）**     | ~2ms         | ~15ms         | ~50ms       | ~500ms             |
| **メモリアクセスパターン**  | 順次         | ランダム      | ランダム    | 順次               |
| **CPU キャッシュ効率**      | ✅ 高い      | ⚠️ 中程度     | ❌ 低い     | ✅ 高い            |
| **スケーラビリティ**        | ✅ 一定      | ⚠️ 負荷率依存 | ✅ O(log n) | ❌ O(log n)        |

## 制約条件別選択

| 制約条件                | 推奨選択           | 理由                     |
| ----------------------- | ------------------ | ------------------------ |
| **メモリ < 1MB**        | Bloom Filter       | 最小メモリフットプリント |
| **偽陽性率 < 0.1%**     | Hash Table         | 完全な正確性             |
| **削除操作必須**        | Cuckoo Filter      | 削除可能な確率的構造     |
| **リアルタイム制約**    | Bloom Filter       | 予測可能な性能           |
| **分散環境**            | Bloom Filter       | ネットワーク効率         |
| **大量データ（TB 級）** | Bloom Filter       | スケーラビリティ         |
| **組み込みシステム**    | Bloom Filter       | リソース効率             |
| **正確性最優先**        | Hash Table/TreeSet | 偽陽性なし               |

# パフォーマンス特性

## Bloom Filter 設計パラメータの影響

| 要素数(n) | FPR 目標 | 最適 m     | 最適 k | 実際の FPR | メモリ使用量 |
| --------- | -------- | ---------- | ------ | ---------- | ------------ |
| 1,000     | 1%       | 9,586      | 7      | 0.82%      | 1.2KB        |
| 10,000    | 1%       | 95,851     | 7      | 0.82%      | 12KB         |
| 100,000   | 1%       | 958,506    | 7      | 0.82%      | 120KB        |
| 1,000,000 | 1%       | 9,585,059  | 7      | 0.82%      | 1.2MB        |
| 1,000,000 | 0.1%     | 14,377,589 | 10     | 0.08%      | 1.8MB        |

## ハッシュ関数数と性能の関係

```
最適ハッシュ関数数: k = (m/n) * ln(2)

k値と偽陽性率の関係（m/n = 10の場合）:
k=1: FPR ≈ 39.3%
k=5: FPR ≈ 6.2%
k=7: FPR ≈ 0.8%  ← 最適
k=10: FPR ≈ 1.2%
k=15: FPR ≈ 7.1%

処理時間とのトレードオフ:
- k値が大きいほど偽陽性率は低下
- ただし処理時間（k * ハッシュ計算時間）は増加
- 最適点での運用が重要
```

# まとめ

Bloom Filter（ブルームフィルター）は、確率的データ構造として非常に少ないメモリで高速な存在確認を実現する優秀なデータ構造です。偽陽性は許容するが偽陰性は発生しない特性を活かし、データベース、ネットワーク、セキュリティなど幅広い分野で活用されています。

Go 言語での実装においても、ビット配列とハッシュ関数を組み合わせることで効率的な Bloom Filter を構築できます。用途や制約に応じて、標準版、Counting 版、Scalable 版などの変種を選択し、適切なパラメータ設計を行うことが重要です。

特に、メモリ制約が厳しい環境や大量データの事前フィルタリングが必要な場面では、Bloom Filter の導入により大幅な性能改善とリソース効率化を実現できます。
