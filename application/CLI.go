package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerInputErrMsg = "Bad value received for winner, expect format of 'PlayerName wins'"

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberOfPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}
	c.game.Start(numberOfPlayers, c.out)
	winner, err := extractWinner(c.readLine())
	if err != nil {
		fmt.Fprint(c.out, BadWinnerInputErrMsg)
		return
	}
	c.game.Finish(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()

}
func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	reader := bufio.NewScanner(in)
	return &CLI{reader, out, game}
}
