% Representation: [Color, Nationality, Beverage, Cigar, Pet]

% Rule to check if X is next to Y in the list
next_to(X, Y, List) :- 
   append(_, [X, Y | _], List);
   append(_, [Y, X | _], List).

% The solution for the zebra puzzle
originalLifeSolution(Houses, ZebraOwner, WaterDrinker) :-
    % Define the houses
    % There are Five houses
    % Norwegian lives in the first house
    % milk is drunk in the middle house
    Houses = [[_, norwegian, _, _, _], _, [_, _, milk, _, _], _, _],
    
    % Clues
    member([red, englishman, _, _, _], Houses),
    member([_, spaniard, _, _, dog], Houses),
    member([green, _, coffee, _, _], Houses),
    member([_, ukranian, tea, _, _], Houses),
    next_to([green, _, _, _, _], [ivory, _, _, _, _], Houses),
    member([_, _, _, oldgold, snails], Houses),
    member([yellow, _, _, kools, _], Houses),

    % man who smokes Chesterfields lives in the house next to the man with the fox
    next_to([_, _, _, chesterfield, _], [_, _, _, _, fox], Houses),
    % Kools are smoked in the house next to the house where the horse is kept
    next_to([_, _, _, kools, _], [_, _, _, _, horse], Houses),
    member([_, _, orangejuice, luckystrike, _], Houses),
    % Japanese smokes Parliaments
    member([_, japanese, _, parliaments, _], Houses),
    % Norwegian lives next to the blue house
    next_to([_, norwegian, _, _, _], [blue, _, _, _, _], Houses),

    % Who owns the fish?
    member([_, ZebraOwner, _, _, zebra], Houses).
    % Who Drinks Water?
    member([_, WaterDrinker, water, _, _], Houses).