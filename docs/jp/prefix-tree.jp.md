# Prefix-tree（トライ木）とは

Prefix-tree（プレフィックス木）は、**トライ木（Trie）**とも呼ばれ、文字列の集合を効率的に格納・検索するための木構造である。各ノードが文字を表し、ルートから葉までのパスが一つの文字列を形成する。共通の接頭辞（prefix）を持つ文字列は同じパスを共有するため、メモリ効率と検索効率を両立できる。

**文字列の前方一致検索、オートコンプリート、辞書検索を高速に実行できる。**

# Prefix-tree の構成要素

## 基本要素

- **ルートノード（Root Node）**: 空文字列を表す最上位ノード
- **ノード（Node）**: 各文字に対応する単位。文字情報と子ノードへのリンクを保持
- **エッジ（Edge）**: 親ノードから子ノードへの文字遷移を表現
- **終端フラグ（End Flag）**: そのノードで文字列が終了することを示すフラグ
- **文字キー（Character Key）**: ノード間の遷移に使用する文字
- **子ノード（Children）**: 各文字に対応する子ノードの集合
- **パス（Path）**: ルートから特定ノードまでの文字列
- **接頭辞（Prefix）**: 共通の文字列開始部分

## 文字列の表現方法

- **パス表現**: ルートからノードまでのパスが文字列を表現
- **終端マーク**: 単語の終了を明示的に記録
- **共有構造**: 共通接頭辞を持つ文字列はパスを共有
- **階層化**: 文字列の階層的分解による効率的格納

# Prefix-tree の種類

## 標準 Prefix-tree（基本トライ木）

- 最も基本的な Prefix-tree
- 各文字が一つのノードに対応
- 実装がシンプル

## 圧縮 Prefix-tree（Radix Tree / Patricia Tree）

- 一本道のパスを圧縮して格納
- メモリ効率を大幅に改善
- 用途: 大量の文字列データの効率的格納

## Suffix Trie

- 文字列のすべての接尾辞を格納
- 部分文字列検索に特化
- 用途: 文字列解析、パターンマッチング

## AC Trie（Aho-Corasick）

- 複数パターンの同時検索に特化
- 失敗リンクで効率化
- 用途: テキスト内の複数パターン検索

## Double Array Trie

- 配列ベースの高速実装
- メモリ効率とアクセス速度を両立
- 用途: 大規模辞書システム

# Prefix-tree の特徴

- **前方一致検索**: O(m) 時間（m: 検索文字列長）
- **接頭辞共有**: 共通接頭辞を持つ文字列でメモリ効率化
- **辞書順列挙**: 辞書順でのすべての文字列取得が容易
- **動的更新**: 文字列の動的な追加・削除が効率的
- **オートコンプリート**: 入力途中での候補提示が高速

# Prefix-tree の構成要素の図解

## 文字列集合 {"cat", "cats", "car", "card", "care", "careful"} の Prefix-tree 例

```
文字列セット:
- cat
- cats
- car
- card
- care
- careful

Prefix-tree構造:
                    Root
                     |
                     c
                     |
                     a
                   /   \
                  t     r
                  |   /   \
                 (cat)   d   e
                  |      |   |
                  s    (card)(care)
                  |          |
                (cats)       f
                             |
                             u
                             |
                             l
                             |
                          (careful)

() = 終端ノード（単語の終了）
```

### 簡略化した詳細構造

```
       Root
        |
        c
        |
        a ──────────────┐
        |               |
        t               r
        |       ┌───────┼───────┐
     [cat*]     |       |       |
        |       d       e       ?
        s       |       |
     [cats*]  [card*] [care*]
                        |
                        f
                        |
                        u
                        |
                        l
                        |
                   [careful*]

* = isEndOfWord = true
```

## 各要素の詳細説明

### ノード（Node）構造

```go
type TrieNode struct {
    Children     map[rune]*TrieNode  // 子ノードのマップ
    IsEndOfWord  bool                // 終端フラグ
    Value        interface{}         // オプション: 関連する値
    Count        int                 // オプション: このパスの出現回数
}
```

### パスと文字列の対応

- **Root → c → a → t**: "cat" （終端あり）
- **Root → c → a → t → s**: "cats" （終端あり）
- **Root → c → a → r**: "car" （終端あり）
- **Root → c → a → r → d**: "card" （終端あり）
- **Root → c → a → r → e**: "care" （終端あり）
- **Root → c → a → r → e → f → u → l**: "careful" （終端あり）

