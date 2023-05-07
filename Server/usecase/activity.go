package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Risuii/Server/domain"
)

type activityUseCase struct {
	activityRepository domain.ActivityRepository
}

func NewActivityUsecase(activityRepo domain.ActivityRepository) domain.ActivityUseCase {
	return &activityUseCase{
		activityRepository: activityRepo,
	}
}

func (au *activityUseCase) AddActivity(ctx context.Context, data domain.AddActivityReq) (err error) {

	activity := domain.Activity{
		Days:        data.Days,
		Description: data.Description,
		CreatedAt:   time.Now(),
	}

	err = au.activityRepository.Create(ctx, activity)
	if err != nil {
		log.Println(err)
		err = errors.New("errInternalServer")
		return
	}

	return
}

func (au *activityUseCase) ReadOneActivity(ctx context.Context, id int64) (domain.Activity, error) {
	activity, err := au.activityRepository.GetByID(ctx, id)
	if err != nil {
		log.Println(err)
		return domain.Activity{}, errors.New("errNotFound")
	}

	return activity, nil
}

func (au *activityUseCase) ReadAllActivity() ([]domain.Activity, error) {
	activity, err := au.activityRepository.GetAll()
	if err != nil {
		log.Println(err)
		return []domain.Activity{}, err
	}
	return activity, nil
}

func (au *activityUseCase) UpdateActivity(ctx context.Context, data domain.UpdateActivityReq) error {
	activityData, err := au.activityRepository.GetByID(ctx, data.ID)
	if err != nil {
		log.Println(err)
		return errors.New("errNotFound")
	}

	updatedData := domain.Activity{
		Id:          activityData.Id,
		Days:        activityData.Days,
		Description: data.Description,
		UpdatedAt:   time.Now(),
	}

	err = au.activityRepository.Update(ctx, updatedData.Id, updatedData)
	if err != nil {
		log.Println(err)
		switch err.Error() {
		case "errNotFound":
			return errors.New("errNotFound")
		default:
			return err
		}
	}

	return nil
}

func (au *activityUseCase) DeleteActivity(ctx context.Context, id int64) error {

	activity, err := au.activityRepository.GetByID(ctx, id)
	if err != nil {
		log.Println(err)
		switch err.Error() {
		case "errNotFound":
			return errors.New("errNotFound")
		default:
			return err
		}
	}

	if err := au.activityRepository.Delete(ctx, activity.Id); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
