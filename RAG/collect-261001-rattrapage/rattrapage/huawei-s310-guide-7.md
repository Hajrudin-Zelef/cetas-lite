---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-7
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [1196, 1448]
sha256: 82a48b68b19bc28bc50ff4afdbb51a9aada6443a5fb847733db35bb6d0532fb9
---

# Guide ultra-complet — Huawei eKit S310

Deux câbles entre deux switches (ou un switch bouclé sur lui-même) = **tempête
de broadcast** = réseau mort en quelques secondes. Le Spanning Tree bloque
les ports redondants et ne les ouvre qu'en cas de panne du lien principal.

**Les 3 protocoles sur le S310 (valeur datasheet) :**

| Protocole | Standard | Usage |
|---|---|---|
| STP | 802.1D | Historique, lent (30–50 s de convergence) — à éviter |
| **RSTP** | 802.1w | **Le choix par défaut** : convergence en quelques secondes |
| MSTP | 802.1s | Multiples instances par VLAN, pour les topologies complexes |

Le S310 supporte aussi **VBST** (VLAN-based Spanning Tree, spécifique Huawei,
interopérable avec PVST/PVST+/RPVST côté Cisco) et **ERPS** (protection en
anneau ITU-T G.8032).

⚠️ **Règle absolue :** STP/RSTP **toujours activé** sur les ports qui relient
des switches entre eux. Le seul endroit où on l'accélère (edge port), c'est
vers les équipements terminaux (section 48).

---

## 44. Activation de RSTP : la base

```
system-view
stp mode rstp
stp enable
quit
save
```

🔧 Vérifications :

```
display stp brief
```

Tu dois voir : le mode (RSTP), l'ID du root bridge élu, et pour chaque port
son rôle (Root/Designated/Alternate) et son état (Forwarding/Discarding).

⚠️ **Active STP AVANT de brancher le deuxième câble** d'une redondance.
L'ordre « je câble puis je configure » crée la boucle qu'on voulait éviter.
Procédure : configure les deux switches, active RSTP des deux côtés, **puis**
branche le lien redondant.

---

## 45. Choisir le root bridge : la priorité

