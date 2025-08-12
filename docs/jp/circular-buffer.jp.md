# リングバッファ（Circular Buffer）とは

リングバッファは、固定サイズの配列を使用してデータを循環的に管理するデータ構造です。配列の末尾が先頭に接続されるように設計されています。

**このデータ構造を使う一番の利点は、メモリを効率的に使用しながら、一定のサイズ内でデータを循環させることができる点です。**

固定長の配列をリング状に扱い、古いデータを自動で上書きしていく、キューの一種のデータ構造です。

# 基本的な特徴

- **固定サイズ**: メモリ使用量が一定。
- **効率的なメモリ利用**: 配列を再利用するため、メモリ割り当てや解放のオーバーヘッドが少ない。
- **FIFO 原則**: キューと同様に、先入れ先出しのデータ構造。
- **用途**: 音声処理、ネットワークバッファ、リアルタイムシステムなど。

# 行える処理

## リングバッファ（Circular Buffer）

| 機能         | 説明                                     | 計算量 | 戻り値                | 得意/苦手 | 補足（その他）                                       |
| ------------ | ---------------------------------------- | ------ | --------------------- | --------- | ---------------------------------------------------- |
| Enqueue/Push | バッファの末尾に要素を追加               | O(1)   | bool（成功/失敗）     | ✅        | 満杯時は失敗またはオーバーライト。                   |
| Dequeue/Pop  | バッファの先頭から要素を取り出し         | O(1)   | 要素の値              | ✅        | 空の場合はエラーまたはゼロ値を返す。                 |
| Front/Peek   | 先頭要素を削除せずに参照                 | O(1)   | 要素の値              | ✅        | キューの次に取り出される要素を確認。                 |
| Rear/Back    | 末尾要素を削除せずに参照                 | O(1)   | 要素の値              | ✅        | 最後に追加された要素を確認。                         |
| IsFull       | バッファが満杯かを確認                   | O(1)   | bool（満杯/空きあり） | ✅        | 書き込み前のチェックに使用。                         |
| IsEmpty      | バッファが空かを確認                     | O(1)   | bool（空/データあり） | ✅        | 読み取り前のチェックに使用。                         |
| Size/Count   | 現在の要素数を取得                       | O(1)   | int（要素数）         | ✅        | 内部カウンタまたは計算により取得。                   |
| Capacity     | バッファの最大容量を取得                 | O(1)   | int（容量）           | ✅        | 固定値のため定数時間。                               |
| Clear/Reset  | バッファを空の状態にリセット             | O(1)   | なし（void）          | ✅        | インデックスとカウンタをリセット。                   |
| Get/Access   | インデックスを指定して要素を取得         | O(1)   | 要素の値              | ✅        | 論理インデックスから物理インデックスへの変換が必要。 |
| Set/Update   | インデックスを指定して要素を更新         | O(1)   | なし（void）          | ✅        | 既存要素の値変更。範囲チェックが必要。               |
| Overwrite    | 満杯時に古い要素を上書きして新要素を追加 | O(1)   | なし（void）          | ✅        | データロスが発生するが常に成功。                     |
| Search/Find  | 指定した値を検索                         | O(n)   | インデックスまたは-1  | ❌        | 全要素を順次チェックする必要あり。                   |
| Copy         | バッファの複製を作成                     | O(n)   | 新しいバッファ        | -         | 全要素をコピーする必要あり。                         |
| ToArray      | 論理順序で線形配列に変換                 | O(n)   | 配列                  | -         | 先頭から末尾まで順序を保って配列化。                 |
| Insert       | 任意位置に要素を挿入                     | O(n)   | bool（成功/失敗）     | ❌        | 要素のシフトが必要で非効率。                         |
| Remove       | 任意位置の要素を削除                     | O(n)   | 削除された要素        | ❌        | 要素のシフトが必要で非効率。                         |
| Resize       | バッファサイズを変更                     | O(n)   | bool（成功/失敗）     | ❌        | 新しいバッファに全要素をコピー。                     |

## リングバッファの特殊操作

