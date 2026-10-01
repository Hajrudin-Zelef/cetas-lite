---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-13
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [2251, 2386]
sha256: d1587dd2da064ea17e46fff33d8ec3c388d0fdb3b3224c732c8a07095c70a636
---

# --- VLAN et LAN ---
[AGENCE-DAKAR-AR720]vlan batch 10 20 30
[AGENCE-DAKAR-AR720]interface Vlanif 10
[AGENCE-DAKAR-AR720-Vlanif10]ip address 192.168.10.1 255.255.255.0
[AGENCE-DAKAR-AR720-Vlanif10]dhcp select global
[AGENCE-DAKAR-AR720-Vlanif10]quit
[AGENCE-DAKAR-AR720]interface Vlanif 20
[AGENCE-DAKAR-AR720-Vlanif20]ip address 192.168.20.1 255.255.255.0
[AGENCE-DAKAR-AR720-Vlanif20]dhcp select global
[AGENCE-DAKAR-AR720-Vlanif20]quit
[AGENCE-DAKAR-AR720]interface Vlanif 30
[AGENCE-DAKAR-AR720-Vlanif30]ip address 192.168.30.1 255.255.255.0
[AGENCE-DAKAR-AR720-Vlanif30]quit
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/2
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]port link-type access
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]port default vlan 10
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]stp edged-port enable
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/2]quit
# GE0/0/2 en accès VLAN 10 ; les autres ports à répartir selon le brassage.

# --- DHCP ---
[AGENCE-DAKAR-AR720]dhcp enable
[AGENCE-DAKAR-AR720]ip pool lan-bureautique
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]gateway-list 192.168.10.1
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]network 192.168.10.0 mask 255.255.255.0
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]dns-list 192.168.10.1 8.8.8.8
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]excluded-ip-address 192.168.10.1 192.168.10.20
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]lease day 1 hour 0 minute 0
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]static-bind ip-address 192.168.10.21 mac-address aaaa-bbbb-cccc
[AGENCE-DAKAR-AR720-ip-pool-lan-bureautique]quit
[AGENCE-DAKAR-AR720]ip pool lan-invites
[AGENCE-DAKAR-AR720-ip-pool-lan-invites]gateway-list 192.168.20.1
[AGENCE-DAKAR-AR720-ip-pool-lan-invites]network 192.168.20.0 mask 255.255.255.0
[AGENCE-DAKAR-AR720-ip-pool-lan-invites]dns-list 192.168.20.1 8.8.8.8
[AGENCE-DAKAR-AR720-ip-pool-lan-invites]lease day 0 hour 6 minute 0
[AGENCE-DAKAR-AR720-ip-pool-lan-invites]quit
# Bail court pour les invités (6 h), long pour la bureautique (1 j).

# --- WAN PPPoE ---
[AGENCE-DAKAR-AR720]dialer-rule
[AGENCE-DAKAR-AR720-dialer-rule]dialer-rule 1 ip permit
[AGENCE-DAKAR-AR720-dialer-rule]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]dialer user agence.dakar@fai.exemple
[AGENCE-DAKAR-AR720-Dialer1]dialer-group 1
[AGENCE-DAKAR-AR720-Dialer1]dialer bundle 1
[AGENCE-DAKAR-AR720-Dialer1]ppp chap user agence.dakar@fai.exemple
[AGENCE-DAKAR-AR720-Dialer1]ppp chap password cipher MotDePasseFAIFictif
[AGENCE-DAKAR-AR720-Dialer1]ppp pap local-user agence.dakar@fai.exemple password cipher MotDePasseFAIFictif
[AGENCE-DAKAR-AR720-Dialer1]ip address ppp-negotiate
[AGENCE-DAKAR-AR720-Dialer1]tcp adjust-mss 1400
[AGENCE-DAKAR-AR720-Dialer1]quit
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/0
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]pppoe-client dial-bundle-number 1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/0]quit

# --- WAN secours ---
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]ip address dhcp-alloc
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]quit
# La box 4G donne 192.168.8.0/24, passerelle 192.168.8.1.

