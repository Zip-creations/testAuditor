# base Output:
# test/test_identity.py::test_copy
# test/test_simple.py::test_add_item
# test/test_simple.py::test_removing_items
# test/test_simple.py::test_skipping

OUTPUT=$(PYTHONPATH=code python3 -m pytest --collect-only -q)

RESULT="("
FIRST=1

while IFS= read -r line; do
    file="${line%%::*}"
    test="${line##*::}"

    if [ $FIRST -eq 0 ]; then
        RESULT+=", "
    fi

    RESULT+="(${file}, ${test})"
    FIRST=0
done <<< "$OUTPUT"

RESULT+=")"

echo "$RESULT"
