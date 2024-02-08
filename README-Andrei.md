# Paper Rock Scissors Exercise

The application is built as modular and decoupled as possible within the given time frame, considering also a very
basic test coverage.

## Console - main function
The user interface in the form of a console application is built in the class GameConsole. 

This is decoupled from the rest of the game implementation, interacting with the rest of the game via the Game interface.

## GameFactory
The factory pattern is used for decoupling the game user interface from the rest of the game implementation and fetching an instance of the Game execution.
Dependencies are managed inside the game factory.

## Stateless implementation
All the classes used in building the game are stateless. The most obvious choice for this simple example is to implement these classes as PoJos
with static methods and no instance dependency.

However, my choice was to implement aggregations with object references, to show clear structure of dependency and also to
follow the pattern of server side frameworks such as Spring, with the Dependency Injection and IoC principles.

This also makes unit testing easier and shows dependencies clearer in tests as well.

## Testing
A basic coverage has been implemented in the src/test/java root folder based on Junit 5 and Mockito.

## Dependency management
Maven has been used.

## Missing requirement:

After completion I have noticed that I have unfortunately forgotten to implement a history of the current game to be displayed upon completion.
This could be done in a new class, that could hold a list of such games.
Each game is already displaying the winner and the model of the GameResult already contains the showed hands and result.

The SimpleGame could add a dependency to this class and add a new entry each time the game is played.

Then a function could be added to SimpleGam to return the list of played games for the current instance.

Finally, the console would call this function when the user enters the quit input. And display this summary.

