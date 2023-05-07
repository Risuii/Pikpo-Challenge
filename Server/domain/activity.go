package domain

import (
	"context"
	"time"
)

type Activity struct {
	Id          int64     `json:"id"`
	Days        string    `json:"days"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AddActivityReq struct {
	Days        string    `json:"days" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

type AddactivityRes struct {
	Message string `json:"message"`
}

type UpdateActivityReq struct {
	ID          int64     `json:"id" validate:"required"`
	Days        string    `json:"days"`
	Description string    `json:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateActivityRes struct {
	Message string `json:"message"`
}

type GetActivityReq struct {
	Id int64 `json:"id" validate:"required"`
}

type DeleteActivityReq struct {
	Id int64 `json:"id" validate:"required"`
}

type ActivityUseCase interface {
	AddActivity(ctx context.Context, data AddActivityReq) error
	ReadOneActivity(ctx context.Context, id int64) (Activity, error)
	ReadAllActivity() ([]Activity, error)
	UpdateActivity(ctx context.Context, data UpdateActivityReq) error
	DeleteActivity(ctx context.Context, id int64) error
}

type ActivityRepository interface {
	GetByID(ctx context.Context, id int64) (Activity, error)
	GetAll() ([]Activity, error)
	Create(ctx context.Context, data Activity) error
	Update(ctx context.Context, id int64, data Activity) error
	Delete(ctx context.Context, id int64) error
}
