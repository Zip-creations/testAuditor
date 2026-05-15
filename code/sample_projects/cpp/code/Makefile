CXX = g++
CXXFLAGS = -Wall -Wextra -std=c++23

# gTest
GTEST_BIN = build/gtest_tests
GTEST_SRC = tests/gtest/gtest.cpp

# Catch2
CATCH_BIN = build/catch_tests
CATCH_SRC = tests/catch2/catch2.cpp

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
	@./$(GTEST_BIN) --gtest_list_tests | awk '\
	/\.$$/ { suite=$$0; sub(/\.$$/, "", suite); next } \
	{ gsub(/^ +/, "", $$0); print suite "::" $$0 }'

discover_catch: $(CATCH_BIN)
	@./$(CATCH_BIN) --list-tests | sed '1d;$$d'

clean:
	rm -rf build *.xml
