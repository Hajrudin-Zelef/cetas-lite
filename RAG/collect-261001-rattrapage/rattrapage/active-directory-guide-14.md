---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-14
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [2156, 2251]
sha256: a8bf8abba2b82ec074d97739952cd7393c3562c374ecac3b904c252aac084ed0
---

# Active Directory & GPO en entreprise — Guide technique ultra-complet

| Terme | Définition |
|---|---|
| **ACL** | Access Control List : liste des autorisations sur un objet (qui peut faire quoi). |
| **ADAC** | Centre d'administration Active Directory (`dsacls.msc`), console moderne. |
| **AD DS** | Active Directory Domain Services : le rôle serveur d'annuaire. |
| **ADMT** | Active Directory Migration Tool : migration inter-forêts. |
| **AGDLP / AGUDLP** | Stratégie d'imbrication : Comptes → Global → (Universel) → Local de Domaine → Permissions. |
| **Approbation (Trust)** | Relation de confiance entre domaines/forêts. |
| **Arbre** | Domaines de même espace de noms DNS dans une forêt. |
| **Attribut** | Propriété d'un objet AD (ex. `telephoneNumber`). |
| **Bridgehead** | DC tête de pont pour la réplication inter-sites. |
| **Catalogue global (GC)** | Index partiel de tous les objets de la forêt (ports 3268/3269). |
| **Classe d'objet** | Type défini dans le schéma (user, group, computer...). |
| **CN** | Common Name, composant du DN. |
| **Corbeille AD** | Fonctionnalité de restauration des objets supprimés avec leurs attributs. |
| **DACL / SACL** | Parties de l'ACL : autorisations (DACL) et audit (SACL). |
| **DC** | Contrôleur de domaine. |
| **DFSR** | DFS Replication : réplication de SYSVOL (remplace FRS). |
| **DN** | Distinguished Name : chemin unique d'un objet. |
| **Domaine** | Frontière de réplication et d'administration. |
| **DSRM** | Directory Services Restore Mode : mode de restauration d'un DC. |
| **Enforced** | Option de lien GPO : s'applique malgré le blocage d'héritage, prioritaire. |
| **FGPP** | Fine-Grained Password Policy : stratégies de mot de passe par groupe (PSO). |
| **Forêt** | Ensemble des domaines partageant schéma et configuration. Frontière de sécurité. |
| **FSMO** | Flexible Single Master Operations : les 5 rôles maître unique. |
| **GC** | Voir Catalogue global. |
| **GPC / GPT** | Conteneur (AD) et modèle (SYSVOL) d'une GPO. |
| **GPMC** | Console de gestion des stratégies de groupe (`gpmc.msc`). |
| **GPO** | Group Policy Object : stratégie de groupe. |
| **GPP** | Préférences de stratégie de groupe (coche verte). |
| **GUID** | Identifiant unique d'objet (ex. identifiant d'une GPO). |
| **Héritage (GPO)** | Application cumulative des GPO parentes (LSDOU). |
| **IFM** | Install From Media : promotion d'un DC depuis un média pour liaisons lentes. |
| **Infrastructure Master** | FSMO maintenant les références inter-domaines. |
| **ISTG** | Inter-Site Topology Generator : DC générant la topologie inter-sites. |
| **Item-Level Targeting** | Ciblage conditionnel des éléments GPP. |
| **KCC** | Knowledge Consistency Checker : calcule la topologie de réplication. |
| **Kerberos** | Protocole d'authentification principal d'AD (port 88). |
| **LDAP / LDAPS** | Protocole d'accès à l'annuaire (389 / 636). |
| **LDS / AD LDS** | Instance d'annuaire légère (hors domaine). |
| **Lingering object** | Objet rémanent : supprimé partout mais réinjecté par un DC trop longtemps isolé. |
| **Loopback (bouclage)** | Mode GPO où la config utilisateur dépend du poste (Fusion/Remplacement). |
| **LSDOU** | Ordre d'application : Local, Site, Domaine, OU. |
| **Maître de schéma** | FSMO autorisant les modifications du schéma (forêt). |
| **Maître RID** | FSMO distribuant les pools de RID (domaine). |
| **Metadata cleanup** | Nettoyage des métadonnées d'un DC mort. |
| **MS16-072** | Correctif imposant le droit Lecture aux ordinateurs sur les GPO filtrées. |
| **Naming Master** | FSMO d'attribution des noms de domaine (forêt). |
| **NTDS.dit** | Fichier de la base AD sur chaque DC. |
| **NTP / w32time** | Hiérarchie de temps, pilotée par l'émulateur PDC. |
| **OU** | Organizational Unit : conteneur pour GPO et délégation. |
| **PDC Emulator** | FSMO : temps, verrouillages, mots de passe urgents, édition GPO. |
| **Phantom** | Objet fantôme : référence inter-domaine maintenue par l'Infrastructure Master. |
| **PSO** | Password Settings Object : objet de stratégie FGPP. |
| **RDS** | Remote Desktop Services : bureaux distants (cas typique du bouclage). |
| **repadmin** | Outil de diagnostic de la réplication. |
| **RID** | Relative Identifier : fin du SID, alloué par pool. |
| **RODC** | Read-Only Domain Controller : DC en lecture seule. |
| **RSOP** | Resultant Set of Policy : stratégies résultantes (`rsop.msc`, `gpresult`). |
| **Schéma** | Définition des classes et attributs, commun à la forêt. |
| **Seize** | Saisie forcée d'un rôle FSMO (détenteur mort). |
| **SID** | Security Identifier : identifiant unique portant les droits. |
| **SID History** | Conservation des anciens SID lors d'une migration (ADMT). |
| **Site** | Emplacement réseau bien connecté (réplication, proximité client). |
| **Site Link** | Liaison logique entre sites (coût, fréquence). |
| **sAMAccountName** | Login court pré-Windows 2000 (`DOMAINE\login`, 20 car. max). |
| **SYSVOL** | Partage répliqué (`\\domaine\SYSVOL`) contenant les GPO et scripts. |
| **Tombstone** | Objet supprimé conservé temporairement pour la réplication (180 j). |
| **Transfert (FSMO)** | Déplacement propre d'un rôle (détenteur en ligne). |
| **UPN** | User Principal Name : `prenom.nom@domaine`, login moderne. |
| **USN** | Update Sequence Number : compteur d'écritures par DC. |
| **USN rollback** | Divergence quand l'USN d'un DC recule (restore de snapshot). |
| **WMI Filter** | Filtre GPO basé sur une requête WMI (version OS, matériel...). |

