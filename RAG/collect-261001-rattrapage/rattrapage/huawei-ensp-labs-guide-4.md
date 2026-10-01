---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-4
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "cost", "distribution", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [444, 605]
sha256: 8d0ba3c680fd9c04ff2e4773cb71fca9459be24619249cb14a381b9891959eba
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

1. **Trunk configuré d'un seul côté** : l'autre côté reste en hybrid/PVID 1 → les trames taggées arrivent sur un port qui ne les attend pas. Symptôme : ping KO dans les deux VLAN. Diagnostic : `display port vlan` des deux côtés et comparer.
2. **`port trunk allow-pass vlan` oublié ou incomplet** : par défaut un trunk ne laisse passer que le VLAN 1. Si on autorise seulement le VLAN 10, le VLAN 20 ne passe pas → **le cas "VLAN qui ne passe pas"** de l'énoncé. Diagnostic :
   ```
   [SW2]display port vlan   # VLAN List du trunk : 10 seul → le 20 est filtré
   ```
   Réparation : `port trunk allow-pass vlan 10 20`.
3. **VLAN non créé d'un côté** : un trunk qui autorise le VLAN 20 alors que le VLAN 20 n'existe pas sur SW2 → le VLAN est "inactif". `display vlan` le montre.
4. **PVID (native VLAN) différent des deux côtés** : en 802.1Q pur avec que des trames taggées, pas d'impact ici, mais à signaler comme cause classique de fuite inter-VLAN en production.
5. **Câbler sur le mauvais port** : configurer GE0/0/1 alors que le câble est sur GE0/0/2. `display interface brief` montre le lien up sur le port réellement câblé.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| VLAN créés des deux côtés | 3 |
| Ports access corrects | 4 |
| Trunk configuré des deux côtés (link-type + allow-pass 10 et 20) | 6 |
| Tests ping conformes | 3 |
| Diagnostic et réparation de la panne provoquée (méthode + commandes) | 4 |

### Durée estimée
**1 heure** (dont 15 min de dépannage guidé).

### Fiche animateur — points à insister
- Le trunk ne "crée" pas de connectivité : il **transporte** des VLAN existants des deux côtés. Les trois conditions : VLAN créés des deux côtés + trunk des deux côtés + allow-pass symétrique.
- La capture Wireshark du tag 802.1Q est le moment "aha" du TP : prévoir 10 minutes.
- Lien terrain : entre vos S310 d'étages, les liens montants sont toujours des trunks ; un VLAN oublié dans le allow-pass = le symptôme "ça marche sur un switch mais pas sur l'autre".
- Faire verbaliser la méthode de dépannage : 1) le lien est-il up ? 2) les VLAN existent-ils des deux côtés ? 3) le trunk est-il symétrique ? (`display port vlan` des deux côtés, côte à côte).

---

## 5. TP3 — STP/RSTP : boucle volontaire, élection root, coûts, portfast/edge

### Objectif
Comprendre pourquoi les boucles de niveau 2 sont dangereuses, comment STP/RSTP les neutralise, et comment contrôler l'élection du root bridge.

### Prérequis
TP1, TP2 (VLAN, trunk).

### Topologie

```
                    ┌──────────┐
   PC1 ──E0/0/1     │  SW1     │     E0/0/1── PC3
   (VLAN 10)        │ (S3700)  │
                    └────┬─────┘
              GE0/0/1    │    GE0/0/2        <- boucle volontaire à 3 switches
                         │         (triangle)
              ┌──────────┼──────────┐
              │          │          │
        GE0/0/1    SW2 (S3700)   GE0/0/1
              │     E0/0/1│          │
              │   PC2     │    SW3 (S3700)
              │  (VLAN10) │     E0/0/1── PC4 (VLAN 10)
              └──────────┴──────────┘
                   GE0/0/2 ↔ GE0/0/2  (SW2 ↔ SW3)
```

- 3 switches **S3700** en triangle : SW1↔SW2 (GE0/0/1↔GE0/0/1), SW1↔SW3 (GE0/0/2↔GE0/0/1), SW2↔SW3 (GE0/0/2↔GE0/0/2).
- Trunks 802.1Q sur les 3 liens inter-switch, VLAN 10 autorisé partout.
- 4 PC en VLAN 10 : 192.168.10.11 à 192.168.10.14.

### Énoncé

