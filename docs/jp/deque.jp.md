# デック（Deque）とは

デック（Deque: Double-Ended Queue）は、キューとスタックの機能を組み合わせた、両端から要素の追加と削除が可能なデータ構造。

配列のランダムアクセスの利点を保ちつつ、両端での操作を効率化した、非常に汎用性の高いデータ構造です。

## 特徴

- **両端操作**: 要素の追加と削除が両端で可能。
- **効率的な操作**: 両端での操作が O(1)で実行可能
- **柔軟性**: スタックやキューとしても利用可能。
- **動的サイズ**: 実行時にサイズを変更可能（実装方法による）。
- **ランダムアクセス**: インデックスによる要素アクセスが可能
- **用途**: スライディングウィンドウ、取り消し機能、パーサー実装など

## 注意点

1. **中間要素へのアクセスが非効率**: deque は両端以外の要素への挿入・削除が非効率的（O(n)）です。
2. **実装による性能差**: 配列ベースと連結リストベースで性能特性が大きく異なります。
3. **メモリレイアウト**: 実装方式によってキャッシュ効率に差が生じます。
4. **ランダムアクセス**: 実装によってはインデックスアクセスが非効率になる場合があります。

# 実装方法

## 実装方式別の比較

| 実装方式             | メリット                                     | デメリット                               | 適用場面                     | 推奨実装   |
| -------------------- | -------------------------------------------- | ---------------------------------------- | ---------------------------- | ---------- |
| **循環配列**         | - メモリ効率が良い<br>- キャッシュ効率が高い | - サイズ制限あり<br>- 実装がやや複雑     | 固定サイズバッファ           | 固定配列   |
| **動的循環配列**     | - 動的サイズ<br>- O(1)両端操作               | - 再割当コスト<br>- メモリ断片化         | 汎用的な deque 実装          | スライス   |
| **双方向連結リスト** | - 真の O(1)挿入削除<br>- メモリ効率          | - キャッシュ効率悪い<br>- ポインタ管理   | 大量の中間操作が必要         | 連結リスト |
| **チャンク配列**     | - バランス良い性能<br>- メモリ効率           | - 実装が複雑<br>- インデックス計算が複雑 | 高性能が必要な大規模システム | スライス   |

## 循環配列・固定配列の実装例

```go
type CircularArrayDeque struct {
    data     [100]int  // 固定サイズ
    front    int
    rear     int
    size     int
    capacity int
}

func NewCircularArrayDeque() *CircularArrayDeque {
    return &CircularArrayDeque{
        capacity: 100,
    }
}

func (d *CircularArrayDeque) PushFront(value int) bool {
    if d.size == d.capacity {
        return false // 満杯
    }

    d.front = (d.front - 1 + d.capacity) % d.capacity
    d.data[d.front] = value
    d.size++
    return true
}

func (d *CircularArrayDeque) PushBack(value int) bool {
    if d.size == d.capacity {
        return false // 満杯
    }

    d.data[d.rear] = value
    d.rear = (d.rear + 1) % d.capacity
    d.size++
    return true
}

func (d *CircularArrayDeque) PopFront() (int, bool) {
    if d.size == 0 {
        return 0, false // 空
    }

    value := d.data[d.front]
    d.front = (d.front + 1) % d.capacity
    d.size--
    return value, true
}

func (d *CircularArrayDeque) PopBack() (int, bool) {
    if d.size == 0 {
        return 0, false // 空
    }

    d.rear = (d.rear - 1 + d.capacity) % d.capacity
    value := d.data[d.rear]
    d.size--
    return value, true
}
```

### 特徴

- **最高のキャッシュ効率**: 連続したメモリレイアウト
- **予測可能な性能**: 全操作が O(1)で一定時間
- **メモリ効率**: 無駄なメモリ使用なし
- **シンプルな実装**: 配列とインデックス操作のみ

### 注意事項

- **サイズ制限**: 固定容量を超えると操作が失敗
- **容量設計**: 事前の適切な容量設計が重要
- **満杯時の処理**: アプリケーション側での対応が必要
- **インデックス計算**: モジュロ演算によるわずかなオーバーヘッド

## 動的循環配列・スライスの実装例

