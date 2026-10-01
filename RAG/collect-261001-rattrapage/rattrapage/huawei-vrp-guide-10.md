---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-10
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [1915, 2123]
sha256: 503ab60affa42ce9ab8da7455735c225d60b8228b7a6faf8215d9e54ed0198bb
---

# VRP — Le système d'exploitation transversal Huawei

Astuce parsing : les sorties VRP sont stables d'une version à l'autre pour
les commandes de base — parsez avec des regex simples (`^sysname\s+(\S+)`,
`^(\S+) +up +up` dans `display interface brief`), ou utilisez TextFSM/
ntc-templates (le template `huawei_display_version` existe dans
ntc-templates — à vérifier pour votre version exacte).

---


## 101. Méthodologie de dépannage VRP — les 6 réflexes

```text
1. display version        -> quel boîtier, quelle version, quel uptime ?
2. display logbuffer      -> que s'est-il passé juste avant ?
3. display interface brief -> les liens sont-ils up/up ?
4. display ip routing-table -> la route existe-t-elle ?
5. ping / tracert         -> où s'arrête le trafic ?
6. display current-configuration | include <objet> -> la config est-elle là ?
```

Toujours du plus bas vers le plus haut (physique → liaison → IP →
routage → applicatif). 70 % des incidents se résolvent aux étapes 2 et 3.

---

## 102. Cas n°1 — Interface « down/down »

**Symptômes** : `display interface brief` montre `down/down` ; pas de ping.

**Diagnostic** :
```vrp
<AR720>display interface GigabitEthernet 0/0/1
  GigabitEthernet0/0/1 current state : DOWN
  Line protocol current state : DOWN
<AR720>display logbuffer | include GE0/0/1
```
**Causes probables** : câble débranché/cassé, SFP défectueux ou incompatible,
port distant éteint, `shutdown` configuré (`*down`).

**Solution** :
```vrp
[AR720]interface GigabitEthernet 0/0/1
[AR720-GigabitEthernet0/0/1]display this     # vérifier l'absence de "shutdown"
[AR720-GigabitEthernet0/0/1]undo shutdown
```
Vérifier le voyant du port, tester un autre câble/SFP, vérifier le distant.

## 103. Cas n°2 — Configuration perdue après reboot (`save` oublié)

**Symptômes** : après une coupure, l'équipement a « oublié » les modifs de
la veille ; `display current-configuration` ne montre plus les changements.

**Diagnostic** :
```vrp
<AR720>display startup
  Startup saved-configuration file: flash:/vrpcfg.zip
<AR720>display saved-configuration | include <objet-manquant>   # absent = jamais sauvé
```

**Solution** : reconfigurer + `save`. **Prévention** :
- Checklist de fin de session (section 70).
- `autosave interval on` en filet de sécurité.
- Sauvegarde externe automatisée (section 98) : même sans `save`, la config
  du matin est récupérable dans le backup.

## 104. Cas n°3 — SSH refuse la connexion

**Symptômes** : `ssh admin@192.168.10.1` → « Connection refused » ou timeout.

**Diagnostic** :
```vrp
<AR720>display ssh server status              # serveur actif ?
<AR720>display current-configuration | include stelnet
<AR720>display rsa local-key-pair public      # clé présente ?
<AR720>display users                          # VTY saturés (0-4 occupés) ?
<AR720>display current-configuration | include "user-interface vty" -A 6
```

**Causes / solutions** :
| Cause | Solution |
|---|---|
| Clé RSA absente | `rsa local-key-pair create` |
| `stelnet server enable` manquant | l'ajouter + `save` |
| `protocol inbound telnet` seulement | passer à `protocol inbound ssh` |
| Utilisateur sans `service-type ssh` | `local-user X service-type ssh` |
| ACL sur VTY qui bloque votre IP | ajuster la règle ACL |
| Les 5 VTY occupés | `display users` puis libérer / augmenter la plage |

## 105. Cas n°4 — Upgrade qui échoue (fichier rejeté ou boot en échec)

**Symptômes** : `startup system-software` refuse le fichier, ou l'équipement
boucle au boot après `reboot`.

**Diagnostic** :
```vrp
<AR720>dir                    # taille du .cc : complète ? (comparer à la source)
<AR720>display startup        # le bon fichier est-il en "Next startup" ?
```

**Causes / solutions** :
- Fichier tronqué (TFTP interrompu) → retransférer, vérifier la taille à
  l'octet près, idéalement la somme MD5.
