---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-6
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [597, 804]
sha256: fde21565e22cf461c330f413bbc65ef9f6f3fb74c26f5731928a07897008a10b
---

# Créer une entrée pour booter une fois sur un périphérique (BootNext) :
# (souvent plus simple via la touche F12 du constructeur au démarrage)
```

### Côté Linux (pour info, en multi-boot)

```bash
efibootmgr -v        # lister les entrées avec chemins complets
efibootmgr -o 0001,0003,0002   # changer l'ordre
efibootmgr -n 0003   # booter une fois sur 0003 (BootNext)
```

### Pourquoi c'est utile en entreprise

- Après un clonage ou un changement de disque, des **entrées fantômes** peuvent subsister → nettoyage via `bcdedit /delete` ou le setup.
- En déploiement (MDT/SCCM), on peut forcer l'ordre de boot par script.
- Un malware/bootkit peut ajouter une entrée : `bcdedit /enum firmware` fait partie de l'audit.

---

## 13. BCD (Boot Configuration Data)

Le **BCD** remplace le `boot.ini` depuis Windows Vista. C'est une base de registre (ruche) stockée dans :
- **UEFI :** `S:\EFI\Microsoft\Boot\BCD` (sur l'ESP, sans extension),
- **Legacy :** `C:\Boot\BCD`.

### Objets principaux

| Objet (GUID) | Rôle |
|---|---|
| `{bootmgr}` | **Windows Boot Manager** : comportement du menu (timeout, affichage, ordre) |
| `{current}` / `{default}` | L'entrée OS démarrée par défaut |
| `{fwbootmgr}` | Le Boot Manager du **firmware** (variables NVRAM) |
| `{memdiag}` | Diagnostic mémoire Windows |
| GUIDs divers | Entrées secondaires (WinRE, autres OS, options de démarrage) |

### Paramètres courants d'une entrée OS

| Paramètre | Signification |
|---|---|
| `device` / `osdevice` | Partition contenant le chargeur / l'OS (`partition=C:`) |
| `path` | `\Windows\system32\winload.efi` (UEFI) ou `winload.exe` (Legacy) |
| `systemroot` | `\Windows` |
| `nx` | DEP : `OptIn` (défaut) / `OptOut` / `AlwaysOn` / `AlwaysOff` |
| `pae` | Physical Address Extension |
| `testsigning` | `Yes` = accepte les drivers en mode test (à éviter en prod !) |
| `nointegritychecks` | Désactive la vérification de signature des drivers (danger) |
| `recoveryenabled` | Active WinRE en cas d'échec |
| `bootlog` | Journal `ntbtlog.txt` |

```powershell
# Vue d'ensemble
bcdedit /enum

# Détail de l'entrée courante
bcdedit /enum '{current}'

# Vue firmware (NVRAM)
bcdedit /enum firmware

# Exporter le BCD avant toute modification (réflexe !)
bcdedit /export C:\BCD-backup-2026-09-27
```

---

## 14. bcdedit : le couteau suisse du boot Windows

`bcdedit.exe` (console administrateur obligatoire) lit et modifie le BCD. **Toujours exporter avant de modifier.**

### Commandes de lecture

```powershell
bcdedit                          # équivalent de /enum ACTIVE
bcdedit /enum ALL                 # tout : firmware + bootmgr + OS + WinRE
bcdedit /enum '{current}' /v      # détail verbeux (GUIDs complets)
bcdedit /v                        # affiche les GUIDs au lieu des alias
```

### Commandes d'écriture courantes

```powershell
# --- Menu de démarrage ---
# Timeout du menu (secondes)
bcdedit /timeout 10

# Entrée par défaut
bcdedit /default '{current}'

# Ordre d'affichage des OS
bcdedit /displayorder '{current}' '{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}'

# --- Options de démarrage avancées ---
# Activer le menu F8 hérité (utile en dépannage, désactivé par défaut depuis Win8)
bcdedit /set '{current}' bootmenupolicy Legacy
# Revenir au comportement moderne :
bcdedit /set '{current}' bootmenupolicy Standard

# Démarrage en mode sans échec minimal (puis reboot)
bcdedit /set '{current}' safeboot Minimal
# Annuler :
bcdedit /deletevalue '{current}' safeboot

# Activer le journal de démarrage (C:\Windows\ntbtlog.txt)
bcdedit /set '{current}' bootlog Yes

# Désactiver la signature des pilotes (DÉPANNAGE UNIQUEMENT, jamais en prod)
bcdedit /set '{current}' testsigning Yes
# ⚠️ Laisse un filigrane "Test Mode" et désactive une protection clé.

