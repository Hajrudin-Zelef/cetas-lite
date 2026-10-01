---
id: collect-261001-general-networking/general-networking/fuite-de-donnees-24-md-d-identifiants-124m-reels-2026-1
title: "fuite-de-donnees-24-md-d-identifiants-124m-reels-2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2026-06"]
keywords: ["data breach", "incident"]
source: docs/RAG/collect-261001-general-networking/fuite-de-donnees-24-md-d-identifiants-124m-reels-2026.md
source_anchor: ""
source_lines: [1, 42]
sha256: beaf82baefd0b14f8f3498066db75a840d3c4cffbe738860fe7befeb649c492c
---

# fuite-de-donnees-24-md-d-identifiants-124m-reels-2026

Le 15 juin 2026, Have I Been Pwned a ajouté à sa base un nouveau jeu de données baptisé « June 2026 Stealer Logs ». Le chiffre affiché à côté donne le vertige : 24 milliards de lignes d’identifiants. Mais ce chiffre brut cache une réalité plus nuancée, et c’est justement cette nuance qui occupe les professionnels de la cybersécurité en France et en Europe cette semaine.

Derrière l’annonce se trouve un cluster Elasticsearch resté ouvert sur internet, repéré par les chercheurs de Cybernews, et rempli de journaux volés par des logiciels infostealers. Une fois dédupliquées, les données représentent 56,3 millions d’adresses e-mail uniques et 124 millions de mots de passe uniques : un total déjà considérable, mais dix fois inférieur à l’annonce initiale. Cette fuite de données s’inscrit dans une tendance longue, celle des compilations géantes d’identifiants volés qui se multiplient depuis 2019, dont la taille brute sert de plus en plus d’argument de communication pour ceux qui les découvrent.

Cet article détaille ce que l’on sait de cette fuite de données, pourquoi l’écart entre 24 milliards et 124 millions compte autant, ce que cela change pour les entreprises et les particuliers européens, et comment vérifier concrètement son exposition.

## Que révèle la fuite « June 2026 Stealer Logs » ajoutée à Have I Been Pwned

Le nom donné par Have I Been Pwned (HIBP) résume bien la nature du jeu de données. Il ne s’agit pas de mots de passe issus d’un piratage unique, mais de journaux collectés par des logiciels malveillants de type infostealer sur des ordinateurs infectés dans le monde entier. Ces programmes aspirent tout ce qu’ils trouvent dans le navigateur d’un poste compromis, identifiants enregistrés et cookies de session actifs compris, parfois même des captures d’écran ou l’historique du presse-papiers.

Troy Hunt, qui exploite Have I Been Pwned depuis 2013, a intégré le jeu de données le 15 juin 2026 sous le nom « June 2026 Stealer Logs ». Environ 56 millions d’utilisateurs peuvent désormais vérifier leur exposition dans la section dédiée aux stealer logs de leur tableau de bord, et les organisations disposent d’un accès via l’API réservée aux domaines professionnels.

Le dataset ne provient pas d’une entreprise unique piratée. Il agrège des identifiants volés sur des appareils infectés partout dans le monde, compilés à partir d’environ 36 sources distinctes : anciennes compilations de fuites, dépôts publiés sur des forums de piratage, et canaux Telegram spécialisés dans la revente de données bancaires et d’identifiants. Ces derniers représenteraient à eux seuls près de 1,7 milliard des lignes brutes du cluster.

## 24 milliards de lignes, 124 millions de mots de passe : d’où vient l’écart

Le chiffre de 24 milliards, largement repris, correspond au nombre brut de lignes présentes dans le cluster Elasticsearch, soit environ 8,3 téraoctets de données. Une grande partie de ce volume est redondante. Les mêmes identifiants apparaissent des dizaines, parfois des centaines de fois, à force d’être recopiés d’une compilation à l’autre depuis plusieurs années.

