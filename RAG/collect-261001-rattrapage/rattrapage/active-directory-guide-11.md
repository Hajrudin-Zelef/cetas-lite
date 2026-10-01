---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-11
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1673, 1824]
sha256: 9ec25ee4d1e4d46a36fedf6551e6975ba06eb3e0f8ee52cea1a41fb000a9252d
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

```powershell
# Voir l'ordre des liens GPO sur une OU (LinkOrder : 1 = appliqué en dernier = gagne)
Get-GPInheritance -Target "OU=Fixes,OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr" |
  Select-Object -ExpandProperty GpoLinks |
  Select-Object DisplayName, Enabled, Enforced, Order
```

---

## 58. Héritage, blocage d'héritage et application forcée (Enforced)

- **Héritage** : par défaut, un objet reçoit les GPO de tous ses parents (LSDOU).
- **Blocage d'héritage** (clic droit sur l'OU > « Bloquer l'héritage ») : l'OU **ignore** les GPO des niveaux supérieurs (sauf celles marquées Enforced). À utiliser avec parcimonie : ça rend le dépannage infernal.
- **Application forcée (Enforced)** : la GPO s'applique **quoi qu'il arrive**, même avec blocage d'héritage, et **gagne** sur les conflits (elle remonte en tête de précédence).

```powershell
# Bloquer l'héritage sur une OU
Set-GPInheritance -Target "OU=DMZ,DC=ad,DC=entreprise,DC=fr" -IsBlocked Yes

# Marquer un lien comme Enforced
Set-GPLink -Name "SECU - Base durcissement" `
  -Target "DC=ad,DC=entreprise,DC=fr" -Enforced Yes

# Voir l'héritage effectif
Get-GPInheritance -Target "OU=Fixes,OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr"
```

**Règle d'or** : préférez le **filtrage de sécurité** (section 59) au blocage d'héritage. Le blocage est un marteau ; le filtrage est un scalpel.

---

## 59. Filtrage de sécurité des GPO

Par défaut, une GPO liée s'applique aux **Utilisateurs authentifiés**. Le filtrage de sécurité restreint à un groupe précis.

**Exemple** : la GPO « Imprimantes atelier » ne doit toucher que les PC de l'atelier :

```powershell
# 1. Créer le groupe (ordinateurs de l'atelier)
New-ADGroup -Name "GG_PC_Atelier" -GroupScope Global -GroupCategory Security `
  -Path "OU=Groupes,DC=ad,DC=entreprise,DC=fr"
# (ajouter les comptes d'ordinateurs dedans)

# 2. Remplacer le filtrage par défaut
Set-GPPermissions -Name "Imprimantes atelier" `
  -TargetName "Utilisateurs authentifiés" -TargetType Group -PermissionLevel None
Set-GPPermissions -Name "Imprimantes atelier" `
  -TargetName "GG_PC_Atelier" -TargetType Group -PermissionLevel GpoApply
```

⚠️ **Piège classique** (erreur n°3, section 78) : retirer « Utilisateurs authentifiés » **sans** ajouter le droit **Lecture** à un groupe. Depuis MS16-072 (2016), les **comptes d'ordinateurs** ont besoin du droit **Lecture** sur la GPO pour la traiter. Si vous filtrez, donnez au minimum `GpoRead` aux ordinateurs concernés (ou laissez « Utilisateurs authentifiés » en lecture seule + votre groupe en « Appliquer »).

```powershell
# Méthode sûre : lecture pour les ordinateurs du domaine, application pour le groupe cible
Set-GPPermissions -Name "Imprimantes atelier" -TargetName "Ordinateurs du domaine" `
  -TargetType Group -PermissionLevel GpoRead
Set-GPPermissions -Name "Imprimantes atelier" -TargetName "GG_PC_Atelier" `
  -TargetType Group -PermissionLevel GpoApply
```

---

## 60. Filtrage WMI

Le **filtrage WMI** applique une GPO uniquement si une requête WQL est vraie sur le poste (version d'OS, modèle, RAM...).

```powershell
# Créer un filtre WMI : Windows 11 uniquement
New-GPWmiFilter -Name "Filtre - Windows 11" `
  -Query 'SELECT * FROM Win32_OperatingSystem WHERE Version LIKE "10.0.226%" AND ProductType = "1"' `
  -Description "Cible Windows 11 22H2+"

# L'associer à une GPO
Set-GPO -Name "ORDI - Config Windows 11" -WmiFilter "Filtre - Windows 11"
```

**Filtres WMI utiles** :

```sql
-- Windows 10 uniquement
SELECT * FROM Win32_OperatingSystem WHERE Version LIKE "10.0.190%" AND ProductType = "1"
-- Postes avec au moins 8 Go de RAM
SELECT * FROM Win32_ComputerSystem WHERE TotalPhysicalMemory >= 8589934592
-- Modèle de PC précis (ex. déploiement pilote)
SELECT * FROM Win32_ComputerSystem WHERE Model LIKE "%Latitude 5440%"
-- Serveurs uniquement
SELECT * FROM Win32_OperatingSystem WHERE ProductType = "3"
```

⚠️ **Performance** : chaque filtre WMI est évalué à chaque traitement de stratégie. Un WMI lent ou une requête mal écrite **ralentit toutes les ouvertures de session**. Préférez le filtrage de sécurité quand c'est possible ; réservez le WMI aux critères non exprimables en groupes (version d'OS, matériel).

