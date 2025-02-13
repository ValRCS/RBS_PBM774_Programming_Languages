% let's make define a fact that contains a list of numbers
% then we will define a predicate that prints the list number by number

% Defining a fact that contains a list of numbers
numbers([1,2,3,4,5,6,7,8,9,10]).
% numbers([one, two, three, four, five, six, seven, eight, nine, ten]).
numbers([5,8,7,1000,9001,999]).
% Defining a predicate that prints the list number by number
print_numbers([]) :- 
    write('\n').

print_numbers([H|T]) :- 
    write(H),
    write(' '),
    print_numbers(T).


% let's create a predicate that prints numbers in reverse
print_numbers_reverse([]) :- 
    write('\n').

print_numbers_reverse([H|T]) :- 
    print_numbers_reverse(T),
    write(H),
    write(' ').

%  how about predicate that prints even numbers
print_even_numbers([]) :- 
    write('\n').

print_even_numbers([H|T]) :-
    0 is H mod 2,
    write(H),
    write(' '),
    print_even_numbers(T).

print_even_numbers([H|T]) :-
    1 is H mod 2,
    print_even_numbers(T).

% let's create a query that will print the numbers
% ?- numbers(X), print_numbers(X).
% https://stackoverflow.com/questions/44232140/how-to-run-prolog-queries-from-within-the-prolog-file-in-swi-prolog
:- forall(numbers(X), print_numbers(X)).
:- forall(numbers(Y), print_numbers_reverse(Y)).
:- forall(numbers(Z), print_even_numbers(Z)).
:- halt.