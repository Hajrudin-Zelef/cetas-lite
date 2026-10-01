---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-12
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [2358, 2541]
sha256: 7078cb8187155cab63e55f02f60f119487709a3eadebd172dea2f77decb57696
---

# VRP — Le système d'exploitation transversal Huawei

Commandes VRP associées : sections 91-92 (SNMP), 69 (syslog), 58 (NTP).

## 125. Pense-bête de poche — une page

```text
=================== VRP PENSE-BÊTE ===================
VUES
  <H>  user view (supervision)      [H]  system view (config)
  [H-GE0/0/0] interface             [H-ospf-1] / [H-bgp] protocoles
  quit=remonter  return/Ctrl+Z=user view  run display X (depuis config)

AFFICHER (display, JAMAIS show)
  display version | device elabel | startup | clock
  display interface brief | display ip interface brief
  display ip routing-table | display ospf peer | display bgp peer
  display arp | display mac-address | display vlan
  display logbuffer | display trapbuffer
  display cpu-usage | display memory
  display current-configuration | display saved-configuration
  display users | display ssh server status

CONFIGURER
  system-view | sysname X | interface GE0/0/0 | ip address 1.1.1.1 24
  undo shutdown | description TEXTE | quit
  vlan 10 | port link-type access | port default vlan 10
  ip route-static 0.0.0.0 0 192.168.1.254
  ospf 1 router-id 1.1.1.1 | area 0 | network 192.168.1.0 0.0.0.255

SAUVER / FICHIERS
  save  (OBLIGATOIRE !) | compare configuration
  dir | copy | delete | undelete | reset recycle-bin
  startup saved-configuration X.zip | display startup

DIAGNOSTIC
  ping 1.1.1.1 | ping -a SRC DST | tracert DST
  debugging X + terminal debugging | undo debugging all
  display diagnostic-information

ACCÈS
  rsa local-key-pair create | stelnet server enable
  user-interface vty 0 4 | authentication-mode aaa | protocol inbound ssh
  aaa -> local-user X password irreversible-cipher Y | privilege level 15

UPGRADE
  tftp S get F.cc | startup system-software flash:/F.cc
  display startup (vérifier !) | save | reboot

SECOURS
  Ctrl+B au boot -> Clear password for console user -> Boot with default mode
  reset saved-configuration + reboot = reset usine
=====================================================
```

## 126. Glossaire VRP

| Terme | Signification |
|---|---|
| VRP | Versatile Routing Platform — l'OS Huawei |
| User view | Vue de supervision (`<H>`) |
| System view | Vue de configuration globale (`[H]`) |
| `display` | Commande d'affichage (équivalent `show` Cisco) |
| `undo` | Annule une commande (équivalent `no` Cisco) |
| `save` | Sauvegarde la config RAM vers flash |
| vrpcfg.zip | Fichier de configuration de démarrage par défaut |
| `.cc` | Fichier logiciel système VRP |
| `.pat` | Fichier de patch logiciel |
| BootROM/BootLoad | Menu de démarrage bas niveau (Ctrl+B) |
| VTY | Lignes virtuelles pour accès distant (Telnet/SSH) |
| AAA | Authentication, Authorization, Accounting |
| STelnet | SSH côté Huawei (Secure Telnet) |
| Eth-Trunk | Agrégation de liens (équivalent EtherChannel) |
| Vlanif | Interface virtuelle de VLAN (équivalent SVI Cisco) |
| `startup` | Fichiers utilisés au prochain démarrage |
| Candidate database | Base de config en attente de `commit` (VRP8) |
| `commit` | Valide la candidate database (VRP8) |
| `rollback` | Annule les changements non commités (VRP8) |
| HRP | Huawei Redundancy Protocol (HA des USG) |
| iStack / CSS | Empilage / cluster de switchs Huawei |
| PAF | Fichier de licence/fonctionnalité (à vérifier selon version) |
| info-center | Sous-système de journalisation VRP |
| NQA | Network Quality Analyzer (sondes de supervision) |

## 127. Quiz — 10 questions (réponses en section 128)

**Q1.** Dans quelle vue tape-t-on `display version`, et que se passe-t-il si
on la tape en system view ?

**Q2.** Un collègue configure un VLAN un vendredi soir, tout fonctionne. Le
lundi matin après une coupure électrique, le VLAN a disparu. Pourquoi, et
quelle commande aurait évité cela ?

**Q3.** Traduisez en VRP : `show ip interface brief`, `conf t`,
`no shutdown`, `copy run start`, `reload`.

