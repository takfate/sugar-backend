package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func urlsRoot(c *gin.Context) {
	urls := map[string]string{
		"GET_code_url": "/code",
	}
	j, _ := json.Marshal(urls)
	r := string(j)
	c.String(http.StatusOK, r)
}

func initRouter() {
	r.GET("/", urlsRoot)
	// account start
	r.GET("/accounts/code", accountSendVerificationCode)
	r.POST("/accounts/create", accountRegister)
	r.GET("/accounts/login", accountLogin)
	r.POST("/accounts/alter", accountAlterInformation)
	r.GET("/accounts/info", accountGetUserInformation)
	r.POST("/accounts/alter/password", accountAlterPassword)
	r.GET("/accounts/privacy", accountGetUserPrivacySetting)
	r.POST("/accounts/alter/privacy", accountAlterUserPrivacySetting)
	r.GET("/accounts/logout", accountLogout)
	r.GET("/accounts/follower", accountGetUserFollowerList)
	r.GET("/accounts/following", accountGetUserFollowingList)
	r.POST("/accounts/following/follow", accountFollowUser)
	r.POST("/accounts/following/ignore", accountIgnoreUser)
	// account end

	// bbs start
	r.POST("/bbs/topic/publish")
	r.POST("/bbs/topic/lord-reply/publish")
	r.POST("/bbs/topic/layer-reply/publish")
	// bbs end

	// school start
	r.GET("/school/article/:id", schoolGetArticlePage)
	r.GET("/school/article", schoolGetArticle)
	r.GET("/school/articles", schoolGetArticleList)
	r.POST("/school/article/:id/comment", schoolPublishArticleComment)
	r.GET("/school/article/:id/comments", schoolGetArticleCommentList)
	r.GET("/school/articles/search", schoolSearchArticle)
	r.GET("/school/articles/user-collected", schoolGetUserCollectedArticleList)
	r.GET("/school/articles/user-comments", schoolGetUserArticleCommentList)
	r.POST("/school/articles/user-collect", schoolCollectArticle)
	r.POST("/school/articles/user-cancel-collect", schoolCancelCollectArticle)
	// school end
}
