# デック（Deque）とは

デック（Deque: Double-Ended Queue）は、両端から要素の追加と削除が可能なデータ構造です。キューとスタックの機能を組み合わせた柔軟なデータ構造として設計されています。

**このデータ構造を使う一番の利点は、先頭と末尾の両方から効率的にデータの追加・削除ができる点です。**

配列のランダムアクセスの利点を保ちつつ、両端での操作を効率化した、非常に汎用性の高いデータ構造です。

## 基本的な特徴

- **両端操作**: 要素の追加と削除が両端で可能。
- **効率的な操作**: 両端での操作が O(1)で実行可能
- **柔軟性**: スタックやキューとしても利用可能。
- **動的サイズ**: 実行時にサイズを変更可能（実装方法による）。
- **ランダムアクセス**: インデックスによる要素アクセスが可能
- **用途**: スライディングウィンドウ、取り消し機能、パーサー実装など

# 行える処理

## Deque（Double-ended Queue / 両端キュー）

| 機能            | 説明                             | 計算量 | 戻り値                | 得意/苦手 | 補足（その他）                             |
| --------------- | -------------------------------- | ------ | --------------------- | --------- | ------------------------------------------ |
| PushFront       | 先頭に要素を追加                 | O(1)   | なし（void）          | ✅        | deque の最大の利点。配列では非効率な操作。 |
| PushBack        | 末尾に要素を追加                 | O(1)   | なし（void）          | ✅        | 通常のキューと同様の操作。                 |
| PopFront        | 先頭から要素を削除して返す       | O(1)   | 削除された要素        | ✅        | キューの基本操作。                         |
| PopBack         | 末尾から要素を削除して返す       | O(1)   | 削除された要素        | ✅        | スタックの基本操作。                       |
| Front/PeekFront | 先頭要素を削除せずに参照         | O(1)   | 要素の値              | ✅        | 次に取り出される要素を確認。               |
| Back/PeekBack   | 末尾要素を削除せずに参照         | O(1)   | 要素の値              | ✅        | 最後に追加された要素を確認。               |
| Get/Access      | インデックスを指定して要素を取得 | O(1)   | 要素の値              | ✅        | ランダムアクセス可能（実装による）。       |
| Set/Update      | インデックスを指定して要素を更新 | O(1)   | なし（void）          | ✅        | 既存要素の値変更。                         |
| Size/Length     | 現在の要素数を取得               | O(1)   | int（要素数）         | ✅        | 内部カウンタで管理。                       |
| IsEmpty         | deque が空かを確認               | O(1)   | bool（空/データあり） | ✅        | サイズが 0 かをチェック。                  |
| Clear           | 全要素を削除                     | O(1)   | なし（void）          | ✅        | インデックスとカウンタをリセット。         |
| Insert          | 指定位置に要素を挿入             | O(n)   | なし（void）          | ❌        | 中間挿入は要素のシフトが必要。             |
| Remove/Delete   | 指定位置の要素を削除             | O(n)   | 削除された要素        | ❌        | 中間削除は要素のシフトが必要。             |
| Search/Find     | 指定した値を検索                 | O(n)   | インデックスまたは-1  | ❌        | 全要素を順次チェックする必要あり。         |
| Contains        | 指定値が存在するかチェック       | O(n)   | bool（存在/非存在）   | ❌        | 線形検索が必要。                           |
| IndexOf         | 指定値の最初のインデックスを取得 | O(n)   | インデックスまたは-1  | ❌        | 線形検索が必要。                           |
| Reverse         | 要素の順序を逆順にする           | O(n)   | なし（void）          | -         | 全要素の入れ替えまたは論理的な反転。       |
| Copy/Clone      | deque の複製を作成               | O(n)   | 新しい deque          | -         | 全要素をコピーする必要あり。               |
| ToArray         | 配列に変換                       | O(n)   | 配列                  | -         | 先頭から末尾まで順序を保って配列化。       |
| Slice/SubDeque  | 指定範囲の部分 deque を取得      | O(k)   | 新しい deque          | ✅        | k は範囲のサイズ。                         |

## Deque の特殊操作

| 機能        | 説明                                   | 計算量 | 戻り値       | 得意/苦手 | 補足（その他）                     |
| ----------- | -------------------------------------- | ------ | ------------ | --------- | ---------------------------------- |
| RotateLeft  | 要素を左に回転（先頭要素を末尾に移動） | O(1)   | なし（void） | ✅        | deque の先頭・末尾操作を活用。     |
| RotateRight | 要素を右に回転（末尾要素を先頭に移動） | O(1)   | なし（void） | ✅        | deque の先頭・末尾操作を活用。     |
| ExtendLeft  | 別のコレクションの要素を先頭側に追加   | O(k)   | なし（void） | ✅        | k は追加要素数。一括処理で効率的。 |
| ExtendRight | 別のコレクションの要素を末尾側に追加   | O(k)   | なし（void） | ✅        | k は追加要素数。一括処理で効率的。 |
| MaxLength   | 最大長制限付き deque での要素追加      | O(1)   | なし（void） | ✅        | 古い要素を自動削除してサイズ制限。 |
| Shrink      | 不要な容量を削除してメモリを節約       | O(n)   | なし（void） | -         | メモリ使用量の最適化。             |

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

