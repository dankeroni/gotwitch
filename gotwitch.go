package gotwitch

// TwitchAPI struct
type TwitchAPI struct {
	v5 *TwitchAPIV5

	id *TwitchAPIID

	helix *TwitchAPIHelix

	tmi *TwitchAPITMI
}

// V5 returns the V5 API object
func (a *TwitchAPI) V5() *TwitchAPIV5 {
	return a.v5
}

// ID returns the ID API object
func (a *TwitchAPI) ID() *TwitchAPIID {
	return a.id
}

// Helix returns the Helix API object
func (a *TwitchAPI) Helix() *TwitchAPIHelix {
	return a.helix
}

// TMI returns the TMI API object
func (a *TwitchAPI) TMI() *TwitchAPITMI {
	return a.tmi
}

// New instantiates a new TwitchAPI object
func New(clientID string) *TwitchAPI {
	api := &TwitchAPI{}

	api.v5 = NewTwitchAPIV5(clientID)
	api.id = NewTwitchAPIID(clientID)
	api.helix = NewTwitchAPIHelix(clientID, api.id)
	api.tmi = NewTwitchAPITMI(clientID)

	return api
}

// SetDebug sets the debug flag on all api objects resty client
func (a *TwitchAPI) SetDebug(v bool) {
	a.v5.c.SetDebug(v)
	a.id.c.SetDebug(v)
	a.helix.c.SetDebug(v)
	a.tmi.c.SetDebug(v)
}

// SetClientSecret sets the client secret for the ID api
func (a *TwitchAPI) SetClientSecret(clientSecret string) {
	a.ID().SetClientSecret(clientSecret)
}
