---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-29
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "agent", "cyber", "diffusion", "incident", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [3768, 3916]
sha256: c46b751bb3a0358f3ec3de8674cd43cf194307be463929a45942270cdd1f696e
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

- **Prétexting** : « je suis du support Thales, j'ai besoin de votre carte pour une
  mise à jour urgente ». Contre-mesure : procédure écrite — **aucune opération sur
  convocation non planifiée**, vérification par un second canal.
- **Corruption** : acheter un custodian. Contre-mesure : le quorum (il faut corrompre
  M personnes), rotation des custodians, séparation avec l'exploitation.
- **Perte/vol de carte** : carte égarée = part potentiellement compromise.
  Contre-mesure : déclaration immédiate, révocation de la part, régénération du quorum.

La sécurité du MofN est **organisationnelle** avant d'être technique — d'où la
formation (section 196) et l'exercice de crise (section 190).

## 233. Tableau de synthèse : menaces → contre-mesures → sections

| Menace | Contre-mesure principale | Sections |
|---|---|---|
| Vol/copie de clé | HSM, clés non extractibles | 1-4, 39 |
| Perte de clé | Sauvegarde chiffrée + MofN, tests | 41-43, 138 |
| Rootkit UEFI | Secure Boot + measured + attestation | 54-57 |
| BMC compromis | Durcissement, VLAN, patch | 59-60 |
| Interception réseau crypto | VLAN dédié, TLS mutuel | 35, 210 |
| Evil maid / vol disque | TPM seal, chiffrement disque | 53, 170 |
| DMA / cold boot | IOMMU, chiffrement mémoire | 171, 154 |
| Supply chain | Circuit officiel, SPDM, quarantaine | 63-66 |
| Harvest now/decrypt later | TLS hybride, PQC, crypto-agilité | 147-153 |
| Ransomware | Immuabilité, air gap | 208-209 |
| Insider | Séparation des rôles, journaux, MofN | 44, 110, 120 |
| Coupure électrique | Onduleur secouru, séquence de redémarrage | 140-141, 145 |
| Surchauffe | Monitoring, airflow, clim N+1 | 142, 144 |
| Ingénierie sociale | Procédures, formation, quorum | 190, 196, 232 |
| Attaque physique TEE (DDRop) | Défense en profondeur, contrôle d'accès | 91-99, 155 |

> Ce tableau est la **carte de visite** du guide : chaque menace a sa contre-mesure,
> chaque contre-mesure sa section détaillée. L'utiliser en revue de direction pour
> montrer la couverture — et les trous à combler en priorité.

# PARTIE 19 — POUR ALLER PLUS LOIN : LABO, RESSOURCES, COMMUNAUTÉS

## 234. Monter un labo crypto à moins de 2 000 €

Le meilleur moyen d'apprendre : manipuler. Configuration labo recommandée :

| Poste | Choix | Coût indicatif |
|---|---|---|
| HSM | 1× YubiHSM 2 | ~650 € |
| Serveur | Mini-PC x86 d'occasion (TPM 2.0 vérifié) | ~300 € |
| OS | Debian/Ubuntu LTS | 0 € |
| PKI | EJBCA (container) | 0 € |
| KMS | Vault OSS (3 nœuds en VM sur le mini-PC) | 0 € |
| TSA | OpenTSA | 0 € |
| Attestation | Keylime (vérificateur + 1 agent) | 0 € |
| **Total** | | **~1 000 €** |

Parcours : générer une AC racine sur le YubiHSM (cérémonie simulée avec 2 collègues),
émettre des certificats, les révoquer, tester le backup/restauration, casser
volontairement (restaurer sans la sauvegarde — constater la perte), puis tout
documenter. Ce labo vaut mieux qu'une formation théorique.

## 235. Outils open source à maîtriser

- **EJBCA** — PKI complète (émission, OCSP, CMP/SCEP). Le standard open source.
- **HashiCorp Vault** — secrets, PKI, Transit encryption, auto-unseal.
- **Keylime** — attestation TPM à distance (CNCF).
- **OpenTSA / openssl ts** — horodatage RFC 3161.
- **SoftHSM** — HSM logiciel pour les tests (jamais en production pour des clés
  critiques — pas de tamper resistance).
- **TPM2-Tools / tpm2-pkcs11** — manipuler le TPM sous Linux.
- **YubiHSM SDK** (yubihsm-shell, libyubihsm) — C/Python, open source.
- **Wireshark + pkcs11** — comprendre les échanges (en labo uniquement).
- **Caliptra / OpenTitan** — designs open source de RoT (pour les curieux du silicium ;
  maturité variable — à vérifier).

