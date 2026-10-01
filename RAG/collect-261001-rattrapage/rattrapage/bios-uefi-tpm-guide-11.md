---
id: collect-261001-rattrapage/rattrapage/bios-uefi-tpm-guide-11
title: "BIOS / UEFI — Secure Boot — TPM 2.0"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/bios_uefi_tpm_guide.md
source_anchor: ""
source_lines: [1543, 1730]
sha256: e4fdcdae4dc219c1638fea6d0f4626e72d4d7dec304f7a2a428ec0c7a6d7f8b6
---

# BIOS / UEFI — Secure Boot — TPM 2.0

```powershell
# Via tpm.msc → pas d'affichage direct des PCR.
# Via PowerShell (nécessite le module TPM) : pas de cmdlet natif simple.
# Méthode : journal des mesures dans l'observateur d'événements
Get-WinEvent -LogName "Microsoft-Windows-TPM/Operational" -MaxEvents 20 |
    Select-Object TimeCreated, Id, Message | Format-Table -Wrap
```

> 💡 **Lien direct avec la section 37 :** c'est parce que BitLocker scelle sa clé sur **PCR 0 (firmware)** que toute mise à jour du BIOS invalide le scellé et déclenche la demande de clé de récupération.

---

## 36. BitLocker et TPM : les protecteurs

BitLocker chiffre le volume ; les **protecteurs** (*key protectors*) contrôlent *comment* la clé de chiffrement (FVEK) est déverrouillée au démarrage.

### Les protecteurs principaux

| Protecteur | Commande | Sécurité | Ergonomie |
|---|---|---|---|
| **TPM seul** | `Enable-BitLocker C: -TpmProtector` | Bonne (lié au matériel + PCR) | ✅ Transparent (pas de saisie) |
| **TPM + PIN** | `Add-BitLockerKeyProtector C: -TpmAndPinProtector` | **Excellente** (double facteur) | ⚠️ PIN à chaque démarrage |
| **TPM + clé USB** | `-TpmAndStartupKeyProtector` | Excellente | ⚠️ Clé USB requise |
| **Mot de passe** | `-PasswordProtector` | Moyenne (attaques dico) | Saisie à chaque démarrage |
| **Clé de récupération** | `-RecoveryPasswordProtector` | Secours (48 chiffres) | À conserver hors machine |

### Recommandations par profil

| Profil de poste | Protecteur recommandé |
|---|---|
| Bureautique standard | **TPM seul** (+ clé de récupération séquestrée) |
| Portable / nomade | **TPM + PIN** (6+ chiffres, anti-vol) |
| Direction / finance / sensible | **TPM + PIN** (+ interdiction du mode veille simple → hibernation) |
| Serveur en datacenter | TPM seul (redémarrage sans intervention) |

### Commandes essentielles

```powershell
# Activer BitLocker (TPM seul) avec clé de récupération
Enable-BitLocker -MountPoint "C:" -TpmProtector -RecoveryPasswordProtector

# Ajouter un PIN (TPM+PIN) — demande le PIN de manière interactive
Add-BitLockerKeyProtector -MountPoint "C:" -TpmAndPinProtector

# Lister les protecteurs
(Get-BitLockerVolume -MountPoint "C:").KeyProtector |
    Select-Object KeyProtectorType, KeyProtectorId

# Sauvegarder la clé de récupération dans l'AD
Backup-BitLockerKeyProtector -MountPoint "C:" -KeyProtectorId '{GUID-de-la-recovery}'

# État du chiffrement
Get-BitLockerVolume -MountPoint "C:" |
    Select-Object MountPoint, VolumeStatus, EncryptionPercentage, KeyProtector
```

### Choisir l'algorithme de chiffrement (GPO)

```
Stratégie : Configuration ordinateur → Modèles d'administration →
            Composants Windows → BitLocker →
            "Choisir la méthode de chiffrement de lecteur"
Recommandé : XTS-AES 256 bits (disques fixes et OS, Windows 10 1511+).
```

---

## 37. Pourquoi une MAJ du BIOS déclenche la récupération BitLocker

### Le mécanisme (rappel PCR)

```
1. BitLocker scelle la clé (FVEK) : "ne te déchiffre que si
   PCR[0,2,4,11] == valeurs enregistrées au moment du chiffrement".
2. Mise à jour du BIOS → le code du firmware change
   → au prochain boot, PCR 0 = NOUVELLE valeur.
3. Le TPM refuse de desceller : les PCR ne correspondent plus.
4. BitLocker bascule en mode récupération → demande la clé à 48 chiffres.
```

### Ce n'est PAS un bug, c'est une protection

Si un attaquant remplaçait votre BIOS par une version piégée, **vous voudriez** que BitLocker refuse de déchiffrer silencieusement. La récupération est le comportement de sécurité attendu.

### Procédure de MAJ BIOS sur poste chiffré (sans se faire piéger)

