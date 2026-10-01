---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-7
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["datacenter", "agent", "amd", "distribution", "intel", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [748, 880]
sha256: 6116db0385eb0dff4423c80e79f21f7bca9564cf80f2ff8bb422eac80abb2187
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

Deux opérations à chaque étape : **vérifier** (la signature est-elle valide ? → secure boot,
qui **bloque** si non) et **mesurer** (quel est le hash de ce qui a démarré ? → measured boot,
qui **enregistre** sans bloquer). La vérification sans mesure = on sait que c'est signé, pas
**quoi exactement** a démarré. La mesure sans vérification = on enregistre tout, y compris le
malware. Il faut les deux.

## 51. Complément au guide BIOS/UEFI/TPM de Zelef : ce qu'on ne redit pas

Zelef dispose déjà d'un guide **BIOS/UEFI/TPM** (bases : activer le TPM, secure boot sur
poste, BitLocker, etc.). Cette partie **ne duplique pas** : elle couvre ce qui est
**spécifique datacenter/serveurs** — TPM sur serveurs (dTPM/fTPM/DICE), measured boot et
attestation à distance en production, BMC sécurisés (le point aveugle n°1), supply chain et
pièces contrefaites. Pour les fondamentaux (qu'est-ce qu'un PCR, comment activer le TPM dans
l'UEFI), se reporter au guide existant.

## 52. TPM 2.0 sur serveur : dTPM, fTPM, DICE

Sur serveur, trois implémentations :

| Type | Principe | Avantages | Limites |
|---|---|---|---|
| **dTPM** (discret) | Puce dédiée (ex. Infineon OPTIGA) sur la carte mère | Isolation physique maximale, certifié Common Criteria | Coût, une puce par serveur |
| **fTPM** (firmware) | TPM émulé dans le firmware (AMD fTPM, Intel PTT) | Coût nul, déjà présent | Partage le TEE du CPU ; vulnérabilités historiques (ex. failles AMD fTPM — patcher) |
| **DICE** | Identité dérivée à chaque boot (chaîne de certificats) | Idéal pour BMC/périphériques sans TPM complet | Écosystème plus jeune |

En datacenter : exiger le **dTPM 2.0** sur les serveurs critiques (PKI, HSM-adjacents,
hyperviseurs), accepter le fTPM sur le parc standard **à condition de maintenir les
microcodes/firmwares à jour**. Vérifier la présence : `tpm2_pcrread` sous Linux, ou via
l'inventaire Redfish du BMC.

## 53. PCR : les registres de mesure, en pratique

Les **PCR** (Platform Configuration Registers) sont 24 registres du TPM où s'accumulent les
mesures (hash étendu : PCR_new = hash(PCR_old || mesure)). Répartition standard :

- **PCR 0-7** : firmware/UEFI (code, configuration, secure boot).
- **PCR 8-15** : OS (bootloader, noyau, initrd).
- **PCR 16-23** : usage libre (applications, conteneurs).

Usage datacenter n°1 : **sceller un secret au PCR** (seal) — ex. la clé de déchiffrement du
disque (LUKS) n'est utilisable que si les PCR 0-7 correspondent au firmware attendu. Si le
firmware est modifié (rootkit UEFI), les PCR changent, le secret reste scellé. Usage n°2 :
**politiques d'accès** — n'autoriser le déverrouillage que sur un état mesuré connu.

## 54. Secure Boot UEFI : focus serveurs datacenter

Sur serveur, le Secure Boot UEFI vérifie la signature de chaque binaire de boot (shim →
GRUB → noyau). Spécificités datacenter :

- **Clés** : utiliser le **mode custom** avec vos propres clés (PK/KEK/db) plutôt que les
  clés Microsoft par défaut si vous voulez contrôler ce qui boote — indispensable pour des
  images maison signées.
- **PXE/iPXE** : le boot réseau doit lui aussi être signé ; un serveur qui netboot une image
  non signée contourne tout le secure boot local.
- **Cartes d'extension** : les **Option ROM** (cartes réseau, contrôleurs RAID) sont
  vérifiées par le secure boot UEFI — d'où l'importance de firmwares signés par le vendeur
  (voir section 62).
- **Inventaire** : auditer régulièrement l'état Secure Boot du parc via Redfish/BMC
  (un serveur avec Secure Boot désactivé « pour dépanner » et jamais réactivé = classique).

## 55. Measured Boot : la différence qui compte

Le measured boot **n'empêche pas** de démarrer un composant non signé : il **enregistre**
son hash dans les PCR. Intérêt : **détection** plutôt que prévention. En production :

- Les mesures sont collectées par un **agent** (ex. Keylime, ou l'infrastructure
  d'attestation du cloud) et comparées à des **valeurs de référence** (golden measurements).
- Tout écart = alerte (firmware modifié, bootloader remplacé).
- Combiné au secure boot : le secure boot **bloque** le connu-mauvais (non signé), le
  measured boot **détecte** le signé-mais-modifié ou le downgrade.

Sans infrastructure de collecte et de comparaison, le measured boot ne sert à rien : c'est
un investissement **système**, pas une case à cocher.

## 56. Attestation à distance : le protocole

L'attestation à distance permet à un **vérificateur** (serveur central) de s'assurer qu'une
machine distante a démarré dans un état sain, **sans lui faire confiance a priori** :

1. Le vérificateur envoie un **nonce** (défi aléatoire anti-rejeu).
2. Le TPM de la machine signe (**quote**) les PCR + le nonce avec sa clé d'attestation (AK).
3. Le vérificateur contrôle la signature, la fraîcheur du nonce, et **compare les PCR aux
   valeurs de référence**.
4. Décision : sain → la machine rejoint le cluster / reçoit ses secrets ; non sain →
   quarantaine.

C'est le mécanisme derrière l'**admission des nœuds** dans les clusters sensibles
(Kubernetes, hyperviseurs) et la **distribution de secrets** (un secret n'est délivré qu'à
une machine attestée — voir Trustee/Keylime, section 156).

## 57. Attestation : qui vérifie quoi (le vérificateur)

Le vérificateur (verifier) a besoin de trois choses :

- **La chaîne de confiance du TPM** : le certificat de la clé EK, signé par le fabricant
  (Infineon, STMicro...), pour prouver que le TPM est authentique.
- **Les valeurs de référence** : hashes attendus du firmware, bootloader, noyau — à
  maintenir à **chaque mise à jour** (sinon faux positifs en cascade après un patch).
- **Une politique** : quels PCR sont critiques, quelle tolérance (ex. PCR 8-15 variables
  selon le noyau, PCR 0-7 stricts).

Solutions : **Keylime** (open source, CNCF), les services d'attestation des clouds,
**Intel Trust Authority** (service managé). En 2026, l'attestation devient aussi la brique
de base du **confidential computing** (voir sections 154-156).

## 58. Intel Boot Guard / AMD Platform Secure Boot (à vérifier)

Les constructeurs CPU proposent des racines de confiance matérielles au tout premier stade :

- **Intel Boot Guard** : vérifie la signature du firmware (BIOS/UEFI) dès le reset, avec une
  clé programmée en usine (fuses). En mode « Verified Boot », un firmware non signé ne
  démarre pas.
- **AMD Platform Secure Boot (PSB)** : équivalent AMD, la clé est programmée chez l'OEM et
  lie le CPU à la plateforme.

Ces mécanismes **complètent** le TPM (ils sont en amont). Détails d'activation par
plateforme et par OEM (HPE, Dell, Supermicro) : **à vérifier** dans la documentation du
serveur — c'est souvent une option à commander ou à activer en usine, pas un réglage
rétroactif.

## 59. BMC : le talon d'Achille des serveurs

Le **BMC** (Baseboard Management Controller : iLO chez HPE, iDRAC chez Dell, XCC chez
Lenovo, OpenBMC sur les plateformes ouvertes) est un **ordinateur dans l'ordinateur** :
SoC ARM séparé, son propre OS (Linux), sa propre pile réseau, accès **total** au serveur
(allumage/extinction, console KVM, montage d'images ISO, lecture des capteurs, et souvent
accès au bus système). Un BMC compromis = contrôle total persistant, **invisible depuis
l'OS** (l'OS ne voit pas le BMC). C'est la cible n°1 des attaquants sur un datacenter, et
le point le plus négligé : mots de passe par défaut jamais changés, firmwares jamais
patchés, BMC exposés sur le réseau de production.

## 60. Durcissement BMC : checklist technique

