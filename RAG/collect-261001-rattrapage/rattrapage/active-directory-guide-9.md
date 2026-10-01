---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-9
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1344, 1512]
sha256: 41cd0e1c8a50217fc9719111f201a7737d63ce3c80d7d152cbff03570b55d054
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

> ⚠️ La stratégie définie dans une GPO liée à une **OU** ne s'applique **pas** aux mots de passe (sauf FGPP, section suivante). Seule la Default Domain Policy (ou une GPO liée à la racine avec précédence) compte.

---

## 46. FGPP : stratégies de mot de passe affinées (PSO)

Les **FGPP** (Fine-Grained Password Policies) permettent des politiques **différentes par groupe** via des objets **PSO** (Password Settings Object). Ordre de précédence : PSO lié directement à l'utilisateur > PSO du groupe (le plus petit `msDS-PasswordSettingsPrecedence` gagne).

```powershell
# Créer un PSO strict pour les administrateurs
New-ADFineGrainedPasswordPolicy -Name "PSO-Admins" `
  -Precedence 10 `
  -MinPasswordLength 16 `
  -PasswordHistoryCount 24 `
  -ComplexityEnabled $true `
  -ReversibleEncryptionEnabled $false `
  -MinPasswordAge "1.00:00:00" `
  -MaxPasswordAge "60.00:00:00" `
  -LockoutThreshold 3 `
  -LockoutObservationWindow "00:30:00" `
  -LockoutDuration "00:30:00"

# L'appliquer au groupe des admins du domaine
Add-ADFineGrainedPasswordPolicySubject -Identity "PSO-Admins" -Subjects "Admins du domaine"

# Créer un PSO assoupli pour les comptes de service legacy (exemple)
New-ADFineGrainedPasswordPolicy -Name "PSO-Services" -Precedence 50 `
  -MinPasswordLength 20 -PasswordHistoryCount 24 -ComplexityEnabled $true `
  -ReversibleEncryptionEnabled $false -MinPasswordAge "1.00:00:00" `
  -MaxPasswordAge "180.00:00:00" -LockoutThreshold 0

# Voir le PSO effectif d'un utilisateur
Get-ADUserResultantPasswordPolicy -Identity "a.diallo" | Select-Object Name, Precedence
```

**Bonnes pratiques** : un PSO « Admins » strict, un PSO « Utilisateurs » standard (= défaut du domaine), precedence basse = prioritaire. Documentez les PSO comme des GPO.

---

## 47. Bonnes pratiques mots de passe et verrouillage

1. **Longueur > complexité** : une passphrase de 16 caractères (`correct-cheval-batterie-agrafes`) bat `P@ssw0rd!` en sécurité et en mémorisation.
2. **Ne forcez pas** les rotations trop fréquentes (42 j par défaut pousse aux `Motdepasse1`, `Motdepasse2`...). 90-180 j + surveillance des compromissions.
3. **Bannissez les mots de passe connus** : comparez les hachés aux listes de fuites (HaveIBeenPwned, outils type Lithnet Password Protection).
4. **MFA partout** où c'est possible, surtout admins et accès distants.
5. **Comptes de service** : gMSA (section 27), jamais de compte utilisateur avec « le mot de passe n'expire jamais ».
6. **Verrouillage** : 5 essais / 30 min est un bon équilibre. Surveillez les verrouillages répétés (attaque par dictionnaire ?).
7. **Communication** : à l'embauche, mot de passe initial remis en main propre ou via canal séparé, changement imposé à la première ouverture.

```powershell
# Rapport : comptes avec "mot de passe n'expire jamais" (à réduire au minimum)
Get-ADUser -Filter { PasswordNeverExpires -eq $true -and Enabled -eq $true } `
  -Properties PasswordNeverExpires |
  Select-Object Name, SamAccountName, DistinguishedName

# Rapport : mots de passe vieux de plus de 180 jours
$seuil = (Get-Date).AddDays(-180)
Get-ADUser -Filter { PasswordLastSet -lt $seuil -and Enabled -eq $true } `
  -Properties PasswordLastSet |
  Select-Object Name, SamAccountName, PasswordLastSet | Sort-Object PasswordLastSet
```

---

## 48. Corbeille AD : activation et fonctionnement

La **corbeille AD** (AD Recycle Bin) permet de **restaurer des objets supprimés avec tous leurs attributs** (y compris les appartenances aux groupes). Sans elle, la restauration ne récupère qu'une coquille (tombstone).

**Prérequis** : niveau fonctionnel de forêt **2008 R2 minimum** (donc OK en 2019/2022/2025). **Activation irréversible**.

```powershell
# Activer la corbeille (cible : nom de la forêt)
Enable-ADOptionalFeature -Identity "Recycle Bin Feature" `
  -Scope ForestOrConfigurationSet -Target "ad.entreprise.fr" -Confirm:$false

# Vérifier
Get-ADOptionalFeature -Filter { Name -eq "Recycle Bin Feature" } |
  Select-Object Name, @{N="Activée";E={$_.EnabledScopes}}
```

