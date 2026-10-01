---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-5
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["apache", "incident", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [468, 602]
sha256: 463bb67d5a5c5231751d8a9507b93393bbd5245335599ce73c2c90955bd3647a
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Principes : le HSM **n'est jamais sur le VLAN utilisateurs**, jamais exposé sur Internet,
jamais sur le même VLAN que les serveurs web. Un **VLAN crypto dédié** (/24 suffit), firewall
avec règles explicites par IP cliente + authentification mutuelle TLS/PKCS#11. La racine
offline (YubiHSM ou Luna USB) dort dans un **coffre**, pas dans un serveur allumé.

## 34. Où placer le HSM : schéma PME (budget serré)

```
  ┌─ Serveur PKI/AD CS (ou VM dédiée) ─────────────────────┐
  │  - Windows Server + AD CS  OU  EJBCA sur Linux          │
  │  - YubiHSM 2 sur port USB INTERNE (invisible de l'ext.) │
  │  - yubihsm-connector si partage entre 2 serveurs        │
  │  - Sauvegarde : 2e YubiHSM (même clé, cérémonie) au coffre│
  └────────────────────────────────────────────────────────┘
  Coût : ~1 300 € (2 × YubiHSM) + serveur existant.
  Quand passer à l'appliance : > 100 signatures/heure, besoin HA,
  audit externe, ou clés de plus de 3 applications.
```

## 35. Zone réseau du HSM : VLAN dédié, firewall

Checklist réseau minimale, quel que soit le format :

- [ ] VLAN ou sous-réseau **dédié** au HSM (pas de routage vers Internet).
- [ ] Firewall stateful : n'autoriser que les **IP des clients PKCS#11** + le poste
      d'administration, sur les **ports documentés** du constructeur uniquement.
- [ ] **TLS mutuel** (NTLS ou équivalent) entre clients et HSM : certificats clients
      générés à l'initialisation, stockés hors du HSM si besoin.
- [ ] Pas de DHCP sur ce VLAN : IP statiques, documentation à jour.
- [ ] Supervision : SNMP/syslog du HSM vers le SIEM, alertes sur échecs d'authentification
      (tentatives = incident potentiel).
- [ ] Le poste d'administration du HSM est une **machine dédiée**, durcie, sans Internet,
      accès physique contrôlé.

## 36. Dimensionnement : méthode de calcul ops/s

La question « quel modèle de HSM ? » se répond en 4 étapes :

1. **Inventorier les opérations** : pour chaque application, type d'opération (signature RSA,
   déchiffrement, génération), taille de clé, fréquence moyenne et **pic**.
2. **Convertir en ops/s au pic** : ex. une AC qui émet 3 600 certificats/heure en pic de
   renouvellement = 1 signature/s. Un serveur de signature de factures à 36 000 factures/heure
   = 10 signatures/s. Une TSA à 1 M de jetons/jour = ~12/s en moyenne, ~50/s en pic.
3. **Appliquer le facteur de sécurité ×2** minimum (pics, croissance 3 ans, overhead réseau).
4. **Choisir le modèle** dont le débit constructeur ÷ 2 couvre le besoin (voir section 31).

Note : les opérations **symétriques** (AES) sont 10 à 100× moins coûteuses que l'asymétrique ;
le dimensionnement se fait presque toujours sur **RSA/ECDSA**. Et préférez **ECDSA P-256**
à RSA-2048 pour les nouveaux usages : ~2× plus rapide à sécurité équivalente (chiffres Luna :
2 000 vs 1 000 tps sur A700).

## 37. Dimensionnement : 3 exemples chiffrés

**Exemple A — PME, PKI interne (AD CS + EJBCA), 500 utilisateurs.**
Émission : ~2 000 certificats/an + renouvellements. Soit < 1 signature/heure en moyenne,
pic à ~20 signatures/heure lors d'un renouvellement massif. **Besoin : < 1 op/s.**
→ **YubiHSM 2** largement suffisant. Coût : ~650 €.

**Exemple B — ETI, signature de documents + TDE + Vault, 3 applications.**
Signature : 50 000 documents/jour = ~0,6/s moyen, pic 5/s. TDE : négligeable. Vault :
unseal et rotation, ~10 ops/s en pic. **Besoin : ~20 ops/s avec marge.**
→ **Luna Network T-2000** (1 400 tps RSA-2048 ÷ 2 = 700 ops/s utiles) ou **Utimaco Se2k**,
largement couverts, avec 3 partitions (une par usage). Coût indicatif : 15–25 k€.

