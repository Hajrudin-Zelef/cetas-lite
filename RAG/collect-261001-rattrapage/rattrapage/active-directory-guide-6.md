---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-6
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution", "parameters"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [808, 984]
sha256: 0af2b28ece6efa2c33ec5bdc7e4e33b6497b47c1854c443af488cdb2552c5b3b
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

**Pourquoi** : quand Karim change de service, on le déplace d'un groupe global à un autre — **zéro modification des ACL**. Quand un nouveau partage ouvre, on crée un DL et on y imbrique les groupes existants.

```powershell
# Mettre en place la chaîne
Add-ADGroupMember -Identity "GG_Exploitation_Paris" -Members "a.diallo","k.benali"
Add-ADGroupMember -Identity "GU_Exploitation" -Members "GG_Exploitation_Paris"
Add-ADGroupMember -Identity "DL_Partage_Atelier_RW" -Members "GU_Exploitation"

# Vérifier l'appartenance récursive d'un utilisateur
Get-ADUser "a.diallo" -Properties MemberOf | Select-Object -ExpandProperty MemberOf
# Ou : tous les groupes effectifs (récursif)
(Get-ADUser "a.diallo" -Properties tokenGroups).tokenGroups |
  ForEach-Object { (New-Object System.Security.Principal.SecurityIdentifier $_).Translate([System.Security.Principal.NTAccount]) }
```

**Limites** : évitez plus de 3-4 niveaux d'imbrication (lisibilité, temps d'ouverture de session). Documentez la logique dans la description des groupes.

---

## 25. Groupes dynamiques : état des lieux et alternatives

⚠️ **AD DS natif ne propose pas de groupes dynamiques** (contrairement à Entra ID / Azure AD). Les appartenances sont statiques. Alternatives :

1. **Script planifié** : une tâche qui synchronise un groupe selon un attribut (ex. `Department=Exploitation` → membre de `GG_Exploitation`).

```powershell
# Exemple : synchroniser GG_Exploitation_Paris depuis l'attribut Department+Site
$group = "GG_Exploitation_Paris"
$attendus = Get-ADUser -Filter { Department -eq "Exploitation" } -Properties Department |
  Where-Object { $_.DistinguishedName -like "*OU=Paris*" } |
  Select-Object -ExpandProperty SamAccountName
$actuels = Get-ADGroupMember -Identity $group | Select-Object -ExpandProperty SamAccountName
# Ajouter les manquants
$attendus | Where-Object { $_ -notin $actuels } | ForEach-Object {
  Add-ADGroupMember -Identity $group -Members $_
}
# Retirer les intrus
$actuels | Where-Object { $_ -notin $attendus } | ForEach-Object {
  Remove-ADGroupMember -Identity $group -Members $_ -Confirm:$false
}
```

2. **Requêtes LDAP enregistrées** (Saved Queries) dans la console : affichent dynamiquement mais ne sont **pas** des groupes utilisables dans les ACL/GPO.
3. **Entra ID** : si synchronisé (Entra Connect), les groupes dynamiques Entra existent côté cloud uniquement.

---

## 26. Contacts, ordinateurs et comptes de service

**Contacts** : objets annuaire sans identifiant de sécurité (pas de SID), utilisés pour le carnet d'adresses (prestataires externes).

```powershell
New-ADObject -Name "Jean Prestataire" -Type contact `
  -Path "OU=Contacts,DC=ad,DC=entreprise,DC=fr" `
  -OtherAttributes @{ mail = "jean@prestataire.fr"; telephoneNumber = "0123456789" }
```

**Ordinateurs** : créés automatiquement à la jonction au domaine, ou pré-créés (pré-staging) pour contrôler l'OU de destination.

```powershell
# Pré-créer un compte d'ordinateur dans la bonne OU
New-ADComputer -Name "PC-ATELIER-042" `
  -Path "OU=Fixes,OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr" `
  -Description "Poste atelier - ligne 2"

# Rediriger le conteneur par défaut des nouveaux ordinateurs vers une OU gérée
redircmp "OU=Postes,DC=ad,DC=entreprise,DC=fr"
# Idem pour les utilisateurs (rarement utile, mais existe)
redirusr "OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr"
```

> 💡 `redircmp` : indispensable. Sinon les PC jonchés tombent dans `CN=Computers`, hors de vos GPO/OU. À faire **une fois** par domaine.

**Comptes de service classiques** : bannissez les comptes « utilisateur » avec mot de passe qui n'expire jamais partagés entre applis. Préférez les **gMSA** (section suivante).

---

## 27. Comptes de service gérés (gMSA)

