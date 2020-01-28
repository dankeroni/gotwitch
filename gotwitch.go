package gotwitch

// TwitchAPI struct
type TwitchAPI struct {
	v5 *TwitchAPIV5

	helix *TwitchAPIHelix

	id *TwitchAPIID
}

func (a *TwitchAPI) V5() *TwitchAPIV5 {
	return a.v5
}

func (a *TwitchAPI) Helix() *TwitchAPIHelix {
	return a.helix
}

func (a *TwitchAPI) ID() *TwitchAPIID {
	return a.id
}

// New instantiates a new TwitchAPI object
func New(clientID string) *TwitchAPI {
	api := &TwitchAPI{
		v5: NewTwitchAPIV5(clientID),

		helix: NewTwitchAPIHelix(clientID),

		id: NewTwitchAPIID(clientID),
	}

	return api
}
