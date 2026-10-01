---
id: collect-261001-fortinet/fortinet/bradley-rose-pyfortinetapi-5c090740
title: "bradley-rose-pyfortinetapi-5c090740"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/bradley-rose-pyfortinetapi-5c090740.md
source_anchor: ""
source_lines: [1, 21]
sha256: 10d6ef2990753c3d27a41f691606612b91c3b5fc917ff1196efc617b9d2a6f90
---

# bradley-rose-pyfortinetapi-5c090740

This is a base Python wrapper for the FortiOS API. Create custom logic as necessary, including creating new CRUD actions in apiWrapper.py. Of note, this doesn't include any amount of inventory management. Integrate your source of truth for inventory management as you see fit.
To use apiWrapper.py:
import apiWrapper as fgApi
def actionsToPerform(*, api: dict):
    with api.FortiGate(**{
        "hostAddress": "<IP address of FortiGate management interface>",
        "port": "<HTTPS port where web interface is accessible>",
        # If authenticating via apiKey:
        "apiKey": "<apiKey>".
        # Or, if authenticating via username/password:
        "username": "<username>",
        "password": "<password>"
    }) as fortiGate:
        # Syntax: result = fortiGate.<functionFrom_apiWrapper.py>
        # Example:
        result = fortiGate.getUserGroup(name="<groupName>")
        return result
def main():
    resultFromApiCall = actionsToPerform(api = fgApi)
if __name__ == "__main__":
    main()
