package com.hw.rockpaperscissors.game.engine;

import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.model.GameHand;
import com.hw.rockpaperscissors.game.model.GameResultEnum;
import com.hw.rockpaperscissors.game.model.Player;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import static com.hw.rockpaperscissors.game.model.GameHand.*;
import static com.hw.rockpaperscissors.game.model.GameResultEnum.*;

/**
 * Class that represents a stateless game engine.
 * Uses an internal representation of rankings.
 * Thread safe implementation using ConcurrentHashMap and a stateless engine function
 */
public final class RockPaperScissorsEngine {
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
     * @param gameHandPlayer1 shown hand of player 1
     * @param gameHandPlayer2 shown hand of player 2
     * @return {@link GameResultEnum} the result of the hand
     * @throws {@link GameEngineException} in case the unlikely case that the comparison cannot be carried out
     */
    public static GameResultEnum runGame(final GameHand gameHandPlayer1, final GameHand gameHandPlayer2) throws GameEngineException{
        if (gameHandPlayer1 == gameHandPlayer2) {
            return TIE;
        }

        if (winnerToLoserHandMap.get(gameHandPlayer1) == gameHandPlayer2) {
            return PLAYER_1;
        }

        if (winnerToLoserHandMap.get(gameHandPlayer2) == gameHandPlayer1) {
            return PLAYER_2;
        }

        throw new GameEngineException("Unexpected input, could not match; Player 1 hand : " + gameHandPlayer1 + "; Player 2 hand : " + gameHandPlayer1);
    }
}


