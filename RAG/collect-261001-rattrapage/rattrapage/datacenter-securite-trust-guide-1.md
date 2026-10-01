---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-1
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["datacenter", "amd", "asic", "gpu", "incident", "intel", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1, 125]
sha256: 4d3693caa869be91cc86a8833e0accf5de3cb423778d3fca3db01ff31dad697f
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

> Guide technique en français — rédigé le 27/09/2026 pour Zelef, chef de service systèmes & énergies.
> Angle : chiffres, BOMs, dimensionnements, lien énergie (conso, refroidissement, alimentation).
> Références produits **vérifiées le 27/09/2026** par recherche web. Prix = ordres de grandeur, **à vérifier**
> sur devis. Référence introuvable = **« non trouvée au 27/09/2026 »**. Aucune invention.
> Complète le guide BIOS/UEFI/TPM existant de Zelef (focus datacenter/serveurs, sans dupliquer).

---

## Sommaire

- **Partie 1 — HSM** (sections 1 à 48) : principe, usages, form factors, fiches produits vérifiées,
  comparatifs, placement réseau, dimensionnement, sauvegardes, cérémonie, quorum MofN, HA, prix, TCO.
- **Partie 2 — Root of Trust** (sections 49 à 70) : TPM 2.0 serveur, secure boot, measured boot,
  attestation, BMC sécurisés, supply chain.
- **Partie 3 — FPGA** (sections 71 à 90) : principe, cartes AMD Alveo / Intel Agilex vérifiées,
  comparatif FPGA vs GPU vs ASIC, cas d'usage, grille de décision.
- **Partie 4 — Sécurité physique** (sections 91 à 104) : contrôle d'accès, cages, vidéo,
  intrusion, visiteurs, destruction de médias.
- **Partie 5 — KMS / gestion des clés** (sections 105 à 120) : cycle de vie, rotation, rôles,
  HSM↔KMS, Vault, EJBCA, sauvegardes.
- **Partie 6 — Conformité** (sections 121 à 128) : ISO 27001, PCI DSS, RGPD, eIDAS (généraliste,
  pas de conseil juridique).
- **Partie 7 — Checklists pratiques** (sections 129 à 138) : mise en service serveur, BMC, HSM,
  cérémonie, runbooks incident.
- **Partie 8 — Lien énergie** (sections 139 à 146) : consos, onduleurs, refroidissement,
  continuité des services de clés.
- **Partie 9 — À venir, annonces vérifiées au 27/09/2026** (sections 147 à 157) : PQC NIST,
  confidential computing.
- **Glossaire** (40 termes), **Quiz** (10 questions + réponses), **Pièges terrain** (25).

---

# PARTIE 1 — HSM : LE COFFRE-FORT DES CLÉS

## 1. Principe : pourquoi un HSM plutôt qu'un fichier de clés

Un HSM (Hardware Security Module) est un boîtier ou une carte dédiée qui **génère, stocke et utilise
des clés cryptographiques sans jamais les laisser sortir en clair**. La clé privée d'une autorité de
certification, la clé maîtresse qui chiffre une base de données, la clé de signature d'un firmware :
si elle est dans un fichier sur disque, quiconque lit le disque (ou la sauvegarde, ou le snapshot
VM) la possède. Dans un HSM, la clé naît à l'intérieur du composant sécurisé, y reste, et seules
des **opérations** (signer, déchiffrer) sont exposées via une API. Le matériel est conçu pour
**s'auto-détruire logiquement** (effacement des clés) en cas d'ouverture ou d'attaque physique.

Règle de base : **toute clé dont la compromission coûte plus cher que le HSM doit être dans un HSM.**
Clé racine d'AC, clé de chiffrement des sauvegardes, clé de signature de code, clé DNSSEC KSK :
oui. Clé TLS d'un serveur web interne renouvelée tous les 90 jours : discutable, un KMS logiciel
suffit souvent.

## 2. Génération de clés : TRNG et entropie matérielle

Un HSM embarque un **générateur d'aléa vrai (TRNG)** basé sur un phénomène physique (bruit
thermique, oscillateurs en anneau). Les clés générées dans le HSM ne dépendent donc pas de
l'entropie de l'OS hôte — point critique sur les VM cloud au démarrage, historiquement pauvres en
entropie. Les HSM sérieux combinent une source de bruit physique et un DRBG conforme
**NIST SP 800-90A** (ex. CTR-DRBG). Exemple vérifié le 27/09/2026 : les Thales Luna annoncent un
générateur « conçu pour être conforme AIS 20/31 (DRG.4) avec source de bruit matérielle + CTR-DRBG
NIST 800-90A » (product brief Luna Network HSM, cpl.thalesgroup.com).

