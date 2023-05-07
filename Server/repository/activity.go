package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Risuii/Server/domain"
)

type activityRepository struct {
	Conn      *sql.DB
	TableName string
}

func NewActivityRepository(conn *sql.DB, tableName string) domain.ActivityRepository {
	return &activityRepository{
		Conn:      conn,
		TableName: tableName,
	}
}

func (ar *activityRepository) Create(ctx context.Context, data domain.Activity) (err error) {
	query := fmt.Sprintf("INSERT INTO %s (days, description, created_at) VALUES ($1, $2, $3)", ar.TableName)
	stmt, err := ar.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		data.Days,
		data.Description,
		data.CreatedAt,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	return
}

func (ar *activityRepository) GetByID(ctx context.Context, id int64) (domain.Activity, error) {
	activity := domain.Activity{}

	query := fmt.Sprintf("SELECT id, days, description, created_at, update_at FROM %s WHERE id = %d", ar.TableName, id)
	stmt, err := ar.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return domain.Activity{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx)

	err = row.Scan(
		&activity.Id,
		&activity.Days,
		&activity.Description,
		&activity.CreatedAt,
		&activity.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return domain.Activity{}, errors.New("errNotFound")
	}

	return activity, nil
}

func (ar *activityRepository) GetAll() ([]domain.Activity, error) {
	activity := []domain.Activity{}

	rows, err := ar.Conn.Query(fmt.Sprintf("SELECT id, days, description, created_at, update_at FROM %s", ar.TableName))
	if err != nil {
		log.Println(err)
		return activity, err
	}

	defer rows.Close()

	for rows.Next() {
		var c domain.Activity
		rows.Scan(
			&c.Id,
			&c.Days,
			&c.Description,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		activity = append(activity, c)
	}

	return activity, nil
}

func (ar *activityRepository) Update(ctx context.Context, id int64, data domain.Activity) error {
	query := fmt.Sprintf("UPDATE %s SET days = $1, description = $2, update_at = $3 WHERE id = %d", ar.TableName, id)
	stmt, err := ar.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		data.Days,
		data.Description,
		data.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return errors.New("errInternalServer")
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected < 1 {
		return errors.New("errNotFound")
	}

	return nil
}

func (ar *activityRepository) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d", ar.TableName, id)
	stmt, err := ar.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	result, _ := stmt.ExecContext(
		ctx,
	)

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return errors.New("errNotFound")
	}

	return nil
}
