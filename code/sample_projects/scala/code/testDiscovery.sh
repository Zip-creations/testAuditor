# WIP - script not working currently!
TEST_LIST=$(sbt "show Test/definedTestNames")

# TEST_LIST=$(sbt --batch --quiet "export Test/definedTestNames")

# if [ -z "$TEST_LIST" ]; then
#     echo "No tests found. Check your sbt-jupiter-interface configuration."
# else
#     echo "Discovered Tests:"
#     # Clean up sbt output formatting (removes 'List(' and ')')
#     echo "$TEST_LIST" | sed 's/List(//g' | sed 's/)//g' | tr ',' '\n' | xargs -n1
# fi
