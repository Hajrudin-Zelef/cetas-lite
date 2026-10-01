---
id: collect-261001-rattrapage/rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026-2
title: "Calculer l'empreinte SHA-256 de l'archive (PowerShell)"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026.md
source_anchor: ""
source_lines: [51, 141]
sha256: 6304aa529b7dcec5c815cdfa8e8f506652089b26dcf2670f547b27111b5f5f75
---

# Calculer l'empreinte SHA-256 de l'archive (PowerShell)

Deux précautions complètent la liste. D’abord, **installez et exécutez DDU depuis un disque local** (C: ou D:) : l’outil refuse de fonctionner depuis un lecteur réseau. Ensuite, prévoyez de quoi vous connecter sans le code PIN : sur certaines configurations Windows 11 24H2, l’authentification Windows Hello est indisponible en mode sans échec (voir la section dédiée plus bas). Ayez votre mot de passe de compte Microsoft sous la main.

## Étape 1 – Télécharger DDU depuis une source officielle

Ne téléchargez jamais **Display Driver Uninstaller** depuis un site de partage inconnu : les versions repackagées sont un vecteur classique d’adware. Deux sources font autorité : le site de l’éditeur Wagnardsoft et le miroir historique Guru3D. Les deux distribuent désormais une archive auto-extractible nommée `DDU_v18.1.5.6.exe`, mise en ligne le 17 juillet 2026, pour un poids d’environ 1,6 Mo une fois l’installateur récupéré.

Une fois le fichier récupéré, vérifiez son intégrité avant de l’ouvrir. PowerShell calcule l’empreinte SHA-256 en une commande, que vous comparerez à celle publiée sur la page de téléchargement :

```
# Calculer l'empreinte SHA-256 de l'archive (PowerShell)
Get-FileHash "$env:USERPROFILE\Downloads\DDU_v18.1.5.4.exe" -Algorithm SHA256
# Sortie attendue (extrait) :
# Algorithm       Hash
# ---------       ----
# SHA256          9F3C...     C:\Users\vous\Downloads\DDU_v18.1.5.4.exe
```
Si l’empreinte correspond, vous tenez bien la version officielle. Conservez l’archive dans un dossier facile d’accès : nous la décompresserons à l’étape suivante. Inutile de la lancer maintenant – nous voulons d’abord cartographier votre configuration graphique.

## Étape 2 – Extraire l’archive dans un dossier dédié

L’exécutable de DDU est une archive 7-Zip auto-extractible. En le lançant, il propose de décompresser son contenu dans un sous-dossier. Choisissez un emplacement court et sans espace, par exemple `C:\DDU` : cela simplifiera les commandes en ligne de commande de l’étape 12. Vous pouvez aussi automatiser l’extraction en silence :

```
:: Extraire DDU dans C:\DDU sans interaction (invite de commandes)
"%USERPROFILE%\Downloads\DDU_v18.1.5.4.exe" -o"C:\DDU" -y
:: Vérifier que l'exécutable principal est bien présent
dir "C:\DDU\Display Driver Uninstaller.exe"
```
Le dossier obtenu contient l’exécutable `Display Driver Uninstaller.exe`, un fichier `ReadMe` et quelques ressources. Ouvrez le ReadMe : il liste les nouveautés de la version et confirme les commutateurs en ligne de commande disponibles. Ne lancez pas encore DDU en mode normal : pour un nettoyage optimal, nous passerons par le mode sans échec.

## Étape 3 – Identifier votre GPU et vos pilotes actuels

Avant tout nettoyage, notez précisément quelle carte vous possédez et quelle version de pilote elle utilise. Cette information sert de point de comparaison après réinstallation et évite de télécharger le mauvais paquet. PowerShell interroge directement le contrôleur vidéo :

```
# Lister les cartes graphiques, leur pilote et leur memoire (PowerShell)
Get-CimInstance Win32_VideoController |
  Select-Object Name, DriverVersion, @{N='VRAM_Go';E={[math]::Round($_.AdapterRAM/1GB,1)}} |
  Format-Table -AutoSize
# Detail des pilotes d'affichage et de leur etat
Get-PnpDevice -Class Display | Select-Object FriendlyName, Status, InstanceId
```
La sortie ressemble à ceci pour une carte NVIDIA :

