---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-55-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-55.md
source_anchor: ""
source_lines: [150, 166]
sha256: 55007c33bc932306c87ac5d5e2c647b91dbc969630c115cdd36df8c5c212adc9
---

# Summary

| Property | Description | 
|---|---|
| **user** (*string* ; Default: ) | system user to which the SSH key has been assigned | 
| **key-owner** (*string* ) | SSH key owner | 
| **key-type** (read-only*)* | key type | 
| **bits** (read-only*)* | key length | 

### Import private SSH key

On private SSH key import, must specify key file, system user to which SSH key will been assigned, optional it is possible to provide key passhrase and specify key owner.

| Property | Description | 
|---|---|
| **user** (*string* ; Default: ) | system user to which the SSH key has been assigned | 
| **key-owner** (*string* ) | SSH key owner | 
| **passphrase** *(string) sensitive* | key file passphrase | 
| **private-key-file** (*string* ) | file name in the router's root directory containing private key |
