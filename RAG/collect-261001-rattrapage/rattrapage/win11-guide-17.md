---
id: collect-261001-rattrapage/rattrapage/win11-guide-17
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [2743, 2938]
sha256: 19cae9878255963c19053c0cbf128d63df59342f84b9cb106f1574103df960e4
---

# Windows 11 en entreprise — Guide technique ultra-complet

```
1. Mettre à jour le BIOS/UEFI et les pilotes chipset/stockage (constructeur)
2. Débrancher TOUT les périphériques USB non essentiels (dongles, disques externes)
3. Désinstaller l'antivirus tiers (temporairement)
4. Suspendre BitLocker : Suspend-BitLocker -MountPoint "C:" -RebootCount 1
5. Relancer via l'ISO : setup.exe /auto upgrade /dynamicupdate enable
6. Si échec persistant : installation propre + USMT/OneDrive (section 21)
   plutôt que de s'acharner sur une 3e tentative
```

**Vérification** : `winver` = 24H2 (build 26100+), applications métier testées, BitLocker réactivé (`Resume-BitLocker` + vérifier `ProtectionStatus`).

---

## 47. Cas pratique 6 : pilote défectueux — écran bleu, rollback, signature

**Symptômes** : BSOD après MàJ d'un pilote (souvent GPU, Wi-Fi, stockage), code `DRIVER_IRQL_NOT_LESS_OR_EQUAL`, `INACCESSIBLE_BOOT_DEVICE`.

**Diagnostic** :

```powershell
# 1. Identifier le pilote fautif (fichier .sys dans le BSOD ou le dump)
# Analyser le minidump avec WinDbg ou BlueScreenView (NirSoft)
# Emplacement : C:\Windows\Minidump\*.dmp

# 2. Voir l'historique des pilotes
Get-WinEvent -FilterHashtable @{LogName='System'; Id=20001,20003; StartTime=(Get-Date).AddDays(-7)} |
    Select-Object TimeCreated, Message | Format-List
```

**Résolution** :

```powershell
# Option A : restaurer le pilote précédent (si Windows démarre, même en mode sans échec)
# Gestionnaire de périphériques → Propriétés → Pilote → "Restaurer le pilote"
# En PowerShell (via pnputil) :
pnputil /enum-drivers | Select-String "oem" -Context 2

# Option B : désinstaller le pilote fautif depuis WinRE
# Identifier l'oemXX.inf du pilote, puis :
dism /Image:C:\ /Remove-Driver /Driver:oem12.inf

# Option C : démarrer en mode sans échec (pilotes minimaux)
bcdedit /set {default} safeboot minimal
# (puis `bcdedit /deletevalue {default} safeboot` pour revenir)

# Option D : bloquer la réinstallation auto du pilote fautif par Windows Update
# GPO : "Ne pas inclure les pilotes avec les mises à jour Windows" (section 26.3)
```

**Signature des pilotes** : Windows 11 exige des pilotes **signés**. Désactiver l'obligation de signature (`bcdedit /set nointegritychecks on`) = **à proscrire** en production (casse HVCI et Secure Boot).

---

## 48. Cas pratique 7 : performances dégradées — disque, mémoire, processus

**Symptômes** : lenteurs générales, ventilateur à fond, disque à 100 %.

