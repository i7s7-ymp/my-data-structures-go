# Suffix-tree（接尾辞木）とは

Suffix-tree（接尾辞木）は、文字列のすべての**接尾辞（suffix）**を効率的に管理するための特殊な木構造です。文字列の各位置から始まるすべての部分文字列を圧縮された形で格納し、文字列の様々な操作を高速に実行できます。

**このデータ構造の最大の利点は、文字列検索、最長共通部分文字列の発見、パターンマッチングを線形時間で実行できる点です。**

# Suffix-tree の構成要素

## 基本要素

- **接尾辞（Suffix）**: 文字列の任意の位置から末尾までの部分文字列
- **ノード（Node）**: 文字列の分岐点を表現する単位
- **エッジ（Edge）**: ノード間を結ぶ線。文字列の一部を表現
- **リーフノード（Leaf Node）**: 接尾辞の終端を表すノード
- **内部ノード（Internal Node）**: 複数の接尾辞が分岐する点
- **ルートノード（Root Node）**: 木構造の最上位ノード
- **接尾辞リンク（Suffix Link）**: 構築効率化のための補助リンク
- **終端記号**: 文字列の終端を示す特殊文字（通常 `$`）

## エッジラベルの圧縮

- **パスラベル**: ルートからノードまでのエッジラベルを連結した文字列
- **エッジ圧縮**: 一意パスのエッジを圧縮して格納
- **開始・終了インデックス**: エッジラベルを元文字列の位置で表現
- **グローバル終端**: すべてのリーフに共通の終端位置

# Suffix-tree の種類

## 基本的な Suffix-tree

- 最も基本的な接尾辞木
- すべての接尾辞を明示的に格納

## 圧縮 Suffix-tree

- エッジラベルを圧縮して格納
- メモリ効率を向上
- 用途: 大量の文字列データ処理

## 一般化 Suffix-tree（GST）

- 複数の文字列の接尾辞を同一の木に格納
- 文字列間の共通部分文字列を効率的に発見
- 用途: 比較ゲノミクス、プラジャリズム検出

## Suffix Array + LCP Array

- Suffix-tree の機能を配列で実現
- メモリ効率が高い
- 用途: 大規模テキスト処理

# Suffix-tree の特徴

- **線形空間**: O(n) のメモリ使用量（n: 文字列長）
- **線形時間構築**: Ukkonen's algorithm で O(n) 時間で構築
- **高速検索**: パターン検索が O(m) 時間（m: パターン長）
- **豊富な文字列操作**: 最長共通部分文字列、最長回文、文字列比較等
- **動的構築**: オンラインでの文字追加が可能

# Suffix-tree の構成要素の図解

## 文字列 "banana$" の Suffix-tree 例

```
文字列: "banana$"
接尾辞:
- banana$ (0)
- anana$  (1)
- nana$   (2)
- ana$    (3)
- na$     (4)
- a$      (5)
- $       (6)

Suffix-tree構造:
                  Root
                /  |  \
            [1,1] /   |   \ [$,6]
               a/    |    \$
              /   [2,2]    \
           Node1    n       Leaf6
          /  |  \           ($)
      [2,6]/   |   \[5,6]
        na/ [3,3]  \a$
        /     a     \
    Leaf1          Leaf5
   (banana$)       (a$)
       |
   [4,6]/na$
      /
   Leaf2
  (anana$)
       |
   [4,6]/na$
      /
   Leaf3
   (nana$)
       |
   [4,6]/na$
      /
   Leaf4
   (ana$)
```

### 簡略化した構造

```
       Root
      /  |  \
   (a)   (n) ($)
    |     |    |
  Node   Node  Leaf6
 /  |     |
(na$) (a$) (ana$)
  |    |     |
Leaf1 Leaf5 Node
              |
           (na$)
              |
            Leaf2,3,4
```

## 各要素の説明

### ノード（Node）

- **Root**: すべての接尾辞の共通開始点
- **内部ノード**: 複数の接尾辞が分岐する点
- **リーフノード**: 各接尾辞の終端（接尾辞のインデックスを保持）

### エッジ（Edge）

```go
type Edge struct {
    StartIndex int    // 元文字列での開始位置
    EndIndex   *int   // 元文字列での終了位置（グローバル終端の場合はポインタ）
    Label      string // エッジのラベル（デバッグ用）
}
```

### パスラベル

