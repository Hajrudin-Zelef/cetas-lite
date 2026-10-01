---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-8
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [1490, 1662]
sha256: c324bcfc2e143515583e24d2e4d548085b7fcd5d7dbc5ab819873b5b896b12ad
---

# VRP — Le système d'exploitation transversal Huawei

| # | Cisco IOS | Huawei VRP | Notes |
|---|---|---|---|
| 1 | `show running-config` | `display current-configuration` | Config active (RAM) |
| 2 | `show startup-config` | `display saved-configuration` | Config sauvée (flash) |
| 3 | `show version` | `display version` | Version, uptime |
| 4 | `show ip interface brief` | `display ip interface brief` | État IP des interfaces |
| 5 | `show interfaces status` | `display interface brief` | État PHY/Protocol |
| 6 | `show interfaces Gi0/0` | `display interface GigabitEthernet 0/0/0` | Détail interface |
| 7 | `show ip route` | `display ip routing-table` | Table de routage |
| 8 | `show arp` | `display arp` | Table ARP |
| 9 | `show mac address-table` | `display mac-address` | Table MAC (switch) |
| 10 | `show vlan brief` | `display vlan` | VLANs |
| 11 | `show spanning-tree` | `display stp brief` | STP |
| 12 | `show ip ospf neighbor` | `display ospf peer` | Voisins OSPF |
| 13 | `show ip bgp summary` | `display bgp peer` | Pairs BGP |
| 14 | `show logging` | `display logbuffer` | Logs |
| 15 | `show processes cpu` | `display cpu-usage` | CPU |
| 16 | `show memory` / `show processes memory` | `display memory` | Mémoire |
| 17 | `show users` | `display users` | Sessions connectées |
| 18 | `show clock` | `display clock` | Horloge |
| 19 | `show ntp status` | `display ntp-service status` | NTP |
| 20 | `show inventory` | `display device elabel` | N° de série |
| 21 | `show file systems` / `dir` | `dir` | Contenu flash |
| 22 | `configure terminal` | `system-view` | Mode configuration |
| 23 | `end` / `Ctrl+Z` | `return` / `Ctrl+Z` | Retour user view |
| 24 | `exit` | `quit` | Remonter d'un niveau |
| 25 | `do show ...` | `run display ...` | Display depuis config |
| 26 | `interface Gi0/0` | `interface GigabitEthernet 0/0/0` | Vue interface |
| 27 | `ip address 1.1.1.1 255.255.255.0` | `ip address 1.1.1.1 24` | Masque en CIDR |
| 28 | `no shutdown` | `undo shutdown` | Activer l'interface |
| 29 | `shutdown` | `shutdown` | Couper l'interface |
| 30 | `description ...` | `description ...` | Identique |
| 31 | `hostname X` | `sysname X` | Nom d'hôte |
| 32 | `copy running-config startup-config` | `save` | Sauvegarder |
| 33 | `reload` | `reboot` | Redémarrer |
| 34 | `no <commande>` | `undo <commande>` | Annuler |
| 35 | `username X privilege 15 secret Y` | `local-user X password irreversible-cipher Y` + `local-user X privilege level 15` (vue aaa) | Utilisateur local |
| 36 | `line vty 0 4` | `user-interface vty 0 4` | Lignes VTY |
| 37 | `transport input ssh` | `protocol inbound ssh` (vue vty) | Protocole entrant |
| 38 | `ip ssh version 2` + `crypto key generate rsa` | `rsa local-key-pair create` + `stelnet server enable` | Activer SSH |
| 39 | `snmp-server community X RO` | `snmp-agent community read cipher X` (+ `snmp-agent`) | SNMP v2c |
| 40 | `ntp server 1.2.3.4` | `ntp-service unicast-server 1.2.3.4` | NTP |
| 41 | `ip route 0.0.0.0 0.0.0.0 1.2.3.4` | `ip route-static 0.0.0.0 0.0.0.0 1.2.3.4` | Route statique |
| 42 | `router ospf 1` | `ospf 1` (+ `router-id`) | OSPF |
| 43 | `network 1.0.0.0 0.255.255.255 area 0` | `area 0` puis `network 1.0.0.0 0.255.255.255` (vue area) | Annonce OSPF |
| 44 | `traceroute` | `tracert` | Traceroute |
| 45 | `debug ...` | `debugging ...` (+ `terminal debugging`) | Debug |
| 46 | `undebug all` | `undo debugging all` | Couper debugs |
| 47 | `show | include X` | `display ... \| include X` | Filtre include |
| 48 | `show | begin X` | `display ... \| begin X` | Filtre begin |
| 49 | `terminal length 0` | `screen-length 0 temporary` | Pas de pagination |
| 50 | `write erase` + `reload` | `reset saved-configuration` + `reboot` | Reset usine |