**Diagnostic** (dans l'ordre) :

```powershell
# 1. Qui consomme ? (top 10 CPU / mémoire)
Get-Process | Sort-Object CPU -Descending | Select-Object -First 10 Name, CPU, WorkingSet
Get-Process | Sort-Object WorkingSet64 -Descending | Select-Object -First 10 Name,
    @{N='RAM(Mo)';E={[math]::Round($_.WorkingSet64/1MB)}}

# 2. Disque : file d'attente et temps de réponse
Get-Counter '\PhysicalDisk(_Total)\% Disk Time', '\PhysicalDisk(_Total)\Avg. Disk Queue Length'

# 3. Démarrage : qu'est-ce qui se lance ?
Get-CimInstance Win32_StartupCommand | Select-Object Name, Command, Location

# 4. Santé du disque (SMART) — via CrystalDiskInfo ou :
Get-PhysicalDisk | Select-Object FriendlyName, MediaType, HealthStatus, OperationalStatus
```

**Résolutions typiques** :

| Cause | Action |
|---|---|
| Antivirus tiers + Defender en conflit | Un seul antivirus actif ; désinstaller l'autre |
| Indexation / SysMain sur HDD | Passer au SSD (la vraie solution) |
| Trop d'apps au démarrage | Désactiver via Gestionnaire des tâches → Démarrage |
| Mémoire saturée (8 Go + Teams + Edge) | Ajouter de la RAM ou limiter les onglets/processus |
| Disque système plein à >90 % | Nettoyage (cas pratique 14) |
| Profil utilisateur énorme | Nettoyer / recréer le profil |

```powershell
# Désactiver proprement SysMain (ex-Superfetch) si pertinent — rarement utile sur SSD
Stop-Service SysMain -Force
Set-Service SysMain -StartupType Disabled

# Analyser le démarrage avec Autoruns (Sysinternals) — à faire en admin
```

---

## 49. Cas pratique 8 : profil utilisateur corrompu / temporaire

**Symptômes** : « Vous avez été connecté avec un profil temporaire », bureau vide, paramètres perdus.

**Diagnostic** :

```powershell
# 1. Vérifier l'état du profil dans le registre
Get-ChildItem "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion\ProfileList" |
    ForEach-Object {
        $p = Get-ItemProperty $_.PSPath
        [pscustomobject]@{
            SID        = $_.PSChildName
            Path       = $p.ProfileImagePath
            State      = $p.State          # 0 = OK
            RefCount   = $p.RefCount
            Temporaire = ($_.PSChildName -like "*.bak") -or ($p.State -ne 0)
        }
    } | Format-Table -AutoSize

# Signes typiques : une clé SID.bak + une clé SID "neuve" pointant vers C:\Users\TEMP
```

**Résolution** :

```
1. Sauvegarder les données du profil corrompu (C:\Users\<nom>\Desktop, Documents...)
   — même s'il semble vide, vérifier via un compte admin
2. Déconnecter l'utilisateur (pas de session verrouillée !)
3. Dans le registre : supprimer la clé SID.bak ET la clé SID temporaire
4. Supprimer/renommer C:\Users\<nom> (ou le laisser, Windows recréera)
   → méthode propre : Remove-CimInstance sur Win32_UserProfile (section 19.3)
5. L'utilisateur se reconnecte → nouveau profil sain
6. Restaurer ses données (pas tout le profil : uniquement les dossiers utiles)
```

> ⛔ **Ne jamais** « réparer » en renommant `.bak` à la main sans comprendre : si le `NTUSER.DAT` est corrompu, le profil replantera. Recréation propre + restauration sélective des données = plus fiable.

---

## 50. Cas pratique 9 : impossible de joindre le domaine / relation d'approbation

**Symptômes** : « La relation d'approbation entre cette station et le domaine a échoué ».

**Causes** : mot de passe du compte machine désynchronisé (restauration d'image, clone, poste éteint > 30 jours, double SID).

**Résolution** (sans quitter/réintégrer le domaine si possible) :

```powershell
# Option A : réinitialiser le mot de passe du compte machine (le poste doit voir un DC)
Reset-ComputerMachinePassword -Server "DC01.entreprise.local" -Credential (Get-Credential)
Restart-Computer

# Option B : quitter puis rejoindre (si A échoue)
Remove-Computer -UnjoinDomainCredential (Get-Credential) -Restart
# après redémarrage :
Add-Computer -DomainName "entreprise.local" -OUPath "OU=Postes Win11,DC=entreprise,DC=local" -Credential (Get-Credential) -Restart
```

**Diagnostic réseau préalable** (50 % des « échecs de jointure » sont du DNS) :

```powershell
# Le poste résout-il le domaine ? Le DC est-il joignable ?
Resolve-DnsName entreprise.local
Test-NetConnection DC01.entreprise.local -Port 389
nltest /dsgetdc:entreprise.local
# L'heure est-elle synchronisée ? (Kerberos tolère ±5 min !)
w32tm /query /status
```

---

## 51. Cas pratique 10 : GPO qui ne s'applique pas

**Symptômes** : le réglage n'est pas effectif malgré la GPO liée.

**Checklist de diagnostic** (dans l'ordre) :

```powershell
# 1. La GPO est-elle reçue ?
gpresult /r
# → Chercher la GPO dans "Objets de stratégie de groupe appliqués"
# → Si absente : filtrage, liaison, OU, réplication ?

# 2. Détail complet
gpresult /h C:\Admin\gpresult.html

# 3. Le poste est-il dans la bonne OU ?
Get-ADComputer $env:COMPUTERNAME | Select-Object DistinguishedName

# 4. La GPO est-elle liée ET activée ? (console GPMC)
# Vérifier : lien activé, GPO activée (pas "désactivée"), pas de "Block Inheritance" au-dessus

# 5. Filtrage de sécurité : le groupe "Ordinateurs authentifiés" a-t-il "Lire" + "Appliquer" ?
# (depuis la KB3163622, le filtrage par groupe exige ces droits — cause classique !)

# 6. Filtre WMI : s'évalue-t-il à vrai ?
Get-CimInstance -Query 'SELECT * FROM Win32_OperatingSystem WHERE Version LIKE "10.0.26%"'

