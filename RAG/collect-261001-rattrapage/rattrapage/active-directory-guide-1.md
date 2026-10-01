---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-1
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "attribution"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1, 144]
sha256: 9a4798235804222681ed886c686068bd3ff745012252ae78e8ef3337b6d25707
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

> **Public** : Zelef, chef de service systèmes & énergies.
> **Objectif** : tout maîtriser d'Active Directory Domain Services (AD DS) et des stratégies de groupe (GPO) en environnement d'entreprise, de l'architecture au dépannage.
> **Versions couvertes** : Windows Server 2019, 2022 et 2025.
> **Ton** : direct, dense, pratique. Chaque section = théorie + commandes + pièges à éviter.
> **Avertissement global** : testez toujours les commandes destructrices (suppression, restauration, FSMO) dans un labo avant la production.

---

## Sommaire

1. Introduction : pourquoi Active Directory ?
2. Vocabulaire et concepts fondamentaux
3. Forêt, arbre, domaine : architecture logique
4. Approbations (trusts) entre domaines et forêts
5. Catalogue global (Global Catalog)
6. Partitions d'annuaire
7. Contrôleurs de domaine : rôles et architecture
8. Prérequis avant de promouvoir un contrôleur de domaine
9. Promouvoir le premier DC d'une nouvelle forêt (PowerShell)
10. Ajouter un DC supplémentaire dans un domaine existant
11. Rétrograder un contrôleur de domaine proprement
12. Retrait forcé d'un DC mort (metadata cleanup)
13. RODC : contrôleur de domaine en lecture seule
14. Scénario agence : déployer un RODC en site distant
15. Design des unités d'organisation (OU)
16. Bonnes pratiques de structure OU : par site vs par département
17. Protection contre la suppression accidentelle des OU
18. Délégation d'administration
19. Créer des utilisateurs : ADAC vs PowerShell
20. Création en masse via CSV (New-ADUser)
21. Attributs utilisateur importants et bonnes pratiques
22. Désactiver, déplacer, supprimer des comptes : cycle de vie
23. Groupes : portées (global, universel, local de domaine)
24. Imbrication de groupes : stratégie AGDLP/AGUDLP
25. Groupes dynamiques : état des lieux et alternatives
26. Contacts, ordinateurs et comptes de service
27. Comptes de service gérés (gMSA)
28. Les 5 rôles FSMO : vue d'ensemble
29. Maître de schéma (Schema Master)
30. Maître d'attribution des noms de domaine (Domain Naming Master)
31. Maître RID (Relative ID Master)
32. Émulateur PDC (PDC Emulator)
33. Maître d'infrastructure (Infrastructure Master)
34. Transférer un rôle FSMO (Move-ADDirectoryServerOperationMasterRole)
35. Saisir un rôle FSMO (seize via ntdsutil) — procédure d'urgence
36. Bonnes pratiques de placement des rôles FSMO
37. Sites et services : concepts
38. Créer et configurer des sites, sous-réseaux
39. Liaisons de sites : coût, planification, bridgehead
40. Réplication AD : KCC et topologie
41. Surveiller la réplication : repadmin
42. Diagnostiquer avec dcdiag
43. Dépannage réplication : USN rollback
44. Dépannage réplication : lingering objects (objets rémanents)
45. Stratégie de mot de passe par défaut du domaine
46. FGPP : stratégies de mot de passe affinées (PSO)
47. Bonnes pratiques mots de passe et verrouillage de compte
48. Corbeille AD : activation et fonctionnement
49. Restaurer un objet supprimé (Restore-ADObject)
50. Sauvegarde : état système avec wbadmin
51. Sauvegarder via ntdsutil et bonnes pratiques de sauvegarde
52. DSRM : mode de restauration des services d'annuaire
53. Restauration non faisant autorité vs faisant autorité
54. Migration inter-forêts avec ADMT
55. Mise à niveau du niveau fonctionnel (forêt et domaine)
56. GPO : principes fondamentaux
57. Ordre d'application : LSDOU
58. Héritage, blocage d'héritage et application forcée (Enforced)
59. Filtrage de sécurité des GPO
60. Filtrage WMI
61. Bouclage de traitement (loopback) : fusion vs remplacement
62. Préférences de stratégie de groupe (GPP)
63. GPP : lecteurs réseau, imprimantes, registre, raccourcis
64. Ciblage au niveau élément (Item-Level Targeting)
65. GPO de sécurité : audit et droits utilisateur
66. Déploiement de logiciels via GPO (MSI)
67. Scripts de démarrage / arrêt / ouverture / fermeture de session
68. Dépannage GPO : gpresult, rsop.msc, journaux
69. GPO qui ne s'applique pas : méthodologie complète
70. Bonnes pratiques de nommage et gestion du cycle de vie des GPO
71. Santé du domaine : checklist quotidienne
72. Santé du domaine : checklist hebdomadaire et mensuelle
73. Scripts de rapport automatisés
74. Supervision : compteurs et alertes critiques
75. Pense-bête de poche : commandes essentielles
76. Glossaire
77. Quiz : 10 questions + réponses
78. Erreurs classiques (18)
79. Cas pratiques commentés (16)
80. Pour aller plus loin
81. Checklists de mise en production
82. Annexe : versions et niveaux fonctionnels

