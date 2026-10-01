---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-8
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["datacenter", "agent", "gpu", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [881, 1013]
sha256: d7cf6373ae7322ebd2c9a6d314508a149536d386bcebbb08692050e4324e586c
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

- [ ] **Mot de passe** : changer le défaut à la réception (chaque BMC un mot de passe
      unique, stocké au coffre-fort de mots de passe). Désactiver les comptes inutiles.
- [ ] **Réseau dédié** : le BMC sur un **VLAN management** séparé, jamais sur le VLAN
      de production, jamais exposé sur Internet. Firewall : seuls les postes d'admin.
- [ ] **Firmware** : mettre à jour à la réception puis suivre les bulletins (les failles
      iLO/iDRAC sont régulières et critiques).
- [ ] **Services** : désactiver Telnet/HTTP non chiffré, IPMI 1.5 (authentification
      faible), ne garder que HTTPS/Redfish + SSH si besoin.
- [ ] **IPMI** : si utilisé, imposer **IPMI 2.0 + chiffrement**, mots de passe longs ;
      le hash RAKP d'IPMI se craque offline — un BMC sur réseau non fiable = compromis.
- [ ] **Certificats** : remplacer le certificat auto-signé par un certificat de votre PKI
      interne (qui vit idéalement sur HSM — boucle bouclée).
- [ ] **Journaux** : syslog du BMC vers le SIEM ; alerter sur les connexions anormales.
- [ ] **2FA** : activer l'authentification multifacteur sur l'interface web quand
      disponible.

## 61. OpenBMC, Redfish, DMTF : les standards

- **Redfish** (DMTF) : API REST standard de gestion des serveurs — remplace progressivement
  IPMI. Exigez-le sur les nouveaux serveurs : requêtes HTTPS + JSON, inventaire, mise à
  jour firmware, gestion des comptes.
- **OpenBMC** : stack BMC open source (Linux Foundation) adoptée par les hyperscalers et
  des ODM (Supermicro, etc.). Avantage : **auditabilité** du code du BMC, mises à jour
  maîtrisées. Si vous achetez des serveurs ODM, demandez OpenBMC.
- **DMTF SPDM** (Security Protocol and Data Model) : attestation des **composants**
  (cartes réseau, disques, GPU) — le serveur vérifie l'identité et le firmware de ses
  périphériques. À exiger sur les achats 2026+ pour les composants critiques.

## 62. Mises à jour firmware signées

Tout firmware (UEFI/BMC, cartes réseau, disques, FPGA, alimentations !) doit être **signé
par le vendeur** et **vérifié avant application**. Procédure :

- Télécharger uniquement depuis le site officiel du constructeur (vérifier le hash/TLS).
- Ne jamais appliquer un firmware « trouvé sur un forum ».
- Maintenir un **registre des versions** par serveur (inventaire automatisé via Redfish).
- Tester sur un serveur pilote avant déploiement massif.
- Après mise à jour : **re-mesurer** (les PCR changent → mettre à jour les valeurs de
  référence d'attestation, sinon tout le parc passe en alerte).

## 63. Supply chain : pièces contrefaites, le risque réel

Le risque supply chain n'est pas théorique : disques, barrettes mémoire, transceivers,
alimentations et même cartes mères **contrefaites ou reconditionnées vendues comme neuves**
circulent sur le marché gris. Conséquences : firmwares modifiés, portes dérobées, fiabilité
dégradée, et **invalidation des certifications**. Les vecteurs : achats hors circuit
officiel pour « faire des économies », pièces de rechange d'occasion, sous-traitants peu
regardants.

## 64. Supply chain : SBOM, attestations constructeur

Contre-mesures 2026 :

- **Acheter en circuit officiel** (constructeur ou distributeur agréé) pour tout composant
  critique ; conserver factures et numéros de série.
- **SBOM** (Software Bill of Materials) : exiger du vendeur la liste des composants
  logiciels/firmwares — c'est une exigence croissante (EO 14028 aux US, CRA en UE).
- **Attestation SPDM** (section 61) : vérification cryptographique des composants à
  l'allumage.
- **Scellés et traçabilité** : à la réception, vérifier l'intégrité des emballages et la
  concordance des numéros de série (voir section 65).

## 65. Réception serveur : contrôle à la livraison

- [ ] Emballage intact, scellés constructeur présents.
- [ ] Numéros de série (châssis, carte mère, disques, alimentations) = bon de livraison.
- [ ] Premier allumage **en zone de quarantaine** (pas directement en production).
- [ ] Mise à jour UEFI/BMC vers les versions validées, Secure Boot activé, TPM provisionné.
- [ ] Mots de passe BMC changés, inventaire Redfish initial capturé (référence).
- [ ] Photo des étiquettes et archivage du PV de réception.

## 66. Serveurs d'occasion/reconditionnés : précautions

Le reconditionné fait sens économiquement, mais impose un **reconditionnement de confiance** :

- Effacement crypto des disques (section 101) ou remplacement pur et simple.
- **Reflash complet** des firmwares (UEFI, BMC) depuis les images officielles.
- Réinitialisation d'usine du TPM (`tpm2_clear`) et du BMC.
- Vérification SPDM/attestation des composants si disponible.
- Ne jamais acheter de **HSM d'occasion sans effacement certifié** par un professionnel
  (voir le marché du reconditionné — exiger le PV de reset et tester avant production).

## 67. Cas terrain : BMC compromis (retour d'expérience type)

Scénario vécu dans l'industrie (synthèse de cas publics) : un parc de serveurs avec BMC en
**DHCP sur le VLAN de production**, mots de passe **par défaut** (constructeur connus),
firmware vieux de 3 ans. Un attaquant ayant un pied sur le réseau scanne le port 623 (IPMI),
se connecte en admin, monte une **image ISO** via le KVM virtuel et redémarre les serveurs
dessus — prise de contrôle totale, **persistante** (le BMC survit à la réinstallation),
**invisible** (aucun agent OS ne voit le BMC). Coût de remédiation : reflash de tout le
parc, changement de tous les mots de passe, segmentation réseau — plusieurs semaines.
Leçon : le BMC est un **serveur à part entière** avec son cycle de patch, pas un « détail
hardware ».

## 68. Tableau : mécanismes RoT par couche

| Couche | Mécanisme | Ce qu'il garantit | Outil de vérification |
|---|---|---|---|
| Silicium | Boot Guard / PSB (à vérifier) | Le 1er firmware est signé | Fuses usine, doc OEM |
| Firmware UEFI | Secure Boot | Seuls les binaires signés bootent | Setup UEFI, Redfish |
| Firmware UEFI | Measured Boot | Le boot est enregistré (PCR 0-7) | `tpm2_pcrread`, Keylime |
| OS | Secure/measured boot (shim/GRUB) | Noyau signé et mesuré (PCR 8-15) | Attestation distante |
| BMC | Firmware signé, mots de passe | Le BMC n'est pas compromis | Redfish, audit firmware |
| Composants | SPDM | Cartes/disques authentiques | Attestation SPDM |
| Clés | TPM seal, HSM | Secrets liés à l'état mesuré | Politiques TPM |

## 69. Ce que le TPM ne protège pas (limites honnêtes)

- Un TPM ne protège pas contre un **firmware signé mais vulnérable** (il mesure, il ne
  juge pas la qualité du code).
- Il ne protège pas la **mémoire vive** en cours d'exécution (attaques DMA, cold boot —
  d'où le chiffrement mémoire des TEE, voir partie 9).
- Un **fTPM** partage le silicium avec le reste : des vulnérabilités ont existé, patcher
  est obligatoire.
- Sans **infrastructure d'attestation**, les mesures restent des nombres que personne ne
  vérifie.
- Le TPM ne remplace pas le **HSM** : il protège les secrets d'**une** machine, pas les
  clés d'une organisation.

## 70. Lien avec le HSM : sceller des clés au TPM

Architecture type : la **clé maîtresse** vit dans le HSM (datacenter) ; chaque serveur en
dérive une clé locale **scellée à ses PCR** via son TPM. Au boot, le serveur demande au
HSM/KMS sa clé enveloppe, le TPM la descelle uniquement si les PCR sont conformes, et la
clé n'existe en clair qu'en RAM. Compromis élégant : révocation centrale possible (via le
HSM), vol du disque inutile (clé scellée), et pas de clé maîtresse qui traîne sur les
serveurs. C'est le pattern derrière le **sealed storage** des OS modernes et l'auto-unseal
de Vault (section 113).

# PARTIE 3 — FPGA : LE MATÉRIEL REPROGRAMMABLE

## 71. FPGA : le principe en 5 minutes

