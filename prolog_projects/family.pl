% Defining parent-child relationships
parent(dino, albert).
parent(albert, jim).
parent(albert, peter).
parent(jim, brian).
parent(jane, brian).
parent(peter, anna).
% Grandparent relationship
grandparent(GP, GC) :- 
    parent(GP, P),      % GP is the parent of P
    parent(P, GC).      % P is the parent of GC
% let's define great-grandparent relationship
great_grandparent(GGP, GGC) :- 
    parent(GGP, GP),    % GGP is the parent of GP
    grandparent(GP, GGC). % GP is the grandparent of GGC
% A rule to check if someone is a child of another
child(C, P) :- 
    parent(P, C).

% A rule to check if someone is a sibling (i.e., they have a common parent)
sibling(X, Y) :- 
    parent(P, X), 
    parent(P, Y), 
    X \= Y.            % Ensure they aren't the same person

% recursion can be used to define ancestor relationships
ancestor(A, D) :- 
    parent(A, D).       % A is a parent of D
ancestor(A, D) :- 
    parent(A, X),       % A is a parent of X
    ancestor(X, D).     % X is an ancestor of D

% Aunt/Uncle relationship
aunt_or_uncle(AU, N) :- 
    sibling(AU, P),     % AU is the sibling of P
    parent(P, N).       % P is the parent of N

% Cousin relationship
cousin(C, N) :- 
    parent(PC, C),      % PC is the parent of C
    parent(PN, N),      % PN is the parent of N
    sibling(PC, PN).    % PC and PN are siblings