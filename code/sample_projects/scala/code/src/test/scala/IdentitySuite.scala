import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions._
import org.junit.jupiter.api.Disabled
import todolist.*
import scala.collection.mutable.Map

class TestCopy {
  @Test
  def testEquality(): Unit = {
    println("Pos4")
    var list1 = ToDoList.empty()
    var list2 = ToDoList.empty()
    var list3 = list1.copy()
    assertEquals(list1, list2)
    assertEquals(list1, list3)
    assertEquals(list2, list3)
  }

  @Test
  def testCopy(): Unit = {
    println("Pos5")
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
    println("Pos6")
    var list1 = ToDoList.empty()
    var list2 = list1.copy()
    assertEquals(list1, list2)
  }
}