- Mauvais fichier (modèle différent) → vérifier la référence exacte
  (`AR720-...cc` vs `S310-...cc` : NON interchangeables).
- Espace flash insuffisant → nettoyer (`reset recycle-bin`, vieux `.cc`).
- Boot en échec → BootROM (Ctrl+B) → recharger l'ancien `.cc` (section 80).

## 106. Cas n°5 — CPU à 100 %

**Symptômes** : lenteurs CLI, pertes de paquets, `display cpu-usage` > 90 %.

**Diagnostic** :
```vrp
<AR720>display cpu-usage
<AR720>display cpu-usage history 60m          # pic récent ou permanent ?
<AR720>display logbuffer | include -i "attack|storm|loop"
<AR720>display interface brief                # un port avec erreurs qui explosent ?
```

**Causes fréquentes** : boucle de switching (broadcast storm), debug oublié
actif (`display debugging`), attaque (SYN flood), processus SNMP qui
scanne trop vite.

**Solution** :
```vrp
<AR720>undo debugging all
[AR720]undo terminal debugging
```
Puis : casser la boucle (débrancher un lien redondant, vérifier STP :
`display stp brief`), ajuster la supervision, activer `cpu-defend`
(`display cpu-defend policy` — à vérifier selon modèle).

## 107. Cas n°6 — Boucle réseau / tempête de broadcast

**Symptômes** : réseau saturé, voyants qui clignotent frénétiquement, CPU haut
sur les switchs.

**Diagnostic** :
```vrp
<S310>display interface brief    # InUti/OutUti proches de 100 % sur plusieurs ports
<S310>display stp brief          # ports en FORWARDING qui devraient être bloqués ?
<S310>display mac-address        # une MAC qui "flappe" entre plusieurs ports ?
<S310>display logbuffer | include -i "loop|storm"
```

**Solution** : identifier le lien fautif (débranchement séquentiel en
fenêtre de maintenance), activer/durcir STP (`stp enable`, `stp mode mstp`),
configurer `loopback-detect` / `storm-control` sur les ports d'accès :

```vrp
[S310]interface GigabitEthernet 0/0/5
[S310-GigabitEthernet0/0/5]storm-control broadcast min-rate 1000 max-rate 2000
[S310-GigabitEthernet0/0/5]quit
```

## 108. Cas n°7 — Voisin OSPF bloqué (pas « Full »)

**Symptômes** : `display ospf peer` montre `Init`, `ExStart` ou `Exchange`
durable, ou aucun voisin.

**Diagnostic** :
```vrp
<AR720>display ospf peer
<AR720>display ospf error
<AR720>display ospf interface GigabitEthernet 0/0/0
<AR720>display logbuffer | include -i ospf
```

**Checklist OSPF** :
```text
[ ] Même area-id des deux côtés
[ ] Même masque / sous-réseau IP
[ ] Hello/Dead timers identiques (défaut 10/40 s)
[ ] MTU cohérente (un classique en ExStart)
[ ] Authentification identique (ou absente des deux côtés)
[ ] Router-ID unique (pas de doublon !)
[ ] network déclaré dans la bonne area (vue area, pas vue processus)
```

## 109. Cas n°8 — Route manquante dans la table

**Symptômes** : `ping` échoue vers un réseau qui « devrait » être connu.

**Diagnostic** :
```vrp
<AR720>display ip routing-table 192.168.20.0 24
<AR720>display ip routing-table protocol ospf | include 192.168.20
<AR720>display current-configuration | include route-static
```

**Causes** : `ip route-static` mal saisie (next-hop injoignable), redistribution
oubliée en OSPF/BGP, route filtrée par une route-policy, VRF/VPN-instance
incorrecte (regarder la bonne table : `display ip routing-table vpn-instance X`).

## 110. Cas n°9 — Pas d'accès Internet derrière l'USG6000

**Symptômes** : le LAN ne sort pas, alors que l'USG a sa route par défaut.

**Diagnostic** :
```vrp
<USG6000>display security-policy rule all     # une règle trust->untrust en permit ?
<USG6000>display firewall session table       # des sessions se créent-elles ?
<USG6000>display nat outbound                 # le NAT sortant est-il configuré ?
<USG6000>display ip routing-table
```

**Causes classiques** : security-policy manquante (deny par défaut !), NAT
`easy-ip` oublié sur l'interface untrust, interface pas dans la bonne zone.

