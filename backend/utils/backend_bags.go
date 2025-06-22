package utils

import (
	"cmp"
	"math/rand/v2"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h25s_09/handler/middleware"
	"github.com/traP-jp/h25s_09/repository"
)

type Bug struct {
	Name         string
	Probability  float64 // 発生確率(0.0-1.0) default:0.25
	ValidTimeSec int     // 有効期間(秒) default:10sec
}

var (
	BackendBugs = map[int]Bug{
		1:   {Name: "投稿の日時がおかしい", Probability: 0.2},
		2:   {Name: "データが取得できない", Probability: 0.05},
		3:   {Name: "ダイヤルアップ", Probability: 0.25, ValidTimeSec: 10}, // 画像が半分しか出ない
		4:   {Name: "同じ投稿が複数ある"},
		5:   {Name: "API制限", ValidTimeSec: 5},
		6:   {Name: "ユーザーが全部同じに見える"},
		7:   {Name: "リプ増殖:motto:", Probability: 0.8},
		8:   {Name: "解像度が低いな"},
		9:   {Name: "阿部寛で爆速に"},
		10:  {Name: "レスポンスが遅い"},
		11:  {Name: "いいねが無限に増やせる"},
		12:  {Name: "TL/ツイート増殖"},
		100: {Name: "リプライ増殖", Probability: 0.65},
		404: {Name: "404 このバグは存在しないよ……？"},
	}
)

func DetermineDispatchBug(ctx echo.Context, repo repository.Repository, bugID int) bool {
	if IsValidBugNow(ctx, bugID) {
		return true
	}
	if bug, exists := BackendBugs[bugID]; exists {
		p := cmp.Or(bug.Probability, 0.25)
		if p >= 1.0 || (0 < p && rand.Float64() < p) {
			if bug.ValidTimeSec > 0 {
				AddOrUpdateBugState(ctx, bugID, cmp.Or(bug.ValidTimeSec, 10))
			}
			repo.InsertUserAchievement(ctx.Get(middleware.UsernameKey).(string), bug.Name)
			return true
		}
	}
	return false
}
