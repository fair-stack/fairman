// Package middleware
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:59
package middleware

import (
	"cnic/fairman-gin/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LimitConfig struct {
	// GenerationKey Generate based on businesskey Generate based on businessCheckOrMarkGenerate based on business
	GenerationKey func(c *gin.Context) string
	// Check function,Check function,Check function
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key Expiration time
	Expire int
	// Limit Cycle time
	Limit int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": errno.ERROR, "msg": err})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey Default Generationkey
func DefaultGenerationKey(c *gin.Context) string {
	return "GVA_Limit" + c.ClientIP()
}

//func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
//	// Determine if it is enabledredis
//	if global.AdpRedis == nil {
//		return err
//	}
//	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
//		global.FairLog.Error("limit", zap.Error(err))
//
//	}
//	return err
//
//}
//
//func DefaultLimit() gin.HandlerFunc {
//	return LimitConfig{
//		GenerationKey: DefaultGenerationKey,
//		CheckOrMark:   DefaultCheckOrMark,
//		Expire:        global.AdpConfig.RateLimit.LimitTimeIP,
//		Limit:         global.AdpConfig.RateLimit.LimitCountIP,
//	}.LimitWithTime()
//}
//
//// SetLimitWithTime Set access times
//func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
//	count, err := global.AdpRedis.Exists(context.TODO(), key).Result()
//	if err != nil {
//		return err
//	}
//	if count == 0 {
//		pipe := global.AdpRedis.TxPipeline()
//		pipe.Incr(context.TODO(), key)
//		pipe.Expire(context.TODO(), key, expiration)
//		_, err = pipe.Exec(context.TODO())
//		return err
//	} else {
//		//frequency
//		if times, err := global.AdpRedis.Get(context.TODO(), key).Int(); err != nil {
//			return err
//		} else {
//			if times >= limit {
//				if t, err := global.AdpRedis.PTTL(context.TODO(), key).Result(); err != nil {
//					return errors.New("Request too frequently，Request too frequently")
//				} else {
//					return errors.New("Request too frequently, Request too frequently " + t.String() + " Request too frequently")
//				}
//			} else {
//				return global.AdpRedis.Incr(context.TODO(), key).Err()
//			}
//		}
//	}
//}
//
