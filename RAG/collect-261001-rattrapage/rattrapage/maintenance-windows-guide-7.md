---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-7
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "dpo"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [1152, 1333]
sha256: 4978a1b8c95c2ec96d01c9cfef67e9c5c418abe242afec469816877895aef117
---

# Maintenance et exploitation Windows en entreprise

```xml
<Subscription xmlns="http://schemas.microsoft.com/2006/03/windows/events/subscription">
  <SubscriptionId>Erreurs-Systeme-Serveurs</SubscriptionId>
  <Description>Erreurs du journal Systeme des serveurs</Description>
  <Uri>http://schemas.microsoft.com/wbem/wsman/1/windows/EventLog</Uri>
  <ConfigurationMode>MinLatency</ConfigurationMode>
  <Query><![CDATA[
    <QueryList><Query Id="0"><Select Path="System">*[System[Level=2]]</Select></Query></QueryList>
  ]]></Query>
  <LogFile>ForwardedEvents</LogFile>
</Subscription>
```

**Côté sources** (GPO) :

```text
Configuration ordinateur > Stratégies > Modèles d'administration >
Composants Windows > Transfert d'événements > Configurer le gestionnaire
d'abonnement cible :
  Server=http://collecteur.contoso.local:5985/wsman/SubscriptionManager/WEC,Refresh=3600
```

Points de vigilance : WinRM/HTTP ouvert (5985), compte du collecteur autorisé
(groupe « Lecteurs du journal des événements » sur les sources), taille du journal
`ForwardedEvents` dimensionnée (plusieurs Go sur un gros parc).

---

## 37. Journal Système : ce qu'on y cherche

| ID | Source | Signification | Réflexe |
|---|---|---|---|
| 41 | Kernel-Power | Arrêt brutal (coupure, plantage) | Corréler avec onduleur/BSOD |
| 1074 | User32 | Arrêt/redémarrage initié (qui, pourquoi) | Vérifier les reboots planifiés |
| 6005/6006/6008 | EventLog | Démarrage/arrêt du journal ; 6008 = arrêt inattendu | Compter les 6008/semaine |
| 55, 98, 140, 153 | Ntfs / disk | Corruption / erreurs disque | SMART + chkdsk (§25, §52) |
| 7, 11, 15 | disk | Erreur du pilote de disque | Câble/contrôleur/disque |
| 1001 | BugCheck | BSOD enregistré (avec code) | Analyse dump (§72) |
| 7023/7031/7034 | Service Control Manager | Service en échec / arrêté inopinément | Dépendances, compte de service |
| 7045 | Service Control Manager | **Nouveau service installé** | Suspect si non planifié (persistance malveillante) |
| 19/20 | WindowsUpdateClient | Échec d'installation de MAJ | Code d'erreur Windows Update |

Requête hebdo type (à mettre dans la vue personnalisée §35) : niveaux
Critique/Erreur du journal Système, groupés par ID (§33).

---

## 38. Journal Sécurité : audit et alertes

Prérequis : la **stratégie d'audit** doit être activée (GPO : `Configuration
ordinateur > Paramètres Windows > Paramètres de sécurité > Stratégies locales >
Stratégie d'audit`, ou audit avancé). Sans audit, le journal Sécurité est vide
d'intérêt.

| ID | Signification | Alerte si… |
|---|---|---|
| 4624 | Ouverture de session réussie (type 2 locale, 3 réseau, 10 RDP) | Connexion RDP inhabituelle (heure, IP) |
| 4625 | Échec d'ouverture de session | Pic soudain = attaque par force brute |
| 4634/4647 | Fermeture de session | — |
| 4648 | Ouverture de session avec identifiants explicites | Usage anormal (runas massif) |
| 4672 | Privilèges spéciaux attribués | Associé à un logon admin |
| 4720/4728/4732/4756 | Création de compte / ajout à un groupe (sensibles) | Toute occurrence non planifiée |
| 4726 | Suppression de compte | Vérifier le bien-fondé |
| 4738 | Modification de compte | — |
| 4740 | Compte verrouillé | Pic = brute force ou service avec vieux mot de passe |
| 4768/4769 | Tickets Kerberos TGT/TGS | Anomalies (chiffrements faibles) |
| 1102 | Journal d'audit effacé | **Toujours suspect** |

Surveillance minimale sur les DC : 4625 (pics), 4740 (pics), 4720/4728/4732/4756
(toute occurrence), 1102 (toute occurrence). Remontez-les via WEF (§36) ou votre
SIEM.

