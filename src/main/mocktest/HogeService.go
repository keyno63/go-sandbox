package mocktest

type User struct {
	ID string
}

type HogeService interface {
	Fuga() (*User, error)
}
type hogeServiceImpl struct {
	UserRepository HogeRepository
}

func NewHogeService(
	userRepository HogeRepository,
) HogeService {
	return &hogeServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *hogeServiceImpl) Fuga() (*User, error) {
	user, _ := s.UserRepository.Get()
	return &User{
		ID: user.ID,
	}, nil
}
