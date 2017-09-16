package dbmodels

// Project represents each project (from Github / Bitbucket / Gitlab / CI) that gets created
type Project struct {
	ID        int64
	URL       string
	Type      string
	CreatedAt int64
	UpdatedAt int64
	UserID    int64
}
