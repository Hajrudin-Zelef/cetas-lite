---
id: collect-261001-rattrapage/rattrapage/win11-guide-16
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [2526, 2742]
sha256: 0e522fc5bc866f50b0b3ada868c701d044f5efd5d70516766b4eba3c2855b4a1
---

# Windows 11 en entreprise — Guide technique ultra-complet

| Journal | Chemin Event Viewer |
|---|---|
| Système | Journaux Windows → Système |
| Application | Journaux Windows → Application |
| Démarrage/arrêt | Système, ID 6005/6006/6008/6009, 1074 |
| Écrans bleus | Système, ID 1001 (BugCheck) |
| Windows Update | `Applications and Services Logs\Microsoft\Windows\WindowsUpdateClient\Operational` |
| GPO | `...\Microsoft\Windows\GroupPolicy\Operational` |
| BitLocker | `...\Microsoft\Windows\BitLocker\BitLocker-Operational` |
| Defender | `...\Microsoft\Windows\Windows Defender\Operational` |
| Inscription Entra | `...\Microsoft\Windows\User Device Registration\Admin` |

```powershell
# Extraire les erreurs récentes d'un journal (ex. : 7 derniers jours, niveau erreur)
Get-WinEvent -FilterHashtable @{
    LogName   = 'System'
    Level     = 2          # 2=Erreur, 3=Avertissement
    StartTime = (Get-Date).AddDays(-7)
} | Select-Object TimeCreated, Id, ProviderName,
    @{N='Message';E={ $_.Message.Substring(0, [Math]::Min(120, $_.Message.Length)) }} |
  Format-Table -AutoSize
```

---

## 42. Cas pratique 1 : le PC ne démarre plus (écran noir, BOOTMGR, BCD)

**Symptômes** : « Boot Device Not Found », écran noir avec curseur, ou boucle sur le logo.

**Diagnostic** :

```
1. Le BIOS/UEFI voit-il le disque ? (setup du firmware)
   → Non : problème matériel (câble, disque HS, contrôleur) — voir SMART
   → Oui : problème de démarrage logiciel → WinRE
2. Démarrer sur WinRE (clé USB ou démarrage avancé)
3. diskpart → list vol : la partition Windows est-elle visible et saine ?
```

**Réparation** (dans WinRE, invite de commandes) :

```powershell
# Identifier les lettres (en WinRE, C: n'est pas toujours Windows !)
diskpart
list vol
exit

# Cas UEFI : reconstruire le chargeur
bcdboot C:\Windows /s S: /f UEFI
# (S: = partition EFI montée au préalable via diskpart → assign letter=S)

# Cas Legacy/BIOS :
bootrec /fixmbr
bootrec /fixboot
bootrec /rebuildbcd

# Si le disque a des erreurs :
chkdsk C: /f /r
```

**Si BitLocker est actif** : déverrouiller d'abord avec `manage-bde -unlock` (section 18.4), sinon `chkdsk`/`bcdboot` travaillent sur du chiffré.

**Vérification** : 2 redémarrages complets (froid + chaud). Si le disque montre des secteurs réalloués (CrystalDiskInfo) → **remplacer le disque**, ne pas « réparer ».

---

## 43. Cas pratique 2 : boucle de réparation automatique / WinRE en boucle

**Symptômes** : « Réparation automatique » → « Votre PC n'a pas démarré correctement » → redémarrage → rebelote.

**Diagnostic** :

```
1. Dans WinRE → Options avancées → Invite de commandes
2. Lire le journal de la réparation :
   notepad C:\Windows\System32\LogFiles\Srt\SrtTrail.txt
   → la dernière section indique souvent la cause (pilote, KB, disque)
3. Vérifier les MàJ récentes : une KB installée juste avant ?
```

**Réparation** :

```powershell
# Option A : désinstaller la dernière MàJ qualité depuis WinRE
# WinRE → Dépannage → Désinstaller les mises à jour → "Désinstaller la dernière mise à jour qualité"

# Option B : en ligne de commande (identifier le package)
dism /Image:C:\ /Get-Packages /Format:Table | findstr "Installé"
dism /Image:C:\ /Remove-Package /PackageName:Package_for_RollupFix~31bf3856ad364e35~amd64~~26100.XXXX

# Option C : restauration système
# WinRE → Dépannage → Restauration du système → choisir un point antérieur

# Option D : désactiver la réparation auto pour voir le vrai BSOD
bcdedit /set {default} recoveryenabled No
# → au prochain boot, l'écran bleu affiche le code d'erreur au lieu de boucler
```

**Vérification** : réactiver la réparation après résolution (`recoveryenabled Yes`), puis 48 h d'observation.

