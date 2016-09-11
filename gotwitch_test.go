package gotwitch

/*
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	newTests := []struct {
		clientID          string
		expectedTwitchAPI *TwitchAPI
		expectedResult    bool
		expectError       bool
	}{
		{
			clientID: "908kfajosdia8sd8hjkkj",
			expectedTwitchAPI: &TwitchAPI{
				ClientID: "908kfajosdia8sd8hjkkj",
			},
			expectedResult: true,
		},
		{
			clientID: "888sdfgh83489xcvsdifi",
			expectedTwitchAPI: &TwitchAPI{
				ClientID: "8f8gfdghkudflkjgho8",
			},
			expectedResult: false,
		},
	}

	for _, test := range newTests {
		twitchAPI := New(test.clientID)
		RunSimpleTest(t, test.expectedTwitchAPI, twitchAPI, test.expectedResult)
	}
}

func RunSimpleTest(t *testing.T, expected interface{}, actual interface{}, expectedResult bool) {
	if expectedResult {
		assert.Equal(t, expected, actual)
	} else {
		assert.NotEqual(t, expected, actual)
	}
}
*/
