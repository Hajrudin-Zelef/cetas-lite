---
id: collect-261001-general-networking/general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026-3
title: "Windows (PowerShell)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "mai"]
source: docs/RAG/collect-261001-general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [105, 144]
sha256: a7c5d7d97239cbb8c459201d8fb57b1165c5f89e9fc6ac40de6b16563321fdf5
---

# Windows (PowerShell)

Une fois le travail terminé, ne vous contentez jamais d’éjecter le volume comme une simple clé USB : utilisez le bouton **Dismount** (rebaptisé **Unmount All** pour l’option globale depuis la mise à jour 1.26.24 du 30 mai 2025, qui a aussi actualisé les traductions chinoise et russe de l’interface) dans VeraCrypt. Cette étape force l’écriture de toutes les données en attente et referme proprement le volume chiffré. Un démontage brutal (extinction de l’ordinateur, débranchement d’un disque externe) pendant qu’un volume reste monté peut corrompre les données qu’il contient, exactement comme pour n’importe quel système de fichiers.

## Étape 5 : Chiffrer une clé USB ou un disque externe entier

Plutôt qu’un simple conteneur, vous pouvez chiffrer un périphérique complet : une clé USB, un disque dur externe ou une carte SD. Dans l’assistant de création, choisissez cette fois **Encrypt a non-system partition/drive**, puis sélectionnez le périphérique cible via **Select Device** plutôt que **Select File**.

Soyez particulièrement vigilant sur le choix du périphérique à cette étape : VeraCrypt formate entièrement le support sélectionné, effaçant toutes les données déjà présentes. Vérifiez deux fois la lettre de lecteur ou l’identifiant du périphérique avant de valider, surtout si plusieurs disques externes sont branchés simultanément. Une fois le chiffrement terminé, le périphérique doit être monté via VeraCrypt à chaque branchement : il n’apparaîtra pas comme un disque lisible normalement dans l’explorateur de fichiers d’un ordinateur ne disposant pas de VeraCrypt.

## Étape 6 : Chiffrer le disque système complet (Windows)

Le chiffrement du disque système protège l’intégralité de Windows, y compris le fichier d’échange, les fichiers temporaires et les données de mise en veille prolongée, qui peuvent tous contenir des fragments de données sensibles. Dans l’assistant, choisissez **Encrypt the system partition or entire system drive**, puis **Normal** (le mode caché est traité à l’étape 9).

- Sélectionnez **Encrypt the whole drive** si votre disque ne contient qu’une seule partition Windows, ou**Encrypt the Windows system partition** pour ne cibler que la partition système en cas de multi-boot
- Sur un poste avec un seul système d’exploitation installé, répondez **Single-boot** à la question posée par l’assistant
- Choisissez l’algorithme de chiffrement et la fonction de hachage, comme pour un conteneur classique
- Définissez un mot de passe de pré-démarrage (pre-boot authentication password) : c’est ce mot de passe qui sera demandé avant même le chargement de Windows
- Laissez VeraCrypt collecter de l’entropie en bougeant la souris, puis passez à la génération des clés

## Étape 7 : Créer le disque de secours (Rescue Disk), l’étape qu’il ne faut jamais sauter

Avant de chiffrer quoi que ce soit sur le disque système, VeraCrypt oblige à créer un **VeraCrypt Rescue Disk**, sous la forme d’une image ISO à graver ou à copier sur une clé USB bootable. Ce disque de secours n’est pas une option accessoire : si le chargeur de démarrage est un jour endommagé (mise à jour de firmware ratée, secteur de démarrage corrompu, panne de courant pendant une opération sensible), c’est l’unique moyen de redémarrer le système ou de déchiffrer le disque en cas de problème durable.

Une fois l’image générée, VeraCrypt propose de vérifier automatiquement que le disque de secours démarre correctement avant de continuer. Ne sautez jamais cette vérification : un disque de secours corrompu ou mal gravé ne sert à rien le jour où vous en avez réellement besoin, et ce jour arrive toujours au pire moment. Conservez ce support dans un endroit différent de l’ordinateur lui-même, par exemple avec vos documents importants, plutôt que dans le même sac ou le même tiroir.

## Étape 8 : Test de pré-chiffrement et lancement du chiffrement système

Avant de chiffrer réellement le disque, VeraCrypt propose un **test de pré-chiffrement** (pretest) : l’ordinateur redémarre, affiche l’écran de mot de passe de pré-démarrage, puis revient sous Windows sans qu’aucune donnée n’ait encore été modifiée. Ce test confirme que l’écran d’authentification fonctionne correctement sur votre configuration matérielle précise (clavier, résolution d’écran, pilotes de démarrage) avant que le chiffrement irréversible ne commence.

Si le test réussit, cliquez sur **Encrypt** pour lancer le chiffrement effectif. Une barre de progression affiche l’avancement en pourcentage et en Mo restants. Vous pouvez continuer à utiliser l’ordinateur normalement pendant le chiffrement : l’opération se déroule en tâche de fond, avec un léger ralentissement des accès disque. Il est possible de mettre le chiffrement en pause via VeraCrypt et de le reprendre plus tard, ce qui est utile sur un ordinateur portable si la batterie devient faible malgré la recommandation de rester branché au secteur.

## Étape 9 : Créer un volume caché pour le déni plausible

La fonctionnalité de volume caché (hidden volume) permet de dissimuler un second volume chiffré à l’intérieur de l’espace libre d’un premier volume, dit « volume extérieur ». Sans connaître le mot de passe du volume caché, il est mathématiquement impossible de prouver qu’il existe : l’espace qu’il occupe apparaît comme de l’espace libre non utilisé du volume extérieur. C’est la base du concept de **déni plausible** : si quelqu’un vous contraint à révéler un mot de passe, vous pouvez communiquer celui du volume extérieur, qui contient des données plausibles mais non sensibles, sans révéler l’existence du volume caché.

Pour créer un volume caché, choisissez **Hidden VeraCrypt volume** dans l’assistant de création plutôt que **Standard**. VeraCrypt vous demande de créer d’abord le volume extérieur, de le remplir avec des fichiers crédibles (jamais un volume vide, ce qui éveillerait immédiatement les soupçons), puis de définir la taille et le mot de passe du volume caché à l’intérieur. Un point technique important, confirmé par les notes de version de VeraCrypt 1.26.29 : la structure actuelle prend en charge **un volume caché à l’intérieur d’un volume extérieur**, et non un empilement de plusieurs couches cachées imbriquées les unes dans les autres. Si vous avez créé des volumes cachés avec une version antérieure à 1.26.6, ou entre 1.26.6 et 1.26.29, consultez les notes de version officielles : la base nationale américaine de vulnérabilités (NIST NVD) a confirmé le 21 août 2026 que ces versions intermédiaires étaient exposées à la faille CVE-2026-54073, susceptible d’affecter la déniabilité plausible et corrigée depuis dans la 1.26.29 ; les volumes créés dans cet intervalle doivent être recréés par précaution.

Attention à un piège technique classique : si vous montez le volume extérieur et y écrivez de nouveaux fichiers sans protection particulière, vous risquez d’écraser silencieusement le volume caché qu’il contient. Lorsque vous devez modifier le volume extérieur après coup, montez-le en cochant l’option **Protect hidden volume** et en indiquant le mot de passe du volume caché : VeraCrypt empêchera alors toute écriture qui risquerait d’endommager les données cachées.

## Étape 10 : Renforcer l’authentification avec des fichiers-clés (keyfiles)

