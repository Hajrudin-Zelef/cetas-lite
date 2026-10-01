---
id: collect-261001-general-networking/general-networking/memtest86-tester-sa-ram-en-13-etapes-2026-2
title: "memtest86-tester-sa-ram-en-13-etapes-2026"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["arr", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/memtest86-tester-sa-ram-en-13-etapes-2026.md
source_anchor: ""
source_lines: [37, 102]
sha256: db8a7e75f966b815f809559789df82b067599fe4d542742243c9613693fb3728
---

# memtest86-tester-sa-ram-en-13-etapes-2026

Avant de commencer, réunissez les éléments suivants. La liste inclut les versions vérifiées au moment de la rédaction (MemTest86 11.7 Build 1000, référencé par Uptodown sous le paquet Windows 1.5.1004.0 depuis le 6 janvier 2026, après les 11.5 Build 1000 du 18 septembre 2025, 11.3 du 7 mai 2025 puis 11.2 du 5 février 2025), à utiliser telles quelles ou en version ultérieure. Le développement de l’écosystème ne s’arrête pas là : côté fork open source, la branche de développement de Memtest86+ affichait encore un commit récent (référence #7726853) daté de septembre 2026 sur la page dev de Memtest.org, signe que les outils de diagnostic mémoire continuent d’évoluer au-delà de cette version de référence.

| Outil | Version testée | Usage dans ce tutoriel | 
|---|---|---|
| MemTest86 (PassMark) | 11.7 (Build 1000) ou ultérieure | Test de RAM principal, édition Free | 
| Image USB (PassMark) | Fournie avec le pack MemTest86 | Création de la clé USB officielle | 
| Rufus (alternative) | 4.14 ou ultérieure | Création manuelle de la clé USB bootable | 
| Clé USB | 1 Go minimum | Support de démarrage, entièrement effacé | 
| BIOS/UEFI | Mode UEFI activé | Requis par MemTest86 v11.7 (le mode Legacy impose l’ancienne édition V4) | 
| Windows | 10 ou 11 | Système utilisé pour préparer la clé USB | 

Un point mérite d’être clarifié avant de télécharger quoi que ce soit : MemTest86 v11.7 ne démarre qu’en UEFI. Sur une machine ancienne qui ne propose pas ce mode, PassMark met à disposition une édition V4 distincte, basée sur l’ancien BIOS historique. Vérifiez ce point dans le BIOS de votre carte mère avant de perdre du temps avec la mauvaise image.

Sauvegardez également le contenu de la clé USB choisie : la création de la clé de démarrage efface entièrement son contenu existant, sans confirmation supplémentaire une fois le processus lancé. Comptez environ 45 à 60 minutes pour dérouler l’ensemble des 13 étapes de ce tutoriel jusqu’à un premier verdict, sachant qu’une validation complète et fiable demande de laisser tourner le test plus longtemps, souvent toute une nuit.

## Étapes 1 à 4 : préparer et créer votre clé USB MemTest86

**Étape 1 : identifier les symptômes qui justifient le test.** Avant de perdre du temps, notez les circonstances exactes des plantages : toujours sous charge, toujours au repos, uniquement en jeu, ou dès le démarrage. Un crash reproductible dans un seul logiciel oriente plutôt vers un pilote ou une application. Des plantages erratiques, sans lien avec une tâche précise, ou des fichiers qui se corrompent silencieusement, pointent davantage vers la mémoire ou le stockage. Si le doute persiste entre les deux, gardez en tête un contrôle de santé SSD avec CrystalDiskInfo comme second diagnostic possible.

**Étape 2 : télécharger la bonne édition de MemTest86.** Rendez-vous sur la page de téléchargement officielle de MemTest86 et choisissez l’édition Free pour un usage personnel : sur MajorGeeks, ce téléchargement freemium ne pèse que 12 Mo et reste répertorié compatible de Windows 7 à Windows 11, fiche mise à jour le 3 mai 2026 pour la 11.7 Build 1000. Elle couvre l’essentiel : test multi-thread, détection UEFI, rapport d’erreurs à l’écran. Les éditions Pro (58 $US, environ 54 €) et Site (5 557 $US) ajoutent l’export de rapports, l’automatisation par script et le démarrage réseau PXE pour tester plusieurs machines à la fois, des fonctions surtout utiles en entreprise ou en atelier de réparation.

**Étape 3 : vérifier l’intégrité du fichier téléchargé.** Un fichier corrompu pendant le téléchargement peut produire une clé USB qui ne démarre pas ou, pire, qui donne de faux résultats. Sous Windows, comparez le hash SHA-256 du fichier téléchargé avec celui publié sur le site de PassMark :

`certutil -hashfile memtest86-usb-installer.exe SHA256`
Si la chaîne affichée ne correspond pas exactement à celle du site officiel, retéléchargez le fichier avant d’aller plus loin.

**Étape 4 : créer la clé USB bootable.** La méthode la plus simple consiste à lancer l’utilitaire Image USB fourni dans le pack MemTest86 : sélectionnez votre clé, cochez la case de création automatique, validez. En quelques minutes, la clé est prête et reconnue comme périphérique de démarrage UEFI. Si vous préférez une méthode manuelle ou si l’utilitaire fourni ne détecte pas votre clé, Rufus (version 4.14 ou ultérieure) reste l’alternative de référence : sélectionnez l’image ISO de MemTest86, laissez le schéma de partition sur GPT pour UEFI, et lancez la création.

Pour les utilisateurs à l’aise en ligne de commande, une clé bootable peut aussi se préparer manuellement avec l’outil intégré à Windows. Cette méthode demande de la rigueur : une erreur sur le numéro de disque efface le mauvais support.

```
diskpart
list disk
select disk 2
clean
create partition primary
active
format fs=fat32 quick
assign
exit
```
Remplacez le numéro « 2 » par celui de votre clé USB, identifié grâce à sa taille dans la sortie de `list disk`. Une fois le formatage terminé, copiez le contenu de l’image MemTest86 à la racine de la clé.

## Étapes 5 à 7 : configurer le BIOS/UEFI pour démarrer sur la clé USB

**Étape 5 : accéder au BIOS.** Redémarrez le PC et appuyez sur la touche d’accès au BIOS dès l’écran de démarrage constructeur. Cette touche varie selon le fabricant de la carte mère ou du portable.

| Fabricant | Touche BIOS/UEFI | Touche menu de boot rapide | 
|---|---|---|
| ASUS | Suppr (Del) ou F2 | F8 | 
| MSI | Suppr (Del) | F11 | 
| Gigabyte | Suppr (Del) | F12 | 
| ASRock | F2 ou Suppr (Del) | F11 | 
| Dell | F2 | F12 | 
| HP | Échap (Esc) puis F10 | F9 | 
| Lenovo | F1 ou F2 | F12 | 
| Acer | F2 ou Suppr (Del) | F12 | 

**Étape 6 : vérifier le Secure Boot.** MemTest86 v11.7 est signé numériquement et démarre normalement même avec le Secure Boot actif. Sur certaines configurations strictes ou d’anciennes révisions de BIOS, il arrive que le démarrage soit malgré tout bloqué. Si la clé USB n’apparaît pas dans la liste des périphériques de démarrage, désactivez temporairement le Secure Boot dans l’onglet Sécurité du BIOS, testez, puis réactivez-le une fois le diagnostic terminé.

**Étape 7 : modifier l’ordre de démarrage.** Deux options s’offrent à vous : passer par le menu de boot rapide (touche dédiée listée dans le tableau ci-dessus) pour un démarrage ponctuel sans toucher à la configuration permanente, ou entrer dans l’onglet Boot du BIOS pour placer la clé USB en première position de la liste. La première méthode est recommandée pour ce tutoriel : elle évite d’oublier de remettre le disque système en priorité après le test.

## Étapes 8 et 9 : lancer et configurer MemTest86

**Étape 8 : le menu de configuration au démarrage.** Une fois la clé USB sélectionnée, MemTest86 affiche un écran de configuration avant de lancer le moindre test. C’est ici que vous choisissez la disposition du clavier et, si besoin, la langue de l’interface (le français fait partie des langues proposées). Prenez le temps de vérifier que la quantité de RAM détectée correspond bien à ce qui est physiquement installé : un écart signale déjà un problème, avant même le premier test.