### 共有構造の効果

- **"c" ノード**: すべての文字列で共有
- **"ca" パス**: すべての文字列で共有
- **"car" パス**: "car", "card", "care", "careful" で共有
- **"care" パス**: "care", "careful" で共有

## より複雑な例：英単語辞書

```
辞書: {"a", "an", "and", "ant", "any", "app", "apple", "apply"}

Prefix-tree:
                Root
                 |
                 a ──── [a*]
                 |
                 n ──── [an*]
              ┌──┼──┐
              d  t  y
              |  |  |
           [and*] [ant*] [any*]

             p ──── [ap] (NOT終端)
             |
             p ──── [app*]
             |
             l ──── [appl] (NOT終端)
          ┌──┼──┐
          e  i  y
          |  |  |
      [apple*] ? [apply*]

* = 単語終端
```

## 圧縮 Prefix-tree（Radix Tree）の例

```
標準Prefix-tree → 圧縮Prefix-tree

標準版:
  Root-c-a-r-e-f-u-l
       |   |
       t   d
       |   |
      (cat) (card)
       |
       s
       |
     (cats)

圧縮版:
    Root
     |
   "ca"
  ┌──┼──┐
"t"  "rd" "re"──┐
 |    |    |    |
(cat) (card)(care)"ful"
 |              |
"s"         (careful)
 |
(cats)
```

### 圧縮の効果

- **メモリ削減**: 一本道のノード列を一つのノードに圧縮
- **アクセス効率**: 長い文字列でもノード数を削減
- **実装複雑化**: 可変長文字列の管理が必要

# Prefix-tree で行える処理

## 基本操作

| 機能        | 説明                 | 計算量 | 戻り値              | 得意/苦手 | 補足                      |
| ----------- | -------------------- | ------ | ------------------- | --------- | ------------------------- |
| Insert      | 文字列を挿入         | O(m)   | bool（成功/失敗）   | ✅        | m: 文字列長、非常に効率的 |
| Search      | 文字列を検索         | O(m)   | bool（存在/非存在） | ✅        | 前方一致で高速検索        |
| Delete      | 文字列を削除         | O(m)   | bool（成功/失敗）   | ⚠️        | 共有ノードの処理が複雑    |
| Contains    | 文字列の存在確認     | O(m)   | bool（存在/非存在） | ✅        | Search と同等             |
| IsEmpty     | 木が空かチェック     | O(1)   | bool（空/非空）     | ✅        | ルートの子ノード確認      |
| Size        | 格納文字列数を取得   | O(n)   | int（文字列数）     | ❌        | 全ノードを走査する必要    |
| Clear       | 全文字列を削除       | O(n)   | なし（void）        | ⚠️        | メモリ解放が必要          |
| GetAllWords | すべての文字列を取得 | O(n×m) | 文字列配列          | ✅        | DFS で効率的に取得        |

## 前方一致・接頭辞操作

| 機能                 | 説明                         | 計算量     | 戻り値              | 得意/苦手 | 補足                       |
| -------------------- | ---------------------------- | ---------- | ------------------- | --------- | -------------------------- |
| StartsWith           | 接頭辞で始まる文字列があるか | O(m)       | bool（存在/非存在） | ✅        | Prefix-tree の最大の利点   |
| GetWordsWithPrefix   | 指定接頭辞の全文字列         | O(m + k×l) | 文字列配列          | ✅        | k: 結果数、l: 平均文字列長 |
| AutoComplete         | オートコンプリート候補       | O(m + k×l) | 候補文字列配列      | ✅        | 入力補完に最適             |
| LongestPrefix        | 最長共通接頭辞を検索         | O(m)       | 最長接頭辞文字列    | ✅        | 共通部分を効率的に発見     |
| CountWordsWithPrefix | 接頭辞を持つ文字列数         | O(m + k)   | int（個数）         | ✅        | 候補数の高速カウント       |
| HasPrefix            | 接頭辞の存在確認             | O(m)       | bool（存在/非存在） | ✅        | 部分的なパス検索           |

## 文字列解析・パターン検索

