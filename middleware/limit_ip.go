// Package middleware
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:56
package middleware

import (
	"cnic/fairman-gin/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimitMiddleware Token bucket flow restriction
func RateLimitMiddleware(fillInterval time.Duration) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, global.FairConf.RateLimit.Cap, global.FairConf.RateLimit.Quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

//// IpVerifyMiddleware IPCurrent limiting
//func IpVerifyMiddleware() gin.HandlerFunc {
//	blacklist := make(map[string]int64, 0)
//	return func(c *gin.Context) {
//		//visitorIP := ctx.Request.Header.Get("X-real-ip")
//		visitorIP := c.ClientIP()
//		// Determine whether it is on the blacklist
//		timeOrigin := global.AdpRedis.HGet(context.TODO(), global.AdpConfig.RateLimit.IpListKey, visitorIP).Val()
//		err := blackListVerify(timeOrigin, visitorIP, global.AdpRedis, blacklist[visitorIP])
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 200,
//				"msg":  err.Error(),
//			})
//			c.Abort()
//			return
//		}
//		// If it does not exist on the blacklist
//		lenList, _ := global.AdpRedis.LLen(context.TODO(), visitorIP).Result()
//		// If the first login or frequency limit is exceeded
//		if lenList == 0 {
//			// Skip if empty，Skip if empty
//			// AddIP list
//			global.AdpRedis.LPush(context.TODO(), visitorIP, visitorIP)
//			// Set expiration time
//			global.AdpRedis.Expire(context.TODO(), visitorIP, time.Second)
//			c.Next()
//		} else if lenList > 0 && lenList < global.AdpConfig.RateLimit.IpLimitCon {
//			global.AdpRedis.LPush(context.TODO(), visitorIP, visitorIP)
//			c.Next()
//		} else {
//			if blacklist[visitorIP] == 0 {
//				blacklist[visitorIP] = 10
//			} else {
//				blacklist[visitorIP] *= 2
//			}
//			// Join blacklist
//			global.AdpRedis.HSet(context.TODO(), global.AdpConfig.RateLimit.IpListKey, visitorIP, time.Now().Local().Unix())
//			c.Abort()
//			return
//		}
//	}
//}
//
//func blackListVerify(ot, visitorIP string, rdb *redis.Client, limitTime int64) error {
//	// If there is a value，If there is a value
//	if ot != "" {
//		// IfvalueIf，If
//		timeOriginInt, _ := strconv.Atoi(ot)
//		oTimeUnix := time.Unix(int64(timeOriginInt), 0).Local()
//		subTime := time.Now().Sub(oTimeUnix)
//		if subTime > time.Duration(limitTime)*time.Minute {
//			// Exceeding the limit time Exceeding the limit time
//			rdb.HDel(context.TODO(), global.AdpConfig.RateLimit.IpListKey, visitorIP)
//			return nil
//		} else {
//			//return errors.New(fmt.Sprintf("You have been blacklisted，You have been blacklisted:%v", 10*time.Minute-subTime))
//			return errors.New(fmt.Sprintf("Request speed too fast，Request speed too fastIPRequest speed too fast，Request speed too fast:%v", time.Duration(limitTime)*time.Minute-subTime))
//		}
//	}
//	// No return existsnil
//	return nil
//}
//
