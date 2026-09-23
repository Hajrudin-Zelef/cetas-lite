---
id: etape6-phasee2-ansible-nornir-terraform/06-rulebooks-interface-flap-yml/overview
title: "rulebooks/interface_flap.yml"
domain: rulebooks-interface-flap-yml
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [390, 415]
section: "rulebooks/interface_flap.yml"
sha256: 7ebb19ad54b37b8b9e51d9dd6d04402db737a28fa391a4d6ce1884bbe9c5b125
---

# rulebooks/interface_flap.yml
- name: Remediate flapping access ports
  hosts: all
  sources:
    - name: monitoring_webhook
      # webhook source plugin
  rules:
    - name: Port flapped 3x in 5 minutes
      condition: event.interface_status == "flapping"
      # debounce: fire at most once per 10 minutes per interface
      throttle:
        once_within: 10 minutes
        group_by_attributes:
          - event.device
          - event.interface
      action:
        run_job_template:
          name: "Bounce port with pre/post checks"
          job_args:
            device: "{{ event.device }}"
            interface: "{{ event.interface }}"
```

### 15e. CI gate — plan-only validation (GitHub Actions sketch)

```yaml
