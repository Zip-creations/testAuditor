#!/usr/bin/bash

PYTHONPATH=src python3 -m pytest tests/tests_simple.py tests/tests_identity.py --junit-xml="tests/test-reports/report.xml"
