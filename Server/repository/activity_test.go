package repository_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Risuii/Server/domain"
	"github.com/Risuii/Server/domain/mocks"
	"github.com/Risuii/Server/helpers/constant"
	"github.com/Risuii/Server/repository"
	"github.com/stretchr/testify/assert"
)

var activityStruct = domain.Activity{
	Id:          1,
	Days:        "test-days",
	Description: "test-description",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

func TestCreate(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		db, mock := mocks.NewMock()
		ctx := context.TODO()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("INSERT INTO %s", constant.TableActivity)

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.Days, activityStruct.Description, activityStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(ctx, activityStruct)

		assert.NoError(t, err)
	})

	t.Run("Create Error Query", func(t *testing.T) {
		db, mock := mocks.NewMock()
		ctx := context.TODO()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("INSERT %s", constant.TableActivity)

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.Days, activityStruct.Description, activityStruct.CreatedAt).WillReturnResult(sqlmock.NewErrorResult(errors.New("errInternalServer")))

		err := repo.Create(ctx, activityStruct)

		assert.Error(t, err)
	})

	t.Run("Create Error Exec", func(t *testing.T) {
		db, mock := mocks.NewMock()
		ctx := context.TODO()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("INSERT INTO %s", constant.TableActivity)

		mock.ExpectPrepare(query).ExpectExec().WithArgs(nil).WillReturnResult(sqlmock.NewErrorResult(errors.New("errInternalServer")))

		err := repo.Create(ctx, activityStruct)

		assert.Error(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID Success", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("SELECT id, days, description, created_at, update_at FROM %s WHERE id = %d", constant.TableActivity, activityStruct.Id)
		rows := sqlmock.NewRows([]string{"id", "days", "description", "created_at", "update_at"}).AddRow(activityStruct.Id, activityStruct.Days, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdatedAt)

		mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

		activityStruct, err := repo.GetByID(ctx, activityStruct.Id)

		assert.NotNil(t, activityStruct)
		assert.NoError(t, err)
	})

	t.Run("Get By ID Error Query", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("SELECT FROM %s WHERE id = %d", constant.TableActivity, activityStruct.Id)
		rows := sqlmock.NewRows([]string{"id", "days", "description", "created_at", "update_at"}).AddRow(activityStruct.Id, activityStruct.Days, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdatedAt)

		mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

		_, err := repo.GetByID(ctx, activityStruct.Id)

		assert.Error(t, err)
	})

	t.Run("Get By ID Error Scan", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()
		expectedError := errors.New("errNotFound")

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("SELECT id, days, description, created_at, update_at FROM %s WHERE id = %d", constant.TableActivity, activityStruct.Id)
		rows := sqlmock.NewRows([]string{"id", "days", "description", "created_at", "update_at"})

		mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

		_, err := repo.GetByID(ctx, activityStruct.Id)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedError.Error())
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Get All Success", func(t *testing.T) {
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		query := fmt.Sprintf("SELECT id, days, description, created_at, update_at FROM %s", constant.TableActivity)
		rows := sqlmock.NewRows([]string{"id", "days", "description", "created_at", "update_at"}).AddRow(activityStruct.Id, activityStruct.Days, activityStruct.Description, activityStruct.CreatedAt, activityStruct.UpdatedAt)

		mock.ExpectQuery(query).WillReturnRows(rows)

		activityStruct, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, activityStruct)
	})

	t.Run("Get All Error Query", func(t *testing.T) {
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		query := fmt.Sprintf("SELECT  FROM %s", constant.TableActivity)
		rows := sqlmock.NewRows([]string{"id", "days", "description", "created_at", "update_at"})

		mock.ExpectQuery(query).WillReturnRows(rows)

		_, err := repo.GetAll()

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("UPDATE %s SET", constant.TableActivity)

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.Days, activityStruct.Description, activityStruct.UpdatedAt).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(ctx, activityStruct.Id, activityStruct)

		assert.NoError(t, err)
	})

	t.Run("Update Error Query", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("SELECT %s", constant.TableActivity)

		mock.ExpectPrepare(query)

		err := repo.Update(ctx, activityStruct.Id, activityStruct)

		assert.Error(t, err)
	})

	t.Run("Update Error Exec", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("UPDATE %s SET", constant.TableActivity)

		mock.ExpectPrepare(query)

		err := repo.Update(ctx, activityStruct.Id, activityStruct)

		assert.Error(t, err)
	})

	t.Run("Update Error Not Found", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("UPDATE %s SET", constant.TableActivity)

		mock.ExpectPrepare(query).ExpectExec().WithArgs(activityStruct.Days, activityStruct.Description, activityStruct.UpdatedAt).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Update(ctx, activityStruct.Id, activityStruct)

		assert.Equal(t, errors.New("errNotFound"), err)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete Success", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("DELETE FROM %s WHERE id = %d", constant.TableActivity, activityStruct.Id)

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(ctx, activityStruct.Id)

		assert.NoError(t, err)
	})

	t.Run("Delete Error Query", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("DELETE FROM %s id = %d", constant.TableActivity, activityStruct.Id)

		mock.ExpectPrepare(query)

		err := repo.Delete(ctx, activityStruct.Id)

		assert.Error(t, err)
	})

	t.Run("Delete Error Not Found", func(t *testing.T) {
		ctx := context.TODO()
		db, mock := mocks.NewMock()

		repo := repository.NewActivityRepository(db, constant.TableActivity)

		defer db.Close()

		query := fmt.Sprintf("DELETE FROM %s WHERE id = %d", constant.TableActivity, activityStruct.Id)

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Delete(ctx, activityStruct.Id)

		assert.Equal(t, errors.New("errNotFound"), err)
		assert.Error(t, err)
	})
}
