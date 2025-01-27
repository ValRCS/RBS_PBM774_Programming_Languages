(ns factorial)
; so namespace should match the file name

; let's make a factorial function
; the function name could be anything, but it's a convention to start with a hyphen
(defn -factorial
  [n]
  (if (= n 0)
    1
    (* n (-factorial (- n 1)))))

; let's test it by printing value for 5

;; (println (-factorial 5))

; let's make a function that reads arguments from the command line

(defn -main
  [& args]
  (if (empty? args)
    (println "Please provide a number as an argument")
    (println (-factorial (Integer/parseInt (first args))))))

; let's test it by running the program without any arguments

(-main "10")

; let's read command line arguments and run them with main



