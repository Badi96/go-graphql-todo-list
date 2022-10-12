package todo

//use to store data/todos
var TodoItems = []Todo{}

// used to assign unique id to each todo
var count = 1

// The queries

//Get all todos
func GetTodos() []Todo {
	return TodoItems
}

//Get a todo by id
func GetTodo(id int) Todo {
	for _, todo := range TodoItems {
		if todo.ID == id {
			return todo
		}
	}
	return Todo{}
}

//Mutations
func AddTodo(title string) *Todo {
	//Create a new todo
	temp := &Todo{
		ID:        count,
		Title:     title,
		Completed: false,
	}

	//add it to the list of todos
	TodoItems = append(TodoItems, *temp)
	//increment the count for id
	count++
	return temp
}

func UpdateTodo(id int, title string, compelted bool) bool {
	for i, todo := range TodoItems {
		if todo.ID == id {
			TodoItems[i].Title = title
			TodoItems[i].Completed = compelted
			return true
		}
	}
	return false
}

func DeleteTodo(id int) bool {
	for i, todo := range TodoItems {
		if todo.ID == id {
			TodoItems = append(TodoItems[:i], TodoItems[i+1:]...)
			return true
		}
	}
	return false
}