```powershell
# AVANT le flash :
# 1. Vérifier que la clé de récupération est séquestrée (AD/Entra ID)
manage-bde -protectors -get C: -type RecoveryPassword

# 2. Suspendre BitLocker (1 redémarrage suffit pour un flash)
Suspend-BitLocker -MountPoint "C:" -RebootCount 1
# → le scellé PCR est temporairement contourné, pas le chiffrement :
#    les données restent chiffrées, seul le contrôle d'intégrité est levé.

# 3. Flasher le BIOS (section 52-55).

# 4. APRÈS le flash : BitLocker re-scelle automatiquement sur les
#    nouvelles valeurs PCR au redémarrage suivant.
#    Vérifier :
Get-BitLockerVolume -MountPoint "C:" | Select-Object VolumeStatus
# VolumeStatus doit revenir à "FullyEncrypted" (pas "EncryptionSuspended").
```

### Si la récupération se déclenche quand même

```
1. Récupérer la clé à 48 chiffres : AD (Attribut msFVE-RecoveryPassword),
   Entra ID (fiche appareil), ou copie papier/USB de l'utilisateur.
2. La saisir à l'invite bleue BitLocker.
3. Une fois dans Windows : Suspend-BitLocker puis Resume-BitLocker
   pour re-sceller sur les nouveaux PCR.
```

> 🔴 **Erreur classique n°1 en entreprise :** flasher 200 BIOS un vendredi soir sans suspendre BitLocker et sans vérifier le séquestre des clés → 200 utilisateurs bloqués lundi matin. **Toujours** : séquestre vérifié + suspension avant flash.

---

## 38. Gestion des clés de récupération BitLocker (AD / Entra ID)

Sans séquestre centralisé, une clé de récupération perdue = données perdues. Trois modèles :

### Modèle 1 : Active Directory local (GPO)

```
GPO : Configuration ordinateur → Stratégies → Modèles d'administration →
      Composants Windows → BitLocker → Lecteurs du système d'exploitation →
      "Choisir comment les lecteurs protégés par BitLocker peuvent être récupérés"
  ✅ Stocker les informations de récupération dans AD DS
  ✅ Exiger la sauvegarde AD avant activation (bloque si échec)
```

Attributs AD utilisés : `msFVE-RecoveryPassword`, `msFVE-KeyPackage`, `msFVE-VolumeGuid` (sur l'objet ordinateur).

```powershell
# Retrouver la clé de récupération d'un poste dans l'AD
Get-ADObject -Filter "objectClass -eq 'msFVE-RecoveryInformation'" `
    -SearchBase (Get-ADComputer PC-0427).DistinguishedName `
    -Properties msFVE-RecoveryPassword |
    Select-Object Name, msFVE-RecoveryPassword
```

### Modèle 2 : Microsoft Entra ID (Intune / AAD Join)

- La clé est sauvegardée automatiquement si la stratégie Intune l'exige.
- Consultation : **Centre d'administration Entra → Appareils → [appareil] → Clés de récupération**.
- En PowerShell (module Microsoft.Graph) :

```powershell
# Lister les clés BitLocker d'un appareil Entra
Connect-MgGraph -Scopes BitLockerKey.Read.All
$device = Get-MgDevice -Filter "displayName eq 'PC-0427'"
Get-MgInformationProtectionBitlockerRecoveryKey -Filter "deviceId eq '$($device.DeviceId)'" |
    Select-Object Id, CreatedDateTime
# Puis récupérer la clé elle-même :
Get-MgInformationProtectionBitlockerRecoveryKey -BitlockerRecoveryKeyId '<id>' |
    Select-Object -ExpandProperty Key
```

### Modèle 3 : fichier / papier (petites structures)

```powershell
# Exporter la clé dans un fichier (à stocker CHIFFRÉ, hors du poste)
(Get-BitLockerVolume -MountPoint "C:").KeyProtector |
    Where-Object KeyProtectorType -eq RecoveryPassword |
    Select-Object KeyProtectorId, RecoveryPassword |
    Export-Csv C:\sequestre\cles-bitlocker.csv -NoTypeInformation
```

> 🔒 **La clé de récupération est aussi sensible que les données** : stockage chiffré, accès restreint (groupe dédié), journalisation des consultations.

---

## 39. Windows 11 : exigences TPM 2.0 et Secure Boot

### Exigences matérielles officielles (disque système)

| Exigence | Détail |
|---|---|
| **TPM 2.0** | Activé et provisionné (fTPM/PTT accepté) |
| **Secure Boot** | Compatible et **activé** |
| **UEFI** | Mode natif (CSM désactivé) |
| **GPT** | Style de partition du disque système |
| CPU / RAM / stockage | Selon liste Microsoft (CPU 1 GHz+ 2 cœurs 64 bits, 4 Go RAM, 64 Go disque) |

### Vérifier la compatibilité (méthodes officielles)

