---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-5
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [390, 596]
sha256: 27b8b15074949404d5224e136472e979530be0eca0112889b76ea143c1c8d91e
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```powershell
# 1. Vérifier l'état actuel (depuis Windows, en administrateur)
Get-Disk -Number 0 | Select-Object PartitionStyle   # doit afficher MBR

# 2. Suspendre BitLocker si actif (indispensable)
Suspend-BitLocker -MountPoint "C:" -RebootCount 1
# ou : manage-bde -protectors -disable C:

# 3. VALIDER la conversion (ne modifie rien)
mbr2gpt /validate /disk:0 /allowFullOS

# 4. CONVERTIR (depuis Windows ; redémarrage non nécessaire pour la conversion elle-même)
mbr2gpt /convert /disk:0 /allowFullOS
```

Sortie typique en cas de succès :

```
MBR2GPT: Attempting to convert disk 0
MBR2GPT: Retrieving layout of disk
MBR2GPT: Validating layout, disk sector size is: 512 bytes
MBR2GPT: Trying to shrink the system partition
MBR2GPT: Trying to shrink the OS partition
MBR2GPT: Creating the EFI system partition
MBR2GPT: Installing the new boot files
MBR2GPT: Performing the layout conversion
MBR2GPT: Migrating default boot entry
MBR2GPT: Adding recovery boot entry
MBR2GPT: Fixing drive letter mapping
MBR2GPT: Conversion completed successfully
MBR2GPT: Before the new system can boot properly you need to switch the
firmware to boot to UEFI mode!
```

### 5. Basculer le firmware en UEFI (après la conversion)

1. Redémarrer, entrer dans le setup (F2 / DEL / F10 selon constructeur).
2. Passer le mode de boot de **Legacy** à **UEFI** (désactiver le CSM).
3. Vérifier que **Secure Boot** est activable/activé.
4. Sauvegarder et redémarrer.

```powershell
# 6. Vérification post-migration
Get-Disk -Number 0 | Select-Object PartitionStyle   # doit afficher GPT
Confirm-SecureBootUEFI                              # idéalement True
```

### En cas d'échec de validation

