---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-7
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [930, 1102]
sha256: c6b02bfabadf74a5362efd6e9e044ba6c05059cc8289a16502127af640655fc7
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

### Fiche animateur — points à insister
- La table de routage se lit en 4 colonnes : **Destination, Préférence, Coût, Next-hop**. Faire réciter.
- Le routage est toujours une question d'**aller ET retour** : 80 % des pannes "ça ping dans un sens" viennent d'une route de retour manquante.
- La floating static est le mécanisme de backup le plus simple ; ses limites (pas de détection de panne distale) justifient les protocoles dynamiques → transition naturelle vers le TP6 (OSPF).
- Lien terrain : sur vos AR720, les routes statiques servent pour les petits sites et les backup 4G ; la préférence permet de basculer automatiquement entre fibre et 4G.

---

## 8. TP6 — OSPF mono-zone

### Objectif
Déployer OSPF en zone 0 sur 3-4 routeurs : adjacences, élection DR/BDR, coûts, et redistribution d'une route par défaut.

### Prérequis
TP5 (routage, tables de routage).

### Topologie

```
                    R1 (AR2220, Router-ID 1.1.1.1)
                   /        \
   10.0.12.0/30   /          \   10.0.13.0/30
                 /            \
   R2 (2.2.2.2) ────────────── R3 (3.3.3.3)
   10.0.23.0/30 (lien R2-R3)
       |                           |
  GE0/0/1                       GE0/0/1
  192.168.20.254/24             192.168.30.254/24
  LAN B (SW2, PC)               LAN C (SW3, PC)
       |
  (R1 a aussi le LAN A : 192.168.10.254/24 sur GE0/0/1)
```

- 3 routeurs **AR2220** en triangle (triangle = élection DR/BDR intéressante sur chaque segment).
- Chaque routeur a un LAN (switch S3700 + 1 PC).
- R1 simule la sortie Internet : route statique par défaut vers un Cloud/ISP, à **redistribuer** dans OSPF.

### Énoncé

