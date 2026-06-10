package services

import (
	"context"

	profileDTO "github.com/RadithyaR/nihongo-learning-journal/backend/internal/dto/profile"
	repositoryInterfaces "github.com/RadithyaR/nihongo-learning-journal/backend/internal/repositories/interfaces"
	serviceInterface "github.com/RadithyaR/nihongo-learning-journal/backend/internal/services/interfaces"
	"github.com/google/uuid"
)

type profileService struct {
	userRepository repositoryInterfaces.UserRepository
}

func NewProfileService(
	userRepository repositoryInterfaces.UserRepository,
) serviceInterface.ProfileService {
	return &profileService{
		userRepository: userRepository,
	}
}

func (s *profileService) GetProfile(
	ctx context.Context,
	userID uuid.UUID,
) (*profileDTO.ProfileResponse, error) {

	user, err := s.userRepository.FindByID(
		ctx,
		userID,
	)

	if err != nil {
		return nil, err
	}

	return &profileDTO.ProfileResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		AvatarURL:  user.AvatarURL,
		IsVerified: user.IsVerified,
	}, nil
}

func (s *profileService) UpdateProfile(
	ctx context.Context,
	userID uuid.UUID,
	dto profileDTO.UpdateProfileRequest,
) (*profileDTO.ProfileResponse, error) {

	user, err := s.userRepository.FindByID(
		ctx,
		userID,
	)

	if err != nil {
		return nil, err
	}

	user.Name = dto.Name
	user.AvatarURL = dto.AvatarURL

	err = s.userRepository.Update(
		ctx,
		user,
	)

	if err != nil {
		return nil, err
	}

	return &profileDTO.ProfileResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		AvatarURL:  user.AvatarURL,
		IsVerified: user.IsVerified,
	}, nil
}