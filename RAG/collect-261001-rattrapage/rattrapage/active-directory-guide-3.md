---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-3
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [287, 470]
sha256: ef9cf5ae4803994545706df912048293ceb63b911c6e8c5097802ced9d251c86
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

- [ ] **Nom du serveur** définitif (renommer un DC après coup = douloureux, possible mais à éviter).
- [ ] **IP statique** configurée, DNS pointant vers lui-même (ou vers le DC existant pour un DC supplémentaire, puis vers lui-même une fois promu).
- [ ] **Rôle AD DS** installé : `Install-WindowsFeature AD-Domain-Services -IncludeManagementTools`.
- [ ] **Fuseau horaire et heure corrects** (Kerberos tolère 5 min de décalage par défaut).
- [ ] **Disque NTFS** avec espace suffisant pour `ntds.dit` + SYSVOL (jamais sur un lecteur compressé/chiffré EFS).
- [ ] **DNS** : le nom de domaine doit se résoudre ; pour une nouvelle forêt, le serveur peut s'auto-héberger.
- [ ] **Compte** membre des Admins de l'entreprise (nouvelle forêt) ou Admins du domaine (DC supplémentaire).
- [ ] **Niveau fonctionnel** cible décidé (2016 mini recommandé ; 2025 si tout le parc est en 2022/2025, voir section 82).
- [ ] **Mot de passe DSRM** choisi et **stocké dans le coffre** (section 52).

```powershell
# Installer le rôle AD DS
Install-WindowsFeature AD-Domain-Services -IncludeManagementTools

# Tester les prérequis d'une nouvelle forêt (ne fait rien, rapporte seulement)
Test-ADDSDomainControllerInstallation -DomainName "entreprise.lan" -Credential (Get-Credential)
```

---

## 9. Promouvoir le premier DC d'une nouvelle forêt (PowerShell)

```powershell
# Promouvoir le serveur en premier DC d'une nouvelle forêt
Install-ADDSForest `
  -DomainName "ad.entreprise.fr" `
  -DomainNetbiosName "ENTREPRISE" `
  -ForestMode "WinThreshold" `      # 2016+ : "WinThreshold" ; 2025 : "Win2025" (voir section 82)
  -DomainMode "WinThreshold" `
  -InstallDns `
  -CreateDnsDelegation:$false `
  -DatabasePath "C:\Windows\NTDS" `
  -LogPath "C:\Windows\NTDS" `
  -SysvolPath "C:\Windows\SYSVOL" `
  -NoRebootOnCompletion:$false `
  -SafeModeAdministratorPassword (Read-Host "Mot de passe DSRM" -AsSecureString) `
  -Force
```

**Ce qui se passe** : création de la forêt, du domaine racine, du schéma, du partage SYSVOL, installation et configuration du DNS intégré, le serveur devient GC et détient les 5 rôles FSMO.

**Après la promotion** :

```powershell
# Vérifier que le DC fonctionne
dcdiag /v | Select-String "passed|failed"
Get-ADDomain | Select-Object Name, DomainMode
Get-ADForest | Select-Object Name, ForestMode
```

---

## 10. Ajouter un DC supplémentaire dans un domaine existant

```powershell
# Sur le futur DC : pointer son DNS vers un DC existant, joindre le domaine, puis :
Install-ADDSDomainController `
  -DomainName "ad.entreprise.fr" `
  -Credential (Get-Credential "ENTREPRISE\administrateur") `
  -InstallDns `
  -CreateDnsDelegation:$false `
  -DatabasePath "C:\Windows\NTDS" `
  -LogPath "C:\Windows\NTDS" `
  -SysvolPath "C:\Windows\SYSVOL" `
  -NoRebootOnCompletion:$false `
  -SafeModeAdministratorPassword (Read-Host "Mot de passe DSRM" -AsSecureString) `
  -Force
```

**Options utiles** :

```powershell
# Installer depuis un média (IFM) pour les liaisons lentes : réplication initiale allégée
ntdsutil "activate instance ntds" "ifm" "create sysvol full C:\IFM" quit quit
# Puis :
Install-ADDSDomainController -DomainName "ad.entreprise.fr" `
  -InstallationMediaPath "C:\IFM" -Credential (Get-Credential) -Force
