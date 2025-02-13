% let's make define a fact that contains a list of numbers
% then we will define a predicate that prints the list number by number

% Defining a fact that contains a list of numbers
numbers([1,2,3,4,5,6,7,8,9,10]).
numbers([one, two, three, four, five, six, seven, eight, nine, ten]).
% Defining a predicate that prints the list number by number
print_numbers([]) :- 
    write('\n').

print_numbers([H|T]) :- 
    write(H),
    write(' '),
    print_numbers(T).

% let's create a query that will print the numbers
% ?- numbers(X), print_numbers(X).
% https://stackoverflow.com/questions/44232140/how-to-run-prolog-queries-from-within-the-prolog-file-in-swi-prolog
:- forall(numbers(X), print_numbers(X)).
:- halt.