Une fois le nettoyage effectué, HIBP retient 56,3 millions d’adresses e-mail uniques et 124 millions de mots de passe uniques. C’est cet écart qui a poussé plusieurs analystes à qualifier le chiffre de 24 milliards de trompeur, davantage un argument de communication qu’une mesure fidèle du risque réel. Une analyse de Phishing Tackle consacrée à l’incident souligne que la majorité des identifiants stockés dans ce type de cluster ont déjà circulé ailleurs, ce qui réduit leur nouveauté opérationnelle pour les attaquants, mais pas leur dangerosité pour les victimes qui n’ont toujours pas changé leurs mots de passe.

Cette confusion entre volume brut et exposition réelle n’est pas nouvelle. Elle explique en partie pourquoi les titres autour des fuites de données massives restent difficiles à interpréter pour le grand public, et pourquoi la lecture des chiffres dédupliqués demeure la seule façon fiable d’évaluer un risque individuel.

## Comment Cybernews a mis au jour le cluster non sécurisé

Mi-juin 2026, les chercheurs de Cybernews sont tombés sur un cluster Elasticsearch accessible sans authentification. Ce type d’erreur de configuration, banal mais dévastateur, revient régulièrement dans l’actualité de la cybersécurité. Une base de données pensée pour un usage interne, mise en ligne sans mot de passe ni pare-feu, devient consultable par n’importe qui muni de la bonne adresse IP.

Une fois le signalement effectué, le cluster a été mis hors ligne. Mais le temps qu’il reste exposé suffit largement. Des scanners automatisés parcourent en permanence les plages d’adresses IP publiques à la recherche de ce genre d’oubli, si bien qu’une base non protégée peut être repérée par des acteurs malveillants en quelques heures à peine, bien avant qu’un chercheur en sécurité ne la signale publiquement.

Aucune action réglementaire ou judiciaire spécifique à cet incident n’a été rendue publique à ce stade. Il ne s’agit pas d’un piratage d’entreprise identifiable, ce qui complique la question de la responsabilité : qui doit notifier qui, quand la fuite de données est elle-même une compilation de fuites déjà anciennes ?

## Qu’est-ce qu’un stealer log, et pourquoi ces compilations explosent

Un infostealer est un logiciel malveillant conçu pour un seul objectif : extraire discrètement tout ce qui a de la valeur sur un poste infecté, puis l’envoyer à son opérateur. Contrairement à un rançongiciel, il ne chiffre rien et ne réclame aucune rançon. Il agit en silence, souvent pendant des semaines, avant que la victime ne se rende compte de quoi que ce soit.

Le marché de l’infostealer a connu une croissance spectaculaire. Selon les données de Flashpoint, environ 1,8 milliard d’identifiants ont été volés en 2025 via des appareils infectés par ce type de malware, un total obtenu à partir de 5,8 millions de machines compromises, soit une hausse de 800 % sur une période de quatre mois seulement. Le Data Breach Investigations Report 2026 de Verizon, qui couvre les incidents survenus entre novembre 2024 et octobre 2025 et reste, en juin 2026, la référence du secteur, confirme cette dynamique en plaçant les identifiants volés parmi les vecteurs d’intrusion les plus cités. Cette explosion alimente directement les compilations comme celle découverte en juin 2026 : plus il y a d’appareils infectés, plus les journaux volés s’accumulent sur les forums et les canaux Telegram, et plus il devient tentant pour quelqu’un de les regrouper en un seul jeu de données consultable.

## Vérifier son exposition sur Have I Been Pwned

La meilleure façon de savoir si l’on figure dans cette fuite de données reste de consulter directement Have I Been Pwned avec son adresse e-mail. Le service, gratuit pour un usage individuel, indique précisément si une adresse apparaît dans le jeu de données « June 2026 Stealer Logs » ou dans l’une des centaines d’autres compilations déjà indexées.

Pour les mots de passe, HIBP propose depuis plusieurs années une API en k-anonymat, qui permet de vérifier si un mot de passe a fuité sans jamais transmettre ce mot de passe en clair. Le principe : on hache le mot de passe en SHA-1, on envoie uniquement les cinq premiers caractères du hash, et le service retourne toutes les correspondances possibles pour que la comparaison finale se fasse en local, côté client.

