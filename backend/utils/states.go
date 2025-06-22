package utils

import (
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const StatesSessionKey = "bugStates"

type bugState struct {
	BugID       int       `json:"bugId"`
	ValidBefore time.Time `json:"validBefore"`
}

func getValidBugStates(ctx echo.Context, ss sessions.Store) map[int]bugState {
	if ss == nil {
		return map[int]bugState{}
	}
	sess, err := ss.Get(ctx.Request(), StatesSessionKey)
	if err != nil {
		return map[int]bugState{}
	}
	states := make(map[int]bugState)
	for _, v := range sess.Values {
		if s, ok := v.(bugState); ok {
			if time.Now().Before(s.ValidBefore) {
				states[s.BugID] = s
			}
		}
	}
	clear(sess.Values)
	for _, s := range states {
		sess.Values[s.BugID] = s
	}
	sess.Save(ctx.Request(), ctx.Response())
	return states
}

func AddOrUpdateBugState(ctx echo.Context, ss sessions.Store, bugID int, validTimeSec int) {
	if ss == nil {
		return
	}
	sess, err := ss.Get(ctx.Request(), StatesSessionKey)
	if err != nil {
		return
	}
	sess.Values[bugID] = bugState{
		BugID:       bugID,
		ValidBefore: time.Now().Add(time.Duration(validTimeSec) * time.Second),
	}
	sess.Save(ctx.Request(), ctx.Response())
}

func IsValidBugNow(ctx echo.Context, bugID int, ss sessions.Store) bool {
	_, ok := getValidBugStates(ctx, ss)[bugID]
	return ok
}
