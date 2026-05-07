class ToDoItem:
    def __init__(self, item: str):
        self.item = item

class ToDoList:
    def __init__(self, author: str):
        self.author = author
        self.items: dict[int, ToDoItem] = {}
        self.id = 0

    def newID(self):
        self.id += 1
        return self.id

    def addItem(self, item: ToDoItem):
        self.items[self.newID()] = item

    def removeItemByID(self, id: int):
        del self.items[id]