| 機能              | 説明                       | 計算量   | 戻り値           | 得意/苦手 | 補足                       |
| ----------------- | -------------------------- | -------- | ---------------- | --------- | -------------------------- |
| GetLongestWord    | 最長文字列を取得           | O(n×m)   | 最長文字列       | ⚠️        | 全文字列の走査が必要       |
| GetShortestWord   | 最短文字列を取得           | O(n×m)   | 最短文字列       | ⚠️        | 全文字列の走査が必要       |
| FindPattern       | パターンマッチング         | O(n×m)   | マッチ文字列配列 | ❌        | ワイルドカード検索は非効率 |
| GetWordsByLength  | 指定長の文字列を取得       | O(n×m)   | 指定長文字列配列 | ❌        | 長さ制限での検索は非効率   |
| ContainsSubstring | 部分文字列を含む文字列検索 | O(n×m²)  | 含有文字列配列   | ❌        | 部分文字列検索は不得意     |
| RegexMatch        | 正規表現マッチング         | O(n×m×r) | マッチ文字列配列 | ❌        | 正規表現は複雑             |

## 辞書・検索機能

| 機能            | 説明                   | 計算量     | 戻り値           | 得意/苦手 | 補足                 |
| --------------- | ---------------------- | ---------- | ---------------- | --------- | -------------------- |
| GetSortedWords  | 辞書順ソート済み文字列 | O(n×m)     | ソート済み配列   | ✅        | DFS で自然にソート順 |
| GetWordsInRange | 辞書順範囲内文字列     | O(m + k×l) | 範囲内文字列配列 | ✅        | 範囲検索に効率的     |
| GetNextWord     | 辞書順で次の文字列     | O(m + l)   | 次の文字列       | ✅        | 辞書的順序での移動   |
| GetPrevWord     | 辞書順で前の文字列     | O(m + l)   | 前の文字列       | ⚠️        | 実装が複雑           |
| SpellCheck      | スペルチェック         | O(m×d)     | 候補修正文字列   | ⚠️        | 編集距離計算が必要   |
| FuzzySearch     | あいまい検索           | O(m×d×n)   | 類似文字列配列   | ❌        | 計算量が大きい       |

## 構造解析・統計

| 機能                | 説明               | 計算量 | 戻り値            | 得意/苦手 | 補足                     |
| ------------------- | ------------------ | ------ | ----------------- | --------- | ------------------------ |
| GetHeight           | 木の高さを取得     | O(n)   | int（高さ）       | ✅        | 最長文字列の長さ         |
| GetNodeCount        | ノード数を取得     | O(n)   | int（ノード数）   | ✅        | DFS で全ノードをカウント |
| GetLeafCount        | 葉ノード数を取得   | O(n)   | int（葉ノード数） | ✅        | 終端ノードの数           |
| GetMemoryUsage      | メモリ使用量を取得 | O(n)   | int（バイト数）   | ✅        | ノード毎のサイズ計算     |
| GetCompressionRatio | 圧縮率を計算       | O(n×m) | float（圧縮率）   | ✅        | 元文字列とのサイズ比較   |
| Validate            | 木構造の妥当性確認 | O(n×m) | bool（妥当/不正） | ⚠️        | デバッグ・テスト用       |

## 苦手な処理・制限

| 機能          | 説明                         | 計算量  | 戻り値         | 得意/苦手 | 補足                     |
| ------------- | ---------------------------- | ------- | -------------- | --------- | ------------------------ |
| SuffixSearch  | 後方一致検索                 | O(n×m)  | 後方一致文字列 | ❌        | Suffix-tree が適している |
| MiddleMatch   | 中間一致検索                 | O(n×m²) | 中間一致文字列 | ❌        | 全文字列の部分文字列検索 |
| RandomAccess  | インデックスでの直接アクセス | O(n×m)  | i 番目の文字列 | ❌        | 順序アクセスが必要       |
| ReverseSearch | 逆順検索                     | O(n×m)  | 逆順文字列     | ❌        | 逆向きトライが必要       |
| Concatenation | 文字列の連結                 | O(n×m)  | 連結済みトライ | ❌        | 新しいトライの再構築     |
| SetOperations | 集合演算（和・差・積）       | O(n×m)  | 演算結果トライ | ❌        | 複雑なマージ処理         |

# Prefix-tree の実装方法

## 一般的な実装方法

### マップベース実装

- **子ノード管理**: `map[rune]*TrieNode` で文字から子ノードへのマッピング
- **動的サイズ**: 必要な文字のみノードを作成
- **メモリ効率**: 使用しない文字のメモリを節約

### 配列ベース実装

- **固定サイズ配列**: ASCII 文字セット（256 要素）や小文字英字（26 要素）
- **高速アクセス**: O(1)での子ノードアクセス
- **メモリトレードオフ**: 未使用要素でもメモリを消費

