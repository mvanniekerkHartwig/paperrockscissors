package com.hw.rockpaperscissors.game.engine;

import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.exceptions.HandProcessorException;
import com.hw.rockpaperscissors.game.model.GameHand;
import com.hw.rockpaperscissors.game.model.GameResultEnum;
import com.hw.rockpaperscissors.game.model.Player;
import lombok.RequiredArgsConstructor;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import static com.hw.rockpaperscissors.game.model.GameHand.*;
import static com.hw.rockpaperscissors.game.model.GameResultEnum.*;

/**
 * Class that represents a stateless game engine.
 * Uses an internal representation of rankings.
 * Thread safe implementation using ConcurrentHashMap and a stateless engine function
 */
@RequiredArgsConstructor
public final class RockPaperScissorsEngine {
    private final HandProcessor handProcessor;

    /**
     * Executes a comparison of show of hands of 2 players.
     * @param gameHandPlayer1 shown hand of player 1
     * @param gameHandPlayer2 shown hand of player 2
     * @return {@link GameResultEnum} the result of the hand
     * @throws GameEngineException in case the unlikely case that the comparison cannot be carried out
     */
    public GameResultEnum runGame(final GameHand gameHandPlayer1, final GameHand gameHandPlayer2) throws GameEngineException{
        if (gameHandPlayer1 == gameHandPlayer2) {
            return TIE;
        }

        try {
            GameHand winningHand = handProcessor.getWinningHand(gameHandPlayer1, gameHandPlayer2);

            return winningHand == gameHandPlayer1 ? PLAYER_1
                    : PLAYER_2;
        } catch (HandProcessorException e) {
            throw new GameEngineException("Unexpected input, could not match; Player 1 hand : " + gameHandPlayer1 + "; Player 2 hand : " + gameHandPlayer1, e);
        }
    }
}


