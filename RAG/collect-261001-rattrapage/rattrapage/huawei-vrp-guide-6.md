---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-6
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [1081, 1308]
sha256: 7ce00510f797794cec30e52e4e9eb63406c01bd8ef7ed0aade1dceda43d62187
---

# VRP — Le système d'exploitation transversal Huawei

> L'autosave écrit dans le fichier de démarrage configuré. En production,
> préférez une sauvegarde **externalisée** (TFTP/FTP/SFTP, voir section 93)
> plutôt que de compter sur l'autosave local seul.

## 64. Choisir le fichier de démarrage

```vrp
<AR720>startup saved-configuration backup_manuelle.zip
Info: Succeeded in setting the configuration for booting system.
<AR720>display startup
MainBoard:
  Startup system software:        flash:/AR720-V300R022C00SPC500.cc
  Next startup system software:   flash:/AR720-V300R022C00SPC500.cc
  Startup saved-configuration file:       flash:/vrpcfg.zip
  Next startup saved-configuration file:  flash:/backup_manuelle.zip
  ...
```

`display startup` est LE contrôle avant tout reboot/upgrade : il montre ce
qui sera chargé au prochain démarrage (logiciel + config + patch).

## 65. `reset saved-configuration` — repartir de zéro

```vrp
<AR720>reset saved-configuration
  The action will delete the saved configuration in the device.
  Continue? [y/n]:y
<AR720>reboot
  ...
  System will reboot? [y/n]:y      # répondre N à "save" si demandé
```

Séquence de **reset usine** : `reset saved-configuration` puis `reboot`
(en répondant `n` à la sauvegarde proposée). L'équipement redémarre avec la
configuration par défaut.

## 66. `configuration rollback` (VRP8) et restauration manuelle

Sur VRP8 (base candidate) :

```vrp
[~AR720]rollback configuration                 # annule les modifs non commitées
[AR720]display configuration commit list       # historique des commits
[AR720]display configuration commit changes 3  # détail du commit n°3 (à vérifier)
```

Sur VRP5 (sans candidate) : pas de rollback natif — la restauration passe
par le fichier de sauvegarde :

```vrp
<AR720>copy tftp: 192.168.100.50 backup_avant_modif.zip flash:/vrpcfg.zip
<AR720>startup saved-configuration vrpcfg.zip
<AR720>reboot
```

D'où l'importance de la discipline : **sauvegarde externe avant chaque
fenêtre de changement** (voir section 93).

## 67. Espace flash : surveiller et nettoyer

```vrp
<AR720>dir
  42,123,456 bytes total (1,234,567 bytes free)   # <-- presque plein !
<AR720>reset recycle-bin
<AR720>delete /unreserved vieux_fichier.cc
<AR720>format flash:                            # FORMATAGE - EFFACE TOUT, dernier recours
```

Un flash plein bloque le `save` et l'upgrade. En routine : vider la
corbeille, supprimer les vieux `.cc` après validation du nouvel OS, garder
2 générations de config max en local (le reste sur serveur).

## 68. Journaux persistants

```vrp
<AR720>dir flash:/logfile/
<AR720>more flash:/logfile/log.log              # (nom exact à vérifier)
[AR720]info-center logfile enable              # journalisation persistante (à vérifier)
```

Les logs en RAM (`display logbuffer`) disparaissent au reboot ; activez la
journalisation vers fichier et/ou un serveur syslog pour l'historique.

## 69. Syslog distant

```vrp
[AR720]info-center loghost 192.168.100.60
[AR720]info-center loghost 192.168.100.60 facility local7
[AR720]info-center source default channel 2 log state on trap state off debug state off
```

Envoyez les logs vers votre SIEM/supervision (Zabbix, Loki — voir les
guides dédiés de Zelef). C'est non négociable en production.

## 70. Checklist « fin de session de configuration »

```text
[ ] 1. display this                    -> relire chaque vue modifiée
[ ] 2. compare configuration           -> diff RAM vs fichier
[ ] 3. save                            -> figer en flash
[ ] 4. display saved-configuration     -> vérifier le contenu sauvé
[ ] 5. display startup                 -> vérifier le fichier de boot
[ ] 6. Copie externe (TFTP/SFTP)       -> si changement important
```

---


## 71. Upgrade firmware — vue d'ensemble

Un upgrade VRP suit toujours le même cycle :