Par défaut, le switch avec la plus petite priorité (puis la plus petite MAC)
devient root. **Ne laisse pas le hasard décider** : le root doit être ton
switch le plus central et le plus fiable (le cœur, pas un switch d'accès
au fond d'un atelier).

```
system-view
stp root primary        # sur le switch cœur : devient root (priorité 0)
quit
save
```

Sur le switch de secours :

```
system-view
stp root secondary      # priorité 4096 : prend le relais si le primary tombe
quit
save
```

🔧 Vérification : `display stp` → « CIST Root » doit être l'adresse MAC du
switch cœur, sur **tous** les switches.

⚠️ Un root élu au hasard (souvent le plus vieux switch avec la plus petite
MAC) = chemins sous-optimaux + reconvergences bizarres. **Toujours** fixer
primary/secondary explicitement.

---

## 46. Edge port (l'équivalent du portfast) : pour les terminaux

Un port vers un PC, une imprimante ou un AP **ne doit pas attendre** la
convergence STP (30 s en STP classique, même si RSTP est plus rapide) :
le DHCP timeout et l'utilisateur appelle.

```
system-view
interface GigabitEthernet0/0/5
 stp edged-port enable
quit
save
```

En série :

```
system-view
port-group pg-users
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 stp edged-port enable
quit
save
```

⚠️ **JAMAIS d'edge port vers un autre switch.** Si un BPDU arrive sur un edge
port, le port perd son statut edge et repasse en mode normal (comportement de
protection standard) — mais pendant ce temps, une boucle a pu faire des dégâts.
Combine avec **BPDU guard** (section 47) sur les ports utilisateurs.

---

## 47. BPDU guard : le coupe-circuit anti-boucle utilisateur

**Scénario :** un utilisateur branche un petit switch « de bureau » (5 ports à
20 €) sous son bureau, avec deux câbles vers la prise murale → boucle → réseau
à genoux. **BPDU guard** désactive automatiquement le port dès qu'il reçoit un
BPDU (donc dès qu'un switch est branché côté utilisateur).

```
system-view
interface GigabitEthernet0/0/5
 stp bpdu-protection
quit
save
```

En série sur tous les ports utilisateurs :

```
system-view
port-group pg-users
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 stp edged-port enable
 stp bpdu-protection
quit
save
```

🔧 Quand un port est coupé par BPDU guard : `display interface` montre le port
down, les logs (`display logbuffer`) indiquent la raison. Pour le réarmer :
`shutdown` puis `undo shutdown` sur le port — **après avoir retiré le switch
sauvage**, sinon ça recommence.

✅ **Combo gagnant sur chaque port utilisateur :** `stp edged-port enable` +
`stp bpdu-protection`. C'est la config « set and forget » qui évite 90 % des
boucles accidentelles.

---

## 48. Root guard et BPDU filter : les compléments

**Root guard** : sur un port descendant vers l'accès, empêche qu'un switch
d'accès mal configuré (ou malveillant) ne devienne root.

```
system-view
interface GigabitEthernet0/0/28     # uplink vers un switch d'accès
 stp root-protection
quit
save
```

**BPDU filter** : bloque totalement l'émission/réception de BPDU sur un port.
⚠️ À utiliser avec une extrême prudence (jamais vers un switch), typiquement
vers un équipement qui « n'aime pas » les BPDU.

> Disponibilité exacte de `stp root-protection` / `stp bpdu-filter` : **à vérifier
> sur la version logicielle du modèle exact** (`stp ?` en vue système).

**Résumé des protections STP :**

| Protection | Où | Effet |
|---|---|---|
| edged-port | Ports terminaux | Pas d'attente de convergence |
| bpdu-protection | Ports terminaux | Coupe le port si un BPDU arrive |
| root-protection | Ports vers l'accès | Bloque un root illégitime |
| bpdu-filter | Cas très spécifiques | Ignore totalement les BPDU |

---

## 49. Exemple STP complet : deux switches d'accès redondants

**Topologie :** SW-ACC-01 et SW-ACC-02 reliés au cœur par un lien chacun,
**plus** un lien direct entre eux (secours). RSTP doit bloquer le lien
inter-accès en temps normal.

**Sur le cœur (déjà root primary, section 45) :** rien de plus.

**Sur SW-ACC-01 :**

```
system-view
stp mode rstp
stp enable
stp root secondary      # si le cœur tombe, c'est lui le root de secours... 
                       # (en pratique : secondary sur UN seul des deux accès)
#
interface GigabitEthernet0/0/27
 description LIEN_VERS_ACC02_SECOURS
 port link-type trunk
 port trunk allow-pass vlan 10 20 30 50 99
 undo shutdown
quit
save
```

**Sur SW-ACC-02 :** identique, sans `stp root secondary`.

🔧 Vérification sur les trois switches : `display stp brief`.
Le lien inter-accès doit apparaître en **Alternate / Discarding** sur l'un des
deux. Débranche le lien principal d'un accès → le lien de secours passe en
Forwarding en quelques secondes. **Teste-le en vrai** avant la mise en
production : une redondance jamais testée n'existe pas.

---

## 50. Vérifications STP : la routine

```
display stp brief        # rôles et états par port (le plus utile)
display stp              # détail : root, priorités, timers
display stp interface GigabitEthernet0/0/27
display logbuffer | include STP   # événements STP récents
```

**Lecture d'un `display stp brief` sain :**

- Un seul root pour tout le domaine (même « CIST Root » partout).
- Les ports vers les terminaux : `DESI / FORWARDING`.
- Les liens redondants : un en `FORWARDING`, l'autre en `ALTE / DISCARDING`.
- ⚠️ Deux ports en FORWARDING sur une boucle physique = STP ne fait pas son
  travail → vérifie que `stp enable` est actif **des deux côtés**.

---

## 51. ERPS : la protection en anneau (mention)

Pour les topologies en **anneau** (sites industriels, campus), le S310 supporte
**ERPS** (ITU-T G.8032, valeur datasheet) : convergence < 50 ms, plus rapide
et plus propre que STP en anneau. Configuration plus complexe (R-APS, RPL) :
à réserver aux topologies en anneau volontaires, pas au câblage « en boucle
par accident ». Détail de la config : **à vérifier sur le guide de la version
logicielle exacte**.

---

## 52. MSTP : quand RSTP ne suffit plus

RSTP = **une seule** topologie pour tous les VLAN. Si tu veux que le VLAN 10
passe par le lien A et le VLAN 20 par le lien B (répartition de charge),
il faut **MSTP** avec des instances.

