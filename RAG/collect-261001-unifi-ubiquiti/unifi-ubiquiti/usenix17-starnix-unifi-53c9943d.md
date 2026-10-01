---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/usenix17-starnix-unifi-53c9943d
title: "usenix17-starnix-unifi-53c9943d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/usenix17-starnix-unifi-53c9943d.md
source_anchor: ""
source_lines: [1, 36]
sha256: f1e06671a5f62b016a781ec4e3693fad79d0515da91f044b02d26a42ad014743
---

# usenix17-starnix-unifi-53c9943d

Manage the UniFi zone-based firewall (Policy Engine) and networks through the official UniFi Network Integration API (v1).
Existing UniFi collections target the legacy firewall-rules API or AP
provisioning; none cover the new zone firewall. This collection fills that gap
with idiomatic, idempotent, check_mode-aware modules.
Early development (0.1.0). Building order and progress live in
tasks/todo.md. The live API schema this is built against is
captured in docs/schema-discovery.md.
| Module | Manages | 
|---|---|
| unifi_site_info | Read-only: resolve/list sites | 
| unifi_firewall_zone | Firewall zones | 
| unifi_firewall_group | Traffic-matching lists (address/port groups) | 
| unifi_firewall_policy | Firewall policies | 
| unifi_firewall_policy_order | Policy evaluation order per zone pair | 
| unifi_network | Networks (incl. DHCP DNS) | 
- ansible-core >= 2.21
- A UniFi controller with an Integration API key (X-API-KEY ).
All modules share connection options (see the starnix.unifi.auth doc
fragment). Set them once per play with module_defaults:
- hosts: localhost
  gather_facts: false
  module_defaults:
    group/starnix.unifi.unifi:
      host: heimdal.starnix.net
      api_key: "{{ lookup('env', 'UNIFI_API_KEY') }}"
      site: default
  tasks:
    - name: Ensure the internal-DNS matching list exists
      starnix.unifi.unifi_firewall_group:
        name: internal-dns
        type: IPV4_ADDRESSES
        items:
          - { type: IP_ADDRESS, value: 192.168.1.53 }
          - { type: IP_ADDRESS, value: 10.17.89.53 }
        state: present
GPL-3.0-or-later. See LICENSE.
