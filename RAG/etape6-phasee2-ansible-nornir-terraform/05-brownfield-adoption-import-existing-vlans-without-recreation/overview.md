---
id: etape6-phasee2-ansible-nornir-terraform/05-brownfield-adoption-import-existing-vlans-without-recreation/overview
title: "Brownfield adoption: import existing VLANs without recreation"
domain: brownfield-adoption-import-existing-vlans-without-recreation
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [379, 389]
section: "Brownfield adoption: import existing VLANs without recreation"
sha256: a74a16e71ad81816eeb4e7c8fcb1c9d3206c9ab70aceb620c72eb3b756a79ceb
---

# Brownfield adoption: import existing VLANs without recreation
import {
  for_each = toset(["100", "200"])
  to = netbox_vlan.this[each.key]
  id = each.key
}
```

### 15d. EDA rulebook sketch — flap-aware remediation

```yaml
