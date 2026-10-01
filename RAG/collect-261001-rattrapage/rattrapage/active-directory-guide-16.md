---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-16
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [2322, 2405]
sha256: b3b5786a5201be774add25b736946156fcdfd1dfbe1d1ae73ab3a67e476ef272
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

### Erreur n°10 — GPO « fourre-tout » modifiée sans sauvegarde
**Symptôme** : un changement dans une GPO géante casse 300 postes ; impossible de revenir en arrière précisément.
**Cause** : pas de `Backup-GPO` avant modification, pas de commentaire, pas d'OU pilote.
**Remède** : restaurer depuis la dernière sauvegarde GPO (`Restore-GPO`) ; sinon, corriger à la main paramètre par paramètre.
**Prévention** : une GPO = un objectif ; `Backup-GPO` systématique avant changement ; commentaire obligatoire ; déploiement pilote.

### Erreur n°11 — Bouclage (loopback) oublié sur un serveur RDS
**Symptôme** : les GPO utilisateur « ne s'appliquent pas » sur le RDS alors que tout est correct côté utilisateur.
**Cause** : le bouclage en mode Remplacement ignore les GPO utilisateur de l'utilisateur ; en mode Fusion, l'ordre de précédence surprend.
**Remède** : vérifier le paramètre `UserPolicyMode` dans `gpresult /r` (section « Stratégie de groupe ») ; ajuster Fusion/Remplacement selon le besoin (section 61).
**Prévention** : documenter les serveurs en bouclage ; GPO dédiée `SRV-RDS - Bouclage` bien nommée.

### Erreur n°12 — Filtre WMI qui ralentit toutes les ouvertures de session
**Symptôme** : logons de 2-3 minutes après l'ajout d'un filtre WMI « anodin ».
**Cause** : requête WMI lente ou évaluée sur des milliers de postes (ex. `Win32_Product`, notoirement lent).
**Remède** : remplacer par un filtrage de sécurité (groupes) ; si WMI indispensable, requête minimale sur `Win32_OperatingSystem`/`Win32_ComputerSystem`.
**Prévention** : tester le temps d'évaluation (`Measure-Command { Get-WmiObject -Query ... }`) ; bannir `Win32_Product`.