## 236. Livres et références de fond

- **NIST SP 800-57** (parties 1-3) — la référence sur la gestion des clés (gratuit,
  csrc.nist.gov).
- **NIST SP 800-88** — sanitization des médias (gratuit).
- **NIST SP 800-147/147B/193** — protection des firmwares BIOS/plateforme (gratuit).
- **ANSSI** — guides d'hygiène, recommandations sur l'administration sécurisée
  (gratuit, cyber.gouv.fr).
- **ENISA** — rapports sur la transition PQC (gratuit).
- « Bulletproof SSL and TLS » (Ivan Ristić) — pour la partie TLS/PKI opérationnelle.
- Documentation constructeurs : les **product briefs** et **guides d'administration**
  (Thales, Utimaco, Yubico) — gratuits, indispensables avant achat.

## 237. Formations et certifications utiles

- **Formations constructeurs** : Thales/Utimaco (administration HSM, 3-5 jours) —
  inclure dans le devis d'achat.
- **CISSP / CCSP** — culture générale sécurité (pas spécifiques crypto, mais
  reconnues).
- **ISO 27001 Lead Implementer/Auditor** — pour porter la conformité (partie 6).
- **Formations ANSSI / CyberEdu** — selon disponibilité (à vérifier).
- **En interne** : le parcours labo (section 234) + l'exercice de crise annuel
  (section 190) valent autant qu'une certification pour l'opérationnel.

## 238. Communautés et veille

- **Listes de diffusion** : Keylime, EJBCA, Vault, OpenTitan — annonces et failles.
- **pqca (Post-Quantum Cryptography Alliance)** — suivi de la migration PQC, dont le
  tracking HSM cité dans ce guide (github.com/pqca).
- **Confidential Computing Consortium** (Linux Foundation) — TEE, attestation.
- **ANSSI / CERT-FR** — alertes et bulletins (abonnement gratuit).
- **Meetups locaux** (clubs SSI, OWASP chapters) — les retours terrain s'y échangent
  mieux que dans les webinars vendeurs.
- **Rituel d'équipe** : 1 h/mois de veille partagée (section 197) — chacun apporte une
  nouveauté, on décide quoi en faire.

## 239. Maquette POC : le protocole en 30 jours

Pour valider un HSM avant achat (à exiger du vendeur, section 164) :

**Semaine 1 — Prise en main** : installation, initialisation, création d'une partition
de test, génération d'une clé, signature test via PKCS#11.
**Semaine 2 — Intégration** : brancher **votre** application réelle (EJBCA, AD CS,
nginx, Vault) — c'est là que les incompatibilités sortent.
**Semaine 3 — Charge** : test de débit (atteint-on le besoin dimensionné ?), test de
bascule HA (débrancher un nœud), test de sauvegarde/restauration.
**Semaine 4 — Décision** : grille de notation (section 193), TCO affiné, GO/NOGO
documenté.

Règle : un POC qui « se passe bien » sans avoir testé la panne n'a rien prouvé. Casser
volontairement fait partie du protocole.

## 240. Mot de la fin : la sécurité est un sport d'équipe

Ce guide fait 240 sections, des dizaines de tableaux et de checklists — mais la
sécurité d'un datacenter ne tient pas dans un document. Elle tient dans :

- **2-3 personnes compétentes** qui se relaient (jamais une seule),
- des **procédures écrites** qu'on suit même un dimanche à 3 h du matin,
- des **tests réguliers** qui prouvent que ça marche (pas qui le supposent),
- une **direction** qui a compris que la sécurité coûte moins cher que l'incident,
- et l'**humilité** de relire ce guide dans un an — parce que les menaces, les
  produits et les standards auront bougé.

Zelef, chef de service systèmes & énergies : vous avez déjà le réflexe le plus rare —
penser **énergie et sécurité ensemble**. Un HSM sans courant, un FPGA sans airflow,
une PKI sans onduleur : la sécurité cryptographique est un système **électromécanique**
autant que mathématique. C'est cette vision qui fait la différence entre une crypto
« sur le papier » et une crypto qui tient la coupure de courant, la canicule et le
lundi matin.

> *Bon courage — et que vos clés restent à l'abri, vos firmwares signés, et vos
> onduleurs chargés.*

# FICHES RÉFLEXES — À IMPRIMER ET AFFICHER

## Fiche R1 — HSM en panne (à afficher près de la baie)

