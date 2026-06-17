#!/usr/bin/bash

PYTHONPATH_DIR="code"
REPORT_PATH="test/out/reportNew.xml"

mkdir -p "$(dirname "$REPORT_PATH")"

# Unfortunately, pytest forces a path to write into.
# The file created by the test execution from testAuditor can be ignored, since it's already stored in the gitnotes
PYTHONPATH="$PYTHONPATH_DIR" python3 -m pytest "$@" --junit-xml="$REPORT_PATH" >&2

cat "$REPORT_PATH"
