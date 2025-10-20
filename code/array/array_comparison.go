package main

import (
    "fmt"
    "time"
)

func main() {
    // 組み込みの配列とスライスのデモンストレーション
    demonstrateBuiltinArrays()
    
    // カスタム実装のデモンストレーション
    demonstrateCustomArrays()
    
    // パフォーマンス比較
    performanceComparison()
}

func demonstrateBuiltinArrays() {
    fmt.Println("========================================")
    fmt.Println("=== 組み込みの配列とスライスのデモ ===")
    fmt.Println("========================================")
    
    // 配列の基本操作
    fmt.Println("=== 配列の操作 ===")
    array := [5]int{1, 2, 3, 4, 5} // 固定サイズの配列
    fmt.Println("配列:", array)

    // インデックスアクセス (O(1))
    fmt.Println("インデックス 2 の要素:", array[2])

    // 配列のサイズは固定のため、挿入や削除はできない
    // array[5] = 6 // コンパイルエラー: 配列のサイズを超える操作は不可

    // スライスの基本操作
    fmt.Println("\n=== スライスの操作 ===")
    slice := []int{1, 2, 3, 4, 5} // 動的サイズのスライス
    fmt.Println("スライス:", slice)

    // インデックスアクセス (O(1))
    fmt.Println("インデックス 2 の要素:", slice[2])

    // 末尾への要素追加 (平均 O(1))
    start := time.Now()
    slice = append(slice, 6)
    fmt.Println("要素追加後のスライス:", slice)
    fmt.Println("追加処理時間:", time.Since(start))

    // 中間への挿入 (O(n))
    start = time.Now()
    index := 2
    slice = append(slice[:index], append([]int{99}, slice[index:]...)...)
    fmt.Println("中間挿入後のスライス:", slice)
    fmt.Println("中間挿入処理時間:", time.Since(start))

    // 中間からの削除 (O(n))
    start = time.Now()
    slice = append(slice[:index], slice[index+1:]...)
    fmt.Println("中間削除後のスライス:", slice)
    fmt.Println("中間削除処理時間:", time.Since(start))

    // スライスの容量確認
    fmt.Println("\n=== スライスの容量確認 ===")
    fmt.Printf("スライスの長さ: %d, 容量: %d\n", len(slice), cap(slice))

    // スライスの容量が増える挙動
    fmt.Println("\n=== スライスの容量が増える挙動 ===")
    for i := 0; i < 10; i++ {
        slice = append(slice, i)
        fmt.Printf("追加後: 長さ: %d, 容量: %d\n", len(slice), cap(slice))
    }
}

func demonstrateCustomArrays() {
    fmt.Println("\n========================================")
    fmt.Println("=== カスタム実装の配列のデモ ===")
    fmt.Println("========================================")
    
    // 静的配列のテスト
    fmt.Println("=== 静的配列のテスト ===")
    staticArray := NewStaticArray(5)
    fmt.Println("初期状態:", staticArray)
    
    // 要素の設定
    for i := 0; i < staticArray.Size(); i++ {
        staticArray.Set(i, (i+1)*10)
    }
    fmt.Println("値設定後:", staticArray)
    
    // 要素の取得
    if val, err := staticArray.Get(2); err == nil {
        fmt.Printf("インデックス 2 の要素: %d\n", val)
    }
    
    // 範囲外アクセスのテスト
    if _, err := staticArray.Get(10); err != nil {
        fmt.Println("エラー:", err)
    }
    
    // 動的配列のテスト
    fmt.Println("\n=== 動的配列のテスト ===")
    dynamicArray := NewDynamicArray()
    fmt.Printf("初期状態: %s (長さ: %d, 容量: %d)\n", 
        dynamicArray, dynamicArray.Length(), dynamicArray.Capacity())
    
    // 要素の追加
    fmt.Println("\n--- 要素の追加テスト ---")
    for i := 1; i <= 8; i++ {
        start := time.Now()
        dynamicArray.Append(i * 10)
        fmt.Printf("追加後: %s (長さ: %d, 容量: %d) 処理時間: %v\n", 
            dynamicArray, dynamicArray.Length(), dynamicArray.Capacity(), time.Since(start))
    }
    
    // メモリ使用量の表示
    fmt.Println("\n--- メモリ使用量 ---")
    dynamicArray.MemoryInfo()
    
    // 中間への挿入
    fmt.Println("\n--- 中間への挿入テスト ---")
    start := time.Now()
    err := dynamicArray.Insert(3, 999)
    if err == nil {
        fmt.Printf("インデックス 3 に 999 を挿入: %s 処理時間: %v\n", 
            dynamicArray, time.Since(start))
    }
    
    // 要素の取得
    if val, err := dynamicArray.Get(3); err == nil {
        fmt.Printf("インデックス 3 の要素: %d\n", val)
    }
    
    // 要素の変更
    dynamicArray.Set(0, 1111)
    fmt.Printf("インデックス 0 を 1111 に変更: %s\n", dynamicArray)
    
    // 検索機能のテスト
    fmt.Println("\n--- 検索機能テスト ---")
    index := dynamicArray.IndexOf(999)
    fmt.Printf("値 999 のインデックス: %d\n", index)
    fmt.Printf("値 999 が含まれている: %t\n", dynamicArray.Contains(999))
    fmt.Printf("値 777 が含まれている: %t\n", dynamicArray.Contains(777))
    
    // 末尾削除のテスト
    fmt.Println("\n--- 末尾削除テスト ---")
    if val, err := dynamicArray.Pop(); err == nil {
        fmt.Printf("末尾から削除した値: %d, 現在の配列: %s\n", val, dynamicArray)
    }
    
    // 中間からの削除
    fmt.Println("\n--- 中間からの削除テスト ---")
    start = time.Now()
    err = dynamicArray.Delete(3)
    if err == nil {
        fmt.Printf("インデックス 3 を削除: %s 処理時間: %v\n", 
            dynamicArray, time.Since(start))
    }
    
    // 配列の詳細情報表示
    fmt.Println("\n--- 配列詳細情報 ---")
    dynamicArray.PrintInfo()
}

