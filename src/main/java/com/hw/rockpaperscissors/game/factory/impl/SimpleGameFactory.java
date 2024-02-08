package com.hw.rockpaperscissors.game.factory.impl;

import com.hw.rockpaperscissors.game.Game;
import com.hw.rockpaperscissors.game.engine.HandProcessor;
import com.hw.rockpaperscissors.game.engine.RockPaperScissorsEngine;
import com.hw.rockpaperscissors.game.engine.generator.GameHandGenerator;
import com.hw.rockpaperscissors.game.factory.GameFactory;
import com.hw.rockpaperscissors.game.impl.SimpleGame;

public class SimpleGameFactory implements GameFactory {
    private static final SimpleGameFactory instance = new SimpleGameFactory();

    public static SimpleGameFactory getFactory() {
        return instance;
    }

    @Override
    public Game getGame() {
        HandProcessor handProcessor = new HandProcessor();
        RockPaperScissorsEngine rockPaperScissorsEngine = new RockPaperScissorsEngine(handProcessor);
        GameHandGenerator gameHandGenerator = new GameHandGenerator();

        return new SimpleGame(rockPaperScissorsEngine, gameHandGenerator);
    }
}
