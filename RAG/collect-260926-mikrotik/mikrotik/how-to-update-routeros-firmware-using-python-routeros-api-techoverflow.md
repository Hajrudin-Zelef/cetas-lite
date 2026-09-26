---
id: collect-260926-mikrotik/mikrotik/how-to-update-routeros-firmware-using-python-routeros-api-techoverflow
title: "Update RouterOS firmware using Python RouterOS-api"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/how-to-update-routeros-firmware-using-python-routeros-api-techoverflow.md
source_anchor: ""
source_lines: [1, 73]
sha256: da2d1ed3a2c3cedaf4f07819fccc0d2e7cfe1650144130b2379f3cb151a5cfe3
---

# Update RouterOS firmware using Python RouterOS-api

Use the following Python code to update the RouterOS firmware using the Python RouterOS-api package:

update_routeros_fw.py

```
# Update RouterOS firmware using Python RouterOS-api
print(f"Updating RouterOS firmware")
output = api.get_resource('/system/package/update').call('install')
if update_output[-1]["status"] == 'System is already up to date':
    print("Firmware is already up to date")
# Print output (for this example)
print(output)
```
### Example output (if firmware can be updated)

routeros_fw_update_output.txt

```
[{'channel': 'stable',
  'installed-version': '7.13.2',
  'status': 'finding out latest version...',
  '.section': '0'},
 {'channel': 'stable',
  'installed-version': '7.13.2',
  'latest-version': '7.15.3',
  'status': 'Downloaded 22% (2.6MiB)',
  '.section': '1'},
 {'channel': 'stable',
  'installed-version': '7.13.2',
  'latest-version': '7.15.3',
  'status': 'Downloaded 62% (7.2MiB)',
  '.section': '2'},
 {'channel': 'stable',
  'installed-version': '7.13.2',
  'latest-version': '7.15.3',
  'status': 'Downloaded, rebooting...',
  '.section': '3'}]
```
### Example output (if firmware is already up-to-date)

routeros_fw_uptodate_output.txt

```
[{'channel': 'stable',
  'installed-version': '7.15.3',
  'status': 'finding out latest version...',
  '.section': '0'},
 {'channel': 'stable',
  'installed-version': '7.15.3',
  'latest-version': '7.15.3',
  'status': 'System is already up to date',
  '.section': '1'}]
```
## Full example

update_routeros_fw_full.py

```
#!/usr/bin/env python3
import routeros_api
connection = routeros_api.RouterOsApiPool(
    '192.168.88.1', username='admin', password='L9TYAUV7R8',
    plaintext_login=True
)
api = connection.get_api()
# Update RouterOS firmware using Python RouterOS-api
print(f"Updating RouterOS firmware")
output = api.get_resource('/system/package/update').call('install')
if update_output[-1]["status"] == 'System is already up to date':
    print("Firmware is already up to date")
# Print output (for this example)
print(output)
```
