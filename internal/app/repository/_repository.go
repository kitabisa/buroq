package repository

type RepositoryOption struct {
	DbMysql *gorp.DbMap
	DbPostgre *gorp.DbMap
	CachePool *redis.Pool
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

// TODO: set function for each repository
// eg
/*
func (r *Repository) SetUserRepository(userRepository IUserRepository) {
	r.User = userRepository
}
*/