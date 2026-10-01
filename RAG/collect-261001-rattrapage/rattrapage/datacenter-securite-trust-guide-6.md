---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-6
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["datacenter", "arr", "attention", "aws", "capex", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [603, 747]
sha256: 1dde31cc068003ca38fe7b4f7d3e4d02c1d47bc3bdb3fcc7b0b0d3db22bc55a9
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

1. **Préparation** : script écrit et relu, salle sécurisée, HSM initialisé et vérifié
   (scellés, firmware), participants convoqués (voir quorum).
2. **Participants** : au minimum 2-3 personnes de confiance + un **témoin/auditeur**
   (interne ou externe). Chacun a un rôle écrit.
3. **Génération** : la clé est générée **dans** le HSM (TRNG interne), paramètres
   vérifiés à l'écran (algorithme, taille — ex. RSA-4096 ou ECDSA P-384 pour une racine).
4. **Sauvegarde** : duplication vers le HSM de backup, ou partage MofN (section 43).
5. **Procès-verbal** : qui, quand, où, numéros de série, empreintes de clés publiques,
   signatures des participants. Ce PV est un **document d'audit** (eIDAS, WebTrust).

Durée : 1 à 3 heures pour une racine. Coût : essentiellement du temps humain — mais c'est
le moment le plus critique de la vie de la PKI. **Ne jamais générer une racine « vite fait »
sur un poste admin.**

## 43. Quorum M of N : le principe

Le **quorum MofN** (ex. 3-of-5) : la clé (ou le secret permettant de la restaurer) est
découpée en N parts (smartcards, PED keys), et **M parts** sont nécessaires pour la
reconstituer. Personne seul ne peut restaurer la clé ; il faut la collusion de M personnes.
En pratique :

- Les parts sont remises à des **custodians** nommés (direction, RSSI, admin senior...),
  stockées en **lieux séparés** (coffres différents, sites différents).
- M est typiquement la **majorité** (2-of-3, 3-of-5) : compromis entre sécurité et
  disponibilité (si un custodian est injoignable).
- Chez Thales, le **PED** (PIN Entry Device) gère les clés de partition sur smartcards ;
  chez Utimaco, le **« n out of m »** est natif avec smartcards d'opérateurs.

Piège classique : mettre les N parts dans le **même coffre** « pour simplifier ». Le quorum
ne vaut que par la séparation physique et organisationnelle.

## 44. Custodians, PED, smartcards : qui porte quoi

| Rôle / objet | Fonction | Bonne pratique |
|---|---|---|
| **Crypto Officer (CO)** | Administre une partition au quotidien | 2 CO minimum, jamais le même que l'auditeur |
| **Custodian** | Détient une part du quorum MofN | Personne de confiance, hors équipe d'exploitation si possible |
| **PED / smartcard** | Support physique des parts et de l'authentification forte | Cartes numérotées, inventaire annuel |
| **Auditeur/témoin** | Assiste aux cérémonies, vérifie les PV | Indépendant de l'exploitation |
| **SO (Security Officer)** | Initialise le HSM (niveau boîtier) | Utilisé uniquement à l'initialisation |

Règle d'or : **séparation des tâches** — celui qui génère ne sauvegarde pas seul, celui qui
exploite ne détient pas le quorum seul.

## 45. HA, clustering, reprise après sinistre

Un HSM est un **point critique** : s'il tombe, les signatures/chiffrements s'arrêtent. Les
mécanismes :

- **Cluster HA** : 2+ HSM réseau en **réplication synchrone** (ex. Luna HA : un groupe,
  un membre actif, les autres en miroir ; bascule automatique côté client). Prévoir 2
  boîtiers minimum en production.
- **Client multi-HSM** : le client PKCS#11 connaît les 2 membres et bascule
  automatiquement.
- **Site distant** : un 3e HSM (ou le backup) sur le site de secours pour le PRA,
  réplication asynchrone ou manuelle.
- **RPO/RTO** : pour des clés (petit volume), RPO ≈ 0 envisageable en synchrone ;
  RTO = temps de bascule client (secondes à minutes). À tester (section 138).

Attention : la HA protège de la **panne**, pas de la **compromission logique** (une clé
compromise est répliquée compromise). D'où la sauvegarde offline déconnectée.

## 46. Durée de vie : batterie, MTBF, fin de vie

- **Batterie** : les clés en SRAM sont alimentées par une pile interne (durée typique
  5-10 ans selon modèles). Une batterie faible = alerte à traiter **avant** la panne :
  à son épuisement, le HSM s'efface (comportement de sécurité normal). **Contrôler
  l'état batterie dans la supervision** et planifier le remplacement (intervention
  constructeur, souvent avec cérémonie).
