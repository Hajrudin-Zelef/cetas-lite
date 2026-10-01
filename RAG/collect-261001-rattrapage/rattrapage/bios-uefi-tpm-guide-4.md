---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-4
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "intel"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [215, 389]
sha256: 4918ac4fad0bb29ca0485b56ca261969d3bcc8254779bb251db32d5b22bb169f
---

# BIOS / UEFI — Secure Boot — TPM 2.0

Le **CSM** est un module de compatibilité intégré à certains firmwares UEFI : il **émule un BIOS Legacy** pour démarrer des OS et périphériques qui ne connaissent pas l'UEFI.

### Comment ça marche

1. Le firmware démarre en mode UEFI.
2. Si le périphérique de boot n'a pas de chargeur EFI valide (pas d'ESP, pas de `\EFI\BOOT\bootx64.efi`), le CSM prend le relais.
3. Il fournit les interruptions 16 bits (INT 13h...) et exécute le MBR comme un vrai BIOS.

### Options typiques dans le setup

| Option | Signification |
|---|---|
| `UEFI` / `UEFI only` | Pas de CSM, boot UEFI natif uniquement. **Recommandé.** |
| `Legacy` / `Legacy only` / `CSM` | Émulation BIOS, boot MBR uniquement. |
| `UEFI + Legacy` / `Both` | Essaie UEFI puis Legacy. Pratique en transition, à éviter à terme. |
| `Legacy ROM` / `Option ROM` | Active les ROMs d'extension 16 bits (vieilles cartes réseau pour PXE Legacy, cartes RAID...). |

### Pourquoi désactiver le CSM

1. **Sécurité :** le mode Legacy contourne Secure Boot. Un attaquant qui active le CSM peut booter un OS non signé.
2. **Windows 11 :** exige le mode UEFI natif (le CSM doit être désactivé).
3. **Démarrage :** le CSM ralentit le POST (initialisation des ROMs 16 bits).
4. **Avenir :** Intel et AMD ont annoncé la fin du CSM (Intel : plus de CSM sur les plateformes récentes ; les constructeurs le retirent de leurs firmwares depuis ~2020-2022). Sur du matériel récent, l'option n'existe tout simplement plus.

### Piège classique

> Un disque GPT n'est **pas bootable** en mode Legacy/CSM (sauf cas exotiques avec partition BIOS boot). Inversement, un disque MBR avec installation Windows Legacy ne boote pas en UEFI natif. **Le mode du firmware et le style de partition du disque système doivent correspondre.** C'est la cause n°1 des « plus de boot » après un changement de réglage.

Vérifier le mode de démarrage actuel de Windows :

```powershell
# Méthode 1 : variable d'environnement système
# Si le dossier existe, Windows a démarré en UEFI
Test-Path "$env:SystemRoot\Panther"

# Méthode 2 : fiable — interroger le firmware via WMI
$bios = Get-CimInstance -ClassName Win32_BIOS
$bios | Select-Object Manufacturer, Name, SMBIOSBIOSVersion, ReleaseDate

# Méthode 3 : msinfo32 → "Mode BIOS" = UEFI ou Hérité
msinfo32
```

```powershell
# Méthode 4 (la plus directe) : le chemin du chargeur
# En UEFI : \EFI\Microsoft\Boot\bootmgfw.efi — En Legacy : \Windows\system32\winload.exe
bcdedit /enum '{current}' | Select-String 'path'
```

---

## 5. MBR : structure et fonctionnement

Le **MBR** (*Master Boot Record*) occupe le **secteur 0** (512 octets) du disque.

### Layout des 512 octets

| Offset | Taille | Contenu |
|---|---|---|
| 0x000 | 446 octets | **Code de démarrage** (bootstrap) exécuté par le BIOS |
| 0x1BE | 16 octets | Entrée de partition n°1 |
| 0x1CE | 16 octets | Entrée de partition n°2 |
| 0x1DE | 16 octets | Entrée de partition n°3 |
| 0x1EE | 16 octets | Entrée de partition n°4 |
| 0x1FE | 2 octets | **Signature** `55 AA` (indique un MBR valide) |

### Une entrée de partition (16 octets)

| Offset | Taille | Champ |
|---|---|---|
| 0x00 | 1 | Indicateur de boot (`0x80` = active/bootable) |
| 0x01 | 3 | Début CHS (obsolète) |
| 0x04 | 1 | **Type de partition** (0x07 = NTFS, 0x0B/0x0C = FAT32, 0x83 = Linux...) |
| 0x05 | 3 | Fin CHS (obsolète) |
| 0x08 | 4 | LBA de début (32 bits) |
| 0x0C | 4 | Nombre de secteurs (32 bits) |

### Conséquences directes

- **4 entrées × 16 octets = 4 partitions primaires maximum.** Pour en avoir plus, l'une d'elles devient une partition **étendue** contenant des partitions **logiques** (chaînage EBR).
- **Adressage 32 bits** : 2^32 secteurs × 512 o = **2 To** (2,2 To exactement : 2 199 023 255 552 octets). Au-delà, l'espace est inutilisable.
- **Une seule copie** : pas de redondance. Si le secteur 0 est corrompu → disque illisible (récupérable avec des outils comme TestDisk).

