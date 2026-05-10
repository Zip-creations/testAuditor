#include <iostream>
# include "classes/ToDoList.cpp"

int main() {
    std::cout << "Hello World!";
    ToDoList list;
    list.addItem(ToDoItem("testItem1"));
    std::cout << list.getItemByID(0).content;
    return 0;
}
