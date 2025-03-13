// https://kotlinlang.org/docs/collection-transformations.html#test-predicates

//our main will take args
fun main(args: Array<String>) {

    val numbers = listOf("one", "two", "three", "four")

    println(numbers.any { it.endsWith("e") })
    println(numbers.none { it.endsWith("a") })
    println(numbers.all { it.endsWith("e") })

    println(emptyList<Int>().all { it > 5 })   // vacuous truth, returns true for any rule on an empty list

    //now let's look at some filtering and mapping examples
    //lets create a list of numbers from 0 to 12 inclusive
    val numbers2 = (0..12).toList()
    //let's filter out odd numbers
    val evens = numbers2.filter { it % 2 == 0 }
    println(evens)
    //let's do odd numbers
    val odds = numbers2.filterNot { it % 2 == 0 }
    println(odds)
    //let's map odd numbers to their squares
    val oddSquares = numbers2.filter { it % 2 != 0 }.map { it * it }
    println(oddSquares)
    //since we already had odds we could just map them
    val oddSquares2 = odds.map { it * it }
    println(oddSquares2)

    //let's look at partitioning
    //let's split into less than 5 and more than 5
    val (above5, below5) = numbers2.partition { it > 5 }
    println(above5)
    println(below5) //this actually includes 5 since the first one is the one that is greater than 5

    //now let's do some map stuff
    val numbersMap = mapOf("one" to 1, "two" to 2, "three" to 3)
    println(numbersMap.get("one"))
    println(numbersMap["one"])
    println(numbersMap.getOrDefault("four", 10))
    println(numbersMap["five"])               // null
    //numbersMap.getValue("six")      // exception!

    //let's test toCodePoint on Valdis
    val valdisCodePoints = toCodePoint("Valdis")
    println(valdisCodePoints)
    //let's print out the code points
    for ((letter, code) in valdisCodePoints) {
        println("$letter -> $code")
    }
}

//let's make a function that takes in a string and returns a map of the letters and their unicode code points
fun toCodePoint(s: String): Map<Char, Int> {
    // return s.map { it to it.toInt() }.toMap()
    //let's try using no Char.code property instead
    //so we we map each character to its code point
    return s.map { it to it.code }.toMap()
}

//for more map collections check out : https://kotlinlang.org/docs/map-operations.html