---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-2
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft", "Oracle", "United States"]
dates: ["2026-09-17", "2026-09-27"]
keywords: ["datacenter", "capex", "luna", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [126, 235]
sha256: b9a3d84ec206e61c853c555e0ba1dc343c2e6cab9721333b06a56c71591f5395
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

- **PCI HSM** (anciennement PCI PTS HSM) : obligatoire si le HSM manipule des clés liées aux
  cartes bancaires (PIN, clés de chiffrement de données carte). Thales propose une gamme
  **Luna Payment HSM** dédiée ; Utimaco a **Atalla** (voir section 25).
- **eIDAS** (UE) : pour la signature électronique qualifiée et l'horodatage qualifié, le
  dispositif de création de signature doit être **QSCD** (Qualified Signature Creation Device) —
  en pratique un HSM certifié Common Criteria. Si vous visez la signature qualifiée, le choix du
  HSM est contraint par cette exigence.

## 8. Usage 1 — PKI et autorités de certification

C'est l'usage historique n°1. La **clé privée de l'AC racine** (root CA) est générée dans le HSM
lors d'une cérémonie (voir section 42), n'en sort jamais, et ne sert qu'à signer les certificats
des AC intermédiaires — quelques signatures par an. Les AC intermédiaires/émettrices peuvent
aussi être sur HSM (souvent une partition du même boîtier, voir section 38). Dimensionnement :
quelques opérations par jour → le plus petit modèle suffit. Outils : **EJBCA** (open source,
voir section 114) ou Microsoft AD CS branché en PKCS#11 sur le HSM. Pour une PME, un **YubiHSM 2**
(~650 €, voir section 29) suffit pour une AC interne ; pour une AC publique ou qualifiée eIDAS,
appliance réseau FIPS 140-2 L3 / Common Criteria obligatoire.

## 9. Usage 2 — Chiffrement bases de données et disques (TDE)

Transparent Data Encryption (Oracle TDE, SQL Server, PostgreSQL via pgcrypto ou chiffrement au
niveau stockage) : la clé maîtresse (master key) qui protège les DEK est stockée dans le HSM, les
opérations de déchiffrement des DEK passent par PKCS#11. Intérêt : même avec les fichiers de base
et les sauvegardes, sans le HSM les données sont inexploitables. Volume : quelques opérations par
minute (rotation, démarrage d'instance) → faible. Point de vigilance : **la disponibilité du HSM
conditionne le démarrage des bases** — prévoir le cluster HA (section 45) et la procédure de
redémarrage après coupure électrique (section 145).

## 10. Usage 3 — Signature électronique et code signing

Signer des documents (factures, contrats), des binaires, des images de firmware, des conteneurs.
Le HSM garantit que la clé de signature ne peut pas être copiée par un développeur ou exfiltrée
par un malware du poste de build. Cas datacenter : signature des **firmwares/updates internes**,
signature des images de déploiement (PXE, golden images). Volume : de quelques signatures/jour
(code) à des milliers/heure (facturation électronique à grande échelle) — dimensionner en
conséquence (section 36). Pour le code signing « moderne » (Sigstore/Fulcio), le HSM reste
pertinent pour la clé racine ; les certificats éphémères sont gérés par l'infrastructure.

## 11. Usage 4 — Horodatage qualifié (RFC 3161)

Une autorité d'horodatage (TSA) signe des jetons d'horodatage prouvant l'existence d'un document
à un instant T. La clé de la TSA vit dans le HSM, et l'horloge doit être fiable (NTP sécurisé,
idéalement NTS). Usage PME : horodatage des journaux d'audit, des sauvegardes, des signatures.
Volume : potentiellement élevé (un jeton par événement) → prévoir le débit. Produits : la plupart
des TSA logicielles (ex. OpenTSA) se branchent en PKCS#11 sur n'importe quel HSM.

## 12. Usage 5 — DNSSEC

La **KSK** (Key Signing Key) d'une zone DNSSEC est typiquement générée et conservée dans un HSM ;
la ZSK (Zone Signing Key) peut tourner plus souvent, parfois en logiciel. Pour un hébergeur ou
une entreprise gérant ses zones, c'est un usage à faible volume mais à forte criticité (la KSK
compromise = usurpation de tout le domaine). Outils : BIND 9 / Knot DNS avec PKCS#11, ou
PowerDNS. Bonne nouvelle : un petit HSM suffit.

