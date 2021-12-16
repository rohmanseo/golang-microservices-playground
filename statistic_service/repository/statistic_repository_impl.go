package repository

import (
	"github.com/nats-io/nats.go"
	"github.com/rohmanseo/golang-clean-arch/repository/memory"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type StatisticRepositoryImpl struct {
	cacheDb memory.ICacheDataStore
	conn    *nats.Conn
}

func (s *StatisticRepositoryImpl) AddUser() {
	s.cacheDb.AddUser()
}

func (s *StatisticRepositoryImpl) GetTotalTweets() (int, error) {
	wg := sync.WaitGroup{}
	var res uint32
	wg.Add(1)

	go func() {
		msg, _ := s.conn.Request("tweet.total", []byte("gimme total tweets"), 300*time.Millisecond)
		temp, _ := strconv.Atoi(string(msg.Data))
		atomic.StoreUint32(&res, uint32(temp))
		wg.Done()
	}()
	wg.Wait()
	return int(res), nil
}

func (s *StatisticRepositoryImpl) GetTotalUsers() (int, error) {
	return s.cacheDb.GetStatistic(), nil
}

func NewStatisticRepository(db *memory.ICacheDataStore, natsConn *nats.Conn) IStatisticRepository {
	return &StatisticRepositoryImpl{cacheDb: *db, conn: natsConn}
}