Inspecter le style de partition en PowerShell :

```powershell
Get-Disk | Select-Object Number, FriendlyName, PartitionStyle, Size,
    @{n='SizeGB';e={[math]::Round($_.Size/1GB,1)}}
```

---

## 6. GPT : structure et fonctionnement

Le **GPT** (*GUID Partition Table*), défini dans la spécification UEFI, remplace le MBR.

### Layout du disque GPT

```
LBA 0    : Protective MBR (faux MBR : 1 partition type 0xEE couvrant tout le disque)
LBA 1    : En-tête GPT primaire (signature "EFI PART", CRC32, LBA de la sauvegarde)
LBA 2-33 : Table des entrées de partitions primaire (128 entrées × 128 octets)
LBA 34.. : Partitions de données
   ...    : Table des entrées secondaire (copie de secours)
   fin-1  : En-tête GPT secondaire (copie de secours)
```

### Points clés

| Aspect | Détail |
|---|---|
| **Protective MBR** | Protège le disque des vieux outils MBR : une partition fictive de type `0xEE` couvre tout le disque pour éviter tout écrasement accidentel. |
| **128 partitions** | 128 entrées par défaut (extensible), chacune identifiée par un **GUID unique**. |
| **Types par GUID** | Chaque partition a un GUID de type : ESP (`C12A7328-F81F-11D2-BA4B-00A0C93EC93B`), Microsoft Basic Data, Microsoft Reserved (MSR), Linux filesystem, etc. |
| **Noms** | Chaque partition peut porter un nom UTF-16 (36 caractères). |
| **Redondance** | En-tête + table dupliqués en fin de disque → **auto-réparation** si le primaire est corrompu. |
| **CRC32** | Intégrité de l'en-tête et de la table vérifiée à chaque lecture. |
| **Taille max** | LBA sur 64 bits → 9,4 Zo (Zettaoctets). En pratique, limité par l'OS et le système de fichiers. |

### Les partitions typiques d'un Windows en UEFI/GPT

| N° | Type (GUID) | Taille | Rôle |
|---|---|---|---|
| 1 | ESP | 100–260 Mo | **EFI System Partition** (FAT32) : chargeurs de boot |
| 2 | MSR | 16 Mo | Microsoft Reserved : zone de travail pour conversions |
| 3 | Basic Data | le reste − 500 Mo | `C:` Windows (NTFS) |
| 4 | Recovery | ~500 Mo–1 Go | WinRE (environnement de récupération) |

Voir le détail :

```powershell
Get-Disk -Number 0 | Get-Partition |
    Select-Object PartitionNumber, DriveLetter, Size,
    @{n='SizeGB';e={[math]::Round($_.Size/1GB,2)}},
    Type, GptType
```

---

## 7. MBR vs GPT : limites et comparatif

| Critère | MBR | GPT |
|---|---|---|
| Âge | 1983 | 2000s (UEFI) |
| Taille max disque | **2 To** | **9,4 Zo** |
| Partitions primaires | **4** (ou 3 + étendue) | **128** |
| Redondance | Aucune (1 secteur) | En-tête + table dupliqués |
| Contrôle d'intégrité | Aucun | CRC32 |
| Nommage des partitions | Non | Oui (UTF-16) |
| Boot UEFI natif | ❌ Non | ✅ Oui |
| Boot BIOS Legacy | ✅ Oui | ❌ Non (sauf CSM + cas spéciaux) |
| Windows 11 (disque système) | ❌ Incompatible | ✅ Requis |
| Outils de réparation | TestDisk, bootrec | bcdboot, bcdedit, diskpart |

### Règle d'or

> **Disque système Windows moderne = GPT + UEFI.** Le MBR ne se justifie plus que pour : vieux OS, disques de données sur très vieilles machines, clés USB de compatibilité universelle (et encore, l'exFAT/GPT est lisible partout aujourd'hui).

---

## 8. Conversion MBR → GPT : mbr2gpt

`mbr2gpt.exe` (présent depuis Windows 10 1703) convertit un disque système MBR en GPT **sans perte de données** et **sans réinstallation**, puis reconfigure le boot en UEFI.

### Préconditions (toutes obligatoires)

- [ ] Windows 10 version 1703+ / Windows 11.
- [ ] Le disque est le **disque système** (celui qui contient Windows).
- [ ] **3 partitions maximum** sur le disque MBR (le GPT a besoin de place pour l'ESP + MSR).
- [ ] Pas de partition étendue/logique compliquée (l'outil gère les cas simples).
- [ ] **BitLocker suspendu ou désactivé** sur le volume (sinon la conversion échoue ou déclenche une récupération).
- [ ] Sauvegarde à jour (principe de précaution, même si l'outil est fiable).

### Procédure complète

