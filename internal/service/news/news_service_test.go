package service

import (
	"testing"
)

func TestNewNews(t *testing.T) {

	t.Run("setup", func(t *testing.T) {
		_ = NewNewsService(nil)
	})
}
