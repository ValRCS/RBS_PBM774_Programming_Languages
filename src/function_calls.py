# lets explore function calls

# lets define a function
def eat(food):
    print("I am eating " + food)

def drink(drink):
    print("I am drinking " + drink)

# let's define a function that eats tons of food
def eat_tons(food, amount):
    for i in range(amount):
        eat(food)

# define function that eats food froma list
def eat_list(food_list):
    for food in food_list:
        eat(food)

my_favorite_foods = ["pizza", "ice cream", "chocolate", "sushi", "tacos"]

# call the function
eat("pizza")
drink("water")
eat_tons("pizza", 5)

# call the function with a list
eat_list(my_favorite_foods)

print("We are good and full now!")
