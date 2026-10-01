---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-4
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [471, 636]
sha256: ec8331c78ef2c57e2e68450d6d39ccc9ef6b089de03387756a3974c26e7811db
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

- Pré-créez le compte RODC et définissez la stratégie de mot de passe **avant** la promotion.
- **Refusez** explicitement les comptes sensibles (admins du domaine, comptes de service critiques) via `DeniedList`.
- Placez le RODC dans un **site AD dédié** avec le sous-réseau de l'agence.
- Si un RODC est **volé/compromis** : réinitialisez les mots de passe de tous les comptes dont le secret était en cache + supprimez le compte RODC.

---

## 14. Scénario agence : déployer un RODC en site distant

Contexte : agence de 25 personnes à Lyon, liaison VPN vers Paris, pas de local sécurisé.

```powershell
# 1. Créer le site et le sous-réseau (sur un DC inscriptible à Paris)
New-ADReplicationSite -Name "Site-Lyon"
New-ADReplicationSubnet -Name "10.10.20.0/24" -Site "Site-Lyon"

# 2. Créer une liaison de site vers le site principal
New-ADReplicationSiteLink -Name "Lien-Paris-Lyon" `
  -SitesIncluded "Site-Paris","Site-Lyon" -Cost 100 `
  -ReplicationFrequencyInMinutes 60

# 3. Pré-créer le compte RODC avec délégation à un technicien local
Add-ADDomainControllerAccount -DomainControllerAccountName "RODC-Lyon" `
  -DomainName "ad.entreprise.fr" -SiteName "Site-Lyon" `
  -AllowPasswordReplicationAccountName "GG_Utilisateurs_Lyon"

# 4. Sur le serveur de l'agence (joint au domaine), promouvoir en RODC
Install-ADDSDomainController -DomainName "ad.entreprise.fr" -ReadOnlyReplica `
  -SiteName "Site-Lyon" -UseExistingAccount -Credential (Get-Credential) -Force

# 5. Configurer le DNS des postes de Lyon : RODC en primaire, DC Paris en secondaire
```

**Checklist agence** : DHCP local distribuant le RODC en DNS → test d'ouverture de session câble débranché du VPN (échoue si mot de passe non en cache : normal) → puis avec VPN (met en cache) → re-test sans VPN (réussit grâce au cache).

---

## 15. Design des unités d'organisation (OU)

Les OU servent à **trois** choses : appliquer des GPO, déléguer l'administration, structurer l'annuaire. Elles ne sont **pas** une frontière de sécurité (contrairement aux domaines).

**Principes** :

1. **Concevez pour les GPO d'abord** : une OU = un périmètre de stratégie homogène.
2. **Peu profondes** : 2 à 4 niveaux suffisent. Au-delà, l'ordre d'application devient illisible.
3. **Stables** : évitez de réorganiser les OU tous les 6 mois (les GPO liées suivent, mais les scripts et la doc cassent).
4. **Nommage clair** : pas d'accents ni de caractères spéciaux dans les noms d'OU (compatibilité scripts).

**Structure type recommandée** (entreprise multi-sites) :

```
DC=ad,DC=entreprise,DC=fr
├── OU=Admin                    # comptes admin, groupes sensibles
├── OU=Postes
│   ├── OU=Paris                # GPO spécifiques site
│   │   ├── OU=Fixes
│   │   └── OU=Portables
│   └── OU=Lyon
├── OU=Utilisateurs
│   ├── OU=Paris
│   │   ├── OU=Direction
│   │   ├── OU=Exploitation
│   │   └── OU=Maintenance
│   └── OU=Lyon
├── OU=Serveurs
│   ├── OU=ControleursDomaine   # OU par défaut, GPO renforcée
│   ├── OU=Applicatifs
│   └── OU=Hyperviseurs
└── OU=Groupes
```

> 💡 Séparez **utilisateurs** et **ordinateurs** dans des OU distinctes : les GPO « configuration utilisateur » et « configuration ordinateur » ne s'appliquent proprement que si les objets sont bien rangés (sauf bouclage, section 61).

---

## 16. Bonnes pratiques de structure OU : par site, par département

| Approche | Avantages | Inconvénients |
|---|---|---|
| **Par site** (Paris/Lyon) | GPO réseau/imprimantes locales simples, délégation au technicien local | Départements éclatés sur plusieurs OU |
| **Par département** (Compta/Atelier) | GPO métier homogènes | Sites éclatés, imprimantes par site compliquées |
| **Mixte** (recommandée) | Site au niveau 1, département au niveau 2 (ou inverse selon le besoin dominant) | Un peu plus profonde |

**Règle pratique** : si vos GPO diffèrent surtout par **site** (imprimantes, proxy, Wi-Fi), mettez le site au niveau 1. Si elles diffèrent surtout par **métier** (logiciels, droits), mettez le département au niveau 1. Le filtrage de sécurité et le ciblage GPP (sections 59, 64) compensent le reste.

```powershell
# Créer une arborescence d'OU complète
$ous = @(
  "OU=Postes,DC=ad,DC=entreprise,DC=fr",
  "OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr",
  "OU=Fixes,OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr",
  "OU=Portables,OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr",
  "OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr",
  "OU=Paris,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr"
)
foreach ($ou in $ous) {
  $name = ($ou -split ",")[0] -replace "OU=",""
  $path = ($ou -split ",",2)[1]
  New-ADOrganizationalUnit -Name $name -Path $path -ProtectedFromAccidentalDeletion $true
}
```

---

## 17. Protection contre la suppression accidentelle des OU

Chaque OU (et beaucoup d'objets) porte un drapeau **« Protéger l'objet contre une suppression accidentelle »**. C'est votre ceinture de sécurité n°1.

```powershell
# Vérifier la protection d'une OU
Get-ADOrganizationalUnit -Identity "OU=Postes,DC=ad,DC=entreprise,DC=fr" |
  Select-Object Name, ProtectedFromAccidentalDeletion

