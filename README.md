# optimize_CI_deterministic_builds
This repository contains both the latex files for my bachelor thesis, as well as the accompanying code of the practical implementation.

# Damit ein Projekt diese Tool benutzen kann, muss gelten:

- Das Test-Framework ist in der Lage einzelne Tests aus einer Suite gezielt ausführen
- Das Test-Framework benutzt JUnit XML als Ausgabeformat
- Es existiert ein Script, das alle vorhandenen Testcases findet & in einem näher zu definierendem Format ausgibt
- Es existiert ein zweites Script, dass das Format des ersten Scripts versteht, und ausgehend davon einzelne Tests gezielt ausführen kann
