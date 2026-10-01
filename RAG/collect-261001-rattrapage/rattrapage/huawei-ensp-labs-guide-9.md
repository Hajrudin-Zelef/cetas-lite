---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-9
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1252, 1416]
sha256: 17e231eeef543678a7b0d6a9bfff3d9c68da944630ef539674200ed52ee247d1
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

### Objectif
Comprendre l'établissement d'une session PPPoE (cas classique des accès xDSL/fibre via boîtier opérateur en mode bridge) : discovery, session, authentification PAP/CHAP, IP négociée.

### Prérequis
TP5 (routage), TP7 (NAT — le PPPoE débouche souvent sur du NAT).

### Topologie

```
PC1 ──SW1── GE0/0/1 R-Client (AR2220) GE0/0/2 ══════════ GE0/0/2 R-Fournisseur (AR2220) GE0/0/1── "Internet" (Server 203.0.113.10)
              (client PPPoE, dialer 1)   lien Ethernet   (serveur PPPoE, virtual-template 1)
```

- **R-Fournisseur** : joue le rôle du BAS/BRAS de l'opérateur : serveur PPPoE sur GE0/0/2, pool d'adresses publiques simulées.
- **R-Client** : client PPPoE sur GE0/0/2 via une interface **Dialer1**, authentification CHAP (login `client-lab`, mot de passe `huawei123` — identifiants de lab uniquement).
- Après établissement, R-Client reçoit une IP "publique" sur le Dialer et fait du NAT (Easy IP) pour son LAN.

### Énoncé

