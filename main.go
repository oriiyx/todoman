package main

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed todos.db
var embeddedDB embed.FS

type Model struct {
	db      *sql.DB
	todos   []Todo
	cursor  int
	content string
}

type Todo struct {
	ID      int
	Title   string
	Content string
}

func main() {
	db := initDB()
	defer db.Close()

	p := tea.NewProgram(Model{db: db})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB() *sql.DB {
	// Extract embedded database
	data, err := fs.ReadFile(embeddedDB, "todos.db")
	if err != nil {
		log.Fatal(err)
	}

	// Write to a temporary file
	tmpFile, err := os.CreateTemp("", "todos-*.db")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	// Open the database
	db, err := sql.Open("sqlite3", tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"content" TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "TODO App\nPress ctrl+c to quit."
}