- **Root → Node1**: "a"
- **Root → Node1 → Leaf1**: "a" + "nana$" = "anana$"
- **Root → Leaf6**: "$"

### 接尾辞リンク（Suffix Link）

```
内部ノードからその接尾辞に対応するノードへのリンク
例: "ana"のノード → "na"のノード
```

# Suffix-tree で行える処理

## 基本操作

| 機能            | 説明                          | 計算量   | 戻り値              | 得意/苦手 | 補足                           |
| --------------- | ----------------------------- | -------- | ------------------- | --------- | ------------------------------ |
| Build           | 文字列から Suffix-tree を構築 | O(n)     | Suffix-tree         | ✅        | Ukkonen's algorithm で線形時間 |
| Search          | パターン文字列を検索          | O(m)     | bool（存在/非存在） | ✅        | m: パターン長、高速検索        |
| FindOccurrences | パターンのすべての出現位置    | O(m + k) | 位置の配列          | ✅        | k: 出現回数                    |
| Insert          | 新しい文字を動的追加          | O(1)償却 | 更新済み木          | ✅        | オンライン構築                 |
| GetSuffixes     | すべての接尾辞を取得          | O(n²)    | 接尾辞配列          | ❌        | 出力サイズが大きい             |

## 文字列解析操作

| 機能                     | 説明                 | 計算量   | 戻り値           | 得意/苦手 | 補足                 |
| ------------------------ | -------------------- | -------- | ---------------- | --------- | -------------------- |
| LongestCommonSubstring   | 最長共通部分文字列   | O(n + m) | 共通部分文字列   | ✅        | 複数文字列間の比較   |
| LongestRepeatedSubstring | 最長反復部分文字列   | O(n)     | 反復部分文字列   | ✅        | 内部ノードの最深パス |
| CountDistinctSubstrings  | 異なる部分文字列の数 | O(n)     | int（個数）      | ✅        | 木の構造から直接計算 |
| FindLongestPalindrome    | 最長回文を検索       | O(n)     | 回文文字列       | ⚠️        | 拡張アルゴリズム必要 |
| SuffixArray              | 接尾辞配列を生成     | O(n)     | ソート済み接尾辞 | ✅        | DFS で線形時間       |
| LCP                      | 最長共通接頭辞配列   | O(n)     | LCP 配列         | ✅        | 隣接接尾辞の共通長   |

## 高度な文字列操作

| 機能              | 説明                 | 計算量      | 戻り値          | 得意/苦手 | 補足                   |
| ----------------- | -------------------- | ----------- | --------------- | --------- | ---------------------- |
| SubstringCount    | 部分文字列の出現回数 | O(m)        | int（出現回数） | ✅        | リーフ数をカウント     |
| UniqueSubstrings  | 一意部分文字列の検索 | O(n)        | 一意文字列配列  | ✅        | リーフが 1 つのパス    |
| PrefixSuffixMatch | 接頭辞と接尾辞の一致 | O(n)        | 一致する文字列  | ✅        | 特殊な走査アルゴリズム |
| StringComparison  | 文字列の辞書順比較   | O(min(n,m)) | 比較結果        | ✅        | Suffix-tree 上での比較 |
| EditDistance      | 編集距離の計算       | O(n×m)      | int（編集距離） | ❌        | 動的計画法が必要       |
| ApproximateMatch  | 近似文字列検索       | O(n×m)      | 近似一致位置    | ❌        | 複雑なアルゴリズム     |

## 構造解析・最適化

| 機能        | 説明               | 計算量 | 戻り値            | 得意/苦手 | 補足               |
| ----------- | ------------------ | ------ | ----------------- | --------- | ------------------ |
| GetHeight   | 木の高さを取得     | O(n)   | int（高さ）       | ✅        | 最長パスの長さ     |
| NodeCount   | ノード数を取得     | O(1)   | int（ノード数）   | ✅        | 構築時にカウント   |
| EdgeCount   | エッジ数を取得     | O(1)   | int（エッジ数）   | ✅        | ノード数-1         |
| MemoryUsage | メモリ使用量を取得 | O(1)   | int（バイト数）   | ✅        | 線形空間の確認     |
| Validate    | 木構造の妥当性確認 | O(n)   | bool（妥当/不正） | ⚠️        | デバッグ・テスト用 |
| Optimize    | 木構造の最適化     | O(n)   | 最適化済み木      | ⚠️        | エッジ圧縮の最適化 |

