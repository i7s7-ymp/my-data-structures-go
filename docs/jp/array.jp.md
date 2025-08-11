# 配列（Array）とは

配列は、同じ型の要素を連続したメモリ領域に格納するデータ構造です。各要素にはインデックス（添字）が割り当てられ、O(1)でランダムアクセスが可能です。
**このデータ構造を使う一番の利点は、ランダムアクセスが高速である点です。**

# 基本的な特徴

## 静的配列の特徴

- **固定サイズ**: 宣言時にサイズが決まり、実行時に変更不可
- **値型**: 配列全体が値として扱われる（コピー時は全要素がコピーされる）
- **メモリ効率**: 連続したメモリ領域に格納され、メモリ使用量が予測可能
- **型安全**: コンパイル時に型チェックが行われる

## 動的配列の特徴

- **動的サイズ**: 実行時にサイズを変更可能
- **参照型**: 内部的には配列への参照を持つ
- **自動リサイズ**: 容量が足りなくなると自動的に拡張される
- **高性能**: 要素追加操作は平均的に O(1)

## 共通の特徴

- **ランダムアクセス**: インデックスを使用して O(1) で任意の要素にアクセス可能。
- **メモリ効率**: 連続したメモリ領域を使用するため、キャッシュ効率が高い。
- **用途**: 固定サイズのデータ管理、数値計算、バッファ管理など。

# 行える処理

| 機能       | 説明                             | 計算量 | 補足                                                                       |
| ---------- | -------------------------------- | ------ | -------------------------------------------------------------------------- |
| 読み取り   | 指定したインデックスの要素を取得 | O(1)   |                                                                            |
| 書き込み   | 指定したインデックスに要素を設定 | O(1)   |                                                                            |
| サイズ取得 | 配列のサイズを取得               | O(1)   |                                                                            |
| 挿入       | 配列の途中に要素を挿入           | O(n)   | 挿入位置より後ろの全要素を 1 つずつ後ろにシフトする必要がある              |
| 削除       | 配列の途中から要素を削除         | O(n)   | 削除位置より後ろの全要素を 1 つずつ前にシフトする必要がある                |
| 検索       | 指定した値を持つ要素を検索       | O(n)   | 線形探索の場合、最悪全ての要素を確認する必要がある。二分探索などで改善可能 |

# 注意点

## 静的配列の注意点

1. **サイズ制限**: 静的配列のサイズは作成時に固定され、実行時に変更できません。
2. **メモリの事前確保**: サイズが決まっているため、使わない領域もメモリを占有する可能性があります。
3. **型の固定性**: 配列のサイズも型の一部となる場合があります

## 動的配列の注意点

1. **メモリ再割り当て**: 容量が不足すると新しいメモリ領域を確保してデータをコピーするため、パフォーマンスに影響する場合があります。
2. **メモリ使用量の予測困難**: 動的にサイズが変わるため、メモリ使用量を事前に予測することが困難です。
3. **参照の管理**: 動的配列は参照型である場合が多いため、複数の変数が同じ配列を参照する可能性があります

## 共通の注意点

1. **メモリの連続性**: 連続したメモリ領域が必要なため、大きな配列を作成する際にメモリ不足が発生する可能性があります。
2. **挿入・削除のコスト**: 配列の途中に要素を挿入または削除する場合、要素のシフトが必要となり O(n) の計算量が発生します。

# 実装方法

## 静的配列（Array）

```go
package main

import "fmt"

func main() {
    var staticArray [5]int // サイズが5の静的配列
    staticArray[0] = 10
    staticArray[1] = 20

    fmt.Println(staticArray) // [10 20 0 0 0]
}
```

### 値型の特性

```go
func modifyArray(arr [3]int) {
    arr[0] = 100  // 元の配列は変更されない（コピーが渡される）
}

func main() {
    original := [3]int{1, 2, 3}
    modifyArray(original)
    fmt.Println(original) // [1 2 3] - 変更されない
}
```

### Go 特有の注意点

1. **型の厳密性**: `[3]int` と `[5]int` は完全に異なる型として扱われる
2. **関数渡しのコスト**: 大きな配列を関数に渡すと全要素がコピーされ、パフォーマンスに影響
3. **初期化**: 宣言時に明示的にサイズを指定する必要がある

## 静的配列のカスタム実装

Go の組み込み配列を使用せずに静的配列を実装する場合、固定サイズの構造体を使用します。

```go
package main

import "fmt"

// 静的配列構造体
type StaticArray struct {
    data [5]int
}

func (a *StaticArray) Set(index int, value int) {
    if index >= 0 && index < len(a.data) {
        a.data[index] = value
    }
}

func (a *StaticArray) Get(index int) int {
    if index >= 0 && index < len(a.data) {
        return a.data[index]
    }
    return 0 // エラーハンドリングは省略
}

func main() {
    arr := StaticArray{}
    arr.Set(0, 10)
    arr.Set(1, 20)

    fmt.Println(arr.Get(0)) // 10
    fmt.Println(arr.Get(1)) // 20
}
```

## スライス（Slice）- 動的配列

```go
// 動的サイズのスライス
var slice []int = []int{1, 2, 3, 4, 5}

// インデックスアクセス (O(1))
fmt.Println("インデックス 2 の要素:", slice[2])

// 末尾へ動的に要素追加 (平均 O(1))
slice = append(slice, 6)

// 中間への挿入 (O(n))
index := 2
slice = append(slice[:index], append([]int{99}, slice[index:]...)...)

// 中間からの削除 (O(n))
slice = append(slice[:index], slice[index+1:]...)
```

