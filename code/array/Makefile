# Makefile for Go Array Implementations

.PHONY: all demo comparison custom clean

# すべてのプログラムを実行
all: demo comparison custom

# カスタム実装のデモのみ
demo:
	@echo "=== カスタム配列実装のデモ ==="
	go run custom_array.go custom_array_demo.go

# 組み込み配列との比較
comparison:
	@echo "=== 配列の比較・パフォーマンステスト ==="
	go run custom_array.go array_comparison.go

# 元のファイル（後方互換性のため）
custom:
	@echo "=== 元のarray.goファイル ==="
	go run array.go

# クリーンアップ
clean:
	rm -f *.exe *.out

# ヘルプ
help:
	@echo "利用可能なコマンド:"
	@echo "  make demo       - カスタム配列実装のデモを実行"
	@echo "  make comparison - 組み込み配列との比較を実行"
	@echo "  make custom     - 元のarray.goファイルを実行"
	@echo "  make all        - すべてのプログラムを実行"
	@echo "  make clean      - 生成されたファイルをクリーンアップ"