1. Adresser toutes les interfaces (plan ci-dessus, /30 sur les liens inter-routeurs).
2. Configurer OSPF **process 1, area 0** sur les 3 routeurs, avec des **Router-ID manuels** (1.1.1.1, 2.2.2.2, 3.3.3.3).
3. Annoncer les réseaux : les /30 et les LAN.
4. Vérifier les voisins (`display ospf peer`), identifier **DR et BDR** sur chaque segment.
5. Forcer R1 comme DR du segment R1-R2 (priorité d'interface).
6. Modifier un **coût OSPF** pour influencer un chemin, observer le changement dans la table de routage.
7. Sur R1 : créer une route par défaut statique (vers l'ISP simulé) et la **redistribuer** dans OSPF (`default-route-advertise`).
8. Vérifier depuis R3 que la défaut est apprise en OSPF (O_ASE) et tester.

### Correction pas à pas

**Adressage (exemple R1) :**

```
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet0/0/1]undo portswitch
[R1-GigabitEthernet0/0/1]ip address 192.168.10.254 24
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]undo portswitch
[R1-GigabitEthernet0/0/2]ip address 10.0.12.1 30
[R1]interface GigabitEthernet 0/0/3
[R1-GigabitEthernet0/0/3]undo portswitch
[R1-GigabitEthernet0/0/3]ip address 10.0.13.1 30
```

(R2 : GE0/0/1 = 192.168.20.254/24, GE0/0/2 = 10.0.12.2/30, GE0/0/3 = 10.0.23.1/30 ; R3 : GE0/0/1 = 192.168.30.254/24, GE0/0/2 = 10.0.13.2/30, GE0/0/3 = 10.0.23.2/30.)

**OSPF sur R1 :**

```
[R1]ospf 1 router-id 1.1.1.1
[R1-ospf-1]area 0
[R1-ospf-1-area-0.0.0.0]network 192.168.10.0 0.0.0.255
[R1-ospf-1-area-0.0.0.0]network 10.0.12.0 0.0.0.3
[R1-ospf-1-area-0.0.0.0]network 10.0.13.0 0.0.0.3
[R1-ospf-1-area-0.0.0.0]quit
[R1-ospf-1]quit
```

> Le `network` OSPF utilise un **masque inversé** (wildcard) : /24 → 0.0.0.255, /30 → 0.0.0.3. C'est la source d'erreur n°1.

**R2 et R3 :** même structure avec leurs réseaux et leurs Router-ID.

**Vérifier les adjacences :**

```
[R1]display ospf peer
# État Full = adjacence établie. Sur un segment à 3 routeurs (pas ici, nos liens sont point-à-point
# à 2 routeurs : pas de DR/BDR sur les /30... voir note ci-dessous).

[R1]display ospf peer brief
```

> **Note DR/BDR** : sur des liens /30 à 2 routeurs, OSPF ne élit pas de DR/BDR utile (élection triviale). Pour observer une vraie élection, ajouter un **switch entre R1, R2 et R3** (un S3700 avec les 3 routeurs branchés dessus en 10.0.123.0/29) : le segment multi-accès élit alors un DR et un BDR. **Variante conseillée par l'animateur** : remplacer le triangle par une étoile via SW0.

**Variante étoile (recommandée pour DR/BDR) :**

```
            SW0 (S3700)
   ┌─────────┼─────────┐
  R1        R2        R3
  .1        .2        .3   (10.0.123.0/29)
```

```
# Sur chaque routeur, l'interface vers SW0 en mode L3 (undo portswitch), IP 10.0.123.x/29.
# OSPF : network 10.0.123.0 0.0.0.7 area 0

# Forcer R1 DR :
[R1]interface GigabitEthernet 0/0/4
[R1-GigabitEthernet0/0/4]ospf dr-priority 100
# (défaut = 1 ; 0 = inéligible)

[R1]display ospf peer
# Colonne "DR/BDR" : R1 doit apparaître comme DR.

# Pour que l'élection se refasse proprement après changement de priorité :
[R1]reset ospf 1 peer   # (ou reboot du process)
```

**Jouer sur les coûts :**

```
# Rendre le chemin via R2 moins attractif depuis R1 vers le LAN C :
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]ospf cost 100
# (coût par défaut = 1 sur GE... en fait calculé sur la bande passante de référence ;
#  sur eNSP le coût affiché par défaut est souvent 1)

[R1]display ip routing-table 192.168.30.0
# Le chemin doit basculer sur l'autre lien. Comparer avant/après.
```

**Redistribution de la défaut (sur R1) :**

```
# Simuler l'ISP : route statique par défaut (next-hop = une IP du lab ou un Cloud)
[R1]ip route-static 0.0.0.0 0.0.0.0 10.0.99.2

# L'injecter dans OSPF :
[R1]ospf 1
[R1-ospf-1]default-route-advertise always
[R1-ospf-1]quit
# "always" : annonce même si R1 n'a pas lui-même de défaut active (pratique en lab).

# Vérification depuis R3 :
[R3]display ip routing-table 0.0.0.0
# Proto = O_ASE, Pre = 150, NextHop = vers R1.
```

### Vérifications

```
[R1]display ospf peer brief        # voisins et états
[R1]display ospf lsdb              # base de données des LSAs (Router-LSA, Network-LSA...)
[R1]display ospf routing           # table de routage calculée par OSPF
[R1]display ospf interface         # interfaces OSPF : area, cost, DR/BDR, timers
[R1]display ip routing-table protocol ospf   # routes OSPF dans la table globale
```

**Capture Wireshark** sur un lien inter-routeur : filtrer `ospf` → **Hello** (multicast 224.0.0.5, toutes les 10 s), puis provoquer une coupure et observer les **LSU/LSAck** (mises à jour). Le Hello contient : Router-ID, Area, timers, liste des voisins vus.

### Pièges classiques

1. **Wildcard inversé faux** : `network 10.0.12.0 0.0.0.255` au lieu de `0.0.0.3` → le réseau n'est pas forcément activé comme voulu (ici ça activerait OSPF sur trop d'interfaces). Toujours calculer : wildcard = 255.255.255.255 − masque.
2. **Router-ID dupliqué ou 0.0.0.0** : sans `router-id` manuel et sans IP de loopback, eNSP peut prendre une IP d'interface ; deux routeurs avec le même Router-ID = adjacence impossible. **Toujours fixer le Router-ID manuellement** (bonne pratique).
3. **Area différente des deux côtés** : les Hello sont rejetés (area mismatch) → état `Down`, sans message très explicite. `display ospf interface` des deux côtés pour comparer.
4. **Timers Hello/Dead différents** : même symptôme (mismatch). En lab, ne pas y toucher.
5. **Interface passive oubliée** : sur les interfaces LAN (vers les switchs), pas besoin de voisins OSPF : `silent-interface` (ou `ospf silent-interface`) évite d'envoyer des Hello inutiles vers les PC. Pas bloquant, mais bonne pratique à enseigner.
6. **MTU mismatch** : les AR eNSP ont une MTU homogène, mais en production un MTU différent bloque l'adjacence à l'état `ExStart`. À connaître.
7. **`default-route-advertise` sans `always`** : si R1 n'a pas de route par défaut active dans sa table, rien n'est annoncé. En lab, `always` simplifie.

### Barème indicatif (20 points)