### 内部構造

```go
// スライスは内部的に以下の構造を持つ
type slice struct {
    ptr unsafe.Pointer // 配列への参照
    len int           // 長さ
    cap int           // 容量
}
```

### 容量（Capacity）の概念

```go
s := make([]int, 3, 5)  // 長さ3、容量5のスライス
fmt.Println(len(s))     // 3
fmt.Println(cap(s))     // 5

// 容量内であれば再割り当てなしで要素追加可能
```

### nil スライスと空スライス

```go
var nilSlice []int              // nil スライス
emptySlice := []int{}           // 空スライス
madeSlice := make([]int, 0)     // make で作成した空スライス

fmt.Println(nilSlice == nil)    // true
fmt.Println(emptySlice == nil)  // false
fmt.Println(madeSlice == nil)   // false
```

### スライスの共有

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:3]  // [2, 3]
slice2 := original[2:4]  // [3, 4]

slice1[1] = 99
fmt.Println(original)    // [1 2 99 4 5] - 元の配列も変更される
fmt.Println(slice2)      // [99 4] - slice2も影響を受ける
```

### append の挙動

```go
s1 := []int{1, 2, 3}
s2 := s1

s1 = append(s1, 4)  // 容量が足りない場合、新しい配列が作成される

// s1とs2が異なる配列を参照する可能性がある
```

### Go 特有の注意点

1. **nil スライスの操作**:

```go
var s []int
s = append(s, 1)  // OK: nil スライスにも append 可能
// s[0] = 1       // パニック: nil スライスにはインデックスアクセス不可
```

2. **スライスの比較**:

```go
s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
// fmt.Println(s1 == s2)  // コンパイルエラー: スライスは直接比較できない
fmt.Println(s1 == nil)   // OK: nil との比較は可能
```

3. **メモリリーク**:

```go
func getSmallSlice() []int {
    largeSlice := make([]int, 1000000)
    // 小さな部分だけを返すが、大きな配列への参照が残る
    return largeSlice[:5]  // メモリリークの可能性
}


// 対処法: コピーを作成
func getSmallSliceSafe() []int {
    largeSlice := make([]int, 1000000)
    smallSlice := make([]int, 5)
    copy(smallSlice, largeSlice[:5])
    return smallSlice
}
```

4. **関数への渡し方**:

```go
// スライスを変更する関数
func modifySliceContent(s []int) {
    s[0] = 100  // 元のスライスの内容が変更される
}

func modifySliceLength(s []int) {
    s = append(s, 4)  // 元のスライスの長さは変更されない
}

// スライス自体を変更したい場合はポインタを使用
func modifySliceItself(s *[]int) {
    *s = append(*s, 4)  // 元のスライスの長さも変更される
}
```

### パフォーマンスの考慮事項

1. **append の効率性**:

```go
// 非効率: 容量を考慮しない
s := []int{}
for i := 0; i < 1000; i++ {
    s = append(s, i)  // 何度も再割り当てが発生
}

// 効率的: 事前に容量を確保
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    s = append(s, i)  // 再割り当てが発生しない
}
```

2. **copy vs append**:

```go
// copy を使用（より効率的）
dst := make([]int, len(src))
copy(dst, src)

// append を使用
dst := append([]int(nil), src...)
```

## 動的配列のカスタム実装

Go のスライスを使用せずに動的配列を実装する場合、内部で配列を管理し、サイズ変更時に新しい配列を作成してコピーします。

```go
package main

import "fmt"

// 動的配列構造体
type DynamicArray struct {
    data []int
    size int
}

func NewDynamicArray() *DynamicArray {
    return &DynamicArray{
        data: make([]int, 0),
    }
}

func (a *DynamicArray) Append(value int) {
    if a.size == len(a.data) {
        newData := make([]int, a.size*2+1) // サイズを2倍に拡張
        copy(newData, a.data)
        a.data = newData
    }
    a.data = append(a.data, value)
    a.size++
}

func (a *DynamicArray) Get(index int) int {
    if index >= 0 && index < a.size {
        return a.data[index]
    }
    return 0 // エラーハンドリングは省略
}

func main() {
    arr := NewDynamicArray()
    arr.Append(10)
    arr.Append(20)

    fmt.Println(arr.Get(0)) // 10
    fmt.Println(arr.Get(1)) // 20
}
```

# 配列の応用例

1. **数値計算**: 行列やベクトルの計算。
2. **バッファ管理**: 固定サイズのデータバッファ。
3. **検索アルゴリズム**: バイナリサーチなどの効率的な検索。

# 他のデータ構造との比較

| 特徴                 | 静的配列 | 動的配列 | リスト | ハッシュテーブル |
| -------------------- | -------- | -------- | ------ | ---------------- |
| **ランダムアクセス** | 可能     | 可能     | 不可   | 不可             |
| **サイズ変更**       | 不可     | 動的     | 動的   | 動的             |
| **挿入・削除の効率** | O(n)     | O(n)     | O(1)   | O(1)             |
| **メモリ効率**       | 高い     | 中程度   | 中程度 | 中程度           |
| **型安全性**         | 高い     | 高い     | 中程度 | 中程度           |
| **初期化時のサイズ** | 必須     | 不要     | 不要   | 不要             |