**Partie A — Observer le danger (5 min) :**
1. Câbler le triangle **sans activer STP** (par défaut, sur les S eNSP, STP est... à vérifier : l'animateur le désactive au préalable avec `stp disable` sur les 3 switches pour la démo).
2. Depuis PC1, pinger PC4 en continu. Observer : tempête de broadcast, ping qui part en timeout, CPU des switches qui monte. **Arrêter vite** (débrancher un câble du triangle) pour retrouver un réseau sain.
   > En lab eNSP l'effet est atténué par la virtualisation, mais les doublons de trames et l'instabilité MAC sont visibles dans `display mac-address` (adresses qui "flappent").

**Partie B — Activer RSTP et observer :**
3. Réactiver STP en mode **RSTP** sur les 3 switches.
4. Identifier le **root bridge** élu, les **ports designated / root / alternate (bloqué)**.
5. Forcer **SW1** comme root primaire et **SW2** comme secondaire (priorités).
6. Observer la convergence après coupure d'un lien du triangle (débrancher/rebrancher un câble).

**Partie C — Edge ports :**
7. Configurer les ports vers les PC en **edge** (portfast) : `stp edged-port enable`.
8. Constater la différence de temps de passage à l'état forwarding (immédiat vs ~30 s en STP classique / ~2 s en RSTP avec handshake).

### Correction pas à pas

**Préparation (animateur, avant la séance) : topologie câblée, STP désactivé pour la partie A :**

```
[SWx]stp disable
```

**Partie B — activation RSTP :**

```
# Sur les 3 switches :
[SW1]stp mode rstp
[SW1]stp enable

[SW2]stp mode rstp
[SW2]stp enable

[SW3]stp mode rstp
[SW3]stp enable
```

**Identifier le root :**

```
[SW1]display stp brief
# Affiche pour chaque port : rôle (DESI/ROOT/ALTE), état (FORWARDING/DISCARDING)
# Le switch dont tous les ports sont DESI est le root.

[SW1]display stp
# Détail : "CIST Bridge" avec priorité + MAC ; "Root" indique le root élu.
# Par défaut, priorité 32768 partout : le root = la plus petite MAC.
```

**Forcer l'élection :**

```
[SW1]stp root primary
# Équivaut à stp priority 0 (ou 4096 selon version) : SW1 devient root.

[SW2]stp root secondary
# Priorité 4096 : SW2 devient root si SW1 disparaît.

# Vérification :
[SW3]display stp
# Root ID doit correspondre au Bridge ID de SW1.
```

**Partie C — edge ports vers les PC :**

```
[SW1]interface Ethernet 0/0/1
[SW1-Ethernet0/0/1]stp edged-port enable
[SW1-Ethernet0/0/1]quit
# Répéter sur chaque port connecté à un PC (jamais sur un lien inter-switch !)
```

**Test de convergence :** débrancher le câble SW1↔SW2, observer dans `display stp brief` le port alternate de SW3 passer en forwarding (RSTP : quasi immédiat grâce au proposal/agreement). Rebrancher : retour à l'état initial.

### Vérifications

```
[SWx]display stp brief
# Carte d'identité du spanning tree : rôles + états de chaque port.

[SWx]display stpee inconsistent-port   # (si supporté) ports en incohérence

[SW1]display stp history
# Historique des changements de topologie (TCN).
```

**Capture Wireshark** sur un lien inter-switch : observer les **BPDU** (destination multicast `01:80:C2:00:00:00`), avec les champs Root Identifier, Bridge Identifier, Port Identifier. En RSTP, noter le flag "Proposal/Agreement".

### Pièges classiques

1. **Activer `stp edged-port` sur un lien inter-switch** : si une BPDU arrive sur un edge port, le port perd son statut edge (protection intégrée), mais en formation c'est l'erreur à ne jamais commettre en production : un edge port ne doit jamais recevoir de BPDU.
2. **Oublier `stp enable` après `stp mode rstp`** : le mode seul ne suffit pas.
3. **Root bridge subi** : sans priorité configurée, le switch avec la plus petite MAC devient root — parfois un switch d'accès en bout de chaîne. En production : **toujours** fixer root primary/secondaire sur les switches de cœur/distribution.
4. **Confondre coût de port et priorité** : la priorité élit le root ; le **coût** (`stp cost`) influence le choix du chemin vers le root quand plusieurs chemins existent. Pour préférer un lien : `stp cost 20000` sur le port à défavoriser (coût plus élevé = moins préféré).
5. **VLAN 1 vs instances** : eNSP/VRP de base = une seule instance (CIST). MSTP multi-instances existe sur les vrais S5700 mais sort du cadre de ce TP.

### Barème indicatif (20 points)