# --- Gestion des entrées ---
# Copier l'entrée courante pour créer une entrée de test
bcdedit /copy '{current}' /d "Windows 11 (Test)"

# Supprimer une entrée (utiliser le GUID affiché par /v)
bcdedit /delete '{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}'

# Réparer l'entrée : repointer vers la bonne partition
bcdedit /set '{current}' device partition=C:
bcdedit /set '{current}' osdevice partition=C:
```

### Sauvegarde / restauration

```powershell
# Export (à faire AVANT toute modification)
bcdedit /export C:\BCD-backup-avant-modif

# Import (depuis WinRE ou un OS fonctionnel)
bcdedit /import C:\BCD-backup-avant-modif
```

> 💡 **Réflexe d'entreprise :** avant toute intervention sur le BCD (changement de disque, clonage, multi-boot), exportez le BCD et notez le `diskpart → list volume`. 30 secondes qui évitent des heures de dépannage.

---

## 15. bcdboot : reconstruire les fichiers de démarrage

`bcdboot.exe` **recrée** les fichiers de boot (chargeurs + BCD) à partir d'une installation Windows existante. C'est l'outil n°1 quand « BOOTMGR is missing » / `0xc000000f` / ESP effacée.

### Cas d'usage

| Symptôme | Commande |
|---|---|
| ESP formatée ou corrompue | `bcdboot C:\Windows /s S: /f UEFI` |
| BCD supprimé/corrompu | `bcdboot C:\Windows` (recrée tout) |
| Nouveau disque cloné qui ne boote pas | `bcdboot` + vérification `bcdedit` |
| Ajouter Windows au menu après install Linux | `bcdboot C:\Windows` (recrée l'entrée `{bootmgr}`) |

### Procédure depuis WinRE (environnement de récupération)

```
1. Démarrer sur une clé USB d'installation Windows → "Réparer l'ordinateur"
   → Dépannage → Invite de commandes.
2. Identifier les lettres (elles peuvent changer en WinRE !) :
     diskpart → list volume → noter C: (Windows) et l'ESP (FAT32, ~100 Mo)
3. Assigner une lettre à l'ESP si besoin :
     select volume 2   (l'ESP)
     assign letter=S
     exit
4. Reconstruire :
     bcdboot C:\Windows /s S: /f UEFI
     → "Boot files successfully created."
5. Vérifier :
     bcdedit /store S:\EFI\Microsoft\Boot\BCD /enum
6. Retirer la lettre (optionnel) puis redémarrer.
```

### Syntaxe détaillée

```powershell
# Forme complète
bcdboot C:\Windows /s S: /f UEFI /l fr-FR
#        ^source     ^cible ^firmware ^langue du menu
# /f ALL    : copie les fichiers pour UEFI ET BIOS (clé USB universelle)
# /l fr-FR  : langue ; par défaut celle de l'OS
```

### bcdboot vs bootrec (Legacy)

| Outil | Firmware | Rôle |
|---|---|---|
| `bcdboot` | UEFI (et BIOS) | Recrée chargeurs + BCD depuis `C:\Windows` |
| `bootrec /fixmbr` | BIOS Legacy | Réécrit le code MBR (pas la table !) |
| `bootrec /fixboot` | BIOS Legacy | Réécrit le secteur de boot de la partition |
| `bootrec /rebuildbcd` | BIOS Legacy | Reconstruit le BCD en scannant les installations |

> ⚠️ Sur un système UEFI/GPT, `bootrec /fixmbr` est inutile voire contre-productif : utilisez `bcdboot`.

---

## 16. Secure Boot : principe général

**Secure Boot** est une fonctionnalité UEFI (chapitre 27 de la spec) qui garantit que **seuls des logiciels signés par une autorité de confiance** s'exécutent pendant le démarrage : firmware → chargeur → noyau → pilotes critiques.

### L'idée en une phrase

> Chaque maillon de la chaîne de démarrage **vérifie la signature cryptographique** du maillon suivant avant de lui passer la main. Un chargeur modifié ou inconnu = démarrage refusé.

### Ce que Secure Boot protège

- ✅ Chargeurs de boot (bootmgfw.efi, shim Linux signé...)
- ✅ Noyau de l'OS et HAL
- ✅ Pilotes de démarrage critiques (via ELAM — *Early Launch Anti-Malware*)
- ✅ Pilotes en mode noyau (sur Windows 64 bits : signature obligatoire de toute façon)

### Ce que Secure Boot ne protège PAS

