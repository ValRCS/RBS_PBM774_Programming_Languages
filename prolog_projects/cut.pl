max(X, Y, X) :- X > Y, !. % AND CUT when X>=Y cut if we get the first one
max(X, Y, theyaresame) :- X == Y, !.
max(_, Y, Y). % fallback case

print_solution() :-
    max(10,20,X),
    % max(20,10,X).
    writeln(X),
    max(25, 10, Y),
    writeln(Y),
    max(20, 20, Z),
    writeln(Z),
    writeln("Looks like max works!").

:- print_solution.
:- halt.