## 苦手な処理・制限

| 機能          | 説明                         | 計算量     | 戻り値     | 得意/苦手 | 補足                   |
| ------------- | ---------------------------- | ---------- | ---------- | --------- | ---------------------- |
| RandomAccess  | インデックスでの直接アクセス | O(n)       | 文字       | ❌        | 木構造の特性上非効率   |
| Concatenation | 文字列の連結                 | O(n + m)   | 新しい木   | ❌        | 再構築が必要           |
| Deletion      | 文字の削除                   | O(n)       | 更新済み木 | ❌        | 木の再構築が必要       |
| Modification  | 文字の変更                   | O(n)       | 更新済み木 | ❌        | 部分的な再構築         |
| Compression   | 木の圧縮                     | O(n log n) | 圧縮済み木 | ❌        | 複雑な圧縮アルゴリズム |
| Parallel      | 並列処理                     | -          | -          | ❌        | 木構造の依存関係       |

# Suffix-tree の実装方法

## 一般的な実装方法

### Ukkonen's Algorithm（最も効率的）

- **線形時間構築**: O(n) 時間で Suffix-tree を構築
- **オンライン構築**: 文字を一つずつ追加して構築
- **複雑度**: 実装が複雑だが最も効率的

### Naive Algorithm（単純法）

- **理解しやすい**: 各接尾辞を順次挿入
- **時間計算量**: O(n²) だが実装が簡単
- **教育用**: アルゴリズムの理解に適している

### McCreight's Algorithm

- **効率的**: O(n) 時間構築
- **バッチ処理**: すべての接尾辞を一度に処理
- **実装の複雑度**: 中程度

## Go による基本的な Suffix-tree 実装例

