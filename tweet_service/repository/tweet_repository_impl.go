package repository

import (
	"github.com/rohmanseo/golang-clean-arch/entity"
	"github.com/rohmanseo/golang-clean-arch/repository/permanent"
)

type TweetRepositoryImpl struct {
	permanentDb permanent.IPermanentDb
}

func (t *TweetRepositoryImpl) GetTotalTweet() int {
	return t.permanentDb.GetTotalTweet()
}

func (t *TweetRepositoryImpl) AddTweet(userId int64, tweetContent string) (entity.Tweet, error) {
	return t.permanentDb.AddTweet(userId, tweetContent)
}

func NewTweetRepository(permanent *permanent.IPermanentDb) ITweetRepository {
	return &TweetRepositoryImpl{
		permanentDb: *permanent,
	}
}
