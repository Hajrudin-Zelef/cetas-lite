---
id: collect-261001-rattrapage/rattrapage/win11-guide-20
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [3306, 3494]
sha256: 188cb51a6e7606badc459510dd62905ce08b5260003bca5566d92bb50ea898e2
---

# Windows 11 en entreprise — Guide technique ultra-complet

```powershell
<#
.SYNOPSIS
    Audit express : affiche l'état de santé/sécurité d'un poste Windows 11.
#>
Write-Host "=== AUDIT $($env:COMPUTERNAME) ===" -ForegroundColor Cyan
$ver = Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion"
Write-Host "OS      : $($ver.ProductName) $($ver.DisplayVersion) (build $($ver.CurrentBuild).$($ver.UBR))"
Write-Host "Installé: $((Get-CimInstance Win32_OperatingSystem).InstallDate)"
Write-Host "Dernier boot : $((Get-CimInstance Win32_OperatingSystem).LastBootUpTime)"
Write-Host ""
Write-Host "--- Sécurité ---" -ForegroundColor Yellow
Write-Host "TPM 2.0 prêt : $((Get-Tpm).TpmReady)"
Write-Host "Secure Boot : $(try { Confirm-SecureBootUEFI } catch { 'N/A' })"
$bl = Get-BitLockerVolume -MountPoint "C:"
Write-Host "BitLocker   : $($bl.ProtectionStatus) ($($bl.VolumeStatus), $($bl.EncryptionPercentage)%)"
$mp = Get-MpComputerStatus
Write-Host "Defender    : AV=$($mp.AntivirusEnabled) TempsRéel=$($mp.RealTimeProtectionEnabled)"
Write-Host "Signatures  : $($mp.AntivirusSignatureLastUpdated)"
Write-Host "LAPS (stratégie) : $((Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\LAPS' -ErrorAction SilentlyContinue) -ne $null)"
Write-Host ""
Write-Host "--- Réseau ---" -ForegroundColor Yellow
Get-NetIPAddress -AddressFamily IPv4 | Where-Object { $_.InterfaceAlias -notlike "*Loopback*" } |
    ForEach-Object { Write-Host "$($_.InterfaceAlias) : $($_.IPAddress)" }
Write-Host "Domaine/Entra : $(dsregcmd /status | Select-String 'AzureAdJoined|DomainJoined')"
Write-Host ""
Write-Host "--- Disque ---" -ForegroundColor Yellow
Get-PSDrive C | Select-Object @{N='Libre(Go)';E={[math]::Round($_.Free/1GB,1)}},
    @{N='Utilisé(Go)';E={[math]::Round($_.Used/1GB,1)}} | Format-Table -AutoSize
```

### 61.2 Redémarrage planifié + notification (maintenance)

```powershell
# Redémarrer un parc à heure fixe avec avertissement
$computers = Get-Content C:\Admin\postes-vague2.txt
foreach ($pc in $computers) {
    shutdown /m \\$pc /r /t 600 /c "Maintenance planifiée : redémarrage dans 10 minutes. Sauvegardez votre travail."
}
# Annuler si besoin :
# shutdown /m \\$pc /a
```

### 61.3 Nettoyage à distance (disque plein)

```powershell
# Nettoyer les profils inutilisés depuis plus de 90 jours (à distance)
Invoke-Command -ComputerName "PC-COMPTA-042" -ScriptBlock {
    Get-CimInstance Win32_UserProfile | Where-Object {
        -not $_.Special -and $_.LastUseTime -lt (Get-Date).AddDays(-90) -and -not $_.Loaded
    } | ForEach-Object {
        Write-Output "Suppression : $($_.LocalPath)"
        Remove-CimInstance -InputObject $_
    }
    # Vider les temporaires Windows
    Remove-Item "C:\Windows\Temp\*" -Recurse -Force -ErrorAction SilentlyContinue
}
```

### 61.4 Forcer l'inventaire / la remontée Intune

```powershell
# Forcer la synchronisation Intune (côté client)
# Via le planificateur : tâche "PushLaunch" du fournisseur MDM, ou :
Get-ScheduledTask -TaskName "PushLaunch" -TaskPath "\Microsoft\Windows\EnterpriseMgmt\*" |
    Start-ScheduledTask
```

---

## 62. Sécurité avancée : Secure Boot, measured boot, attestation

### 62.1 Chaîne de confiance au démarrage

```
UEFI Secure Boot  →  vérifie la signature du bootloader
   → Measured Boot →  le TPM "mesure" chaque étape (hash dans les PCR)
      → BitLocker  →  ne déchiffre que si les mesures sont conformes
         → HVCI    →  le noyau ne charge que du code signé
            → ELAM →  l'antimalware démarre en premier
```

### 62.2 Vérifier l'état de la chaîne

