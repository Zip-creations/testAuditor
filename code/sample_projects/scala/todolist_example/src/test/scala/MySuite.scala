import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions._
import org.junit.jupiter.api.Disabled
import todolist.*
import scala.collection.mutable.Map


class TestAddingItems {
  @Test
  def testTodo(): Unit = {
    var testlist = ToDoList.empty()
    var a = testlist.addItem(ToDoItem("testItem1"))
    assertEquals(testlist.items(1).content, "testItem1")
  }
}

class TestRemovingItems {
  @Test
  def testTodo(): Unit = {
    var testlist = ToDoList.empty()
    var a = testlist.addItem(ToDoItem("testItem1"))
    testlist.removeItemByID(1)
    assertFalse(testlist.items.contains(1))
  }

  @Test // This test is expected to fail, to show how the tooling is handling failing tests
  def testPositive(): Unit = {
    var testlist = ToDoList.empty()
    testlist.addItem(ToDoItem("testItem1"))
    assertEquals(testlist, ToDoList.empty())
  }
}

class TestCopy {
  @Test
  def testEquality(): Unit = {
    var list1 = ToDoList.empty()
    var list2 = ToDoList.empty()
    var list3 = list1.copy()
    assertEquals(list1, list2)
    assertEquals(list1, list3)
    assertEquals(list2, list3)
  }

  @Test
  def testCopy(): Unit = {
    var list1 = ToDoList.empty()
    var list2 = list1.copy()
    assertEquals(list1, list2)
    list2.addItem(ToDoItem("testItem1"))
    var list3 = list2.copy()
    assertNotEquals(list1, list2)
    list2.removeAllItems()
    assertEquals(list1, list2)
    assertNotEquals(list2, list3)
  }

  @Test
  @Disabled("This test will be skipped, to show how the pipeline tooling is handling this")
  def testDisabled(): Unit = {
    var list1 = ToDoList.empty()
    var list2 = list1.copy()
    assertEquals(list1, list2)
  }
}