---

## 61. Bouclage de traitement (loopback) : fusion vs remplacement

Problème : les GPO « configuration utilisateur » suivent l'**utilisateur**, pas le poste. Mais sur un **serveur TSE/RDS** ou un **poste d'atelier partagé**, on veut que la config utilisateur dépende du **poste** (ex. verrouiller le bureau sur les PC d'atelier quel que soit l'utilisateur).

Solution : le **bouclage** (loopback), paramètre ordinateur :
`Configuration ordinateur > Stratégies > Modèles d'administration > Système > Stratégie de groupe > Configurer le mode de traitement bouclé`.

- **Fusion (Merge)** : les GPO utilisateur du poste **s'ajoutent** à celles de l'utilisateur. En cas de conflit, **celles de l'utilisateur gagnent**.
- **Remplacement (Replace)** : seules les GPO utilisateur liées au poste s'appliquent ; celles de l'utilisateur sont **ignorées**.

```powershell
# Activer le bouclage en mode Fusion via registre (équivalent du paramètre de stratégie)
# À déployer via une GPO "configuration ordinateur" liée à l'OU des serveurs RDS :
# Clé : HKLM\SOFTWARE\Policies\Microsoft\Windows\System\UserPolicyMode = 1 (Fusion) / 2 (Remplacement)
```

**Cas d'usage** : serveurs RDS (bureau verrouillé identique pour tous), postes en libre-service, salles de réunion, kiosques. **Piège** : oublier le bouclage actif quand on dépannage une GPO utilisateur « qui ne s'applique pas » sur un RDS.

---

## 62. Préférences de stratégie de groupe (GPP)

Les **préférences** (coche verte) vs les **stratégies** (policies) :

| | Stratégies | Préférences |
|---|---|---|
| Effet | **Imposées**, non modifiables par l'utilisateur (grisé) | Appliquées par défaut, **modifiables** ensuite par l'utilisateur |
| Persistance | Réappliquées, retirées si la GPO disparaît | Selon l'option (appliquer une fois, retirer si hors périmètre...) |
| Périmètre | Paramètres du système | Lecteurs, imprimantes, registre, fichiers, raccourcis, variables d'env... |

**Options communes des GPP** (onglet « Commun ») :

- **Appliquer une fois et ne pas réappliquer** : idéal pour un paramétrage initial que l'utilisateur peut changer.
- **Ciblage au niveau de l'élément** (section 64) : le vrai super-pouvoir des GPP.
- **Exécuter dans le contexte de sécurité de l'utilisateur** vs SYSTEM.

> 💡 Règle simple : **stratégie** pour la sécurité et la conformité (verrouillage, pare-feu), **préférence** pour le confort (lecteurs réseau, imprimantes par défaut, raccourcis).

---

## 63. GPP : lecteurs réseau, imprimantes, registre, raccourcis

**Lecteurs réseau** (`Configuration utilisateur > Préférences > Paramètres Windows > Mappages de lecteurs`) :

```powershell
# Équivalent script (mais préférez la GPP avec ciblage) :
# New-PSDrive -Name "S" -PSProvider FileSystem -Root "\\srv-fichiers\Commun" -Persist
```

En GPP : Nouvel élément > Lecteur mappé : lettre `S:`, chemin `\\srv-fichiers\Commun`, action **Mettre à jour**, reconnecter coché, ciblage sur le groupe `GG_Exploitation_Paris`.

**Imprimantes** (`Préférences > Paramètres du Panneau de configuration > Imprimantes`) : déployer l'imprimante partagée `\\srv-print\Atelier-HP`, la définir par défaut avec ciblage sur l'OU/le site.

**Registre** (`Préférences > Paramètres Windows > Registre`) : pousser une valeur, ex. page d'accueil du navigateur d'entreprise, avec action Mettre à jour.

**Raccourcis** (`Préférences > Paramètres Windows > Raccourcis`) : déposer un raccourci vers l'intranet sur le Bureau de tous les utilisateurs, ciblé par groupe.

**Bonnes pratiques GPP** :

