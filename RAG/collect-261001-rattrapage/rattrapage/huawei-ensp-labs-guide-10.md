---
id: collect-261001-rattrapage/rattrapage/huawei-ensp-labs-guide-10
title: "Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ensp_labs_guide.md
source_anchor: ""
source_lines: [1417, 1576]
sha256: 6d4c48d1f3d7b1a14be2b070aec2593c2894b256ebe13c12c4d1a1c496881a5f
---

# Labs eNSP pour former son équipe — Travaux pratiques Huawei corrigés

1. Configurer la **phase 1 (IKE)** : proposal (AES-256, SHA-256, DH group 14), peer avec pre-shared-key, des deux côtés.
2. Configurer la **phase 2 (IPSec)** : proposal ESP (AES-256 + SHA-256), ACL d'**interesting traffic** (trafic Site A ↔ Site B), policy qui lie le tout.
3. Appliquer la policy IPSec sur les interfaces externes (GE0/0/2).
4. Vérifier : `display ike sa`, `display ipsec sa` → tunnel établi.
5. Tester : PC1 ping PC2 → passe **dans** le tunnel (vérifier les compteurs IPSec qui augmentent).
6. **Pièges à provoquer** (un par un, diagnostiquer à chaque fois) :
   - a) ACL non miroir (oubli du sens retour sur R2).
   - b) Clé pré-partagée différente d'un côté.
   - c) Policy IPSec appliquée sur la mauvaise interface.

### Correction pas à pas

**R1 — Phase 1 (IKE) :**

```
[R1]ike proposal 10
[R1-ike-proposal-10]encryption-algorithm aes-256
[R1-ike-proposal-10]authentication-algorithm sha2-256
[R1-ike-proposal-10]dh group14
[R1-ike-proposal-10]quit

[R1]ike peer R2
[R1-ike-peer-R2]pre-shared-key simple Cle-IPSec-Lab-2026
[R1-ike-peer-R2]ike-proposal 10
[R1-ike-peer-R2]remote-address 100.64.1.2
[R1-ike-peer-R2]quit
```

**R1 — Phase 2 (IPSec) :**

```
[R1]ipsec proposal PROP-A
[R1-ipsec-proposal-PROP-A]esp authentication-algorithm sha2-256
[R1-ipsec-proposal-PROP-A]esp encryption-algorithm aes-256
[R1-ipsec-proposal-PROP-A]quit

# Interesting traffic : le trafic à protéger (les deux sens du point de vue de R1)
[R1]acl 3000
[R1-acl-adv-3000]rule permit ip source 192.168.10.0 0.0.0.255 destination 192.168.20.0 0.0.0.255
[R1-acl-adv-3000]quit

[R1]ipsec policy POLICY-A 10 isakmp
[R1-ipsec-policy-POLICY-A-10]security acl 3000
[R1-ipsec-policy-POLICY-A-10]ike-peer R2
[R1-ipsec-policy-POLICY-A-10]proposal PROP-A
[R1-ipsec-policy-POLICY-A-10]quit

[R1]interface GigabitEthernet 0/0/2
[R1-GigabitEthernet0/0/2]ipsec policy POLICY-A
[R1-GigabitEthernet0/0/2]quit
```

**R2 — miroir exact :**

```
[R2]ike proposal 10
[R2-ike-proposal-10]encryption-algorithm aes-256
[R2-ike-proposal-10]authentication-algorithm sha2-256
[R2-ike-proposal-10]dh group14
[R2-ike-proposal-10]quit

[R2]ike peer R1
[R2-ike-peer-R1]pre-shared-key simple Cle-IPSec-Lab-2026
[R2-ike-peer-R1]ike-proposal 10
[R2-ike-peer-R1]remote-address 100.64.1.1
[R2-ike-peer-R1]quit

[R2]ipsec proposal PROP-A
[R2-ipsec-proposal-PROP-A]esp authentication-algorithm sha2-256
[R2-ipsec-proposal-PROP-A]esp encryption-algorithm aes-256
[R2-ipsec-proposal-PROP-A]quit

# ACL MIROIR : source/destination inversées par rapport à R1
[R2]acl 3000
[R2-acl-adv-3000]rule permit ip source 192.168.20.0 0.0.0.255 destination 192.168.10.0 0.0.0.255
[R2-acl-adv-3000]quit

[R2]ipsec policy POLICY-A 10 isakmp
[R2-ipsec-policy-POLICY-A-10]security acl 3000
[R2-ipsec-policy-POLICY-A-10]ike-peer R1
[R2-ipsec-policy-POLICY-A-10]proposal PROP-A
[R2-ipsec-policy-POLICY-A-10]quit

[R2]interface GigabitEthernet 0/0/2
[R2-GigabitEthernet0/0/2]ipsec policy POLICY-A
```

> **Important** : le tunnel ne monte que quand du trafic "intéressant" le déclenche (ou avec `sa trigger` selon version). Lancer `PC1> ping 192.168.20.11` pour déclencher la négociation.

### Vérifications

