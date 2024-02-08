package com.hw.rockpaperscissors.game.impl;

import com.hw.rockpaperscissors.game.engine.RockPaperScissorsEngine;
import com.hw.rockpaperscissors.game.engine.generator.GameHandGenerator;
import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.model.GameHand;
import com.hw.rockpaperscissors.game.model.GameResult;
import com.hw.rockpaperscissors.game.model.GameResultEnum;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.when;

public class SimpleGameTest {
    @Mock
    RockPaperScissorsEngine rockPaperScissorsEngineMock;

    @Mock
    GameHandGenerator gameHandGeneratorMock;

    private SimpleGame simpleGame;

    @BeforeEach
    public void init() {
        simpleGame = new SimpleGame(rockPaperScissorsEngineMock, gameHandGeneratorMock);
    }

    @Test
    public void testGameResult() throws GameEngineException {
        when(gameHandGeneratorMock.generateRandom()).thenReturn(GameHand.SCISSORS);
        when(rockPaperScissorsEngineMock.runGame(GameHand.ROCK, GameHand.SCISSORS)).thenReturn(GameResultEnum.PLAYER_2);

        GameResult result = simpleGame.runGame("r");

        assertNotNull(result, "result is not null");
        assertEquals(GameResultEnum.PLAYER_2, result.getWinner(), "Winner should be player 2");
        assertEquals(GameHand.ROCK, result.getPlayer1(), "Player 1 should have shown rock");
        assertEquals(GameHand.SCISSORS, result.getPlayer2(), "Player 2 should have shown scissors");
    }

    @Test
    public void test_Throws_exception() throws GameEngineException {
        when(gameHandGeneratorMock.generateRandom()).thenReturn(GameHand.SCISSORS);
        when(rockPaperScissorsEngineMock.runGame(GameHand.ROCK, GameHand.SCISSORS))
                .thenThrow(new GameEngineException("error"));

        assertThrows(GameEngineException.class,
                () -> simpleGame.runGame("r"),
                    "Should throw GameEngineException when engine throws exception"
                );
    }
}
