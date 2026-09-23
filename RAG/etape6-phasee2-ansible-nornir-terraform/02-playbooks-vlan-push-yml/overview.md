---
id: etape6-phasee2-ansible-nornir-terraform/02-playbooks-vlan-push-yml/overview
title: "playbooks/vlan_push.yml"
domain: playbooks-vlan-push-yml
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [299, 339]
section: "playbooks/vlan_push.yml"
sha256: 3784742d5b42deeafedd73644d7e2ab743ccadb49594938c194e6542a4599200
---

# playbooks/vlan_push.yml
- name: Push VLAN 200 to access switches
  hosts: access_switches
  gather_facts: false
  tasks:
    - name: Render intended VLAN config (dry-run artifact)
      ansible.builtin.set_fact:
        rendered: "{{ lookup('ansible.builtin.template', 'templates/vlan.j2') }}"

    - name: Apply VLANs on IOS
      cisco.ios.ios_vlans:
        config:
          - name: "{{ vlan_name }}"
            vlan_id: "{{ vlan_id }}"
            state: active
        state: merged
      when: ansible_network_os == 'cisco.ios.ios'
      check_mode: "{{ dry_run | default(true) }}"

    - name: Apply VLANs on EOS
      arista.eos.eos_vlans:
        config:
          - name: "{{ vlan_name }}"
            vlan_id: "{{ vlan_id }}"
        state: merged
      when: ansible_network_os == 'arista.eos.eos'

    - name: Apply VLANs on Junos (juniper.device FQCN)
      juniper.device.junos_config:
        src: templates/vlan_junos.j2
        comment: "VLAN {{ vlan_id }} via Ansible"
      when: ansible_network_os == 'juniper.junos.junos'
```

### 15b. Nornir — parallel backup with NAPALM (Python)

```python
from nornir import InitNornir
from nornir_napalm.plugins.tasks import napalm_cli, napalm_get

nr = InitNornir(config_file="config.yaml")
