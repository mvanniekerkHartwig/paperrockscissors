package com.hw.rockpaperscissors.game.engine;

import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.exceptions.HandProcessorException;
import com.hw.rockpaperscissors.game.model.GameHand;
import com.hw.rockpaperscissors.game.model.GameResultEnum;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import static com.hw.rockpaperscissors.game.model.GameHand.PAPER;
import static com.hw.rockpaperscissors.game.model.GameHand.ROCK;
import static com.hw.rockpaperscissors.game.model.GameResultEnum.*;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class RockPaperScissorsEngineEngineTest {
    @Mock
    private HandProcessor handProcessorMock;

    private RockPaperScissorsEngine rockPaperScissorsEngine;

    @BeforeEach
    public void init() {
        rockPaperScissorsEngine = new RockPaperScissorsEngine(handProcessorMock);
    }

    @Test
    public void test_Tie() throws GameEngineException {
        GameResultEnum result = rockPaperScissorsEngine.runGame(GameHand.ROCK, GameHand.ROCK);

        assertEquals(TIE, result, "Result should be a tie");
    }

    @Test
    public void test_Player1_Wins() throws GameEngineException, HandProcessorException {
        when(handProcessorMock.getWinningHand(PAPER, ROCK)).thenReturn(PAPER);

        GameResultEnum result = rockPaperScissorsEngine.runGame(PAPER, ROCK);

        assertEquals(PLAYER_1, result, "Winner should be Player 1 PAPER");
    }

    @Test
    public void test_Player2_Wins() throws GameEngineException, HandProcessorException {
        when(handProcessorMock.getWinningHand(ROCK, PAPER)).thenReturn(PAPER);

        GameResultEnum result = rockPaperScissorsEngine.runGame(ROCK, PAPER);

        assertEquals(PLAYER_2, result, "Winner should be Player 2 PAPER");
    }

    @Test
    public void test_processor_exception() throws HandProcessorException {
        when(handProcessorMock.getWinningHand(ROCK, PAPER)).thenThrow(new HandProcessorException("error"));

        assertThrows(GameEngineException.class,
                () -> rockPaperScissorsEngine.runGame(ROCK, PAPER),
                "Should throw a GameEngineException when Processor throws an exception");
    }

}
