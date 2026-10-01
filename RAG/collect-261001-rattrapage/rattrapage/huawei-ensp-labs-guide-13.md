---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-13
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1899, 2038]
sha256: c63f634eee04511a70020d6b1c865e2cbecfffffd3b07c5b493e27f7d70881ba
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

### Fiche animateur — points à insister
- Le raisonnement "zones" : on ne filtre plus des IP, on autorise des **flux entre zones**. C'est le changement de paradigme par rapport aux ACL de routeur.
- **Deny implicite partout** : chaque flux doit être explicitement autorisé. En audit, partir du principe "tout ce qui n'est pas autorisé est interdit".
- Stateful : une règle aller suffit pour le retour d'une session établie. Démonstration : `display firewall session table` pendant un ping.
- Ordre : **NAT destination → politique interzone** : la politique voit l'adresse **après** translation (d'où `policy destination 192.168.10.100`, l'adresse interne).
- Lien terrain : vos USG6000 = même logique, syntaxe `security-policy` (règles nommées avec `rule name ... source-zone ... destination-zone ... action permit`). Le raisonnement appris ici s'applique tel quel.
- Question piège : "Faut-il une règle untrust→trust pour que le retour d'une navigation web passe ?" → Non (stateful). "Et pour exposer un serveur ?" → Oui, c'est un flux **initié** depuis untrust.

---

## 14. TP12 — Dépannage guidé : la topologie cassée

### Objectif
Mettre en pratique une **méthode de diagnostic** structurée sur une topologie contenant 5 pannes volontaires couvrant les couches 2, 3 et la sécurité.

### Prérequis
Tous les TP précédents (ou au minimum TP1, TP2, TP5, TP6, TP7).

### Topologie (saine de référence)

```
PC1 (VLAN 10, 192.168.10.11/24) ──E0/0/1 SW1 (S3700) GE0/0/1 ═══ GE0/0/1 SW2 (S3700) E0/0/1── PC3 (VLAN 10, 192.168.10.13/24)
PC2 (VLAN 20, 192.168.20.12/24) ──E0/0/2        (trunk)                    E0/0/2── PC4 (VLAN 20, 192.168.20.14/24)
                                            GE0/0/2 │ (trunk vers R1)
                                                    │ GE0/0/1.10 (192.168.10.254) / GE0/0/1.20 (192.168.20.254)
                                                    │ R1 (AR2220, OSPF area 0)
                                            GE0/0/2 │ (10.0.12.1/30)
                                                    │ GE0/0/2 (10.0.12.2/30)
                                                    │ R2 (AR2220, OSPF area 0, NAT Easy IP sur GE0/0/3)
                                            GE0/0/3 │ (198.51.100.1/30)
                                                    │ "Internet" : R-ISP (198.51.100.2/30) + Server (203.0.113.10)
```

- Routage inter-VLAN via R1 (sous-interfaces), OSPF entre R1 et R2, NAT sur R2 vers l'extérieur.
- Objectif final : PC1 peut joindre PC3 (L2), PC2 (routage), et 203.0.113.10 (NAT+Internet).

### Les 5 pannes (à introduire par l'animateur, une fiche par stagiaire/groupe)