```powershell
# Secure Boot + VBS (rappel section 32)
Confirm-SecureBootUEFI
Get-CimInstance -ClassName Win32_DeviceGuard -Namespace root\Microsoft\Windows\DeviceGuard

# Journal du démarrage mesuré (TPM)
Get-WinEvent -FilterHashtable @{LogName='Microsoft-Windows-TPM-WMI/Operational'; StartTime=(Get-Date).AddDays(-1)} |
    Select-Object -First 5 TimeCreated, Id, Message
```

### 62.3 Attestation (flottes critiques)

L'**attestation d'intégrité** (Microsoft Intune / Azure Attestation) permet de vérifier à distance que le démarrage d'un poste est sain avant de lui donner accès (accès conditionnel : « exiger un appareil sain »).

```powershell
# Vérifier que l'attestation est configurée (registre, poste Entra join)
Get-ItemProperty "HKLM:\SYSTEM\CurrentControlSet\Control\DeviceGuard" -Name "Attestation" -ErrorAction SilentlyContinue
```

> 💡 Pour un chef de service : l'attestation + accès conditionnel = « les postes non sains n'accèdent pas aux données sensibles ». C'est le niveau au-dessus du simple antivirus.

---

## 63. Chiffrement avancé : BitLocker réseau (Network Unlock), DRA

### 63.1 BitLocker Network Unlock

Permet le **déchiffrement automatique au boot** pour les postes **filaire sur le LAN d'entreprise** (sans PIN), tout en exigeant le PIN hors réseau.

Prérequis :

- Serveurs WDS avec le rôle Network Unlock ;
- Certificat (clé publique) déployé sur les clients via GPO ;
- DHCP avec option 066/067 vers le WDS ;
- Postes avec protecteur **TPM+PIN** (le Network Unlock ajoute un protecteur réseau).

```powershell
# Côté client : ajouter le protecteur Network Unlock (après config serveur)
Add-BitLockerKeyProtector -MountPoint "C:" -NetworkUnlockProtector
```

### 63.2 Agent de récupération de données (DRA)

Le DRA est un certificat dont la clé privée permet de déchiffrer **n'importe quel volume** de l'organisation : c'est le « passe-partout » ultime.

```powershell
# Créer un DRA (sur une machine isolée, à conserver HORS LIGNE) :
cipher /R:C:\Admin\DRA-Entreprise
# → génère DRA-Entreprise.cer (public, à déployer par GPO) et DRA-Entreprise.pfx (PRIVÉ !)

# GPO : Composants Windows\BitLocker\Lecteurs du système d'exploitation
# → "Choisir comment les lecteurs chiffrés peuvent être récupérés"
# → ajouter le certificat DRA
```

> 🔐 **Règles d'or du DRA** : clé privée sur support amovible chiffré, stockée en coffre-fort, jamais sur le réseau. Deux personnes minimum connaissent son existence (anti-« bus factor »), mais une seule détient le support. Documenter la procédure d'usage d'urgence.

---

## 64. Conformité : stratégies de conformité Intune, base de référence

### 64.1 Stratégie de conformité type (Windows 11)

Intune → Appareils → **Conformité** → créer une stratégie Windows 11 :

| Paramètre | Valeur |
|---|---|
| Version minimale de l'OS | 10.0.26100 (24H2) |
| Chiffrement BitLocker requis | Oui |
| Antivirus requis / à jour | Oui |
| Pare-feu requis | Oui |
| Mot de passe requis (complexité) | Oui |
| Intégrité (attestation) | Selon criticité |
| Action en cas de non-conformité | Marquer non conforme → accès conditionnel bloque |

### 64.2 Accès conditionnel (le bras armé)

Entra ID → Protection → Accès conditionnel :

```
SI  (utilisateur du groupe "Tous") 
ET  (application = "Office 365" / "Toutes les applications cloud")
ALORS exiger : appareil conforme ET/OU appareil joint (hybride/Entra)
```

> ⚠️ **Toujours** tester une stratégie d'accès conditionnel en mode « Rapport seul » avant de l'appliquer, et prévoir un **compte d'urgence (break glass)** exclu de toutes les stratégies.

### 64.3 Rapports de conformité

```powershell
# Via Microsoft Graph : lister les appareils non conformes
# (nécessite le module Microsoft.Graph + consentement DeviceManagementManagedDevices.Read.All)
Connect-MgGraph -Scopes "DeviceManagementManagedDevices.Read.All"
Get-MgDeviceManagementManagedDevice -Filter "complianceState eq 'noncompliant'" |
    Select-Object DeviceName, OperatingSystem, OSVersion, ComplianceState, LastSyncDateTime
```

---

## 65. Kiosque et postes partagés : Assigned Access, invité

### 65.1 Assigned Access (kiosque mono-application)

Verrouille le poste sur **une seule application** (UWP ou navigateur) :