```powershell
# Comptes verrouillés dans les dernières 24 h (sur un DC)
Get-WinEvent -FilterHashtable @{ LogName='Security'; Id=4740; StartTime=(Get-Date).AddDays(-1) } |
  Select-Object TimeCreated, @{n='Compte'; e={ $_.Properties[0].Value }} |
  Format-Table -AutoSize
```

---

## 39. Journaux Active Directory : DS, DNS Server, DFSR

Sur un contrôleur de domaine, ajoutez ces journaux à la surveillance :

| Journal (Applications et services) | IDs clés | Sens |
|---|---|---|
| Directory Service | 1084, 1311, 1864 | Réplication AD : erreurs, latence |
| DNS Server | 4015, 4004, 4521 | Échecs DNS sur le DC |
| DFS Replication | 2213, 2104, 5002 | Réplication SYSVOL/DFSR en panne |
| Active Directory Web Services | — | Requêtes ADWS (PowerShell AD) |

Contrôle express de la réplication (quotidien, §5) :

```cmd
:: Résumé de la réplication : toute erreur > 0 est à traiter
repadmin /replsummary

:: Détail des échecs
repadmin /showrepl * /errorsonly

:: État des contrôleurs du domaine
dcdiag /q
```

> `dcdiag` sans `/q` est verbeux : en routine, `/q` (erreurs seules) suffit.
> Un `dcdiag` propre + `repadmin` propre = un AD sain à 95 %.

---

## 40. Journaux DHCP, NPS et autres rôles

| Rôle | Journal / outil | À surveiller |
|---|---|---|
| DHCP | `Microsoft > Windows > DHCP-Server` (opérationnel) + `C:\Windows\System32\dhcp\DhcpSrvLog-*.log` | Étendue saturée (ID 1020), conflits, NACK |
| NPS (RADIUS) | Journaux NPS (`C:\Windows\System32\LogFiles\IN*.log`) | Échecs d'authentification Wi-Fi/VPN |
| Serveur de fichiers | Journal Sécurité (audit d'accès si activé) + SMBServer/SMBClient | Erreurs SMB, partages inaccessibles |
| RDS | `Microsoft > Windows > TerminalServices-*` | Échecs de connexion, licences |
| Hyper-V | `Microsoft > Windows > Hyper-V-*` | Arrêts VM, réplication |

Exemple : détecter une étendue DHCP saturée :

```powershell
Get-WinEvent -LogName 'Microsoft-Windows-DHCP-Server/Operational' |
  Where-Object Id -eq 1020 | Select-Object -First 5 TimeCreated, Message
```

> Les logs texte DHCP (`DhcpSrvLog-Lun.log`) sont au format CSV daté : pensez à
> les purger/rotater comme les logs IIS (§30).

---

## 41. Rétention et archivage des journaux

Dimensionnement par défaut (souvent trop juste sur serveur) :

```powershell
# Voir la taille max et la rétention actuelles
wevtutil gl System
wevtutil gl Security

# Augmenter : journal Sécurité à 1 Go, écrasement des plus anciens si besoin
wevtutil sl Security /ms:1073741824
# Ou : archiver au lieu d'écraser (ne pas écraser, archiver)
wevtutil sl Security /rt:false /ab:true
```

**Politique type :**

- Serveurs : Système/Sécurité ≥ 512 Mo–1 Go, rétention ≥ 90 jours (ou transfert
  vers collecteur/SIEM).
- DC : Sécurité ≥ 1 Go (les audits sont volumineux).
- Archivage : export `.evtx` mensuel vers un partage d'archives
  (`wevtutil epl`), conservé 1 an (contraintes légales/assurance à vérifier
  avec votre DPO).

> Un journal Sécurité qui écrase tous les 3 jours = aucune capacité d'enquête
> a posteriori. C'est un problème d'exploitation, pas de « place disque ».

---

## 42. Planificateur de tâches : concepts

Le **Planificateur de tâches** (`taskschd.msc`) exécute les automatismes Windows.
Vocabulaire :

| Concept | Rôle |
|---|---|
| **Action** | Quoi faire (programme, script, e-mail déprécié) |
| **Déclencheur** | Quand (horaire, au démarrage, à l'ouverture de session, sur événement) |
| **Condition** | Contraintes (secteur, réseau, batterie) |
| **Principal** | Sous quel compte (SYSTEM, utilisateur, gMSA) |
| **Paramètres** | Comportement (relance en cas d'échec, arrêt si trop long) |

Toujours configurer : **« Arrêter la tâche si elle s'exécute plus de X »**
(évite les scripts bloqués qui s'accumulent) et **« Si la tâche échoue, redémarrer
toutes les X minutes »** (résilience).

---

## 43. Créer une tâche en PowerShell

