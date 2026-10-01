---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/o-x-l-opnsense-api-client-blob-head-docs-source-usage-1-intro-rst-95420deb
title: "o-x-l-opnsense-api-client-blob-head-docs-source-usage-1-intro-rst-95420deb"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/o-x-l-opnsense-api-client-blob-head-docs-source-usage-1-intro-rst-95420deb.md
source_anchor: ""
source_lines: [1, 19]
sha256: 8c3f1a8e5510540744655cc36903b3596e3561fb134ebb987ddde360f9dc959d
---

# o-x-l-opnsense-api-client-blob-head-docs-source-usage-1-intro-rst-95420deb

This is a Python3 client for interacting with the official OPNsense API.

It enables easy management and automation of OPNsense firewalls.

The base-code is a Fork of this OPNsense Ansible-Collection that was refactored for use within raw Python.

This can be useful if you want to automate your Infrastructure and do not use Ansible.

`pip install oxl-opnsense-client````
from oxl_opnsense_client import Client
with Client(
    firewall='192.168.10.20',
    port=443,  # default
    credential_file='/tmp/.opnsense.txt',
    # token='0pWN/C3tnXem6OoOp0zc9K5GUBoqBKCZ8jj8nc4LEjbFixjM0ELgEyXnb4BIqVgGNunuX0uLThblgp9Z',
    # secret='Vod5ug1kdSu3KlrYSzIZV9Ae9YFMgugCIZdIIYpefPQVhvp6KKuT7ugUIxCeKGvN6tj9uqduOzOzUlv',
) as c:
    c.test()
```
