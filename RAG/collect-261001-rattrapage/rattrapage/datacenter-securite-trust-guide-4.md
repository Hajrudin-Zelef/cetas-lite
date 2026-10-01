---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-4
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "EU", "Microsoft"]
dates: ["2026-09-27"]
keywords: ["datacenter", "aws", "capex", "hyperscaler", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [356, 467]
sha256: baecddc8731fae7a107eb84e7c22f71a999fdb5af8a61c1af2bad1fe05bb9a72
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Entrust (ex-nCipher) **nShield** est le troisième grand acteur historique avec Thales et
Utimaco : gammes **nShield Connect** (réseau), **nShield Edge** (USB), **nShield 5c**
(PCIe). Points différenciants connus : architecture **Security World** (clés partagées et
protégées entre plusieurs HSM via cartes d'opérateurs), outil **nShield Monitor**.
Fiche technique détaillée et tarifs : **non trouvés au 27/09/2026** (constructeur sur devis)
— **à vérifier** avant tout comparatif d'achat.

## 27. Futurex — panorama (détails à vérifier)

**Futurex** (Texas, USA) : HSM généralistes et de paiement (séries Vectra, Excrypt),
positionnés sur la performance et le marché américain. Intégration PKCS#11/JCE classique,
clustering. Fiche détaillée : **non trouvée au 27/09/2026** — **à vérifier** si un appel
d'offres les inclut.

## 28. Marvell LiquidSecurity — le HSM paiement cloud (vérifié)

**Marvell LiquidSecurity** est la brique matérielle derrière **Azure Payment HSM v2**
(voir section 25) : des HSM conçus pour le cloud, exposés en service managé. Tendance 2026 :
le HSM « cloud-scale » opéré par l'hyperscaler, le client gardant la souveraineté des clés.
À suivre pour les workloads paiement qui refusent le CAPEX.

## 29. YubiHSM 2 — fiche vérifiée le 27/09/2026 + prix

(Sources : Yubico via revendeurs — prix constatés septembre 2026 ; specs constructeur.)

| Caractéristique | Valeur vérifiée |
|---|---|
| Format | Nano **USB-A**, 12 × 13 × 3,1 mm, **1 g** — tient dans un port USB interne |
| Alimentation | **20 mA moyen / 30 mA max** (négligeable) |
| Algos | RSA 2048/3072/4096, ECC P-256/P-384/P-521/Curve25519, **Ed25519**, AES 128/192/256, SHA-1/256/384/512, HMAC |
| Interfaces | PKCS#11, Microsoft CNG (KSP), SDK natif C/Python, yubihsm-connector (partage réseau) |
| Audit | Journal chaîné par hash (**tamper-evident**) |
| Version FIPS | **YubiHSM 2 FIPS** : FIPS 140-2 (niveau global 2, physique niveau 3) |
| Prix constatés | **~650 €** version standard ; **~949 €** version FIPS (revendeur EU, 2026) — **à vérifier** au jour de l'achat |
| PQC | Firmware 2.4+ : PQC **en test pilote**, pas de support complet (suivi pqca, 09/2026) |

C'est le **meilleur rapport sécurité/prix** pour : AC racine offline, AD CS en PME,
signature de code d'équipe, labo. Ce n'est PAS un remplacement d'appliance réseau pour des
milliers d'opérations/seconde.

## 30. Tableau comparatif général des HSM

| Critère | Thales Luna Network | Thales Luna PCIe | Utimaco Se Gen2 / u.trust GP | YubiHSM 2 | Cloud HSM (AWS/Azure) |
|---|---|---|---|---|---|
| Format | Appliance 1U | Carte PCIe LP | Appliance ou PCIe | Nano USB | Service managé |
| Partage réseau | Oui (partitions) | Non | Oui (31 conteneurs u.trust) | Via connector | Oui (cluster) |
| Débit RSA-2048 | 1 400–14 000 tps | 1 000–10 000 tps | Gamme Se100→Se40k (à vérifier) | Dizaines de tps | ~10 000 ops/s/HSM (indicatif) |
| FIPS | 140-2 L3 (gamme) | 140-2 L3 | 140-2 L3 (L4 en cours) | 140-2 (version FIPS) | 140-2 L3 |
| Conso | 84 W typ. | 14 W typ. | À vérifier (~50-100 W appliance) | 0,1 W | 0 (chez le provider) |
| Prix d'entrée | ~15–30 k€ (indicatif) | ~5–12 k€ (indicatif) | ~10–25 k€ (indicatif) | 650–949 € | ~1 100 €/mois/HSM (AWS) |
| Idéal pour | Datacenter mutualisé | Serveur dédié haut débit | Datacenter, PQC-ready | PME, racine offline | Cloud-native, OPEX |

Prix matériels : **ordres de grandeur constatés sur le marché, à vérifier sur devis** — les
constructeurs ne publient pas de tarifs publics.

## 31. Tableau des performances chiffrées (tps = transactions/s)

Chiffres **constructeur vérifiés** (briefs avril 2024) — maximums soutenus en conditions de
test, pas des garanties SLA :

| Produit / modèle | RSA-2048 | RSA-4096 | ECC P-256 | AES-GCM |
|---|---|---|---|---|
| Luna Network T-2000 | 1 400 | 350 | 3 000 | — |
| Luna Network T-5000 | 14 000 | 3 500 | 16 000 | — |
| Luna PCIe A700/S700 | 1 000 | — | 2 000 | 2 000 |
| Luna PCIe A750/S750 | 5 000 | — | 10 000 | 10 000 |
| Luna PCIe A790/S790 | 10 000 | — | 22 000 | 17 000 |
| Luna Payment HSM | — | — | — | 2 000 vérif. PIN/s |

Règle d'ingénieur : **diviser par 2** les chiffres constructeur pour le dimensionnement
(marge réseau, latence applicative, pics), et **ne jamais dimensionner au pic théorique**.

## 32. Tableau des certifications par produit (vérifié le 27/09/2026)

| Produit | FIPS 140-2/3 | Common Criteria | PCI HSM | eIDAS/QSCD |
|---|---|---|---|---|
| Thales Luna (gammes) | L3 (vérifié, gamme) | Selon modèle (à vérifier) | Gamme Payment dédiée | Selon modèle (à vérifier) |
| Utimaco Se/u.trust GP | L3, L4 en cours (vérifié 09/2026) | Oui — CCN Espagne 09/2026 (vérifié) | PCI PTS HSM V3 (vérifié) | Selon modèle (à vérifier) |
| YubiHSM 2 FIPS | 140-2 (global L2, physique L3) | Non trouvé | Non | Non |
| AWS CloudHSM | 140-2 L3 | — | Via configuration | Non |

Toujours vérifier la **version exacte firmware/matériel** sur la liste CMVP du NIST au moment
de l'achat : une certification « de gamme » ne couvre pas automatiquement votre modèle.

## 33. Où placer le HSM : schéma datacenter

```
                    ┌─────────────────────────────────────────────────┐
                    │              DATACENTER / SALLE SÉCURISÉE       │
                    │                                                 │
  Internet / WAN ──▶│  ┌─────────┐   ┌──────────┐   ┌──────────────┐   │
                    │  │ Firewall│──▶│  VLAN    │──▶│  HSM réseau  │   │
                    │  │ périm.  │   │  MGMT    │   │  (cluster    │   │
                    │  └─────────┘   │  SÉCURISÉ│   │   2 nœuds)   │   │
                    │                │ 10.99.0.0│   └──────┬───────┘   │
                    │                │  /24     │          │ PKCS#11   │
                    │                └──────────┘          │ (TLS mutuel)
                    │       ┌──────────────┐              │           │
                    │       │ VLAN SERVEURS│◀─────────────┘           │
                    │       │ applicatifs  │   seuls les clients     │
                    │       │ (PKI, AD CS, │   déclarés (IP+cert)    │
                    │       │  Vault, TSA) │                         │
                    │       └──────────────┘                         │
                    │                                                 │
                    │  ┌──────────────┐   ┌──────────────────────┐    │
                    │  │ Coffre-fort  │   │ Baie 42U verrouillée │    │
                    │  │ YubiHSM racine│  │ alim. secourue (UPS) │    │
                    │  │ (AC offline) │   │ double PDU A+B       │    │
                    │  └──────────────┘   └──────────────────────┘    │
                    └─────────────────────────────────────────────────┘
```

