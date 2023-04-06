package service

import (
	"testing"
	"time"

	"infogpt/internal/conf"

	"github.com/stretchr/testify/assert"
)

func TestNewTelegramLimiter(t *testing.T) {
	testCases := []struct {
		Name           string
		TelegramConfig *conf.Telegram
		LimitNil       bool
		UserLimitNil   bool
		ExcludeKeyMap  map[string]struct{}
	}{
		{
			Name:           "1 nil config",
			TelegramConfig: &conf.Telegram{},
			LimitNil:       true,
			UserLimitNil:   true,
			ExcludeKeyMap:  make(map[string]struct{}),
		},
		{
			Name: "2 ratelimit not nil",
			TelegramConfig: &conf.Telegram{
				Ratelimit: "10-H",
			},
			LimitNil:      false,
			UserLimitNil:  true,
			ExcludeKeyMap: make(map[string]struct{}),
		},
		{
			Name: "3 userRatelimit not nil",
			TelegramConfig: &conf.Telegram{
				UserRatelimit: "10-H",
			},
			LimitNil:      true,
			UserLimitNil:  false,
			ExcludeKeyMap: make(map[string]struct{}),
		},
		{
			Name: "4 both ratelimit not nil",
			TelegramConfig: &conf.Telegram{
				Ratelimit:     "10-H",
				UserRatelimit: "10-H",
			},
			LimitNil:      false,
			UserLimitNil:  false,
			ExcludeKeyMap: make(map[string]struct{}),
		},
		{
			Name: "5 exclude keys not nil",
			TelegramConfig: &conf.Telegram{
				ExcludeKeys:   []string{"foo", "bar"},
				Ratelimit:     "10-H",
				UserRatelimit: "10-H",
			},
			LimitNil:     false,
			UserLimitNil: false,
			ExcludeKeyMap: map[string]struct{}{
				"foo": {},
				"bar": {},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			res := NewTelegramLimiter(tc.TelegramConfig)
			assert.NotNil(t, res)

			if tc.LimitNil {
				assert.Nil(t, res.Limiter)
			} else {
				assert.NotNil(t, res.Limiter)
			}

			if tc.UserLimitNil {
				assert.Nil(t, res.UserLimiter)
			} else {
				assert.NotNil(t, res.UserLimiter)
			}

			for k, v := range tc.ExcludeKeyMap {
				assert.Equal(t, v, res.ExcludeKeyMap[k])
			}
		})
	}
}

func TestUserReachedLimit(t *testing.T) {
	// 1. 设置 userLimit 没有excludeKeys
	telConf := &conf.Telegram{
		UserRatelimit: "2-H",
	}
	tLimiter := NewTelegramLimiter(telConf)
	assert.NotNil(t, tLimiter)
	assert.NotNil(t, tLimiter.UserLimiter)

	assert.Equal(t, false, tLimiter.userReachedLimit("foo"))
	assert.Equal(t, false, tLimiter.userReachedLimit("foo"))
	assert.Equal(t, true, tLimiter.userReachedLimit("foo"))

	assert.Equal(t, false, tLimiter.userReachedLimit("bar"))
	assert.Equal(t, false, tLimiter.userReachedLimit("bar"))
	assert.Equal(t, true, tLimiter.userReachedLimit("bar"))
	assert.Equal(t, true, tLimiter.userReachedLimit("bar"))

	// 2. 没有设置 userLimit
	telConf2 := &conf.Telegram{}
	tLimiter2 := NewTelegramLimiter(telConf2)
	assert.NotNil(t, tLimiter2)
	assert.Nil(t, tLimiter2.UserLimiter)

	assert.Equal(t, false, tLimiter2.userReachedLimit("foo"))
	assert.Equal(t, false, tLimiter2.userReachedLimit("foo"))
	assert.Equal(t, false, tLimiter2.userReachedLimit("bar"))
	assert.Equal(t, false, tLimiter2.userReachedLimit("bar"))

	// 3. 设置 exclude keys
	telConf3 := &conf.Telegram{
		UserRatelimit: "2-H",
		ExcludeKeys:   []string{"foo"},
	}
	tLimiter3 := NewTelegramLimiter(telConf3)
	assert.NotNil(t, tLimiter3)
	assert.NotNil(t, tLimiter3.UserLimiter)
	assert.Equal(t, false, tLimiter3.userReachedLimit("foo"))
	assert.Equal(t, false, tLimiter3.userReachedLimit("foo"))
	assert.Equal(t, false, tLimiter3.userReachedLimit("foo"))
	assert.Equal(t, false, tLimiter3.userReachedLimit("foo"))

	assert.Equal(t, false, tLimiter3.userReachedLimit("bar"))
	assert.Equal(t, false, tLimiter3.userReachedLimit("bar"))
	assert.Equal(t, true, tLimiter3.userReachedLimit("bar"))
}

