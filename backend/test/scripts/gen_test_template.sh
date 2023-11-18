#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SOURCE_DIR="${SCRIPT_DIR}/../../internal/"
TEST_FILE_DIR="${SCRIPT_DIR}/../test/"
FILELIST="$SCRIPT_DIR/gofiles.txt"

# 引数の数が正しいかチェック
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 FILE"
    exit 1
fi

FILE=$1
BASE_NAME="$(basename "$FILE")"
DIR_PATH="$(dirname "$FILE")"
FILE_NAME="${BASE_NAME%.*}"  # 拡張子を取り除いたファイル名
EXTENSION="${BASE_NAME##*.}"  # 拡張子

# コマンドの実行
OUTPUT_FILE=${TEST_FILE_DIR}${DIR_PATH}/${FILE_NAME}_test.${EXTENSION}
echo "出力先: ${OUTPUT_FILE}"
gotests -exported -template testify "${SOURCE_DIR}${FILE}" > ${OUTPUT_FILE}

# 終了ステータスの確認
if [ $? -eq 0 ]; then
    echo "テストファイルの生成が成功しました。"
else
    echo "エラーが発生しました。"
fi
