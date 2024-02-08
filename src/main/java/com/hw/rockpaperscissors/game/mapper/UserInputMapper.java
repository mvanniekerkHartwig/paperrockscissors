package com.hw.rockpaperscissors.game.mapper;

import com.hw.rockpaperscissors.game.model.GameHand;

import static com.hw.rockpaperscissors.game.model.GameHand.*;

public class UserInputMapper {
    public static GameHand userInputToGameHand (final String gameHand) {
        return switch (gameHand) {
            case "r" -> ROCK;
            case "p"-> PAPER;
            case "s" -> SCISSORS;
            default -> null;
        };
    }
}
