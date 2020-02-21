#!/usr/bin/env python
import sys

path = ""
for line in sys.stdin:
    if line.startswith('/'):
        path = "paths." + line
    else:
        sys.stdout.write(path.replace(':', '.' + line.strip().replace(':', '.' + "operationId: TODO")))