```go
package suffixtree

import (
    "fmt"
    "strings"
)

// エッジの終端位置を表現（グローバル終端対応）
type EdgeEnd struct {
    IsGlobal bool
    Position int
}

// ノード構造体
type SuffixTreeNode struct {
    Children    map[byte]*SuffixTreeNode
    SuffixLink  *SuffixTreeNode
    Start       int
    End         *EdgeEnd
    SuffixIndex int // リーフノードの場合の接尾辞インデックス
}

// Suffix-tree構造体
type SuffixTree struct {
    Root        *SuffixTreeNode
    Text        string
    GlobalEnd   *EdgeEnd  // グローバル終端位置
    ActiveNode  *SuffixTreeNode
    ActiveEdge  int
    ActiveLength int
    Remaining   int
    LeafEnd     int
    NodeCount   int
}

// 新しいSuffix-treeを作成
func NewSuffixTree(text string) *SuffixTree {
    if !strings.HasSuffix(text, "$") {
        text += "$"
    }

    globalEnd := &EdgeEnd{IsGlobal: true, Position: -1}
    root := &SuffixTreeNode{
        Children: make(map[byte]*SuffixTreeNode),
        Start:    -1,
        End:      &EdgeEnd{IsGlobal: false, Position: -1},
    }

    st := &SuffixTree{
        Root:         root,
        Text:         text,
        GlobalEnd:    globalEnd,
        ActiveNode:   root,
        ActiveEdge:   -1,
        ActiveLength: 0,
        Remaining:    0,
        LeafEnd:      -1,
        NodeCount:    1,
    }

    st.build()
    return st
}

// 新しいノードを作成
func (st *SuffixTree) newNode(start int, end *EdgeEnd) *SuffixTreeNode {
    st.NodeCount++
    return &SuffixTreeNode{
        Children:    make(map[byte]*SuffixTreeNode),
        Start:       start,
        End:         end,
        SuffixIndex: -1,
    }
}

// エッジの長さを取得
func (st *SuffixTree) edgeLength(node *SuffixTreeNode) int {
    if node.End.IsGlobal {
        return st.GlobalEnd.Position - node.Start + 1
    }
    return node.End.Position - node.Start + 1
}

// Suffix-treeを構築（Ukkonen's Algorithm）
func (st *SuffixTree) build() {
    for i := 0; i < len(st.Text); i++ {
        st.extendSuffixTree(i)
    }
}

// 文字を一つ追加してSuffix-treeを拡張
func (st *SuffixTree) extendSuffixTree(pos int) {
    st.GlobalEnd.Position = pos
    st.Remaining++
    lastNewNode := (*SuffixTreeNode)(nil)

    for st.Remaining > 0 {
        if st.ActiveLength == 0 {
            st.ActiveEdge = pos
        }

        char := st.Text[st.ActiveEdge]

        // アクティブノードから該当文字のエッジが存在しない場合
        if st.ActiveNode.Children[char] == nil {
            // 新しいリーフノードを作成
            st.ActiveNode.Children[char] = st.newNode(pos, st.GlobalEnd)

            // 前回作成した内部ノードにsuffix linkを設定
            if lastNewNode != nil {
                lastNewNode.SuffixLink = st.ActiveNode
                lastNewNode = nil
            }
        } else {
            // エッジが存在する場合
            next := st.ActiveNode.Children[char]

            // Walk down（アクティブポイントの調整）
            if st.walkDown(next) {
                continue
            }

            // 現在の文字がエッジ上に存在する場合
            if st.Text[next.Start+st.ActiveLength] == st.Text[pos] {
                // Suffix linkの設定
                if lastNewNode != nil && st.ActiveNode != st.Root {
                    lastNewNode.SuffixLink = st.ActiveNode
                    lastNewNode = nil
                }

                st.ActiveLength++
                break
            }

            // 分岐が必要な場合
            splitEnd := &EdgeEnd{IsGlobal: false, Position: next.Start + st.ActiveLength - 1}
            split := st.newNode(next.Start, splitEnd)
            st.ActiveNode.Children[char] = split

            // 新しいリーフノードを作成
            split.Children[st.Text[pos]] = st.newNode(pos, st.GlobalEnd)
            next.Start += st.ActiveLength
            split.Children[st.Text[next.Start]] = next

            // Suffix linkの設定
            if lastNewNode != nil {
                lastNewNode.SuffixLink = split
            }
            lastNewNode = split
        }

        st.Remaining--

        if st.ActiveNode == st.Root && st.ActiveLength > 0 {
            st.ActiveLength--
            st.ActiveEdge = pos - st.Remaining + 1
        } else if st.ActiveNode != st.Root {
            st.ActiveNode = st.ActiveNode.SuffixLink
        }
    }
}

// Walk down処理
func (st *SuffixTree) walkDown(node *SuffixTreeNode) bool {
    edgeLen := st.edgeLength(node)
    if st.ActiveLength >= edgeLen {
        st.ActiveEdge += edgeLen
        st.ActiveLength -= edgeLen
        st.ActiveNode = node
        return true
    }
    return false
}

// パターン検索
func (st *SuffixTree) Search(pattern string) bool {
    return st.searchHelper(st.Root, pattern, 0)
}

func (st *SuffixTree) searchHelper(node *SuffixTreeNode, pattern string, index int) bool {
    if index == len(pattern) {
        return true
    }

    char := pattern[index]
    if child, exists := node.Children[char]; exists {
        edgeLen := st.edgeLength(child)
        edgeStr := st.Text[child.Start:child.Start+edgeLen]

        // エッジラベルとパターンを比較
        i := 0
        for i < len(edgeStr) && index+i < len(pattern) && edgeStr[i] == pattern[index+i] {
            i++
        }

        if i == len(edgeStr) {
            return st.searchHelper(child, pattern, index+i)
        } else if index+i == len(pattern) {
            return true
        }
    }

    return false
}

// パターンの出現位置をすべて取得
func (st *SuffixTree) FindOccurrences(pattern string) []int {
    var occurrences []int
    node := st.findPatternNode(pattern)
    if node != nil {
        st.collectLeafIndices(node, &occurrences)
    }
    return occurrences
}

func (st *SuffixTree) findPatternNode(pattern string) *SuffixTreeNode {
    node := st.Root
    index := 0

    for index < len(pattern) {
        char := pattern[index]
        child, exists := node.Children[char]
        if !exists {
            return nil
        }

        edgeLen := st.edgeLength(child)
        edgeStr := st.Text[child.Start:child.Start+edgeLen]

        i := 0
        for i < len(edgeStr) && index+i < len(pattern) && edgeStr[i] == pattern[index+i] {
            i++
        }

        if i == len(edgeStr) {
            node = child
            index += i
        } else if index+i == len(pattern) {
            return child
        } else {
            return nil
        }
    }

    return node
}

func (st *SuffixTree) collectLeafIndices(node *SuffixTreeNode, indices *[]int) {
    if len(node.Children) == 0 {
        // リーフノード
        *indices = append(*indices, len(st.Text)-st.edgeLength(node))
        return
    }

    for _, child := range node.Children {
        st.collectLeafIndices(child, indices)
    }
}

// 最長反復部分文字列を検索
func (st *SuffixTree) LongestRepeatedSubstring() string {
    maxLen := 0
    var result string
    st.findLongestRepeatedHelper(st.Root, "", &maxLen, &result)
    return result
}

func (st *SuffixTree) findLongestRepeatedHelper(node *SuffixTreeNode, path string, maxLen *int, result *string) {
    if len(node.Children) > 1 && len(path) > *maxLen {
        *maxLen = len(path)
        *result = path
    }

    for _, child := range node.Children {
        edgeLen := st.edgeLength(child)
        edgeStr := st.Text[child.Start:child.Start+edgeLen]
        newPath := path + edgeStr
        st.findLongestRepeatedHelper(child, newPath, maxLen, result)
    }
}

// 異なる部分文字列の数を計算
func (st *SuffixTree) CountDistinctSubstrings() int {
    return st.countSubstringsHelper(st.Root)
}

func (st *SuffixTree) countSubstringsHelper(node *SuffixTreeNode) int {
    if len(node.Children) == 0 {
        return 0
    }

    count := 0
    for _, child := range node.Children {
        count += st.edgeLength(child)
        count += st.countSubstringsHelper(child)
    }
    return count
}

// Suffix Arrayを生成
func (st *SuffixTree) GenerateSuffixArray() []int {
    var suffixArray []int
    st.dfsForSuffixArray(st.Root, &suffixArray)
    return suffixArray
}

func (st *SuffixTree) dfsForSuffixArray(node *SuffixTreeNode, suffixArray *[]int) {
    if len(node.Children) == 0 {
        // リーフノード: 接尾辞のインデックスを追加
        suffixIndex := len(st.Text) - 1
        for current := node; current != st.Root; {
            // 親ノードを見つける（簡略化）
            break
        }
        *suffixArray = append(*suffixArray, suffixIndex)
        return
    }

    // 子ノードを辞書順でソートして訪問
    for i := 0; i < 256; i++ {
        if child, exists := node.Children[byte(i)]; exists {
            st.dfsForSuffixArray(child, suffixArray)
        }
    }
}

// デバッグ用: 木構造を出力
func (st *SuffixTree) PrintTree() {
    st.printHelper(st.Root, "")
}

func (st *SuffixTree) printHelper(node *SuffixTreeNode, prefix string) {
    if node == st.Root {
        fmt.Println("Root")
    } else {
        edgeLen := st.edgeLength(node)
        edgeStr := st.Text[node.Start:node.Start+edgeLen]
        fmt.Printf("%s%s\n", prefix, edgeStr)
    }

    for _, child := range node.Children {
        st.printHelper(child, prefix+"  ")
    }
}
```

