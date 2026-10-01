---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-11
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [2124, 2357]
sha256: b8d8151d11a7159c11fdb65e3af51582b1c99ef3c5948d3fd108dd3a357f8ff3
---

# VRP — Le système d'exploitation transversal Huawei

```vrp
[USG6000]nat-policy
[USG6000-policy-nat]rule name NAT_SORTANT
[USG6000-policy-nat-rule-NAT_SORTANT]source-zone trust
[USG6000-policy-nat-rule-NAT_SORTANT]destination-zone untrust
[USG6000-policy-nat-rule-NAT_SORTANT]action source-nat easy-ip
```

## 111. Cas n°10 — DHCP qui ne distribue plus (S310/AR720)

**Symptômes** : les clients n'obtiennent plus d'adresse.

**Diagnostic** :
```vrp
<AR720>display ip pool                        # pools et adresses restantes
<AR720>display ip pool name POOL1 used        # (syntaxe à vérifier)
<AR720>display dhcp server statistics
<AR720>display logbuffer | include -i dhcp
```

**Causes** : pool épuisé (baux trop longs), `dhcp enable` désactivé, mauvaise
interface (`dhcp select global` vs `interface`), relais DHCP manquant entre
VLANs (`dhcp relay server-ip`).

## 112. Cas n°11 — Mot de passe perdu (sans autre accès)

Voir procédure complète sections 82-84. En résumé : console → reboot →
Ctrl+B → mot de passe BootROM → « Clear password for console user » →
« Boot with default mode » → redéfinir + `save`. **Coupure de service
obligatoire** : planifiez une fenêtre.

## 113. Cas n°12 — `save` impossible (flash plein)

**Symptômes** : `save` échoue avec une erreur d'espace.

**Diagnostic / solution** :
```vrp
<AR720>dir                    # "X bytes free" trop faible
<AR720>reset recycle-bin      # vide la corbeille
<AR720>dir
<AR720>delete /unreserved vieux_backup.zip
<AR720>save
```

Prévention : ne garder que 2 générations de `.cc` et 2 configs en flash,
tout le reste sur serveur.

## 114. Cas n°13 — Vitesse/duplex incohérents (négociation)

**Symptômes** : lien up mais débit anormal, erreurs CRC qui montent.

**Diagnostic** :
```vrp
<AR720>display interface GigabitEthernet 0/0/0 | include -i "speed|duplex|error|CRC"
```

**Solution** : forcer des deux côtés identiquement (ou auto des deux côtés,
jamais un mélange) :

```vrp
[AR720]interface GigabitEthernet 0/0/0
[AR720-GigabitEthernet0/0/0]speed 1000
[AR720-GigabitEthernet0/0/0]duplex full
[AR720-GigabitEthernet0/0/0]undo negotiation auto
```

## 115. Cas n°14 — VLAN qui ne passe pas sur un trunk

**Symptômes** : deux switchs reliés, un VLAN injoignable d'un côté.

**Diagnostic** :
```vrp
<S310>display vlan
<S310>display interface GigabitEthernet 0/0/24   # le trunk
<S310>display port vlan GigabitEthernet 0/0/24   # (à vérifier selon version)
```

**Causes** : VLAN non créé d'un côté, `port trunk allow-pass vlan` incomplet,
`port link-type` resté en `access`/`hybrid` par défaut, native VLAN
différente (PVID).

```vrp
[S310]interface GigabitEthernet 0/0/24
[S310-GigabitEthernet0/0/24]port link-type trunk
[S310-GigabitEthernet0/0/24]port trunk allow-pass vlan 10 20 30
```

## 116. Cas n°15 — Sessions VTY saturées

**Symptômes** : « Too many users » à la connexion SSH/Telnet.

**Diagnostic** :
```vrp
<AR720>display users          # VTY 0-4 tous occupés ?
```

**Solution** : demander aux collègues de se déconnecter, ou en console :
```vrp
<AR720>user-interface vty 5 14               # étendre la plage (si supporté)
```
Mieux : `idle-timeout` court (5-10 min) pour libérer les sessions fantômes.

## 117. Cas n°16 — NTP non synchronisé

**Symptômes** : `display ntp-service status` → `clock status: unsynchronized`.

**Diagnostic** :
```vrp
<AR720>display ntp-service sessions
<AR720>ping <ip-serveur-ntp>
<AR720>display current-configuration | include ntp-service
```

**Causes** : serveur injoignable (route/firewall UDP 123), `clock timezone`
incorrect (l'heure semble fausse mais c'est le fuseau), authentification NTP
requise côté serveur.

## 118. Cas n°17 — Après upgrade : patch manquant / comportement régressé

**Symptômes** : tout semble OK mais un bug corrigé par un `.pat` réapparaît.

