---
id: collect-261001-huawei/huawei/fuite-de-donnees-france-2e-23-5m-comptes-voles-2026-1
title: "Aucune verification des droits d'acces : il suffit d'incrementer le numero"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-huawei/fuite-de-donnees-france-2e-23-5m-comptes-voles-2026.md
source_anchor: ""
source_lines: [1, 35]
sha256: 75b08d650e87631367910ac7ddfe0c7a9495d1f27124926c2c74b55caf98ef9e
---

# Aucune verification des droits d'acces : il suffit d'incrementer le numero

La France a basculé, en six mois, au rang de cible privilégiée de la cybercriminalité mondiale. Selon l’étude de surveillance des violations de données publiée par Surfshark, relayée par **BFMTV**, **43,4 millions de comptes français** ont été compromis sur l’ensemble du premier semestre 2026, soit une hausse de **62,3 %** – dont **23,5 millions** pour le seul premier trimestre, en bond de **108,6 %** par rapport au dernier trimestre 2025. Ce volume hisse l’Hexagone à la **deuxième place mondiale** des pays les plus touchés par une **fuite de données**, juste derrière les États-Unis et leurs 60,3 millions de comptes exposés, devant l’Inde et le Brésil.

Derrière ces chiffres se cache une succession d’incidents retentissants : le piratage de l’**ANTS** et ses 11,7 millions de comptes officiels exposés, la fuite d’**ÉduConnect** touchant 3,5 millions d’élèves, et, plus récemment, la compromission de **Tchap**, la messagerie chiffrée de l’État. Une même classe de vulnérabilité – la faille **IDOR** – revient comme un fil rouge. Tour d’horizon d’une crise qui redéfinit la **cyberattaque en France** en 2026, son impact sur le marché et la réponse réglementaire européenne.

## Fuite de données en France : 23,5 millions de comptes compromis au premier trimestre 2026

Le constat dressé par Surfshark est sans appel, un diagnostic également repris par **Les Numériques** et le tracker **FrenchBreaches**. Avec 23,5 millions de comptes compromis sur les trois premiers mois de l’année, la France enregistre une accélération brutale : la densité des fuites a plus que doublé en un trimestre. À l’échelle planétaire, l’étude estime qu’environ **27 comptes sont compromis chaque seconde**. Rapportée à sa population, la France concentre désormais une part disproportionnée de ce flux. Ce diagnostic est corroboré outre-Atlantique par l’**Identity Theft Resource Center** (ITRC), dont le rapport H1 2026, publié en juillet 2026 et relayé par **HIPAA Journal**, **Security Magazine** et **CNBC**, recense **1 803 compromissions** sur les six premiers mois de l’année aux États-Unis – une cadence qui projette plus de **3 600 incidents** sur l’ensemble de 2026 – et comptabilise **471,2 millions d’avis aux victimes** envoyés sur ce seul semestre, contre 297,5 millions sur toute l’année 2025.

Le contexte historique aggrave le tableau. Depuis 2004, selon la même source, ce sont **740,9 millions de comptes** français qui ont fuité, correspondant à **211,9 millions d’adresses e-mail uniques** et à près de **2,1 milliards d’enregistrements** de données personnelles. Dans le détail figurent 488,8 millions de mots de passe, 196,8 millions de noms d’utilisateur, 77,4 millions d’adresses postales, 65,2 millions de numéros de téléphone, 643 300 numéros de sécurité sociale et près de 19 700 coordonnées bancaires. Sur l’ensemble de la période, la France occupe la quatrième position mondiale ; sur le seul premier trimestre 2026, elle bondit à la deuxième. Un baromètre fondé sur les données de la **CNIL**, relayé par le forum **InCyber**, dénombre **8 613 notifications** de violations de données personnelles entre septembre 2024 et septembre 2025, soit une hausse de **45 %** ; le tracker **Fuites Infos** recense quant à lui **163 fuites** en 2026, ayant affecté **733 469 422 personnes** depuis janvier 2025, tandis qu’**itsense.fr** dénombre plus de **140 fuites** répertoriées en France sur le seul premier trimestre 2026. À titre de comparaison, la **Privacy Rights Clearinghouse** américaine a recensé, dans un rapport publié en août 2026, **4 080 fuites de données uniques** survenues en 2025, touchant au moins **375 millions de personnes** à travers 8 019 notifications distinctes – preuve que l’inflation des violations de données dépasse très largement le seul cas français.

