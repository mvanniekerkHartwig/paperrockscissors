package com.hw.rockpaperscissors.game.engine;

import com.hw.rockpaperscissors.game.exceptions.HandProcessorException;
import com.hw.rockpaperscissors.game.model.GameHand;
import com.hw.rockpaperscissors.game.model.GameResultEnum;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import static com.hw.rockpaperscissors.game.model.GameHand.*;

public final class HandProcessor {
    // Maps each hand to the inferior hand
    private static final Map<GameHand, GameHand> winnerToLoserHandMap;

    static {
        winnerToLoserHandMap = new ConcurrentHashMap<>();
        winnerToLoserHandMap.put(ROCK, SCISSORS);
        winnerToLoserHandMap.put(SCISSORS, PAPER);
        winnerToLoserHandMap.put(PAPER, ROCK);
    }

    /**
     * Executes a comparison of show of hands of 2 players.
     * @param hand1 shown hand of player 1
     * @param hand2 shown hand of player 2
     * @return {@link GameResultEnum} the result of the hand
     */
    public GameHand getWinningHand(final GameHand hand1, final GameHand hand2) throws HandProcessorException {
        if (hand1 == hand2) {
            return hand1;
        }

        if (winnerToLoserHandMap.get(hand1) == hand2) {
            return hand1;
        }

        if (winnerToLoserHandMap.get(hand2) == hand1) {
            return hand2;
        }

        throw new HandProcessorException("Unexpected input, could not match hand 1 " + hand1 + "and " + hand2);
    }
}