### ハイブリッド実装

- **小文字英字**: 配列で O(1)アクセス
- **その他文字**: マップで柔軟に対応
- **最適化**: よく使う文字セットを高速化

## Go による基本的な Prefix-tree 実装例

```go
package prefixtree

import (
    "strings"
)

// トライ木のノード構造体
type TrieNode struct {
    Children    map[rune]*TrieNode
    IsEndOfWord bool
    Value       interface{} // オプション: 関連する値を格納
    Count       int         // オプション: この単語の出現回数
}

// トライ木構造体
type PrefixTree struct {
    Root *TrieNode
    Size int // 格納されている単語数
}

// 新しいトライ木を作成
func NewPrefixTree() *PrefixTree {
    return &PrefixTree{
        Root: &TrieNode{
            Children:    make(map[rune]*TrieNode),
            IsEndOfWord: false,
        },
        Size: 0,
    }
}

// 新しいノードを作成
func newTrieNode() *TrieNode {
    return &TrieNode{
        Children:    make(map[rune]*TrieNode),
        IsEndOfWord: false,
    }
}

// 文字列を挿入
func (pt *PrefixTree) Insert(word string) {
    if pt.insertHelper(pt.Root, []rune(word), 0) {
        pt.Size++
    }
}

func (pt *PrefixTree) insertHelper(node *TrieNode, runes []rune, index int) bool {
    // 文字列の終端に到達
    if index == len(runes) {
        if !node.IsEndOfWord {
            node.IsEndOfWord = true
            node.Count = 1
            return true // 新しい単語を追加
        } else {
            node.Count++
            return false // 既存の単語
        }
    }

    char := runes[index]

    // 子ノードが存在しない場合は作成
    if child, exists := node.Children[char]; !exists {
        node.Children[char] = newTrieNode()
        return pt.insertHelper(node.Children[char], runes, index+1)
    } else {
        return pt.insertHelper(child, runes, index+1)
    }
}

// 文字列を検索
func (pt *PrefixTree) Search(word string) bool {
    node := pt.searchNode(pt.Root, []rune(word), 0)
    return node != nil && node.IsEndOfWord
}

func (pt *PrefixTree) searchNode(node *TrieNode, runes []rune, index int) *TrieNode {
    // 文字列の終端に到達
    if index == len(runes) {
        return node
    }

    char := runes[index]
    if child, exists := node.Children[char]; exists {
        return pt.searchNode(child, runes, index+1)
    }

    return nil
}

// 接頭辞で始まる文字列があるかチェック
func (pt *PrefixTree) StartsWith(prefix string) bool {
    node := pt.searchNode(pt.Root, []rune(prefix), 0)
    return node != nil
}

// 指定接頭辞を持つすべての文字列を取得
func (pt *PrefixTree) GetWordsWithPrefix(prefix string) []string {
    var result []string
    prefixNode := pt.searchNode(pt.Root, []rune(prefix), 0)

    if prefixNode != nil {
        pt.collectWords(prefixNode, prefix, &result)
    }

    return result
}

func (pt *PrefixTree) collectWords(node *TrieNode, currentWord string, result *[]string) {
    if node.IsEndOfWord {
        *result = append(*result, currentWord)
    }

    for char, child := range node.Children {
        pt.collectWords(child, currentWord+string(char), result)
    }
}

// オートコンプリート候補を取得（上位N件）
func (pt *PrefixTree) AutoComplete(prefix string, limit int) []string {
    var result []string
    prefixNode := pt.searchNode(pt.Root, []rune(prefix), 0)

    if prefixNode != nil {
        pt.collectWordsWithLimit(prefixNode, prefix, &result, limit)
    }

    return result
}

func (pt *PrefixTree) collectWordsWithLimit(node *TrieNode, currentWord string, result *[]string, limit int) {
    if len(*result) >= limit {
        return
    }

    if node.IsEndOfWord {
        *result = append(*result, currentWord)
    }

    // 辞書順で子ノードを処理
    for char := rune(0); char <= rune(0x10FFFF); char++ {
        if child, exists := node.Children[char]; exists {
            pt.collectWordsWithLimit(child, currentWord+string(char), result, limit)
            if len(*result) >= limit {
                break
            }
        }
    }
}

// 文字列を削除
func (pt *PrefixTree) Delete(word string) bool {
    if pt.deleteHelper(pt.Root, []rune(word), 0) {
        pt.Size--
        return true
    }
    return false
}

func (pt *PrefixTree) deleteHelper(node *TrieNode, runes []rune, index int) bool {
    // 文字列の終端に到達
    if index == len(runes) {
        if !node.IsEndOfWord {
            return false // 単語が存在しない
        }

        node.IsEndOfWord = false
        node.Count = 0

        // 子ノードがない場合、このノードは削除可能
        return len(node.Children) == 0
    }

    char := runes[index]
    child, exists := node.Children[char]
    if !exists {
        return false // 単語が存在しない
    }

    shouldDeleteChild := pt.deleteHelper(child, runes, index+1)

    if shouldDeleteChild {
        delete(node.Children, char)
        // 現在のノードが終端でなく、子ノードもない場合は削除可能
        return !node.IsEndOfWord && len(node.Children) == 0
    }

    return false
}

// すべての文字列を辞書順で取得
func (pt *PrefixTree) GetAllWords() []string {
    var result []string
    pt.collectWords(pt.Root, "", &result)
    return result
}

// 最長共通接頭辞を取得
func (pt *PrefixTree) LongestCommonPrefix() string {
    if pt.Size == 0 {
        return ""
    }

    var prefix strings.Builder
    current := pt.Root

    for {
        // 終端ノードの場合、または子が複数ある場合は終了
        if current.IsEndOfWord || len(current.Children) != 1 {
            break
        }

        // 子が一つの場合、その文字を追加
        for char, child := range current.Children {
            prefix.WriteRune(char)
            current = child
            break
        }
    }

    return prefix.String()
}

// 指定文字列の出現回数を取得
func (pt *PrefixTree) GetCount(word string) int {
    node := pt.searchNode(pt.Root, []rune(word), 0)
    if node != nil && node.IsEndOfWord {
        return node.Count
    }
    return 0
}

// 木の高さを取得
func (pt *PrefixTree) GetHeight() int {
    return pt.getHeightHelper(pt.Root)
}

func (pt *PrefixTree) getHeightHelper(node *TrieNode) int {
    if len(node.Children) == 0 {
        return 0
    }

    maxHeight := 0
    for _, child := range node.Children {
        height := pt.getHeightHelper(child)
        if height > maxHeight {
            maxHeight = height
        }
    }

    return maxHeight + 1
}

// ノード数を取得
func (pt *PrefixTree) GetNodeCount() int {
    return pt.getNodeCountHelper(pt.Root)
}

func (pt *PrefixTree) getNodeCountHelper(node *TrieNode) int {
    count := 1 // 現在のノード
    for _, child := range node.Children {
        count += pt.getNodeCountHelper(child)
    }
    return count
}

// 辞書順で指定範囲の文字列を取得
func (pt *PrefixTree) GetWordsInRange(start, end string) []string {
    var result []string
    pt.collectWordsInRange(pt.Root, "", start, end, &result)
    return result
}

func (pt *PrefixTree) collectWordsInRange(node *TrieNode, currentWord, start, end string, result *[]string) {
    if node.IsEndOfWord && currentWord >= start && currentWord <= end {
        *result = append(*result, currentWord)
    }

    for char, child := range node.Children {
        newWord := currentWord + string(char)
        // 範囲外の場合はスキップ
        if newWord > end {
            continue
        }
        // 可能性がある場合は継続
        if newWord < start && !strings.HasPrefix(start, newWord) {
            continue
        }
        pt.collectWordsInRange(child, newWord, start, end, result)
    }
}

// メモリ使用量を概算で取得（バイト）
func (pt *PrefixTree) GetMemoryUsage() int {
    return pt.getMemoryUsageHelper(pt.Root)
}

func (pt *PrefixTree) getMemoryUsageHelper(node *TrieNode) int {
    size := 0
    // ノード自体のサイズ（概算）
    size += 32 // 基本構造体
    size += len(node.Children) * 16 // マップエントリ

    for _, child := range node.Children {
        size += pt.getMemoryUsageHelper(child)
    }

    return size
}
```