| 機能           | 説明                                     | 計算量 | 戻り値              | 得意/苦手 | 補足（その他）                             |
| -------------- | ---------------------------------------- | ------ | ------------------- | --------- | ------------------------------------------ |
| BulkEnqueue    | 複数要素を一度に追加                     | O(k)   | int（追加された数） | ✅        | k は追加要素数。メモリコピーで効率化可能。 |
| BulkDequeue    | 複数要素を一度に取り出し                 | O(k)   | 要素の配列          | ✅        | k は取り出し要素数。一括処理で効率的。     |
| Advance        | 先頭ポインタを進める（要素を破棄）       | O(1)   | なし（void）        | ✅        | 実際の読み取りなしで要素をスキップ。       |
| Retreat        | 先頭ポインタを戻す（読み取りを取り消し） | O(1)   | なし（void）        | ✅        | 直前の操作を取り消し（実装による）。       |
| Snapshot       | 現在の状態のスナップショットを取得       | O(1)   | 状態情報            | ✅        | インデックスとカウンタの現在値を保存。     |
| AvailableSpace | 利用可能な空き容量を取得                 | O(1)   | int（空き容量）     | ✅        | 容量 - 現在の要素数。                      |
| UsedSpace      | 使用中の容量を取得                       | O(1)   | int（使用容量）     | ✅        | 現在の要素数と同義。                       |
| ForEach        | 全要素に対して関数を適用                 | O(n)   | なし（void）        | -         | 論理順序で要素を処理。                     |

## 注意点

1. **サイズ制限**: リングバッファは固定サイズのため、サイズを超えるとデータが上書きされる可能性があります。
2. **空の状態**: バッファが空の状態で `Dequeue` や `Peek` を呼び出すとエラーになるため、事前に空かどうかを確認する必要があります。
3. **動的サイズ変更が不可**: サイズを変更する場合は新しいバッファを作成する必要があります。

# 実装方法

## 実装パターン別の比較

| 実装パターン       | 満杯時の動作         | メリット                                       | デメリット                           | 使用場面                       | 推奨実装                                         |
| ------------------ | -------------------- | ---------------------------------------------- | ------------------------------------ | ------------------------------ | ------------------------------------------------ |
| **上書き型**       | 古いデータを上書き   | データロスしても動作継続<br>常に最新データ保持 | データロスが発生<br>ロスト検出が困難 | ログ、リアルタイム監視         | 固定配列(シンプルさ・パフォーマンス重視)         |
| **拒否型**         | 新しいデータを拒否   | データロスなし<br>明確なエラーハンドリング     | 書き込み失敗の処理が必要<br>性能劣化 | 重要データ処理、バッチ処理     | 固定配列(明確な容量制限が必要)                   |
| **ブロッキング型** | 空きができるまで待機 | データロスなし<br>フロー制御が自動             | デッドロックリスク<br>応答性劣化     | プロデューサー・コンシューマー | 固定配列 or スライス(同期プリミティブとの親和性) |
| **動的拡張型**     | バッファサイズ拡張   | 柔軟性が高い<br>データロスなし                 | メモリ使用量増加<br>実装が複雑       | 可変長データ処理               | スライス or 連結リスト(動的サイズ変更が必須)     |

## 上書き型・固定配列の実装例

```go
type OverwriteBuffer struct {
    data     [100]int  // 固定サイズ
    head     int
    tail     int
    isFull   bool
}

func (b *OverwriteBuffer) Enqueue(value int) {
    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)

    if b.isFull {
        b.head = (b.head + 1) % len(b.data) // 古いデータを上書き
    }

    b.isFull = (b.tail == b.head)
}
```

### 特徴

- **最もシンプルな実装**: head/tail ポインタのみで制御
- **最高のパフォーマンス**: 分岐処理が最小限
- **予測可能な動作**: 常に一定時間で処理完了
- **メモリ効率**: 無駄なメモリ使用なし

### 注意事項

- **データロスの検出困難**: 上書きされたデータの追跡が難しい
- **ロストデータ通知**: アプリケーション側での対策が必要
- **並行アクセス**: 単純な実装では競合状態が発生しやすい
- **デバッグの困難さ**: データロスによるバグの原因特定が困難

### 上書き型の改良実装例（ロスト検出付き）

```go
type OverwriteBufferWithLossDetection struct {
    data      [100]int
    head      int
    tail      int
    isFull    bool
    lostCount int64  // ロストしたデータの数
    mu        sync.RWMutex
}

func (b *OverwriteBufferWithLossDetection) Enqueue(value int) int64 {
    b.mu.Lock()
    defer b.mu.Unlock()

    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)

    var lost int64
    if b.isFull {
        b.head = (b.head + 1) % len(b.data)
        b.lostCount++
        lost = b.lostCount
    }

    b.isFull = (b.tail == b.head)
    return lost
}
```

