---
id: etape6-phasee2-ansible-nornir-terraform/01-requirements-yml-pin-everything/overview
title: "requirements.yml — pin everything"
domain: requirements-yml-pin-everything
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [288, 298]
section: "requirements.yml — pin everything"
sha256: 594cbffd290cebf62938db3eb29548a93b6f28e677a7b96aa99c7bfec43d68d4
---

# requirements.yml — pin everything
collections:
  - name: cisco.ios
    version: ">=11.0.0,<12.0.0"
  - name: arista.eos
    version: ">=12.0.0,<13.0.0"
  - name: juniper.device
    version: ">=2.0.0,<3.0.0"
```

```yaml
