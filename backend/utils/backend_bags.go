package utils

import (
	"cmp"
	"math/rand/v2"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/repository"
)

type Bug struct {
	Name         string
	Probability  float64 // 発生確率(0.0-1.0) default:0.25
	ValidTimeSec int     // 有効期間(秒) default:10sec
}

var (
	BackendBugs = map[int]Bug{
		1:  {Name: "投稿の日時がおかしい"},
		2:  {Name: "データが取得できない"},
		3:  {Name: "画像が半分しか表示されない"},
		4:  {Name: "同じ投稿が複数ある"},
		5:  {Name: "API制限"},
		6:  {Name: "ユーザーが全部同じに見える"},
		7:  {Name: "投稿が増殖する"},
		8:  {Name: "画質が悪い"},
		9:  {Name: "阿部寛で爆速に"},
		10: {Name: "レスポンスが遅い"},
		11: {Name: "いいねが無限に増やせる"},
	}
)

func DetermineDispatchBugAndRecord(id int, repo repository.Repository) (Bug, bool) {
	if bug, exists := BackendBugs[id]; exists {
		p := cmp.Or(bug.Probability, 0.25)
		if p >= 1.0 || (0 < p && rand.Float64() < p) {
			return bug, true
		}
	}
	return Bug{}, false
}

func DetermineDispatchBug(ctx echo.Context, repo repository.Repository, ss sessions.Store, bugID int) bool {
	if IsValidBugNow(ctx, bugID, ss) {
		return true
	}
	if bug, exists := BackendBugs[bugID]; exists {
		p := cmp.Or(bug.Probability, 0.25)
		if p >= 1.0 || (0 < p && rand.Float64() < p) {
			AddOrUpdateBugState(ctx, ss, bugID, cmp.Or(bug.ValidTimeSec, 10))
			return true
		}
	}
	return false
}
