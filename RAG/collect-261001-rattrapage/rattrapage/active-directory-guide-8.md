---
id: collect-261001-rattrapage/rattrapage/active-directory-guide-8
title: "Active Directory & GPO en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["advisory", "parameters"]
source: docs/RAG/collect-261001-rattrapage/active_directory_guide.md
source_anchor: ""
source_lines: [1164, 1343]
sha256: a583f5af3447943fd2ffc87474ea327edf9416f6ae03dbf80c2c1feeaa9dbd57
---

# Planification : répliquer uniquement la nuit (ex. 20h-6h)
# Via la console : propriétés de la liaison > Modifier la planification
# En PowerShell, le schedule est un tableau d'octets ; préférez la console pour ce réglage fin

# Désigner manuellement un bridgehead (déconseillé sauf besoin précis)
# Sites et services > serveur > Propriétés > cocher le transport IP comme "serveur tête de pont préféré"
```

**Coûts** : reflétez la **qualité des liaisons** (bande passante, fiabilité). Le KCC calcule les chemins les moins coûteux, y compris via des sites intermédiaires (pontage de liaisons de sites, activé par défaut : « Bridge all site links »).

**Fréquence** : 15 min minimum via la console (valeur basse = 15). En dessous, passez par la notification inter-sites (possible mais rarement nécessaire).

---

## 40. Réplication AD : KCC et topologie

Le **KCC** (Knowledge Consistency Checker) tourne sur chaque DC (toutes les 15 min) et construit automatiquement la topologie de réplication :

- **Intra-site** : anneau bidirectionnel + connexions supplémentaires (chaque DC a au moins 2 partenaires).
- **Inter-sites** : via les bridgeheads, selon les liaisons de sites.

**Types de réplication** :

| Partition | Mode |
|---|---|
| Domaine, Configuration, Schéma | Multi-maître (tout DC inscriptible accepte les écritures) |
| SYSVOL (stratégies) | DFSR (depuis 2008 ; FRS est mort, ne l'utilisez plus) |

```powershell
# Forcer le KCC à recalculer la topologie
repadmin /kcc

# Voir les partenaires de réplication entrante
repadmin /showrepl DC01

# Topologie inter-sites
repadmin /bridgeheads
```

**À retenir** : dans 99 % des cas, **ne touchez pas** aux objets de connexion créés par le KCC (`<automatically generated>`). Créez des connexions manuelles uniquement pour des besoins temporaires (dépannage) et supprimez-les après.

---

## 41. Surveiller la réplication : repadmin

`repadmin` est l'outil n°1 du dépannage réplication.

```powershell
# Résumé de l'état de réplication de tous les DC (le premier à lancer)
repadmin /replsummary

# Détail par DC : derniers succès/échecs, USN
repadmin /showrepl
repadmin /showrepl DC02

# File d'attente de réplication
repadmin /queue

# Forcer la réplication d'une partition vers tous les partenaires
repadmin /syncall DC01 /AdeP
# /A : toutes les partitions, /d : noms distinctifs, /e : tous les sites, /P : push

# Vérifier l'état du SYSVOL/DFSR
dfsrdiag pollad
Get-DfsrState  # (module DFSR)

# Connaître l'USN le plus élevé et les "high watermarks"
repadmin /showutdvec DC01 "DC=ad,DC=entreprise,DC=fr"
```

**Lecture de `/replsummary`** : la colonne des échecs doit être à **0** partout. Un échec isolé et récent peut être transitoire ; un échec qui grandit = investigation immédiate (réseau, DNS, heure, espace disque).

---

## 42. Diagnostiquer avec dcdiag

```powershell
# Diagnostic complet du DC local
dcdiag

# Diagnostic verbeux, tous les tests
dcdiag /v

# Tester tous les DC du site / de l'entreprise
dcdiag /s:DC02
dcdiag /a          # tous les DC du site
dcdiag /e          # tous les DC de l'entreprise

