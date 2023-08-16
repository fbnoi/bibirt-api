package data

import (
	"bibirt-api/internal/biz"
	"bibirt-api/internal/data/ent/user"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) error {
	ux, err := r.data.db.User.Create().
		SetUUID(u.Uuid).
		SetType(u.Type).
		SetName(u.Name).
		SetEmail(u.Email.String).
		SetPhone(u.Phone.String).
		SetStatus(u.Status).
		SetCreatedAt(u.CreatedAt).
		Save(ctx)
	u.Id = int64(ux.ID)
	return err
}

func (r *userRepo) FindByUUID(ctx context.Context, uuid string, dist *biz.User) error {
	ux, err := r.data.db.User.Query().Where(user.UUIDEQ(uuid)).First(ctx)
	if err != nil {
		return err
	}
	dist.Id = int64(ux.ID)
	dist.Uuid = ux.UUID
	dist.Type = ux.Type
	dist.Name = ux.Name
	dist.Email.Scan(ux.Email)
	dist.Phone.Scan(ux.Phone)
	dist.Status = ux.Status
	dist.CreatedAt = ux.CreatedAt
	return nil
}
