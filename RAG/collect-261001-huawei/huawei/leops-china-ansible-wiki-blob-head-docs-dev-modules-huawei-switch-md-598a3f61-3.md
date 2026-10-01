---
id: collect-261001-huawei/huawei/leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61-3
title: "leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61.md
source_anchor: ""
source_lines: [325, 351]
sha256: 1dcc40ba084ad8c436fca01d9a39594f54838cbf9c5be532b7b04527273da32d
---

# leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61

    - debug: msg={{result.stdout_lines | d()}}
    - name: add vlan 900 and int 0/0/13
      hwos_command:
        sport: "{{ sport }}"
        shost: "{{ shost }}"
        suser: "{{ suser }}"
        spass: "{{ spass }}"
        save: true
        command: |
          system-view
          vlan 900
          quit
          interface GigabitEthernet 0/0/13
          port link-type access
          vlan 900
          port GigabitEthernet 0/0/13
          
    - name: display vlan 900
      hwos_command:
        sport: "{{ sport }}"
        shost: "{{ shost }}"
        suser: "{{ suser }}"
        spass: "{{ spass }}"
        command: display vlan 900
      register: result
 
    - debug: msg={{result.stdout_lines}}
