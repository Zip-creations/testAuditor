import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions._
import todolist.*
import scala.collection.mutable.Map


class TestAddingItems {
  @Test
  def testTodo(): Unit = {
    var testlist = TODOList.empty()
    var a = testlist.addItem(TODOItem("testItem1"))
    assertEquals(testlist.items(1).content, "testItem1")
  }
}

class TestRemovingItems {
  @Test
  def testTodo(): Unit = {
    var testlist = TODOList.empty()
    var a = testlist.addItem(TODOItem("testItem1"))
    testlist.removeItemByID(1)
    assertFalse(testlist.items.contains(1))
  }
}

class TestCopy {
  @Test
  def testEquality(): Unit = {
    var list1 = TODOList.empty()
    var list2 = TODOList.empty()
    var list3 = list1.copy()
    assertEquals(list1, list2)
    assertEquals(list1, list3)
    assertEquals(list2, list3)
  }

  @Test
  def testCopy(): Unit = {
    var list1 = TODOList.empty()
    var list2 = list1.copy()
    assertEquals(list1, list2)
    list2.addItem(TODOItem("testItem1"))
    var list3 = list2.copy()
    assertNotEquals(list1, list2)
    list2.removeAllItems()
    assertEquals(list1, list2)
    assertNotEquals(list2, list3)
  }
}
