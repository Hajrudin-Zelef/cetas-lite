---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-4
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [521, 708]
sha256: af4bbff1e0ed37e34d6a66452d82a36433c6b618a014dead39e243bfccd39d1c
---

# Maintenance et exploitation Windows en entreprise

```powershell
# Décliner les mises à jour remplacées et non nécessaires (à adapter)
[void][reflection.assembly]::LoadWithPartialName("Microsoft.UpdateServices.Administration")
$wsus = [Microsoft.UpdateServices.Administration.AdminProxy]::GetUpdateServer()
$updates = $wsus.GetUpdates() | Where-Object {
    $_.IsSuperseded -and -not $_.IsDeclined
}
$updates | ForEach-Object { $_.Decline() }
"Déclinées : $($updates.Count)"
```

2. **Assistant Nettoyage du serveur** : cocher tout (mises à jour expirées,
   remplacées, fichiers inutiles, ordinateurs inactifs, révisions).
3. **Réindexer la base WID** (script `WsusDBMaintenance.sql` fourni par Microsoft,
   à exécuter mensuellement via sqlcmd sur `\\.\pipe\MICROSOFT##WID\tsql\query`).
4. **Surveiller l'espace** : le dossier `WsusContent` ne doit pas saturer le volume.

> Si la console MMC ne s'ouvre plus : augmentez la mémoire privée du pool
> d'applications `WsusPool` dans IIS (recommandé : 4 Go+) et recyclez-le. C'est le
> symptôme n°1 d'un WSUS négligé.

---

## 16. Windows Update for Business : principes

WUfB = pilotage des mises à jour **sans serveur intermédiaire** : les clients
téléchargent depuis Windows Update / Microsoft Update, mais **vous** contrôlez le
*quand* via stratégies (GPO, Intune, MDM). C'est la voie recommandée par Microsoft.

| Élément | WSUS | WUfB |
|---|---|---|
| Source des binaires | Serveur local | Microsoft Update (CDN) |
| Bande passante WAN | Faible (1 téléchargement/site) | Plus élevée (optimisable : Delivery Optimization, §ci-dessous) |
| Infrastructure | Serveur + base + maintenance | Aucune (stratégies uniquement) |
| Granularité | Approbation par KB | Reports en jours, pauses, anneaux |
| Avenir | Déprécié | Stratégique |

**Delivery Optimization** (optimisation de la distribution) : indispensable avec WUfB
pour ne pas saturer les liens WAN — les postes s'échangent les morceaux en P2P sur
le LAN. À activer par GPO/Intune (`Téléchargement > Optimisation de la distribution`,
mode « LAN » ou « Groupe »).

---

## 17. WUfB : stratégies de report (quality/feature)

Deux familles de reports (en jours) pilotent le déploiement :

| Stratégie | Contenu | Report conseillé (Prod) |
|---|---|---|
| Mises à jour qualité (Quality) | Correctifs mensuels | 7-14 jours |
| Mises à jour de fonctionnalités (Feature) | Nouvelles versions Windows 11 | 60-120 jours |

GPO correspondantes :

```text
Configuration ordinateur > Stratégies > Modèles d'administration >
Composants Windows > Windows Update > Gérer les mises à jour proposées par
Windows Update Entreprise
  - Sélectionner le moment de réception des builds Preview et des mises à
    jour de fonctionnalités : canal « Général », report 90 jours
  - Sélectionner le moment de réception des mises à jour qualité : report 10 jours
```

**Pause** : possibilité de suspendre 35 jours max (qualité) en cas de correctif
défectueux — le bouton d'urgence du §11.

**Anneaux via Intune** : créez des profils « Update rings for Windows 10 and later »
par anneau (Test : report 0 j / Pilote : 5 j / Prod : 10 j), assignés à des groupes
Entra ID. C'est l'équivalent moderne des groupes WSUS.

---

## 18. Intune vs GPO : pilotage des mises à jour

| Critère | GPO (AD local) | Intune (cloud) |
|---|---|---|
| Prérequis | Domaine AD, postes joints | Licences Intune, postes joints Entra ID / hybrides |
| Postes nomades | Non gérés hors VPN | Gérés partout |
| Rapports | Limités (scripts) | Update Compliance / rapports natifs |
| Coexistence | — | La co-gestion est possible (SCCM + Intune) |

