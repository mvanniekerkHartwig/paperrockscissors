package nl.pdekker.hartwig;

import java.util.Random;
import java.util.Scanner;

public interface GameContext {

    Random getRandomGenerator();

    Scanner getUserInputScanner();
}
