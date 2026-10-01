---
id: collect-261001-general-networking/general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026-4
title: "Calculer l'empreinte SHA-256 de l'installeur téléchargé"
domain: general-networking
role: reference
task: reference
actors: ["Samsung"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/crystaldiskinfo-tutoriel-sante-ssd-en-12-etapes-2026.md
source_anchor: ""
source_lines: [161, 258]
sha256: b0639ca103e368a3b99ed263d74c76adabd5d96a04b00d41371498e8d16dc19a
---

# Calculer l'empreinte SHA-256 de l'installeur téléchargé

```
[Setting]
Language=French
AutoRefresh=10          ; actualisation toutes les 10 minutes
Startup=1               ; lancement au demarrage de Windows
ResidentMinimize=1      ; demarrer reduit dans la zone de notification
AlarmTemperature=55     ; seuil d'alarme temperature (deg C)
AlarmHealthStatus=1     ; alarme si l'etat passe sous "Bon"
MailAlert=1             ; activer l'envoi d'e-mail sur alerte
TrayTemperature=1       ; afficher la temperature dans la barre des taches
```
Les noms de clés exacts peuvent varier légèrement selon la version ; le plus fiable reste de configurer une machine « modèle » via l’interface graphique, puis de récupérer son `DiskInfo.ini` comme gabarit à répliquer.

## Étape 10 – Automatiser l’export des rapports S.M.A.R.T.

Pour l’administration d’un parc, l’idéal est de collecter périodiquement l’état de tous les disques sans intervention. CrystalDiskInfo expose pour cela des **options en ligne de commande** officielles, particulièrement pratiques pour le scripting :

- `/Copy` – met à jour les données puis écrit le résultat de « Édition → Copier » dans un fichier`DiskInfo.txt` .
- `/CopyExit` – identique, mais ferme l’application ensuite (parfait pour une tâche planifiée).
- `/Exit` – met à jour les valeurs S.M.A.R.T. et l’état AAM/APM, puis quitte.

Créons une tâche planifiée Windows qui exporte l’état de tous les disques chaque nuit. Le script batch suivant lance CrystalDiskInfo en mode `/CopyExit`, puis archive le rapport horodaté :

```
@echo off
REM export-smart.bat -- export nocturne du rapport S.M.A.R.T.
set "CDI=C:\Outils\CrystalDiskInfo\DiskInfo64.exe"
set "LOGDIR=C:\Logs\SMART"
if not exist "%LOGDIR%" mkdir "%LOGDIR%"
REM Genere DiskInfo.txt dans le dossier du programme, puis quitte
"%CDI%" /CopyExit
REM Horodate et archive le rapport
set "STAMP=%date:~-4%%date:~3,2%%date:~0,2%"
copy /Y "C:\Outils\CrystalDiskInfo\DiskInfo.txt" "%LOGDIR%\smart-%COMPUTERNAME%-%STAMP%.txt"
```
Planifiez ensuite ce script via le Planificateur de tâches Windows, en une seule ligne PowerShell exécutée en administrateur :

```
# Cree une tache quotidienne a 03h00 qui exporte l'etat S.M.A.R.T.
$action  = New-ScheduledTaskAction -Execute "C:\Outils\CrystalDiskInfo\export-smart.bat"
$trigger = New-ScheduledTaskTrigger -Daily -At 3am
$princ   = New-ScheduledTaskPrincipal -UserId "SYSTEM" -RunLevel Highest
Register-ScheduledTask -TaskName "Export-SMART-CrystalDiskInfo" `
  -Action $action -Trigger $trigger -Principal $princ `
  -Description "Export nocturne des rapports S.M.A.R.T. via CrystalDiskInfo"
```
Vous disposez désormais d’un historique daté, poste par poste, exploitable pour repérer les tendances (un secteur réalloué qui apparaît, une température qui grimpe semaine après semaine). Cette approche par ligne de commande est la même que celle documentée par la communauté DevOps, notamment sur CyberDrain pour l’intégrer à un outil de supervision RMM.

## Étape 11 – Surveiller spécifiquement les SSD NVMe (TBW et endurance)

Les SSD NVMe imposent une lecture différente. Ici, pas de secteurs mécaniques : le risque n’est pas la panne brutale mais l’**épuisement de l’endurance d’écriture**. Trois indicateurs concentrent l’attention. Le *Percentage Used* exprime, en pourcentage, l’usure estimée par le contrôleur : à 100 %, le disque a consommé l’endurance garantie (il continue souvent de fonctionner, mais hors garantie). L’*Available Spare* indique la réserve de blocs de secours ; s’il chute sous son seuil (« Spare Threshold »), le SSD est en fin de vie. Enfin, le *Total Host Writes* (ou *Data Units Written*) totalise les octets écrits, à comparer au TBW inscrit sur la fiche technique.

Un exemple concret : un SSD NVMe de 2 To garanti pour 1 200 TBW. Si CrystalDiskInfo affiche 150 To écrits après deux ans, vous avez consommé 12,5 % de l’endurance – le disque tiendra largement au-delà de sa garantie de cinq ans à ce rythme. Voici un extrait de rapport `DiskInfo.txt` réel pour un tel SSD, tel qu’exporté à l’étape 10 :

```
----------------------------------------------------------------------------
 (1) Samsung SSD 990 PRO 2TB
----------------------------------------------------------------------------
           Modele : Samsung SSD 990 PRO 2TB
    Micrologiciel : 4B2QJXD7
        Interface : NVM Express
   Mode transfert : PCIe 4.0 x4 | PCIe 4.0 x4
   Lettre lecteur : C:
    Etat de sante : Bon (94 %)
      Temperature : 43 C (109 F)
   Heures fonct.  : 6821 heures
    Nb demarrages : 742
  Total ecritures : 150331 GB
   Total lectures : 214887 GB
-- S.M.A.R.T. --------------------------------------------------------------
ID Actuel Pire Seuil  Valeurs brutes   Nom de l'attribut
01  100   100   __0   000000000000     Critical Warning
05  100   100   __0   000000000006     Percentage Used  (6 %)
07  100   100   __0   000000000064     Available Spare  (100 %)
0C  100   100   __0   00000002C7A1     Unsafe Shutdowns
```
Pour recouper ces valeurs, Windows propose un complément natif via PowerShell. Les compteurs de fiabilité de stockage exposent une partie des mêmes données, utile pour un contrôle croisé ou une supervision sans installer de logiciel :

```
# Lire l'usure et la temperature via PowerShell (complement a CrystalDiskInfo)
Get-PhysicalDisk | Get-StorageReliabilityCounter |
  Select-Object DeviceId, Wear, Temperature, ReadErrorsTotal, PowerOnHours |
  Format-Table -AutoSize
# Afficher l'etat de sante global vu par Windows
Get-PhysicalDisk | Select-Object FriendlyName, MediaType, HealthStatus, OperationalStatus
```
La colonne `Wear` renvoyée par Windows correspond au *Percentage Used* de CrystalDiskInfo. Si les deux outils divergent nettement, faites confiance à CrystalDiskInfo, qui lit directement le journal S.M.A.R.T. du contrôleur plutôt qu’une abstraction du système.

## Étape 12 – Gérer l’AAM/APM et exporter un diagnostic complet

Dernière étape, souvent méconnue : CrystalDiskInfo peut aussi *agir* sur certains disques via `Fonction → Paramètres avancés → Contrôle AAM/APM`. L’**AAM** (Automatic Acoustic Management) ajuste le compromis bruit/performance de la tête de lecture d’un disque dur ; l’**APM** (Advanced Power Management) règle son agressivité de mise en veille. Ces réglages ne concernent que les disques durs mécaniques compatibles – la plupart des SSD les ignorent. Sur un vieux disque bruyant, réduire l’AAM peut sensiblement diminuer le cliquetis, au prix d’un léger recul de performance.

Pour clore votre diagnostic, générez un rapport complet à archiver ou à joindre à une demande de garantie (RMA). Le menu `Édition → Copier` place l’intégralité du rapport dans le presse-papiers ; collez-le dans un fichier texte. Ce rapport contient le modèle, le numéro de série, le micrologiciel et l’ensemble des attributs S.M.A.R.T. – exactement ce qu’un service après-vente demandera pour valider un remplacement sous garantie. Conservez-le : en cas de litige avec un vendeur, un rapport S.M.A.R.T. daté fait office de preuve technique objective.

## Projet complet : un pipeline de surveillance automatisée

Assemblons maintenant les briques précédentes en un système de surveillance autonome. L’objectif : un poste qui exporte son état S.M.A.R.T. chaque nuit (étape 10), qu’un script analyse pour détecter tout état « Attention » ou « Mauvais » et vous alerte immédiatement. Le script PowerShell suivant lit le dernier rapport `DiskInfo.txt`, en extrait chaque disque et son état de santé, et envoie un courriel si un problème est détecté :

