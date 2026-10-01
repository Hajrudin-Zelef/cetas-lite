---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-8
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "cost"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1103, 1251]
sha256: 899e334416f33c9cf0554719e497efb2220b2e01fa3e0b2aeb7c04f4f5ba187f
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

| Critère | Points |
|---|---|
| Adressage + Router-ID manuels | 3 |
| OSPF area 0 sur les 3 routeurs, networks corrects (wildcards) | 5 |
| Adjacences Full vérifiées (`display ospf peer`) | 3 |
| DR/BDR : élection observée + R1 forcé DR (variante étoile) | 4 |
| Coût modifié avec effet observé sur la table de routage | 2 |
| Défaut redistribuée et vue en O_ASE sur R3 | 3 |

### Durée estimée
**1 h 30** (dont 20 min de capture OSPF).

### Fiche animateur — points à insister
- OSPF ne se "configure" pas, il se **déclare** : on déclare les réseaux dans l'area, le protocole fait le reste (voisins, calcul SPF, convergence). C'est le changement de paradigme après le TP5.
- Les 3 conditions d'une adjacence : **même area, mêmes timers, Router-ID uniques** (+ même MTU, même subnet en production).
- DR/BDR : concept souvent mystifié — simplifier : "sur un segment à plus de 2 routeurs, on élit un représentant (DR) pour limiter les échanges ; la priorité la plus haute gagne, à égalité la plus haute Router-ID".
- Le coût = la métrique : **chemin de coût cumulé le plus faible gagne**. Faire manipuler `ospf cost` pour "sentir" la métrique.
- Lien terrain : OSPF est le protocole de vos réseaux d'agence multi-sites ; la redistribution de la défaut depuis le routeur de sortie est le standard.
- Transition vers TP7 : "maintenant que tout le monde se parle en interne, ouvrons vers l'extérieur : le NAT".

---

## 9. TP7 — NAT : NAPT et port mapping

### Objectif
Partager une adresse publique entre plusieurs postes (NAPT / Easy IP) et exposer un serveur interne vers l'extérieur (port mapping / NAT server).

### Prérequis
TP5 (routage), adressage.

### Topologie

```
LAN privé (192.168.10.0/24)                        "Internet" simulé (203.0.113.0/30)
PC1 (.11) ──SW1── GE0/0/1 R1 (AR2220) GE0/0/2 ══════════ GE0/0/2 R-ISP (AR2220) GE0/0/1── Serv-Public (203.0.113.10/24)
PC2 (.12)          .254          198.51.100.1/30 (IP publique)      198.51.100.2/30
Serveur (.100) ── (serveur web interne à exposer)
```

- **R1** : GE0/0/1 = 192.168.10.254/24 (inside), GE0/0/2 = 198.51.100.1/30 (outside, IP "publique").
- **R-ISP** : routeur simulant l'opérateur/Internet, GE0/0/2 = 198.51.100.2/30, GE0/0/1 = 203.0.113.1/24, avec un **Server** (203.0.113.10) jouant le rôle d'un site web public.
- **Serveur interne** : un objet **Server** eNSP en 192.168.10.100 (HTTP activé dans sa configuration).

### Énoncé

1. Adresser et router : R1 connaît 192.168.10.0/24 (direct) ; R-ISP a une route vers 198.51.100.0/30 (direct) ; ajouter sur R-ISP une route vers... attention : **R-ISP ne doit PAS connaître 192.168.10.0/24** (c'est le principe du NAT : le privé n'est pas routé sur Internet).
2. Configurer sur R1 le **NAPT (Easy IP)** : tout le LAN sort avec l'adresse de GE0/0/2.
3. Depuis PC1/PC2, pinger 203.0.113.10 et ouvrir la page web (le client HTTP du PC eNSP ou `curl` si dispo). Vérifier les translations.
4. Configurer le **port mapping** : le serveur web interne (192.168.10.100:80) doit être joignable depuis "Internet" via 198.51.100.1:8080.
5. Depuis le Server public (ou un PC côté ISP), accéder à `http://198.51.100.1:8080` → doit afficher le site du serveur interne.
6. Vérifier les sessions NAT et la table de translation.

### Correction pas à pas

**Routage de base :**

```
# R1 : défaut vers l'ISP
[R1]ip route-static 0.0.0.0 0.0.0.0 198.51.100.2

# R-ISP : RIEN vers 192.168.10.0/24 (volontaire). Route vers 198.51.100.0/30 = direct.
```

**NAPT Easy IP sur R1 :**

