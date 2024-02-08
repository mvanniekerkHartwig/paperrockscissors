package com.hw.rockpaperscissors.game.engine;

import com.hw.rockpaperscissors.game.exceptions.HandProcessorException;
import com.hw.rockpaperscissors.game.model.GameHand;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static com.hw.rockpaperscissors.game.model.GameHand.*;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class HandProcessorTest {
    private HandProcessor handProcessor;

    @BeforeEach
    public void init() {
        handProcessor = new HandProcessor();
    }
    @Test
    public void test_paper_beats_rock() throws HandProcessorException {
        GameHand result = handProcessor.getWinningHand(ROCK, PAPER);

        assertEquals(result, PAPER, "Paper beats Rock");
    }

    @Test
    public void test_paper_beats_rock_2() throws HandProcessorException {
        GameHand result = handProcessor.getWinningHand(PAPER, ROCK);

        assertEquals(result, PAPER, "Paper beats Rock");
    }

    @Test
    public void test_rock_beats_scissors() throws HandProcessorException {
        GameHand result = handProcessor.getWinningHand(ROCK, SCISSORS);

        assertEquals(result, ROCK, "Rock beats Scissors");
    }

    @Test
    public void test_rock_beats_scissors_2() throws HandProcessorException {
        GameHand result = handProcessor.getWinningHand(SCISSORS, ROCK);

        assertEquals(result, ROCK, "Rock beats Scissors");
    }

    @Test
    public void test_scissors_beats_paper() throws HandProcessorException {
        GameHand result = handProcessor.getWinningHand(SCISSORS, PAPER);

        assertEquals(result, SCISSORS, "Scissors beats Paper");
    }

    @Test
    public void test_scissors_beats_paper_2() throws HandProcessorException {
        GameHand result = handProcessor.getWinningHand(PAPER, SCISSORS);

        assertEquals(result, SCISSORS, "Scissors beats Paper");
    }
}