1. Configurer **R-Fournisseur** comme serveur PPPoE : virtual-template, pool d'adresses, authentification CHAP, utilisateur local.
2. Configurer **R-Client** : interface Dialer1 (requête d'IP, CHAP, dialer-group), lier GE0/0/2 au PPPoE (`pppoe-client dial-bundle-number 1`).
3. Monter la session, vérifier l'IP obtenue sur le Dialer.
4. Ajouter le NAT Easy IP sur le Dialer + route par défaut via le Dialer.
5. Tester : PC1 ping 203.0.113.10 (à travers PPPoE + NAT).
6. **Dépannage** : l'animateur casse l'authentification (mauvais mot de passe côté client) → diagnostiquer avec les `display` et les debugs.

### Correction pas à pas

**R-Fournisseur (serveur PPPoE) :**

```
[R-F]ip pool POOL-PPPOE
[R-F-ip-pool-POOL-PPPOE]network 198.51.100.0 mask 30
[R-F-ip-pool-POOL-PPPOE]quit
# Le pool fournit ici 198.51.100.1 (le .2 sera pris par... attention : /30 = .1 et .2 utilisables)

[R-F]aaa
[R-F-aaa]local-user client-lab password cipher huawei123
[R-F-aaa]local-user client-lab service-type ppp
[R-F-aaa]quit

[R-F]interface Virtual-Template 1
[R-F-Virtual-Template1]ppp authentication-mode chap
[R-F-Virtual-Template1]ip address 198.51.100.2 30
[R-F-Virtual-Template1]remote address pool POOL-PPPOE
[R-F-Virtual-Template1]quit

[R-F]interface GigabitEthernet 0/0/2
[R-F-GigabitEthernet0/0/2]undo portswitch
[R-F-GigabitEthernet0/0/2]pppoe-server bind virtual-template 1
[R-F-GigabitEthernet0/0/2]quit
```

**R-Client (client PPPoE) :**

```
[R-C]dialer-rule
[R-C-dialer-rule]dialer-rule 1 ip permit
[R-C-dialer-rule]quit

[R-C]interface Dialer 1
[R-C-Dialer1]dialer user client-lab
[R-C-Dialer1]dialer-group 1
[R-C-Dialer1]dialer bundle 1
[R-C-Dialer1]ppp chap user client-lab
[R-C-Dialer1]ppp chap password cipher huawei123
[R-C-Dialer1]ip address ppp-negotiate
[R-C-Dialer1]quit

[R-C]interface GigabitEthernet 0/0/2
[R-C-GigabitEthernet0/0/2]undo portswitch
[R-C-GigabitEthernet0/0/2]pppoe-client dial-bundle-number 1
[R-C-GigabitEthernet0/0/2]quit
```

**Vérifier la session :**

```
[R-C]display pppoe-client session summary
# État UP, Session ID attribuée, interface physique liée.

[R-C]display ip interface brief
# Dialer1 doit avoir une IP (ex. 198.51.100.1) en "PPP-negotiate".
```

**NAT + défaut via le Dialer :**

```
[R-C]acl 2000
[R-C-acl-basic-2000]rule permit source 192.168.10.0 0.0.0.255
[R-C-acl-basic-2000]quit
[R-C]interface Dialer 1
[R-C-Dialer1]nat outbound 2000
[R-C-Dialer1]quit
[R-C]ip route-static 0.0.0.0 0.0.0.0 Dialer 1
```

**Test final :** `PC1> ping 203.0.113.10` → OK (LAN → NAT → PPPoE → fournisseur → Internet simulé).

### Vérifications

```
[R-C]display pppoe-client session summary   # session UP ?
[R-C]display pppoe-client session packet     # compteurs PADI/PADO/PADR/PADS
[R-F]display virtual-access                   # interfaces d'accès virtuelles créées par le serveur
[R-F]display pppoe-server session all         # sessions vues côté serveur
[R-F]display access-user                       # utilisateurs PPP connectés
```

**Capture Wireshark** sur le lien R-C↔R-F : filtrer `pppoe` → observer la phase **Discovery** (PADI → PADO → PADR → PADS) puis la phase **Session** (LCP → authentification CHAP Challenge/Response → IPCP avec l'IP négociée). C'est toute la séquence d'un accès ADSL/fibre résumée en quelques paquets.

### Pièges classiques

1. **Mot de passe CHAP différent des deux côtés** : la session reste Down, `display pppoe-client session summary` ne montre rien. Diagnostic : `debugging ppp chap` (avec modération) ou vérifier les deux configs côte à côte. C'est la panne provoquée de l'énoncé.
2. **PAP vs CHAP incohérents** : serveur en `chap`, client en `pap` → échec d'authentification. Toujours aligner.
3. **`dialer-group` sans `dialer-rule`** : le Dialer ne sait pas quel trafic déclenche l'appel → la session ne monte pas. La règle `dialer-rule 1 ip permit` est obligatoire.
4. **`pppoe-server bind` sur une interface restée en mode switch** : `undo portswitch` d'abord.
5. **Pool PPPoE vide ou mal dimensionné** : avec un /30, une seule adresse attribuable ; un second client échouerait (IPCP nak). En lab c'est voulu pour comprendre, en production dimensionner le pool.
6. **MTU PPPoE** : 1492 au lieu de 1500 (8 octets d'en-tête PPPoE+PPP). Symptôme en production : "les petits sites web passent, les gros chargements bloquent" → régler `mtu 1492` ou activer le TCP MSS adjust. À mentionner.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| Serveur PPPoE complet (virtual-template, pool, AAA, bind) | 6 |
| Client PPPoE complet (Dialer, CHAP, dialer-rule, bind) | 6 |
| Session UP + IP négociée vérifiées | 3 |
| NAT + défaut via Dialer, ping de bout en bout | 3 |
| Diagnostic de la panne d'authentification (méthode + correction) | 2 |

### Durée estimée
**1 h 30** (dont 20 min de capture PPPoE).

### Fiche animateur — points à insister
- PPPoE = 2 phases : **Discovery** (trouver le serveur, 4 paquets PADI/PADO/PADR/PADS) puis **Session** (LCP, authentification, IPCP). Faire retrouver chaque phase dans la capture.
- CHAP ne transmet jamais le mot de passe en clair (challenge/response) contrairement à PAP : d'où sa préférence.
- Lien terrain : quand un boîtier opérateur est en mode **bridge**, c'est l'AR720 qui monte la session PPPoE — exactement ce TP. En mode routeur, le boîtier fait le PPPoE et l'AR720 fait juste du DHCP/NAT derrière.
- Le MTU 1492 est un classique des tickets "ça rame sur certaines applis" sur les accès PPPoE : à garder en tête pour le support.
- Question piège : "Où est configurée l'IP publique du client ?" → Nulle part en statique : elle est **négociée** (IPCP). C'est la différence avec une liaison louée.

---

## 11. TP9 — IPSec site-à-site

### Objectif
Monter un tunnel IPSec entre deux sites : phase 1 (IKE), phase 2 (IPSec), interesting traffic via ACL, vérifications et pièges (ACL miroir, NAT-T).

### Prérequis
TP5 (routage), TP7 (NAT), adressage.

### Topologie

```
Site A (192.168.10.0/24)                  "Internet" (10.0.0.0/30)                  Site B (192.168.20.0/24)
PC1 ──SW1── GE0/0/1 R1 (AR2220) GE0/0/2 ════════════════ GE0/0/2 R2 (AR2220) GE0/0/1──SW2── PC2
              .254          100.64.1.1/30  tunnel IPSec   100.64.1.2/30         .254
```

- **R1** : GE0/0/1 = 192.168.10.254/24, GE0/0/2 = 100.64.1.1/30 (IP publique simulée).
- **R2** : GE0/0/1 = 192.168.20.254/24, GE0/0/2 = 100.64.1.2/30.
- Routage : R1 a une défaut vers 100.64.1.2 ; R2 une défaut vers 100.64.1.1 (ils se joignent directement en lab).
- Clé pré-partagée de lab : `Cle-IPSec-Lab-2026` (jamais en production telle quelle).

### Énoncé

