package permanent

import (
	"database/sql"
	"fmt"
	"github.com/rohmanseo/golang-clean-arch/entity"
)

type permanentDbImpl struct {
	pgConnection *sql.DB
}

func (p permanentDbImpl) AddTweet(userId int64, tweetContent string) (entity.Tweet, error) {
	insertQuery := fmt.Sprintf("INSERT INTO \"tweet\"(user_id, tweet_content) VALUES ('%d', '%s') RETURNING id,user_id, tweet_content;", userId, tweetContent)
	tweet := entity.Tweet{
		Id:           0,
		UserId:       0,
		TweetContent: "",
	}
	err := p.pgConnection.QueryRow(insertQuery).Scan(&tweet.Id, &tweet.UserId, &tweet.TweetContent)
	if err != nil {
		fmt.Println(err)
		return tweet, err
	} else {
		return tweet, nil
	}
}

func NewPermanentDb(db *sql.DB) IPermanentDb {
	return &permanentDbImpl{pgConnection: db}
}