## 高度な実装（圧縮 Prefix-tree / Radix Tree）

```go
// 圧縮トライ木のノード
type RadixNode struct {
    Key         string              // 圧縮された文字列
    IsEndOfWord bool
    Value       interface{}
    Children    map[rune]*RadixNode
}

// 圧縮トライ木
type RadixTree struct {
    Root *RadixNode
    Size int
}

// 新しい圧縮トライ木を作成
func NewRadixTree() *RadixTree {
    return &RadixTree{
        Root: &RadixNode{
            Key:      "",
            Children: make(map[rune]*RadixNode),
        },
        Size: 0,
    }
}

// 文字列を挿入（圧縮版）
func (rt *RadixTree) Insert(word string) {
    rt.insertHelper(rt.Root, word)
}

func (rt *RadixTree) insertHelper(node *RadixNode, word string) {
    if len(word) == 0 {
        if !node.IsEndOfWord {
            node.IsEndOfWord = true
            rt.Size++
        }
        return
    }

    firstChar := rune(word[0])

    // 対応する子ノードを検索
    if child, exists := node.Children[firstChar]; exists {
        // 共通プレフィックスを見つける
        commonLen := findCommonPrefixLength(child.Key, word)

        if commonLen == len(child.Key) {
            // 子ノードのキーが完全に一致
            rt.insertHelper(child, word[commonLen:])
        } else if commonLen == len(word) {
            // 挿入する単語が子ノードのキーのプレフィックス
            // 子ノードを分割
            rt.splitNode(child, commonLen)
            child.IsEndOfWord = true
            rt.Size++
        } else {
            // 部分的な一致 - ノードを分割
            rt.splitNode(child, commonLen)
            rt.insertHelper(child, word[commonLen:])
        }
    } else {
        // 新しい子ノードを作成
        newNode := &RadixNode{
            Key:         word,
            IsEndOfWord: true,
            Children:    make(map[rune]*RadixNode),
        }
        node.Children[firstChar] = newNode
        rt.Size++
    }
}

// ノードを分割
func (rt *RadixTree) splitNode(node *RadixNode, splitPos int) {
    // 分割後の子ノード
    newChild := &RadixNode{
        Key:         node.Key[splitPos:],
        IsEndOfWord: node.IsEndOfWord,
        Value:       node.Value,
        Children:    node.Children,
    }

    // 現在のノードを更新
    node.Key = node.Key[:splitPos]
    node.IsEndOfWord = false
    node.Value = nil
    node.Children = make(map[rune]*RadixNode)

    // 新しい子ノードを追加
    if len(newChild.Key) > 0 {
        node.Children[rune(newChild.Key[0])] = newChild
    }
}

// 共通プレフィックスの長さを計算
func findCommonPrefixLength(str1, str2 string) int {
    minLen := len(str1)
    if len(str2) < minLen {
        minLen = len(str2)
    }

    for i := 0; i < minLen; i++ {
        if str1[i] != str2[i] {
            return i
        }
    }

    return minLen
}
```

