---
id: collect-261001-general-networking/general-networking/fuite-de-donnees-24-md-d-identifiants-124m-reels-2026-2
title: "fuite-de-donnees-24-md-d-identifiants-124m-reels-2026"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Oracle"]
dates: ["2026-06"]
keywords: ["aws", "distribution", "incident"]
source: docs/RAG/collect-261001-general-networking/fuite-de-donnees-24-md-d-identifiants-124m-reels-2026.md
source_anchor: ""
source_lines: [43, 98]
sha256: 21ba30491b92bdd5f1a65d9ad8120d8c2e79cf7cf2df2cca8abd2300ba2bc498
---

# fuite-de-donnees-24-md-d-identifiants-124m-reels-2026

`curl https://api.pwnedpasswords.com/range/5BAA6`
Cette requête renvoie la liste des suffixes de hash correspondant au préfixe 5BAA6, avec le nombre de fois où chaque mot de passe a été vu dans des fuites connues. C’est justement le préfixe que donne le mot de passe « password » une fois haché, ce qui explique pourquoi il revient si souvent en exemple dans la documentation technique de développeurs.

## Historique des grandes compilations d’identifiants volés

Pour resituer l’ampleur de l’incident, voici comment la compilation de juin 2026 se compare aux principales méga-fuites recensées depuis 2019.

| Compilation | Année | Volume annoncé | Nature des données | 
|---|---|---|---|
| Collection #1 | 2019 | 773 millions | Emails et mots de passe compilés | 
| COMB (Compilation of Many Breaches) | 2021 | 3,2 milliards | Emails et mots de passe compilés | 
| Mother of All Breaches (MOAB) | 2024 | 26 milliards | Compilation multi-sources | 
| Compilation infostealer Cybernews | Juin 2025 | 16 milliards (brut) | Journaux infostealer | 
| June 2026 Stealer Logs | Juin 2026 | 24 Md brut / 124M mots de passe uniques | Journaux infostealer (Cybernews / HIBP) | 

La tendance est claire : les chiffres bruts gonflent d’année en année, mais l’écart entre volume brut et données réellement uniques grandit tout autant, à mesure que les mêmes identifiants sont recopiés d’une compilation à l’autre.

## Le marché noir des identifiants volés

Derrière chaque ligne de ce type de compilation se trouve une chaîne économique bien rodée. Un appareil s’infecte, souvent via un logiciel piraté ou une pièce jointe piégée. Le malware infostealer collecte les identifiants stockés dans le navigateur et les envoie à un serveur de commande. Ces journaux bruts sont ensuite vendus à l’unité ou par lot sur des forums spécialisés, puis republiés, recopiés, et mélangés à d’anciennes fuites sur des canaux Telegram dédiés au commerce de données bancaires et d’accès.

C’est cette republication en boucle qui explique pourquoi 1,7 milliard des lignes retrouvées dans le cluster de juin 2026 provenaient de canaux Telegram. Ces canaux fonctionnent comme des points de distribution, où les mêmes journaux circulent, se combinent avec d’autres, et gonflent artificiellement le volume total sans ajouter de victimes réellement nouvelles.

## Ce que cela change pour la France et l’Europe

Les chercheurs à l’origine de la découverte n’ont pas publié de répartition par pays pour ce jeu de données précis. Mais le contexte européen reste déterminant pour comprendre les conséquences. Le Règlement général sur la protection des données impose des obligations de notification aux entreprises victimes d’une violation touchant des données personnelles, et la directive NIS2, en cours de transposition dans plusieurs États membres, durcit encore les exigences pour les secteurs jugés critiques.

Cette compilation arrive alors que la France a déjà été confrontée à une série d’incidents distincts en 2026, dont la mise en vente de 250 000 passeports et cartes d’identité français sur le dark web. Le ministère de l’Intérieur a ainsi annoncé en avril 2026 une divulgation de données touchant l’ANTS et concernant 11,7 millions de comptes, tandis que shattered.io a signalé, en février 2026, l’exposition d’environ 15 millions de personnes liées à Cegedim Santé et de 1,2 million d’enregistrements associés à la DGFiP/FICOBA, dans la foulée d’une fuite ayant déjà touché 1,6 million de personnes chez France Travail en décembre 2025. Le pays est aussi partie prenante d’un contentieux européen : la France et trois autres pays ont été renvoyés devant la Cour de justice de l’Union européenne pour retard dans la transposition de NIS2, ce qui illustre la difficulté des États membres à faire appliquer uniformément ces règles pendant que le volume de données volées continue de croître, comme le rappelle l’agence européenne de cybersécurité ENISA.

