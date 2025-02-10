% Defining parent-child relationships
parent(albert, jim).
parent(albert, peter).
parent(jim, brian).
parent(jane, brian).
parent(peter, anna).
% Grandparent relationship
grandparent(GP, GC) :- 
    parent(GP, P),      % GP is the parent of P
    parent(P, GC).      % P is the parent of GC
% A rule to check if someone is a child of another
child(C, P) :- 
    parent(P, C).

% A rule to check if someone is a sibling (i.e., they have a common parent)
sibling(X, Y) :- 
    parent(P, X), 
    parent(P, Y), 
    X \= Y.            % Ensure they aren't the same person