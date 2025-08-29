# Suffix-tree（接尾辞木）とは

Suffix-tree（接尾辞木）は、ある文字列の**すべての接尾辞（suffix）**を格納した、圧縮されたトライ木である。文字列の各位置から始まるすべての部分文字列を圧縮された形で格納し、文字列の様々な操作を高速に実行できる。テキストを一度前処理しておくことで、その後の複雑な部分文字列クエリに極めて高速に応答することを可能にする。

**文字列検索、最長共通部分文字列の発見、パターンマッチングを線形時間で実行できる。**

## 特徴

ルートノードから各葉ノードへのパスは、文字列の一意な接頭辞に対応する。

全ての接尾辞が必ず葉ノードで終わるようにするため、元の文字列の末尾に、その文字列に現れない特殊な文字（終端文字、$ など）を付加することが一般的である。

内部ノード（2 つ以上の子を持つノード）は、文字列 T 内で 2 回以上出現する反復部分文字列に対応する。特に、根から内部ノードへのパス文字列が、その反復部分文字列となる。

## 計算量

上記の特徴により、文字列 T（長さ n）に対する部分文字列 P（長さ m）の存在確認は、$O(m)$という、元の文字列の長さに依存しない時間計算量で実行できる。これは、Suffix Tree をたどることで、$P$が$T$のいずれかの接尾辞の接頭辞であるかどうかを判定できるためである。

## ユースケース

O(n) の前処理の後、任意のパターン(長さ $m$)がテキスト内に存在するかを $O(m)$ で検索できます。その他、最長共通部分文字列の発見、最長反復部分文字列の発見、DNA シーケンス解析などのバイオインフォマティクス分野で絶大な力を発揮します。

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

サフィックス木は強力である一方、実装が複雑でメモリ消費が大きいという課題がある。
そのため、実用上はより省メモリな**サフィックス配列（Suffix Array）**が代替として好まれることが多い。
サフィックス配列は、すべての接尾辞を辞書順にソートし、その開始インデックスを格納しただけの単純な整数配列である。つまり、SA[i]は、辞書順で i 番目に小さい接尾辞の開始インデックスを保持する。
LCP 配列（Longest Common Prefix array; 最長共通接頭辞）などの補助的なデータ構造と組み合わせることで、サフィックス木の機能の多くを、より少ないメモリで実現できる。

Go の標準ライブラリには、このサフィックス配列を実装した index/suffixarray パッケージが提供されており、効率的な部分文字列検索に利用できる。

# Suffix-tree の特徴

- **線形空間**: O(n) のメモリ使用量（n: 文字列長）
- **線形時間構築**: Ukkonen's algorithm で O(n) 時間で構築
- **高速検索**: パターン検索が O(m) 時間（m: パターン長）
- **最長共通部分文字列（LCS）**: 複数の文字列に共通して現れる最も長い部分文字列を効率的に見つけられる
- **最長反復部分文字列(LRS)**: 一つの文字列内で複数回出現する最も長い部分文字列を見つけられる。Suffix Tree における最も深い内部ノードに対応する。
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

# Suffix-tree の実装方法

## 一般的な実装方法

### Naive Algorithm（単純法）

各接尾辞を一つずつ木構造に挿入する。実装は簡単だが、計算量は O(n²)っとなり、長い文字列に対して利用するのは非現実的である。

### Ukkonen's Algorithm（最も効率的）

Suffix Tree をオンライン（文字列を先頭から一文字ずつ処理する）かつ O(n) の線形時間で構築できるアルゴリズム。

以下のような概念を利用して、効率的なデータ構造を実装できる。

- **オンライン構築**: 文字列を$T[1..i]からT[1..i+1]$へと一文字ずつ拡張しながら、木を段階的に更新していく。
- **サフィックスリンク**:文字列 xA（x は文字、A は文字列）に対応するノードから、文字列 A に対応するノードへのポインタ。これにより、関連する接尾辞を表すノード間を高速に移動でき、構築の効率化に大きく貢献する。
- **案目的ノード**: 実際にノードを作成するのを、分岐が必要になるまで遅延させるテクニック。エッジ上の位置としてノードを表現することで、空間と時間の両方を節約する。

### McCreight's Algorithm

- **効率的**: O(n) 時間構築
- **バッチ処理**: すべての接尾辞を一度に処理
- **実装の複雑度**: 中程度

## Go による基本的な Suffix-tree 実装例

接尾辞木をゼロから正しく実装するのは複雑な作業になるため、実用上は、github.com/spacewander/go-suffix-tree のような既存のライブラリや、類似の機能を持つラディックス木(Radix Tree)の実装である github.com/armon/go-radix などを利用するのが現実的な選択肢となる。

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

# 実装方法のトレードオフ

Suffix Tree と Suffix Array は、同じ問題領域を解決するための代替アプローチであるが、それぞれ適した利用場面が存在する。

| 特徴               | Suffix Tree                                     | Suffix Array (+ LCP Array)                                        |
| ------------------ | ----------------------------------------------- | ----------------------------------------------------------------- |
| **空間計算量**     | ❌ 高い（例：20n ～ 40n バイト）                | ✅ 低い（例：5n ～ 9n バイト）                                    |
| **構築時間**       | ✅ O(n) (Ukkonen 法)                            | O(n) or O(nlogn) (より単純なアルゴリズム)                         |
| **部分文字列検索** | ✅ O(m) (m=パターン長)                          | ⚠️ O(mlogn) or O(m+logn)                                          |
| **実装の複雑さ**   | ❌ 非常に高い                                   | 中程度                                                            |
| **主な用途**       | ⚠️ 理論的なアルゴリズム、$O(m)$検索が必須な応用 | ✅ ほとんどの実用的な応用、バイオインフォマティクス、テキスト索引 |

Suffix Array は空間効率と実装の容易さで Suffix Tree を大幅に上回る。多くの応用では、検索時間の計算量が$O(m log n)$でも十分に高速であり、メモリ使用量の削減という利点がそれを上回るため、Suffix Array が好まれる傾向にある。