# Protéger toutes les OU de premier niveau d'un coup
Get-ADOrganizationalUnit -Filter * -SearchBase "DC=ad,DC=entreprise,DC=fr" -SearchScope OneLevel |
  Set-ADOrganizationalUnit -ProtectedFromAccidentalDeletion $true

# Pour supprimer une OU protégée : d'abord retirer la protection
Set-ADOrganizationalUnit -Identity "OU=Test,DC=ad,DC=entreprise,DC=fr" `
  -ProtectedFromAccidentalDeletion $false
Remove-ADOrganizationalUnit -Identity "OU=Test,DC=ad,DC=entreprise,DC=fr" -Recursive -Confirm:$false
```

> ⚠️ Erreur classique n°1 (voir section 78) : tenter de supprimer une OU et obtenir « accès refusé » → c'est la protection, pas un problème de droits. Et inversement : **activez-la partout**, y compris sur les OU créées par script.

---

## 18. Délégation d'administration

La délégation donne à un groupe des droits **limités** sur une OU, sans en faire des admins du domaine.

**Cas typiques** :

- Techniciens d'agence : réinitialiser les mots de passe de leur site uniquement.
- Équipe helpdesk : déverrouiller les comptes, sans toucher aux groupes sensibles.
- Responsable applicatif : gérer les groupes d'une application.

```powershell
# Méthode 1 : assistant graphique (clic droit sur l'OU > Délégation de contrôle)
# Méthode 2 : PowerShell — ex. permettre au helpdesk de réinitialiser les mots de passe
$ou = "OU=Paris,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr"
$group = Get-ADGroup "GG_Helpdesk"
# Droits : Reset Password + Change Password sur les objets user descendants
dsacls $ou /G "$($group.SamAccountName):CA;Reset Password;user"
```

**Délégations prédéfinies utiles** (assistant) : « Créer, supprimer et gérer des comptes d'utilisateurs », « Réinitialiser les mots de passe », « Gérer les liaisons de stratégie de groupe ».

**Bonnes pratiques** :

- Déléguez à des **groupes**, jamais à des individus.
- Documentez chaque délégation (qui, quoi, où) — un tableau dans votre wiki.
- **Ne déléguez jamais** sur l'OU `Domain Controllers` ni sur `AdminSDHolder`-protégés sans raison impérieuse.
- Auditez régulièrement : script qui liste les ACL non standard sur les OU.

```powershell
# Lister les entrées de contrôle d'accès d'une OU
(Get-Acl "AD:\OU=Paris,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr").Access |
  Where-Object { $_.IdentityReference -notlike "*SYSTEM*" -and $_.IdentityReference -notlike "*Administrateurs*" } |
  Format-Table IdentityReference, ActiveDirectoryRights -AutoSize
```

---

## 19. Créer des utilisateurs : ADAC vs PowerShell