---

## 44. Cas pratique 3 : BitLocker demande la clé de récupération en boucle

**Symptômes** : à chaque démarrage, l'écran bleu BitLocker réclame la clé à 48 chiffres.

**Diagnostic** : voir section 18 — causes typiques : MàJ BIOS sans suspension, changement de périphérique USB au boot, Secure Boot modifié, protecteur TPM corrompu.

**Résolution pas à pas** :

```powershell
# 1. Démarrer avec la clé (procédure section 18.1)
# 2. En session admin :
Suspend-BitLocker -MountPoint "C:" -RebootCount 1
Restart-Computer

# 3. Si ça redemande : recréer le protecteur TPM (section 18.3)
$vol = Get-BitLockerVolume -MountPoint "C:"
$tpmProt = $vol.KeyProtector | Where-Object { $_.KeyProtectorType -eq 'Tpm' }
Remove-BitLockerKeyProtector -MountPoint "C:" -KeyProtectorId $tpmProt.KeyProtectorId
Add-BitLockerKeyProtector -MountPoint "C:" -TpmProtector
Resume-BitLocker -MountPoint "C:"

# 4. Vérifier que la NOUVELLE clé de récupération est bien sauvegardée en AD !
Backup-BitLockerKeyProtector -MountPoint "C:" -KeyProtectorId (
    (Get-BitLockerVolume -MountPoint "C:").KeyProtector |
    Where-Object { $_.KeyProtectorType -eq 'RecoveryPassword' } |
    Select-Object -First 1 -ExpandProperty KeyProtectorId
)
```

**Prévention** : procédure écrite « toute MàJ BIOS = Suspend-BitLocker avant », affichée dans l'atelier.

---

## 45. Cas pratique 4 : mise à jour Windows bloquée / échec répété

**Symptômes** : Windows Update tourne en boucle, erreur 0x800700xx, ou « échec de l'installation » à chaque Patch Tuesday.

**Diagnostic** :

```powershell
# 1. Code d'erreur exact
Get-WinEvent -FilterHashtable @{
    LogName='System'; ProviderName='Microsoft-Windows-WindowsUpdateClient'; Level=2; StartTime=(Get-Date).AddDays(-3)
} | Select-Object TimeCreated, Id, Message | Format-List

# 2. Espace disque (cause n°1 !)
Get-PSDrive C | Select-Object Free, Used

# Codes fréquents :
# 0x80070070 / 0x80070002 : espace disque / fichiers corrompus
# 0x80240034             : téléchargement interrompu
# 0x800F0922             : partition EFI / de récupération trop petite
# 0x80070005             : droits d'accès
```

**Réparation** (dans l'ordre) :

```powershell
# Étape 1 : utilitaire de résolution des problèmes (Paramètres → Dépannage)
# Étape 2 : réinitialiser les composants Windows Update
net stop wuauserv
net stop cryptSvc
net stop bits
net stop msiserver
ren C:\Windows\SoftwareDistribution SoftwareDistribution.old
ren C:\Windows\System32\catroot2 catroot2.old
net start wuauserv
net start cryptSvc
net start bits
net start msiserver

# Étape 3 : réparer l'image puis relancer
dism /Online /Cleanup-Image /RestoreHealth
sfc /scannow

# Étape 4 : installer la KB manuellement (catalogue Microsoft Update)
# Télécharger le .msu correspondant puis :
wusa.exe C:\Admin\windows11.0-kb50xxxxx-x64.msu /quiet /norestart
```

**Désinstaller une KB fautive** :

```powershell
wusa /uninstall /kb:5034123 /quiet /norestart
# Puis mettre en pause les MàJ le temps de l'analyse (section 26.2)
```

**Vérification** : `Paramètres → Windows Update → Historique des mises à jour` = « Installation réussie », puis un cycle complet le mois suivant.

---

(Bloc 3/6 — sections 31 à 45)

---

## 46. Cas pratique 5 : mise à niveau 23H2 → 24H2 qui échoue

**Symptômes** : la feature update échoue à un pourcentage fixe (ex. 71 %), rollback vers 23H2.

**Diagnostic** :

```powershell
# 1. Lancer SetupDiag juste après l'échec (section 7.4)
.\SetupDiag.exe /Output:C:\Admin\SetupDiag.log
# Chercher "Error" / "Last Phase" (ex. : Safe OS, First Boot)

# 2. Lire setuperr.log
Get-Content "C:\`$Windows.~BT\Sources\Panther\setuperr.log" -Tail 30

# 3. Vérifier l'espace et les pilotes
# Erreur 0xC1900101-0x30018 = pilote ; 0xC1900208 = application incompatible
```

**Résolution** :

