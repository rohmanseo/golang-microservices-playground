package permanent

import "github.com/rohmanseo/golang-clean-arch/entity"

type IPermanentDb interface {
	AddTweet(userId int64, tweetContent string) (entity.Tweet, error)
	GetTotalTweet() int
}