Les **gMSA** (group Managed Service Accounts) : mots de passe **gérés automatiquement** par AD (rotation toutes les 30 jours), utilisables par plusieurs serveurs, sans intervention humaine.

```powershell
# 1. Préparer la forêt (une seule fois) : créer la clé racine KDS
Add-KdsRootKey -EffectiveTime ((Get-Date).AddHours(-10))
# ⚠️ La clé n'est utilisable que 10 h après sa création (ou forcez une date passée en labo uniquement)

# 2. Créer le gMSA
New-ADServiceAccount -Name "gmsa_Sauvegarde" `
  -DNSHostName "gmsa_Sauvegarde.ad.entreprise.fr" `
  -PrincipalsAllowedToRetrieveManagedPassword "GG_Serveurs_Sauvegarde"

# 3. Sur chaque serveur autorisé : installer le compte
Install-ADServiceAccount -Identity "gmsa_Sauvegarde"

# 4. L'utiliser dans un service ou une tâche planifiée
# Service : compte "AD.ENTREPRISE.FR\gmsa_Sauvegarde$" (mot de passe vide)
# Tâche planifiée :
$action = New-ScheduledTaskAction -Execute "C:\Scripts\backup.ps1"
Register-ScheduledTask -TaskName "Sauvegarde quotidienne" -Action $action `
  -User "ad.entreprise.fr\gmsa_Sauvegarde$" -RunLevel Highest
```

**Cas d'usage** : services Windows, tâches planifiées, pools IIS, SQL Server. **Limites** : ne fonctionne pas pour les services interactifs nécessitant un profil complet ni hors du domaine.

---

## 28. Les 5 rôles FSMO : vue d'ensemble

Certaines opérations ne supportent pas le multi-maître : 5 rôles **FSMO** (Flexible Single Master Operations) sont attribués à des DC désignés.

| # | Rôle | Portée | Sensible à la panne ? |
|---|---|---|---|
| 1 | Maître de **schéma** | Forêt | Non (sauf modif du schéma en cours) |
| 2 | Maître d'attribution des **noms de domaine** | Forêt | Non (sauf ajout/suppression de domaine) |
| 3 | Maître **RID** | Domaine | Oui à moyen terme (épuisement des pools RID) |
| 4 | **Émulateur PDC** | Domaine | **Oui** (heure, verrouillages, GPO, mots de passe) |
| 5 | Maître d'**infrastructure** | Domaine | Peu (sauf multi-domaines) |

```powershell
# Voir tous les détenteurs FSMO de la forêt
Get-ADForest | Select-Object SchemaMaster, DomainNamingMaster
Get-ADDomain | Select-Object RIDMaster, PDCEmulator, InfrastructureMaster

# Ou en une ligne :
netdom query fsmo
```

---

## 29. Maître de schéma (Schema Master)

**Rôle** : seul DC autorisé à **modifier le schéma** (ajouter des classes/attributs). Portée forêt, un seul détenteur.

**Quand on y touche** : installation d'Exchange, de SCCM, d'applications étendant le schéma, `adprep /forestprep` lors d'une montée de version.

```powershell
# Voir le détenteur
(Get-ADForest).SchemaMaster

# Autoriser les mises à jour du schéma sur ce DC (verrou de sécurité)
# Registre : HKLM\SYSTEM\CurrentControlSet\Services\NTDS\Parameters\Schema Update Allowed = 1
```

**Bonnes pratiques** :

- Ne transférez ce rôle que si nécessaire ; il peut rester sur le premier DC.
- **Sauvegardez l'état système** avant toute extension de schéma (irréversible : on peut désactiver un attribut, jamais le supprimer vraiment).
- Testez l'extension en labo.

---

## 30. Maître d'attribution des noms de domaine (Domain Naming Master)

**Rôle** : autorise l'**ajout et la suppression de domaines** dans la forêt. Vérifie l'unicité des noms.

```powershell
(Get-ADForest).DomainNamingMaster
```

Peu sollicité au quotidien. Comme le maître de schéma, il vit généralement sur le premier DC de la forêt et on l'y laisse.

---

## 31. Maître RID (Relative ID Master)

**Rôle** : distribue des **pools de RID** aux DC. Chaque objet créé consomme un RID (le SID = SID du domaine + RID). Quand un DC épuise son pool (500 par défaut), il en redemande au maître RID.

**Danger** : si le maître RID est indisponible longtemps et que les pools s'épuisent, **impossible de créer** utilisateurs/ordinateurs/groupes.

```powershell
(Get-ADDomain).RIDMaster

# Vérifier la consommation du pool RID sur un DC
dcdiag /test:ridmanager /v
# Événement 16650 dans le journal Directory Service = pool épuisé : urgence
```

