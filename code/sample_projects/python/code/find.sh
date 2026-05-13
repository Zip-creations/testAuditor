# base Output:
# test/test_identity.py::test_copy
# test/test_simple.py::test_add_item
# test/test_simple.py::test_removing_items
# test/test_simple.py::test_skipping

OUTPUT=$(PYTHONPATH=code python3 -m pytest --collect-only -q)

echo '<?xml version="1.0" encoding="utf-8"?>'
echo '<testsuites>'

while IFS= read -r line; do
    [[ "$line" != *"::"* ]] && continue
    echo "    <testcase qualifiedName=\"$line\"/>"
done <<< "$OUTPUT"

echo '</testsuites>'
