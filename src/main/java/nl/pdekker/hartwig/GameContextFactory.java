package nl.pdekker.hartwig;

import java.util.Random;
import java.util.Scanner;

public class GameContextFactory {

    /**
     * use for playing, user input is taken from System.in and computer imput
     * is using random numbers.
     * @return GameContext
     */
    public static GameContext defaultContext() {
        return new GameContext() {

            private final Random random = new Random();

            @Override
            public Random getRandomGenerator() {
                return random;
            }

            @Override
            public Scanner getUserInputScanner() {
                return new Scanner(System.in);
            }
        };
    }

    /**
     * use for testing, user input is taken from input string,
     * computer input is using pseudo random numbers
     * @param userInput  a string used user input, seperate moves with \n
     * @return GameContext used by game
     */
    public static GameContext unitTestContext(final String userInput) {
        return new GameContext() {

            final Random random = new Random(0);
            @Override
            public Random getRandomGenerator() {
                return random;
            }

            @Override
            public Scanner getUserInputScanner() {
                return new Scanner(userInput);
            }
        };
    }
}
