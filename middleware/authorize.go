package middleware

import (
	"context"
	"fmt"
	"strings"
	"todo-list/common"
	"todo-list/component/tokenprovider"
	"todo-list/module/user/model"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

// nếu có lỗi từ header
func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthenHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// Authorization : "Bearer ${token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// requireAuth
// 1. get token from header
// 2. validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB

func RequiredAuth(authStore AuthenStore, tokenProvider tokenprovider.Provider) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId()})

		if err != nil {
			panic(err)
		}

		// check status

		// if user.Status == 0 {
		// 	panic(common.ErrNoPermission())
		// }

		c.Set(common.CurrentUser, user) // set vao context cua gin = c.Set("current_user", user)
		c.Next()
	}
}
