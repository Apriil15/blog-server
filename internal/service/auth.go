package service

import "errors"

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"reuqired"`
	AppSecret string `form:"app_secret" binding:"reuqired"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info doesn't exist.")
}
