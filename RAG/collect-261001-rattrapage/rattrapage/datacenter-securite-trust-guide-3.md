---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-3
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Microsoft", "United States"]
dates: ["2026-09-17", "2026-09-27"]
keywords: ["datacenter", "aws", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [236, 355]
sha256: 681f84a59ca7f26ccc5e8fd8f437bbcccccbfcaaa04f714132bb29effda98709
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Le **YubiHSM 2** (vérifié le 27/09/2026) : format nano USB-A (12 × 13 × 3,1 mm, 1 g), 20 mA
moyens / 30 mA max, RSA 2048/3072/4096, ECC P-256/P-384/P-521/Curve25519, AES, SHA, HMAC,
interfaces PKCS#11 / Microsoft CNG (KSP) / SDK natif C et Python. Il se branche sur un **port USB
interne** du serveur et y reste. Un connecteur réseau (yubihsm-connector) permet de le partager
entre serveurs via sessions mutuellement authentifiées. Cas d'usage roi : **clé d'AC racine
offline** (la clé ne sort jamais, le boîtier tient dans un coffre), PKI interne PME, AD CS.
Limites : débit modeste (dizaines d'opérations/s, pas des milliers), pas de clustering natif —
on duplique en générant la même clé sur 2 modules via cérémonie, ou on accepte le point unique
de défaillance pour une racine offline (acceptable : elle ne sert que quelques fois par an).

## 19. Cloud HSM : principe et limites de confiance

Le provider vous alloue un **HSM physique dédié** (single-tenant) dans son datacenter ;
vous en êtes le seul client (Crypto Officer), le provider n'a pas accès à vos clés. Références
vérifiées le 27/09/2026 :

| Service | Prix indicatif vérifié | Certification | Notes |
|---|---|---|---|
| **AWS CloudHSM** | ~1,45–1,50 $/h/HSM | FIPS 140-2 L3 | PKCS#11, JCE, OpenSSL ; clusters multi-AZ |
| **Azure Dedicated HSM** | **4,85 $/h/HSM** (page tarifs Azure) | FIPS 140-2 L3 | Régions limitées ; basé sur Thales Luna |
| **Azure Managed HSM** | ~3,20 $/h (indicatif, à vérifier) | FIPS 140-2 L3 | Géré, multi-tenant logique |
| **GCP Cloud HSM** (via Cloud KMS) | ~1 $/mois/clé + requêtes (indicatif) | FIPS 140-2 L3 | Intégré à Cloud KMS |

Limites honnêtes : vous **faites confiance au provider** pour l'intégrité physique du datacenter
et l'absence de backdoor d'administration ; la latence dépend de la région ; le coût mensuel
d'un cluster HA (2-3 HSM) tourne autour de **2 200–3 500 $/mois** pour AWS CloudHSM — à comparer
au TCO d'une appliance achetée (section 48). Pour une PME sans datacenter, c'est souvent le
meilleur rapport sécurité/effort.

## 20. Thales Luna Network HSM — fiche vérifiée le 27/09/2026

(Source : product briefs Thales / cpl.thalesgroup.com ; chiffres du brief « BBv37 avril 2024 »
et du brief T-Series.)

- **Formats** : appliance 19" 1U (482,6 × 533,4 × 43,8 mm), 12,7 kg, **double alimentation
  hot-swap**.
- **Électrique** : 100-240 V, 50-60 Hz ; **110 W max, 84 W typique** ; dissipation
  376 BTU/h max. Plage 0-35 °C en fonctionnement.
- **Fiabilité** : MTBF **171 308 h** (~19,5 ans).
- **Modèles** : deux gammes (A : authentification par mot de passe ; S : authentification
  multifacteur PED), 3 niveaux de performance chacune. Ex. T-Series : T-2000 (1 400 tps
  RSA-2048, 3 000 tps ECC P-256) → T-5000 (14 000 tps RSA-2048, 16 000 tps ECC P-256).
- **Partitions** : jusqu'à 20 partitions isolées cryptographiquement (modèle paiement),
  politiques de partitions flexibles.
- **Crypto** : RSA, DSA, DH, ECC (ECDSA/ECDH/Ed25519/ECIES, courbes NIST/Brainpool),
  AES/AES-GCM, 3DES, SHA-1/2/3, SM2/SM3/SM4, BIP32/SLIP10, Milenage/TUAK (5G).
- **RNG** : source de bruit matérielle + CTR-DRBG NIST 800-90A, conception visant AIS 20/31.
- **PQC** : **Luna PQC Functionality Module** mentionné dans la documentation récente
  (mécanismes post-quantiques dans le HSM — voir section 152).
- **APIs** : PKCS#11, JCA/JCE, Microsoft CAPI/CNG, OpenSSL, REST d'administration ;
  Functionality Modules (code custom dans l'enclave).
- **Certifications visées** : FIPS 140-2 niveau 3 (historique de la gamme ; vérifier la
  validation exacte du modèle/firmware sur le site NIST CMVP — **à vérifier** au moment
  de l'achat).

## 21. Thales Luna PCIe HSM — fiche vérifiée le 27/09/2026

(Source : product brief Luna PCIe HSM, « BBv37 avril 2024 ».)

- **Format** : carte PCIe **low profile**, PCIe 2.0, 69,6 × 167 × 18,7 mm.
- **Électrique** : **18 W max, 14 W typique** — à prévoir dans le budget thermique du serveur.
- **Fiabilité** : MTBF **997 508 h** (~114 ans — valeur constructeur, indicateur, pas promesse).
- **Modèles** : A700/A750/A790 (mot de passe) et S700/S750/S790 (PED multifacteur).
  Performances : **RSA-2048 : 1 000 / 5 000 / 10 000 tps** ; **ECC P-256 : 2 000 / 10 000 /
  22 000 tps** ; **AES-GCM : 2 000 / 10 000 / 17 000 tps** selon modèle.
- **Mémoire** : 4 / 32 / 64 Mo selon modèle (nombre de clés/objets stockables).
- **Cas d'usage** : serveur applicatif dédié, appliance de sécurité embarquée, PKI locale
  à fort débit.

## 22. Thales Luna USB HSM et DPoD (cloud)

- **Luna USB HSM** : format portable pour **stockage offline** des clés (sauvegarde,
  transport sécurisé, racine d'AC hors-ligne). Petit volume, usage ponctuel.
- **DPoD (Data Protection on Demand)** : plateforme cloud de Thales — HSM à la demande
  sans acheter de matériel, marketplace de services (Luna Cloud HSM, gestion de clés).
  Modèle OPEX, à comparer avec AWS CloudHSM/Azure (section 19). Tarifs : **à vérifier**
  sur devis.

## 23. Utimaco SecurityServer Se Gen2 — fiche vérifiée le 27/09/2026

(Source : datasheet constructeur ; communiqué Utimaco septembre 2026.)

- **Formats** : **carte PCIe** ou **appliance réseau (LAN)** — même logiciel, deux
  emballages.
- **Crypto** : RSA, DSA, ECDSA (courbes NIST et Brainpool), DH/ECDH, AES, 3DES, MAC/CMAC/
  HMAC, SHA-1/SHA-2/**SHA-3** (depuis firmware 4.10), RIPEMD, DRNG déterministe.
- **Authentification** : **« n out of m »** (quorum MofN natif), séparation des tâches,
  smartcard pour l'authentification forte des opérateurs.
- **Administration** : à distance extensive, SNMP, mises à jour de firmware à distance,
  simulateur logiciel pour les tests d'intégration.
- **Standards** : PKCS#11, Microsoft CSP/CNG, JCE.
- **Certifications** : **FIPS 140-2 niveau 3** (niveau 4 annoncé « en cours » — à vérifier),
  **PCI PTS HSM V3**, certification **Common Criteria** (CCN espagnol, septembre 2026).

## 24. Utimaco u.trust General Purpose HSM — la gamme conteneurisée (vérifié)

(Source : communiqué Utimaco, septembre 2026 — vérifié le 27/09/2026.)

- **Architecture conteneurisée** inspirée du cloud : jusqu'à **31 tenants/conteneurs**,
  plusieurs partitions PKCS#11 par conteneur — séparation applicative fine.
- **Modèles** : Se100, Se2k, Se5k, Se15k, Se40k (montée en charge par licence/field upgrade),
  stockage interne ou externe des clés.
- **Crypto-agile** : support des algorithmes **post-quantiques sans changement de matériel**
  (voir section 152 : package **Quantum Protect**).
- **Packages solutions** : 5G Protect, Blockchain Protect, Double Key Encryption,
  Quantum Protect.
- **Disponible** : en appliance LAN, carte PCIe, ou **as-a-service**.

## 25. Utimaco Atalla AT1000 — le HSM paiement (à vérifier)

La gamme **Atalla** est la référence historique du paiement. Vérifié le 27/09/2026 :
**Azure Payment HSM v2** combine la technologie **Marvell LiquidSecurity**, le module
**Utimaco Atalla Payments** et Azure — public preview depuis le **17/09/2026**
(Western US, Western Europe), ciblant PCI sans gérer de HSM physiques (communiqués Marvell,
Microsoft, Utimaco). Pour un datacenter on-prem avec besoin paiement : Atalla AT1000 en
appliance — caractéristiques détaillées **non trouvées au 27/09/2026**, demander le datasheet
au vendeur.

## 26. Entrust nShield — panorama (détails à vérifier)

