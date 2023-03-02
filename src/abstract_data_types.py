# Abstract Data types

# float can be thought of as a real number
# floating point implementation is abstracted away from the user
# sometimes this abstraction is not perfect
# 0.1 + 0.2 = 0.30000000000000004
# print(0.1+0.2)
# we call it abstract is leaking
# we can use the decimal module to fix this or round and so on
# before 1980s there were competing standards for floating point numbers
# now we have a standard
# https://en.wikipedia.org/wiki/IEEE_754

# Stack
# Stack is a data structure that follows the LIFO principle

# implementation
# lets define a class
class Stack:
    # define the constructor
    def __init__(self):
        # define the data structure
        self.stack = [] # the whole idea is that list or something is not exposed to the user
        # the implementation is abstracted away
        # we could be using a linked list or something else
    # define the push method
    def push(self, value):
        self.stack.append(value)
    # define the pop method
    def pop(self):
        return self.stack.pop()
    # define the peek method
    def peek(self):
        return self.stack[-1]
    # define the is_empty method
    def is_empty(self):
        return len(self.stack) == 0
    # define the size method
    def size(self):
        return len(self.stack)
    
# now we can use it and not worry about the implementation

# lets define a function that checks if a string is a palindrome
def is_palindrome(string):
    # create a stack
    stack = Stack()
    # push all the characters of the string
    for char in string:
        stack.push(char)
    # pop all the characters and check if they are the same
    for char in string:
        if char != stack.pop():
            return False
    return True

# lets test it
print(is_palindrome("racecar"))
print(is_palindrome("hello"))
# abba
print(is_palindrome("abba"))


# lets look at a class with both class and instance methods

# lets define a class House
class House:
    # define the constructor with parameters
    def __init__(self, color, size):
        # define the attributes
        self.color = color
        self.size = size
    # define the class method
    @classmethod
    def from_square_feet(cls, color, square_feet):
        # create a new instance of the class
        # cls is the class itself
        return cls(color, square_feet / 10)
    # class methods can be called without creating an instance
    # useful for utility methods
    @classmethod
    def calculate_square_feet(cls, size):
        size_in_square_feet = size * 10
        print(f"The size in square feet is: {size_in_square_feet}")
        return size_in_square_feet

    # define the instance method
    def paint(self, color):
        self.color = color
    # define the str method
    def __str__(self):
        return f"House({self.color}, {self.size})"
    
# I can call class method BEFORE I create an instance
House.calculate_square_feet(60)

# lets create an instance of the class
my_house = House("blue", 100)
print(my_house)
#paint my house
my_house.paint("red")
print(my_house)
# create a new house from square feet
my_house2 = House.from_square_feet("blue", 500)
print(my_house2)