## 87. Pièges de conversion IOS → VRP

1. **`show` n'existe pas** : `show version` → erreur. Réflexe `display`.
   (Ou `command-alias mapping show display`, section 28.)
2. **Le masque en CIDR** : `ip address 192.168.1.1 24`, pas
   `255.255.255.0`.
3. **`save` obligatoire** : IOS `copy run start` est explicite ; en VRP
   l'oubli est encore plus traître car tout *semble* fonctionner jusqu'au
   reboot.
4. **Numérotation des interfaces** : `GigabitEthernet0/0/0` (slot/sous-carte/
   port), pas `Gi0/0`.
5. **OSPF** : le `network` se fait dans la **vue d'area**, pas dans la vue
   du processus.
6. **VLAN d'interface** : `port link-type access` + `port default vlan 10`
   (pas `switchport mode access` / `switchport access vlan 10`).
7. **Trunk** : `port link-type trunk` + `port trunk allow-pass vlan 10 20`.
8. **Agrégation** : `interface Eth-Trunk 1` + `trunkport GE0/0/1` (pas
   `channel-group`).
9. **Niveaux 0-15** : un `privilege level 15` VRP ≈ `privilege 15` IOS,
   mais la granularité par commande (`command-privilege`) n'a pas
   d'équivalent direct simple en IOS.

## 88. Fiche mémo VRP pour un cisconien pressé

```text
display ...        = show ...
system-view        = conf t
quit               = exit
return / Ctrl+Z    = end
undo X             = no X
save               = copy run start
reboot             = reload
sysname            = hostname
user-interface vty = line vty
```

---


## 89. SSH (STelnet) — configuration complète pas à pas

```vrp
[AR720]rsa local-key-pair create
  The key name will be: AR720_Host
  The range of public key size is (512 ~ 2048).
  NOTES: If the key modulus is greater than 512,
         it will take a few minutes.
  Input the bits in the modulus[default = 2048]:2048
  Generating keys...
  ..........................++++++
  ..........................++++++
  .........++++++++
  .........++++++++
[AR720]display rsa local-key-pair public      # vérifie la clé
[AR720]aaa
[AR720-aaa]local-user admin password irreversible-cipher MotDePasseFort123!
[AR720-aaa]local-user admin privilege level 15
[AR720-aaa]local-user admin service-type ssh
[AR720-aaa]quit
[AR720]ssh user admin authentication-type password
[AR720]ssh user admin service-type stelnet
[AR720]stelnet server enable
[AR720]ssh server-source -i Vlanif 10          # restreint l'interface d'écoute (optionnel)
[AR720]user-interface vty 0 4
[AR720-ui-vty0-4]authentication-mode aaa
[AR720-ui-vty0-4]protocol inbound ssh
[AR720-ui-vty0-4]user privilege level 3
[AR720-ui-vty0-4]idle-timeout 10 0
[AR720-ui-vty0-4]quit
[AR720]quit
<AR720>save
<AR720>display ssh server status              # vérification
```

Depuis le poste admin :

```bash
ssh admin@192.168.10.1
# Première connexion : vérifier l'empreinte de la clé !
```

> Si la clé RSA n'existe pas, `stelnet server enable` échoue ou SSH refuse
> les connexions : **toujours générer la clé d'abord**.

## 90. Telnet — à éviter, mais à connaître

```vrp
[AR720]telnet server enable
[AR720]user-interface vty 0 4
[AR720-ui-vty0-4]protocol inbound telnet
[AR720-ui-vty0-4]authentication-mode aaa
[AR720-ui-vty0-4]quit
```

Telnet transmet tout en clair (mots de passe inclus). En production :
**SSH uniquement** (`protocol inbound ssh`). Ne gardez Telnet que pour du
lab isolé ou un équipement ancien sans SSH.

## 91. SNMP v2c — supervision de base

```vrp
[AR720]snmp-agent
[AR720]snmp-agent community read cipher CommunauteLecture123
[AR720]snmp-agent community write cipher CommunauteEcriture456
[AR720]snmp-agent sys-info contact "Zelef - Service Systemes & Energies"
[AR720]snmp-agent sys-info location "Siege - Baie A1"
[AR720]snmp-agent sys-info version v2c
[AR720]snmp-agent target-host trap-hostname ZABBIX trap address udp-domain 192.168.100.60 params securityname CommunauteLecture123 v2c
[AR720]snmp-agent trap source Vlanif 10
```

Vérification depuis le serveur de supervision :

```bash
snmpwalk -v2c -c CommunauteLecture123 192.168.10.1 sysDescr
snmpwalk -v2c -c CommunauteLecture123 192.168.10.1 1.3.6.1.2.1.2.2.1.8  # ifOperStatus
```

## 92. SNMP v3 — la version sécurisée (recommandée)

