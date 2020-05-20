package poker_test

import (
	"bytes"
	"fmt"
	"github.com/kilsenp/application"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartedWith  int
	StartCalled  bool
	FinishedWith string
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}

func TestCLI(t *testing.T) {

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("3\nCleo wins\n")

		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 3 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

		want := "Cleo"
		if want != game.FinishedWith {
			t.Fatalf("got %q want %q", game.FinishedWith, want)
		}

	})

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("7\nChris wins\n")

		game := &GameSpy{}
		stdout := &bytes.Buffer{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

		want := "Chris"
		if want != game.FinishedWith {
			t.Fatalf("got %q want %q", game.FinishedWith, want)
		}

	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})

	t.Run("Prints an error when a winner is declared incorreclty", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("5\nLloyd is a killer")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.FinishCalled {
			t.Errorf("game shoud not have finished")
		}
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})

}

func assertMessageSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}

}

func assertScheduledAlert(t *testing.T, got scheduledAlert, want scheduledAlert) {
	amountGot := got.amount
	if amountGot != want.amount {
		t.Errorf("got amount %d, want %d", amountGot, want.amount)
	}

	gotScheduledTime := got.at
	if gotScheduledTime != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)

}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}
