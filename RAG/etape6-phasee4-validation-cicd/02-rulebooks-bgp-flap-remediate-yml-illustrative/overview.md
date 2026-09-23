---
id: etape6-phasee4-validation-cicd/02-rulebooks-bgp-flap-remediate-yml-illustrative/overview
title: "rulebooks/bgp_flap_remediate.yml (illustrative)"
domain: rulebooks-bgp-flap-remediate-yml-illustrative
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [743, 769]
section: "rulebooks/bgp_flap_remediate.yml (illustrative)"
sha256: b8008bf27b6e3a050eb7cd38b8a035df1324c75fb81a239df6e46bdaf679b69e
---

# rulebooks/bgp_flap_remediate.yml (illustrative)
- name: Remediate BGP flaps
  hosts: all
  sources:
    - name: alerts
      ansible.eda.webhook:
        host: 0.0.0.0
        port: 5000
  rules:
    - name: flap storm
      condition: event.alert.labels.alertname == "BgpSessionDown"
      throttle:
        once_within: 30 minutes
        group_by_attributes: [ "event.alert.labels.peer" ]
      action:
        run_workflow_template:
          name: "NET - BGP flap triage"
          job_args:
            extra_vars:
              peer: "{{ event.alert.labels.peer }}"
```

- Illustrative only `[analysis]`; follows Red Hat EDA patterns `[secondary]`.

### E4.16.8 Batfish pybatfish sketch (illustrative, not run)

```python
