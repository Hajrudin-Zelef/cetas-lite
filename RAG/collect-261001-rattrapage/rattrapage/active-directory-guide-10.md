---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-10
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-15-10"]
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1513, 1672]
sha256: 23ae4761725d42f6dc40b5774ef2f6918af2c902906cdf8e1c7a120d0621d63c
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

Le **DSRM** (Directory Services Restore Mode) est le mode sans échec spécial des DC pour restaurer la base AD. Il utilise un **mot de passe local propre au DC**, défini à la promotion.

```powershell
# Changer le mot de passe DSRM d'un DC (à faire régulièrement / à chaque changement d'admin)
ntdsutil "set dsrm password" "reset password on server DC01" quit quit
# Ou via PowerShell (2016+) :
Set-DSRMPassword -NewPassword (Read-Host "Nouveau mot de passe DSRM" -AsSecureString)
```

**Bonnes pratiques** :

- Stockez les mots de passe DSRM dans le **coffre à mots de passe** d'entreprise (pas dans un fichier texte !).
- Un mot de passe DSRM **différent par DC** (ou au minimum documenté).
- Pour démarrer en DSRM : `msconfig` > Démarrer > options de démarrage, ou `bcdedit /set safeboot dsrepair`, ou F8 au boot.
- En DSRM, le DC ne fait **pas** autorité : parfait pour restaurer l'état système avant de choisir le type de restauration (section 53).

---

## 53. Restauration non faisant autorité vs faisant autorité

**Restauration non faisant autorité** (cas standard : un DC est mort, on le réinstalle) : on restaure l'état système, puis le DC **se resynchronise** avec ses partenaires (les données les plus récentes gagnent). Aucune manipulation supplémentaire.

**Restauration faisant autorité** (on veut **réimposer** des objets supprimés par erreur à toute la forêt) :

```powershell
# 1. Démarrer en DSRM, restaurer l'état système :
wbadmin start recovery -version:10/15/2026-22:00 -itemType:systemstate -quiet
# (en pratique : wbadmin get versions pour choisir la version)

# 2. Marquer les objets comme "faisant autorité" AVANT de redémarrer normalement
ntdsutil "activate instance ntds" "authoritative restore" `
  "restore object OU=Exploitation,OU=Paris,OU=Utilisateurs,DC=ad,DC=entreprise,DC=fr" `
  quit quit
# Pour toute la base (nucléaire : à éviter sauf catastrophe) :
# ntdsutil "activate instance ntds" "authoritative restore" "restore database" quit quit

# 3. Redémarrer en mode normal : les objets restaurés se répliquent en écrasant les suppressions
```

**Arbre de décision** :

- Objet(s) supprimé(s) par erreur, corbeille AD activée → `Restore-ADObject` (section 49). **C'est 99 % des cas.**
- DC mort → réinstallation + restauration non faisant autorité (ou simplement promouvoir un nouveau DC).
- Suppression massive répliquée partout, pas de corbeille → restauration faisant autorité depuis une sauvegarde antérieure à la suppression.

---

## 54. Migration inter-forêts avec ADMT

**ADMT** (Active Directory Migration Tool) migre utilisateurs, groupes, ordinateurs et profils d'une forêt source vers une forêt cible. Scénarios : fusion d'entreprises, rachat, refonte de forêt.

**Prérequis** :

1. Approbation de forêt bidirectionnelle (section 4).
2. DNS : chaque forêt résout l'autre.
3. ADMT installé sur un serveur de la forêt **cible** (+ SQL Server Express pour sa base).
4. Clé de migration des mots de passe (`admt key`) si on migre les mots de passe (nécessite un DC source 2016+ et la DLL Password Export Server).

**Étapes** :

```powershell
# 1. Sur la source : autoriser la migration (une fois)
#    Registre : HKLM\SYSTEM\CurrentControlSet\Control\Lsa\TcpipClientSupport = 1 (DWORD)
#    + redémarrage du DC source (désactive le filtrage SID pour la migration)

# 2. Migrer les comptes de service et groupes globaux d'abord, puis les utilisateurs
#    (via l'assistant ADMT : "User Account Migration Wizard")

# 3. Migrer les postes (avec traduction des profils locaux)
#    L'agent ADMT se déploie ; le poste redémarre et rejoint le nouveau domaine

# 4. Conserver le SID historique (SID History) pour garder l'accès aux ressources source
#    Option "Migrate user SIDs to target domain" dans l'assistant
```

**Ordre de migration** : groupes globaux → utilisateurs → postes → ressources. **Ne migrez jamais un DC** avec ADMT (on promeut des DC natifs dans la cible).

