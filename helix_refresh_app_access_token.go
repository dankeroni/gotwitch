package gotwitch

func (a *TwitchAPIHelix) RefreshAppAccessToken() (err error) {
	if a.id == nil {
		return ErrMissingID
	}

	a.appAccessToken, err = a.id.GetAppAccessToken()
	a.c.SetAuthToken(a.appAccessToken.AccessToken)
	return
}
