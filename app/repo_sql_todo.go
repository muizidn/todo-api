package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type SQLTodoRepo struct {
	db *sql.DB
}

func (s *SQLTodoRepo) List() ([]*Todo, error) {
	todos := make([]*Todo, 0)
	{
		rows, err := s.db.Query(`SELECT * FROM todo`)
		if err != nil {
			return nil, log.TError(
				grpc.Errorf(
					codes.Internal,
					"Query Todo error",
				),
				err,
			)
		}
		for rows.Next() {
			todo := Todo{}
			if err := rows.Scan(&todo.uuid, &todo.title, &todo.description, &todo.createdAt); err != nil {
				return nil, log.TError(
					grpc.Errorf(
						codes.Internal,
						"Query Todo error",
					),
					err,
				)
			}
			todos = append(todos, &todo)
		}
	}
	return todos, nil
}
func (s *SQLTodoRepo) Get(uuid string) (*Todo, error) {
	todo := Todo{}
	row := s.db.QueryRow(`SELECT * FROM todo WHERE id=?`, uuid)
	if err := row.Scan(&todo.uuid, &todo.title, &todo.description, &todo.createdAt); err != nil {
		return nil, log.TError(
			grpc.Errorf(
				codes.Internal,
				"Get Todo error",
			),
			err,
		)
	}
	return &todo, nil
}
func (s *SQLTodoRepo) Create(title string, description string) (*Todo, error) {
	newId := uuid.New()
	{
		cs, ps, vs := sqlStatementBuildWithValues(SqlPair{
			"id":          newId,
			"title":       title,
			"description": description,
			"created_at":  time.Now().UnixNano(),
		})
		stmt := fmt.Sprintf(`INSERT INTO todo %s VALUES %s`, cs, ps)
		_, err := s.db.Exec(stmt, vs...)
		if err != nil {
			return nil, log.TError(
				grpc.Errorf(
					codes.Internal,
					"Create Todo error",
				),
				err,
			)
		}
	}
	todo := Todo{}
	{
		row := s.db.QueryRow(`SELECT * FROM todo WHERE id=?`, newId)
		if err := row.Scan(&todo.uuid, &todo.title, &todo.description, &todo.createdAt); err != nil {
			return nil, log.TError(
				grpc.Errorf(
					codes.Internal,
					"Create Todo error",
				),
				err,
			)
		}
	}
	return &todo, nil
}
func (s *SQLTodoRepo) Update(todo Todo) (*Todo, error) {
	{
		us := sqlUpdateStatementBuildWithValues(SqlPair{
			"title":       todo.title,
			"description": todo.description,
		})
		stmt := fmt.Sprintf(`UPDATE todo SET %s WHERE id=?`, us)
		_, err := s.db.Exec(stmt, todo.uuid)
		if err != nil {
			return nil, log.TError(
				grpc.Errorf(
					codes.Internal,
					"Update Todo error",
				),
				err,
			)
		}
	}
	updated := Todo{}
	{
		row := s.db.QueryRow(`SELECT * FROM todo WHERE id=?`, todo.uuid)
		if err := row.Scan(&updated.uuid, &updated.title, &updated.description, &updated.createdAt); err != nil {
			return nil, log.TError(
				grpc.Errorf(
					codes.Internal,
					"Update Todo error",
				),
				err,
			)
		}
		return &updated, nil
	}
}
func (s *SQLTodoRepo) Delete(uuid string) error {
	{
		if _, err := s.db.Exec(`DELETE FROM todo WHERE id=?`, uuid); err != nil {
			return log.TError(
				grpc.Errorf(
					codes.Internal,
					"Delete Todo error",
				),
				err,
			)
		}
	}
	return nil
}