## より高度な実装（一般化 Suffix-tree）

```go
// 一般化Suffix-tree（複数文字列対応）
type GeneralizedSuffixTree struct {
    Root      *SuffixTreeNode
    Strings   []string
    NodeCount int
}

// 複数文字列からGSTを構築
func NewGeneralizedSuffixTree(strings []string) *GeneralizedSuffixTree {
    gst := &GeneralizedSuffixTree{
        Root: &SuffixTreeNode{
            Children: make(map[byte]*SuffixTreeNode),
            Start:    -1,
            End:      &EdgeEnd{IsGlobal: false, Position: -1},
        },
        Strings:   strings,
        NodeCount: 1,
    }

    // 各文字列に一意の終端文字を追加
    combinedText := ""
    for i, str := range strings {
        terminatorChar := byte('$') + byte(i)
        combinedText += str + string(terminatorChar)
    }

    // Suffix-treeを構築
    st := NewSuffixTree(combinedText)
    gst.Root = st.Root

    return gst
}

// 最長共通部分文字列を検索
func (gst *GeneralizedSuffixTree) LongestCommonSubstring() string {
    maxLen := 0
    var result string
    gst.findLCSHelper(gst.Root, "", &maxLen, &result)
    return result
}

func (gst *GeneralizedSuffixTree) findLCSHelper(node *SuffixTreeNode, path string, maxLen *int, result *string) {
    // 複数の文字列からの接尾辞を含むノードかチェック
    stringSet := make(map[int]bool)
    gst.collectStringIndices(node, stringSet)

    if len(stringSet) == len(gst.Strings) && len(path) > *maxLen {
        *maxLen = len(path)
        *result = path
    }

    for _, child := range node.Children {
        // エッジラベルを構築（簡略化）
        newPath := path + "..." // 実際にはエッジの文字列を取得
        gst.findLCSHelper(child, newPath, maxLen, result)
    }
}

func (gst *GeneralizedSuffixTree) collectStringIndices(node *SuffixTreeNode, stringSet map[int]bool) {
    if len(node.Children) == 0 {
        // リーフノード: どの文字列の接尾辞かを判定
        // 実装は簡略化
        return
    }

    for _, child := range node.Children {
        gst.collectStringIndices(child, stringSet)
    }
}
```

