---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-15
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [2336, 2547]
sha256: f0bf707fbb192e60d73defd469cc6ce83d23f0d6ca93a1e4b647274e94fe7cfa
---

# BIOS / UEFI — Secure Boot — TPM 2.0

| PXE classique | iPXE |
|---|---|
| TFTP uniquement (lent, UDP) | **HTTP** (rapide, TCP, cacheable) |
| Pas de script | Scripts de boot conditionnels |
| Menu WDS uniquement | Menus dynamiques, boot multi-OS |
| Pas de HTTPS | **HTTPS**, SAN boot (iSCSI) |

### Exemple de script iPXE

```ipxe
#!ipxe
# Menu de boot réseau d'entreprise
set server http://deploy.lan.lan

:start
menu Boot reseau - ${hostname}
item --gap --             -------------------------
item win11               Installer Windows 11 (WDS)
item winpe               WinPE de depannage
item diag                Diagnostics memoire (memtest)
item --gap --             -------------------------
item shell               Shell iPXE
item reboot              Redemarrer
choose option || goto start

:win11
kernel ${server}/wdsboot.cgi/arch=x64 ||
chain ${server}/menu.ipxe || goto start

:winpe
kernel ${server}/winpe/wimboot
initrd ${server}/winpe/boot.wim boot.wim
boot

:diag
kernel ${server}/memtest86.bin
boot

:shell
shell

:reboot
reboot
```

### Chaînage depuis le PXE existant

```
1. Le client PXE classique télécharge undionly.kpxe / ipxe.efi via TFTP.
2. iPXE prend la main → DHCP → télécharge le script via HTTP.
3. Le script propose le menu → charge les images via HTTP.
```

> 💡 **En datacenter :** iPXE + HTTP est le standard pour provisionner des serveurs (avec Foreman, MAAS, ou NetBox + scripts maison). Sur les serveurs récents, l'**HTTP Boot UEFI natif** (sans iPXE) fait la même chose directement depuis le firmware.

---

## 52. Mises à jour du firmware : précautions générales

### Règle d'or

> 🔴 **Ne jamais interrompre un flash du BIOS/UEFI.** Une coupure pendant l'écriture = carte mère potentiellement inutilisable (sauf fonction de récupération, §77). C'est l'opération la plus risquée sur un poste.

### Check-list avant tout flash

