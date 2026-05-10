#!/usr/bin/bash

PYTHONPATH=code python3 -m pytest --collect-only -q #tests/tests_simple.py tests/tests_identity.py --junit-xml="tests/out/report.xml"
# TEST_FILES=("tests/tests_simple.py" "tests/tests_identity.py")

# if [ "$#" -eq 0 ]; then
#     PYTHONPATH=src python3 -m pytest "${TEST_FILES[@]}" --junit-xml="tests/out/report.xml"
#     exit 0
# fi

# TEST_TARGETS=()

# for test_name in "$@"; do
#     for file in "${TEST_FILES[@]}"; do
#         TEST_TARGETS+=("${file}::${test_name}")
#     done
# done

# PYTHONPATH=src python3 -m pytest "${TEST_TARGETS[@]}" --junit-xml="tests/out/report.xml"
