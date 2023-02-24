import random
small_list = [1, 2, 3, 4, 5]
medium_list = list(range(1000))
big_list = list(range(1_000_000))
# calculate memory usage
print("The size of the small list is: " + str(small_list.__sizeof__()))

# curent system memory usage
# https://stackoverflow.com/questions/938733/total-memory-used-by-python-process
import os
print("The current memory usage is: " + str(os.getpid()))
huge_list = list(range(100_000_000))
print("The size of the huge list is: " + str(huge_list.__sizeof__()))
# calculate real memory usage in os
# current memory usage
# print("The current memory usage is: " + str(os.getpid()))

# lets define a function
# that takes as size argument and creates a random list and calculates average

def calculate_average(size):
    # create a list of random numbers
    random_list = []
    for i in range(size):
        random_list.append(random.randint(0, 100))
    # calculate average
    average = sum(random_list) / len(random_list)
    print("The average of the list is: " + str(average))
    # i am not returning it

# call the function
calculate_average(1000)
calculate_average(1_000_000)
calculate_average(100_000_000)


print("All good!")

