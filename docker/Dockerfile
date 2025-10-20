# ベースイメージとして公式の Go イメージを使用
FROM golang:1.24

# 作業ディレクトリを設定
WORKDIR /app

# ホストのコードをコンテナにコピー
COPY code/ .

# 必要なモジュールをダウンロード
RUN go mod init my-data-structures-go || true
RUN go mod tidy

# デフォルトのエントリーポイントを設定
# 実行する .go ファイルを引数で指定可能
ENTRYPOINT ["go", "run"]
