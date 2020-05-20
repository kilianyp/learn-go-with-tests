package poker

import "testing"

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWind want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store the correct winner got %q want %q", store.winCalls[0], winner)
	}

}
