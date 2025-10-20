package main

import (
    "fmt"
)

func main() {
    fmt.Println("========================================")
    fmt.Println("=== カスタム配列実装のデモ ===")
    fmt.Println("========================================")
    
    // 静的配列のデモ
    demonstrateStaticArray()
    
    // 動的配列のデモ
    demonstrateDynamicArray()
}

func demonstrateStaticArray() {
    fmt.Println("\n=== 静的配列のデモ ===")
    
    // 静的配列の作成
    arr := NewStaticArray(5)
    fmt.Printf("作成: %s\n", arr)
    
    // 値の設定
    for i := 0; i < arr.Size(); i++ {
        arr.Set(i, (i+1)*10)
    }
    fmt.Printf("値設定後: %s\n", arr)
    
    // 値の取得
    if val, err := arr.Get(2); err == nil {
        fmt.Printf("arr[2] = %d\n", val)
    }
    
    // エラーハンドリング
    if _, err := arr.Get(10); err != nil {
        fmt.Printf("エラー: %s\n", err)
    }
    
    fmt.Printf("配列サイズ: %d\n", arr.Size())
}

func demonstrateDynamicArray() {
    fmt.Println("\n=== 動的配列のデモ ===")
    
    // 動的配列の作成
    arr := NewDynamicArray()
    fmt.Printf("初期状態: %s\n", arr)
    arr.PrintInfo()
    
    // 要素の追加
    fmt.Println("\n--- 要素の追加 ---")
    for i := 1; i <= 5; i++ {
        arr.Append(i * 10)
        fmt.Printf("追加後 %d: %s (長さ: %d, 容量: %d)\n", 
            i, arr, arr.Length(), arr.Capacity())
    }
    
    // 中間への挿入
    fmt.Println("\n--- 中間への挿入 ---")
    arr.Insert(2, 999)
    fmt.Printf("インデックス2に999を挿入: %s\n", arr)
    
    // 要素の変更
    fmt.Println("\n--- 要素の変更 ---")
    arr.Set(0, 111)
    fmt.Printf("インデックス0を111に変更: %s\n", arr)
    
    // 検索
    fmt.Println("\n--- 検索 ---")
    index := arr.IndexOf(999)
    fmt.Printf("値999のインデックス: %d\n", index)
    fmt.Printf("値999が含まれている: %t\n", arr.Contains(999))
    fmt.Printf("値777が含まれている: %t\n", arr.Contains(777))
    
    // 削除
    fmt.Println("\n--- 削除 ---")
    if val, err := arr.Pop(); err == nil {
        fmt.Printf("末尾から削除: %d, 結果: %s\n", val, arr)
    }
    
    arr.Delete(2)
    fmt.Printf("インデックス2を削除: %s\n", arr)
    
    // 配列の状態確認
    fmt.Println("\n--- 最終状態 ---")
    arr.PrintInfo()
    
    // スライスとしてエクスポート
    slice := arr.ToSlice()
    fmt.Printf("スライスとしてエクスポート: %v\n", slice)
    
    // クリア
    arr.Clear()
    fmt.Printf("クリア後: %s\n", arr)
    fmt.Printf("空かどうか: %t\n", arr.IsEmpty())
}