- [ ] **Secteur branché** + batterie chargée à > 50 % (portable). Idéalement : onduleur (lien avec votre métier énergie !).
- [ ] **BitLocker suspendu** (`Suspend-BitLocker -RebootCount 1`) si chiffré (§37).
- [ ] **Sauvegarde** des données critiques (principe).
- [ ] **Bonne version** : vérifier le modèle exact (référence constructeur, pas « ça ressemble »).
- [ ] **Lire le changelog** : la MAJ corrige-t-elle quelque chose qui vous concerne ? (Si tout fonctionne et qu'aucune faille critique : on peut planifier plutôt que subir.)
- [ ] **Ne pas flasher pendant un orage** / sur une prise douteuse.
- [ ] Fermer les applications, désactiver temporairement l'antivirus si le constructeur le recommande.
- [ ] Noter la **version actuelle** (pour rollback éventuel).

### Méthodes de flash (général)

| Méthode | Description | Quand l'utiliser |
|---|---|---|
| **Depuis Windows** (utilitaire OEM) | `.exe` du constructeur | Parc standard, flash simple |
| **Depuis le setup UEFI** | Clé USB FAT32 + fichier `.bin`/`.cap` | Le plus sûr (pas d'OS), serveurs |
| **Capsule UEFI** (Windows Update) | MAJ proposée par WU | Parc géré, avec validation pilote |
| **À distance** (outils OEM) | Dell Command Update, HP Image Assistant... | Déploiement de masse |

### Après le flash

```
1. Le PC redémarre (parfois 2-3 fois, avec écran noir prolongé : NORMAL).
2. Entrer dans le setup : vérifier la nouvelle version.
3. ⚠️ Les réglages peuvent être réinitialisés → re-vérifier :
   ordre de boot, Secure Boot, TPM, virtualisation, mot de passe.
4. BitLocker : vérifier qu'il est "FullyEncrypted" (re-scellé).
5. Tester un redémarrage complet + une mise en veille/reprise.
```

---

## 53. Mise à jour BIOS/UEFI : Dell

### Méthodes

| Méthode | Outil | Usage |
|---|---|---|
| Exécutable Windows | `Latitude_5x40_1.2.3.exe` téléchargé du support Dell | Un poste, manuel |
| **Dell Command Update** | `dcu-cli.exe` | Parc : scan + déploiement planifié |
| Depuis le BIOS | F12 → BIOS Flash Update (fichier sur USB) | Sans OS, le plus sûr |
| Via SCCM/Intune | Dell Command Configure + packages | Masse |

### Dell Command Update en CLI (déploiement)

```powershell
# Installer silencieusement (téléchargé depuis dell.com/support)
Start-Process .\Dell-Command-Update-Application.msi -ArgumentList '/qn' -Wait

# Scanner les mises à jour disponibles
& 'C:\Program Files\Dell\CommandUpdate\dcu-cli.exe' /scan

# Appliquer BIOS + firmware (silencieux, avec reboot auto)
& 'C:\Program Files\Dell\CommandUpdate\dcu-cli.exe' /applyUpdates -silent -reboot=enable

# Rapport
& 'C:\Program Files\Dell\CommandUpdate\dcu-cli.exe' /report
```

### BIOS Flash Update (F12, sans OS)

```
1. Télécharger le .exe du BIOS → l'exécuter UNE fois sur un PC Windows :
   il propose "créer un support de flash" OU extraire le fichier .rcv/.hdr.
   (Astuce : beaucoup de .exe Dell s'extraient avec /s /e=C:\temp\bios)
2. Copier le fichier sur une clé USB FAT32.
3. Sur la cible : F12 → BIOS Flash Update → sélectionner le fichier.
4. Ne pas interrompre. Le PC redémarre seul.
```

### BIOSConnect / SupportAssist (récupération)

Les Dell récents ont une **récupération BIOS automatique** : si le flash échoue, le système restaure depuis une copie de secours (voir §77).

---

## 54. Mise à jour BIOS/UEFI : HP

### Méthodes

| Méthode | Outil | Usage |
|---|---|---|
| Exécutable Windows | `spXXXXX.exe` (SoftPaq) | Un poste |
| **HP Image Assistant** / HP Manageability | Analyse du parc | Masse |
| Depuis le BIOS | F10 → Update BIOS (HP Sure Start) | Sans OS |
| Via USB recovery | Touche `Win + B` au démarrage | **Récupération** après flash raté |

### Exemple SoftPaq en silencieux

```powershell
# Télécharger le SoftPaq BIOS depuis support.hp.com (ex. sp123456.exe)
# Installation silencieuse avec flash :
Start-Process .\sp123456.exe -ArgumentList '/s /f' -Wait
# /s = silencieux, /f = forcer le flash (selon SoftPaq, vérifier la doc)
```

### HP Sure Start (protection)

Les HP Pro/Elite récents intègrent **Sure Start** : le BIOS est vérifié cryptographiquement à chaque démarrage, avec **copie dorée de restauration automatique**. En cas de corruption → restauration sans intervention. C'est une excellente protection contre les flashs ratés et les attaques firmware.

### Récupération d'urgence HP

```
Écran noir après flash :
1. Éteindre (10 s sur le bouton).
2. Maintenir Win + B (ou Win + V sur certains modèles).
3. Brancher le secteur TOUT en maintenant les touches.
4. Allumer : l'utilitaire de récupération BIOS HP démarre
   (chercher un BIOS sur USB ou en ligne).
```

---

## 55. Mise à jour BIOS/UEFI : Lenovo

### Méthodes

| Méthode | Outil | Usage |
|---|---|---|
| Exécutable Windows | `bios_update_utility.exe` | Un poste |
| **Lenovo System Update** | `tvsu.exe` | Parc ThinkPad/ThinkCentre |
| Depuis le BIOS | F12 / Enter → BIOS Update via USB | Sans OS |
| Via SCCM | Lenovo Patch / ThinkVantage | Masse |

### Lenovo System Update en CLI

```powershell
# Scanner
& 'C:\Program Files (x86)\Lenovo\System Update\tvsu.exe' /CM -search A -action INSTALL `
    -repository C:\temp\lsu -noreboot

# Installer les critiques (dont BIOS) sans interaction
& 'C:\Program Files (x86)\Lenovo\System Update\tvsu.exe' /CM -search C -action INSTALL -noreboot
# Codes : C = critiques, R = recommandées, A = toutes
```

### ThinkShield / self-healing

Les ThinkPad récents (ThinkShield) incluent une vérification d'intégrité du firmware avec restauration. Comme chez HP/Dell, la tendance est au **firmware auto-réparant** : vérifiez que cette option est activée dans le setup (Security → Secure Rollback Prevention, etc.).

### Particularité Lenovo : le mot de passe supervisor

Sans le mot de passe supervisor, **impossible de flasher** le BIOS sur beaucoup de ThinkPad. Prévoyez-le dans la procédure de masse (déverrouillage via WMI avec le mot de passe, §43).

---

## 56. Mise à jour du firmware TPM