| Message | Cause probable | Solution |
|---|---|---|
| `Disk layout validation failed` | > 3 partitions, ou partition étendue | Supprimer/fusionner les partitions superflues (après sauvegarde). |
| `Cannot find OS partition` | Partition Windows non reconnue | Vérifier avec `diskpart` → `list volume`. |
| Échec avec BitLocker | Chiffrement actif | `Suspend-BitLocker` puis réessayer. |
| `ValidateLayout: Too many MBR partitions` | 4 partitions primaires | Supprimer une partition (souvent la partition de récupération d'usine). |

> ⚠️ **mbr2gpt ne convertit pas dans l'autre sens.** Pour GPT → MBR, il faut sauvegarder, recréer la table (diskpart `clean` + `convert mbr`) et réinstaller : opération destructive.

---

## 9. Cas d'usage : quand utiliser MBR ou GPT

| Situation | Choix recommandé |
|---|---|
| Poste Windows 10/11 neuf | **GPT + UEFI** (obligatoire pour Win11) |
| Migration vers Windows 11 | Convertir en GPT via `mbr2gpt` |
| Disque de données > 2 To | **GPT** (le MBR ne verra que 2 To) |
| Disque de données ≤ 2 To sur machine récente | **GPT** (redondance, CRC, nommage) |
| Vieux PC en BIOS Legacy (Windows 7, DOS, industriel) | **MBR** (le seul bootable en Legacy) |
| Clé USB multi-usage (boot + stockage) | **GPT** + FAT32 si UEFI only ; **MBR** + FAT32 si compatibilité Legacy requise |
| Serveur / NAS | **GPT** systématiquement |

---

## 10. Séquence de boot UEFI pas à pas

Chronologie complète, de l'appui sur le bouton à l'écran de connexion Windows :

```
① Mise sous tension
   └─ Le CPU sort du reset, exécute le firmware (SEC → PEI)
② Initialisation matérielle (DXE)
   ├─ Pilotes UEFI : NVMe/SATA, USB, GOP (affichage), réseau
   ├─ Mesure TPM (Measured Boot) : chaque composant est "mesuré" (hash)
   │   dans les PCR du TPM
   └─ Secure Boot : vérification des signatures des pilotes/chargeurs
③ BDS — Boot Device Selection
   ├─ Lecture des variables Boot#### en NVRAM
   ├─ Application de BootOrder
   └─ (Optionnel) menu de boot F12 / timeout
④ Exécution du chargeur : \EFI\Microsoft\Boot\bootmgfw.efi (depuis l'ESP)
   └─ Secure Boot vérifie sa signature Microsoft
⑤ bootmgfw.efi lit le BCD
   ├─ Affiche le menu (multi-boot éventuel)
   └─ Sélectionne l'entrée OS
⑥ Chargement de winload.efi (+ signature vérifiée)
   ├─ winload.efi vérifie à son tour les composants critiques :
   │   noyau ntoskrnl.exe, HAL, drivers de démarrage (ELAM)
   └─ ExitBootServices : fin des services de boot UEFI
⑦ Noyau Windows : initialisation, chargement des drivers, session
⑧ Écran de connexion
```

### Measured Boot vs Secure Boot : ne pas confondre

| | Secure Boot | Measured Boot (TPM) |
|---|---|---|
| Rôle | **Bloquer** ce qui n'est pas signé | **Enregistrer** (mesurer) ce qui a démarré |
| Action en cas de problème | Refuse de booter | Boote quand même, mais les PCR changent |
| Utilisé par | Vérification de signature | BitLocker (scelle la clé sur les PCR), attestation |

---

## 11. ESP (EFI System Partition)

L'**ESP** est la partition FAT32 qui contient les chargeurs de démarrage UEFI.

### Caractéristiques

| Propriété | Valeur |
|---|---|
| Système de fichiers | **FAT32** (obligatoire par la spec UEFI) |
| GUID de type | `C12A7328-F81F-11D2-BA4B-00A0C93EC93B` |
| Taille Windows | 100 Mo (disques ≤ 16 Go : 100 Mo ; en pratique 100–260 Mo) |
| Lettre de lecteur | **Aucune** par défaut (montage manuel possible) |
| Contenu Windows | `\EFI\Microsoft\Boot\bootmgfw.efi`, `\EFI\Boot\bootx64.efi` (fallback), polices, BCD |

### Arborescence typique

```
ESP (FAT32)
└── EFI
    ├── Boot
    │   └── bootx64.efi          ← chargeur de secours (chemin "removable media")
    ├── Microsoft
    │   └── Boot
    │       ├── bootmgfw.efi     ← Windows Boot Manager
    │       ├── bootmgr.efi
    │       ├── BCD              ← base de configuration de boot
    │       ├── memtest.efi
    │       └── Fonts\ / Resources\
    └── Ubuntu                  ← exemple multi-boot
        └── grubx64.efi / shimx64.efi
```

### Manipuler l'ESP (lecture seule recommandée)

```powershell
# Assigner une lettre à l'ESP pour l'inspecter (administrateur)
# 1. Identifier l'ESP
Get-Partition | Where-Object { $_.GptType -eq '{c12a7328-f81f-11d2-ba4b-00a0c93ec93b}' } |
    Select-Object DiskNumber, PartitionNumber, Size

# 2. Lui assigner une lettre temporaire via diskpart
diskpart
#   list disk
#   select disk 0
#   list partition
#   select partition 1        (l'ESP)
#   assign letter=S
#   exit

# 3. Explorer
Get-ChildItem S:\EFI -Recurse | Select-Object FullName, Length

# 4. Retirer la lettre après inspection (hygiène)
#   diskpart → select partition 1 → remove letter=S
```

> ⚠️ **Ne jamais formater l'ESP** sauf à savoir la reconstruire avec `bcdboot` (section 15). Ne jamais y stocker de fichiers personnels.

---

## 12. UEFI Boot Manager et NVRAM

Le **Boot Manager UEFI** est le menu/logiciel du firmware qui choisit *quoi* démarrer. Ses choix sont stockés dans des **variables NVRAM** :

| Variable | Rôle |
|---|---|
| `Boot0001`, `Boot0002`... | Une entrée de boot = chemin vers un `.efi` + description + options |
| `BootOrder` | Ordre de priorité (ex. `0001,0003,0002`) |
| `BootNext` | Boot unique au prochain redémarrage (puis effacé) |
| `Timeout` | Délai d'affichage du menu |
| `BootCurrent` | Entrée utilisée pour le boot en cours |

### Lire/écrire les entrées (Windows)

```powershell
# Lister les entrées de boot du firmware (vue Windows via BCD)
bcdedit /enum firmware

# Exemple de sortie (abrégée) :
# Firmware Boot Manager
# ---------------------
# identifier              {fwbootmgr}
# displayorder            {bootmgr}
#                         {6a5f2e1a-...}   ← entrée "ubuntu" par ex.
# timeout                 2
```

```powershell
# Changer l'ordre de boot du firmware (exemple : remettre Windows en premier)
bcdedit /set '{fwbootmgr}' displayorder '{bootmgr}' /addfirst

# Définir le timeout du menu firmware
bcdedit /set '{fwbootmgr}' timeout 5

