package com.hw.rockpaperscissors.console;


import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.factory.impl.SimpleGameFactory;
import com.hw.rockpaperscissors.game.model.GameResult;

import java.util.Scanner;

import static java.lang.System.out;

public class GameConsole {
    private static final String EXIT_COMMAND = ":q";
    private static final String ROCK = "r";
    private static final String PAPER = "p";
    private static final String SCISSORS = "s";

    private static final String VALID_INPUT_STRING_REGEX = "(?i)" + ROCK
            + "|" + PAPER
            + "|" + SCISSORS
            + "|" + EXIT_COMMAND;

    public static void main(String[] args) {
        var game = SimpleGameFactory.getFactory().getGame();

        var scanner = new Scanner(System.in);

        out.println("***************  ROCK PAPER SCISSORS ***************");
        var input = "";

        do {
            printWelcomeText();

            do {
                input = scanner.nextLine();
            } while (input.isBlank());

            if (! isValid(input)) {
                out.println("Invalid input : " + input);
                printWelcomeText();
                continue;
            }

            if (input.equalsIgnoreCase(EXIT_COMMAND)) {
                out.println("Goodbye");
                return;
            }

            GameResult result;

            try {
                result = game.runGame(input);

                out.println("Player 1: " + result.getPlayer1() + " ; Computer : " + result.getPlayer2());

                switch (result.getWinner()) {
                    case TIE: out.println("It's a Draw. Try again?");
                    break;

                    case PLAYER_1: out.println("Player 1 won!");
                    break;

                    case PLAYER_2: out.println("Computer won!");
                    break;
                }
            } catch (GameEngineException e) {
                out.println("Error while processing your input");

            }
        } while (!input.equalsIgnoreCase(EXIT_COMMAND));
    }

    private static boolean isValid(String input) {
        return input.matches(VALID_INPUT_STRING_REGEX);
    }
    private static void printWelcomeText () {
        out.println("Enter your choice : " + ROCK + ", " + PAPER + ", or " + SCISSORS
                + "? \n Or type " + EXIT_COMMAND + " to quit game.");
    }
}