---

## 77. Quiz : 10 questions + réponses

**Q1. Dans quel ordre s'appliquent les GPO ?**
<details><summary>Réponse</summary>LSDOU : Local → Site → Domaine → OU (de la plus éloignée à la plus proche de l'objet). En cas de conflit, la dernière appliquée gagne, sauf lien Enforced ou blocage d'héritage.</details>

**Q2. Un utilisateur change de service. Selon AGDLP, que modifiez-vous ?**
<details><summary>Réponse</summary>Uniquement son appartenance aux groupes **globaux** (ex. le retirer de `GG_Compta` et l'ajouter à `GG_Exploitation`). Les groupes locaux de domaine et les ACL sur les ressources ne changent pas.</details>

**Q3. Le détenteur du rôle PDC tombe en panne un vendredi soir. Quelles conséquences immédiates ?**
<details><summary>Réponse</summary>La référence de temps du domaine est perdue (dérive possible), les verrouillages de compte ne sont plus centralisés, l'édition des GPO via la GPMC peut poser problème (elle cible le PDC). La plupart des authentifications continuent. Si la panne est définitive : saisir le rôle (seize) sur un autre DC ; si temporaire : attendre ou transférer.</details>

**Q4. Pourquoi ne faut-il jamais restaurer un DC à partir d'un snapshot hyperviseur ?**
<details><summary>Réponse</summary>Risque d'**USN rollback** : l'USN du DC restauré recule, ses partenaires croient déjà avoir répliqué ces numéros et ignorent les écritures intermédiaires → divergence silencieuse (mots de passe, comptes). On restaure un DC via une sauvegarde d'état système, jamais par snapshot.</details>

**Q5. Une GPO liée à une OU ne s'applique pas à un PC qui est bien dans l'OU. Citez 4 vérifications.**
<details><summary>Réponse</summary>1) Filtrage de sécurité : le PC a-t-il les droits Lecture + Appliquer (piège MS16-072) ? 2) Héritage bloqué sur l'OU ? 3) Filtre WMI faux sur ce poste ? 4) Réplication SYSVOL/GPO pas encore arrivée sur le DC qui sert le client (`gpresult /r`, `Get-GPInheritance`).</details>

**Q6. Quelle est la différence entre transférer et saisir un rôle FSMO ?**
<details><summary>Réponse</summary>Le **transfert** (`Move-ADDirectoryServerOperationMasterRole`) se fait avec le détenteur en ligne : propre et réversible. La **saisie** (`ntdsutil` → seize) se fait quand le détenteur est mort et ne reviendra jamais : l'ancien DC ne doit plus jamais être reconnecté (risque de RID dupliqués).</details>