```
[R1]acl 2000
[R1-acl-basic-2000]rule permit source 192.168.10.0 0.0.0.255
[R1-acl-basic-2000]quit

[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]nat outbound 2000
[R1-GigabitEthernet0/0/2]quit
```

> `nat outbound 2000` **sans adresse** = Easy IP : utilise l'IP de l'interface de sortie. C'est le cas standard d'un site avec une seule IP publique (box, AR720 avec IP WAN).

**Port mapping (NAT server) sur R1 :**

```
[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]nat server protocol tcp global 198.51.100.1 8080 inside 192.168.10.100 80
[R1-GigabitEthernet0/0/2]quit
```

> Avec Easy IP, `global 198.51.100.1` = l'IP de l'interface elle-même. On peut aussi écrire `global current-interface 8080`.

**Tests :**

```
# Depuis PC1 (client) :
PC1> ping 203.0.113.10        # OK via NAT
# Accès web : utiliser le navigateur du PC eNSP vers http://203.0.113.10

# Depuis le Server public (côté ISP) :
# http://198.51.100.1:8080  →  doit afficher le site hébergé en 192.168.10.100
```

### Vérifications

```
[R1]display nat outbound
# Politiques NAT configurées par interface.

[R1]display nat server
# Mappings statiques (port mapping).

[R1]display nat session all
# Sessions actives : Proto, SrcIP:Port -> DstIP:Port, translation.
# Exemple : TCP 192.168.10.11:1234 -> 203.0.113.10:80  traduit en  198.51.100.1:2048 -> 203.0.113.10:80

[R1]display firewall session table   # (selon version, équivalent)
```

**Capture Wireshark** : capturer simultanément côté LAN (SW1↔R1) et côté WAN (R1↔R-ISP) pendant un ping PC1→203.0.113.10 : constater que l'IP source **change** en traversant R1 (192.168.10.11 → 198.51.100.1). C'est la démonstration la plus parlante du NAT.

### Pièges classiques

1. **ACL du NAT trop restrictive ou dans le mauvais sens** : `rule permit source 192.168.10.0 0.0.0.255` — le wildcard est 0.0.0.255 (pas 0.0.255.255 !). Une ACL qui ne matche pas = pas de translation = ping KO sans message d'erreur.
2. **`nat outbound` sur la mauvaise interface** : il se configure sur l'interface de **sortie** (outside, GE0/0/2), pas sur l'interface LAN.
3. **Oublier la route par défaut** sur R1 : le NAT ne crée pas le routage ; sans défaut, le paquet est droppé avant même la translation.
4. **Port mapping : tester depuis le LAN** : depuis PC1, `http://198.51.100.1:8080` peut échouer (hairpin NAT non supporté par défaut) alors que ça marche depuis l'extérieur. Toujours tester le port mapping **depuis l'extérieur**.
5. **Conflit de port** : exposer le port 80 alors que l'interface externe... en lab ça passe, mais en production le port 80 externe peut être pris par l'interface web du routeur lui-même → utiliser 8080 comme dans l'énoncé.
6. **Le serveur interne sans passerelle** : le Server eNSP en 192.168.10.100 doit avoir 192.168.10.254 comme gateway, sinon le retour du port mapping n'arrive jamais.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| Routage de base + absence volontaire de route privée côté ISP (comprise) | 3 |
| ACL 2000 correcte (source + wildcard) | 3 |
| NAPT Easy IP fonctionnel (ping + web depuis PC1/PC2) | 5 |
| Port mapping fonctionnel depuis l'extérieur | 5 |
| Vérifications (`display nat session`, `display nat server`) + double capture commentée | 4 |

### Durée estimée
**1 h 15**.

### Fiche animateur — points à insister
- NAT = **translation d'adresse**, pas du routage et pas de la sécurité (même si ça masque le réseau interne). Distinguer les trois rôles.
- Easy IP couvre 95 % des sites PME (une seule IP publique). Le port mapping couvre l'exposition de services (vidéosurveillance, site web, accès distant).
- La double capture (avant/après translation) est le moment clé : faire formuler par les stagiaires ce qui a changé dans l'en-tête IP.
- Lien terrain : sur vos AR720 en sortie d'agence, le NAT sortant est identique ; sur les USG6000, le NAT se configure dans les politiques (voir TP11) — même logique, autre syntaxe.
- Question piège : "Pourquoi R-ISP ne connaît-il pas 192.168.10.0/24 ?" → Parce que sur Internet, les adresses privées ne sont pas routées ; c'est précisément pour ça que le NAT existe.

---

## 10. TP8 — PPPoE client et serveur