## 配列ベース高速実装（英小文字専用）

```go
// 英小文字専用の高速トライ木
type FastTrie struct {
    Children    [26]*FastTrie // a-z の固定配列
    IsEndOfWord bool
    Count       int
}

// 新しい高速トライ木を作成
func NewFastTrie() *FastTrie {
    return &FastTrie{}
}

// 文字列を挿入（英小文字専用）
func (ft *FastTrie) Insert(word string) {
    current := ft

    for _, char := range word {
        if char < 'a' || char > 'z' {
            return // 英小文字以外は無視
        }

        index := char - 'a'
        if current.Children[index] == nil {
            current.Children[index] = &FastTrie{}
        }
        current = current.Children[index]
    }

    if !current.IsEndOfWord {
        current.IsEndOfWord = true
        current.Count = 1
    } else {
        current.Count++
    }
}

// 文字列を検索（英小文字専用）
func (ft *FastTrie) Search(word string) bool {
    current := ft

    for _, char := range word {
        if char < 'a' || char > 'z' {
            return false
        }

        index := char - 'a'
        if current.Children[index] == nil {
            return false
        }
        current = current.Children[index]
    }

    return current.IsEndOfWord
}

// 接頭辞で始まる文字列があるかチェック（英小文字専用）
func (ft *FastTrie) StartsWith(prefix string) bool {
    current := ft

    for _, char := range prefix {
        if char < 'a' || char > 'z' {
            return false
        }

        index := char - 'a'
        if current.Children[index] == nil {
            return false
        }
        current = current.Children[index]
    }

    return true
}
```

# Prefix-tree の応用例

## 1. 検索エンジン・オートコンプリート

- **検索候補**: ユーザー入力に基づく検索候補の高速提示
- **クエリ補完**: 部分入力からの完全なクエリ提案
- **人気度順**: 検索頻度に基づく候補順位付け
- **リアルタイム検索**: 入力と同時の候補更新

