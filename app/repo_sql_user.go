package app

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type SQLUserRepo struct {
	db *sql.DB
}

func (r *SQLUserRepo) Create(username string, email string, hashedPassword string) (*User, error) {
	{
		user, err := r.GetUsername(username)
		if err != nil {
			return nil, err
		}
		if user != nil {
			return nil, log.TError(grpc.Errorf(codes.AlreadyExists, "user is already exists"), err)
		}
	}
	{
		cs, ps, vs := sqlStatementBuildWithValues(SqlPair{
			"uuid":     uuid.New(),
			"username": username,
			"email":    email,
			"password": hashedPassword,
		})
		stmt := fmt.Sprintf(`INSERT INTO user %s VALUES %s`, cs, ps)
		_, err := r.db.Exec(stmt, vs...)
		if err != nil {
			return nil, log.TError(grpc.Errorf(codes.Internal, "Create User error"), err)
		}
	}
	user := &User{}
	{
		row := r.db.QueryRow(`SELECT * FROM user WHERE username=?`, username)
		err := row.Scan(&user.uuid, &user.username, &user.email, &user.hashedPassword)
		if err != nil {
			return nil, log.TError(grpc.Errorf(codes.Internal, "Create User error"), err)
		}
	}
	return user, nil
}
func (r *SQLUserRepo) GetUsername(username string) (*User, error) {
	user := &User{}
	{
		row := r.db.QueryRow(`SELECT * FROM user WHERE username=?`, username)
		err := row.Scan(&user.uuid, &user.email, &user.username, &user.hashedPassword)
		if err != nil {
			return nil, log.TError(grpc.Errorf(codes.Internal, "Get User error"), err)
		}
	}
	return user, nil
}
