# Prefix-tree（トライ木）とは

Prefix-tree（プレフィックス木）は、**トライ木（Trie）**とも呼ばれ、文字列の集合を効率的に格納・検索するための木構造である。各ノードが文字を表し、ルートから葉までのパスが一つの文字列を形成する。共通の接頭辞（prefix）を持つ文字列は同じパスを共有するため、メモリ効率と検索効率を両立できる。

**文字列の前方一致検索、オートコンプリート、辞書検索を高速に実行できる。**

## 特徴

以下のような直感的な構造を持つ。

- 各ノードは、アルファベットの 1 文字に対応する（根ノードを除く）
- 根ノードからあるノードまでのパスは、文字列の接頭辞（プレフィックス）を表現
- ノードには、そのノードで終わる完全な単語が存在するかどうかを示すフラグ（例えば、真偽値）が付与されることが多い

最も重要な特性は、共通の接頭辞を持つ文字列が木の同じパスを共有する点にある。例えば、「text」と「team」という 2 つの単語を格納する場合、「te」という接頭辞は木の共通の部分として一度だけ格納される。この性質により、接頭辞の重複が多いデータセット（例えば、英単語の辞書）では、大幅なメモリ削減が期待できる。

また、木構造の性質から保存した文字は自然に辞書順でデータが配置されるため、深さ優先探索を行うだけで辞書の操作が可能である。

## 計算量

文字列の長さを L とすると、検索や挿入は O(L) で完了する。これは、木の深さが最長文字列の長さに依存し、探索が文字列の長さに比例するためであり、Prefix-tree に保存されているキーの総数には依存しない。この特性が、大量の文字列の中から高速に検索を行いたい場合に Trie が強力な選択肢となる理由である。
ただし、ノードごとに子ノードのポインタを複数所持する必要があるため、文字の種類が多い場合、メモリ降雨率が悪くなる可能性がある。

## ユースケース

その接頭辞に基づいた検索性能から、検索エンジンのオートコンプリート機能やスペルチェッカー、IP ルーティングテーブル(最長接頭辞一致)などで広く利用されている。

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
- メモリ効率で課題あり
  - 分岐のない長い文字列を格納するさい、子ノードを一つしか持たないノードが連続した一本道の連結が形成されてしまう
  - ノードの増加とメモリ占有の原因となる

## 圧縮 Prefix-tree（Radix Tree / Patricia Tree）

- 一本道のパスを圧縮して格納
  - 子ノードが一つだけの中間ノードを親ノードと結合
  - これにより文字列でラベル付されたノードができる
- メモリ効率を大幅に改善
  - スパースなデータセット（共通の接頭辞が少ないデータ）や長い非分岐部分を持つ文字列を扱いやすい
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

# Prefix-tree の実装方法

## 一般的な実装方法

子ノードへのポインタをどうやって保持するかで方法が変わる。
それぞれの実装方法は、柔軟性とパフォーマンスの間でトレードオフが存在する。

### マップベース実装

柔軟性が高い。文字列のサイズを事前に知る必要がなく、Unicode の全文字をキーとして扱える。ユーザからの自由なテキスト入力を扱うような汎用的ケースに適している。

- **子ノード管理**: `map[rune]*TrieNode` で文字から子ノードへのマッピング
- **動的サイズ**: 必要な文字のみノードを作成
- **メモリ効率**: 使用しない文字のメモリを節約

懸念点として、ハッシュ計算時のオーバーヘッドが挙げられる。これにより配列ベースの実装に比べて、各要素へのアクセスが愛知則人ある可能性がある。メモリについても map 自体の内部構造によって消費量が大きくなる。

### 配列ベース実装

文字列の長さ N が既知で小さい場合、非常に高速かつメモリ効率よく計算できる。子ノードへのアクセスは、文字コードを index に変換するだけの単純計算なので、CPU キャッシュの観点からも有利である。

- **固定サイズ配列**: ASCII 文字セット（256 要素）や小文字英字（26 要素）
- **高速アクセス**: O(1)での子ノードアクセス
- **メモリトレードオフ**: 未使用要素でもメモリを消費

懸念点は、汎用性に欠けることである。扱う文字セットが限定され、Unicode のような巨大な文字セットに用いるのは現実的ではない。

### ハイブリッド実装

- **小文字英字**: 配列で O(1)アクセス
- **その他文字**: マップで柔軟に対応
- **最適化**: よく使う文字セットを高速化

## Go による基本的な Prefix-tree 実装例

children map[rune]\*Node と isEndOfWord bool フラグを持つ Node 構造体を用いた実装例である。

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

# 他のデータ構造との比較

文字列集合の格納・検索を高速にあつかうデータ構造として、トライ木とハッシュテーブルがよく比較される。

- 完全一致検索や存在確認などの基本的な用途が共通している
- 大量の文字列データを扱う辞書・検索・補完などで利用される
- 計算量やメモリ効率、前方一致検索の可否など、用途や用件によって選択が分かれる

| 特徴                   | Prefix Tree (Trie)                            | ハッシュテーブル               |
| ---------------------- | --------------------------------------------- | ------------------------------ |
| **点検索（完全一致）** | O(m)                                          | ✅ 平均 O(m)                   |
| **接頭辞検索**         | ✅ O(p) (p=接頭辞長)                          | ❌ 非効率（O(n・m)）           |
| **順序付き走査**       | ✅ 効率的（深さ優先探索を行えば可能）         | ❌ サポートしない              |
| **空間計算量**         | ⚠️ 可変（共通接頭辞が多いと効率的）           | ✅ O(n・m)                     |
| **主な用途**           | オートコンプリート、IP ルーティング、辞書検索 | 汎用的なキーバリュー格納、存在 |

Trie の最大の利点は、ハッシュテーブルが苦手とする操作、接頭辞検索と辞書順での走査を効率的に実行できる点にある。ハッシュテーブルで接頭辞検索を行うには、格納されている全てのキーを線形走査する必要があるため非効率である。一方、Trie では接頭辞の長さに比例した時間で効率的に完了する。この機能的な違いが、Trie を特定のアプリケーション(例: オートコンプリート)において不可欠なものにしている。

空間効率に関しては、一概にどちらが優れているとは言えない。共通接頭辞が多いデータセッ トでは Trie が有利だが、キーが互いに大きく異なる場合はハッシュテーブルの方がコンパクトになることがある。