```
Name                          DriverVersion   VRAM_Go
----                          -------------   -------
NVIDIA GeForce RTX 4070        32.0.15.7652      12,0
```
Vous pouvez aussi taper `dxdiag` dans le menu Démarrer pour ouvrir l’outil de diagnostic DirectX, dont l’onglet « Affichage » récapitule le modèle, le fabricant et la version du pilote. Notez ces valeurs. Si vous hésitez encore sur le choix de votre prochaine carte, notre comparatif RX 9070 XT vs RTX 5070 détaille le rapport performances/prix actuel.

## Étape 4 – Télécharger à l’avance les nouveaux pilotes

C’est l’étape que la plupart des tutoriels négligent, et c’est pourtant la plus importante. Une fois **DDU** passé et le redémarrage effectué, Windows bascule sur un pilote d’affichage générique : résolution limitée, pas de multi-écran fluide, pas d’accélération matérielle. Récupérer un pilote dans ces conditions est pénible. Téléchargez donc dès maintenant le paquet adapté à votre carte :

- **NVIDIA** : pilote GeForce Game Ready ou Studio depuis la page officielle NVIDIA.
- **AMD** : suite Adrenalin Edition depuis le centre de téléchargement AMD.
- **Intel** : pilote Arc & Graphics pour les GPU intégrés et les cartes Intel Arc.

Placez le fichier d’installation sur le Bureau ou dans un dossier facile à retrouver, idéalement sur le même disque que DDU. Si vous changez de marque de carte, téléchargez le pilote de la *nouvelle* carte, pas de l’ancienne. Gardez aussi sous la main votre pilote de chipset (carte mère) : il arrive qu’un nettoyage agressif touche des composants audio HDMI partagés.

## Étape 5 – Créer un point de restauration système

**DDU modifie le registre en profondeur.** Même si l’outil est mûr et fiable, un point de restauration vous offre un filet de sécurité en cas de mauvaise manipulation (mauvais fabricant sélectionné, par exemple). Sur les machines récentes, la protection système est parfois désactivée par défaut : activez-la avant de créer le point.

```
# A executer dans une console PowerShell ouverte EN ADMINISTRATEUR
# 1. Activer la protection systeme sur le disque C: (si necessaire)
Enable-ComputerRestore -Drive "C:\"
# 2. Creer le point de restauration avant d'utiliser DDU
Checkpoint-Computer -Description "Avant nettoyage DDU" -RestorePointType "MODIFY_SETTINGS"
# 3. Verifier que le point a bien ete cree
Get-ComputerRestorePoint | Select-Object SequenceNumber, Description, CreationTime
```
Si la commande `Checkpoint-Computer` renvoie une erreur de fréquence (Windows limite la création à un point toutes les 24 heures), créez le point manuellement via « Créer un point de restauration » dans les propriétés système, ou réglez temporairement la valeur de registre `SystemRestorePointCreationFrequency` sur 0. Notez le numéro de séquence affiché : il vous servira si vous devez revenir en arrière.

## Étape 6 – Empêcher Windows Update de réinstaller les pilotes

Voici le piège qui ruine la moitié des nettoyages : à peine DDU terminé, Windows Update détecte une carte « sans pilote » et réinstalle automatiquement une version générique, souvent ancienne. Résultat : vos résidus reviennent et votre nettoyage est annulé. Il faut donc **bloquer le téléchargement automatique des pilotes** avant de redémarrer. Deux clés de registre suffisent :

```
:: Invite de commandes ouverte EN ADMINISTRATEUR
:: 1. Ne plus rechercher les pilotes via Windows Update
reg add "HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\DriverSearching" /v SearchOrderConfig /t REG_DWORD /d 0 /f
:: 2. Exclure les pilotes des mises a jour qualite de Windows Update
reg add "HKLM\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate" /v ExcludeWUDriversInQualityUpdate /t REG_DWORD /d 1 /f
```
DDU propose lui-même une option « Empêcher les téléchargements de pilotes via Windows Update » dans son menu Options : cochez-la également, ceinture et bretelles. La méthode la plus radicale reste de **débrancher le câble réseau (ou couper le Wi-Fi)** pendant toute la durée de l’opération, jusqu’à la réinstallation manuelle de votre pilote. Une fois vos pilotes réinstallés, vous pourrez réactiver la recherche en repassant les valeurs de registre à 1, ou en décochant l’option dans DDU.

## Étape 7 – Démarrer Windows en mode sans échec

