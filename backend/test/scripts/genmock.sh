#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SOURCE_DIR="${SCRIPT_DIR}/../../internal/"
MOCK_RELATIVE_DIR="${SCRIPT_DIR}/../mock/"
FILELIST="$SCRIPT_DIR/gofiles.txt"

while IFS= read -r FILE; do
  FULL_PATH="$SOURCE_DIR${FILE}"
  BASE_NAME="$(basename "$FILE")"
  DIR_PATH="$(dirname "$FILE")"
  FILE_NAME="${BASE_NAME%.*}"  # 拡張子を取り除いたファイル名
  EXTENSION="${BASE_NAME##*.}"  # 拡張子
  if [ -f "$FULL_PATH" ]; then
    echo "Processing $FULL_PATH..."
    mockgen --source $FULL_PATH -destination "${MOCK_RELATIVE_DIR}/${DIR_PATH}/${FILE_NAME}_mock.${EXTENSION}"
  else
    echo "Dir not found: $FULL_PATH"
    echo "モック生成が失敗しました。"
    exit 1;
  fi
done < "$FILELIST"

echo "モック生成が完了しました。"
