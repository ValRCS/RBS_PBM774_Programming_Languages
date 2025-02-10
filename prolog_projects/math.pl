% you would use these by having constant first argument and variable second
add_3_and_double(X,Y) :- Y is (X+3)*2.
is_double(X,Y) :- Y is X + X.
% let's do a is_factorial predicate
is_factorial(1,1).
is_factorial(N,F) :- 
    N > 1,
    N1 is N - 1,
    is_factorial(N1,F1),
    F is N * F1.
% let's do fibonacci
is_fibonacci(1,1).
is_fibonacci(2,1).
is_fibonacci(N,F) :- 
    N > 2,
    N1 is N - 1,
    N2 is N - 2,
    is_fibonacci(N1,F1),
    is_fibonacci(N2,F2),
    F is F1 + F2.
% above implementation will not work on values past say 30, let's do a tail recursive one
% above smells like a naive implementation, let's do a tail recursive one next time
% TODO tail recursive fibonacci