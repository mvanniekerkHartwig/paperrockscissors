package nl.pdekker.hartwig;

import org.junit.jupiter.api.Test;

import java.util.Scanner;

import static nl.pdekker.hartwig.Move.*;
import static org.junit.jupiter.api.Assertions.*;

public class GameTest {

    @Test
    public void getUserInput() {
        var game = new Game(GameContextFactory.defaultContext());

        assertNull(game.getUserMove(new Scanner("quit\n")));
        assertEquals(PAPER, game.getUserMove(new Scanner("Paper\n")));
        assertEquals(SCISSORS, game.getUserMove(new Scanner("SCISSORS\n")));
        assertEquals(ROCK, game.getUserMove(new Scanner("wronginput\nrock\n")));
    }

    @Test
    public void getComputerMove() {
        Game game = new Game(GameContextFactory.unitTestContext(""));
        //the unittestContext has fixed random seed, so we get pseudo random numbers
        assertEquals(PAPER, game.getComputerMove());
        assertEquals(ROCK, game.getComputerMove());
        assertEquals(ROCK, game.getComputerMove());
    }


    @Test
    public void playGame() {
        var userInput = "Scissors\nRock\nScissors\nquit\n";
        var game = new Game(GameContextFactory.unitTestContext(userInput));
        //the computer will play Paper, Rock, Rock see getComputerTest above
        //user will play Scissor, Rock, Scissor
        game.play();

        assertAll(
                () -> assertEquals(1, game.getStatistics(Result.TIE), "TIE"),
                () -> assertEquals(1, game.getStatistics(Result.WON), "WON"),
                () -> assertEquals(1, game.getStatistics(Result.LOST), "LOST")
        );
    }
}
