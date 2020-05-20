package poker

import (
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func (g *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}
func (g *TexasHoldem) Finish(player string) {
	g.store.RecordWin(player)
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter, store}
}