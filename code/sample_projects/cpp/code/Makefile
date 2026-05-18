CXX = g++
CXXFLAGS = -Wall -Wextra -std=c++23

# gTest
GTEST_BIN = build/gtest_tests
GTEST_SRC = tests/gtest/gtest.cpp

# Catch2
CATCH_BIN = build/catch_tests
CATCH_SRC = tests/catch2/catch2.cpp

# testAuditor
AUDITOR_BIN = ./testAuditor

all: test_gtest test_catch

# gtest build
$(GTEST_BIN): $(GTEST_SRC)
	mkdir -p build
	$(CXX) $(CXXFLAGS) $(GTEST_SRC) -o $(GTEST_BIN) -lgtest -lgtest_main -pthread

test_gtest: $(GTEST_BIN)
	./$(GTEST_BIN) --gtest_output=xml:tests/out/gtest/report_gtest.xml

# catch2 build
$(CATCH_BIN): $(CATCH_SRC)
	mkdir -p build
	$(CXX) $(CXXFLAGS) $(CATCH_SRC) -o $(CATCH_BIN)

test_catch: $(CATCH_BIN)
	./$(CATCH_BIN) -r junit -o tests/out/catch2/report_catch.xml

discover_gtest: $(GTEST_BIN)
	@echo '<?xml version="1.0" encoding="utf-8"?>'
	@echo '<testsuite>'
	@./$(GTEST_BIN) --gtest_list_tests | awk '\
	/Running main/ { next } \
	/\.$$/ { suite=$$0; sub(/\.$$/, "", suite); next } \
	/^[[:space:]]/ { \
		gsub(/^ +/, "", $$0); \
		name=$$0; \
		class=suite; \
		printf("    <testcase classname=\"%s\" name=\"%s\" qualifiedName=\"%s::%s\"/>\n", class, name, class, name); \
	}'
	@echo '</testsuite>'

discover_catch: $(CATCH_BIN)
	@echo '<?xml version="1.0" encoding="utf-8"?>'
	@echo '<testsuite>'
	@./$(CATCH_BIN) --list-tests | awk '\
	/^All available/ { next } \
	/^[0-9]+ test case/ { next } \
	/^[[:space:]]*$$/ { next } \
	{ \
		gsub(/^ +/, "", $$0); \
		name=$$0; \
		printf("    <testcase classname=\"catch_tests.global\" name=\"%s\" qualifiedName=\"%s\"/>\n", name, name); \
	}'
	@echo '</testsuite>'

audit_gtest: test_gtest $(AUDITOR_BIN)
	./$(AUDITOR_BIN) \
		--disc "make discover_gtest" \
		--junit tests/out/gtest \
		--out tests/out/gtest/audit_gtest.xml

audit_catch: test_catch $(AUDITOR_BIN)
	./$(AUDITOR_BIN) \
		--disc "make discover_catch" \
		--junit tests/out/catch2 \
		--out tests/out/catch2/audit_catch.xml

clean:
	rm -rf build *.xml
