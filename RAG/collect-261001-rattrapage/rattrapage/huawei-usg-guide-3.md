---
id: collect-261001-rattrapage/rattrapage/huawei-usg-guide-3
title: "Huawei USG — Guide CLI complet (USG6000 series, V500R005 / V600)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention", "license", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg_guide.md
source_anchor: ""
source_lines: [528, 684]
sha256: 71d38ab2f39ddedc7e6f033843cf7f8092fd6f04c14c7b075990a9309b5cd83c
---

# VRRP sur une interface (VIP partagée)
[USG] interface GigabitEthernet 0/0/1
[USG-GigabitEthernet0/0/1] vrrp vrid 1 virtual-ip 192.168.1.254
[USG-GigabitEthernet0/0/1] vrrp vrid 1 priority 120
[USG-GigabitEthernet0/0/1] vrrp vrid 1 preempt-mode timer delay 60
[USG-GigabitEthernet0/0/1] quit

display hrp state verbose
display vrrp
```

- Le lien heartbeat HRP doit être direct et dédié.
- VGMP : les groupes VRRP basculent ensemble avec le HRP.

---

## 20. Logs et supervision

```shell
# Journalisation locale
display logbuffer
display trapbuffer
display logbuffer level warnings

# Envoi vers un syslog distant
system-view
[USG] info-center enable
[USG] info-center loghost 192.168.1.100
[USG] info-center source default channel 2 log state on

# Sessions en cours (très utile en dépannage)
display firewall session table
display firewall session table source-ip 192.168.1.50
display firewall statistics

# SNMP
[USG] snmp-agent
[USG] snmp-agent community read cipher Public@123
[USG] snmp-agent sys-info version v2c
[USG] snmp-agent target-host trap address udp-domain 192.168.1.100 params securityname public v2c
```

- Les logs de sessions/politiques se configurent par règle
  (`session-logging` dans la règle security-policy).
- NetStream disponible pour l'analyse de flux (`ip netstream`).

---

## 21. Diagnostic

```shell
display version
display device
display cpu-usage
display memory-usage
display fan
display power
display interface brief
display ip routing-table
display arp all
display mac-address
display security-policy rule all
display nat-policy rule all
display ike sa
display ipsec sa
display hrp state verbose
ping -a 192.168.1.1 8.8.8.8
tracert 8.8.8.8

# Debug temps réel (attention en production !)
terminal monitor
terminal debugging
debugging ip packet
undo debugging all
```

---

## 22. Maintenance : fichiers, upgrade, licence

```shell
dir                         # contenu flash:
tftp 192.168.1.100 get USG6600-V600R007.bin
startup system-software flash:/USG6600-V600R007.bin
display startup
reboot                      # redémarre (confirmer)

# Sauvegarde / restauration de config
save
display saved-configuration
reset saved-configuration   # efface la config sauvegardée
# puis reboot -> retour config d'usine

# Licence
display license
# activation via web UI (fichier .dat) ou commande license active

# Mot de passe console perdu : BootWare au démarrage (Ctrl+B),
# menu « Clear password for console user ».
```

---

## 23. Exemple complet : USG en passerelle d'entreprise

```shell
system-view
[USG] sysname USG-SIEGE
# --- Interfaces & zones
[USG] interface GigabitEthernet 0/0/1
[USG-GigabitEthernet0/0/1] ip address 192.168.1.1 24
[USG-GigabitEthernet0/0/1] quit
[USG] interface GigabitEthernet 0/0/2
[USG-GigabitEthernet0/0/2] ip address 202.10.10.10 30
[USG-GigabitEthernet0/0/2] quit
[USG] firewall zone trust
[USG-zone-trust] add interface GigabitEthernet 0/0/1
[USG-zone-trust] quit
[USG] firewall zone untrust
[USG-zone-untrust] add interface GigabitEthernet 0/0/2
[USG-zone-untrust] quit
# --- Routage & DNS
[USG] ip route-static 0.0.0.0 0.0.0.0 202.10.10.9
[USG] dns resolve
[USG] dns server 8.8.8.8
# --- NAT Internet
[USG] nat-policy
[USG-policy-nat] rule name SNAT
[USG-policy-nat-rule-SNAT] source-zone trust
[USG-policy-nat-rule-SNAT] destination-zone untrust
[USG-policy-nat-rule-SNAT] source-address 192.168.1.0 24
[USG-policy-nat-rule-SNAT] action source-nat easy-ip
[USG-policy-nat-rule-SNAT] quit
[USG-policy-nat] quit
# --- Politique : LAN->WAN + UTM, admin vers local
[USG] security-policy
[USG-policy-security] rule name LAN_TO_WAN
[USG-policy-security-rule-LAN_TO_WAN] source-zone trust
[USG-policy-security-rule-LAN_TO_WAN] destination-zone untrust
[USG-policy-security-rule-LAN_TO_WAN] source-address 192.168.1.0 24
[USG-policy-security-rule-LAN_TO_WAN] profile av AV_DEFAULT
[USG-policy-security-rule-LAN_TO_WAN] action permit
[USG-policy-security-rule-LAN_TO_WAN] quit
[USG-policy-security] rule name ADMIN_TO_FW
[USG-policy-security-rule-ADMIN_TO_FW] source-zone trust
[USG-policy-security-rule-ADMIN_TO_FW] destination-zone local
[USG-policy-security-rule-ADMIN_TO_FW] action permit
[USG-policy-security-rule-ADMIN_TO_FW] quit
[USG-policy-security] quit
[USG] quit
save
```

---

*Fin du guide — ~400 lignes. Pour aller plus loin : les Command Reference
officiels par produit sont la référence exacte de chaque commande.*
