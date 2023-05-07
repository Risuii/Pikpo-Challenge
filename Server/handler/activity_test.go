package handler_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Risuii/Server/domain"
	"github.com/Risuii/Server/domain/mocks"
	"github.com/Risuii/Server/handler"
	"github.com/Risuii/Server/helpers/exception"
	"github.com/Risuii/Server/helpers/response"
	pb "github.com/Risuii/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddActivity(t *testing.T) {
	t.Run("Add Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.AddActivityReq{
			Days:        "test-days",
			Description: "test-description",
		}

		expectedRes := &pb.AddActivityRes{
			Message: response.Success,
		}

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("AddActivity", mock.Anything, mock.AnythingOfType("domain.AddActivityReq")).Return(nil)

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.AddActivity(ctx, &mockData)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("Add Activity Error Validation", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.AddActivityReq{
			Days:        "",
			Description: "",
		}

		expectedRes := status.Error(codes.InvalidArgument, exception.ErrBadRequest)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("AddActivity", mock.Anything, mock.AnythingOfType("domain.AddActivityReq")).Return(nil)

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.AddActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Add Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.AddActivityReq{
			Days:        "test-days",
			Description: "test-description",
		}

		expectedRes := status.Error(codes.Internal, exception.ErrInternalServer)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("AddActivity", mock.Anything, mock.AnythingOfType("domain.AddActivityReq")).Return(errors.New("errInternalServer"))

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.AddActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})
}

func TestGetActivity(t *testing.T) {
	t.Run("Get Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := &pb.GetActivityReq{
			Id: 1,
		}

		expectedRes := &pb.GetActivityResp{
			Activity: &pb.Activity{
				Id:          1,
				Days:        "test-days",
				Description: "test-description",
				CreatedAt:   "0001-01-01T00:00:00Z",
				UpdateAt:    "0001-01-01T00:00:00Z",
			},
		}

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("ReadOneActivity", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{
			Id:          1,
			Days:        "test-days",
			Description: "test-description",
		}, nil)

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.GetActivity(ctx, mockData)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("Get Activity Error Validation", func(t *testing.T) {
		ctx := context.Background()
		mockData := &pb.GetActivityReq{
			Id: 0,
		}

		expectedRes := status.Error(codes.InvalidArgument, exception.ErrBadRequest)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("ReadOneActivity", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil)

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.GetActivity(ctx, mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Get Activity Error Not Found", func(t *testing.T) {
		ctx := context.Background()
		mockData := &pb.GetActivityReq{
			Id: 1,
		}

		expectedRes := status.Error(codes.NotFound, exception.ErrNotFound)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("ReadOneActivity", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("errNotFound"))

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.GetActivity(ctx, mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Get Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		mockData := &pb.GetActivityReq{
			Id: 1,
		}

		expectedRes := status.Error(codes.Internal, exception.ErrInternalServer)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("ReadOneActivity", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("errInternalServer"))

		handler := handler.NewActivityService(
			activityUseCase,
		)

		res, err := handler.GetActivity(ctx, mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})
}

func TestGetAllActivity(t *testing.T) {
	t.Run("Get All Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := []domain.Activity{}
		expectedRes := &pb.GetAllResp{}

		for _, val := range mockData {
			expectedRes.Activity = append(expectedRes.Activity, &pb.Activity{
				Id:          val.Id,
				Days:        val.Days,
				Description: val.Description,
				CreatedAt:   val.CreatedAt.Format(time.RFC3339),
				UpdateAt:    val.UpdatedAt.Format(time.RFC3339),
			})
		}

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("ReadAllActivity").Return([]domain.Activity{}, nil)

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.GetAllActivity(ctx, &pb.GetAllReq{})

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("Get All Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		expectedRes := status.Error(codes.Internal, exception.ErrInternalServer)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("ReadAllActivity").Return([]domain.Activity{}, errors.New("errInternalServer"))

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.GetAllActivity(ctx, &pb.GetAllReq{})

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})
}

func TestUpdateActivity(t *testing.T) {
	t.Run("Update Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.UpdateActivityReq{
			Id:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		expectedRes := &pb.UpdateActivityResp{
			Message: response.Success,
		}

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("domain.UpdateActivityReq")).Return(nil)

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.UpdateActivity(ctx, &mockData)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("Update Activity Error Validation", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.UpdateActivityReq{
			Id:          0,
			Days:        "",
			Description: "",
		}

		expectedRes := status.Error(codes.InvalidArgument, exception.ErrBadRequest)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("domain.UpdateActivityReq")).Return(nil)

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.UpdateActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Update Activity Error Not Found", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.UpdateActivityReq{
			Id:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		expectedRes := status.Error(codes.NotFound, exception.ErrNotFound)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("domain.UpdateActivityReq")).Return(errors.New("errNotFound"))

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.UpdateActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Update Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.UpdateActivityReq{
			Id:          1,
			Days:        "test-days",
			Description: "test-description",
		}

		expectedRes := status.Error(codes.Internal, exception.ErrInternalServer)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("UpdateActivity", mock.Anything, mock.AnythingOfType("domain.UpdateActivityReq")).Return(errors.New("errInternalServer"))

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.UpdateActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})
}

func TestDeleteActivity(t *testing.T) {
	t.Run("Delete Activity Success", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.DeleteReq{
			Id: 1,
		}

		expectedRes := &pb.DeleteResp{
			Message: response.Success,
		}

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("DeleteActivity", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.DeleteActivity(ctx, &mockData)

		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("Delete Activity Error Validation", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.DeleteReq{
			Id: 0,
		}

		expectedRes := status.Error(codes.InvalidArgument, exception.ErrBadRequest)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("DeleteActivity", mock.Anything, mock.AnythingOfType("int64")).Return(nil)

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.DeleteActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Delete Activity Error Not Found", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.DeleteReq{
			Id: 1,
		}

		expectedRes := status.Error(codes.NotFound, exception.ErrNotFound)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("DeleteActivity", mock.Anything, mock.AnythingOfType("int64")).Return(errors.New("errNotFound"))

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.DeleteActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})

	t.Run("Delete Activity Error Internal Server", func(t *testing.T) {
		ctx := context.Background()
		mockData := pb.DeleteReq{
			Id: 1,
		}

		expectedRes := status.Error(codes.Internal, exception.ErrInternalServer)

		activityUseCase := &mocks.ActivityUseCase{}
		activityUseCase.On("DeleteActivity", mock.Anything, mock.AnythingOfType("int64")).Return(errors.New("errInternalServer"))

		handler := handler.NewActivityService(activityUseCase)

		res, err := handler.DeleteActivity(ctx, &mockData)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedRes.Error())
		assert.Nil(t, res)
	})
}
