Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit instead of a loop.

Exercise 5.2: Write a function to populate a mapping from element names - p, div, span, and so on- to the number of elments with that name in an HTML document tree.

Exercise 5.3: Write a function to print the contents of all text nodes in an HTML document tree. Do not descend into <script> or <style> elements, since their contents are not visible in a web browser.

Exercise 5.4: Extend the visit function so that it extracts other kinds of links from teh document, such as images, scripts, and style sheets.

Exercise 5.5: Implement countWordsAndImages. (See Exercise 4.9 for word-splitting)

Exercise 5.6: Modify the corner function in gopl.io/ch3/surface to use named results and a bare return statement.


Exercise 5.7: Develop startElement and endElement into a general HTML pretty-printer. Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use short forms like <img/> instead of <img></img> when an element has no children. Write a test to ensure that the output can be parsed successfully.


Exercise 5.8: Modify forEachNode so that the pre and post functions return a boolean result indicateing whether to continue the traversal. Use it to write a function ElementByID with the following signature that finds the first HTML element with the specified id attribute. The function should stop the traversal as soon as a match is found.

`func ElementByID(doc *html.Node, id string) *html.Node`

Exercise 5.9: Write a function expand(s string, f func(string) string) string that replace each substring "$foo" within s by the text returned by f("foo").


Exercise 5.10: Rewrite topoSort to use maps insteaad of slices and eliminate the initial sort. Verify that the results, though nondeterministic, are valid topological orderings.

Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument.

Exercise 5.16: Write a variadic version of strings.Join.

Exercise 5.17: Write a variadic function ElementsByTagName that, given an HTML node tree and zero or more names, returns all the elements that match one of those names. Here are two example calls:
```
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

images := ElementByTagName(doc, "img")
headings := ElementByTagName(doc, "h1", "h2", "h3", "h4")
```