| # | Panne | Couche | Symptôme | Commande de sabotage (animateur) |
|---|---|---|---|---|
| 1 | Port E0/0/1 de SW2 resté en **VLAN 1** au lieu du VLAN 10 | L2 / VLAN | PC1 ne ping pas PC3, mais PC1 ping SW1... (en fait PC1 ping PC2 ? non : PC3 injoignable) | `interface E0/0/1` → `undo port default vlan` (retour PVID 1) |
| 2 | Trunk SW1↔SW2 : `allow-pass` **sans le VLAN 20** côté SW1 | L2 / trunk | PC2 ne ping pas PC4 ; PC1 ping PC3 OK | `port trunk allow-pass vlan 10` seul sur GE0/0/1 de SW1 |
| 3 | OSPF : **Router-ID dupliqué** (R2 configuré avec 1.1.1.1 comme R1) | L3 / OSPF | Adjacence Down, pas de route OSPF, PC1 ne ping pas... (en fait ici R2 ne route plus vers Internet : PC1 ne ping pas 203.0.113.10) | `ospf 1 router-id 1.1.1.1` sur R2 |
| 4 | NAT : **ACL 2000 avec mauvais wildcard** (`0.0.255.255` au lieu de `0.0.0.255`) | L3 / NAT | Le ping vers Internet échoue (pas de translation) alors que le routage est OK | `rule permit source 192.168.0.0 0.0.255.255` |
| 5 | **ACL sur GE0/0/2 de R1** bloquant tout (oubliée d'un TP précédent) | Sécurité | Tout le trafic inter-sites bloqué, pings KO partout sauf en local | `traffic-filter inbound acl 3000` avec une acl 3000 `rule deny` |

### Énoncé (distribué aux stagiaires)

> "Le réseau ci-dessus **fonctionnait hier soir**. Ce matin, plusieurs utilisateurs se plaignent : certains PC ne se joignent plus, Internet est coupé par intermittence... Votre mission : **trouver et réparer les 5 pannes**, en suivant la méthode de diagnostic. Pour chaque panne : symptômes observés, hypothèse, commande de vérification, correction, test de validation. Temps imparti : 1 h 30."

### Correction : la méthode pas à pas (à faire appliquer)

**Méthode en 6 étapes (à afficher en salle) :**

```
1. DÉFINIR le symptôme précisément (qui ne joint pas qui ? depuis quand ?).
2. DÉLIMITER : le problème est-il local (un PC), par VLAN, par site, global ?
3. VÉRIFIER la couche 1/2 : liens up ? (display interface brief)
4. VÉRIFIER la couche 2 : VLAN, trunk, STP (display port vlan, display vlan)
5. VÉRIFIER la couche 3 : IP, routage, protocoles (display ip routing-table, display ospf peer)
6. VÉRIFIER les fonctions avancées : DHCP, NAT, ACL, firewall (display nat session, display acl all)
   → Corriger UNE panne à la fois, re-tester, documenter.
```

**Déroulé type attendu :**

**Étape 1 — Cartographier les symptômes :**

```
PC1> ping 192.168.10.13 (PC3)   → KO   # même VLAN, switches différents
PC1> ping 192.168.10.12 (PC2)   → ...  # autre VLAN, via R1
PC2> ping 192.168.20.14 (PC4)   → KO   # même VLAN, switches différents
PC1> ping 203.0.113.10          → KO   # Internet
```

> Premier tri : les pannes intra-VLAN inter-switch (1 et 2) vs le reste.

**Panne 1 — PC1 ne ping pas PC3 :**

```
[SW2]display port vlan
# E0/0/1 : PVID 1, pas membre du VLAN 10 → le port n'est pas dans le bon VLAN.
[SW2]interface Ethernet 0/0/1
[SW2-Ethernet0/0/1]port link-type access
[SW2-Ethernet0/0/1]port default vlan 10
# Test : PC1 ping PC3 → OK.
```

**Panne 2 — PC2 ne ping pas PC4 (mais PC1 ping PC3 OK après réparation 1) :**

```
[SW1]display port vlan
# GE0/0/1 (trunk) : allow-pass vlan 10 seul → le VLAN 20 ne traverse pas.
[SW2]display port vlan
# GE0/0/1 : allow-pass vlan 10 20 → asymétrie détectée !
[SW1]interface GigabitEthernet 0/0/1
[SW1-GigabitEthernet0/0/1]port trunk allow-pass vlan 10 20
# Test : PC2 ping PC4 → OK.
```

**Panne 3 — OSPF ne monte pas (pas de route vers Internet / inter-sites KO) :**

```
[R1]display ospf peer
# Aucun voisin (ou état anormal).
[R1]display ospf interface
# Area 0 des deux côtés, timers OK...
[R2]display current-configuration | include router-id
# R2 a router-id 1.1.1.1 → DOUBLON avec R1 !
[R2]ospf 1
[R2-ospf-1]router-id 2.2.2.2
[R2-ospf-1]quit
[R2]reset ospf 1   # (ou reboot du process pour prendre en compte)
[R1]display ospf peer
# État Full → adjacence rétablie.
```

**Panne 4 — Internet KO malgré OSPF OK :**

```
# Le ping vers 198.51.100.2 (IP publique de R2) passe depuis R1,
# mais PC1> ping 203.0.113.10 échoue → le routage est bon, c'est le NAT.
[R2]display acl 2000
# rule permit source 192.168.0.0 0.0.255.255 → wildcard faux !
# (matche 192.168.0.0-192.168.255.255... en fait le wildcard 0.0.255.255 matche bien
#  192.168.10.0 ! → subtilité : la règle matche, donc le NAT devrait marcher...)
```

> **Variante plus parlante** : l'animateur mettra plutôt `rule permit source 192.168.11.0 0.0.0.255` (mauvais réseau) → aucun paquet de 192.168.10.0 ne matche → pas de translation. `display nat session all` reste vide pendant les pings = preuve que le NAT ne s'enclenche pas.

```
[R2]acl 2000
[R2-acl-basic-2000]undo rule 5   # (numéro de la règle fautive)
[R2-acl-basic-2000]rule permit source 192.168.10.0 0.0.0.255
[R2-acl-basic-2000]rule permit source 192.168.20.0 0.0.0.255
# Test : PC1 ping 203.0.113.10 → OK (display nat session all montre les sessions).
```

**Panne 5 — l'ACL parasite sur R1 :**

