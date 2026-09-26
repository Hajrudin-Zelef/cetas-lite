---
id: collect-260926-mikrotik/mikrotik/how-to-access-routeros-api-in-python-using-the-routeros-library-minimal-example-techoverfl
title: "Extract the one identity string from the list of dictionaries"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/how-to-access-routeros-api-in-python-using-the-routeros-library-minimal-example-techoverflow.md
source_anchor: ""
source_lines: [1, 12]
sha256: 8e722129c7fa6f7c207b14cc19ed35db11051fd88c236b7f028c59aba2d23e43
---

# Extract the one identity string from the list of dictionaries

This example uses the routeros library from PyPI (GitHub) to access the MikroTik API and extract the system identity.

routeros_minimal_example.py

```
#!/usr/bin/env python3
from routeros import login
routeros = login('admin', 'abc123abc', '192.168.88.1')
output = routeros('/system/identity/print')
# Extract the one identity string from the list of dictionaries
print(output[0]['name'])
```
