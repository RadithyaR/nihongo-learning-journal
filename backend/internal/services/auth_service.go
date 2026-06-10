package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/config"
	authDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/auth"
	"github.com/RadithyaR/nihongo-learning-journal/backend/internal/models"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
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
	emailVerificationRepository repositoryInterfaces.EmailVerificationRepository
	passwordResetRepository repositoryInterfaces.PasswordResetRepository
}

func NewAuthService(
	userRepository repositoryInterfaces.UserRepository,
	userSessionRepository repositoryInterfaces.UserSessionRepository,
	emailVerificationRepository repositoryInterfaces.EmailVerificationRepository,
	passwordResetRepository repositoryInterfaces.PasswordResetRepository,
) serviceInterfaces.AuthService {
	return &authService{
		userRepository:        userRepository,
		userSessionRepository: userSessionRepository,
		emailVerificationRepository: emailVerificationRepository,
		passwordResetRepository: passwordResetRepository,
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
		IsVerified:   false,
	}

	if err := s.userRepository.Create(
		ctx,
		&user,
	); err != nil {
		return nil, err
	}

	token, err := s.generateVerificationToken()

	if err != nil {
		return nil, err
	}

	verification := models.EmailVerification{
		UserID: user.ID,
		Token:  token,
		ExpiresAt: time.Now().Add(
			24 * time.Hour,
		),
	}

	err = s.emailVerificationRepository.Create(
		ctx,
		&verification,
	)

	if err != nil {
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

	if !user.IsVerified {
		return nil, customErrors.ErrEmailNotVerified
	}

	return s.createLoginSession(
		ctx,
		user,
	)
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
	newTokenID := uuid.New()

	newSecret, err := jwt.GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	newRefreshToken :=
		newTokenID.String() +
		"." +
		newSecret
	
	hashedSecret, err := hash.HashPassword(
	newSecret,
	)

	if err != nil {
		return nil, err
	}
	session.TokenID = newTokenID

	session.RefreshTokenHash = hashedSecret

	session.LastUsedAt = func() *time.Time {
		now := time.Now()
		return &now
	}()
	err = s.userSessionRepository.Update(
		ctx,
		session,
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
		RefreshToken: newRefreshToken,
		User: responses.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			AvatarURL: user.AvatarURL,
		},
	}, nil
}

func (s *authService) Logout(
	ctx context.Context,
	refreshToken string,
) error {

	parts := strings.Split(
		refreshToken,
		".",
	)

	if len(parts) != 2 {
		return customErrors.ErrInvalidRefreshToken
	}

	tokenID, err := uuid.Parse(
		parts[0],
	)

	if err != nil {
		return customErrors.ErrInvalidRefreshToken
	}

	session, err := s.userSessionRepository.
		FindByTokenID(
			ctx,
			tokenID,
		)

	if err != nil {
		return customErrors.ErrSessionNotFound
	}

	err = s.userSessionRepository.Delete(
		ctx,
		session.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) LogoutAll(
	ctx context.Context,
	userID uuid.UUID,
) error {

	return s.userSessionRepository.DeleteByUserID(
		ctx,
		userID,
	)
}


func (s *authService) createLoginSession(
	ctx context.Context,
	user *models.User,
) (*responses.LoginResponse, error){
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

func (s *authService) generateVerificationToken() (
	string,
	error,
) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (s *authService) generatePasswordResetToken() (
	string,
	error,
) {
	return s.generateVerificationToken()
}

func (s *authService) VerifyEmail(
	ctx context.Context,
	token string,
) error {

	verification, err :=
		s.emailVerificationRepository.FindByToken(
			ctx,
			token,
		)

	if err != nil {
		return err
	}

	if verification.VerifiedAt != nil {
		return errors.New(
			"email already verified",
		)
	}

	if time.Now().After(
		verification.ExpiresAt,
	) {
		return errors.New(
			"verification token expired",
		)
	}

	user, err := s.userRepository.FindByID(
		ctx,
		verification.UserID,
	)

	if err != nil {
		return err
	}

	user.IsVerified = true

	err = s.userRepository.Update(
		ctx,
		user,
	)

	if err != nil {
		return err
	}

	now := time.Now()

	verification.VerifiedAt = &now

	err = s.emailVerificationRepository.Update(
		ctx,
		verification,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ForgotPassword(
	ctx context.Context,
	dto authDTO.ForgotPasswordDTO,
) (string, error) {

	user, err := s.userRepository.FindByEmail(
		ctx,
		dto.Email,
	)

	if err != nil {
		return "", err
	}

	err = s.passwordResetRepository.DeleteByUserID(
		ctx,
		user.ID,
	)

	if err != nil {
		return "", err
	}

	token, err := s.generatePasswordResetToken()

	if err != nil {
		return "", err
	}

	reset := models.PasswordReset{
		UserID: user.ID,
		Token: token,
		ExpiresAt: time.Now().Add(
			1 * time.Hour,
		),
	}

	err = s.passwordResetRepository.Create(
		ctx,
		&reset,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) ResetPassword(
	ctx context.Context,
	dto authDTO.ResetPasswordDTO,
) error {

	reset, err := s.passwordResetRepository.
		FindByToken(
			ctx,
			dto.Token,
		)

	if err != nil {
		return err
	}

	if reset.UsedAt != nil {
		return errors.New(
			"reset token already used",
		)
	}

	if time.Now().After(
		reset.ExpiresAt,
	) {
		return errors.New(
			"reset token expired",
		)
	}

	user, err := s.userRepository.FindByID(
		ctx,
		reset.UserID,
	)

	if err != nil {
		return err
	}

	hashedPassword, err := hash.HashPassword(
		dto.NewPassword,
	)

	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword

	err = s.userRepository.Update(
		ctx,
		user,
	)

	if err != nil {
		return err
	}

	now := time.Now()

	reset.UsedAt = &now

	err = s.passwordResetRepository.Update(
		ctx,
		reset,
	)

	if err != nil {
		return err
	}

	err = s.userSessionRepository.DeleteByUserID(
		ctx,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ChangePassword(
	ctx context.Context,
	userID uuid.UUID,
	dto authDTO.ChangePasswordDTO,
) error {

	user, err := s.userRepository.FindByID(
		ctx,
		userID,
	)

	if err != nil {
		return err
	}

	err = hash.ComparePassword(
		user.PasswordHash,
		dto.CurrentPassword,
	)

	if err != nil {
		return customErrors.ErrInvalidCredentials
	}

	hashedPassword, err := hash.HashPassword(
		dto.NewPassword,
	)

	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword

	err = s.userRepository.Update(
		ctx,
		user,
	)

	if err != nil {
		return err
	}

	err = s.userSessionRepository.DeleteByUserID(
		ctx,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}