## 拒否型・固定配列の実装例

```go
type RejectBuffer struct {
    data     [100]int
    head     int
    tail     int
    size     int
}

func (b *RejectBuffer) Enqueue(value int) bool {
    if b.size == len(b.data) {
        return false // 拒否
    }

    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)
    b.size++
    return true
}
```

### 特徴

- **データ保護**: 重要なデータのロスを完全に防止
- **明確なエラー処理**: 失敗時の処理が明確
- **バックプレッシャー**: 自然なフロー制御メカニズム
- **デバッグの容易さ**: 問題の原因が特定しやすい

### 注意事項

- **エラーハンドリング**: 拒否時の処理が必須
- **デッドロック**: 書き込み失敗による処理停止リスク
- **性能劣化**: 満杯状態での頻繁な失敗による性能低下
- **リトライ機構**: アプリケーション側での再試行戦略が必要

### 拒否型の改良実装例（統計情報付き）

```go
type RejectBufferWithStats struct {
    data         [100]int
    head         int
    tail         int
    size         int
    rejectCount  int64
    enqueueCount int64
    mu           sync.RWMutex
}

func (b *RejectBufferWithStats) Enqueue(value int) (bool, *Stats) {
    b.mu.Lock()
    defer b.mu.Unlock()

    if b.size == len(b.data) {
        b.rejectCount++
        return false, &Stats{
            Rejects:  b.rejectCount,
            Enqueues: b.enqueueCount,
            FullRate: float64(b.rejectCount) / float64(b.enqueueCount + b.rejectCount),
        }
    }

    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)
    b.size++
    b.enqueueCount++

    return true, nil
}

type Stats struct {
    Rejects  int64
    Enqueues int64
    FullRate float64
}
```

## ブロッキング型・スライスの実装例

```go
type BlockingBuffer struct {
    data     []int
    head     int
    tail     int
    size     int
    capacity int
    mu       sync.Mutex
    notEmpty *sync.Cond
    notFull  *sync.Cond
}

func (b *BlockingBuffer) Enqueue(value int) {
    b.mu.Lock()
    defer b.mu.Unlock()

    for b.size == b.capacity {
        b.notFull.Wait() // ブロッキング
    }

    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)
    b.size++
    b.notEmpty.Signal()
}
```

### 特徴

- **データ保護**: データロス完全防止
- **自動フロー制御**: プロデューサーの速度を自動調整
- **簡潔な API**: 呼び出し側でのエラー処理が不要
- **並行安全**: 適切な同期機構で安全な並行アクセス

### 注意事項

- **デッドロックリスク**: 不適切な同期でデッドロック発生
- **応答性劣化**: 長時間ブロックによるシステム応答性低下
- **タイムアウト設計**: 無限待機を避けるタイムアウト機構が必要
- **キャンセレーション**: Context による操作キャンセル対応が重要

### ブロッキング型の改良実装例（タイムアウト・Context 対応）

```go
type BlockingBufferWithTimeout struct {
    data     []int
    head     int
    tail     int
    size     int
    capacity int
    mu       sync.Mutex
    notEmpty *sync.Cond
    notFull  *sync.Cond
}

func (b *BlockingBufferWithTimeout) EnqueueWithTimeout(
    ctx context.Context,
    value int,
    timeout time.Duration,
) error {
    // タイムアウト付きContext作成
    timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()

    done := make(chan error, 1)

    go func() {
        b.mu.Lock()
        defer b.mu.Unlock()

        for b.size == b.capacity {
            // Contextキャンセルチェック
            select {
            case <-timeoutCtx.Done():
                done <- timeoutCtx.Err()
                return
            default:
            }
            b.notFull.Wait()
        }

        b.data[b.tail] = value
        b.tail = (b.tail + 1) % len(b.data)
        b.size++
        b.notEmpty.Signal()
        done <- nil
    }()

    select {
    case err := <-done:
        return err
    case <-timeoutCtx.Done():
        return timeoutCtx.Err()
    }
}
```

## 動的拡張型・スライスの実装例

```go
type DynamicBuffer struct {
    data []int
    head int
    tail int
    size int
}

func (b *DynamicBuffer) Enqueue(value int) {
    if b.size == len(b.data) {
        b.resize() // 動的拡張
    }

    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)
    b.size++
}

func (b *DynamicBuffer) resize() {
    newSize := len(b.data) * 2
    if newSize == 0 {
        newSize = 1
    }

    newData := make([]int, newSize)
    for i := 0; i < b.size; i++ {
        newData[i] = b.data[(b.head+i)%len(b.data)]
    }

    b.data = newData
    b.head = 0
    b.tail = b.size
}
```