- **MTBF** : 171 308 h (Luna Network) / 997 508 h (Luna PCIe) — valeurs constructeur
  vérifiées, utiles pour comparer, pas des garanties.
- **Fin de vie** : à la mise au rebut, **zeroization** documentée + destruction physique
  si le modèle le permet (section 100). Ne jamais revendre un HSM sans effacement certifié.

## 47. Prix indicatifs : tableau (ordres de grandeur, à vérifier sur devis)

| Produit | Prix d'achat indicatif | Récurrent annuel indicatif |
|---|---|---|
| YubiHSM 2 (standard / FIPS) | 650 € / 949 € (constaté 2026) | 0 € |
| Luna PCIe HSM (A700→A790) | ~5–12 k€ | Support ~15-20 %/an |
| Luna Network HSM (T-2000→T-5000) | ~15–35 k€ | Support ~15-20 %/an |
| Utimaco Se Gen2 / u.trust (Se2k→Se15k) | ~10–30 k€ | Support ~15-20 %/an |
| AWS CloudHSM (1 HSM) | 0 € CAPEX | ~13 k$/an (~1,45 $/h) |
| AWS CloudHSM (cluster 2 HSM, 2 AZ) | 0 € CAPEX | ~26 k$/an |
| Azure Dedicated HSM | 0 € CAPEX | ~42 k$/an (4,85 $/h — tarif public vérifié) |

Tous les prix matériels sont **à vérifier** : les constructeurs vendent via partenaires,
avec des remises selon volumes et support. Le support annuel (15-20 %) inclut firmware et
remplacement — **ne pas le négliger** dans le budget.

## 48. TCO : calcul exemple sur 5 ans

**Scénario** : ETI, 1 appliance réseau + 1 backup, support 18 %/an.

| Poste | Année 1 | Années 2-5 (par an) | Total 5 ans |
|---|---|---|---|
| 2× appliance (25 k€ pièce, indicatif) | 50 000 € | — | 50 000 € |
| Support 18 % | 9 000 € | 9 000 € | 45 000 € |
| Installation + cérémonie (presta) | 8 000 € | — | 8 000 € |
| Formation (2 pers.) | 4 000 € | — | 4 000 € |
| Électricité (2× 84 W × 24h × 365 × 0,20 €/kWh) | ~300 € | ~300 € | ~1 500 € |
| **Total** | **~71 300 €** | **~9 300 €** | **~108 500 €** |

Comparaison : le même service en **AWS CloudHSM** (2 HSM) ≈ 26 k$/an ≈ **130 k$ sur 5 ans**,
sans CAPEX ni gestion matérielle. Le **seuil de rentabilité** de l'achat se situe vers
**3-4 ans** d'utilisation continue — en-deçà, le cloud HSM gagne ; au-delà, l'appliance
gagne. Pour un usage < 2 ans ou très variable : cloud HSM sans hésiter.

# PARTIE 2 — ROOT OF TRUST : LA CONFIANCE DEPUIS LE SILICIUM

## 49. Root of Trust : définition et pourquoi c'est la base

Le **Root of Trust (RoT)** est le composant (matériel, firmware ou les deux) en qui le système
a **implicitement confiance** pour démarrer la chaîne de vérification. Tout le reste —
bootloader, OS, applications — est vérifié **à partir de lui**. S'il est compromis, tout
l'édifice s'effondre : un rootkit dans le firmware survit à la réinstallation de l'OS, au
changement de disque, à tout. En datacenter, le RoT concerne **chaque serveur** : des centaines
de machines avec BMC, BIOS/UEFI, cartes réseau à firmware — autant de surfaces d'attaque.

On distingue classiquement :
- **RTM (Root of Trust for Measurement)** : mesure chaque étape du démarrage (hash dans le TPM).
- **RTS (Root of Trust for Storage)** : protège les clés/secrets (TPM, HSM).
- **RTR (Root of Trust for Reporting)** : signe les attestations (clé EK/AK du TPM).

## 50. Chaîne de confiance : du silicium à l'OS

```
  [Silicium / Boot ROM immuable]
        │ vérifie (signature)
        ▼
  [Bootloader / firmware initial] ──mesure──▶ [TPM PCR]
        │ vérifie
        ▼
  [UEFI / BIOS] ──mesure──▶ [TPM PCR]
        │ vérifie
        ▼
  [Bootloader OS (shim, GRUB)] ──mesure──▶ [TPM PCR]
        │ vérifie
        ▼
  [Noyau OS] ──mesure──▶ [TPM PCR]
        │
        ▼
  [Applications / conteneurs]
```