```text
1. Vérifier la version actuelle et l'espace libre
2. Sauvegarder la configuration (locale + externe)
3. Transférer le nouveau fichier .cc (TFTP/FTP/SFTP/USB)
4. Désigner le fichier pour le prochain boot (startup system-software)
5. Vérifier avec display startup
6. Rebooter en fenêtre de maintenance
7. Vérifier la version et le fonctionnement après reboot
```

Temps typique : 10 à 30 minutes d'indisponibilité par équipement (reboot
inclus). Sur USG6000 en cluster HRP, on upgrade en bascule pour limiter la
coupure (voir section 79).

## 72. Étape 1 : vérifications préalables

```vrp
<AR720>display version                          # version actuelle
<AR720>display startup                          # fichiers de boot actuels
<AR720>display patch-information                # patchs installés (.pat)
<AR720>dir                                      # espace libre : il faut > taille du .cc + marge
<AR720>display device                           # état matériel OK ?
<AR720>display alarm active                     # aucune alarme bloquante ?
```

Checklist pré-upgrade (à cocher) :

```text
[ ] Version cible compatible avec le modèle (notes de version Huawei)
[ ] Chemin d'upgrade supporté (ex. V300R019 -> V300R022 direct ? ou palier ?)
[ ] Fichier .cc téléchargé depuis le site officiel, somme de contrôle vérifiée (MD5/SHA)
[ ] Espace flash suffisant (taille du .cc x 2 idéalement)
[ ] Sauvegarde config : save + copie externe (TFTP/SFTP)
[ ] Fenêtre de maintenance validée, utilisateurs prévenus (display users)
[ ] Plan de rollback écrit (ancien .cc conservé en flash)
[ ] Console accessible (en cas de problème au boot)
```

## 73. Transfert via TFTP (méthode simple)

Côté serveur : un serveur TFTP (tftpd-hpa sous Linux, Tftpd64 sous Windows)
avec le fichier `.cc` dans son répertoire racine.

```vrp
<AR720>ping 192.168.100.50                      # le serveur TFTP doit répondre
<AR720>tftp 192.168.100.50 get AR720-V300R022C00SPC500.cc
  ...
  TFTP: Downloading the file successfully. 84934656 bytes received.
<AR720>dir                                      # vérifier la taille du fichier reçu
```

Limites TFTP : pas d'authentification, UDP (moins fiable sur WAN), limite
historique ~32 Mo par transfert sur certaines implémentations — préférez
FTP/SFTP pour les gros `.cc` (> 100 Mo).

## 74. Transfert via FTP (méthode recommandée en LAN)

```vrp
<AR720>ftp 192.168.100.50
Trying 192.168.100.50 ...
Press CTRL+K to abort
Connected to 192.168.100.50.
220 FTP service ready.
User(192.168.100.50:(none)):ftpuser
331 Password required for ftpuser.
Password:
230 User logged in.
[ftp]get AR720-V300R022C00SPC500.cc
[ftp]bye
221 Server closing.
<AR720>dir
```

En une ligne (si supporté) :

```vrp
<AR720>copy ftp://ftpuser:motdepasse@192.168.100.50/AR720-V300R022C00SPC500.cc flash:/
```

> ⚠️ Le mot de passe en clair dans la commande reste dans l'historique :
> préférez la session FTP interactive, ou SFTP.

## 75. Transfert via SFTP / SCP (méthode sécurisée)

```vrp
[AR720]sftp server enable                       # si l'équipement est serveur (dépôt)
```

Pour **pousser** le fichier depuis un poste Linux vers l'équipement
(l'équipement en serveur SFTP) :

```bash
# Sur le poste Linux :
sftp admin@192.168.1.1
sftp> put AR720-V300R022C00SPC500.cc flash:/
sftp> ls -l
sftp> bye
```

Prérequis côté VRP : utilisateur avec `service-type ssh sftp`, `sftp
server enable`, clés RSA générées (`rsa local-key-pair create`). Vérifiez
la taille après transfert avec `dir`.

## 76. Transfert via clé USB

```vrp
<AR720>dir usb0:/
<AR720>copy usb0:/AR720-V300R022C00SPC500.cc flash:/
```

Pratique quand le réseau est indisponible ou pour un site distant. Formatez
la clé en FAT32 ; certains modèles n'acceptent que des clés ≤ 32 Go
(à vérifier selon modèle).

## 77. Désigner le logiciel de démarrage et rebooter

