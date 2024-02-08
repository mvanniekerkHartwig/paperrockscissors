package com.hw.rockpaperscissors.game.engine.generator;

import com.hw.rockpaperscissors.game.model.GameHand;

import java.util.UUID;

import static com.hw.rockpaperscissors.game.model.GameHand.*;

public class GameHandGenerator {
    public GameHand generateRandom() {
        var random = UUID.randomUUID().getMostSignificantBits();

        if (random % 2 == 0) {
            if (random % 3 == 0) {
                return ROCK;
            }

            return PAPER;
        }

        if (random % 3 == 0) {
            return ROCK;
        }

        return SCISSORS;
    }
}
