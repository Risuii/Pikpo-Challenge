package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Risuii/Server/domain"
	"github.com/Risuii/Server/domain/mocks"
	"github.com/Risuii/Server/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddActivity(t *testing.T) {
	t.Run("Add Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.AddActivityReq{
			Days:        "test-days",
			Description: "test-description",
			CreatedAt:   time.Now(),
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Activity")).Return(nil)

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res := activityUseCase.AddActivity(ctx, mockData)
		assert.NoError(t, res)
		activityRepo.AssertExpectations(t)
	})

	t.Run("Add Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.AddActivityReq{
			Days:        "test-days",
			Description: "test-description",
			CreatedAt:   time.Now(),
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Activity")).Return(errors.New("errInternalServer"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res := activityUseCase.AddActivity(ctx, mockData)
		assert.Error(t, res)
		activityRepo.AssertExpectations(t)
	})
}

func TestReadOneActivity(t *testing.T) {
	t.Run("Read One Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.GetActivityReq{
			Id: 1,
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res, err := activityUseCase.ReadOneActivity(ctx, mockData.Id)

		assert.NoError(t, err)
		assert.NotNil(t, res)

		activityRepo.AssertExpectations(t)
	})

	t.Run("Read One Activity Error Not Found", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.GetActivityReq{
			Id: 1,
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("errNotFound"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		_, err := activityUseCase.ReadOneActivity(ctx, mockData.Id)

		assert.Error(t, err)

		activityRepo.AssertExpectations(t)
	})
}

func TestReadAllActivity(t *testing.T) {
	t.Run("Read All Activity Success", func(t *testing.T) {
		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetAll").Return([]domain.Activity{}, nil)

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res, err := activityUseCase.ReadAllActivity()

		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Read All Activity Error", func(t *testing.T) {
		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetAll").Return([]domain.Activity{}, errors.New("error"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		_, err := activityUseCase.ReadAllActivity()

		assert.Error(t, err)

	})
}

func TestUpdateActivity(t *testing.T) {
	t.Run("Update Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.UpdateActivityReq{
			ID:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)
		activityRepo.On("Update", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("domain.Activity")).Return(nil)

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res := activityUseCase.UpdateActivity(ctx, mockData)

		assert.NoError(t, res)

		activityRepo.AssertExpectations(t)
	})

	t.Run("Update Activity Error Not Found Get By ID", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.UpdateActivityReq{
			ID:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("errNotFound"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res := activityUseCase.UpdateActivity(ctx, mockData)

		assert.Error(t, res)

		activityRepo.AssertExpectations(t)
	})

	t.Run("Update Activity Error Not Found Update", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.UpdateActivityReq{
			ID:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)
		activityRepo.On("Update", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("domain.Activity")).Return(errors.New("errNotFound"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res := activityUseCase.UpdateActivity(ctx, mockData)

		assert.Error(t, res)

		activityRepo.AssertExpectations(t)
	})

	t.Run("Update Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		mockData := domain.UpdateActivityReq{
			ID:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)
		activityRepo.On("Update", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("domain.Activity")).Return(errors.New("errInternalServer"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		res := activityUseCase.UpdateActivity(ctx, mockData)

		assert.Error(t, res)

		activityRepo.AssertExpectations(t)
	})
}

func TestDeleteActivity(t *testing.T) {
	t.Run("Delete Activity Success", func(t *testing.T) {
		ctx := context.Background()

		mockData := domain.DeleteActivityReq{
			Id: 1,
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)
		activityRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		err := activityUseCase.DeleteActivity(ctx, mockData.Id)

		assert.NoError(t, err)
		activityRepo.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Not Found GetByID", func(t *testing.T) {
		ctx := context.Background()

		mockData := domain.DeleteActivityReq{
			Id: 1,
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("errNotFound"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		err := activityUseCase.DeleteActivity(ctx, mockData.Id)

		assert.Error(t, err)
		activityRepo.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Internal Server GetByID", func(t *testing.T) {
		ctx := context.Background()

		mockData := domain.DeleteActivityReq{
			Id: 1,
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("errInternalServer"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		err := activityUseCase.DeleteActivity(ctx, mockData.Id)

		assert.Error(t, err)
		activityRepo.AssertExpectations(t)
	})

	t.Run("Delete Activity Error Internal Server Delete", func(t *testing.T) {
		ctx := context.Background()

		mockData := domain.DeleteActivityReq{
			Id: 1,
		}

		activityRepo := &mocks.ActivityRepository{}
		activityRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)
		activityRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(errors.New("errInternalServer"))

		activityUseCase := usecase.NewActivityUsecase(activityRepo)

		err := activityUseCase.DeleteActivity(ctx, mockData.Id)

		assert.Error(t, err)
		activityRepo.AssertExpectations(t)
	})
}
