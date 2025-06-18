package repositories



import (
	"database/sql"
)
// TodoRepository is a struct that holds the database connection

type TodoRepository struct{
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]Todo, error) {

	rows, err := r.db.Query("Select id, title, completed FROM todos")
	if err !=nil{
		return nil, err
	}
// Ensure that the rows are closed after we're done
	defer func() {
		rows.Close()
	}()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}