package repository

type IStatisticRepository interface {
	GetTotalUsers() (int, error)
	GetTotalTweets() (int, error)
	AddUser()
}
