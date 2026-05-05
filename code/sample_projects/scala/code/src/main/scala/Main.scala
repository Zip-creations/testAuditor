import scala.collection.mutable.Map
import todolist.*

@main def hello(): Unit =
  println("Hello world!")
  println(msg)
  var testlist = ToDoList("sample", Map.empty[Int, ToDoItem])
  var a = testlist.addItem(ToDoItem("testItem1"))
  // print(testlist.items.toString + "\n")
  testlist.removeItemByID(1)
  // print(testlist.items.toString + "\n")

def msg = "I was compiled by Scala 3. :)"