## 2. 辞書・言語処理システム

- **英語辞書**: 単語の存在確認と意味検索
- **スペルチェック**: 入力ミスの検出と修正候補
- **多言語辞書**: 複数言語の効率的な管理
- **語幹解析**: 単語の語幹と活用形の管理

## 3. プログラミング言語・IDE

- **コード補完**: 変数名・関数名の自動補完
- **シンタックスハイライト**: キーワードの高速認識
- **リファクタリング**: 識別子の一括置換
- **API ドキュメント**: メソッド名の検索と提案

## 4. ネットワーク・ルーティング

- **IP ルーティング**: IP アドレスの最長プレフィックスマッチ
- **URL マッチング**: Web ルーティングのパス解決
- **DNS 解決**: ドメイン名の階層的解決
- **ファイアウォール**: パケットフィルタリングルール

## 5. ゲーム・エンターテイメント

- **単語ゲーム**: Scrabble、クロスワードパズルの単語検証
- **チャットフィルタ**: 不適切な単語の検出とフィルタリング
- **プレイヤー名**: ユーザー名の重複チェックと提案
- **コマンド解析**: ゲーム内コマンドの解析

## 6. データベース・情報検索

- **フルテキスト検索**: 文書内の単語インデックス
- **カテゴリ分類**: 階層的なタグシステム
- **メタデータ検索**: ファイル属性の効率的検索
- **ログ解析**: ログパターンの高速マッチング

## 7. Web サービス・API

- **RESTful API**: エンドポイントのルーティング
- **パラメータ検証**: URL パラメータの妥当性チェック
- **アクセス制御**: 権限パスの階層的管理
- **キャッシュキー**: 階層的なキャッシュ戦略

## 8. セキュリティ・監視

- **マルウェア検出**: 既知の悪意あるファイル名パターン
- **侵入検知**: ネットワークトラフィックのパターンマッチング
- **ドメインブラックリスト**: 危険なドメインの高速チェック
- **ログ監視**: セキュリティイベントの検出

## 9. 電子商取引・推薦

- **商品検索**: 商品名・カテゴリの前方一致検索
- **カテゴリナビゲーション**: 階層的な商品分類
- **タグシステム**: 商品タグの効率的管理
- **ユーザー検索**: 顧客名の部分検索

## 10. 科学・研究

- **遺伝子配列**: DNA/RNA 配列の部分マッチング
- **化学式検索**: 分子構造の部分検索
- **文献検索**: 論文タイトル・キーワード検索
- **分類学**: 生物の階層的分類システム

## 11. 地理・位置情報

- **住所検索**: 住所の階層的検索（国 → 州 → 市 → 町）
- **郵便番号**: 郵便番号の前方一致検索
- **地名辞書**: 地名の効率的な格納と検索
- **ルート検索**: 道路名の部分検索

## 12. ファイル・ストレージシステム

- **ファイル名検索**: ファイル名の前方一致検索
- **パス解決**: ディレクトリパスの効率的解決
- **権限管理**: ファイルパス単位での権限制御
- **バックアップ**: ファイルパターンによる選択的バックアップ

# Prefix-tree と他のデータ構造との比較

## 文字列検索データ構造比較

| 特徴                   | Prefix-tree | ハッシュテーブル | 二分探索木    | Suffix-tree | 配列（ソート済み） | KMP       | Boyer-Moore |
| ---------------------- | ----------- | ---------------- | ------------- | ----------- | ------------------ | --------- | ----------- |
| **前方一致検索**       | ✅ O(m)     | ❌ 完全一致のみ  | ⚠️ O(m log n) | ⚠️ O(m)     | ⚠️ O(m log n)      | ❌ 不適   | ❌ 不適     |
| **完全一致検索**       | ✅ O(m)     | ✅ O(1)平均      | ✅ O(log n)   | ✅ O(m)     | ✅ O(log n)        | ✅ O(n+m) | ✅ O(n+m)   |
| **オートコンプリート** | ✅ O(m+k)   | ❌ 不適          | ❌ 困難       | ❌ 不適     | ❌ 困難            | ❌ 不適   | ❌ 不適     |
| **辞書順列挙**         | ✅ O(n)     | ❌ 不適          | ✅ O(n)       | ⚠️ 複雑     | ✅ O(n)            | ❌ 不適   | ❌ 不適     |
| **挿入・削除**         | ✅ O(m)     | ✅ O(1)平均      | ✅ O(log n)   | ❌ 複雑     | ❌ O(n)            | ❌ 不適   | ❌ 不適     |
| **メモリ効率**         | ⚠️ 中程度   | ✅ 高効率        | ✅ 高効率     | ❌ 大量     | ✅ 最高            | ✅ 最小   | ✅ 最小     |
| **動的更新**           | ✅ 効率的   | ✅ 効率的        | ✅ 効率的     | ❌ 困難     | ❌ 非効率          | ❌ 不適   | ❌ 不適     |
| **実装複雑度**         | ⚠️ 中程度   | ✅ 簡単          | ⚠️ 中程度     | ❌ 複雑     | ✅ 簡単            | ⚠️ 中程度 | ⚠️ 中程度   |
| **部分文字列検索**     | ❌ 前方のみ | ❌ 不適          | ❌ 不適       | ✅ 最適     | ❌ 線形探索        | ✅ 最適   | ✅ 最適     |

