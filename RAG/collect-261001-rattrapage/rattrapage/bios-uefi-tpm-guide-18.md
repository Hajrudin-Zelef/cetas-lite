---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-18
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["acquisition", "arr", "gpu", "intel"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [2916, 3132]
sha256: ffd743b049030b7008a06a8cdf7c634bc97e02ba53aef282bb6c743ac6a85cbc
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```
Le firmware, avec Secure Boot activé, vérifie aussi la signature des
ROMs d'extension (Option ROM) des cartes PCIe.
→ La ROM de cette carte n'est pas signée par une clé de db
  (carte exotique, vieux firmware de carte, ou contrefaçon).
→ Secure Boot refuse d'exécuter la ROM → bloque le boot.
```

### Résolution (par ordre de préférence)

```
Option A (recommandée) : mettre à jour le firmware de LA CARTE
  → les firmwares récents des cartes pro sont signés (clé "Microsoft
    Option ROM UEFI CA" dans db). Télécharger le firmware chez le
    fabricant de la carte, flasher, retester.

Option B : désactiver le boot sur cette carte
  → setup UEFI → désactiver l'Option ROM de ce slot PCIe
    (si la carte ne sert pas au boot : ex. GPU de calcul, carte
    d'acquisition). Le système démarrera, la carte sera vue par l'OS.

Option C (temporaire, tracée) : désactiver Secure Boot
  → uniquement pour valider que c'est bien la cause, puis réactiver
    après application de l'option A ou B.

Option D (avancée) : mode Custom + enrôlement de la clé du fabricant
  → si le fabricant fournit son certificat, l'importer en db.
```

### Vérification

```powershell
# Après résolution
Confirm-SecureBootUEFI   # → True (on n'a PAS laissé Secure Boot désactivé)
# Tester 2-3 redémarrages à froid + un arrêt prolongé.
```

> 💡 **Leçon :** intégrer la vérification « Secure Boot + nouvelle carte » dans la procédure d'ajout matériel. Un pilote Windows signé ne suffit pas : c'est la **ROM de la carte** qui est vérifiée au boot.

---

## 66. Cas pratique 3 : boucle de démarrage (boot loop) après une mise à jour

### Contexte

Après une mise à jour Windows (ou un flash BIOS), le PC redémarre en boucle : logo constructeur → écran bleu fugace → reboot. Impossible d'atteindre Windows.

### Diagnostic

```
1. Interrompre la boucle : forcer 3 extinctions pendant le démarrage
   (bouton 10 s) → au 3e, Windows lance WinRE automatiquement.
   (Alternative : clé USB d'installation → "Réparer l'ordinateur".)
2. En WinRE : Dépannage → Options avancées → voir les options :
   - "Désinstaller les mises à jour" (si MAJ Windows en cause).
   - "Restauration du système" (si point de restauration).
   - "Invite de commandes" (diagnostic manuel).
```

### Résolution en invite WinRE

```powershell
# Identifier les lettres (en WinRE, C: n'est pas toujours Windows !)
diskpart
#   list volume   → noter la lettre du volume Windows (ex. D:) et de l'ESP
#   exit

# Cas A : BCD endommagé par la MAJ
bcdboot D:\Windows /s S: /f UEFI
#   → "Boot files successfully created."

# Cas B : pilote fautif (écran bleu au boot)
#   Démarrer en mode sans échec :
bcdedit /store S:\EFI\Microsoft\Boot\BCD /set '{default}' safeboot Minimal
#   → reboot, désinstaller le pilote/MAJ, puis :
bcdedit /store S:\EFI\Microsoft\Boot\BCD /deletevalue '{default}' safeboot

# Cas C : MAJ Windows fautive → la désinstaller
dism /image:D:\ /get-packages | findstr "Package_for"
dism /image:D:\ /remove-package /packagename:Package_for_RollupFix~...

# Cas D : restauration système
rstrui.exe
```

### Si c'est le flash BIOS qui a causé la boucle

```
→ Le firmware a peut-être réinitialisé des réglages (SATA RAID→AHCI,
  Secure Boot, ordre de boot). Vérifier dans le setup :
  - Mode SATA identique à l'installation d'origine (sinon BSOD INACCESSIBLE_BOOT_DEVICE).
  - Ordre de boot correct.
→ BitLocker peut demander la clé (PCR 0 changé) : l'avoir sous la main (§37).
```

---

## 67. Cas pratique 4 : BCD corrompu — écran bleu 0xc000000f

### Contexte

Message au démarrage :

```
Recovery — Your PC needs to be repaired
The Boot Configuration Data for your PC is missing or contains errors.
File: \EFI\Microsoft\Boot\BCD
Error code: 0xc000000f
```

Causes : coupure pendant une MAJ, clonage raté, ESP effacée, disque remplacé.

### Résolution complète (WinRE)

```powershell
# 1. Démarrer sur clé USB Windows → Réparer → Invite de commandes.

# 2. Identifier les volumes
diskpart
#   list disk        → repérer le disque système
#   select disk 0
#   list partition   → repérer l'ESP (FAT32, ~100 Mo) et Windows (NTFS)
#   select partition 1   (l'ESP)
#   assign letter=S
#   list volume      → noter la lettre de Windows (ex. C:)
#   exit

# 3. Reconstruire le boot
bcdboot C:\Windows /s S: /f UEFI /l fr-FR

# 4. Vérifier
bcdedit /store S:\EFI\Microsoft\Boot\BCD /enum
#   → doit afficher Windows Boot Manager + une entrée Windows.

# 5. Nettoyer et redémarrer
#   diskpart → select partition 1 → remove letter=S → exit
#   → "Continuer" (redémarrer sur le disque).
```

### Si bcdboot échoue (« Failure when attempting to copy boot files »)

```
Cause n°1 : mauvaise lettre Windows (en WinRE, vérifier avec dir C:\Windows).
Cause n°2 : ESP non en FAT32 ou sans lettre → reformater PROPREMENT :
  diskpart → select partition 1 → format fs=fat32 quick → assign letter=S
  (⚠️ ne formater QUE si on est sûr que c'est l'ESP !)
Cause n°3 : disque en lecture seule → attributes disk clear readonly.
```

---

## 68. Cas pratique 5 : PXE-E53 / PXE-M0F — le client ne trouve pas le serveur de boot

### Contexte

Un poste en boot PXE (F12 → Network) affiche `PXE-E53: No boot filename received` puis `PXE-M0F: Exiting Intel PXE ROM`. Le déploiement est bloqué.

### Arbre de diagnostic (dans l'ordre, du plus probable)

```
① Le client a-t-il une IP ?
   → Non (ou 169.254.x.x) : problème DHCP de base.
     Vérifier : câble, port switch (VLAN correct ?), serveur DHCP joignable,
     relais DHCP (ip helper) si autre sous-réseau.
② Le client a une IP mais pas de "boot filename" ?
   → Le serveur WDS ne répond pas :
     a) Service WDSServer démarré ? (Get-Service WDSServer)
     b) Pare-feu serveur : UDP 67/68 (DHCP), UDP 69 (TFTP), UDP 4011 (PXE) ?
        → Test rapide : désactiver temporairement le pare-feu pour valider.
     c) DHCP et WDS sur le même serveur ? → options "Ne pas écouter le
        port 67" + "DHCP option 60" cochées ? (§50)
     d) Stratégie de réponse WDS : "Répondre aux clients connus" alors que
        la machine est inconnue ? → passer en "Répondre à tous" ou
        pré-créer le compte ordinateur (wdsutil /Add-Device).
③ TFTP timeout (PXE-E32) après l'IP ?
   → Le fichier de boot est-il présent ?
     E:\RemoteInstall\Boot\x64\wdsmgfw.efi (UEFI)
     → Régénérer : redémarrer le service WDS, vérifier les droits NTFS.
④ Le client est en UEFI mais le DHCP pousse wdsnbp.com (Legacy) ?
   → Corriger l'option 67 : boot\x64\wdsmgfw.efi pour l'UEFI.
```

### Commandes serveur

```powershell
# État du service et des ports
Get-Service WDSServer
netstat -ano | Select-String ':69|:4011'

# Journaux WDS (les 20 dernières erreurs)
Get-WinEvent -LogName 'Microsoft-Windows-Deployment-Services-Diagnostics/Operational' `
    -MaxEvents 20 | Where-Object LevelDisplayName -eq 'Erreur' |
    Select-Object TimeCreated, Id, Message | Format-List

# Tester le TFTP depuis un poste (client TFTP Windows à activer)
tftp <ip-wds> GET boot\x64\wdsmgfw.efi C:\temp\test.efi
# → si ça échoue : pare-feu ou service.
```

---

## 69. Cas pratique 6 : mot de passe BIOS oublié

### Contexte

Personne ne connaît le mot de passe admin du firmware. Impossible de changer l'ordre de boot ou d'activer le TPM. (Situation fréquente sur du matériel récupéré ou après le départ d'un technicien.)

### Ce qui NE marche PAS (ou plus)

| Méthode | Statut |
|---|---|
| « Master password » universel trouvé sur internet | ❌ Arnaque ou obsolète ; les algos par constructeur changent |
| Enlever la pile CMOS | ❌ Inefficace sur 99 % des machines modernes (mot de passe en NVRAM/flash, pas en CMOS) |
| Court-circuiter la puce au hasard | ❌ Dangereux, peut détruire la carte mère |

### Procédures par constructeur (voies officielles)