## Comparaison avec les autres grandes fuites de 2026

La compilation de juin 2026 n’est pas un cas isolé. Elle s’ajoute à une série d’incidents distincts touchant des organisations et des citoyens européens au cours de l’année.

| Incident | Période | Échelle | Type | 
|---|---|---|---|
| June 2026 Stealer Logs | Juin 2026 | 124M mots de passe uniques (24 Md brut) | Compilation infostealer | 
| Passeports français sur le dark web | Juin 2026 | 250 000 documents | Vente de documents d’identité | 
| ShinyHunters vs Commission européenne | 2026 | 350 Go de données | Intrusion cloud AWS | 
| Faille Oracle PeopleSoft | 2026 | Environ 100 organisations ciblées | Vulnérabilité CVSS 9,8 | 

La juxtaposition de ces incidents montre une même logique à l’œuvre. La donnée volée circule, se recombine, et repart. Le groupe de revendication Ostraca a par exemple annoncé le 10 juin 2026 détenir 3,6 millions d’enregistrements issus de CapiFrance, avant de revendiquer, le 2 juillet 2026, 3,8 millions de personnes et 966 Go de documents liés à IAD Group ; plus récemment encore, FrenchBreaches a signalé le 18 août 2026 une fuite touchant Foodtrack.fr. Contrairement à la fuite de passeports français, qui vise des documents directement exploitables pour de la fraude, ou à l’intrusion attribuée à ShinyHunters contre la Commission européenne, qui cible une institution précise, la compilation de juin 2026 n’a pas de victime unique désignée. Elle agrège les dégâts de milliers d’infections individuelles, ce qui la rend à la fois moins spectaculaire à raconter et plus difficile à corriger d’un coup.

## 91 % de données déjà vues : le recyclage permanent des identifiants

Un travail publié par le chercheur Benjamin Brundage, de la société Synthient, sur un jeu de données comparable de 183 millions d’identifiants, a montré que 91 % de ces identifiants avaient déjà été observés dans des fuites antérieures. L’essentiel de ce qui circule aujourd’hui n’est donc pas neuf. Ce sont souvent les mêmes mots de passe, parfois vieux de plusieurs années, qui repassent de main en main entre acteurs malveillants.

Ce recyclage a une conséquence directe pour les utilisateurs qui réutilisent un même mot de passe sur plusieurs services. Chaque nouvelle compilation, même composée en majorité d’anciennes données, leur fait courir un risque de bourrage d’identifiants (credential stuffing) tant qu’ils n’ont pas changé ce mot de passe partout où ils l’utilisent.

## La riposte des autorités : Europol et l’opération Endgame

Face à cette économie de la donnée volée, les autorités européennes ne restent pas inactives. Europol a mené l’opération Endgame, une action coordonnée contre les infrastructures utilisées pour diffuser des logiciels infostealers et d’autres familles de malwares, qui a permis la saisie de 326 serveurs utilisés par des réseaux criminels. Cette opération illustre la stratégie privilégiée par les forces de l’ordre européennes : s’attaquer à l’infrastructure de distribution plutôt qu’à chaque fuite de données individuellement, sur le principe qu’un serveur de commande démantelé coupe l’alimentation de dizaines de compilations futures.

Le problème, reconnu par les chercheurs en sécurité, est que ces infrastructures se reconstituent rapidement. Les groupes derrière les infostealers changent de nom, migrent vers de nouveaux hébergeurs, et reprennent leur activité en quelques semaines. La compilation de juin 2026, qui inclut des données bien antérieures à l’opération Endgame, montre que le stock de données déjà volées continue de circuler longtemps après le démantèlement d’un réseau.

## Have I Been Pwned face à la concurrence

