package biz

import (
	"context"
	"todo-list/common"
	"todo-list/component/tokenprovider"
	"todo-list/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. hash pass from input and compare with pass in db
// 3. provider issuse JWT token for client
// 3.1 access token and refresh token
// 4 return token

func (business *loginBusiness) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// refeshToken, err := business.tokenProvider.Gennerate(payload, business.tkCfg.GetRtExp())
	// if err != nil {
	// 	return nil, common.ErrInternal(err)
	// }

	return accessToken, nil

}
