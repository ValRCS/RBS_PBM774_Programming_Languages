(require '[clojure.string :as str])

;; (defn ask-name [] ; so no parameters are passed to this function
;;   (println "What is your name?")
;;   (let [name (read-line)] ; so name is bound to the value of the user's input, read-line reads a line of text from the standard input
;;     (println "Hello," name))
;;   )

; let's do it in two steps make a get-name function that returns the name
(defn get-name []
  (println "What is your name?")
  ;let's return the name split into words
  ;documentations for str/split: https://clojuredocs.org/clojure.string/split
  (str/split (read-line) #" ")
 )

;then let's have a function fancy-print that takes a name as an argument and prints a greeting
(defn fancy-print [name]
  (println "Hello," name))


;; (ask-name)

(fancy-print (get-name))