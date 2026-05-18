# optimize_CI_deterministic_builds
This repository contains both the latex files for my bachelor thesis, as well as the accompanying code of the practical implementation.

Tool can be build with
`go build -o testAuditor`
# Damit ein Projekt diese Tool benutzen kann, muss gelten:

- Das Test-Framework ist in der Lage einzelne Tests aus einer Suite gezielt ausführen
- Das Test-Framework benutzt JUnit XML als Ausgabeformat
- Es existiert ein Ordner, der alle von dem/den Testframeworks(s) erezugten JUnit-XML Dateien enthält
- Es existiert ein Script, das alle vorhandenen Testcases findet & in einem spezifizierten XML Format ausgibt
- Es wird eine Datei config.json im selben Verzeichnis wie die binary angelegt, nach diesem Format:
```
{
    "testDiscoveryPath": {
        "command": "path/to/sricpt.sh",
        "args": []
    },
    "jUnitXMLDirectory": "path/to/directory",
    "outputPath": "./out/report.xml"
}
```
<!-- - Es existiert ein zweites Script, dass das Format des ersten Scripts versteht, und ausgehend davon einzelne Tests gezielt ausführen kann -->
