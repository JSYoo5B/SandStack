package image

import "time"

func (s *Service) CreateMember(imageID string, memberID string) (Member, error) {
	if _, err := s.repository.Get(imageID); err != nil {
		return Member{}, err
	}

	now := s.clock.Now().UTC().Format(time.RFC3339)
	member := Member{
		ImageID:   imageID,
		MemberID:  memberID,
		Status:    "pending",
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.memberRepository.Create(member), nil
}

func (s *Service) ListMembers(imageID string) ([]Member, error) {
	if _, err := s.repository.Get(imageID); err != nil {
		return nil, err
	}

	return s.memberRepository.List(imageID), nil
}

func (s *Service) GetMember(imageID string, memberID string) (Member, error) {
	if _, err := s.repository.Get(imageID); err != nil {
		return Member{}, err
	}

	return s.memberRepository.Get(imageID, memberID)
}

func (s *Service) UpdateMember(
	imageID string,
	memberID string,
	status string,
) (Member, error) {
	member, err := s.GetMember(imageID, memberID)
	if err != nil {
		return Member{}, err
	}

	member.Status = status
	member.UpdatedAt = s.clock.Now().UTC().Format(time.RFC3339)

	return s.memberRepository.Update(member)
}

func (s *Service) DeleteMember(imageID string, memberID string) error {
	if _, err := s.repository.Get(imageID); err != nil {
		return err
	}

	return s.memberRepository.Delete(imageID, memberID)
}