### 特徴

- **柔軟性**: 実行時のサイズ調整が可能
- **データ保護**: 容量不足によるデータロスなし
- **簡単な使用**: 容量を意識せずに使用可能
- **Go 言語との親和性**: スライスの特性を活用

### 注意事項

- **メモリ使用量**: 予期しないメモリ増大のリスク
- **GC 負荷**: 頻繁な拡張によるガベージコレクション負荷
- **性能の不安定性**: リサイズ時の一時的な性能劣化
- **リサイズ戦略**: 効率的な拡張アルゴリズムの実装が重要

### 動的拡張型の改良実装例（拡張戦略・制限付き）

```go
type DynamicBufferWithLimits struct {
    data        []int
    head        int
    tail        int
    size        int
    maxCapacity int     // 最大容量制限
    growthRate  float64 // 成長率
    shrinkCount int     // 縮小カウンタ
    mu          sync.RWMutex
}

func NewDynamicBufferWithLimits(initialCap, maxCap int, growthRate float64) *DynamicBufferWithLimits {
    return &DynamicBufferWithLimits{
        data:        make([]int, initialCap),
        maxCapacity: maxCap,
        growthRate:  growthRate,
    }
}

func (b *DynamicBufferWithLimits) Enqueue(value int) error {
    b.mu.Lock()
    defer b.mu.Unlock()

    if b.size == len(b.data) {
        if err := b.resize(); err != nil {
            return fmt.Errorf("resize failed: %w", err)
        }
    }

    b.data[b.tail] = value
    b.tail = (b.tail + 1) % len(b.data)
    b.size++
    b.shrinkCount = 0 // リセット

    return nil
}

func (b *DynamicBufferWithLimits) resize() error {
    currentCap := len(b.data)
    newCap := int(float64(currentCap) * b.growthRate)

    if newCap > b.maxCapacity {
        if currentCap >= b.maxCapacity {
            return fmt.Errorf("maximum capacity %d reached", b.maxCapacity)
        }
        newCap = b.maxCapacity
    }

    newData := make([]int, newCap)
    for i := 0; i < b.size; i++ {
        newData[i] = b.data[(b.head+i)%len(b.data)]
    }

    b.data = newData
    b.head = 0
    b.tail = b.size

    return nil
}

// 自動縮小機能
func (b *DynamicBufferWithLimits) Dequeue() (int, error) {
    b.mu.Lock()
    defer b.mu.Unlock()

    if b.size == 0 {
        return 0, fmt.Errorf("buffer is empty")
    }

    value := b.data[b.head]
    b.head = (b.head + 1) % len(b.data)
    b.size--

    // 縮小判定
    if b.size < len(b.data)/4 && len(b.data) > 4 {
        b.shrinkCount++
        if b.shrinkCount > 10 { // 連続して小さい状態が続いた場合のみ縮小
            b.shrink()
        }
    }

    return value, nil
}

func (b *DynamicBufferWithLimits) shrink() {
    newCap := len(b.data) / 2
    if newCap < 4 {
        newCap = 4
    }

    newData := make([]int, newCap)
    for i := 0; i < b.size; i++ {
        newData[i] = b.data[(b.head+i)%len(b.data)]
    }

    b.data = newData
    b.head = 0
    b.tail = b.size
    b.shrinkCount = 0
}
```

## 実装選択の指針

| 要件                   | 推奨実装               | 理由                               |
| ---------------------- | ---------------------- | ---------------------------------- |
| **リアルタイム性重視** | 上書き型               | 最も予測可能な性能                 |
| **データ整合性重視**   | 拒否型・ブロッキング型 | データロス完全防止                 |
| **スループット重視**   | 上書き型・拒否型       | ブロッキングなしで高スループット   |
| **メモリ制約厳しい**   | 固定配列実装           | 最小限のメモリ使用                 |
| **可変長データ対応**   | 動的拡張型             | 実行時のサイズ変更が必要           |
| **エラー処理の簡素化** | ブロッキング型         | 呼び出し側のエラー処理が最小限     |
| **デバッグの容易さ**   | 拒否型                 | 問題の原因特定が容易               |
| **並行性重視**         | ブロッキング型         | 適切な同期機構で安全な並行アクセス |

