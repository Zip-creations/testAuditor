#!/usr/bin/env bash
# example output of find.sh, pytest style
echo '<?xml version="1.0" encoding="utf-8"?>'
echo '<testsuite>'
echo '    <testcase classname="test.test_identity" name="test_copy" qualifiedName="test/test_identity.py::test_copy"/>'
echo '    <testcase classname="test.test_simple" name="test_add_item" qualifiedName="test/test_simple.py::test_add_item"/>'
echo '    <testcase classname="test.test_simple" name="test_removing_items" qualifiedName="test/test_simple.py::test_removing_items"/>'
echo '    <testcase classname="test.test_simple" name="test_skipping" qualifiedName="test/test_simple.py::test_skipping"/>'
echo '    <testcase classname="TestAddingItems" name="testTodo()" qualifiedName=""/>'
echo '    <testcase classname="TestRemovingItems" name="testPositive()" qualifiedName=""/>'
echo '    <testcase classname="ThisTest" name="WasNotExecuted" qualifiedName=""/>'
echo '</testsuite>'
