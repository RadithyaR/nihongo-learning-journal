package services

import (
	"context"
	"errors"

	authDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/auth"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	authErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	responses "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"

	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/hash"

	"gorm.io/gorm"
)

type authService struct {
	userRepository repositoryInterfaces.UserRepository
}

func NewAuthService(
	userRepository repositoryInterfaces.UserRepository,
) repositoryInterfaces.AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) Register(
	ctx context.Context,
	dto authDTO.RegisterDTO,
) (*responses.UserResponse, error) {

	existingUser, err := s.userRepository.
		FindByEmail(ctx, dto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, authErrors.ErrEmailAlreadyExists
	}

	hashedPassword, err := hash.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: hashedPassword,
	}

	if err := s.userRepository.Create(
		ctx,
		&user,
	); err != nil {
		return nil, err
	}

	return &responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}, nil
}