### Erreur n°13 — Stratégie de mot de passe mise dans une GPO d'OU (sans effet)
**Symptôme** : « j'ai mis 14 caractères minimum dans la GPO de l'OU Compta, ça ne marche pas ».
**Cause** : les paramètres de stratégie de mot de passe/verrouillage ne s'appliquent que depuis la **Default Domain Policy** (ou GPO liée à la racine avec précédence). Une GPO d'OU est ignorée pour ces paramètres.
**Remède** : configurer dans la Default Domain Policy ; pour différencier par population, utiliser les **FGPP/PSO** (section 46).
**Prévention** : le savoir (c'est le but de cette fiche !) ; documenter les PSO existants.

### Erreur n°14 — DC laissé sans sauvegarde d'état système pendant des mois
**Symptôme** : crash disque un lundi matin ; dernière sauvegarde : 4 mois ; la restauration faisant autorité est impossible (tombstone lifetime dépassée ? non, mais perte de données).
**Cause** : tâche `wbadmin` jamais planifiée ou en échec silencieux (alerte non configurée).
**Remède** : reconstruire un DC par promotion (si d'autres DC sains) ; sinon restauration + procédure de crise.
**Prévention** : `wbadmin enable backup -systemstate -schedule`, supervision du journal `Microsoft-Windows-Backup`, test de restauration semestriel en labo.

### Erreur n°15 — Sous-réseaux non déclarés dans Sites et services
**Symptôme** : clients d'une agence authentifiés sur un DC distant (ouvertures lentes), `nltest /dsgetsite` retourne `Default-First-Site-Name`.
**Cause** : le sous-réseau de l'agence n'est pas déclaré/associé au bon site.
**Remède** : créer le sous-réseau et l'associer au site (section 38) ; `nltest /dsgetsite` pour vérifier ; `gpupdate`.
**Prévention** : checklist « nouveau site / nouveau VLAN » incluant systématiquement la déclaration AD ; audit trimestriel des sous-réseaux vs réalité réseau.

### Erreur n°16 — Compte de service avec mot de passe en dur qui expire
**Symptôme** : une application tombe en panne tous les 42/90 jours ; « on change le mot de passe du compte svc_appli partout ».
**Cause** : compte utilisateur classique utilisé comme compte de service, mot de passe avec expiration, documenté nulle part.
**Remède** : migrer vers un **gMSA** (section 27) ; à défaut, PSO dédié + procédure de rotation documentée.
**Prévention** : interdire les comptes de service « utilisateur » ; inventaire des comptes de service ; gMSA par défaut pour tout nouveau besoin.

### Erreur n°17 — Suppression d'un groupe utilisé dans des ACL sans vérifier
**Symptôme** : après suppression d'un groupe « qui ne servait plus », des accès partages/applications cassés.
**Cause** : le groupe était dans des ACL NTFS ou des groupes imbriqués ; la suppression est immédiate et la corbeille ne restaure pas toujours les usages externes.
**Remède** : restaurer le groupe (corbeille), vérifier les ACL ; à défaut, recréer et réassigner (nouveau SID = tout refaire).
**Prévention** : avant suppression, vérifier `memberOf` inverse et les ACL des partages critiques ; désactiver/renommer d'abord (« _OBSOLETE »), supprimer 3 mois après.

### Erreur n°18 — Mise à jour Windows simultanée de tous les DC
**Symptôme** : après un week-end de patching, plus aucun DC joignable pendant les reboots simultanés ; authentifications en échec.
**Cause** : tous les DC redémarrés en même temps (WSUS/GPO mal réglée).
**Remède** : ne jamais patcher les DC en parallèle : **un par un**, en vérifiant `dcdiag`/`repadmin` entre chaque.
**Prévention** : groupes de maintenance étalés ; le PDC en dernier ; fenêtre de maintenance dédiée aux DC ; procédure écrite.

---

## 79. Cas pratiques commentés (16)

### Cas n°1 — Créer une arborescence complète pour une nouvelle agence
**Contexte** : ouverture de l'agence de Lyon (30 utilisateurs, 25 PC, 1 imprimante).
**Objectif** : OU, groupes, GPO de base, tout par script pour reproductibilité.
```powershell
$dom = "DC=ad,DC=entreprise,DC=fr"
# OU
New-ADOrganizationalUnit -Name "Lyon" -Path "OU=Utilisateurs,$dom" -ProtectedFromAccidentalDeletion $true
New-ADOrganizationalUnit -Name "Lyon" -Path "OU=Postes,$dom" -ProtectedFromAccidentalDeletion $true
# Groupes
New-ADGroup "GG_Utilisateurs_Lyon" -GroupScope Global -GroupCategory Security -Path "OU=Groupes,$dom"
New-ADGroup "GG_PC_Lyon" -GroupScope Global -GroupCategory Security -Path "OU=Groupes,$dom"
# GPO imprimante, filtrée sur les PC de Lyon
New-GPO -Name "ORDI - Lyon - Imprimante atelier" -Comment "Déploiement imprimante Lyon"
New-GPLink -Name "ORDI - Lyon - Imprimante atelier" -Target "OU=Lyon,OU=Postes,$dom"
Set-GPPermissions -Name "ORDI - Lyon - Imprimante atelier" -TargetName "Utilisateurs authentifiés" -TargetType Group -PermissionLevel None
Set-GPPermissions -Name "ORDI - Lyon - Imprimante atelier" -TargetName "Ordinateurs du domaine" -TargetType Group -PermissionLevel GpoRead
Set-GPPermissions -Name "ORDI - Lyon - Imprimante atelier" -TargetName "GG_PC_Lyon" -TargetType Group -PermissionLevel GpoApply
```
**Commentaire** : tout est scripté → l'agence suivante se déploie en changeant 3 variables. La GPO imprimante est configurée ensuite dans la GPMC (GPP > Imprimantes partagées).

### Cas n°2 — Intégrer 40 intérimaires en une matinée
**Contexte** : pic d'activité, 40 intérimaires arrivent lundi 8h.
**Objectif** : comptes créés, expirant dans 30 jours, mot de passe initial, groupes par atelier.
**Solution** : CSV + script de la section 20, avec `-AccountExpirationDate (Get-Date).AddDays(30)`. Imprimer la liste login/mot de passe initiaux, remise en main propre par le chef d'atelier.
**Commentaire** : l'expiration automatique évite les comptes fantômes. Prévoir le script inverse (désactivation) — section 22.