**Q4.** Sur un USG6000, deux interfaces sont dans les zones `trust` et
`untrust` mais aucun trafic ne passe. Quelle est la cause la plus probable ?

**Q5.** En OSPF, où déclare-t-on le `network` : dans la vue du processus
`ospf` ou dans la vue d'area ? Écrivez la séquence complète.

**Q6.** `display interface brief` affiche `up/down` sur une interface. Que
signifie chaque état, et par où commence-t-on l'investigation ?

**Q7.** SSH refuse les connexions alors que `stelnet server enable` est
configuré. Citez deux causes probables et leur diagnostic.

**Q8.** Avant de rebooter pour un upgrade, quelles deux sections de
`display startup` vérifiez-vous impérativement ?

**Q9.** Quelle est la différence entre `quit`, `return` et `Ctrl+Z` ?

**Q10.** Vous êtes en VRP8 et le prompt affiche `[~AR720]`. Que signifie le
tilde `~`, et que devez-vous faire avant de quitter ?

## 128. Quiz — réponses

**R1.** En **user view** (`<H>`). En system view, il faut `run display
version` (ou `quit` pour revenir en user view). Sans `run`, VRP répond que la
commande est inconnue dans cette vue.

**R2.** Il a oublié `save` : la configuration vivait en RAM et a été perdue
au reboot. Réflexe : terminer chaque session par `save` (+ `display
saved-configuration` pour vérifier).

**R3.** `display ip interface brief` ; `system-view` ; `undo shutdown` ;
`save` ; `reboot`.

**R4.** Il manque une **security-policy** autorisant le trafic
(`trust` → `untrust`, `action permit`). Sur USG, le défaut est **deny entre
zones** : sans règle, rien ne passe.

**R5.** Dans la **vue d'area**. Séquence :
```vrp
[AR720]ospf 1 router-id 1.1.1.1
[AR720-ospf-1]area 0
[AR720-ospf-1-area-0.0.0.0]network 192.168.1.0 0.0.0.255
```

**R6.** PHY `up` = signal physique OK ; Protocol `down` = problème de couche
liaison (négociation, encapsulation, keepalive). On commence par `display
interface X` (détail : speed/duplex/erreurs) puis `display logbuffer`.

**R7.** (1) Clé RSA absente → `display rsa local-key-pair public`, puis
`rsa local-key-pair create`. (2) Utilisateur sans `service-type ssh` →
`display current-configuration | include local-user`, corriger en vue aaa.
Autres : `protocol inbound` restreint à telnet, ACL sur VTY, VTY saturés.

**R8.** « Next startup system software » (le nouveau `.cc` doit y figurer)
et « Next startup saved-configuration file » (le bon fichier de config).
Si ce n'est pas le cas : NE PAS REBOOTER.

**R9.** `quit` remonte d'**un** niveau (et quitte la session en user view) ;
`return` et `Ctrl+Z` retournent **directement** à la user view d'où que
l'on soit.

**R10.** Le `~` signale des modifications **non commitées** dans la
candidate database (VRP8). Il faut `commit` (valider) ou `rollback
configuration` (annuler) avant de partir, puis `save`.

## 129. Pour aller plus loin — par famille d'équipement

- **AR720** : voir le guide `huawei_ar720_guide.md` (routage WAN, VPN
  IPsec, NAT, QoS) — la couche 3 avancée au-delà du socle VRP.
- **S310** : voir `huawei_s310_guide.md` (VLAN, STP/MSTP, Eth-Trunk,
  iStack, PoE) — la commutation campus en détail.
- **USG6000** : voir `huawei_usg6000_guide.md` (zones, security-policy,
  NAT, IPsec, HRP) — la sécurité périmétrique.
- **CLI Huawei** : les Command Reference PDF par produit/version
  (EDOC sur support.huawei.com) restent la source de vérité pour la
  syntaxe exacte — Zelef les a en partie dans son corpus RAG.

## 130. Pour aller plus loin — certifications et pratique

- **HCIA-Datacom** : la certification d'entrée Huawei ; son premier module
  est exactement ce socle VRP (vues, display, config de base).
- **eNSP** (Enterprise Network Simulation Platform) : le simulateur
  gratuit Huawei pour pratiquer VRP sans matériel — idéal pour rejouer les
  cas de dépannage de ce guide.
- Pratique recommandée : pour chaque cas de dépannage (sections 102-119),
  provoquez la panne en lab (débranchez, oubliez le save, cassez l'OSPF) et
  déroulez la méthode. On retient en cassant.

## 131. Erreurs fréquentes des débutants VRP — top 10

