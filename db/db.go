// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.insertCourseStmt, err = db.PrepareContext(ctx, insertCourse); err != nil {
		return nil, fmt.Errorf("error preparing query InsertCourse: %w", err)
	}
	if q.insertDepartmentStmt, err = db.PrepareContext(ctx, insertDepartment); err != nil {
		return nil, fmt.Errorf("error preparing query InsertDepartment: %w", err)
	}
	if q.insertEnrollStmt, err = db.PrepareContext(ctx, insertEnroll); err != nil {
		return nil, fmt.Errorf("error preparing query InsertEnroll: %w", err)
	}
	if q.insertInstructorStmt, err = db.PrepareContext(ctx, insertInstructor); err != nil {
		return nil, fmt.Errorf("error preparing query InsertInstructor: %w", err)
	}
	if q.insertStudentStmt, err = db.PrepareContext(ctx, insertStudent); err != nil {
		return nil, fmt.Errorf("error preparing query InsertStudent: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.insertCourseStmt != nil {
		if cerr := q.insertCourseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertCourseStmt: %w", cerr)
		}
	}
	if q.insertDepartmentStmt != nil {
		if cerr := q.insertDepartmentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertDepartmentStmt: %w", cerr)
		}
	}
	if q.insertEnrollStmt != nil {
		if cerr := q.insertEnrollStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertEnrollStmt: %w", cerr)
		}
	}
	if q.insertInstructorStmt != nil {
		if cerr := q.insertInstructorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertInstructorStmt: %w", cerr)
		}
	}
	if q.insertStudentStmt != nil {
		if cerr := q.insertStudentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertStudentStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	insertCourseStmt     *sql.Stmt
	insertDepartmentStmt *sql.Stmt
	insertEnrollStmt     *sql.Stmt
	insertInstructorStmt *sql.Stmt
	insertStudentStmt    *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		insertCourseStmt:     q.insertCourseStmt,
		insertDepartmentStmt: q.insertDepartmentStmt,
		insertEnrollStmt:     q.insertEnrollStmt,
		insertInstructorStmt: q.insertInstructorStmt,
		insertStudentStmt:    q.insertStudentStmt,
	}
}
