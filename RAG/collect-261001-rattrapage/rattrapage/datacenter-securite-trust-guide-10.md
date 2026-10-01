---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-10
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["datacenter", "amd", "gpu", "intel", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1145, 1260]
sha256: aa0c3c3e5524ab1f7285bcbb79622d7787f2ae3addbc62bb4b415351d351d673
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Le cas d'usage datacenter n°1 du FPGA : **décharger le CPU**. Jusqu'à ~80 % des cycles CPU
peuvent partir en overhead I/O sur des workloads cloud-natifs (chiffre avancé par Intel
pour justifier ses IPU). Le FPGA/SmartNIC prend en charge : switching, routage, firewall,
**IPsec/MACsec** (moteurs 400G câblés sur l'Alveo V80), encapsulation (VxLAN, GENEVE),
équilibrage de charge. Résultat : les cycles CPU retournent aux applications — c'est un
**gain capacitaire** mesurable en nombre de serveurs évités. Pour Zelef : un parc qui
chiffre beaucoup (IPsec site-à-site massif) est un candidat naturel.

## 82. Cas d'usage 2 — Crypto : TLS, IPsec/MACsec, PQC

- **TLS** : terminaison à haut débit avec clés protégées (le FPGA peut dialoguer avec un
  HSM pour les opérations de clé privée, ou embarquer ses propres clés dans du stockage
  sécurisé).
- **IPsec/MACsec** : les moteurs câblés (3× 400G sur V80) traitent le chiffrement
  ligne-rate sans charger le CPU.
- **PQC** : les algorithmes post-quantiques (ML-KEM, ML-DSA) sont **plus lourds** que
  RSA/ECDH (clés et signatures plus grosses, plus de calculs) — le FPGA est un candidat
  naturel pour les accélérer, et sa **reprogrammabilité** est précieuse pendant la
  transition (les algos peuvent encore évoluer — cf. affaire HAWK, section 153).

## 83. Cas d'usage 3 — Trading haute fréquence

En finance, la latence se mesure en **nanosecondes** et se monnaye. Les FPGA traitent les
flux de marché (décodage des protocoles de bourse, prise de décision, émission d'ordres)
**entièrement en matériel**, sans traverser un OS. C'est le domaine historique du FPGA,
avec des designs propriétaires ultra-optimisés. Hors sujet pour la plupart des
datacenters d'entreprise, mais c'est la démonstration la plus pure de l'avantage FPGA :
**déterminisme absolu**.

## 84. Cas d'usage 4 — IA et inférence

Le FPGA accélère l'**inférence** (pas l'entraînement) sur des modèles **figés et
quantifiés** (INT8) : latence basse et déterministe, excellent perf/watt. Outils : Vitis AI
(AMD), FPGA AI Suite + OpenVINO (Intel/Altera). Limites honnêtes : l'écosystème logiciel
est **très en retard** sur le GPU (CUDA), les gros LLM ne tiennent pas sur FPGA, et chaque
nouveau modèle demande un nouveau design. Pertinent pour : vision industrielle, NLP
embarqué, scoring temps réel à latence garantie. Pour les LLM : restez sur GPU (voir le
guide IA de Zelef).

## 85. Cas d'usage 5 — Stockage et compression

Compression/décompression (GZIP, ZSTD), déduplication, chiffrement au vol, calcul
d'empreintes — le FPGA excelle dans ces pipelines réguliers à haut débit. Cas datacenter :
**appliances de sauvegarde** (accélération de la déduplication), **stockage NVMe-oF**
(offload du protocole), **bases de données** (filtrage/compression près des données).
Souvent le ROI le plus rapide : le FPGA se paie en **débit disque/licences évités**.

## 86. Consommation et refroidissement : les chiffres

| Équipement | TDP / conso typique | Refroidissement | Remarque |
|---|---|---|---|
| Alveo V80 | **190 W TDP** (vérifié) | Passif (flux d'air serveur requis) | Prévoir le slot double + airflow |
| Alveo U55C (gén. précédente) | ~150-225 W (à vérifier) | Passif/actif selon SKU | — |
| SmartNIC FPGA type N6000 | ~75-150 W (à vérifier) | Passif | Vérifier le SKU exact |
| GPU H100 (comparaison) | 700 W | Actif/liquide | 3-4× un FPGA |
| Serveur 2U standard | 300-800 W | Ventilateurs | — |

Point énergie (lien avec le guide onduleurs de Zelef) : un serveur avec 2 FPGA = **+300 à
400 W** à prévoir dans le dimensionnement de la baie, de l'onduleur et de la clim. Le FPGA
ne « consomme peu » qu'en **perf/watt** sur son workload — en absolu, c'est une charge
thermique sérieuse à évacuer (airflow avant-arrière, pas d'obstruction).

## 87. Quand ça vaut le coup : grille de décision

| Question | Si oui → | Si non → |
|---|---|---|
| Le workload est-il **stable** (pas de changement mensuel) ? | FPGA envisageable | Restez CPU/GPU |
| La **latence déterministe** est-elle critique ? | FPGA fort candidat | GPU/CPU suffisent |
| Le débit dépasse-t-il ce qu'un CPU fait (~10-40 Gb/s chiffrés/cœur) ? | FPGA | CPU |
| Avez-vous (ou pouvez-vous acheter) de l'**expertise RTL/HLS** ? | FPGA | GPU (écosystème logiciel) |
| Le volume justifie-t-il 10 k€/carte + développement ? | FPGA | Cloud / CPU |
| Le workload est-il de l'IA générative / LLM ? | — | **GPU**, pas FPGA |

Score : 4+ « oui » → étudier le FPGA sérieusement ; 2-3 → prototype ; moins → oubliez.

## 88. Coûts : prix et TCO d'une carte FPGA

- **Carte** : ~9 500 $ (Alveo V80 MSRP vérifié) ; 3 000-10 000 € selon gamme et
  distributeur (**à vérifier** au jour de l'achat).
- **Serveur hôte** : prévoir un serveur avec slots PCIe x16 libres, airflow adapté,
  alimentation redondante dimensionnée (+190 W/carte).
- **Licences outils** : Vivado/Quartus — les éditions pour ces cartes sont souvent
  incluses ou en abonnement (**à vérifier**).
- **Développement** : le vrai coût — **3 à 12 mois-ingénieur** pour un design sérieux,
  en interne ou via un bureau d'études (50-150 k€). C'est ce poste qui fait ou défait le ROI.
- **TCO 3 ans** (1 carte + serveur + dev. amorti) : de l'ordre de **80-200 k€** — à
  comparer au nombre de serveurs CPU évités ou à la valeur de la latence gagnée.

## 89. Outils : Vivado, Quartus, open source

- **AMD Vivado / Vitis** : flot historique Xilinx, supporte les Alveo (exemple de design
  AVED sur GitHub pour le V80 — vérifié).
- **Intel Quartus Prime** : flot Altera pour Agilex.
- **Open source** : Yosys/nextpnr pour les petits FPGA (Lattice surtout) — pas pour les
  grosses cartes datacenter en production.
- **Simulation** : cocotb (Python), Verilator — indispensables avant de flasher.
- Conseil : figer les **versions d'outils** par projet (un design validé sous Vivado
  2023.2 n'est pas garanti sous 2024.1 sans re-validation).

## 90. FPGA en production : drivers, monitoring, mise à jour bitstream

- **Drivers** : XRT (AMD) / OFS (Intel) — à intégrer dans l'image OS standard, version
  épinglée.
- **Monitoring** : température du FPGA, taux d'utilisation, erreurs ECC mémoire —
  remonter au SIEM/supervision comme tout équipement.
- **Mise à jour du bitstream** : procédure **signée et versionnée** (comme un firmware) ;
  prévoir le **rollback** (garder l'ancien bitstream valide). Un bitstream corrompu =
  carte inutilisable jusqu'au reflash physique.
- **Sécurité du bitstream** : activer le **chiffrement + authentification** du bitstream
  (clé dans le FPGA) — sinon la propriété intellectuelle et l'intégrité du design sont
  exposées à quiconque lit la flash de la carte.

# PARTIE 4 — SÉCURITÉ PHYSIQUE DU DATACENTER

## 91. Défense en profondeur physique : les 4 cercles

