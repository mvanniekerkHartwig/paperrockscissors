package com.hw.rockpaperscissors.game.model;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Builder
@Getter
@Setter
public class Player {
    private String name;
    private GameHand gameHand;
}
