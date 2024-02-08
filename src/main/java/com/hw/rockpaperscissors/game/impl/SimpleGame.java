package com.hw.rockpaperscissors.game.impl;

import com.hw.rockpaperscissors.game.Game;
import com.hw.rockpaperscissors.game.engine.RockPaperScissorsEngine;
import com.hw.rockpaperscissors.game.engine.generator.GameHandGenerator;
import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.mapper.UserInputMapper;
import com.hw.rockpaperscissors.game.model.GameHand;
import com.hw.rockpaperscissors.game.model.GameResult;
import com.hw.rockpaperscissors.game.model.GameResultEnum;
import lombok.RequiredArgsConstructor;

import java.util.Optional;

@RequiredArgsConstructor
public class SimpleGame implements Game {
    private final RockPaperScissorsEngine rockPaperScissorsEngine;
    private final GameHandGenerator gameHandGenerator;

    @Override
    public GameResult runGame(final String input) throws GameEngineException {
        GameHand gameHandPlayer1 = Optional.ofNullable(input)
                .map(UserInputMapper::userInputToGameHand)
                .orElseThrow(() -> new GameEngineException("Invalid input : " + input));

        GameHand gameHandPlayer2 = gameHandGenerator.generateRandom();

        GameResultEnum result = rockPaperScissorsEngine.runGame(gameHandPlayer1, gameHandPlayer2);

        return GameResult.builder()
                .player1(gameHandPlayer1)
                .player2(gameHandPlayer2)
                .winner(result)
                .build();
    }
}
