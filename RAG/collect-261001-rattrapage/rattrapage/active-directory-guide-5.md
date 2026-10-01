---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-5
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["diffusion", "distribution"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [637, 807]
sha256: d7b1f4a7ff0f86ef84cf130c6cd3ec57ef6583c0bd248fee4c8c30a6f79b4329
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

**Centre d'administration Active Directory (ADAC, `dsacls.msc`)** : interface moderne, parfait pour la création unitaire et la corbeille. **Utilisateurs et ordinateurs AD (`dsa.msc`)** : l'ancienne console, encore très utilisée.

```powershell
# Création unitaire en PowerShell
New-ADUser -Name "Amina Diallo" `
  -GivenName "Amina" -Surname "Diallo" `
  -SamAccountName "a.diallo" `
  -UserPrincipalName "a.diallo@ad.entreprise.fr" `
  -Path "OU=Exploitation,OU=Paris,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr" `
  -AccountPassword (ConvertTo-SecureString "MotDePasseTemporaire!2026" -AsPlainText -Force) `
  -Enabled $true -ChangePasswordAtLogon $true `
  -Title "Technicienne maintenance" -Department "Exploitation" -Company "Entreprise"

# Ajouter au groupe du service
Add-ADGroupMember -Identity "GG_Exploitation_Paris" -Members "a.diallo"
```

**Champs à renseigner systématiquement** : nom complet, UPN, service (`Department`), fonction (`Title`), responsable (`Manager`), date d'expiration pour les temporaires (`AccountExpirationDate`). Votre futur vous remerciera pendant les audits.

---

## 20. Création en masse via CSV (New-ADUser)

Le cas le plus courant : une rentrée, 40 intérimaires, un fichier RH.

**Fichier `arrivees.csv`** (UTF-8, séparateur `;`) :

```csv
Prenom;Nom;Login;Service;Site;Fonction
Amina;Diallo;a.diallo;Exploitation;Paris;Technicienne maintenance
Karim;Benali;k.benali;Maintenance;Lyon;Electrotechnicien
```

**Script** :

```powershell
$users = Import-Csv -Path "C:\Scripts\arrivees.csv" -Delimiter ";" -Encoding UTF8
$domain = "ad.entreprise.fr"

foreach ($u in $users) {
  $ou = "OU=$($u.Service),OU=$($u.Site),OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr"
  # Vérifier que l'OU existe
  if (-not (Get-ADOrganizationalUnit -Filter "Name -eq '$($u.Service)'" -SearchBase "OU=$($u.Site),OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr")) {
    Write-Warning "OU introuvable pour $($u.Login), ignoré"
    continue
  }
  # Éviter les doublons
  if (Get-ADUser -Filter "SamAccountName -eq '$($u.Login)'") {
    Write-Warning "Le compte $($u.Login) existe déjà, ignoré"
    continue
  }
  $mdp = ConvertTo-SecureString "Bienvenue!2026" -AsPlainText -Force
  New-ADUser -Name "$($u.Prenom) $($u.Nom)" `
    -GivenName $u.Prenom -Surname $u.Nom `
    -SamAccountName $u.Login `
    -UserPrincipalName "$($u.Login)@$domain" `
    -Path $ou -Department $u.Service -Title $u.Fonction `
    -AccountPassword $mdp -Enabled $true -ChangePasswordAtLogon $true
  Add-ADGroupMember -Identity "GG_$($u.Service)_$($u.Site)" -Members $u.Login
  Write-Host "Créé : $($u.Login)" -ForegroundColor Green
}
```

> 💡 Le mot de passe initial doit être **communiqué de façon sécurisée** (pas par mail en clair) et l'option « changer au premier logon » force l'utilisateur à le personnaliser immédiatement.

---

## 21. Attributs utilisateur importants et bonnes pratiques

| Attribut | Usage |
|---|---|
| `sAMAccountName` | Login court (max 20 car.). Convention : `p.nom`. |
| `userPrincipalName` | Login long `p.nom@domaine`. Alignez-le sur l'email si possible. |
| `displayName` | Affichage dans le carnet d'adresses : « NOM, Prénom » ou « Prénom NOM » (choisissez une convention). |
| `manager` | Responsable hiérarchique (objet lié). Indispensable pour les workflows. |
| `department` / `title` / `company` | Service, fonction, société. Base des tris et des groupes dynamiques. |
| `employeeID` | Matricule RH. Clé de rapprochement avec la SIRH. |
| `accountExpires` | Date de fin pour temporaires/stagiaires/intérimaires. **Toujours renseigner** pour les non-CDI. |
| `info` / `description` | Notes libres (« Intérimaire via Agence X, fin prévue... »). |

```powershell
# Trouver les comptes dont l'expiration arrive dans 30 jours
$limit = (Get-Date).AddDays(30)
Get-ADUser -Filter { AccountExpirationDate -lt $limit -and Enabled -eq $true } `
  -Properties AccountExpirationDate, Department |
  Select-Object Name, SamAccountName, Department, AccountExpirationDate |
  Sort-Object AccountExpirationDate

# Exporter l'annuaire pour la RH (sans données sensibles)
Get-ADUser -Filter * -Properties Department, Title, Enabled |
  Select-Object Name, SamAccountName, Department, Title, Enabled |
  Export-Csv "C:\Rapports\annuaire.csv" -NoTypeInformation -Encoding UTF8
```

---

## 22. Désactiver, déplacer, supprimer des comptes : cycle de vie

**Procédure de départ** (à automatiser autant que possible) :

```powershell
$login = "a.diallo"

# 1. Désactiver immédiatement (coupe tous les accès)
Disable-ADAccount -Identity $login

# 2. Réinitialiser le mot de passe avec une valeur aléatoire
$random = ConvertTo-SecureString ([guid]::NewGuid().ToString()) -AsPlainText -Force
Set-ADAccountPassword -Identity $login -NewPassword $random -Reset

# 3. Déplacer vers l'OU des comptes désactivés (sort des GPO utilisateurs actives)
Move-ADObject -Identity (Get-ADUser $login).DistinguishedName `
  -TargetPath "OU=Desactives,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr"

# 4. Retirer des groupes (sauf "Domain Users")
Get-ADUser $login -Properties MemberOf | Select-Object -ExpandProperty MemberOf |
  ForEach-Object { Remove-ADGroupMember -Identity $_ -Members $login -Confirm:$false }

# 5. Masquer du carnet d'adresses si Exchange, noter la date
Set-ADUser $login -Description "Départ le $(Get-Date -Format 'yyyy-MM-dd')"
```

**Suppression** : attendez le délai légal/RH (souvent 30 à 90 jours), la corbeille AD (section 48) vous protège des suppressions accidentelles mais **ne remplace pas** une procédure.

```powershell
# Comptes inactifs depuis plus de 90 jours (candidats au nettoyage)
$date = (Get-Date).AddDays(-90)
Search-ADAccount -AccountInactive -TimeSpan 90 -UsersOnly |
  Where-Object { $_.Enabled } |
  Select-Object Name, SamAccountName, LastLogonDate
```

---

## 23. Groupes : portées (global, universel, local de domaine)

| Portée | Contenu possible | Visible/utilisable | Usage typique |
|---|---|---|---|
| **Global (G)** | Utilisateurs, ordinateurs, autres groupes globaux **du même domaine** | Tout le domaine (et forêts approuvées) | Regrouper des utilisateurs par métier/site |
| **Universel (U)** | Utilisateurs/groupes globaux de **toute la forêt** | Toute la forêt (via le catalogue global) | Regrouper des globaux de plusieurs domaines |
| **Local de domaine (DL)** | Tout (utilisateurs, globaux, universels, d'autres domaines) | **Domaine local uniquement** | Donner des droits sur des ressources du domaine |

**Moyen mnémotechnique** : **G**lobal = **G**ens (les utilisateurs), **D**L = **D**roits (sur les ressources).

```powershell
# Créer les trois types
New-ADGroup -Name "GG_Exploitation_Paris" -GroupScope Global -GroupCategory Security `
  -Path "OU=Groupes,DC=ad,DC=entreprise,DC=fr"
New-ADGroup -Name "GU_Exploitation" -GroupScope Universal -GroupCategory Security `
  -Path "OU=Groupes,DC=ad,DC=entreprise,DC=fr"
New-ADGroup -Name "DL_Partage_Atelier_RW" -GroupScope DomainLocal -GroupCategory Security `
  -Path "OU=Groupes,DC=ad,DC=entreprise,DC=fr"
```

**Groupes de sécurité vs distribution** : les groupes de **sécurité** ont un SID et servent aux droits ; les groupes de **distribution** servent aux listes de diffusion (messagerie) et n'ont pas de SID utilisable pour les ACL. En cas de doute : sécurité.

---

## 24. Imbrication de groupes : stratégie AGDLP/AGUDLP

**AGDLP** (mono-domaine) : **A**ccounts → **G**lobal → **D**omain **L**ocal → **P**ermissions.
**AGUDLP** (multi-domaines) : Accounts → Global → **U**niversal → Domain Local → Permissions.

```
Utilisateurs (a.diallo, k.benali)
   └── GG_Exploitation_Paris (Global : les gens)
          └── GU_Exploitation (Universel : fédère les sites)
                 └── DL_Partage_Atelier_RW (Local de domaine : les droits)
                        └── ACL du partage \\SRV\Atelier (Modify)
```

