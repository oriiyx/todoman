package main

import "log"

func (m *Model) fetchTodos() {
	rows, err := m.db.Query("SELECT id, title, content FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Content); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	m.todos = todos
}

func (m *Model) saveTodo() {
	_, err := m.db.Exec("INSERT INTO todos (title, content) VALUES (?, ?)", "New Todo", "This is a new todo")
	if err != nil {
		log.Fatal(err)
	}
	m.fetchTodos()
}
