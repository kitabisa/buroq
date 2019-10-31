package service

type ServiceOption struct {
	Repo repository.Repository
	CachePool *redis.Pool
}

type Service struct {}

func NewService() *Service {
	return &Service{}
}

// TODO: set function for each service
// eg
/*
func (s *Service) SetUserService(userService IUserService) {
	s.User = userService
}
*/