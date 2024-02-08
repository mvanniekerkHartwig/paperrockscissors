package com.hw.rockpaperscissors.game.model;

import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class GameResult {
    GameResultEnum winner;
    GameHand player1;
    GameHand player2;
}
