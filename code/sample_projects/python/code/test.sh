#!/usr/bin/bash

PYTHONPATH_DIR="code"
REPORT_PATH="test/out/reportNew.xml"

mkdir -p "$(dirname "$REPORT_PATH")"

PYTHONPATH="$PYTHONPATH_DIR" python3 -m pytest "$@" --junit-xml="$REPORT_PATH" >&2

cat "$REPORT_PATH"