---

## 1. Introduction : pourquoi Active Directory ?

Active Directory Domain Services (AD DS) est le service d'annuaire de Microsoft. En entreprise, il centralise :

- **L'authentification** : un seul identifiant pour ouvrir sa session, accéder aux partages, aux imprimantes, aux applications (SSO via Kerberos).
- **L'autorisation** : qui a le droit de faire quoi, via les groupes de sécurité.
- **La gestion centralisée** : stratégies de groupe (GPO) pour configurer des milliers de postes depuis un point unique.
- **L'inventaire** : chaque utilisateur, ordinateur, imprimante est un objet avec des attributs interrogeables.

Sans AD, chaque poste gère ses comptes locaux : cauchemar dès 20+ machines. Avec AD, un départ = un compte désactivé, et l'accès est coupé partout.

**Ce que ce guide n'est pas** : un remplacement pour la documentation Microsoft ni pour un labo. Chaque commande destructrice doit être testée hors production.

**Versions** : les cmdlets PowerShell du module `ActiveDirectory` sont stables depuis 2012 R2. Les différences 2019/2022/2025 concernent surtout les niveaux fonctionnels, la sécurité (LDAP signing, SMB) et quelques nouveautés (voir section 82). Quand une commande dépend de la version, c'est signalé par ⚠️.

---

## 2. Vocabulaire et concepts fondamentaux

| Terme | Définition |
|---|---|
| **Annuaire** | Base de données hiérarchique stockant des objets (utilisateurs, ordinateurs, groupes...) et leurs attributs. Fichier physique : `C:\Windows\NTDS\ntds.dit`. |
| **Objet** | Élément de l'annuaire (utilisateur, groupe, OU, ordinateur, imprimante...). |
| **Attribut** | Propriété d'un objet (`sAMAccountName`, `userPrincipalName`, `memberOf`...). |
| **Schéma** | Définition des classes d'objets et attributs autorisés. Commun à toute la forêt. |
| **DN (Distinguished Name)** | Chemin unique d'un objet : `CN=Zelef,OU=Utilisateurs,DC=entreprise,DC=lan`. |
| **CN / OU / DC** | Common Name, Organizational Unit, Domain Component (composants du DN). |
| **UPN** | User Principal Name, format `prenom.nom@entreprise.lan` — ce que l'utilisateur tape à l'ouverture de session. |
| **sAMAccountName** | Nom de connexion pré-Windows 2000, format `ENTREPRISE\zelef`, limité à 20 caractères. |
| **SID** | Identifiant de sécurité unique (`S-1-5-21-...`). Les droits portent sur les SID, pas sur les noms. |
| **Forêt** | Ensemble de domaines partageant le même schéma et le même catalogue global. Frontière de sécurité. |
| **Domaine** | Frontière de réplication et d'administration (stratégies de mot de passe, etc.). |

**Point clé à retenir** : les autorisations NTFS et les appartenances aux groupes sont liées aux **SID**, pas aux noms. Renommer un utilisateur ne casse rien ; supprimer puis recréer un compte (nouveau SID) casse tout.

```powershell
# Voir le SID d'un utilisateur
Get-ADUser zelef | Select-Object Name, SID

# Voir le DN
Get-ADUser zelef | Select-Object DistinguishedName
```

---

## 3. Forêt, arbre, domaine : architecture logique

