# optimize_CI_deterministic_builds
Dieses Repository enthält den Code für ein Tool, dass ich im Rahmen meiner Bachelorarbeit entwickelt habe (siehe [Repo der Bachelorarbeit](https://github.com/Zip-creations/BA_latex)).
[Hier](https://github.com/Zip-creations/BA_showcase) befindet sich ein Demoprojekt zur Anwendung des Tool. 

Tool can be build with
`go build -o testAuditor`
# Damit ein Projekt diese Tool benutzen kann, muss gelten:

- Das Test-Framework ist in der Lage einzelne Tests aus einer Suite gezielt ausführen
- Das Test-Framework benutzt JUnit XML als Ausgabeformat
- Es existiert ein Script, das alle vorhandenen Testcases findet & in einem spezifizierten XML Format ausgibt
- Es existiert ein zweites Script, dass die selben Namen für Testcases benutzt wie das erste Scripts, und ausgehend davon einzelne Tests gezielt ausführen kann
- Es wird eine Datei config.json im selben Verzeichnis wie die binary angelegt, nach diesem Format:
```
{
    "testDiscoveryPath": {
        "command": "path/to/example_testDiscovery.sh",
        "args": []
    },
    "testExecutionPath": {
        "command": "path/to/example_testExecution.sh",
        "args": []
    }
}
```