**Exemple C — Prestataire, TSA + signature de masse, 10 M opérations/jour.**
10 M/jour = ~116/s moyen, pic ×4 = ~460/s. **Besoin : ~1 000 ops/s avec marge.**
→ **Luna Network T-5000** (14 000 tps ÷ 2 = 7 000 utiles) ou **2× T-2000 en cluster HA**.
Coût indicatif : 25–50 k€ + redondance.

## 38. Partitions et multi-tenancy

Une **partition** est un HSM logique isolé dans le HSM physique : clés, politiques et
administrateurs séparés, comme si c'était un HSM indépendant. Intérêts :

- **Séparation des usages** : partition « AC racine », partition « TDE », partition
  « signature » — une compromission d'un mot de passe opérateur n'expose qu'une partition.
- **Multi-entités** : héberger les clés de plusieurs clients/entités sur un seul boîtier
  (modèle MSP/hébergeur).
- **Séparation des rôles** : chaque partition a ses propres Crypto Officers.

Limites : les partitions **partagent le débit** du HSM physique et, selon les modèles, un
même firmware. Pour une séparation forte (ex. racine offline vs production), préférez deux
boîtiers physiques. Chez Utimaco u.trust, le découpage va jusqu'à **31 conteneurs**
(voir section 24).

## 39. APIs : PKCS#11, le socle universel

**PKCS#11** (Cryptoki) est l'API standard d'accès aux HSM : ouvrir une session, lister les
objets, signer, déchiffrer, générer. 95 % des logiciels (EJBCA, OpenSSL via engine/provider,
Java via SunPKCS11, BIND, OpenTSA) parlent PKCS#11. Points pratiques :

- Le HSM est vu comme un **slot** contenant un **token** ; chaque partition = un slot.
- Les clés peuvent être marquées **non extractibles** (CKA_EXTRACTABLE=false) : le
  minimum vital pour une clé racine.
- Les **sessions** et les **PIN** (User PIN, SO PIN) : le SO (Security Officer) administre,
  l'utilisateur normal utilise. Ne jamais utiliser le SO au quotidien.
- Tester avec `pkcs11-tool` (OpenSC) avant d'intégrer : `pkcs11-tool --module
  /usr/lib/libCryptoki2.so -L` liste les slots.

## 40. APIs : CNG/CAPI, JCE, OpenSSL engine, REST

- **Microsoft CNG (KSP) / CAPI** : pour AD CS, IIS, les applis Windows. Le YubiHSM 2
  fournit un KSP ; Thales/Utimaco fournissent leurs CSP/KSP.
- **Java JCE** : via le provider **SunPKCS11** (fichier de config pointant sur la lib
  PKCS#11 du HSM) ou les providers natifs (ex. Luna JCP). Indispensable pour EJBCA,
  les applis Java.
- **OpenSSL engine/provider** : permet à nginx, Apache, stunnel d'utiliser la clé du HSM
  sans la copier. En OpenSSL 3.x, le modèle **provider** remplace les engines.
- **REST d'administration** : les appliances modernes exposent une API REST pour le
  provisionnement et le monitoring (pas pour les opérations crypto sensibles, qui restent
  en PKCS#11). Utile pour l'automatisation (Terraform/Ansible — avec parcimonie et
  revue, voir piège n°7).

## 41. Sauvegarde de clés : principes

Un HSM protège les clés contre le vol, **pas contre la perte** : panne matérielle, incendie,
erreur d'effacement. Sans sauvegarde, la clé est perdue à jamais — et avec elle, toutes les
données chiffrées. Les méthodes :

1. **HSM de sauvegarde dédié** (ex. Luna Backup HSM, ou un 2e boîtier) : réplication
   chiffrée de partition à partition, clés jamais en clair hors HSM. La méthode reine.
2. **Sauvegarde chiffrée** : export des clés **wrappées** (chiffrées par une KEK) vers un
   support externe, stocké en coffre. La KEK elle-même est partagée en quorum MofN
   (section 43).
3. **Duplication à la génération** : générer la clé simultanément sur 2 HSM pendant la
   cérémonie (pratique courante avec 2 YubiHSM pour une racine).

Règle : **tester la restauration** au moins une fois par an (voir section 138). Une sauvegarde
jamais testée n'est pas une sauvegarde.

## 42. Cérémonie de génération de clés (key ceremony)

La cérémonie est la procédure **formelle, documentée et auditée** de génération d'une clé
critique (typiquement la racine d'AC). Déroulé type :

