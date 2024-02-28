package nl.pdekker.hartwig;

import org.junit.jupiter.api.Test;

import java.util.Optional;

import static nl.pdekker.hartwig.Move.*;
import static nl.pdekker.hartwig.Result.*;
import static org.junit.jupiter.api.Assertions.*;

public class MoveTest {

    @Test
    public void fromValue() {
        assertEquals(Optional.of(ROCK), Move.fromValue("rock"));
        assertEquals(Optional.of(PAPER), Move.fromValue("Paper"));
        assertEquals(Optional.of(SCISSORS), Move.fromValue("SCISSORS"));
        assertEquals(Optional.empty(), Move.fromValue(null));
        assertEquals(Optional.empty(), Move.fromValue("NotAMove"));
    }

    @Test
    public void beats() {
        assertEquals(TIE, PAPER.beats(PAPER));
        assertEquals(WON, PAPER.beats(ROCK));
        assertEquals(LOST, PAPER.beats(SCISSORS));

        assertEquals(LOST, ROCK.beats(PAPER));
        assertEquals(TIE, ROCK.beats(ROCK));
        assertEquals(WON, ROCK.beats(SCISSORS));

        assertEquals(WON, SCISSORS.beats(PAPER));
        assertEquals(LOST, SCISSORS.beats(ROCK));
        assertEquals(TIE, SCISSORS.beats(SCISSORS));

        assertThrowsExactly(NullPointerException.class, () -> ((Move) null).beats(PAPER));
        assertThrowsExactly(NullPointerException.class, () -> SCISSORS.beats(null));
    }
}
