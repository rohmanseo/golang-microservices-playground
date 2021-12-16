package repository

import (
	"github.com/rohmanseo/golang-clean-arch/entity"
)

type ITweetRepository interface {
	AddTweet(userId int64, tweetContent string) (entity.Tweet, error)
	GetTotalTweet() int
}