```

**Post-installation** :

1. Vérifier la réplication : `repadmin /showrepl`, `repadmin /replsummary`.
2. Vérifier le DNS : le nouveau DC doit apparaître dans la zone `_msdcs`.
3. Le passer en GC si besoin (coché par défaut).
4. Mettre à jour les clients/serveurs DHCP pour inclure son IP en DNS secondaire.
5. Vérifier SYSVOL partagé : `net share`, `dcdiag /test:netlogons`.

---

## 11. Rétrograder un contrôleur de domaine proprement

**Jamais** de simple extinction ou de suppression de VM : toujours **rétrograder** (demote) pour nettoyer les métadonnées.

```powershell
# Transférer d'abord les rôles FSMO si le DC en détient (voir section 34)
Move-ADDirectoryServerOperationMasterRole -Identity "DC02" -OperationMasterRole 0,1,2,3,4

# Rétrograder (le serveur devient membre simple)
Uninstall-ADDSDomainController `
  -DemoteOperationMasterRole `
  -RemoveDnsDelegation:$false `
  -Force `
  -Credential (Get-Credential "ENTREPRISE\administrateur")
# Le serveur redémarre. Le rôle AD DS peut ensuite être désinstallé :
Uninstall-WindowsFeature AD-Domain-Services -IncludeManagementTools
```

⚠️ **Dernier DC du domaine** : `Uninstall-ADDSDomainController` avec `-LastDomainControllerInDomain` ; **dernier DC de la forêt** : `-LastDomainInForest`. Ces cas détruisent le domaine/la forêt : triple vérification + sauvegarde au préalable.

**Vérifications après rétrogradation** :

```powershell
# Le DC ne doit plus apparaître
Get-ADDomainController -Filter *
# Nettoyer les enregistrements DNS résiduels (_msdcs, zone directe)
# Vérifier la réplication sur les DC restants
repadmin /replsummary
```

---

## 12. Retrait forcé d'un DC mort (metadata cleanup)

Quand un DC est **mort sans rétrogradation** (crash disque, VM supprimée), ses métadonnées polluent l'annuaire. Nettoyage :

```powershell
# 1. Identifier le DC mort
Get-ADDomainController -Filter * | Select-Object Name, Enabled

# 2. Supprimer l'objet serveur (fait le metadata cleanup moderne)
# Via Sites et services AD : supprimer l'objet serveur -> confirme la suppression des paramètres NTDS
# Ou en PowerShell :
Remove-ADDomainController -Identity "DC03" -Force
# (échoue si le DC est encore joignable : c'est une sécurité)

# 3. Saisir les rôles FSMO qu'il détenait éventuellement (voir section 35)
# 4. Nettoyer le DNS : supprimer les enregistrements A, NS et _msdcs du DC mort
# 5. Vérifier qu'il ne reste rien :
Get-ADObject -Filter { Name -eq "DC03" }
repadmin /replsummary
dcdiag /v
```

L'ancienne méthode `ntdsutil > metadata cleanup` existe toujours mais `Remove-ADDomainController` et la suppression via la console font le travail proprement depuis 2008 R2.

---

## 13. RODC : contrôleur de domaine en lecture seule

Un **RODC** (Read-Only Domain Controller) héberge une copie **non inscriptible** de la base AD. Cas d'usage typiques :

- **Agence / site distant** sans local technique sécurisé (vol du serveur = pas de base inscriptible à exploiter).
- **DMZ / réseau exposé** : authentifier sans exposer un DC inscriptible.
- **Contrainte réglementaire** : séparer strictement l'administration.

**Caractéristiques** :

- Ne réplique qu'en **entrant** (jamais de réplication sortante).
- Par défaut, **ne met en cache aucun mot de passe** : chaque authentification est relayée vers un DC inscriptible (sauf si la stratégie de réplication de mot de passe l'autorise).
- Ne détient **aucun rôle FSMO**, n'est jamais GC par défaut (peut le devenir).
- Les modifications locales (changement de mot de passe sur site) sont relayées vers un DC inscriptible.

```powershell
# Promouvoir un RODC (le compte doit être pré-créé ou droits suffisants)
Install-ADDSDomainController `
  -DomainName "ad.entreprise.fr" `
  -ReadOnlyReplica `
  -SiteName "Site-Agence-Lyon" `
  -Credential (Get-Credential) `
  -SafeModeAdministratorPassword (Read-Host "DSRM" -AsSecureString) `
  -Force

# Stratégie de réplication des mots de passe : autoriser la mise en cache
# pour un groupe (ex. utilisateurs de l'agence)
Add-ADDomainControllerPasswordReplicationPolicy -Identity "RODC-Lyon" `
  -AllowedList "GG_Utilisateurs_Lyon"

# Voir qui a un mot de passe en cache sur le RODC
Get-ADDomainControllerPasswordReplicationPolicy -Identity "RODC-Lyon" -RevealedList
Get-ADDomainControllerPasswordReplicationPolicyUsage -Identity "RODC-Lyon" -RevealedAccounts
```

**Bonnes pratiques RODC** :