Cette flambée n’a rien d’un accident statistique. Elle traduit une bascule structurelle : les attaquants ne ciblent plus seulement les entreprises privées, mais s’attaquent désormais frontalement aux **grands portails publics**, dépositaires des bases de données les plus volumineuses et les plus sensibles du pays. L’année 2026 restera celle où la **fuite de données administrative** est devenue le mode opératoire dominant, supplantant souvent le rançongiciel classique.

## Piratage de l’ANTS : 11,7 millions de Français exposés par une faille IDOR

L’incident emblématique de l’année reste le piratage de l’**ANTS** (Agence nationale des titres sécurisés, désormais France Titres), l’organisme qui gère cartes d’identité, passeports, permis de conduire et cartes grises. Détectée le **15 avril 2026** et rendue publique le 20 avril, l’attaque a conduit le ministère de l’Intérieur à reconnaître l’exposition de **11,7 millions de comptes**. L’auteur, un pirate opérant sous le pseudonyme **« breach3d »**, revendique pour sa part une base comprise **entre 18 et 19 millions d’enregistrements**, mise en vente sur des forums cybercriminels.

D’après la communication officielle de l’ANTS, les données concernées incluent les noms et prénoms, les dates de naissance, les adresses électroniques, les identifiants de connexion et les identifiants uniques de compte. Les pièces justificatives elles-mêmes (scans de documents) n’auraient pas été compromises. Conformément aux articles 33 et 34 du RGPD, l’agence était tenue de notifier la CNIL dans un délai de 72 heures et d’informer les personnes concernées en cas de risque élevé.

Le plus alarmant tient à la trivialité de l’exploitation. Interrogé sur la méthode employée, le pirate a résumé la situation d’une formule cinglante : *« C’était une faille vraiment stupide »*. Il suffisait, selon lui, de *« modifier un chiffre dans une requête pour accéder aux données d’un autre citoyen »*. Aucun contrôle d’autorisation n’était implémenté côté serveur. C’est la définition même d’une faille **IDOR**.

### Qu’est-ce qu’une faille IDOR et pourquoi elle est si répandue

Une faille **IDOR** (Insecure Direct Object Reference, ou « référence directe à un objet non sécurisée ») survient lorsqu’une application expose un identifiant interne – un numéro de dossier, un identifiant utilisateur – sans vérifier que la personne qui le demande a bien le droit d’y accéder. En pratique, le serveur répond à n’importe quelle requête tant que l’identifiant existe :

```
GET /api/v1/usagers/100123/titres   ->  dossier de l'usager 100123 (le vôtre)
GET /api/v1/usagers/100124/titres   ->  dossier de l'usager 100124 (un inconnu)
GET /api/v1/usagers/100125/titres   ->  dossier de l'usager 100125 (un autre inconnu)
# Aucune verification des droits d'acces : il suffit d'incrementer le numero
# pour parcourir l'integralite de la base, usager apres usager.
```
L’IDOR figure parmi les vulnérabilités les plus courantes du top 10 de l’OWASP, sous la catégorie « Broken Access Control ». Elle est redoutable parce qu’elle ne nécessite ni outil sophistiqué, ni logiciel malveillant : un simple navigateur et un peu de patience suffisent. Pour les administrations qui ont massivement numérisé leurs services sans toujours auditer le contrôle d’accès de leurs API, c’est le talon d’Achille parfait. Renforcer cette couche logicielle – comme le permettent des outils défensifs tels que CrowdSec côté infrastructure – ne suffit pas si l’autorisation applicative reste absente.

## ÉduConnect, CROUS, Compas : l’Éducation nationale en première ligne