# --- NAT ---
[AGENCE-DAKAR-AR720]acl number 2000
[AGENCE-DAKAR-AR720-acl-basic-2000]rule 5 permit source 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2000]rule 10 permit source 192.168.20.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2000]quit
[AGENCE-DAKAR-AR720]acl number 3000
[AGENCE-DAKAR-AR720-acl-adv-3000]rule 5 deny ip source 192.168.10.0 0.0.0.255 destination 192.168.0.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3000]rule 10 permit ip source 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3000]rule 15 permit ip source 192.168.20.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-adv-3000]quit
[AGENCE-DAKAR-AR720]interface Dialer 1
[AGENCE-DAKAR-AR720-Dialer1]nat outbound 3000
[AGENCE-DAKAR-AR720-Dialer1]quit
[AGENCE-DAKAR-AR720]interface GigabitEthernet 0/0/1
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]nat outbound 3000
[AGENCE-DAKAR-AR720-GigabitEthernet0/0/1]quit
# ACL 3000 : le trafic vers le siège n'est PAS naté (tunnel IPSec), le reste oui.
# NAT appliqué sur les DEUX interfaces WAN (sinon pas d'Internet en secours).

# --- Basculement NQA/track ---
[AGENCE-DAKAR-AR720]nqa test-instance admin pppoe-check
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]test-type icmp
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]destination-address ipv4 8.8.8.8
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]frequency 10
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]probe-count 3
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]start now
[AGENCE-DAKAR-AR720-nqa-admin-pppoe-check]quit
[AGENCE-DAKAR-AR720]track 1 nqa admin pppoe-check
[AGENCE-DAKAR-AR720]ip route-static 0.0.0.0 0.0.0.0 Dialer 1 track 1
[AGENCE-DAKAR-AR720]ip route-static 0.0.0.0 0.0.0.0 192.168.8.1 preference 100

# --- Zones de sécurité ---
[AGENCE-DAKAR-AR720]firewall zone trust
[AGENCE-DAKAR-AR720-zone-trust]add interface Vlanif 10
[AGENCE-DAKAR-AR720-zone-trust]add interface Vlanif 30
[AGENCE-DAKAR-AR720-zone-trust]quit
[AGENCE-DAKAR-AR720]firewall zone untrust
[AGENCE-DAKAR-AR720-zone-untrust]add interface Dialer 1
[AGENCE-DAKAR-AR720-zone-untrust]add interface GigabitEthernet 0/0/1
[AGENCE-DAKAR-AR720-zone-untrust]quit
[AGENCE-DAKAR-AR720]firewall zone dmz
[AGENCE-DAKAR-AR720-zone-dmz]add interface Vlanif 20
[AGENCE-DAKAR-AR720-zone-dmz]quit

# --- Politiques ---
[AGENCE-DAKAR-AR720]security-policy
[AGENCE-DAKAR-AR720-policy-security]rule name LAN-vers-Internet
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]source-zone trust
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]destination-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]source-address 192.168.10.0 24
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-LAN-vers-Internet]quit
[AGENCE-DAKAR-AR720-policy-security]rule name Invites-vers-Internet
[AGENCE-DAKAR-AR720-policy-security-rule-Invites-vers-Internet]source-zone dmz
[AGENCE-DAKAR-AR720-policy-security-rule-Invites-vers-Internet]destination-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-Invites-vers-Internet]source-address 192.168.20.0 24
[AGENCE-DAKAR-AR720-policy-security-rule-Invites-vers-Internet]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-Invites-vers-Internet]quit
[AGENCE-DAKAR-AR720-policy-security]rule name IKE-vers-local
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]source-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]destination-zone local
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]service protocol udp destination-port 500
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]service protocol udp destination-port 4500
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-IKE-vers-local]quit
[AGENCE-DAKAR-AR720-policy-security]rule name Admin-SSH-LAN
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]source-zone trust
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]destination-zone local
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]service protocol tcp destination-port 22
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]action permit
[AGENCE-DAKAR-AR720-policy-security-rule-Admin-SSH-LAN]quit
[AGENCE-DAKAR-AR720-policy-security]quit
# Les invités (dmz) vont sur Internet mais n'ont AUCUNE règle vers trust : isolés.

