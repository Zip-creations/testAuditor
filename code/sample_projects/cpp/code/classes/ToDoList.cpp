#include <map>
#include <string>

class ToDoItem {
public:
    std::string content;
    ToDoItem(const std::string& content) {
        this->content = content;
    };
};

class ToDoList {
private:
    std::map<int, ToDoItem> items;
    int nextID = 0;
public:
    void addItem(const ToDoItem& item){
        items.emplace(nextID, item);
        nextID++;
    };
    void removeItem(int id) {
        items.erase(id);
    };
    ToDoItem getItemByID(int id) const {
        return items.at(id);
    };
};
