package main

import (
    "fmt"
    "unsafe"
)

// 静的配列の実装
type StaticArray struct {
    data []int
    size int
}

// 静的配列のコンストラクタ
func NewStaticArray(size int) *StaticArray {
    return &StaticArray{
        data: make([]int, size),
        size: size,
    }
}

// 要素の取得 (O(1))
func (sa *StaticArray) Get(index int) (int, error) {
    if index < 0 || index >= sa.size {
        return 0, fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, sa.size-1)
    }
    return sa.data[index], nil
}

// 要素の設定 (O(1))
func (sa *StaticArray) Set(index int, value int) error {
    if index < 0 || index >= sa.size {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, sa.size-1)
    }
    sa.data[index] = value
    return nil
}

// サイズの取得
func (sa *StaticArray) Size() int {
    return sa.size
}

// 配列の表示
func (sa *StaticArray) String() string {
    return fmt.Sprintf("StaticArray%v", sa.data)
}

// 動的配列の実装
type DynamicArray struct {
    data     []int
    length   int
    capacity int
}

// 動的配列のコンストラクタ
func NewDynamicArray() *DynamicArray {
    initialCapacity := 2
    return &DynamicArray{
        data:     make([]int, initialCapacity),
        length:   0,
        capacity: initialCapacity,
    }
}

// 指定された容量で動的配列を作成
func NewDynamicArrayWithCapacity(capacity int) *DynamicArray {
    if capacity < 1 {
        capacity = 2
    }
    return &DynamicArray{
        data:     make([]int, capacity),
        length:   0,
        capacity: capacity,
    }
}

// 容量を拡張する内部メソッド
func (da *DynamicArray) resize() {
    oldCapacity := da.capacity
    da.capacity *= 2
    newData := make([]int, da.capacity)
    
    // 手動でデータをコピー（copy関数を使わない）
    for i := 0; i < da.length; i++ {
        newData[i] = da.data[i]
    }
    
    da.data = newData
    fmt.Printf("容量を %d から %d に拡張しました\n", oldCapacity, da.capacity)
}

// 要素の取得 (O(1))
func (da *DynamicArray) Get(index int) (int, error) {
    if index < 0 || index >= da.length {
        return 0, fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length-1)
    }
    return da.data[index], nil
}

// 要素の設定 (O(1))
func (da *DynamicArray) Set(index int, value int) error {
    if index < 0 || index >= da.length {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length-1)
    }
    da.data[index] = value
    return nil
}

// 末尾への要素追加 (平均 O(1))
func (da *DynamicArray) Append(value int) {
    if da.length >= da.capacity {
        da.resize()
    }
    da.data[da.length] = value
    da.length++
}

// 指定位置への要素挿入 (O(n))
func (da *DynamicArray) Insert(index int, value int) error {
    if index < 0 || index > da.length {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length)
    }
    
    if da.length >= da.capacity {
        da.resize()
    }
    
    // 要素を右にシフト（手動実装）
    for i := da.length; i > index; i-- {
        da.data[i] = da.data[i-1]
    }
    
    da.data[index] = value
    da.length++
    return nil
}

// 指定位置の要素削除 (O(n))
func (da *DynamicArray) Delete(index int) error {
    if index < 0 || index >= da.length {
        return fmt.Errorf("インデックス %d が範囲外です (0-%d)", index, da.length-1)
    }
    
    // 要素を左にシフト（手動実装）
    for i := index; i < da.length-1; i++ {
        da.data[i] = da.data[i+1]
    }
    
    da.length--
    return nil
}

// 末尾の要素を削除 (O(1))
func (da *DynamicArray) Pop() (int, error) {
    if da.length == 0 {
        return 0, fmt.Errorf("配列が空です")
    }
    
    value := da.data[da.length-1]
    da.length--
    return value, nil
}

// 長さの取得
func (da *DynamicArray) Length() int {
    return da.length
}

// 容量の取得
func (da *DynamicArray) Capacity() int {
    return da.capacity
}

// 配列が空かどうか
func (da *DynamicArray) IsEmpty() bool {
    return da.length == 0
}

// 配列の表示
func (da *DynamicArray) String() string {
    result := "DynamicArray["
    for i := 0; i < da.length; i++ {
        if i > 0 {
            result += ", "
        }
        result += fmt.Sprintf("%d", da.data[i])
    }
    result += "]"
    return result
}

// 配列を空にする
func (da *DynamicArray) Clear() {
    da.length = 0
}

// 指定した値を検索 (O(n))
func (da *DynamicArray) IndexOf(value int) int {
    for i := 0; i < da.length; i++ {
        if da.data[i] == value {
            return i
        }
    }
    return -1 // 見つからない場合
}

// 指定した値が含まれているかチェック (O(n))
func (da *DynamicArray) Contains(value int) bool {
    return da.IndexOf(value) != -1
}

// 配列の内容をスライスとして取得
func (da *DynamicArray) ToSlice() []int {
    result := make([]int, da.length)
    for i := 0; i < da.length; i++ {
        result[i] = da.data[i]
    }
    return result
}

// メモリ使用量の情報を表示
func (da *DynamicArray) MemoryInfo() {
    elementSize := int(unsafe.Sizeof(int(0)))
    usedMemory := da.length * elementSize
    allocatedMemory := da.capacity * elementSize
    fmt.Printf("メモリ情報: 使用中=%dバイト, 確保済み=%dバイト, 使用率=%.1f%%\n", 
        usedMemory, allocatedMemory, float64(usedMemory)/float64(allocatedMemory)*100)
}

// 配列の詳細情報を表示
func (da *DynamicArray) PrintInfo() {
    fmt.Printf("動的配列情報:\n")
    fmt.Printf("  長さ: %d\n", da.length)
    fmt.Printf("  容量: %d\n", da.capacity)
    fmt.Printf("  空: %t\n", da.IsEmpty())
    da.MemoryInfo()
}