# 他のデータ構造との比較

| 特徴                 | Deque（循環配列） | Deque（連結リスト） | 配列（スライス）  | リングバッファ | キュー        | スタック      | ハッシュテーブル      |
| -------------------- | ----------------- | ------------------- | ----------------- | -------------- | ------------- | ------------- | --------------------- |
| **メモリ使用量**     | ✅ 効率的         | ❌ ポインタ分重い   | ✅ 効率的         | ✅ 最効率      | ⚠️ 実装に依存 | ✅ 効率的     | ⚠️ 負荷率に依存       |
| **キャッシュ効率**   | ✅ 最適           | ❌ 非連続メモリ     | ✅ 最適           | ✅ 最適        | ⚠️ 実装に依存 | ✅ 最適       | ✅ 良好               |
| **先頭挿入**         | ✅ O(1)           | ✅ O(1)             | ❌ O(n)シフト     | ❌ 不可        | ❌ 不可       | ❌ 不可       | ❌ 概念なし           |
| **末尾挿入**         | ✅ O(1)           | ✅ O(1)             | ✅ O(1)償却       | ✅ O(1)        | ✅ O(1)       | ✅ O(1)       | ❌ 概念なし           |
| **先頭削除**         | ✅ O(1)           | ✅ O(1)             | ❌ O(n)シフト     | ✅ O(1)        | ✅ O(1)       | ❌ 不可       | ❌ 概念なし           |
| **末尾削除**         | ✅ O(1)           | ✅ O(1)             | ✅ O(1)           | ❌ 不可        | ❌ 不可       | ✅ O(1)       | ❌ 概念なし           |
| **ランダムアクセス** | ✅ O(1)           | ❌ O(n)             | ✅ O(1)           | ✅ O(1)        | ❌ 先頭のみ   | ❌ トップのみ | ✅ O(1)平均           |
| **キー検索**         | ❌ O(n)線形       | ❌ O(n)線形         | ❌ O(n)線形       | ❌ O(n)線形    | ❌ O(n)線形   | ❌ O(n)線形   | ✅ O(1)平均           |
| **容量制限**         | ⚠️ 実装に依存     | ✅ 無制限           | ❌ 実質無制限     | ✅ 厳密な上限  | ⚠️ 実装に依存 | ⚠️ 実装に依存 | ❌ 動的拡張           |
| **順序保持**         | ✅ 両端操作順     | ✅ 両端操作順       | ✅ インデックス順 | ✅ FIFO        | ✅ FIFO       | ✅ LIFO       | ❌ なし               |
| **実装複雑さ**       | ⚠️ 中程度         | ⚠️ 中程度           | ✅ シンプル       | ✅ シンプル    | ✅ シンプル   | ✅ シンプル   | ❌ 複雑               |
| **リアルタイム性**   | ✅ 高い           | ⚠️ GC 影響あり      | ❌ 再割当で変動   | ✅ 最高        | ⚠️ 実装に依存 | ✅ 高い       | ❌ ハッシュ衝突で変動 |
| **用途の柔軟性**     | ✅ 非常に高い     | ✅ 非常に高い       | ✅ 高い           | ❌ 制限的      | ❌ 制限的     | ❌ 制限的     | ❌ 制限的             |

# パフォーマンス特性

| 操作パターン           | Deque（循環配列） | Deque（連結リスト）   | 配列            | リングバッファ |
| ---------------------- | ----------------- | --------------------- | --------------- | -------------- |
| **大量の小データ処理** | ✅ 最適           | ❌ GC 負荷            | ✅ 最適         | ✅ 最適        |
| **一定レート処理**     | ✅ 最適           | ⚠️ メモリ断片化       | ❌ 再割当変動   | ✅ 最適        |
| **バースト処理**       | ✅ 柔軟対応       | ✅ 柔軟対応           | ✅ 自動拡張     | ❌ 容量制限    |
| **長時間稼働**         | ✅ 安定           | ⚠️ メモリリーク可能性 | ❌ メモリ増大   | ✅ 最適        |
| **頻繁な両端操作**     | ✅ 最適           | ✅ 最適               | ❌ 先頭操作重い | ⚠️ 片端のみ    |
| **順次アクセス**       | ✅ 最適           | ⚠️ キャッシュミス     | ✅ 最適         | ✅ 最適        |
