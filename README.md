### General Info

This little go program does two basic things:
1 - Searches for an element in an HTML file using its id value
2 - Searches for any similar elements in another HTML file using the attributes of the first founded element

### Usage

Run the following command:

./main -originalHTML=./sample-0-origin.html -diffCaseHTML=./sample-1-evil-gemini.html (using the executable)

or

go run main.go -originalHTML=./sample-0-origin.html -diffCaseHTML=./sample-1-evil-gemini.html

you can also add the -searchedElement   flag to search for another desired element

-originalHTML flag contains the path to the origin HTML, in which we will search for the element's id and obtain it
-diffCaseHTML flag contains the path to the target HTML, in which we will search for any similar elements to the first one.

In the standard output we will see each similar element obtained alongside its path in the HTML path


### Notes
As im writing this readme it's been 3 hours since I begun the test. Various things can be improved but I know the clock's ticking.
I've finished most of this in little more than 2 hours but the last part (presenting the path) delayed me.
One of the improvements would be reverting the order of the tags presented for each element, and also indicating the relative position of each one.
(Now its something like "div div" instead of "div[2] div[0]" and so on).
Another improvement is to delete redundant copies of elements founded. This could be achieved easily by establishing some parameter of equality between elements.
Other improvement would be to separate all nicely in different packages, separating the functions used for searching in the HTML docs from the main package.
The whole functionality could fit in a service with appropiate endpoints, API, middlewares, logging, metrics,storage,containerization, etc.
