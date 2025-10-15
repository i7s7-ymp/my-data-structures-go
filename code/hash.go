package main

import (
	"fmt"
	"time"
)

// キー・バリューペア
type KeyValue struct {
	Key   string
	Value interface{}
	Next  *KeyValue // チェイン法用のポインタ
}

// ハッシュテーブル構造体
type HashTable struct {
	Buckets  []*KeyValue // バケット配列
	Size     int         // 現在の要素数
	Capacity int         // バケットの容量
}

// ハッシュテーブルの作成
func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		Buckets:  make([]*KeyValue, capacity),
		Size:     0,
		Capacity: capacity,
	}
}

// シンプルなハッシュ関数（djb2アルゴリズム）
func (ht *HashTable) hash(key string) int {
	hash := 5381
	for _, char := range key {
		hash = ((hash << 5) + hash) + int(char)
	}
	return hash % ht.Capacity
}

// 挿入操作 O(1)平均
func (ht *HashTable) Put(key string, value interface{}) {
	index := ht.hash(key)

	// 既存のキーをチェック
	current := ht.Buckets[index]
	for current != nil {
		if current.Key == key {
			current.Value = value // 値を更新
			return
		}
		current = current.Next
	}

	// 新しいキー・バリューペアを先頭に挿入
	newNode := &KeyValue{
		Key:   key,
		Value: value,
		Next:  ht.Buckets[index],
	}
	ht.Buckets[index] = newNode
	ht.Size++
}

// 検索操作 O(1)平均
func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.hash(key)
	current := ht.Buckets[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return nil, false // キーが見つからない
}

// 削除操作 O(1)平均
func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)
	current := ht.Buckets[index]
	var prev *KeyValue

	for current != nil {
		if current.Key == key {
			if prev == nil {
				// 先頭要素の削除
				ht.Buckets[index] = current.Next
			} else {
				// 中間要素の削除
				prev.Next = current.Next
			}
			ht.Size--
			return true
		}
		prev = current
		current = current.Next
	}

	return false // キーが見つからない
}

// 存在確認 O(1)平均
func (ht *HashTable) Contains(key string) bool {
	_, exists := ht.Get(key)
	return exists
}

// 全要素の表示
func (ht *HashTable) Display() {
	fmt.Printf("ハッシュテーブル (サイズ: %d, 容量: %d)\n", ht.Size, ht.Capacity)
	for i, bucket := range ht.Buckets {
		if bucket != nil {
			fmt.Printf("バケット[%d]: ", i)
			current := bucket
			for current != nil {
				fmt.Printf("(%s: %v)", current.Key, current.Value)
				if current.Next != nil {
					fmt.Print(" -> ")
				}
				current = current.Next
			}
			fmt.Println()
		}
	}
	fmt.Println()
}

// 負荷率の計算
func (ht *HashTable) LoadFactor() float64 {
	return float64(ht.Size) / float64(ht.Capacity)
}

// サイズを取得
func (ht *HashTable) GetSize() int {
	return ht.Size
}

// 空かどうかを確認
func (ht *HashTable) IsEmpty() bool {
	return ht.Size == 0
}

// 全てのキーを取得 O(n)
func (ht *HashTable) GetAllKeys() []string {
	keys := make([]string, 0, ht.Size)
	for _, bucket := range ht.Buckets {
		current := bucket
		for current != nil {
			keys = append(keys, current.Key)
			current = current.Next
		}
	}
	return keys
}

// 衝突の統計情報を取得
func (ht *HashTable) GetCollisionStats() (int, int, float64) {
	usedBuckets := 0
	maxChainLength := 0

	for _, bucket := range ht.Buckets {
		if bucket != nil {
			usedBuckets++
			chainLength := 0
			current := bucket
			for current != nil {
				chainLength++
				current = current.Next
			}
			if chainLength > maxChainLength {
				maxChainLength = chainLength
			}
		}
	}

	avgChainLength := 0.0
	if usedBuckets > 0 {
		avgChainLength = float64(ht.Size) / float64(usedBuckets)
	}

	return usedBuckets, maxChainLength, avgChainLength
}

// 計算量を測定するためのテスト関数
func measureTime(operation string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s の実行時間: %v\n", operation, duration)
}

