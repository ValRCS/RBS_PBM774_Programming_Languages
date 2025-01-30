; lets look at a map data structure in Clojure
; a map is a collection of key-value pairs
; lookup speed is O(log32 n)
; if you need O(1) lookup speed, use a hash-map

; let's create a map
(def my-map {:name "Winnie" :age 2 :likes "Mice"})

;let's print out the map
(println my-map)

;let's create a custom function that takes map and prints each key value pair on a new line
(defn print-map [m]
  ;documentation for doseq: https://clojuredocs.org/clojure.core/doseq
  (doseq [[k v] m] ; so map is unpacked into key value pairs
    (println (str k " : " v))))

;let's call the function
(print-map my-map)

;let's count some dice rolls
;first let's make a function that returns a random number between 1 and 6
(defn roll-die []
  (+ (rand-int 6) 1))

;now let's create a function that returns n dice rolls and sums them up
(defn roll-dice [n]
  (reduce + (repeatedly n roll-die)))

;now let's make a function that returns x number of d roll-dice  results
(defn roll-dice-x-times [x d]
  ;documentation for repeatedly: https://clojuredocs.org/clojure.core/repeatedly
  (vec (repeatedly x #(roll-dice d))))

;now let's take a function that counts the number of times each sum appears and returns a map of results
(defn count-dice-rolls [rolls]
  ;documentation for frequencies: https://clojuredocs.org/clojure.core/frequencies
  ; so very similar to how Python's Counter works
  (frequencies rolls))

;let's create a function that takes a map and returns it sorted by value
(defn sort-map-by-value [m]
  ;documentation for sort-by: https://clojuredocs.org/clojure.core/sort-by
  (sort-by val > m))


;let's make a function that prints a list of counted dice rolls
;each item in a list is a vector with the sum and the number of times it appeared
(defn print-dice-roll-counts [rolls]
  (println "Results of rolling the dice:")
  ;documentation for doseq: https://clojuredocs.org/clojure.core/doseq
  (doseq [[k v] rolls]
    (println (str k " : " v))))

;now let's create a function that prints the results of the dice rolls
(defn print-dice-rolls [rolls] 
  (println "Results of rolling the dice:")
  ;now let's sort by value and print the map
;;   (print-dice-roll-counts (count-dice-rolls (sort-map-by-value rolls)))
  ;let's just print without sorting
    (print-dice-roll-counts (count-dice-rolls rolls)))
  

;now let's create a -main function that takes two arguments, x and d, and prints the results of rolling x dice with d sides
(defn -main
  [& args]
  (if (empty? args)
    (println "Please provide two arguments, the number of dice to roll and the number of sides on each die")
    (print-dice-rolls (roll-dice-x-times (Integer/parseInt (first args)) (Integer/parseInt (second args))))))

;let's call the -main function
(-main "1000000" "4")