# Tests ciblés utiles
dcdiag /test:dns /v          # santé DNS (le plus critique)
dcdiag /test:replications
dcdiag /test:netlogons
dcdiag /test:advertising     # le DC s'annonce-t-il correctement ?
dcdiag /test:fsmocheck
dcdiag /test:ridmanager
```

**Tests à surveiller en priorité** : `DNS`, `Replications`, `NetLogons`, `Advertising`, `FSMOCheck`, `RidManager`, `Services`, `SystemLog`. Un `dcdiag` propre sur tous les DC = 80 % de la santé AD.

> 💡 Lancez `dcdiag /e /v` **avant** toute opération sensible (promotion, transfert FSMO, décommission) : c'est votre photo « avant travaux ».

---

## 43. Dépannage réplication : USN rollback

**USN rollback** : le cauchemar des DC virtualisés. Chaque DC maintient un **USN** (Update Sequence Number) qui s'incrémente à chaque écriture. Les partenaires retiennent le dernier USN vu (`high watermark`).

**Scénario** : on restaure un snapshot d'un DC vieux de 3 jours. Son USN **recule**. Ses partenaires croient déjà avoir tout jusqu'à l'USN 1000, le DC restauré est à 800 : les écritures 801-1000 (mots de passe changés, comptes créés) **ne seront jamais répliquées** → divergence silencieuse, comptes verrouillés mystérieusement, etc.

**Symptômes** : événement **2108/1084** (NTDS Replication), `repadmin /showrepl` avec des USN incohérents.

**Prévention** :

1. Ne restaurez **jamais** un DC par snapshot hyperviseur. Utilisez la sauvegarde d'état système.
2. Activez la protection **VM-GenerationID** (hyperviseur récent + OS 2012+) : en cas de restore, l'ID change, le DC s'en rend compte et se met en sécurité (il demande une réplication complète / se désactive proprement avec l'événement 2162).
3. Un seul DC physique ou un DC hors snapshots critiques dans les petits parcs.

**Remède** : le DC en USN rollback doit être **rétrogradé de force, nettoyé (metadata cleanup) et repromu**. Il n'y a pas de « réparation » fiable.

---

## 44. Dépannage réplication : lingering objects (objets rémanents)

**Lingering object** : un objet supprimé sur les autres DC mais qui **réapparaît** sur un DC resté trop longtemps déconnecté (au-delà de la **tombstone lifetime**, 180 jours par défaut). Quand il se reconnecte, il réinjecte l'objet fantôme.

**Symptômes** : événement **1388** (NTDS Replication) : « le DC a reçu une mise à jour pour un objet supprimé » ; `repadmin /showrepl` en erreur.

**Traitement** :

```powershell
# 1. Identifier les objets rémanents (comparer deux DC)
repadmin /removelingeringobjects DC01 <GUID_DC_Reference> "DC=ad,DC=entreprise,DC=fr" /advisory_mode
# /advisory_mode : journalise sans supprimer (toujours commencer par là)

# 2. Suppression réelle (après analyse du journal d'événements 1937)
repadmin /removelingeringobjects DC01 <GUID_DC_Reference> "DC=ad,DC=entreprise,DC=fr"

# 3. Durcir : activer la protection stricte contre la réplication des objets rémanents
# Registre sur chaque DC :
Set-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\NTDS\Parameters" `
  -Name "Strict Replication Consistency" -Value 1 -Type DWord
```

**Prévention** : ne laissez jamais un DC déconnecté plus de 180 jours ; surveillez `repadmin /replsummary` (un DC silencieux depuis des semaines = alerte).

---

## 45. Stratégie de mot de passe par défaut du domaine

La **stratégie de mot de passe du domaine** se définit dans la **Default Domain Policy** (liée à la racine du domaine) : elle s'applique à **tous** les utilisateurs du domaine.

Chemin GPO : `Configuration ordinateur > Stratégies > Paramètres Windows > Paramètres de sécurité > Stratégies de comptes > Stratégie de mot de passe`.

| Paramètre | Défaut | Recommandation entreprise |
|---|---|---|
| Longueur minimale | 7 | **12-14** (16 pour les admins) |
| Complexité | Activée | Activée (ou passphrases longues) |
| Durée de vie maximale | 42 jours | 90-180 j utilisateurs ; **ne pas expirer** les comptes de service gMSA (géré par AD) ; admins : 60-90 j + MFA |
| Durée de vie minimale | 1 jour | 1 jour (anti-contournement du changement immédiat) |
| Historique | 24 | 24 |

**Verrouillage de compte** (`Stratégies de comptes > Stratégie de verrouillage`) :

| Paramètre | Recommandation |
|---|---|
| Seuil de verrouillage | 5 tentatives |
| Durée du verrouillage | 30 min (ou jusqu'au déverrouillage admin) |
| Compteur réinitialisé après | 30 min |

```powershell
# Lire la stratégie effective du domaine
Get-ADDefaultDomainPasswordPolicy | Format-List *

# Déverrouiller un compte
Unlock-ADAccount -Identity "a.diallo"

# Comptes verrouillés en ce moment
Search-ADAccount -LockedOut | Select-Object Name, SamAccountName
```