```go
type DynamicCircularDeque struct {
    data  []int
    front int
    rear  int
    size  int
}

func NewDynamicCircularDeque(initialCap int) *DynamicCircularDeque {
    if initialCap < 4 {
        initialCap = 4
    }
    return &DynamicCircularDeque{
        data: make([]int, initialCap),
    }
}

func (d *DynamicCircularDeque) PushFront(value int) {
    if d.size == len(d.data) {
        d.resize()
    }

    d.front = (d.front - 1 + len(d.data)) % len(d.data)
    d.data[d.front] = value
    d.size++
}

func (d *DynamicCircularDeque) PushBack(value int) {
    if d.size == len(d.data) {
        d.resize()
    }

    d.data[d.rear] = value
    d.rear = (d.rear + 1) % len(d.data)
    d.size++
}

func (d *DynamicCircularDeque) resize() {
    newCap := len(d.data) * 2
    newData := make([]int, newCap)

    for i := 0; i < d.size; i++ {
        newData[i] = d.data[(d.front + i) % len(d.data)]
    }

    d.data = newData
    d.front = 0
    d.rear = d.size
}

func (d *DynamicCircularDeque) Get(index int) (int, bool) {
    if index < 0 || index >= d.size {
        return 0, false
    }

    actualIndex := (d.front + index) % len(d.data)
    return d.data[actualIndex], true
}
```

### 特徴

- **動的サイズ**: 実行時の容量調整が可能
- **柔軟性**: 容量を意識せずに使用可能
- **ランダムアクセス**: インデックスによる効率的なアクセス
- **Go 言語との親和性**: スライスの特性を活用

### 注意事項

- **リサイズコスト**: 容量拡張時の一時的な性能劣化
- **メモリ使用量**: 予期しないメモリ増大のリスク
- **GC 負荷**: 頻繁な拡張によるガベージコレクション負荷
- **リサイズ戦略**: 効率的な拡張アルゴリズムの実装が重要

## 双方向連結リスト・連結リストの実装例

```go
type Node struct {
    value int
    prev  *Node
    next  *Node
}

type DoublyLinkedDeque struct {
    head *Node
    tail *Node
    size int
}

func NewDoublyLinkedDeque() *DoublyLinkedDeque {
    return &DoublyLinkedDeque{}
}

func (d *DoublyLinkedDeque) PushFront(value int) {
    newNode := &Node{value: value}

    if d.head == nil {
        d.head = newNode
        d.tail = newNode
    } else {
        newNode.next = d.head
        d.head.prev = newNode
        d.head = newNode
    }
    d.size++
}

func (d *DoublyLinkedDeque) PushBack(value int) {
    newNode := &Node{value: value}

    if d.tail == nil {
        d.head = newNode
        d.tail = newNode
    } else {
        d.tail.next = newNode
        newNode.prev = d.tail
        d.tail = newNode
    }
    d.size++
}

func (d *DoublyLinkedDeque) PopFront() (int, bool) {
    if d.head == nil {
        return 0, false
    }

    value := d.head.value
    d.head = d.head.next

    if d.head != nil {
        d.head.prev = nil
    } else {
        d.tail = nil
    }

    d.size--
    return value, true
}

func (d *DoublyLinkedDeque) PopBack() (int, bool) {
    if d.tail == nil {
        return 0, false
    }

    value := d.tail.value
    d.tail = d.tail.prev

    if d.tail != nil {
        d.tail.next = nil
    } else {
        d.head = nil
    }

    d.size--
    return value, true
}
```

### 特徴

- **真の O(1)操作**: 全ての両端操作が真の定数時間
- **メモリ効率**: 必要な分だけメモリを使用
- **柔軟性**: ノード単位での柔軟な操作が可能
- **無制限サイズ**: メモリが許す限り無制限に拡張可能

### 注意事項

- **キャッシュ効率**: メモリが非連続でキャッシュミスが発生しやすい
- **ポインタ管理**: prev/next ポインタの適切な管理が必要
- **メモリオーバーヘッド**: 各ノードにポインタが必要
- **ランダムアクセス**: インデックスアクセスが O(n)と非効率

## 実装選択の指針

| 要件                     | 推奨実装         | 理由                           |
| ------------------------ | ---------------- | ------------------------------ |
| **キャッシュ効率重視**   | 循環配列         | 最高のメモリ局所性             |
| **動的サイズ必要**       | 動的循環配列     | 容量制限なしで配列の利点を享受 |
| **メモリ効率重視**       | 双方向連結リスト | 必要最小限のメモリ使用         |
| **ランダムアクセス必要** | 動的循環配列     | インデックスアクセスが効率的   |
| **リアルタイム性重視**   | 循環配列         | 最も予測可能な性能             |
| **大量の中間操作**       | 双方向連結リスト | 中間でのノード操作が効率的     |
| **メモリ制約厳しい**     | 循環配列         | 最小限のメモリオーバーヘッド   |
| **汎用的利用**           | 動的循環配列     | バランスの取れた性能と柔軟性   |