**Diagnostic** :
```vrp
<AR720>display patch-information    # le patch est-il toujours actif ?
<AR720>display startup              # "Next startup patch package" renseigné ?
```

**Solution** : les patchs sont liés à une version ; après un upgrade majeur,
retéléchargez le patch correspondant à la NOUVELLE version et réappliquez
(`startup patch` + reboot si requis).

## 119. Cas n°18 — Lenteurs applicatives, pas de panne franche

**Symptômes** : « le réseau est lent » mais rien n'est down.

**Diagnostic** :
```vrp
<AR720>display interface brief          # InUti/OutUti : saturation ?
<AR720>display cpu-usage
<AR720>display ip routing-table         # la route prend-elle le bon chemin ?
<AR720>tracert <destination>           # latence par saut
<AR720>ping -c 100 -s 1400 <destination>  # pertes ?
<USG6000>display firewall session table | include <ip-client>  # sessions USG
```

**Pistes** : lien saturé (QoS à mettre en place), asymétrie de routage,
MTU/fragmentation (tester avec gros pings), session firewall qui expire,
boucle partielle.

---


## 120. Bonnes pratiques d'exploitation — nommage

Convention proposée pour le parc de Zelef :

```text
Sysname : <MODELE>-<SITE>-<ROLE>-<NUM>
  Ex. AR720-SIEGE-WAN-01, S310-ETAGE1-ACC-02, USG6000-SIEGE-FW-01

Description d'interface : <DESTINATION>_<USAGE>_<VLAN>
  Ex. description S310-ETAGE1_UPLINK_TRUNK
  Ex. description PC-BUREAU-042_ACC_VLAN10

VLAN : id + nom fonctionnel
  vlan 10 -> description USERS
  vlan 20 -> description SERVEURS
  vlan 99 -> description MGMT

Règles USG : <ACTION>_<SRC>_TO_<DST>_<SERVICE>
  Ex. ALLOW_LAN_TO_WAN_HTTP
```

Un nom doit permettre à l'astreinte de comprendre **sans ouvrir la doc**.

## 121. Documentation du parc — ce qu'il faut tenir à jour

```text
[ ] Plan d'adressage IP (sous-réseaux, passerelles, DHCP)
[ ] Matrice des VLANs (id, nom, usage, équipements concernés)
[ ] Topologie L2/L3 (ports d'interconnexion, trunks, Eth-Trunk)
[ ] Tableau des équipements (modèle, n° série, version VRP, rôle)
[ ] Mots de passe dans un coffre (jamais dans un fichier clair !)
[ ] Procédures : upgrade, password recovery, rollback (ce guide)
[ ] Journal des changements (qui, quoi, quand, pourquoi)
[ ] Sauvegardes de config datées (automatiques, section 98)
```

## 122. Sauvegardes planifiées — politique type

| Quoi | Fréquence | Où | Rétention |
|---|---|---|---|
| `display current-configuration` (script) | Quotidienne 02h00 | Serveur admin `/srv/backups/vrp/` | 30 j + 12 mensuels |
| `save` + copie `vrpcfg.zip` externe | Avant chaque changement | Serveur admin | Liée au changement |
| Fichier `.cc` en production | À chaque upgrade | Serveur admin (versionné) | Tant que l'OS est en prod + 1 |
| `display diagnostic-information` | Mensuelle | Serveur admin | 12 mois |

Testez la restauration **une fois par an** sur un équipement de lab :
une sauvegarde jamais testée n'est pas une sauvegarde.

## 123. Fenêtres de changement — rituel

```text
AVANT :
[ ] display version / display startup (état de départ noté)
[ ] Sauvegarde externe de la config
[ ] Commandes préparées à l'avance (copier-coller, pas d'impro)
[ ] Rollback écrit (les commandes undo / l'ancien .cc identifié)
[ ] Fenêtre validée, utilisateurs prévenus, console accessible

PENDANT :
[ ] Une modification à la fois, vérification après chacune
[ ] display this après chaque vue modifiée

APRÈS :
[ ] compare configuration -> relecture
[ ] save
[ ] display saved-configuration (échantillon)
[ ] Tests de non-régression (ping passerelles, display ospf peer, applis)
[ ] Journal des changements mis à jour
```

## 124. Supervision minimale à mettre en place

```text
[ ] SNMP (v3 de préférence) vers Zabbix/supervision : uptime, CPU, mémoire,
    état des interfaces, température
[ ] Syslog centralisé (tous les équipements -> même serveur, même NTP)
[ ] Traps SNMP pour : link up/down, changement de config, reboot
[ ] Ping de surveillance des IP de management
[ ] Alerte sur : CPU > 80 %, mémoire > 85 %, température, alarme active
[ ] Sauvegarde de config quotidienne + alerte si échec
```

