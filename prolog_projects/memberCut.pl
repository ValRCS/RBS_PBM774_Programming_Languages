% let's look at an example of using cut operator to look for element in a list
% Base case
memberCut(X, [X|_]) :- !.
% Recursive case
memberCut(X, [_|T]) :- memberCut(X, T).

% let's test it
testMemberCut() :-
    memberCut(3, [1,2,3,4,5]),
    writeln("3 is a member of the list"),
    \+ memberCut(6, [1,2,3,4,5]), % \+ is the negation operator in swi-prolog
    % so 6 was not found and we negated the result to get true
    % https://www.swi-prolog.org/pldoc/man?predicate=%5C%2B/1
    writeln("6 is not a member of the list").

:- testMemberCut.
:- halt.