package nl.pdekker.hartwig;

import org.jetbrains.annotations.NotNull;

import java.util.Optional;
import java.util.Random;

public enum Move {
    PAPER, ROCK, SCISSORS;

    /**
     * @param value a String
     * @return Optional<Move> when the name() of enum equals value (ignoring case). Or Optional.Empty when value is not present in Move
     */
    public static Optional<Move> fromValue(String value) {
        for (var m : values()) {
            if (m.name().equalsIgnoreCase(value)) {
                return Optional.of(m);
            }
        }
        return Optional.empty();
    }

    public static Move random(Random random) {
        var index = random.nextInt(Move.values().length);
        return Move.values()[index];
    }

    /**
     * this function compare two moves (this and other), if both moves are the same it will return Result.TIE}
     * in the following cases it will return Result.WON
     * PAPER beats ROCK, ROCK beats SCISSOR, SCISSOR beats PAPER
     * in all other case it returns Result.LOST
     *
     * @param other move to compare with this.
     * @return the result of the Comparison Result.WIN, Result.TIE or Result.TIE
     * @throws NullPointerException if move or this is null
     */
    public Result beats(Move other) {
        if (other == null) {
            throw new NullPointerException("Move should not be null");
        }
        if (this == other) {
            return Result.TIE;
        }
        //NOTE: did not move this to the enum field, due to forward referencing.
        boolean beats = switch (this) {
            case PAPER:
                yield other == ROCK;
            case ROCK:
                yield other == SCISSORS;
            case SCISSORS:
                yield other == PAPER;
            default:
                throw new IllegalStateException("Did not expect Move type: " + this.name());
        };
        return beats ? Result.WON : Result.LOST;
    }


}
