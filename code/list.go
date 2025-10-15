package main

import (
	"fmt"
	"time"
)

// 単方向連結リストのノード
type ListNode struct {
	Data int       // データ部分
	Next *ListNode // 次のノードへのポインタ
}

// 連結リスト構造体
type LinkedList struct {
	Head *ListNode // 先頭ノードへのポインタ
	Size int       // リストのサイズ
}

// 先頭への挿入 O(1)
func (list *LinkedList) InsertAtHead(data int) {
	newNode := &ListNode{Data: data, Next: list.Head}
	list.Head = newNode
	list.Size++
}

// 末尾への挿入 O(n)
func (list *LinkedList) InsertAtTail(data int) {
	newNode := &ListNode{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	list.Size++
}

// 指定位置への挿入 O(n)
func (list *LinkedList) InsertAtIndex(index int, data int) bool {
	if index < 0 || index > list.Size {
		return false
	}

	if index == 0 {
		list.InsertAtHead(data)
		return true
	}

	newNode := &ListNode{Data: data}
	current := list.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	list.Size++
	return true
}

// 先頭からの削除 O(1)
func (list *LinkedList) DeleteFromHead() bool {
	if list.Head == nil {
		return false
	}

	list.Head = list.Head.Next
	list.Size--
	return true
}

// 末尾からの削除 O(n)
func (list *LinkedList) DeleteFromTail() bool {
	if list.Head == nil {
		return false
	}

	if list.Head.Next == nil {
		list.Head = nil
		list.Size--
		return true
	}

	current := list.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
	list.Size--
	return true
}

// 指定位置の削除 O(n)
func (list *LinkedList) DeleteAtIndex(index int) bool {
	if index < 0 || index >= list.Size || list.Head == nil {
		return false
	}

	if index == 0 {
		return list.DeleteFromHead()
	}

	current := list.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
		list.Size--
		return true
	}
	return false
}

// 検索 O(n)
func (list *LinkedList) Search(target int) *ListNode {
	current := list.Head
	for current != nil {
		if current.Data == target {
			return current
		}
		current = current.Next
	}
	return nil
}

// インデックスでのアクセス O(n)
func (list *LinkedList) GetAtIndex(index int) *ListNode {
	if index < 0 || index >= list.Size {
		return nil
	}

	current := list.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current
}

// 全要素の表示 O(n)
func (list *LinkedList) Display() {
	fmt.Print("リスト: ")
	current := list.Head
	for current != nil {
		fmt.Printf("%d", current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Print(" -> nil")
	fmt.Printf(" (サイズ: %d)\n", list.Size)
}

// リストが空かどうかを確認
func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

// リストのサイズを取得
func (list *LinkedList) GetSize() int {
	return list.Size
}

// 計算量を測定するためのテスト関数
func measureTime(operation string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s の実行時間: %v\n", operation, duration)
}

func main() {
	fmt.Println("=== 連結リストの基本操作と計算量の確認 ===\n")

	list := &LinkedList{}

	// 1. 先頭への挿入操作（O(1)）
	fmt.Println("1. 先頭への挿入操作（O(1)）")
	measureTime("先頭に1を挿入", func() {
		list.InsertAtHead(1)
	})
	list.Display()

	measureTime("先頭に2を挿入", func() {
		list.InsertAtHead(2)
	})
	list.Display()

	measureTime("先頭に3を挿入", func() {
		list.InsertAtHead(3)
	})
	list.Display()
	fmt.Println()

	// 2. 末尾への挿入操作（O(n)）
	fmt.Println("2. 末尾への挿入操作（O(n)）")
	measureTime("末尾に4を挿入", func() {
		list.InsertAtTail(4)
	})
	list.Display()

	measureTime("末尾に5を挿入", func() {
		list.InsertAtTail(5)
	})
	list.Display()
	fmt.Println()

	// 3. 中間位置への挿入操作（O(n)）
	fmt.Println("3. 中間位置への挿入操作（O(n)）")
	measureTime("インデックス2に99を挿入", func() {
		list.InsertAtIndex(2, 99)
	})
	list.Display()
	fmt.Println()

	// 4. インデックスでのアクセス（O(n)）
	fmt.Println("4. インデックスでのアクセス（O(n)）")
	var node *ListNode
	measureTime("インデックス0の要素取得", func() {
		node = list.GetAtIndex(0)
	})
	if node != nil {
		fmt.Printf("インデックス0の要素: %d\n", node.Data)
	}

	measureTime("インデックス3の要素取得", func() {
		node = list.GetAtIndex(3)
	})
	if node != nil {
		fmt.Printf("インデックス3の要素: %d\n", node.Data)
	}
	fmt.Println()

	// 5. 要素の検索（O(n)）
	fmt.Println("5. 要素の検索（O(n)）")
	var foundNode *ListNode
	measureTime("値99を検索", func() {
		foundNode = list.Search(99)
	})
	if foundNode != nil {
		fmt.Printf("値99が見つかりました: %d\n", foundNode.Data)
	}

	measureTime("値999を検索（存在しない）", func() {
		foundNode = list.Search(999)
	})
	if foundNode == nil {
		fmt.Println("値999は見つかりませんでした")
	}
	fmt.Println()

	// 6. 先頭からの削除（O(1)）
	fmt.Println("6. 先頭からの削除（O(1)）")
	measureTime("先頭から削除", func() {
		list.DeleteFromHead()
	})
	list.Display()
	fmt.Println()

	// 7. 末尾からの削除（O(n)）
	fmt.Println("7. 末尾からの削除（O(n)）")
	measureTime("末尾から削除", func() {
		list.DeleteFromTail()
	})
	list.Display()
	fmt.Println()

	// 8. 中間位置からの削除（O(n)）
	fmt.Println("8. 中間位置からの削除（O(n)）")
	measureTime("インデックス1から削除", func() {
		list.DeleteAtIndex(1)
	})
	list.Display()
	fmt.Println()

	// 9. 大量データでの性能比較
	fmt.Println("9. 大量データでの性能比較")
	bigList := &LinkedList{}

	// 先頭への1000回挿入（各操作がO(1)）
	measureTime("先頭への1000回挿入", func() {
		for i := 0; i < 1000; i++ {
			bigList.InsertAtHead(i)
		}
	})

	// 末尾への100回挿入（各操作がO(n)）
	measureTime("末尾への100回挿入", func() {
		for i := 0; i < 100; i++ {
			bigList.InsertAtTail(i + 1000)
		}
	})

	fmt.Printf("大きなリストのサイズ: %d\n", bigList.GetSize())

	// インデックス500での検索（O(n)）
	measureTime("インデックス500の要素取得", func() {
		node = bigList.GetAtIndex(500)
	})
	if node != nil {
		fmt.Printf("インデックス500の要素: %d\n", node.Data)
	}

	fmt.Println("\n=== 連結リストの特徴まとめ ===")
	fmt.Println("得意な処理（O(1)）:")
	fmt.Println("- 先頭への挿入・削除")
	fmt.Println("- ノードの参照がある場合の挿入・削除")
	fmt.Println()
	fmt.Println("苦手な処理（O(n)）:")
	fmt.Println("- インデックスでのアクセス")
	fmt.Println("- 末尾への挿入・削除")
	fmt.Println("- 中間位置での操作")
	fmt.Println("- 要素の検索")
}