# Suffix-tree の応用例

## 1. バイオインフォマティクス・ゲノム解析

- **ゲノム配列解析**: DNA/RNA 配列の共通部分配列検索
- **遺伝子発見**: 遺伝子の反復パターン検出
- **系統解析**: 種間の遺伝的類似性比較
- **配列アライメント**: 複数配列の最適整列

## 2. テキスト処理・自然言語処理

- **文書検索**: 大量文書からの高速パターン検索
- **プラジャリズム検出**: 文書間の類似部分検出
- **テキストマイニング**: 頻出パターンの抽出
- **機械翻訳**: 翻訳メモリの効率的検索

## 3. データ圧縮・符号化

- **テキスト圧縮**: 反復パターンを利用した圧縮
- **辞書構築**: 圧縮アルゴリズムの辞書生成
- **パターン抽出**: データの規則性発見
- **符号化**: 効率的な文字列符号化

## 4. Web・検索エンジン

- **Web 検索**: ページ内容の高速全文検索
- **インデックス構築**: 検索エンジンのインデックス
- **スペルチェック**: 類似文字列の検索
- **オートコンプリート**: 入力候補の高速提示

## 5. ソフトウェア開発・版管理

- **差分検出**: ファイル間の差分計算
- **リファクタリング**: コードの重複検出
- **版管理**: ソースコードの変更履歴管理
- **コード解析**: プログラムパターンの抽出

## 6. データベース・情報検索

- **文字列インデックス**: データベースの文字列カラムインデックス
- **範囲検索**: 文字列の範囲クエリ最適化
- **パターンマッチング**: 複雑な文字列検索条件
- **全文検索**: 大規模データの全文検索

## 7. セキュリティ・暗号

- **マルウェア検出**: 悪意あるコードパターン検出
- **侵入検知**: ネットワークトラフィックの異常検出
- **フォレンジック**: デジタル証拠の文字列解析
- **パスワード解析**: 弱いパスワードパターン検出

## 8. 画像・音声処理

- **画像内テキスト**: OCR 結果の文字列処理
- **音声認識**: 音素パターンの検索
- **マルチメディア検索**: メタデータの文字列検索
- **パターン認識**: 1 次元信号のパターンマッチング

# Suffix-tree と他のデータ構造との比較

## 文字列データ構造比較

| 特徴                   | Suffix-tree     | Suffix Array | Trie          | KMP     | ハッシュテーブル | 二分探索木   |
| ---------------------- | --------------- | ------------ | ------------- | ------- | ---------------- | ------------ |
| **構築時間**           | O(n)            | O(n log n)   | O(n×m)        | O(m)    | O(n)             | O(n log n)   |
| **空間使用量**         | O(n)            | O(n)         | O(n×m)        | O(m)    | O(n)             | O(n)         |
| **パターン検索**       | O(m)            | O(m log n)   | O(m)          | O(n+m)  | O(m)平均         | O(m log n)   |
| **全出現位置**         | O(m+k)          | O(m+k log n) | O(m+k)        | O(n+m)  | O(n)             | O(m log n+k) |
| **最長共通部分文字列** | O(n)            | O(n²)        | ❌ 不適       | ❌ 不適 | ❌ 不適          | ❌ 不適      |
| **動的更新**           | ✅ O(1)追加     | ❌ 再構築    | ✅ O(m)       | ❌ 不適 | ✅ O(1)平均      | ✅ O(log n)  |
| **実装複雑度**         | ❌ 非常に複雑   | ⚠️ 中程度    | ✅ 簡単       | ✅ 簡単 | ✅ 簡単          | ⚠️ 中程度    |
| **メモリ効率**         | ⚠️ ポインタ多用 | ✅ 高効率    | ❌ 大量メモリ | ✅ 最小 | ⚠️ 負荷率依存    | ⚠️ ポインタ  |

