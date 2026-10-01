---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-16
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "attention", "gpu", "incident", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [2548, 2709]
sha256: 0ebea3b0c2557e9b7c6b731a7cad41fa4a07fc047d5e852738fab452af844c7b
---

# BIOS / UEFI — Secure Boot — TPM 2.0

Le TPM lui-même a un firmware, avec des failles publiées (ex. Infineon RSA 2017, CVE-2017-15361). Le mettre à jour, c'est patcher le coffre-fort.

### Quand mettre à jour

- [ ] Bulletin de sécurité du fabricant du TPM (Infineon, STMicro) ou du constructeur.
- [ ] Outil de détection : `Get-Tpm | Select-Object ManufacturerVersion` → comparer avec la version corrigée publiée.

### Procédure type (exemple Dell/Infineon)

```powershell
# 1. Séquestre BitLocker VÉRIFIÉ + suspension (la MAJ TPM = Clear implicite !)
Suspend-BitLocker -MountPoint "C:" -RebootCount 2
manage-bde -protectors -get C: -type RecoveryPassword

# 2. Lancer l'utilitaire du constructeur (ex. Dell TPM 2.0 Firmware Update)
Start-Process .\TPM_Firmware_Update.exe -ArgumentList '/s' -Wait

# 3. Le PC redémarre, met à jour le TPM (ne pas interrompre !),
#    puis Windows re-provisionne automatiquement (AutoProvisioning).

# 4. Vérifications
Get-Tpm | Select-Object TpmPresent, TpmReady, ManufacturerVersion
Get-BitLockerVolume -MountPoint "C:" | Select-Object VolumeStatus
# → FullyEncrypted, pas "suspendu"
```

> 🔴 **Une MAJ du firmware TPM efface le TPM** (comme un Clear). Sans la clé de récupération BitLocker sous la main, c'est la perte des données. C'est l'opération la plus piégeuse du chapitre : elle ressemble à une simple MAJ, elle agit comme un Clear.

---

## 57. Attaques : les bootkits

Un **bootkit** est un malware qui s'installe **avant l'OS**, dans la chaîne de démarrage : MBR, chargeur, firmware UEFI, ou ROM d'extension. Il est invisible pour l'antivirus classique (qui démarre après lui).

### Vecteurs d'infection

