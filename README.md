# optimize_CI_deterministic_builds
Dieses Repository enthält den Code für ein Tool, dass ich im Rahmen meiner Bachelorarbeit entwickelt habe (siehe [Repo der Bachelorarbeit](https://github.com/Zip-creations/BA_latex)).<br>
[Hier](https://github.com/Zip-creations/BA_showcase) befindet sich ein Demoprojekt zur Anwendung des Tools. 

Tool can be build with
`go build -o testAuditor`
# Damit ein Projekt diese Tool benutzen kann, muss gelten:

- Das Testframework ist in der Lage einzelne Testcases gezielt ausführen
- Das Testframework benutzt JUnit XML als Ausgabeformat
- Das Testframework identifiziert einen Testcase im JUnit-XML durch die Kombination der Attribute `name` und `classname`

testAuditor erwartet auf `stdin` die Menge aller Testcases innerhalb des Projekts, zusammen mit einer Menge von JUnit-XML report aus bisher abgeschlossenen Testläufen in einem [hier](./src/cmd/testAuditor/examples/testInput.xml) spezifizierten XML-Format.<br>
testAuditor gibt dann eine Liste von `qualifiedName` zurück, welche die Menge aller Testcases beschreibt, die bisher noch nicht in den angegeben reports ausgefürt wurden. 