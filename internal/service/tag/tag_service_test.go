package service

import "testing"

func TestNewTag(t *testing.T) {

	t.Run("setup", func(t *testing.T) {
		_ = NewTagService(nil)
	})
}