| Vecteur | Cible | Contre-mesure |
|---|---|---|
| MBR infecté (ex. familles historiques) | Secteur 0 du disque | UEFI + Secure Boot (le MBR n'est plus exécuté) |
| Chargeur EFI remplacé | `\EFI\...\bootmgfw.efi` | **Secure Boot** (signature vérifiée) |
| Pilote malveillant au boot | Drivers de démarrage | Secure Boot + ELAM + HVCI |
| **Firmware UEFI infecté** | Image du BIOS elle-même | Secure Boot + **Boot Guard** (Intel) / PSB (AMD) + MAJ firmware |
| ROM d'extension (carte réseau, GPU) | Option ROM | Secure Boot vérifie aussi les ROMs (si activé) |

### Exemples notables (culture sécurité)

- **BlackLotus (2023)** : contournement de Secure Boot via une faille dans un chargeur signé (CVE-2023-24932) → Microsoft a dû révoquer via dbx + patch du boot. Démonstration que Secure Boot n'est pas invincible : il faut **tenir la dbx à jour** (§21).
- **CosmicStrand / MoonBounce** : implants dans le firmware UEFI (persistants même après formatage du disque !).
- **Bootkits MBR historiques** (TDL4, Olmasco) : rendus obsolètes par l'UEFI.

### Détection

```
- Secure Boot activé + dbx à jour = la base (bloque 95 % des cas).
- Comparer les hashes des fichiers EFI avec une machine de référence saine.
- Outils EDR avec protection firmware (certains EDR d'entreprise scannent
  l'image UEFI).
- Flash d'un firmware officiel connu-sain en cas de doute sérieux
  (un reflash écrase un implant firmware dans la plupart des cas).
```

> 💡 **Message clé pour la direction :** un bootkit firmware survit au formatage et à la réinstallation. La prévention (Secure Boot, MAJ firmware, contrôle physique) coûte 100× moins cher que la remédiation.

---

## 58. Attaques : evil maid et accès physique

L'**evil maid** (« la femme de chambre malveillante ») : un attaquant avec un **accès physique temporaire** au poste (hôtel, bureau partagé, maintenance) compromet la machine en quelques minutes.

### Scénarios

| Scénario | Déroulé | Contre-mesure |
|---|---|---|
| Clé USB bootable | Boot sur Linux live → copie/altération du disque | Mot de passe UEFI + USB boot désactivé + **BitLocker** |
| Modification du chargeur | Remplacement de `bootmgfw.efi` par une version piégée | **Secure Boot** (signature refusée) |
| Reset du mot de passe Windows | Utilitaire offline sur la base SAM | BitLocker (disque illisible hors OS) |
| Keylogger matériel | Dongle USB entre clavier et PC | Inspection physique, ports verrouillés |
| Attaque DMA (Thunderbolt/PCIe) | Lecture mémoire via DMA avant verrouillage | VT-d + **Kernel DMA Protection**, désactiver le pre-boot DMA |
| Cold boot | Refroidir la RAM pour en extraire les clés | BitLocker + extinction complète (pas de veille) sur postes sensibles |

### La défense en profondeur du poste

```
1. Chiffrement : BitLocker TPM+PIN (le disque seul ne vaut rien).
2. Firmware : mot de passe admin, USB boot désactivé, Secure Boot.
3. Physique : câble antivol, ports condamnés si besoin, BIOS verrouillé.
4. Comportement : extinction (pas veille) pour les postes sensibles ;
   méfiance sur les postes laissés sans surveillance en déplacement.
5. Détection : BitLocker en récupération "sans raison" = signal d'alerte
   (quelqu'un a touché au firmware ou au boot !).
```

> 🔴 **Point d'attention manager :** une demande de clé de récupération BitLocker « spontanée » sur un portable de direction n'est pas qu'un incident support : c'est potentiellement un **indicateur de compromission physique**. Procédure : vérifier la cause (MAJ ? changement matériel ?) avant de simplement débloquer.

---

## 59. Durcissement d'un poste : la check-list sécurité

### Niveau 1 — Base (tout le parc)

- [ ] UEFI natif, CSM désactivé
- [ ] Secure Boot activé (vérifié : `Confirm-SecureBootUEFI` = True)
- [ ] TPM 2.0 activé et provisionné (`Get-Tpm` → TpmReady)
- [ ] Disque système GPT
- [ ] BitLocker activé, clé séquestrée (AD/Entra ID)
- [ ] Mot de passe admin UEFI défini (coffre d'équipe)
- [ ] USB boot désactivé (sauf exceptions tracées)
- [ ] Ordre de boot : disque en premier
- [ ] Firmware à jour (moins de 12 mois)

### Niveau 2 — Renforcé (portables, nomades)

- [ ] Tout le niveau 1, plus :
- [ ] BitLocker **TPM + PIN** (6 chiffres minimum)
- [ ] Mot de passe système au boot (si compatible avec le WoL/patching)
- [ ] Kernel DMA Protection activé (vérifier : `msinfo32` → « Protection DMA du noyau »)
- [ ] HVCI / Memory Integrity activé (Sécurité Windows → Isolation du noyau)
- [ ] Extinction exigée (désactiver la veille simple via GPO sur profils sensibles)

### Niveau 3 — Sensible (direction, finance, admin)

- [ ] Tout le niveau 2, plus :
- [ ] TPM discret si disponible
- [ ] Ports USB avant désactivés dans le firmware
- [ ] Windows Hello for Business (clés dans le TPM)
- [ ] Credential Guard activé
- [ ] Inventaire firmware mensuel (script §40)

### Vérification automatisée (conformité)

```powershell
#requires -RunAsAdministrator
function Test-PostureSecurite {
    $tpm = Get-Tpm
    $bde = Get-BitLockerVolume -MountPoint $env:SystemDrive
    $sb  = $false; try { $sb = Confirm-SecureBootUEFI } catch {}
    $dma = (Get-CimInstance -ClassName Win32_DeviceGuard |
        Select-Object -ExpandProperty KernelDmaProtectionStatus -ErrorAction SilentlyContinue)
    [PSCustomObject]@{
        SecureBoot      = $sb
        TPM_Pret        = $tpm.TpmReady
        BitLocker       = $bde.VolumeStatus
        Protecteur_TPM  = ($bde.KeyProtector.KeyProtectorType -contains 'Tpm')
        DMA_Kernel      = $dma
        Score           = (($sb) + ($tpm.TpmReady) + ($bde.VolumeStatus -eq 'FullyEncrypted'))
    }
}
Test-PostureSecurite | Format-List
```

---

## 60. Bonnes pratiques de gestion de parc

### Standardisation

- [ ] **2–3 modèles** de postes maximum par génération (réduit les images, les pilotes, les procédures firmware).
- [ ] **Image master** unique par modèle (avec firmware de référence inclus dans la doc).
- [ ] **Profils BIOS standardisés** : exporter la config d'un poste modèle (Dell `cctk --backup`, HP `BiosConfigUtility /get`, Lenovo `wmi`) et la déployer.

### Cycle de vie du firmware

