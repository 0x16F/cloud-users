package users

import (
	"context"
	"errors"

	"github.com/0x16F/cloud-common/pkg/generator"
	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/entity"
	"github.com/0x16F/cloud-users/pkg/codes"
	"github.com/jackc/pgx/v5"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUser(ctx context.Context, id uint64) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
	GetUsers(ctx context.Context, params entity.GetUsersParams) ([]entity.User, error)
	UpdateEmail(ctx context.Context, id uint64, email string) error
	UpdateUsername(ctx context.Context, id uint64, username string) error
	UpdatePassword(ctx context.Context, id uint64, password string, salt string) error
	DeleteUser(ctx context.Context, id uint64) error
}

type ErrorsService interface {
	GetError(code int) error
}

type Service struct {
	log           logger.Logger
	usersRepo     UsersRepository
	errorsService ErrorsService
}

func New(log logger.Logger, usersRepo UsersRepository, errorsService ErrorsService) *Service {
	return &Service{
		log:           log,
		usersRepo:     usersRepo,
		errorsService: errorsService,
	}
}

func (s *Service) CreateUser(ctx context.Context, dto entity.UserCreateDTO) (entity.User, error) {
	log := s.log.WithFields(logger.Fields{
		"method": "CreateUser",
	})

	user, err := s.usersRepo.GetUserByEmail(ctx, dto.Email)
	if err != nil && errors.Is(err, s.errorsService.GetError(codes.InternalError)) {
		log.Errorf("failed to get user by email: %v", err)

		return entity.User{}, err
	}

	if user.ID != 0 {
		return entity.User{}, s.errorsService.GetError(codes.EmailAlreadyExists)
	}

	user, err = s.usersRepo.GetUserByUsername(ctx, dto.Username)
	if err != nil && errors.Is(err, s.errorsService.GetError(codes.InternalError)) {
		log.Errorf("failed to get user by username: %v", err)

		return entity.User{}, err
	}

	if user.ID != 0 {
		return entity.User{}, s.errorsService.GetError(codes.UsernameAlreadyExists)
	}

	user, err = s.usersRepo.CreateUser(ctx, entity.NewUser(dto))
	if err != nil {
		log.Errorf("failed to create user: %v", err)

		return entity.User{}, s.errorsService.GetError(codes.InternalError)
	}

	return user, nil

}

func (s *Service) GetUser(ctx context.Context, id uint64) (entity.User, error) {
	log := s.log.WithFields(logger.Fields{
		"method": "GetUser",
	})

	user, err := s.usersRepo.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, s.errorsService.GetError(codes.UserNotFound)
		}

		log.Errorf("failed to get user: %v", err)

		return entity.User{}, s.errorsService.GetError(codes.InternalError)
	}

	return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	log := s.log.WithFields(logger.Fields{
		"method": "GetUserByEmail",
	})

	user, err := s.usersRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, s.errorsService.GetError(codes.UserNotFound)
		}

		log.Errorf("failed to get user by email: %v", err)

		return entity.User{}, s.errorsService.GetError(codes.InternalError)
	}

	return user, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	log := s.log.WithFields(logger.Fields{
		"method": "GetUserByUsername",
	})

	user, err := s.usersRepo.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, s.errorsService.GetError(codes.UserNotFound)
		}

		log.Errorf("failed to get user by username: %v", err)

		return entity.User{}, s.errorsService.GetError(codes.InternalError)
	}

	return user, nil
}

func (s *Service) GetUsers(ctx context.Context, params entity.GetUsersParams) ([]entity.User, error) {
	log := s.log.WithFields(logger.Fields{
		"method": "GetUsers",
	})

	users, err := s.usersRepo.GetUsers(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []entity.User{}, nil
		}

		log.Errorf("failed to get users: %v", err)

		return nil, s.errorsService.GetError(codes.InternalError)
	}

	return users, nil
}

func (s *Service) UpdateEmail(ctx context.Context, id uint64, email string) error {
	log := s.log.WithFields(logger.Fields{
		"method": "UpdateEmail",
	})

	user, err := s.GetUserByEmail(ctx, email)
	if err != nil && errors.Is(err, s.errorsService.GetError(codes.InternalError)) {
		log.Errorf("failed to get user by email: %v", err)

		return err
	}

	if user.ID != 0 {
		return s.errorsService.GetError(codes.EmailAlreadyExists)
	}

	if _, err = s.GetUser(ctx, id); err != nil {
		log.Errorf("failed to get user: %v", err)

		return err
	}

	if err = s.usersRepo.UpdateEmail(ctx, id, email); err != nil {
		log.Errorf("failed to update email: %v", err)

		return s.errorsService.GetError(codes.InternalError)
	}

	return nil
}

func (s *Service) UpdateUsername(ctx context.Context, id uint64, username string) error {
	log := s.log.WithFields(logger.Fields{
		"method": "UpdateUsername",
	})

	user, err := s.usersRepo.GetUserByUsername(ctx, username)
	if err != nil && errors.Is(err, s.errorsService.GetError(codes.InternalError)) {
		log.Errorf("failed to get user by username: %v", err)

		return err
	}

	if user.ID != 0 {
		return s.errorsService.GetError(codes.UsernameAlreadyExists)
	}

	if _, err = s.GetUser(ctx, id); err != nil {
		log.Errorf("failed to get user: %v", err)

		return err
	}

	if err = s.usersRepo.UpdateUsername(ctx, id, username); err != nil {
		log.Errorf("failed to update username: %v", err)

		return s.errorsService.GetError(codes.InternalError)
	}

	return nil
}

func (s *Service) UpdatePassword(ctx context.Context, id uint64, oldPassword string, newPassword string) error {
	log := s.log.WithFields(logger.Fields{
		"method": "UpdatePassword",
	})

	user, err := s.GetUser(ctx, id)
	if err != nil {
		log.Errorf("failed to get user: %v", err)

		return err
	}

	if ok := user.ValidatePassword(oldPassword); !ok {
		return s.errorsService.GetError(codes.InvalidOldPassword)
	}

	salt := generator.NewString(entity.SaltLength)
	hash := generator.NewHash(newPassword, salt)

	if err = s.usersRepo.UpdatePassword(ctx, id, hash, salt); err != nil {
		log.Errorf("failed to update password: %v", err)

		return s.errorsService.GetError(codes.InternalError)
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id uint64) error {
	log := s.log.WithFields(logger.Fields{
		"method": "DeleteUser",
	})

	if err := s.usersRepo.DeleteUser(ctx, id); err != nil {
		log.Errorf("failed to delete user: %v", err)

		return s.errorsService.GetError(codes.InternalError)
	}

	return nil
}
