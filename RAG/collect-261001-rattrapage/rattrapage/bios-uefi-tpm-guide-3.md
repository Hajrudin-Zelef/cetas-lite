---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-3
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [91, 214]
sha256: 64a857d7c9f2cc859c3659cbb33fdf876a87d5efc3eb7106e7b1073774843097
---

# BIOS / UEFI — Secure Boot — TPM 2.0

78. [Cas pratique 15 : double boot Windows/Linux cassé après une mise à jour Windows](#78-cas-pratique-15--double-boot-windowslinux-cass%C3%A9-apr%C3%A8s-une-mise-%C3%A0-jour-windows)
79. [Erreurs classiques : les 18 pièges à éviter](#79-erreurs-classiques--les-18-pi%C3%A8ges-%C3%A0-%C3%A9viter)
80. [Checklist : configuration d'un poste neuf en entreprise](#80-checklist--configuration-dun-poste-neuf-en-entreprise)
81. [Pense-bête de poche](#81-pense-b%C3%AAte-de-poche)
82. [Glossaire](#82-glossaire)
83. [Quiz : 10 questions pour valider](#83-quiz--10-questions-pour-valider)
84. [Pour aller plus loin](#84-pour-aller-plus-loin)

---

## 1. Le BIOS Legacy : l'héritage du PC

Le **BIOS** (*Basic Input/Output System*) est le firmware historique des PC, né avec l'IBM PC en 1981. C'est un programme stocké dans une puce mémoire flash de la carte mère, exécuté en premier à la mise sous tension.

### Ce que fait le BIOS au démarrage

1. **POST** (*Power-On Self-Test*) : test rapide du processeur, de la mémoire vive, du clavier, des disques.
2. **Initialisation matérielle** : configuration basique des périphériques via des interruptions 16 bits (INT 13h pour les disques, INT 10h pour l'affichage...).
3. **Recherche d'un périphérique bootable** : lecture du **MBR** (secteur 0) du premier disque trouvé dans l'ordre de boot.
4. **Transfert du contrôle** : chargement en mémoire du code du MBR (446 octets de code exécutable) puis exécution.

### Les limites historiques du BIOS

| Limite | Détail |
|---|---|
| Mode 16 bits réel | Le BIOS s'exécute en mode réel 16 bits du processeur, avec 1 Mo de mémoire adressable. |
| Disques ≤ 2 To | Adressage des secteurs sur 32 bits (LBA 32 bits) → 2^32 secteurs × 512 octets ≈ 2,2 To. |
| 4 partitions primaires | Structure MBR limitée (voir section 5). |
| Pas de réseau natif | Aucune pile réseau : impossible de booter en HTTP, de vérifier une signature, etc. |
| Lenteur | Initialisation séquentielle, pas de parallélisme, POST parfois long. |
| Sécurité quasi nulle | Aucune vérification d'intégrité du chargeur : un bootkit peut s'installer dans le MBR sans être détecté. |
| Interface texte | Configuration via une interface clavier en mode texte, navigation rudimentaire. |

### Pourquoi c'est important de le connaître encore

- Des millions de machines en entreprise tournent encore en mode **Legacy/CSM**.
- Les outils d'imagerie anciens, certains logiciels de clonage et vieux OS (Windows 7 sans UEFI, DOS) exigent le mode Legacy.
- Comprendre le BIOS éclaire *pourquoi* l'UEFI existe : chaque fonctionnalité UEFI répond à une limite du BIOS.

> 💡 **En résumé :** le BIOS est un firmware 16 bits, simple, rapide à comprendre, mais limité (2 To, pas de sécurité, pas de réseau). Il a régné de 1981 à ~2012.

---

## 2. L'UEFI : l'architecture moderne

L'**UEFI** (*Unified Extensible Firmware Interface*) est le successeur du BIOS, standardisé par l'**UEFI Forum**. Intel l'a initié (sous le nom EFI pour Itanium, fin des années 1990) ; il s'est généralisé sur PC à partir de ~2012.

### Différences d'architecture fondamentales

| Aspect | BIOS | UEFI |
|---|---|---|
| Mode processeur | 16 bits réel | 32 ou 64 bits protégé |
| Mémoire adressable | 1 Mo | Toute la RAM |
| Langage des modules | Assembleur 16 bits | C (drivers UEFI compilés) |
| Pilotes | Intégrés au firmware, figés | **Drivers UEFI** chargeables (système de fichiers, réseau...) |
| Interface | Texte, clavier | Graphique possible, souris, police TrueType |
| Partitionnement | MBR uniquement | **GPT** (et MBR via CSM) |
| Taille disque bootable | ≤ 2 To | **> 9 Zo** (Zettaoctets) en théorie |
| Réseau | Non | Pile **TCP/IP** intégrée (HTTP boot possible) |
| Sécurité | Aucune | **Secure Boot**, TPM mesuré |
| Configuration stockée | CMOS (pile) | **NVRAM** (variables UEFI persistantes) |
| Extensibilité | Quasi nulle | Applications UEFI (shell, diagnostics, utilitaires) |

### Les composants d'un firmware UEFI

```
┌─────────────────────────────────────────────────┐
│  SEC (Security)        Phase d'initialisation   │
│  PEI (Pre-EFI Init)    Mémoire, CPU de base     │
│  DXE (Driver Exec.)    Pilotes, périphériques   │
│  BDS (Boot Dev Select) Choix du périphérique    │
│  TSL (Transient Sys.)  Chargeurs, OS            │
│  RT (Runtime)          Services après boot OS   │
└─────────────────────────────────────────────────┘
```

1. **SEC** : le processeur sort du reset, initialise un début d'environnement (cache-as-RAM).
2. **PEI** : initialise la mémoire vive permanente (RAM), découvre le chipset.
3. **DXE** : charge les pilotes UEFI (contrôleur SATA/NVMe, USB, réseau, affichage GOP...). C'est ici que le TPM est généralement interrogé et que Secure Boot vérifie les signatures.
4. **BDS** : le **Boot Manager** lit les variables de boot en NVRAM, affiche éventuellement un menu, puis lance l'application EFI choisie (ex. `\EFI\Microsoft\Boot\bootmgfw.efi`).
5. **TSL** : l'OS démarre ; il peut encore appeler les services de boot UEFI.
6. **RT** : une fois l'OS lancé (`ExitBootServices`), seuls les **services d'exécution** restent disponibles (horloge temps réel, variables NVRAM, reset/shutdown, capsule de mise à jour firmware).

### Les services UEFI en pratique

- **Boot Services** : allocation mémoire, protocoles (Block I/O, Simple File System, Graphics Output...). Disponibles jusqu'au démarrage de l'OS.
- **Runtime Services** : `GetVariable` / `SetVariable` (lecture/écriture NVRAM), `GetTime`, `ResetSystem`, `UpdateCapsule` (mise à jour firmware depuis l'OS !). C'est grâce à `UpdateCapsule` que Windows Update peut proposer des mises à jour de firmware.

> 💡 **En résumé :** l'UEFI est un mini-système d'exploitation pré-boot : pilotes en C, réseau, sécurité cryptographique, variables persistantes. Tout ce que le BIOS ne savait pas faire.

---

## 3. BIOS vs UEFI : comparatif détaillé

Tableau de synthèse pour décider du mode à utiliser sur le parc :

| Critère | BIOS (Legacy) | UEFI natif |
|---|---|---|
| Année d'introduction | 1981 | ~2006 (généralisé ~2012) |
| Architecture CPU au boot | 16 bits réel | 64 bits (long mode) |
| Disque système max | 2 To | 9,4 Zo (limite GPT) |
| Table de partitions | MBR | GPT |
| Nombre de partitions | 4 primaires (ou 3 + étendue) | 128 par défaut |
| Vitesse de démarrage | Lente (POST séquentiel) | Rapide (Fast Boot, init parallèle) |
| Secure Boot | ❌ Impossible | ✅ Natif |
| TPM mesuré au boot | ❌ | ✅ (Measured Boot) |
| Boot réseau | PXE via ROM optionnelle | PXE + **HTTP Boot** natif |
| Shell pré-boot | ❌ | ✅ (UEFI Shell en option) |
| Interface de config | Texte | Graphique / souris |
| Variables persistantes | CMOS 256 octets | NVRAM (plusieurs Mo) |
| Windows 11 | ❌ Incompatible | ✅ Requis |
| BitLocker + TPM 2.0 | ⚠️ Partiel (TPM 1.2) | ✅ Plein support |
| Support constructeurs | Abandonné (CSM retiré depuis ~2020) | Standard actuel |

### Recommandation d'entreprise (2026)

- **Tout poste neuf : UEFI natif, CSM désactivé, Secure Boot activé, TPM 2.0 activé.** C'est le prérequis Windows 11 et la base du durcissement.
- **Parc existant :** planifier la migration Legacy → UEFI (via `mbr2gpt`, section 8) avant toute montée en Windows 11.
- **Exceptions :** vieux logiciels industriels, DOS, Windows 7 32 bits, certains outils de clonage propriétaires → maintenir quelques machines en Legacy/CSM, isolées et documentées.

---

## 4. CSM (Compatibility Support Module)