func performanceComparison() {
    fmt.Println("\n========================================")
    fmt.Println("=== パフォーマンス比較 ===")
    fmt.Println("========================================")
    
    const testSize = 50000
    
    // カスタム動的配列の性能テスト
    fmt.Println("--- カスタム動的配列の性能テスト ---")
    customArrayPerformanceTest(testSize)
    
    // 組み込みスライスの性能テスト
    fmt.Println("\n--- 組み込みスライスの性能テスト ---")
    builtinSlicePerformanceTest(testSize)
}

func customArrayPerformanceTest(size int) {
    da := NewDynamicArray()
    
    // 要素の追加
    start := time.Now()
    for i := 0; i < size; i++ {
        da.Append(i)
    }
    appendTime := time.Since(start)
    
    // 先頭に挿入（最悪ケース）
    start = time.Now()
    da.Insert(0, -1)
    insertTime := time.Since(start)
    
    // 中間から削除
    start = time.Now()
    da.Delete(size / 2)
    deleteTime := time.Since(start)
    
    // ランダムアクセス
    start = time.Now()
    for i := 0; i < 1000; i++ {
        da.Get(i * 10)
    }
    accessTime := time.Since(start)
    
    // 検索
    start = time.Now()
    da.IndexOf(size / 2)
    searchTime := time.Since(start)
    
    fmt.Printf("カスタム動的配列の結果 (%d要素):\n", size)
    fmt.Printf("- 要素追加: %v\n", appendTime)
    fmt.Printf("- 先頭挿入: %v\n", insertTime)
    fmt.Printf("- 中間削除: %v\n", deleteTime)
    fmt.Printf("- 1000回アクセス: %v\n", accessTime)
    fmt.Printf("- 検索: %v\n", searchTime)
    fmt.Printf("最終状態: 長さ=%d, 容量=%d\n", da.Length(), da.Capacity())
    da.MemoryInfo()
}

func builtinSlicePerformanceTest(size int) {
    slice := make([]int, 0)
    
    // 要素の追加
    start := time.Now()
    for i := 0; i < size; i++ {
        slice = append(slice, i)
    }
    appendTime := time.Since(start)
    
    // 先頭に挿入（最悪ケース）
    start = time.Now()
    slice = append([]int{-1}, slice...)
    insertTime := time.Since(start)
    
    // 中間から削除
    index := size / 2
    start = time.Now()
    slice = append(slice[:index], slice[index+1:]...)
    deleteTime := time.Since(start)
    
    // ランダムアクセス
    start = time.Now()
    for i := 0; i < 1000; i++ {
        _ = slice[i*10]
    }
    accessTime := time.Since(start)
    
    // 検索
    start = time.Now()
    searchValue := size / 2
    found := -1
    for i, v := range slice {
        if v == searchValue {
            found = i
            break
        }
    }
    searchTime := time.Since(start)
    _ = found
    
    fmt.Printf("組み込みスライスの結果 (%d要素):\n", size)
    fmt.Printf("- 要素追加: %v\n", appendTime)
    fmt.Printf("- 先頭挿入: %v\n", insertTime)
    fmt.Printf("- 中間削除: %v\n", deleteTime)
    fmt.Printf("- 1000回アクセス: %v\n", accessTime)
    fmt.Printf("- 検索: %v\n", searchTime)
    fmt.Printf("最終状態: 長さ=%d, 容量=%d\n", len(slice), cap(slice))
}
