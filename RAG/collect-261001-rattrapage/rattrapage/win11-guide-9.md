---
id: collect-261001-rattrapage/rattrapage/win11-guide-9
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [1177, 1363]
sha256: 823509809fb5d0fc1b393d57b20305039edaaec3413a56ca8b53e2724792c0be
---

# Windows 11 en entreprise — Guide technique ultra-complet

### 17.2 Déploiement silencieux via GPO / script de démarrage

```powershell
<#
.SYNOPSIS
    Script de démarrage GPO : chiffre C: si ce n'est pas déjà fait,
    avec sauvegarde de la clé dans AD (via la GPO de récupération).
#>
$vol = Get-BitLockerVolume -MountPoint "C:"
if ($vol.VolumeStatus -eq 'FullyDecrypted') {
    # Vérifier que le TPM est prêt avant de chiffrer
    $tpm = Get-Tpm
    if ($tpm.TpmReady) {
        Enable-BitLocker -MountPoint "C:" -TpmProtector -UsedSpaceOnly -SkipHardwareTest
        # La GPO "Choisir comment..." se charge de la sauvegarde AD
        Write-EventLog -LogName Application -Source "BitLockerDeploy" `
            -EventId 1001 -EntryType Information `
            -Message "Chiffrement BitLocker démarré sur $env:COMPUTERNAME"
    } else {
        Write-EventLog -LogName Application -Source "BitLockerDeploy" `
            -EventId 1002 -EntryType Warning `
            -Message "TPM non prêt sur $env:COMPUTERNAME — chiffrement reporté"
    }
}
```

> ⚠️ Créez la source `BitLockerDeploy` une fois (`New-EventLog -LogName Application -Source BitLockerDeploy`) ou le script échouera sur `Write-EventLog`.

### 17.3 MBAM → BitLocker Management (MECM / Intune)

- **MBAM** (Microsoft BitLocker Administration and Monitoring) : historique, lié à MDT/SCCM. **Fin de support** — ne plus déployer.
- **BitLocker Management dans MECM** (à partir de la version 2002) : le successeur. Stratégies de chiffrement + portail de récupération en libre-service + rapports de conformité.
- **Intune** : stratégies de chiffrement de disque (Endpoint Security → Disk encryption), clés visibles dans le portail, conformité.

| Fonction | GPO seule | MECM BitLocker Mgmt | Intune |
|---|---|---|---|
| Imposer le chiffrement | ✅ | ✅ | ✅ |
| Stocker les clés | AD | Base MECM (+ AD) | Entra ID |
| Portail libre-service | ❌ | ✅ | Partiel (myaccount) |
| Rapports de conformité | Scripts maison | ✅ natif | ✅ natif |

### 17.4 Conformité : qui n'est pas chiffré ?

```powershell
# Rapport rapide : postes du domaine sans BitLocker actif
$computers = Get-ADComputer -Filter * -SearchBase "OU=Postes Win11,DC=entreprise,DC=local" |
             Select-Object -ExpandProperty Name
Invoke-Command -ComputerName $computers -ErrorAction SilentlyContinue -ScriptBlock {
    $v = Get-BitLockerVolume -MountPoint "C:" -ErrorAction SilentlyContinue
    [pscustomobject]@{
        Hostname         = $env:COMPUTERNAME
        ProtectionStatus = $v.ProtectionStatus   # On = chiffré et protégé
        VolumeStatus     = $v.VolumeStatus
    }
} | Where-Object { $_.ProtectionStatus -ne 'On' } |
  Export-Csv "C:\Admin\BitLocker_NonConformes.csv" -NoTypeInformation -Encoding UTF8
```

---

## 18. BitLocker : dépannage — écran de récupération, boucles, causes fréquentes

### 18.1 L'écran de récupération : que faire (procédure standard)

```
1. Noter l'ID DE CLÉ affiché (ex. : 4A2B1C3D) — il identifie QUELLE clé utiliser
2. Retrouver la clé dans AD / Intune / fichier (section 16)
3. Saisir les 48 chiffres (tirets automatiques)
4. Le poste démarre → NE PAS s'arrêter là :
5. Suspend-BitLocker -MountPoint "C:" -RebootCount 1
6. Redémarrer → si ça redemande la clé, chercher la CAUSE (ci-dessous)
7. Une fois stable : Resume-BitLocker
```

### 18.2 Causes fréquentes de demande de clé (par ordre de fréquence)

| Cause | Explication | Prévention |
|---|---|---|
| Mise à jour du BIOS/UEFI | Le TPM mesure un firmware différent | `Suspend-BitLocker` avant toute MàJ BIOS |
| Changement matériel (carte mère, disque) | Nouvelle mesure PCR | Suspendre avant intervention |
| Modification de l'ordre de boot / Secure Boot désactivé | PCR 7 modifié | Ne pas toucher au boot sans suspendre |
| Docking station / périphérique USB au boot | Modification de la chaîne de boot | Débrancher, ou configurer le BIOS |
| Mise à jour Windows majeure | Rare, mais possible | — |
| TPM effacé (clear TPM) | Les clés scellées sont perdues | Ne jamais effacer le TPM sans avoir la clé de récupération ! |