**Règle de conflit** : si un poste reçoit à la fois une GPO WSUS (clé `WUServer`) et
une stratégie WUfB/Intune, **le WSUS gagne** (la présence de `WUServer` désactive
WUfB). Pour migrer WSUS → WUfB : supprimez proprement les stratégies WSUS
(`WUServer`, `WUStatusServer`, `UseWUServer=0`) avant d'appliquer les anneaux.

Vérification du mode actif côté client :

```powershell
# 0 = Microsoft Update direct (WUfB), 1 = WSUS configuré
$wu = 'HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU'
(Get-ItemProperty $wu -ErrorAction SilentlyContinue).UseWUServer
# Source réelle des stratégies (MDM vs GPO) : voir aussi rsop / MDM Diagnostics
```

---

## 19. Fenêtres de maintenance et heures d'activité

**Serveurs** : pas de redémarrage automatique sauvage. Pilotez par GPO :

```text
Composants Windows > Windows Update > Gérer l'expérience utilisateur final
  - Ne pas redémarrer automatiquement avec des utilisateurs connectés : Activé
  - Toujours redémarrer automatiquement à l'heure planifiée : selon fenêtre
```

Ou imposez l'installation + redémarrage via script/ordonnanceur pendant la fenêtre
contractuelle (§4), avec notification aux utilisateurs (RDS : `msg *`).

**Postes Windows 11** : les **heures d'activité** (ex. 08:00–19:00) empêchent le
redémarrage pendant le travail ; au-delà, le poste redémarre seul. Configurez-les
par GPO/Intune et laissez un **délai de grâce** (ex. 3-7 jours) avant redémarrage
forcé pour les correctifs qualité.

```powershell
# Consulter / définir les heures d'activité (Windows 11)
Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\WindowsUpdate\UX\Settings' |
  Select-Object ActiveHoursStart, ActiveHoursEnd
# Exemple de notification avant maintenance sur un serveur RDS
msg * /SERVER:srv-rds-01 "Maintenance planifiée à 22h : sauvegarde en cours, merci de fermer vos sessions."
```

---

## 20. Rollback : désinstaller un correctif

Quand un correctif pose problème et qu'aucun contournement n'existe :

```powershell
# 1. Identifier le KB fautif
Get-HotFix | Sort-Object InstalledOn -Descending | Select-Object -First 10 HotFixID, InstalledOn, Description

# 2a. Désinstaller (méthode classique, invite admin)
wusa /uninstall /kb:5034765 /quiet /norestart

# 2b. Via DISM (utile si wusa échoue, notamment sur serveurs)
DISM /Online /Remove-Package /PackageName:Package_for_RollupFix~31bf3856ad364e35~amd64~~26100.1234.1.1

# 3. Bloquer sa réinstallation
#  - WSUS : refuser la mise à jour pour les groupes concernés
#  - WUfB : mettre en pause les mises à jour qualité (35 j max) ou outil wushowhide.diagcab (postes)
```

**Nom du package** pour DISM : `DISM /Online /Get-Packages | findstr RollupFix`.

**Limites** :

- Un cumulatif ne se « répare » pas partiellement : on le retire en entier.
- Après 10 jours (postes) / selon configuration (serveurs), les fichiers de
  désinstallation peuvent être purgés → rollback impossible sans restauration.
- Sur un **contrôleur de domaine**, évitez le rollback sauf nécessité absolue :
  préférez la restauration d'état système (§126) en cas de corruption AD.

> Procédure d'équipe : tout rollback est tracé (ticket : KB, machines, motif,
> validation) et suivi d'une revue au patch suivant (le correctif réapparaît).

---

## 21. DISM : comprendre le magasin de composants

Le **magasin de composants** (`C:\Windows\WinSxS`) est la base de données des
fichiers système Windows. Quand il est corrompu, les mises à jour échouent
(erreur `0x80073712`, `0x800f081f`) et `sfc` ne suffit plus. **DISM**
(Deployment Image Servicing and Management) est l'outil de réparation.

Ordre de bataille en cas de système « bizarre » après un patch raté :

1. `DISM /Online /Cleanup-Image /RestoreHealth` (§22)
2. `sfc /scannow` (§24)
3. Redémarrage, re-test de Windows Update

---

## 22. DISM : /ScanHealth, /CheckHealth, /RestoreHealth

```cmd
:: 1. Vérification rapide (lit un marqueur, quelques secondes)
DISM /Online /Cleanup-Image /CheckHealth

:: 2. Analyse approfondie (plusieurs minutes)
DISM /Online /Cleanup-Image /ScanHealth

:: 3. Réparation : télécharge les fichiers sains depuis Windows Update
DISM /Online /Cleanup-Image /RestoreHealth
```

