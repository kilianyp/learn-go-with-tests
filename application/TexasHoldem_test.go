package poker_test

import (
	"fmt"
	"github.com/kilsenp/application"
	"testing"
	"time"
)

func TestGame(t *testing.T) {

	var dummyPlayerStore = &poker.StubPlayerStore{}

	t.Run("schedules alerts for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}

		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(5)

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}

		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
		game.Start(7)

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	dummyBlindAlerter := &SpyBlindAlerter{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"
	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []scheduledAlert, t *testing.T, blindAlerter *SpyBlindAlerter) {
	t.Helper()
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.at), func(t *testing.T) {
			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}

			alert := blindAlerter.alerts[i]
			assertScheduledAlert(t, alert, c)

		})
	}

}
