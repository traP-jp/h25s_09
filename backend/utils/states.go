package utils

import (
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const StatesSessionKey = "bug_states"

type bugState struct {
	BugID       int       `json:"bugId"`
	ValidBefore time.Time `json:"validBefore"`
}

func getValidBugStates(ctx echo.Context, ss sessions.Store) map[int]bugState {
	if ss == nil {
		return map[int]bugState{}
	}
	sess, err := session.Get(StatesSessionKey, ctx)
	if err != nil {
		return map[int]bugState{}
	}
	result := make(map[int]bugState)
	for k, v := range sess.Values {
		idStr := k.(string)
		timeStr := v.(string)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		datetime, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			continue
		}
		if datetime.After(time.Now()) {
			ctx.Logger().Info("Hit: ", id)
			result[id] = bugState{
				BugID:       id,
				ValidBefore: datetime,
			}
		}
	}
	clear(sess.Values)
	for k, v := range result {
		sess.Values[k] = v.ValidBefore
	}
	sess.Save(ctx.Request(), ctx.Response())
	return result
}

func AddOrUpdateBugState(ctx echo.Context, ss sessions.Store, bugID int, validTimeSec int) {
	if ss == nil {
		return
	}
	sess, err := session.Get(StatesSessionKey, ctx)
	if err != nil {
		return
	}
	sess.Values[strconv.Itoa(bugID)] = time.Now().Add(time.Duration(validTimeSec) * time.Second).Format(time.RFC3339)
	sess.Save(ctx.Request(), ctx.Response())
}

func IsValidBugNow(ctx echo.Context, bugID int, ss sessions.Store) bool {
	_, ok := getValidBugStates(ctx, ss)[bugID]
	return ok
}