### 18.3 Boucle de récupération (demande la clé à chaque démarrage)

```powershell
# 1. Démarrer avec la clé de récupération, ouvrir PowerShell admin
# 2. Vérifier les protecteurs
Get-BitLockerVolume -MountPoint C: | Select-Object -ExpandProperty KeyProtector

# 3. Le cas classique : le protecteur TPM est "cassé" → le recréer
# D'abord suspendre, puis supprimer et recréer le protecteur TPM :
Suspend-BitLocker -MountPoint "C:" -RebootCount 0
$vol = Get-BitLockerVolume -MountPoint "C:"
$tpmProt = $vol.KeyProtector | Where-Object { $_.KeyProtectorType -eq 'Tpm' }
if ($tpmProt) {
    Remove-BitLockerKeyProtector -MountPoint "C:" -KeyProtectorId $tpmProt.KeyProtectorId
}
Add-BitLockerKeyProtector -MountPoint "C:" -TpmProtector
Resume-BitLocker -MountPoint "C:"
```

### 18.4 Réparer le démarrage chiffré depuis WinRE

```powershell
# Dans WinRE (invite de commandes) : déverrouiller le volume
manage-bde -unlock C: -RecoveryPassword 111111-222222-333333-444444-555555-666666-777777-888888
# Puis réparer le BCD :
bootrec /fixmbr
bootrec /fixboot   # peut échouer en UEFI → utiliser bcdboot
bootrec /rebuildbcd
bcdboot C:\Windows /s S: /f UEFI
```

### 18.5 TPM verrouillé (lockout)

Après trop de tentatives de PIN erronées, le TPM se verrouille (anti-bruteforce). Il se déverrouille seul après un délai (jusqu'à 24 h selon la configuration), ou via la clé de récupération.

```powershell
# Voir l'état de verrouillage
Get-Tpm | Select-Object LockedOut, LockoutCount, LockoutTimeRemaining
```

---

## 19. Profils utilisateurs : locaux, itinérants, obligatoires

### 19.1 Les 3 types de profils

| Type | Stockage | Usage moderne |
|---|---|---|
| Local | `C:\Users\<nom>` | Standard (avec OneDrive KFM pour la « mobilité ») |
| Itinérant (roaming) | Copié depuis/vers un partage à l'ouverture/fermeture | **Déconseillé** : lent, corruptible, incompatible avec les apps modernes |
| Obligatoire (mandatory) | `NTUSER.MAN` en lecture seule | Kiosques, salles de formation |

**Doctrine 2026** : profil local + **OneDrive Known Folder Move** (Bureau/Documents/Images synchronisés) + **FSLogix** pour les environnements virtualisés. Les profils itinérants classiques sont une dette technique.

### 19.2 Anatomie d'un profil

```
C:\Users\jdupont\
 ├─ NTUSER.DAT            ← ruche de registre de l'utilisateur
 ├─ AppData\Local         ← données locales (non itinérantes)
 ├─ AppData\LocalLow
 ├─ AppData\Roaming       ← données "itinérantes"
 ├─ Desktop / Documents / Downloads / ...
```

```powershell
# Lister les profils présents sur un poste
Get-CimInstance Win32_UserProfile |
    Where-Object { -not $_.Special } |
    Select-Object LocalPath, Loaded, LastUseTime, @{N='SizeGB';E={
        [math]::Round((Get-ChildItem $_.LocalPath -Recurse -Force -ErrorAction SilentlyContinue |
            Measure-Object Length -Sum).Sum / 1GB, 2)
    }} | Sort-Object SizeGB -Descending
```

### 19.3 Supprimer proprement un profil

```powershell
# Méthode propre (via WMI — nettoie aussi le registre)
$profile = Get-CimInstance Win32_UserProfile -Filter "LocalPath='C:\\Users\\ancien.user'"
Remove-CimInstance -InputObject $profile

# ⛔ Ne JAMAIS supprimer juste le dossier C:\Users\... sans nettoyer
# HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\ProfileList :
# → profil "fantôme" / temporaire au prochain logon (voir cas pratique 8).
```

### 19.4 Profil obligatoire (mandatory) — pour kiosques

```powershell
# 1. Personnaliser un profil modèle en mode audit
# 2. Renommer NTUSER.DAT en NTUSER.MAN dans le profil copié
Rename-Item "D:\Profils\kiosque\NTUSER.DAT" "NTUSER.MAN"
# 3. Le copier sur un partage, le référencer dans l'AD (onglet Profil de l'utilisateur)
```

---

## 20. FSLogix : principe, conteneurs de profils, cas d'usage

### 20.1 Pourquoi FSLogix