# 他のデータ構造との比較

| 特徴                 | 上書き型リングバッファ | 拒否型リングバッファ | 静的配列          | 動的配列（スライス）    | 通常のキュー    | 連結リストキュー  | Deque           | ハッシュテーブル      |
| -------------------- | ---------------------- | -------------------- | ----------------- | ----------------------- | --------------- | ----------------- | --------------- | --------------------- |
| **メモリ使用量**     | ✅ 固定・最効率        | ✅ 固定・最効率      | ✅ 固定・最効率   | ❌ 動的・オーバーヘッド | ❌ 動的・変動   | ❌ ポインタ分重い | ⚠️ 実装に依存   | ⚠️ 負荷率に依存       |
| **キャッシュ効率**   | ✅ 最適                | ✅ 最適              | ✅ 最適           | ✅ 良好                 | ❌ 非連続メモリ | ❌ 非連続メモリ   | ✅ 良好         | ✅ 良好               |
| **ランダムアクセス** | ✅ O(1)                | ✅ O(1)              | ✅ O(1)           | ✅ O(1)                 | ❌ 先頭のみ     | ❌ 不可           | ✅ O(1)         | ✅ O(1)平均           |
| **先頭挿入**         | ❌ 不可                | ❌ 不可              | ❌ 不可           | ❌ O(n)シフト           | ❌ 不可         | ✅ O(1)           | ✅ O(1)         | ❌ 概念なし           |
| **末尾挿入**         | ✅ O(1)                | ⚠️ O(1)満杯時拒否    | ❌ 不可           | ✅ O(1)償却             | ✅ O(1)         | ✅ O(1)           | ✅ O(1)         | ❌ 概念なし           |
| **先頭削除**         | ✅ O(1)                | ✅ O(1)              | ❌ 不可           | ❌ O(n)シフト           | ✅ O(1)         | ✅ O(1)           | ✅ O(1)         | ❌ 概念なし           |
| **末尾削除**         | ❌ 不可                | ❌ 不可              | ❌ 不可           | ✅ O(1)                 | ❌ 不可         | ✅ O(1)           | ✅ O(1)         | ❌ 概念なし           |
| **キー検索**         | ❌ O(n)線形            | ❌ O(n)線形          | ❌ O(n)線形       | ❌ O(n)線形             | ❌ O(n)線形     | ❌ O(n)線形       | ❌ O(n)線形     | ✅ O(1)平均           |
| **容量制限**         | ✅ 厳密な上限          | ✅ 厳密な上限        | ✅ 厳密な上限     | ❌ 実質無制限           | ❌ メモリ次第   | ❌ メモリ次第     | ❌ 実装に依存   | ❌ 動的拡張           |
| **データロス**       | ❌ 上書き時発生        | ✅ なし              | ✅ なし           | ✅ なし                 | ✅ なし         | ✅ なし           | ✅ なし         | ✅ なし               |
| **順序保持**         | ✅ FIFO                | ✅ FIFO              | ✅ インデックス順 | ✅ インデックス順       | ✅ FIFO         | ✅ FIFO           | ✅ 両端操作順   | ❌ なし               |
| **実装複雑さ**       | ✅ 最シンプル          | ✅ シンプル          | ✅ 最シンプル     | ✅ シンプル             | ✅ シンプル     | ⚠️ 中程度         | ⚠️ 中程度       | ❌ 複雑               |
| **リアルタイム性**   | ✅ 最高（予測可能）    | ✅ 高い              | ✅ 最高           | ❌ 再割当で変動         | ⚠️ メモリ次第   | ⚠️ メモリ次第     | ⚠️ 実装に依存   | ❌ ハッシュ衝突で変動 |
| **並行安全性**       | ❌ 追加実装必要        | ❌ 追加実装必要      | ❌ 追加実装必要   | ❌ 追加実装必要         | ❌ 追加実装必要 | ❌ 追加実装必要   | ❌ 追加実装必要 | ❌ 追加実装必要       |

# パフォーマンス特性

| 操作パターン           | リングバッファの優位性                         |
| ---------------------- | ---------------------------------------------- |
| **大量の小データ処理** | メモリ断片化なし、ガベージコレクション負荷軽減 |
| **一定レート処理**     | 予測可能なメモリ使用量、リアルタイム性保証     |
| **バースト処理**       | 固定バッファサイズによる負荷制限               |
| **長時間稼働**         | メモリリークなし、安定した性能                 |