func TestTeleReachedLimit(t *testing.T) {
	// 1. 不限制
	telConf := &conf.Telegram{}
	tLimiter := NewTelegramLimiter(telConf)
	assert.NotNil(t, tLimiter)
	assert.Nil(t, tLimiter.Limiter)
	assert.Equal(t, false, tLimiter.teleReachedLimit())
	assert.Equal(t, false, tLimiter.teleReachedLimit())
	assert.Equal(t, false, tLimiter.teleReachedLimit())
	assert.Equal(t, false, tLimiter.teleReachedLimit())

	// 2. 限制
	telConf2 := &conf.Telegram{
		Ratelimit: "2-H",
	}
	tLimiter2 := NewTelegramLimiter(telConf2)
	assert.NotNil(t, tLimiter2)
	assert.NotNil(t, tLimiter2.Limiter)
	assert.Equal(t, false, tLimiter2.teleReachedLimit())
	assert.Equal(t, false, tLimiter2.teleReachedLimit())
	assert.Equal(t, true, tLimiter2.teleReachedLimit())
}

func TestReachedLimit(t *testing.T) {
	// 1. 不限制
	telConf := &conf.Telegram{}
	tLimiter := NewTelegramLimiter(telConf)
	assert.NotNil(t, tLimiter)
	assert.Nil(t, tLimiter.Limiter)
	assert.Nil(t, tLimiter.UserLimiter)
	assert.EqualValues(t, 0, len(tLimiter.ExcludeKeyMap))
	assert.Equal(t, false, tLimiter.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter.ReachedLimit("foo"))

	// 2. 系统限制
	telConf2 := &conf.Telegram{
		Ratelimit: "2-S",
	}
	tLimiter2 := NewTelegramLimiter(telConf2)
	assert.NotNil(t, tLimiter2)
	assert.NotNil(t, tLimiter2.Limiter)
	assert.Nil(t, tLimiter2.UserLimiter)
	assert.EqualValues(t, 0, len(tLimiter2.ExcludeKeyMap))
	assert.Equal(t, false, tLimiter2.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter2.ReachedLimit("foo"))
	assert.Equal(t, true, tLimiter2.ReachedLimit("foo"))
	// 1s超时后重新开始
	time.Sleep(1100 * time.Millisecond)
	assert.Equal(t, false, tLimiter2.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter2.ReachedLimit("foo"))
	assert.Equal(t, true, tLimiter2.ReachedLimit("foo"))

	// 3. 用户限制
	telConf3 := &conf.Telegram{
		Ratelimit:     "3-S",
		UserRatelimit: "2-S",
	}
	tLimiter3 := NewTelegramLimiter(telConf3)
	assert.NotNil(t, tLimiter3)
	assert.NotNil(t, tLimiter3.Limiter)
	assert.NotNil(t, tLimiter3.UserLimiter)
	assert.EqualValues(t, 0, len(tLimiter3.ExcludeKeyMap))
	assert.Equal(t, false, tLimiter3.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter3.ReachedLimit("foo"))
	assert.Equal(t, true, tLimiter3.ReachedLimit("foo"))

	// 4. exclude keys
	telConf4 := &conf.Telegram{
		Ratelimit:     "3-S",
		UserRatelimit: "2-S",
		ExcludeKeys:   []string{"foo"},
	}
	tLimiter4 := NewTelegramLimiter(telConf4)
	assert.NotNil(t, tLimiter4)
	assert.NotNil(t, tLimiter4.Limiter)
	assert.NotNil(t, tLimiter4.UserLimiter)
	assert.EqualValues(t, 1, len(tLimiter4.ExcludeKeyMap))
	assert.Equal(t, false, tLimiter4.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter4.ReachedLimit("foo"))
	assert.Equal(t, false, tLimiter4.ReachedLimit("foo"))
}
