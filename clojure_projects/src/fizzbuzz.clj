(ns fizzbuzz)

; let's do a simple fizzbuzz
; we will use fizz 5 and buzz 7
; so we will print fizz for multiples of 5, buzz for multiples of 7, and fizzbuzz for multiples of 5 and 7

(defn -fizzbuzz
  [n] ; n has local scope
  (if (and (= (mod n 5) 0) (= (mod n 7) 0))
    "fizzbuzz"
    (if (= (mod n 5) 0)
      "fizz"
      (if (= (mod n 7) 0)
        "buzz"
        n))))


(defn -main
    [& args]
    (if (empty? args)
        (println "Please provide a number as an argument")
        (doseq [n (range 1 (inc (Integer/parseInt (first args))))]
        (println (-fizzbuzz n)))))
    
(-main "40")
