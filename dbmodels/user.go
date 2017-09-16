package dbmodels

// User represents the Users who signup / login into the system
type User struct {
	ID        int64
	Username  string
	IsGithub  bool
	CreatedAt int64
	UpdatedAt int64
	Projects  []*Project
}
