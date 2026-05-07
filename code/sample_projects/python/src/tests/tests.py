from classes.ToDoList import ToDoList
from junit_xml import TestCase, TestSuite

# replace with pytest (with junit xml)
test_cases = [TestCase('Test1', 'some.class.name', 123.345, 'I am stdout!', 'I am stderr!')]
ts = TestSuite("my test suite", test_cases)
print(TestSuite.to_xml_string([ts]))
