;let's demonstrate a simple filter map reduce example
; by filtering even numbers from a list, squaring them, and then summing them

(ns filter_map_reduce)

; let's define a function that maps numbers to their squares
(defn -square
  [n] ; square brackers here mean local scope
  (* n n))

;now let's define a function that gets even numbers from the list
(defn -even? ; so really just a wrapper around the even? function
  [n]
  (even? n))

; now how about summing the numbers in a function
(defn -sum
  [nums]
  ; let's print the list of numbers here as a side effect 
  (println nums)
  (reduce + nums)
  )

; now let's define a function that filters, maps, and reduces
(defn -filter-map-reduce
  [nums]
  (-sum ; finally sum requires just one argument the list
   (map -square ; now map also requires two arguments, the function and the list
        (filter -even? nums))) ; so filter requires two arguments, the function and the list
)

(defn -main
  [& args]
    (if (empty? args)
        (println "Please provide a list of numbers as an argument")
        (println (-filter-map-reduce (map #(Integer/parseInt %) args))))
)

(-main "1" "2" "3" "4" "5" "6" "7" "8" "9" "10")
