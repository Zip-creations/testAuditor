import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions._
import org.junit.jupiter.api.Disabled
import todolist.*
import scala.collection.mutable.Map


// @Suite?
class TestAddingItems {
  @Test
  def testTodo(): Unit = {
    println("Pos1")
    var testlist = ToDoList.empty()
    var a = testlist.addItem(ToDoItem("testItem1"))
    assertEquals(testlist.items(1).content, "testItem1")
  }
}

class TestRemovingItems {
  @Test
  def testTodo(): Unit = {
    println("Pos2")

    var testlist = ToDoList.empty()
    var a = testlist.addItem(ToDoItem("testItem1"))
    testlist.removeItemByID(1)
    assertFalse(testlist.items.contains(1))
  }

  @Test // This test is expected to fail, to show how the tooling is handling failing tests
  def testPositive(): Unit = {
    println("Pos3")
    var testlist = ToDoList.empty()
    testlist.addItem(ToDoItem("testItem1"))
    assertEquals(testlist, ToDoList.empty())
  }
}
