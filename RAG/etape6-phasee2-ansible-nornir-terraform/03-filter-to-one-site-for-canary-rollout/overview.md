---
id: etape6-phasee2-ansible-nornir-terraform/03-filter-to-one-site-for-canary-rollout/overview
title: "Filter to one site for canary rollout"
domain: filter-to-one-site-for-canary-rollout
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [340, 359]
section: "Filter to one site for canary rollout"
sha256: d05af9404ba0cb99607227b1687d78b8b35ad75d16dbf09451ed56243fc29c3f
---

# Filter to one site for canary rollout
canary = nr.filter(site="dc1")

def backup(task):
    # Structured facts + raw running config
    facts = task.run(task=napalm_get, getters=["facts", "interfaces"])
    cfg = task.run(task=napalm_cli, commands=["show running-config"])
    return {"facts": facts.result, "config": cfg.result}

results = nr.run(task=backup)
for host, res in results.items():
    if res.failed:
        print(f"FAILED: {host}")
    else:
        open(f"backups/{host}.cfg", "w").write(res.result["config"]["show running-config"])
```

### 15c. OpenTofu — NetBox-as-code + safe planning

```hcl
