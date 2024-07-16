package router

import (
	"eCommerce/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type Router struct {
	config *config.Config

	engin *gin.Engine
}

func NewRouter(config *config.Config) (*Router, error) {
	r := &Router{
		config: config,
		engin:  gin.New(),
	}

	/** 서버 실행 코드 **/
	r.engin.Use(requestTimeOutMiddleWare(5 * time.Second)) // 미들웨어

	NewMongoRouter(r)

	return r, r.engin.Run(config.ServerInfo.Port)
}

// 수동으로 타임아웃 설정
func requestTimeOutMiddleWare(timeout time.Duration) gin.HandlerFunc {
	// context -> 요청의 생명 주기를 관리하기 위해 사용되는 인터페이스
	return func(c *gin.Context) {
		// 새로운 context를 생성하여 주어진 timeout을 설정
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		// 이 함수가 끝날 때 context 취소를 보장
		defer cancel()

		// 새 context로 요청을 설정
		c.Request = c.Request.WithContext(ctx)
		// 요청 처리가 완료되었음을 알리는 채널을 생성
		done := make(chan struct{})

		// 고루틴(동시성)
		go func() {
			// 이 고루틴은 요청을 처리하는데 사용될 수 있다.
			defer close(done)
			c.Next() // 다음 미들웨어 또는 핸들러로 제어를 넘긴다.
		}()

		// select 문을 사용하여 done 채널과 context의 완료 상태를 대기
		select {
		case <-done:
			return // 요청 처리가 완료되면 미들웨어 종료
		case <-ctx.Done():
			// 타임아웃이 발생하면 요청을 중단하고 504 상태 코드를 반환
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{"error": "Request Time Out"})
		}
	}
}
