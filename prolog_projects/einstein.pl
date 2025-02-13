% Einstein's Riddle, also known as the Zebra Puzzle, is a logic puzzle that's often attributed to Albert Einstein. The puzzle is as follows:

% There are 5 houses in 5 different colors. In each house lives a person with a different nationality. The 5 owners drink a certain drink, smoke a certain brand of cigar, and keep a certain pet. No owners have the same pet, smoke the same brand of cigar, or drink the same beverage.

% The clues are:


% - The Norwegian lives in the first house.
% - The Englishman lives in the red house.
% - The green house is on the left of the white house.
% - The Danish man drinks tea.
% - The German smokes Prince.
% - The green house's owner drinks coffee.
% - The person who smokes Marlboro lives next to the one with the fish.
% - The Swede keeps dogs.
% - The Norwegian lives next to the blue house.
% - The person who keeps birds lives in the yellow house.
% - The person in the third house drinks milk.
% - The person who smokes Chesterfield lives next to the one who drinks water.
% - The owner of the white house drinks beer.
% - The person who smokes Rothmans keeps cats.

% The question is: Who owns the fish?



% Representation: [Color, Nationality, Beverage, Cigar, Pet]

% Rule to check if X is next to Y in the list
next_to(X, Y, List) :- 
   append(_, [X, Y | _], List);
   append(_, [Y, X | _], List).

% The solution for the zebra puzzle
solution(Houses, FishOwner) :-
    % Define the houses
    Houses = [[_, norwegian, _, _, _], _, [_, _, milk, _, _], _, _],
    
    % Clues
    member([red, englishman, _, _, _], Houses),
    next_to([_, norwegian, _, _, _], [blue, _, _, _, _], Houses),
    member([green, _, coffee, _, _], Houses),
    member([_, danish, tea, _, _], Houses),
    member([_, _, _, marlboro, fish], Houses),
    member([_, _, _, chesterfield, _], Houses),
    next_to([_, _, _, marlboro, _], [_, _, water, _, _], Houses),
    member([yellow, _, _, _, birds], Houses),
    member([_, swede, _, _, dogs], Houses),
    member([_, _, beer, _, _], Houses),
    member([white, _, _, _, _], Houses),
    member([_, german, _, prince, _], Houses),
    next_to([green, _, _, _, _], [white, _, _, _, _], Houses),
    member([_, _, _, rothmans, cats], Houses),

    % Who owns the fish?
    member([_, FishOwner, _, _, fish], Houses).

% let's print the solution
print_solution() :-
    solution(Houses, FishOwner), % we bind the variables to the solution
    write('The solution is: '), nl,
    write('House 1: '), write(Houses), nl,
    write('The owner of the fish is: '), write(FishOwner), nl.

:- print_solution.
:- halt.



