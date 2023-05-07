package handler

import (
	"context"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Risuii/Server/domain"
	"github.com/Risuii/Server/helpers/exception"
	"github.com/Risuii/Server/helpers/response"
	pb "github.com/Risuii/proto"
)

type activityService struct {
	usecase domain.ActivityUseCase
	pb.UnimplementedActivityServiceServer
}

func NewActivityService(usecase domain.ActivityUseCase) pb.ActivityServiceServer {
	return &activityService{
		usecase: usecase,
	}
}

func (as *activityService) AddActivity(ctx context.Context, req *pb.AddActivityReq) (*pb.AddActivityRes, error) {
	// get data from payload
	data := domain.AddActivityReq{
		Days:        req.Days,
		Description: req.Description,
	}

	// validate data
	validate := validator.New()
	err := validate.Struct(&data)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, exception.ErrBadRequest)
	}

	// send data to usecase
	err = as.usecase.AddActivity(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, exception.ErrInternalServer)
	}

	res := pb.AddActivityRes{
		Message: response.Success,
	}

	return &res, nil
}

func (as *activityService) GetActivity(ctx context.Context, req *pb.GetActivityReq) (*pb.GetActivityResp, error) {
	// get data from payload
	data := domain.GetActivityReq{
		Id: req.Id,
	}

	// validate data
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, exception.ErrBadRequest)
	}

	// send data to usecase and get return data
	activityData, err := as.usecase.ReadOneActivity(ctx, data.Id)
	if err != nil {
		log.Println(err)
		switch err.Error() {
		case "errNotFound":
			return nil, status.Error(codes.NotFound, exception.ErrNotFound)
		default:
			return nil, status.Error(codes.Internal, exception.ErrInternalServer)
		}
	}

	// use return data for response
	res := pb.GetActivityResp{
		Activity: &pb.Activity{
			Id:          activityData.Id,
			Days:        activityData.Days,
			Description: activityData.Description,
			CreatedAt:   activityData.CreatedAt.Format(time.RFC3339),
			UpdateAt:    activityData.CreatedAt.Format(time.RFC3339),
		},
	}

	return &res, nil
}

func (as *activityService) GetAllActivity(context.Context, *pb.GetAllReq) (*pb.GetAllResp, error) {
	// call usecase to get data
	activityData, err := as.usecase.ReadAllActivity()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, exception.ErrInternalServer)
	}

	// looping data for response
	res := pb.GetAllResp{}

	for _, val := range activityData {
		res.Activity = append(res.Activity, &pb.Activity{
			Id:          val.Id,
			Days:        val.Days,
			Description: val.Description,
			CreatedAt:   val.CreatedAt.Format(time.RFC3339),
			UpdateAt:    val.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &res, nil
}

func (as *activityService) UpdateActivity(ctx context.Context, req *pb.UpdateActivityReq) (*pb.UpdateActivityResp, error) {
	// get data from payload
	data := domain.UpdateActivityReq{
		ID:          req.Id,
		Days:        req.Days,
		Description: req.Description,
	}

	// validate data
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, exception.ErrBadRequest)
	}

	// send data to usecase
	err := as.usecase.UpdateActivity(ctx, data)
	if err != nil {
		switch err.Error() {
		case "errNotFound":
			return nil, status.Error(codes.NotFound, exception.ErrNotFound)
		default:
			return nil, status.Error(codes.Internal, exception.ErrInternalServer)
		}
	}

	// response
	res := pb.UpdateActivityResp{
		Message: response.Success,
	}

	return &res, nil
}

func (as *activityService) DeleteActivity(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteResp, error) {
	// get data from payload
	data := domain.DeleteActivityReq{
		Id: req.Id,
	}

	// validate data
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		log.Println(err)
		return nil, status.Error(codes.InvalidArgument, exception.ErrBadRequest)
	}

	// send data to usecase
	if err := as.usecase.DeleteActivity(ctx, data.Id); err != nil {
		log.Println(err)
		switch err.Error() {
		case "errNotFound":
			return nil, status.Error(codes.NotFound, exception.ErrNotFound)
		default:
			return nil, status.Error(codes.Internal, exception.ErrInternalServer)
		}
	}

	// response
	res := pb.DeleteResp{
		Message: response.Success,
	}

	return &res, nil
}
