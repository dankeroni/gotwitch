package gotwitch

import (
	"testing"
)

func TestGetClipByID(t *testing.T) {
	const clipSlug = `GorgeousAntsyPizzaSaltBae`
	const expectedTitle = `come on let's go`
	clips, _, err := TwitchV5.GetClip(clipSlug)
	if err != nil {
		t.Error("should not error")
	}
	if expectedTitle != clips.Title {
		t.Error("wrong title")
	}
}

func TestGetClipByIDInvalidID(t *testing.T) {
	const clipSlug = `uidgfhuidrghuidrghuidrhguidrhgui`
	_, _, err := TwitchV5.GetClip(clipSlug)
	if err == nil {
		t.Error("must error")
	}
}
