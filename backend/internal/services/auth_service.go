package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
	authDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/auth"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	"github.com/google/uuid"

	customErrors "github.com/RadithyaR/nihongo-learning-journal/backend/pkg/errors"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/hash"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/jwt"
	"github.com/RadithyaR/nihongo-learning-journal/backend/pkg/responses"

	"gorm.io/gorm"
)

type authService struct {
	userRepository repositoryInterfaces.UserRepository
	userSessionRepository repositoryInterfaces.UserSessionRepository
}

func NewAuthService(
	userRepository repositoryInterfaces.UserRepository,
	userSessionRepository repositoryInterfaces.UserSessionRepository,
) repositoryInterfaces.AuthService {
	return &authService{
		userRepository:        userRepository,
		userSessionRepository: userSessionRepository,
	}
}

func (s *authService) Register(
	ctx context.Context,
	dto authDTO.RegisterDTO,
) (*responses.UserResponse, error) {

	existingUser, err := s.userRepository.FindByEmail(
		ctx,
		dto.Email,
	)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, customErrors.ErrEmailAlreadyExists
	}

	hashedPassword, err := hash.HashPassword(
		dto.Password,
	)

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

func (s *authService) Login(
	ctx context.Context,
	dto authDTO.LoginDTO,
) (*responses.LoginResponse, error) {

	user, err := s.userRepository.FindByEmail(
		ctx,
		dto.Email,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customErrors.ErrInvalidCredentials
		}

		return nil, err
	}

	err = hash.ComparePassword(
		user.PasswordHash,
		dto.Password,
	)

	if err != nil {
		return nil, customErrors.ErrInvalidCredentials
	}

	accessToken, err := jwt.GenerateAccessToken(
		user.ID,
		user.Email,
		config.GetEnv("JWT_SECRET"),
	)

	if err != nil {
		return nil, err
	}

	tokenID := uuid.New()

	refreshSecret, err := jwt.GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	refreshToken := tokenID.String() + "." + refreshSecret

	hashedRefreshToken, err := hash.HashPassword(
		refreshSecret,
	)


	if err != nil {
		return nil, err
	}

	session := models.UserSession{
	TokenID: tokenID,
	UserID: user.ID,
	RefreshTokenHash: hashedRefreshToken,
	ExpiresAt: time.Now().Add(
		7 * 24 * time.Hour,
	),
}

	err = s.userSessionRepository.Create(
		ctx,
		&session,
	)

	if err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: responses.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			AvatarURL: user.AvatarURL,
		},
	}, nil
}

func (s *authService) RefreshToken(
	ctx context.Context,
	refreshToken string,
) (*responses.LoginResponse, error) {

	parts := strings.Split(
		refreshToken,
		".",
	)

	if len(parts) != 2 {
		return nil,
			customErrors.ErrInvalidRefreshToken
	}

	tokenID, err := uuid.Parse(
		parts[0],
	)

	if err != nil {
		return nil,
			customErrors.ErrInvalidRefreshToken
	}

	refreshSecret := parts[1]

	session, err := s.userSessionRepository.
		FindByTokenID(
			ctx,
			tokenID,
		)

	if err != nil {
		return nil,
			customErrors.ErrInvalidRefreshToken
	}

	err = hash.ComparePassword(
		session.RefreshTokenHash,
		refreshSecret,
	)

	if err != nil {
		return nil,
			customErrors.ErrInvalidRefreshToken
	}

	if time.Now().After(
		session.ExpiresAt,
	) {
		return nil,
			customErrors.ErrInvalidRefreshToken
	}

	user, err := s.userRepository.FindByID(
		ctx,
		session.UserID,
	)

	if err != nil {
		return nil, err
	}

	accessToken, err := jwt.GenerateAccessToken(
		user.ID,
		user.Email,
		config.GetEnv("JWT_SECRET"),
	)

	if err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		AccessToken: accessToken,
		User: responses.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			AvatarURL: user.AvatarURL,
		},
	}, nil
}