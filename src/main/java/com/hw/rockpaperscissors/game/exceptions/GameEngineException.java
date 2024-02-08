package com.hw.rockpaperscissors.game.exceptions;

import lombok.AllArgsConstructor;

public class GameEngineException extends Exception {
    public GameEngineException(String msg) {
        super(msg);
    }
    public GameEngineException(String msg, Exception cause) {
        super(msg, cause);
    }
}