func main() {
	fmt.Println("=== ハッシュテーブルの基本操作と計算量の確認 ===\n")

	// 1. ハッシュテーブルの作成
	fmt.Println("1. ハッシュテーブルの作成")
	ht := NewHashTable(10)
	fmt.Printf("容量10のハッシュテーブルを作成しました\n")
	fmt.Printf("初期状態 - サイズ: %d, 負荷率: %.2f\n\n", ht.GetSize(), ht.LoadFactor())

	// 2. データの挿入操作（O(1)平均）
	fmt.Println("2. データの挿入操作（O(1)平均）")
	fruits := map[string]int{
		"apple":  100,
		"banana": 200,
		"orange": 150,
		"grape":  300,
		"lemon":  120,
	}

	for fruit, price := range fruits {
		measureTime(fmt.Sprintf("%sを挿入", fruit), func() {
			ht.Put(fruit, price)
		})
	}
	ht.Display()

	// 3. データの検索操作（O(1)平均）
	fmt.Println("3. データの検索操作（O(1)平均）")
	searchKeys := []string{"apple", "banana", "mango"}
	for _, key := range searchKeys {
		measureTime(fmt.Sprintf("%sを検索", key), func() {
			if value, found := ht.Get(key); found {
				fmt.Printf("  → %s: %v が見つかりました\n", key, value)
			} else {
				fmt.Printf("  → %s は見つかりませんでした\n", key)
			}
		})
	}
	fmt.Println()

	// 4. 存在確認操作（O(1)平均）
	fmt.Println("4. 存在確認操作（O(1)平均）")
	checkKeys := []string{"grape", "watermelon"}
	for _, key := range checkKeys {
		measureTime(fmt.Sprintf("%sの存在確認", key), func() {
			exists := ht.Contains(key)
			fmt.Printf("  → %s は存在する: %t\n", key, exists)
		})
	}
	fmt.Println()

	// 5. データの削除操作（O(1)平均）
	fmt.Println("5. データの削除操作（O(1)平均）")
	deleteKeys := []string{"orange", "kiwi"}
	for _, key := range deleteKeys {
		measureTime(fmt.Sprintf("%sを削除", key), func() {
			deleted := ht.Delete(key)
			fmt.Printf("  → %s の削除: %t\n", key, deleted)
		})
	}
	ht.Display()

	// 6. 衝突の確認
	fmt.Println("6. 衝突の確認")
	usedBuckets, maxChain, avgChain := ht.GetCollisionStats()
	fmt.Printf("使用中のバケット数: %d/%d\n", usedBuckets, ht.Capacity)
	fmt.Printf("最大チェーン長: %d\n", maxChain)
	fmt.Printf("平均チェーン長: %.2f\n", avgChain)
	fmt.Printf("負荷率: %.2f\n\n", ht.LoadFactor())

	// 7. 大量データでの性能テスト
	fmt.Println("7. 大量データでの性能テスト")
	bigHt := NewHashTable(100)

	// 1000件のデータを挿入
	measureTime("1000件のデータを挿入", func() {
		for i := 0; i < 1000; i++ {
			key := fmt.Sprintf("key_%d", i)
			value := i * 10
			bigHt.Put(key, value)
		}
	})

	fmt.Printf("大量データ挿入後 - サイズ: %d, 負荷率: %.2f\n", bigHt.GetSize(), bigHt.LoadFactor())

	// 衝突統計
	usedBuckets, maxChain, avgChain = bigHt.GetCollisionStats()
	fmt.Printf("使用中のバケット数: %d/%d\n", usedBuckets, bigHt.Capacity)
	fmt.Printf("最大チェーン長: %d\n", maxChain)
	fmt.Printf("平均チェーン長: %.2f\n", avgChain)

	// 検索性能テスト
	searchTargets := []string{"key_100", "key_500", "key_999", "key_not_exist"}
	for _, target := range searchTargets {
		measureTime(fmt.Sprintf("%sを検索", target), func() {
			_, found := bigHt.Get(target)
			fmt.Printf("  → %s: 見つかった = %t\n", target, found)
		})
	}
	fmt.Println()

	// 8. 苦手な処理の例
	fmt.Println("8. 苦手な処理の例")

	// 全キーの取得（O(n)）
	measureTime("全キーの取得", func() {
		keys := ht.GetAllKeys()
		fmt.Printf("  → 全キー: %v\n", keys)
	})

	// 最小値の検索（O(n)）
	measureTime("最小値の検索", func() {
		minValue := int(^uint(0) >> 1) // int の最大値
		minKey := ""
		for _, bucket := range ht.Buckets {
			current := bucket
			for current != nil {
				if value, ok := current.Value.(int); ok && value < minValue {
					minValue = value
					minKey = current.Key
				}
				current = current.Next
			}
		}
		if minKey != "" {
			fmt.Printf("  → 最小値: %s = %d\n", minKey, minValue)
		}
	})

	fmt.Println("\n=== ハッシュテーブルの特徴まとめ ===")
	fmt.Println("得意な処理（O(1)平均）:")
	fmt.Println("- キーによる検索")
	fmt.Println("- キー・バリューペアの挿入")
	fmt.Println("- キーによる削除")
	fmt.Println("- キーの存在確認")
	fmt.Println()
	fmt.Println("苦手な処理（O(n)）:")
	fmt.Println("- 全要素の走査")
	fmt.Println("- 最小値・最大値の検索")
	fmt.Println("- ソートされた順序での取得")
	fmt.Println("- 範囲検索")
	fmt.Println()
	fmt.Println("注意点:")
	fmt.Println("- 負荷率が高くなると性能が劣化")
	fmt.Println("- ハッシュ関数の品質が性能に大きく影響")
	fmt.Println("- 最悪の場合はO(n)の性能になる可能性")
}