## 13. Usage 6 — Paiement : PIN, DUKPT, PCI HSM

Le monde du paiement a ses propres HSM (« payment HSM ») avec des jeux de commandes spécifiques
(vérification de PIN, génération de CVV, DUKPT pour les TPE). **Ne pas utiliser un HSM généraliste
pour du paiement** si la conformité PCI l'exige : il faut un HSM certifié **PCI HSM**. Acteurs
vérifiés le 27/09/2026 : **Thales Luna Payment HSM**, **Utimaco Atalla** (dont la version cloud
**Azure Payment HSM v2**, annoncée le 17/09/2026 avec Marvell LiquidSecurity — public preview
Western US et Western Europe). Débit typique : jusqu'à ~2 000 vérifications PIN/seconde sur un
Luna Payment HSM (product brief Thales).

## 14. Usage 7 — 5G, blockchain, IoT

- **5G** : les HSM gèrent les clés d'authentification des abonnés (Milenage, TUAK — supportés par
  les Luna, voir product brief) et les certificats des fonctions réseau.
- **Blockchain/custody** : génération et stockage des clés de portefeuilles (BIP32/SLIP10
  supportés par Luna), signature de transactions. Les dépositaires d'actifs numériques utilisent
  des HSM en quorum MofN.
- **IoT/industrie** : AC dédiée aux certificats d'équipements (provisioning en usine), signature
  des firmwares OTA. Volume potentiellement massif à l'émission → dimensionner.

## 15. Form factors : panorama (réseau, PCIe, USB, cloud)

| Format | Principe | Latence typique | Partage | Idéal pour |
|---|---|---|---|---|
| **Appliance réseau** | Boîtier 1U/2U, accès via LAN chiffrée | ~1-5 ms (réseau) | Oui, multi-apps/partitions | Datacenter, PKI centrale, mutualisation |
| **Carte PCIe** | Carte dans le serveur, accès local | ~µs-ms (bus local) | Non (dédiée au serveur) | Serveur applicatif à fort débit, latence critique |
| **USB/nano** | Clé USB interne au serveur | ~ms (USB) | Via logiciel (yubihsm-connector) | AC racine offline, PME, budget serré |
| **Cloud HSM** | HSM dédié chez le provider, accès réseau | ~ms (selon région) | Oui (cluster) | Charges cloud-natives, sans investissement CAPEX |

Le choix se fait sur trois axes : **débit/latence**, **mutualisation**, **modèle économique**
(CAPEX vs OPEX). Les sections 16 à 19 détaillent.

## 16. Appliance réseau : architecture et cas d'usage

L'appliance (ex. Thales Luna Network HSM, Utimaco SecurityServer en version LAN, Entrust nShield
Connect) est un serveur durci 19" avec le module crypto inviolable, des alimentations
redondantes, et une pile réseau minimale. Les applications se connectent via un **client**
(PKCS#11, JCE, CAPI/CNG) sur un **canal chiffré et mutuellement authentifié** (NTLS chez Thales).
Le HSM vit dans un **VLAN dédié**, derrière un firewall qui n'autorise que les clients déclarés
(voir section 35). Avantages : mutualisation entre applications via **partitions** isolées
cryptographiquement (jusqu'à 20 sur un Luna Payment HSM), haute disponibilité par clustering,
administration centralisée. Inconvénients : prix d'entrée le plus élevé, latence réseau.

## 17. Carte PCIe : architecture et cas d'usage

La carte (ex. Luna PCIe HSM A700/A750/A790 ou S700/S750/S790, Utimaco Se en version PCIe) s'enfiche
dans un slot du serveur applicatif. L'application y accède **en local**, latence minimale, pas de
dépendance réseau. Idéal pour : serveur TLS à très fort débit (terminaison SSL avec clé protégée),
serveur de signature à haute cadence, appliance de sécurité embarquant sa crypto. Limites :
**un serveur = un HSM** (pas de mutualisation simple), la HA exige une carte par nœud, et la
carte consomme un slot PCIe + ~14-18 W (chiffres Luna PCIe vérifiés : 14 W typique, 18 W max).
En virtualisation, prévoir le **PCI passthrough** (IOMMU/VT-d) vers la VM dédiée.

## 18. USB/nano : architecture et cas d'usage

