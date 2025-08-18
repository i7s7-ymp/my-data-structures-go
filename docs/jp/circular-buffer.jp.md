# リングバッファ（Circular Buffer）とは

リングバッファは、固定サイズの配列を使用してデータを循環的に管理するデータ構造です。配列の末尾が先頭に接続されるように設計されています。

**このデータ構造を使う一番の利点は、メモリを効率的に使用しながら、一定のサイズ内でデータを循環させることができる点です。**

固定長の配列をリング状に扱い、古いデータを自動で上書きしていく、キューの一種のデータ構造です。

## 基本的な動き

以下の挙動により、バッファが満杯の状態で新しい要素が追加されると、最も古い要素が上書きされる。

- リングバッファは、head(読み取り位置)と tail(書き込み位置)という 2 つのポインタ(index)によって管理される
- 新しい要素が追加されると tail が進み、要素が読み取られると head が進む
- ポインタが配列の最後に来ると、次は配列の先頭(index=0)に戻る

## 特徴

- **定数時間の操作**: head と tail の index を更新するだけで enqueue と dequeue が完了するため、どちらも O(1) の計算量になり、要素のシフトやメモリコピーは発生しない
- **固定サイズ**: 最初に固定サイズで確保されるため、実行中の動的なメモリ確保・解放が発生せず、メモリ使用量が安定しており、リソースに制約のある環境やリアルタイムシステムでは大きな利点となる

## 注意点

- **動的サイズ変更が不可**: サイズを変更する場合は新しいバッファを作成する必要がある

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

## スクラッチでの実装

スライスと 2 つの index を用いて実装できる。
バッファが満杯の際は一番古いデータを上書きする実装例を示す。
（上書きせず、満杯になったらエラーを返す実装もある）

```go
package main

import (
    "fmt"
)

// int型のリングバッファ
type RingBuffer struct {
    bufferint
    size int
    head int // 読み取り位置
    tail int // 書き込み位置
    count int // 現在の要素数
}

// 新しいリングバッファを作成
func NewRingBuffer(size int) *RingBuffer {
    return &RingBuffer{
        buffer: make(int, size),
        size: size,
    }
}

// 要素を追加する
// バッファが満杯なら古い要素を上書きする
func (rb *RingBuffer) Enqueue(item int) {
    rb.buffer[rb.tail] = item
    rb.tail = (rb.tail + 1) % rb.size

    if rb.count < rb.size {
        rb.count ++
    } else {
        // 上書きした場合、head も追従して進める
        rb.head = (rb.head + 1) % rb.size
    }
}

// 要素を取り出す
func (rb *RingBuffer) Dequeue() (int, error) {
    if rb.count = 0 {
        return 0, fmt.Errorf("ring buffer is empty")
    }

    item := rb.buffer[rb.head]
    rb.head = (rb.head + 1) % rb.size
    rb.count-

    return item, nil
}

func (rb *RingBuffer) IsEmpty() bool {
    return rb.count == 0
}

```

## サードパーティーライブラリ

以下のような高機能なリングバッファが必要な場合、github.com/smallnest/ringbuffer のようなライブラリを利用するのが良い。

- ブロッキング/ノンブロッキング動作の切り替え
- io.Reader/io.Writer インターフェースの提供

## 実装選択の指針（要確認）

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

# スライスベースとリングバッファの性能比較

キューとして機能させる（先頭から削除し、末尾に追加する）操作を多数繰り返すベンチマークでは、性能の差が顕著になる。ある報告では、10 万要素を持つデータ構造に対してこの操作を行うと、リングバッファがスライスベースより 2500 倍以上高速らしい（参照: 56）

# リングバッファの応用例

- **I/O バッファリング**: ネットワーク通信やファイル I/O など、データの生成速度と消費速度が異なるプロデューサ・コンシューマ問題において緩衝材として機能する。
- **ストリーミングデータ処理**: 音声や映像のように、常に最新のデータチャンクを保持し、古いものから順に処理していく場合に適している。
- **ログやイベント履歴**: システムの直近 N 件のログやイベント履歴を保持するのに使われる。新しいログが来ると最も古いログが自動的に破棄されるため、履歴のサイズを一定に保つことができる