⚠️ ADMT n'est plus mis à jour par Microsoft depuis longtemps ; pour les migrations modernes, évaluez aussi les outils tiers (Quest, Binary Tree) selon la complexité.

---

## 55. Mise à niveau du niveau fonctionnel (forêt et domaine)

Le **niveau fonctionnel** détermine les fonctionnalités AD disponibles. Tous les DC doivent être au minimum à ce niveau d'OS.

| Niveau | DC minimum requis | Notes |
|---|---|---|
| 2012 R2 | 2012 R2 | Corbeille, etc. |
| 2016 (`WinThreshold`) | 2016 | Dernier niveau « classique » |
| 2025 (`Win2025`) ⚠️ | 2022+ (2025 pour certaines fonctions) | Nouveau avec Server 2025 : pages 32K dans ntds.dit, NUMA, etc. |

```powershell
# Voir les niveaux actuels
Get-ADForest | Select-Object ForestMode
Get-ADDomain | Select-Object DomainMode

# Monter le niveau du domaine (irréversible !)
Set-ADDomainMode -Identity "ad.entreprise.fr" -DomainMode "WinThreshold"

# Monter le niveau de la forêt (après tous les domaines)
Set-ADForestMode -Identity "ad.entreprise.fr" -ForestMode "WinThreshold"
```

**Procédure de montée de version des DC** (ex. 2019 → 2022) :

1. `dcdiag /e`, `repadmin /replsummary` : santé parfaite.
2. Sauvegarde état système.
3. Promouvoir un nouveau DC sous le nouvel OS, transférer les FSMO, vérifier, rétrograder l'ancien.
4. Ne monter le niveau fonctionnel qu'une fois **tous** les DC au nouvel OS (irréversible).

---

## 56. GPO : principes fondamentaux

Une **GPO** (Group Policy Object) = un ensemble de paramètres appliqués aux utilisateurs et/ou ordinateurs. Stockage double :

- **Conteneur de stratégie (GPC)** dans AD (`CN=Policies,CN=System,...`) : métadonnées, version.
- **Modèle de stratégie (GPT)** dans SYSVOL (`\\domaine\SYSVOL\domaine\Policies\{GUID}`) : les fichiers réels (Registry.pol, scripts...).

**Deux moitiés** : **Configuration ordinateur** (appliquée au démarrage, contexte SYSTEM) et **Configuration utilisateur** (appliquée à l'ouverture de session).

**Bonnes pratiques structurelles** :

- **Une GPO = un objectif** (« Durcissement navigateurs », « Imprimantes atelier »). Jamais de GPO « fourre-tout ».
- Nommez explicitement : `ORDI - Fond d'écran entreprise`, `USER - Mappage lecteurs`, avec préfixe du périmètre.
- **Désactivez la moitié inutilisée** (ex. une GPO 100 % « configuration ordinateur » → désactiver la config utilisateur) : accélère le traitement.
- Documentez chaque GPO dans son commentaire (GPMC > Détails > Commentaire).

```powershell
# Lister les GPO du domaine
Get-GPO -All | Select-Object DisplayName, GpoStatus, ModificationTime |
  Sort-Object DisplayName

# Créer une GPO et la lier à une OU
New-GPO -Name "ORDI - Verrouillage session 10 min" -Comment "Sécurité : verrouillage automatique"
New-GPLink -Name "ORDI - Verrouillage session 10 min" `
  -Target "OU=Fixes,OU=Paris,OU=Postes,DC=ad,DC=entreprise,DC=fr"
```

---

## 57. Ordre d'application : LSDOU

Les GPO s'appliquent dans l'ordre **LSDOU**, chaque niveau **écrasant** le précédent en cas de conflit :

1. **L**ocal (stratégie locale du poste, `gpedit.msc`)
2. **S**ite (GPO liée au site AD)
3. **D**omaine (GPO liée au domaine, ex. Default Domain Policy)
4. **OU** (de la racine vers l'OU la plus proche de l'objet — la plus proche gagne)

```
Site-Paris (GPO site)
└── Domaine ad.entreprise.fr (Default Domain Policy, GPO domaine)
    └── OU=Postes (GPO "Postes - Base")
        └── OU=Paris (GPO "Paris - Imprimantes")
            └── OU=Fixes  ← le PC est ici : cette GPO gagne en cas de conflit
```

**Exceptions** : **Blocage d'héritage** (section 58), **Application forcée / Enforced** (section 58), **Bouclage** (section 61), **Filtrage de sécurité/WMI** (sections 59-60).