## 用途別最適選択

| 用途                   | 第 1 選択    | 第 2 選択        | 第 3 選択        | 避けるべき   |
| ---------------------- | ------------ | ---------------- | ---------------- | ------------ |
| **複数パターン検索**   | Suffix-tree  | Suffix Array     | Trie             | KMP          |
| **最長共通部分文字列** | Suffix-tree  | -                | -                | 他すべて     |
| **大量文書検索**       | Suffix Array | Suffix-tree      | ハッシュテーブル | Trie         |
| **リアルタイム検索**   | KMP          | Suffix-tree      | ハッシュテーブル | Suffix Array |
| **メモリ制約環境**     | Suffix Array | KMP              | ハッシュテーブル | Suffix-tree  |
| **プレフィックス検索** | Trie         | Suffix-tree      | 二分探索木       | KMP          |
| **動的文字列更新**     | Suffix-tree  | Trie             | ハッシュテーブル | Suffix Array |
| **ゲノム解析**         | Suffix-tree  | Suffix Array     | -                | Trie、KMP    |
| **テキストマイニング** | Suffix-tree  | Suffix Array     | -                | 他すべて     |
| **単純なパターン検索** | KMP          | ハッシュテーブル | 二分探索木       | Suffix-tree  |

# パフォーマンス特性

## Suffix-tree 実装方法別比較

| 実装方法                    | 構築時間 | 空間使用量 | 検索時間 | 実装複雑度    | 適用場面     |
| --------------------------- | -------- | ---------- | -------- | ------------- | ------------ |
| **Ukkonen's Algorithm**     | O(n)     | O(n)       | O(m)     | ❌ 非常に複雑 | 高性能要求   |
| **McCreight's Algorithm**   | O(n)     | O(n)       | O(m)     | ❌ 複雑       | バッチ処理   |
| **Naive Algorithm**         | O(n²)    | O(n)       | O(m)     | ✅ 簡単       | 教育・小規模 |
| **Compressed Suffix-tree**  | O(n)     | O(n)圧縮   | O(m)     | ❌ 複雑       | メモリ制約   |
| **Generalized Suffix-tree** | O(Σn)    | O(Σn)      | O(m)     | ❌ 非常に複雑 | 複数文字列   |

## データサイズ別推奨実装

| データサイズ          | 推奨実装                | 理由               |
| --------------------- | ----------------------- | ------------------ |
| **小規模（<1KB）**    | Naive Algorithm         | 実装の簡潔性       |
| **中規模（1KB-1MB）** | Ukkonen's Algorithm     | バランスの良い性能 |
| **大規模（1MB-1GB）** | Suffix Array            | メモリ効率         |
| **超大規模（>1GB）**  | 外部メモリ版            | ディスクベース処理 |
| **複数文字列**        | Generalized Suffix-tree | 共通部分文字列検索 |
| **動的更新**          | Ukkonen's Algorithm     | オンライン構築     |

## 文字列特性別最適化

| 文字列特性                           | 最適化手法              | 効果             |
| ------------------------------------ | ----------------------- | ---------------- |
| **DNA 配列（4 文字アルファベット）** | 2bit 符号化             | メモリ使用量 1/4 |
| **テキスト（ASCII）**                | 配列ベース子ノード      | 高速アクセス     |
| **Unicode 文字列**                   | ハッシュマップ子ノード  | 柔軟な文字対応   |
| **大量の反復**                       | エッジ圧縮強化          | 大幅な空間節約   |
| **短い文字列多数**                   | Generalized Suffix-tree | 効率的な管理     |

# まとめ

Suffix-tree は、文字列のすべての接尾辞を効率的に管理する高度なデータ構造です。線形時間での構築と検索が可能で、最長共通部分文字列の発見、パターンマッチング、文字列解析などの複雑な操作を高速に実行できます。

実装は非常に複雑ですが、バイオインフォマティクス、テキスト処理、データ圧縮などの分野で強力な効果を発揮します。Go 言
