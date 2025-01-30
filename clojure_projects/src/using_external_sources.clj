(ns using-external-sources
  (:require [clojure.string :as str]))

; now let's use the clojure.string library to capitalize a string
(defn -main
    [& args]
    (if (empty? args)
      (println "Please provide a string as an argument")
      ;so we wrap our two println calls in a vector and pass it to the vec function
      (vec [(println (str/capitalize (first args)))
       (println (str/upper-case (second args)))])
))

(-main "hello Rbs" "hello Rbs")
