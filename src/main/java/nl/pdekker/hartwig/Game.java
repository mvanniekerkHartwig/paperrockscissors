package nl.pdekker.hartwig;

import java.util.EnumMap;
import java.util.Map;
import java.util.Optional;
import java.util.Scanner;

public class Game {

    private final Map<Result, Integer> summary;
    private final GameContext gameContext;

    public Game(GameContext gameContext) {
        this.gameContext = gameContext;
        summary = new EnumMap<>(Result.class);
        //initialize map, with 0, so we don't have to worry if an entry exists.
        for (var r : Result.values()) {
            summary.put(r, 0);
        }
    }

    public void play() {
        //this will close the scanner when we are done.
        try (var scanner = gameContext.getUserInputScanner()) {

            Move user;
            while ((user = getUserMove(scanner)) != null) {
                Move computer = getComputerMove();
                Result result = user.beats(computer);
                System.out.println("you: " + result.name() + ", your move: " + user + ", computer: " + computer);
                //we initialized the map, so all keys are present with value 0, so disable null warning
                //noinspection ConstantConditions
                summary.compute(result, (k, v) -> v + 1);
            }
        }
        quit();
    }

    /**
     * Reads a line from the scanner for user input, if the user input is a {@link Move} then it will return this move
     * if the input = "quit", it will return null.
     * <p>
     * default access, so it can be accessed by the unittest.
     *
     * @param scanner used to read the input
     * @return {@link  Move} or Null, null indicates "quit"
     */
    Move getUserMove(Scanner scanner) {
        Optional<Move> move;
        do {
            System.out.println("Choose: Paper, Rock, Scissor or type quit to stop.");
            var input = scanner.nextLine();
            if ("quit".equalsIgnoreCase(input)) {
                return null;
            }
            move = Move.fromValue(input);
        } while (move.isEmpty());
        return move.get();
    }

    /**
     * Returns a random {@link Move}, using the random generator from {@link GameContext}
     *
     * @return a random move.
     * <p>
     * default access, so it can be accessed by the unittest.
     */
    Move getComputerMove() {
        return Move.random(gameContext.getRandomGenerator());
    }

    public int getStatistics(Result r) {
        return summary.get(r);
    }

    private void quit() {
        System.out.println("Thanks for playing:");
        summary.forEach((key, value) -> System.out.println(key + " : " + value));
    }

    public static void main(String[] args) {
        Game game = new Game(GameContextFactory.defaultContext());
        game.play();
    }
}