En pratique : générer la clé **dans** le HSM (et non l'y importer) est la méthode recommandée pour
les clés racines. L'import existe (key wrapping, enveloppe chiffrée) pour la migration, mais chaque
import est une fenêtre de risque.

## 3. Tamper resistance : ce que le boîtier fait quand on l'attaque

La résistance à l'effraction (tamper resistance) repose sur plusieurs couches :

- **Capteurs** : interrupteurs d'ouverture du capot, capteurs de température (attaque au froid pour
  figer la RAM), capteurs de tension/fréquence (glitch), parfois capteurs de rayonnement.
- **Mémoire volatile protégée par batterie** : les clés vivent en SRAM alimentée en permanence.
  Coupure d'alimentation anormale ou déclenchement d'un capteur → **effacement immédiat**
  (zeroization).
- **Revêtement** : résine époxy ou mesh conducteur autour du composant : toute tentative de
  perçage/attaque par sonde coupe des pistes et déclenche l'effacement.
- **Blindage** : contre les attaques par canaux auxiliaires (analyse de consommation, émissions
  électromagnétiques).

C'est exactement ce qui justifie la certification **FIPS 140-2/140-3 niveau 3 ou 4** (voir
section 5) : le niveau 3 exige une réponse active à l'effraction.

## 4. Tamper evidence : scellés et traçabilité

Complément de la résistance : la **preuve** d'effraction. Les HSM réseau sont livrés avec des
scellés inviolables sur les vis et les jointures du châssis. À la réception, on photographie et on
consigne l'état des scellés dans le procès-verbal d'installation. Un scellé dégradé par la chaleur
ou le temps n'est pas forcément une compromission, mais c'est un **drapeau rouge** à documenter
avec le vendeur. Les journaux d'audit internes du HSM (tamper-evident logs, parfois chaînés par
hash comme sur le YubiHSM 2) complètent le dispositif : toute tentative laisse une trace.

## 5. FIPS 140-2 / 140-3 : les 4 niveaux, ce qu'ils garantissent vraiment

FIPS 140 est la certification de référence (NIST, reconnue internationalement). Ce qu'il faut en
retenir, sans jargon inutile :

| Niveau | Exigence clé | Usage typique |
|---|---|---|
| 1 | Algorithmes validés, pas d'exigence physique particulière | Logiciel chiffrant (ex. module validé dans un OS) |
| 2 | Scellés inviolables, contrôle d'accès par rôles | Clés USB chiffrées, tokens |
| 3 | **Détection + réponse active à l'effraction**, effacement des clés, authentification forte des opérateurs | **HSM réseau/PCIe** (c'est le niveau visé par Luna, Utimaco, nShield) |
| 4 | Protection contre les attaques physiques poussées (analyse environnementale large, canaux auxiliaires) | HSM haut de gamme, militaire |

Points honnêtes :

- FIPS 140-2 est en **fin de vie** : les nouvelles validations se font en **FIPS 140-3**
  (ISO/IEC 19790:2021). Un HSM « FIPS 140-2 niveau 3 » reste valable, mais vérifiez la date de la
  validation sur le site du NIST (CMVP).
- FIPS valide une **version précise** de firmware : mettre à jour le firmware peut exiger une
  re-validation ou passer par un mode « non validé » temporaire. À planifier.
- La validation porte sur le **module cryptographique**, pas sur tout le produit ni sur votre
  déploiement.

## 6. Common Criteria EAL : le complément européen

Common Criteria (ISO/IEC 15408) évalue le produit contre un **profil de protection** avec des
niveaux d'assurance EAL1 à EAL7. Pour les HSM, on croise typiquement **EAL4+**. Vérifié le
27/09/2026 : Utimaco annonce ses HSM généralistes « FIPS 140-2 niveau 3 (niveau 4 en cours) »
et une certification **Common Criteria** de son HSM par le CCN espagnol (communiqué Utimaco,
septembre 2026, hsm.utimaco.com). En pratique PME/datacenter : FIPS 140-2/3 L3 suffit dans 95 %
des appels d'offres ; Common Criteria devient exigé sur les marchés publics sensibles et
l'administration (ex. exigences ANSSI/CPM en France — à vérifier au cas par cas).

## 7. PCI HSM, eIDAS : certifications sectorielles

