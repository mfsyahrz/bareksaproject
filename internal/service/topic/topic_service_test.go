package service

import "testing"

func TestNewTopic(t *testing.T) {

	t.Run("setup", func(t *testing.T) {
		_ = NewTopicService(nil)
	})
}
