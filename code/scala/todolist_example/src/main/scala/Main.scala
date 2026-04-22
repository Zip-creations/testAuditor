import scala.collection.mutable.Map

@main def hello(): Unit =
  println("Hello world!")
  println(msg)
  var testlist = TODOList("sample", Map.empty[Int, TODOItem])
  testlist.addEntry(TODOItem("testItem1"))
  print(testlist.items)

def msg = "I was compiled by Scala 3. :)"

class TODOList(
  var author: String,
  var items: Map[Int, TODOItem]
) {
  var id = 0
  def newID = () =>
    id += 1
    id
  def addEntry = (item: TODOItem) =>
    items(newID()) = item
}
  // var items = Map.empty[Int, TODOItem]
  // if !itemsInit.isEmpty
  // then 
  //   var item = itemsInit    

case class TODOItem (
  content: String
)