```
[R1]display ike sa
# Phase 1 : état RD|ST (Ready/StayAlive) = tunnel IKE établi.
# Si vide ou état négociant : problème phase 1 (clé, proposal, joignabilité du peer).

[R1]display ipsec sa
# Phase 2 : une SA dans chaque sens (inbound/outbound), avec les SPI, compteurs de paquets.

[R1]display ipsec statistics
# Compteurs chiffrés/déchiffrés : lancer des pings et voir les compteurs augmenter
# = preuve que le trafic passe DANS le tunnel.

[R1]display ike proposal
[R1]display ipsec policy
```

**Capture Wireshark** sur le lien inter-routeurs : sans IPSec on verrait des ICMP en clair ; avec le tunnel, on voit des paquets **ESP** (protocole 50) entre 100.64.1.1 et 100.64.1.2, illisibles. C'est la démonstration de la confidentialité. On distingue aussi au début les échanges **IKE sur UDP 500**.

### Pièges classiques

1. **ACL non miroir** (panne a) : R2 avec `source 192.168.10.0 destination 192.168.20.0` (copié-collé de R1 sans inverser) → la phase 2 ne négocie pas les bons proxy-ID → `display ipsec sa` vide d'un côté. **Règle d'or : l'ACL de R2 est le miroir de celle de R1.**
2. **Clé pré-partagée différente** (panne b) : la phase 1 échoue → `display ike sa` vide. Message peu explicite : toujours revérifier la clé en premier (erreur la plus fréquente en production aussi).
3. **Policy appliquée sur la mauvaise interface** (panne c) : sur GE0/0/1 (LAN) au lieu de GE0/0/2 (WAN) → le trafic n'est jamais intercepté. `display ipsec policy` montre où la policy est appliquée.
4. **Proposals incompatibles** : AES-256 d'un côté, 3DES de l'autre → échec de négociation. Aligner strictement.
5. **NAT devant IPSec** : si un NAT se trouve entre les deux routeurs (cas d'un site derrière une box), l'ESP (protocole 50, pas de ports) ne traverse pas le NAPT → activer **NAT-T** (encapsulation ESP dans UDP 4500). En lab, à évoquer ; sur eNSP : `ike nat-traversal` si supporté. En production derrière box : quasi systématique.
6. **Le trafic interessant inclut le trafic de gestion** : si l'ACL est trop large (ex. `permit ip`), même les pings entre les IP publiques passent dans le tunnel → à éviter, l'ACL doit être précise.
7. **Oublier que le tunnel est déclenché par le trafic** : après configuration, `display ike sa` peut rester vide jusqu'au premier ping. Toujours générer du trafic avant de conclure à une panne.

### Barème indicatif (20 points)

| Critère | Points |
|---|---|
| Phase 1 IKE complète et symétrique (proposal + peer + PSK) | 5 |
| Phase 2 complète (proposal ESP + ACL interesting traffic) | 4 |
| ACL miroir correcte sur R2 | 3 |
| Policy créée et appliquée sur les bonnes interfaces | 3 |
| Vérifications (`ike sa`, `ipsec sa`, compteurs) + capture ESP commentée | 3 |
| Diagnostic d'au moins une panne provoquée | 2 |

### Durée estimée
**1 h 30** (dont 20 min de pannes provoquées).

### Fiche animateur — points à insister
- Deux phases, deux négociants : **IKE** (phase 1 : "peut-on se parler en secret ?") puis **IPSec** (phase 2 : "quel trafic protège-t-on et comment ?"). Faire réciter.
- L'**interesting traffic** est le cœur du policy-based IPSec : tout ce qui matche l'ACL est chiffré, le reste passe en clair. D'où l'importance d'ACL précises et miroirs.
- La capture ESP illisible = la meilleure preuve de la confidentialité. Comparer avec une capture sans IPSec (pings en clair).
- Lien terrain : IPSec site-à-site entre agences sur vos AR720 ; NAT-T dès qu'une box opérateur est devant.
- Question piège : "Le ping entre les deux IP publiques (100.64.1.1 ↔ 100.64.1.2) est-il chiffré ?" → Non, il ne matche pas l'ACL d'interesting traffic. Seul le trafic LAN-à-LAN l'est.

---

## 12. TP10 — WLAN : AC + AP en mode Fit

### Objectif
Déployer un réseau Wi-Fi d'entreprise en mode **Fit AP** : un contrôleur (AC) centralise la configuration, les AP diffusent le SSID, les clients s'associent avec authentification WPA2-PSK.

### Prérequis
TP1 (VLAN), adressage. Notions Wi-Fi (SSID, PSK).

### Topologie

```
STA1 (client Wi-Fi)               STA2
   \  ))  wifi  ((  /              (wifi)
    ╲                ╱
   AP1 (AP6010DN) ──câble── GE0/0/1 SW1 (S3700, PoE simulé) GE0/0/2 ── GE0/0/1 AC1 (AC6605)
   AP2 (AP6010DN) ──câble── GE0/0/3            VLAN 100 (gestion AP)
```

