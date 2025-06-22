package utils

import (
	"math/rand/v2"

	"github.com/traP-jp/h25s_09/repository"
)

type Bug struct {
	Name        string
	Probability float64
}

var (
	BackendBugs = map[int]Bug{
		1: {
			Name:        "投稿の日時がおかしい",
			Probability: 0.25,
		},
		2: {
			Name:        "データが取得できない",
			Probability: 0.25,
		},
		3: {
			Name:        "画像が半分しか表示されない",
			Probability: 0.25,
		},
		4: {
			Name:        "同じ投稿が複数ある",
			Probability: 0.25,
		},
		5: {
			Name:        "API制限",
			Probability: 0.25,
		},
		6: {
			Name:        "ユーザーが全部同じに見える",
			Probability: 0.25,
		},
		7: {
			Name:        "リプが増殖する",
			Probability: 0.25,
		},
		8: {
			Name:        "画質が悪い",
			Probability: 0.25,
		},
		9: {
			Name:        "阿部寛で爆速に",
			Probability: 0.25,
		},
		10: {
			Name:        "レスポンスが遅い",
			Probability: 0.25,
		},
		11: {
			Name:        "いいねが無限に増やせる",
			Probability: 0.25,
		},
	}
)

func DetermineDispatchBugAndRecord(id int, repo repository.Repository) (Bug, bool) {
	if bug, exists := BackendBugs[id]; exists {
		if bug.Probability >= 1.0 || (bug.Probability > 0 && rand.Float64() < bug.Probability) {
			return bug, true
		}
	}
	return Bug{}, false
}