**Cycle de vie d'un objet supprimé** (avec corbeille) : objet supprimé → état **deleted** (restaurable, 180 jours par défaut = `msDS-deletedObjectLifetime`) → objet **recyclé** → destruction physique.

---

## 49. Restaurer un objet supprimé (Restore-ADObject)

```powershell
# 1. Trouver l'objet supprimé
Get-ADObject -Filter { Name -like "*Diallo*" } -IncludeDeletedObjects |
  Select-Object Name, ObjectClass, Deleted, LastKnownParent

# 2. Restaurer à son emplacement d'origine
Get-ADObject -Filter { SamAccountName -eq "a.diallo" } -IncludeDeletedObjects |
  Restore-ADObject

# 3. Restaurer vers une OU différente (plus sûr : on contrôle)
Get-ADObject -Filter { SamAccountName -eq "a.diallo" } -IncludeDeletedObjects |
  Restore-ADObject -NewName "Amina Diallo" `
    -TargetPath "OU=Restaures,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr"

# 4. Restaurer une OU entière avec son contenu
Get-ADObject -Filter { Name -eq "Exploitation" -and ObjectClass -eq "organizationalUnit" } `
  -IncludeDeletedObjects | Restore-ADObject
# Puis restaurer les objets enfants un par un (la restauration est hiérarchique : parent d'abord)
```

**Après restauration** : réactivez le compte (`Enable-ADAccount`), vérifiez les appartenances aux groupes (restaurées avec la corbeille, **pas** sans), réinitialisez le mot de passe (par sécurité).

**Sans corbeille** (tombstone classique) : restauration d'autorité via `ntdsutil authoritative restore` (section 53) ou recréation manuelle — les appartenances aux groupes sont **perdues**.

---

## 50. Sauvegarde : état système avec wbadmin

On ne sauvegarde pas `ntds.dit` à la main : on fait une **sauvegarde de l'état système** (System State) qui inclut AD, SYSVOL, registre, services.

```powershell
# Sauvegarde de l'état système vers un disque dédié (jamais sur le disque système)
wbadmin start systemstatebackup -backupTarget:E: -quiet

# Planifier une sauvegarde quotidienne à 22h
wbadmin enable backup -addtarget:E: -systemstate -schedule:22:00 -quiet

# Vérifier les sauvegardes disponibles
wbadmin get versions

# Sauvegarde bare-metal complète (état système + volumes critiques) — recommandée
wbadmin start backup -backupTarget:E: -allCritical -systemstate -quiet
```

**Bonnes pratiques** :

- Sauvegardez **au moins un DC par domaine** chaque jour (idéalement deux DC différents).
- Conservez plusieurs générations (rotation 7-14 jours).
- **Testez la restauration** en labo isolé au moins une fois par semestre (une sauvegarde non testée = pas de sauvegarde).
- Stockez une copie **hors site** (copie du dossier WindowsImageBackup).
- Surveillez le succès via le journal `Microsoft-Windows-Backup` et une alerte.

---

## 51. Sauvegarder via ntdsutil et bonnes pratiques de sauvegarde

`ntdsutil` permet aussi de créer un snapshot applicatif (utile pour monter une copie et comparer, pas comme sauvegarde principale) :

```powershell
# Créer un snapshot de la base (montable en lecture seule pour inspection)
ntdsutil "activate instance ntds" "snapshot" "create" quit quit
# Lister / monter / démonter
ntdsutil "activate instance ntds" "snapshot" "list all" quit quit
```

**Ce qu'il faut sauvegarder, en résumé** :

| Élément | Moyen | Fréquence |
|---|---|---|
| État système des DC | wbadmin | Quotidien |
| GPO (sauvegardes GPMC) | `Backup-GPO -All` | Hebdo + avant chaque changement |
| Documentation (FSMO, topologie, mots de passe DSRM) | Coffre sécurisé | À chaque changement |
| Scripts de provisionning | Dépôt versionné | Continu |

```powershell
# Sauvegarder toutes les GPO (complément indispensable)
Backup-GPO -All -Path "E:\Sauvegardes\GPO" -Comment "Sauvegarde hebdo $(Get-Date -Format 'yyyy-MM-dd')"
```

---

## 52. DSRM : mode de restauration des services d'annuaire