## 用途別最適選択

| 用途                     | 第 1 選択        | 第 2 選択        | 第 3 選択   | 避けるべき               |
| ------------------------ | ---------------- | ---------------- | ----------- | ------------------------ |
| **オートコンプリート**   | Prefix-tree      | -                | -           | ハッシュテーブル、配列   |
| **辞書・スペルチェック** | Prefix-tree      | ハッシュテーブル | 二分探索木  | 配列、KMP                |
| **IP ルーティング**      | Prefix-tree      | -                | -           | ハッシュテーブル         |
| **コード補完**           | Prefix-tree      | ハッシュテーブル | -           | 配列、Suffix-tree        |
| **URL/パスマッチング**   | Prefix-tree      | ハッシュテーブル | -           | 配列                     |
| **一般的な文字列検索**   | ハッシュテーブル | 二分探索木       | 配列        | Prefix-tree              |
| **部分文字列検索**       | Suffix-tree      | KMP              | Boyer-Moore | Prefix-tree              |
| **大量文字列の完全一致** | ハッシュテーブル | 二分探索木       | 配列        | Prefix-tree              |
| **リアルタイム検索**     | Prefix-tree      | ハッシュテーブル | -           | 配列、Suffix-tree        |
| **メモリ制約環境**       | 配列             | ハッシュテーブル | 二分探索木  | Prefix-tree、Suffix-tree |

## パフォーマンス特性

### Prefix-tree 実装方法別比較

| 実装方法             | 検索時間 | 挿入時間 | 削除時間 | メモリ使用量 | 適用場面       |
| -------------------- | -------- | -------- | -------- | ------------ | -------------- |
| **マップベース**     | O(m)     | O(m)     | O(m)     | 中程度       | 一般的用途     |
| **配列ベース**       | O(m)     | O(m)     | O(m)     | 高い         | 固定文字セット |
| **圧縮 Prefix-tree** | O(m)     | O(m)     | O(m)     | 低い         | メモリ制約     |
| **Double Array**     | O(m)     | ❌ 困難  | ❌ 困難  | 最低         | 読み取り専用   |
| **ハイブリッド**     | O(m)     | O(m)     | O(m)     | 中程度       | バランス重視   |

### データ特性別最適化

| データ特性                 | 推奨実装         | 理由                 |
| -------------------------- | ---------------- | -------------------- |
| **英語単語（小文字のみ）** | 配列ベース       | O(1)子ノードアクセス |
| **多言語テキスト**         | マップベース     | Unicode 対応         |
| **大量の短い文字列**       | 圧縮 Prefix-tree | メモリ効率           |
| **少数の長い文字列**       | 標準 Prefix-tree | 実装の簡潔性         |
| **読み取り専用辞書**       | Double Array     | 最小メモリ           |
| **頻繁な更新**             | マップベース     | 動的更新の効率性     |

## 制約条件別選択

| 制約条件                | 推奨選択                 | 理由                     |
| ----------------------- | ------------------------ | ------------------------ |
| **メモリ < 1MB**        | 圧縮 Prefix-tree         | 最小メモリフットプリント |
| **リアルタイム要求**    | 配列ベース Prefix-tree   | 最速アクセス             |
| **大量データ（GB 級）** | ハッシュテーブル         | スケーラビリティ         |
| **多言語対応**          | マップベース Prefix-tree | Unicode 柔軟性           |
| **組み込みシステム**    | 配列ベース               | 予測可能な性能           |
| **Web サービス**        | ハイブリッド実装         | バランスの良い性能       |
