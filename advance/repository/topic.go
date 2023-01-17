package repository

import "sync"

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

// NewTopicInstance 相当于Java中的单例模式
func NewTopicInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
