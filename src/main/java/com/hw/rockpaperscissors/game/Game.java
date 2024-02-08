package com.hw.rockpaperscissors.game;

import com.hw.rockpaperscissors.game.exceptions.GameEngineException;
import com.hw.rockpaperscissors.game.model.GameResult;

public interface Game {
    GameResult runGame(final String input) throws GameEngineException;
}
