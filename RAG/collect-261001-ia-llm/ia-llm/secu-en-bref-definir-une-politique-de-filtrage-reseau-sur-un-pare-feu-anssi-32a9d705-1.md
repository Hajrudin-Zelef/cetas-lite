---
id: collect-261001-ia-llm/ia-llm/secu-en-bref-definir-une-politique-de-filtrage-reseau-sur-un-pare-feu-anssi-32a9d705-1
title: "Exemple de règle Section 1"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-ia-llm/secu-en-bref-definir-une-politique-de-filtrage-reseau-sur-un-pare-feu-anssi-32a9d705.md
source_anchor: ""
source_lines: [1, 72]
sha256: a7f51f7cbce7c55328f1fb1962117324a9e9afca62ed0f93ed742637e8e6bf80
---

# Exemple de règle Section 1

Sécu en bref : définir une politique de filtrage réseau sur un pare-feu selon l’ANSSI
Cet article décrypte les points clés du guide ANSSI "Recommandations pour la définition d'une politique de filtrage réseau d'un pare-feu" pour vous aider à déterminer la bonne stratégie de filtrage à appliquer sur votre pare-feu.
Vous souhaitez mettre en place une politique de filtrage stricte et efficace sur un SI, mais ne savez pas par où commencer ? Nous allons ici voir les principales bonnes pratiques à travers la revue du guide "Recommandations pour la définition d’une politique de filtrage réseau d’un pare-feu".
Ce guide de l'ANSSI permet d'obtenir les éléments organisationnels pour construire une politique de filtrage correctement structurée. Cela signifie qu’il vous guide dans la conception de votre politique de filtrage en vous indiquant les bonnes pratiques à adopter et celles à éviter.
Il contient 15 recommandations à utiliser comme guidelines et check-list pour valider que la conception et la mise en place de votre politique de filtrage respectent les bonnes pratiques de sécurité.
L'objectif étant de rendre votre politique de filtrage robuste, cohérente et plus facile à maintenir dans le temps. Voyons ici les principales idées et les apprentissages de ce document.
Retrouvez d'autres articles de la série Sécu' en Bref, parmi lesquels :
- Sécu’ en bref – ANSSI : bien architecturer un système de journalisation
- Sécu’ en bref – ANSSI : l’externalisation des SI
Règles, politique, pare-feu et filtrage
Avant de commencer à identifier les principaux points d'attention concernant les bonnes pratiques de conception d'une politique de filtrage, commençons par un bref rappel des termes qui s'y rapportent :
- Règle de filtrage : c'est une instruction appliquée par un pare-feu pour autoriser ou interdire un flux réseau en fonction de critères précis : adresse source, adresse destination, protocole, port, etc. Elle est évaluée séquentiellement dans le cadre d’une politique de filtrage.
- Pare-feu : équipement d’interconnexion qui contrôle les flux réseau entre des zones de confiance distinctes (ex : réseau interne, DMZ, internet). Il applique une politique de filtrage.
- Politique de filtrage : ensemble structuré de règles de filtrage appliquées à un pare-feu. Elle définit quels flux sont autorisés ou interdits entre les différentes zones d’un système d’information. Une politique de filtrage doit être cohérente, documentée et régulièrement révisée.
- Filtrage : processus technique d’analyse et de contrôle des flux réseau en fonction de critères prédéfinis (adresses IP, ports, protocoles). C'est la décision du pare-feu sur chaque paquet qu'il fait transiter.
- Segmentation : division d’un réseau en sous-réseaux ou zones logiques distinctes, afin de limiter la propagation des attaques et de réduire la surface d’exposition. Exemple : séparation des réseaux utilisateurs, serveurs, et invités.
- Cloisonnement : terme général qui combine la segmentation et le filtrage. Il désigne l’ensemble des mesures visant à isoler les zones du réseau et à contrôler les communications entre elles pour réduire la surface d’attaque.
Ces trois dernières notions sont explorées plus en détail dans l'article suivant :
La complexité des SI actuels et la multitude de réseaux, sous-réseaux et pare-feu qui le composent rendent ardue la conception d'une politique de filtrage efficace. Cependant, cette politique de filtrage est un composant clé de la défense en profondeur du SI.
- Elle limite les mouvements latéraux en cas de compromission.
- Elle réduit la surface d’attaque en n’autorisant que les flux strictement nécessaires.
- Elle facilite la détection des anomalies via la journalisation.
Pour aller plus loin, consultez notre article concernant la défense en profondeur.
La définition d'une bonne politique porte donc à la fois sur l'organisation et la cohérence de celle-ci lors de sa mise en place, mais aussi sur la possibilité de la maintenir dans le temps, en anticipant des évolutions certaines dans l'équipe informatique, les besoins opérationnels, les incidents techniques, les audits, les changements de technologies, etc.
Voyons à présent quelles sont ces bonnes pratiques.
Bien organiser sa politique de filtrage réseau
Le guide de l'ANSSI que nous résumons ici présente un modèle clair de conception d'une politique de filtrage réseau selon lequel tout ce qui n'est pas explicitement autorisé est interdit.
Cette approche "liste blanche" permet d'être sûr de ne rien oublier et de ne rien autoriser par omission. Elle est souvent la plus stricte, la plus efficace, mais aussi celle qui demande le plus de préparation avant sa mise en place.
Le tableau suivant nous est fourni afin de traiter les règles de filtrage par catégorie, mais aussi dans le bon ordre :
Dans ce modèle d'organisation, nous retrouvons notamment :
- Les règles de contrôle et protection du pare-feu lui-même (Section 1, 2 et 3). Il s'agit d'un composant du SI qui a, lui aussi, ses faiblesses, sa surface et ses vecteurs d'attaque. On parle notamment de ses services d'administration et de monitoring qui sont nécessairement actifs, mais aussi de l'envoi des logs, alertes et sauvegardes.
# Exemple de règle Section 1
Source : serveurs d'administration
Destination : Pare-feu
Protocole : SSH/HTTPS
Action : Autoriser
Journalisation : Oui
# Exemple de règle Section 2
Source : Pare-feu
Destination : Serveur de centralisation des logs
Protocole : TCP/514
Action : Autoriser
Journalisation : Oui
# Exemple de règle Section 3
Source : Any
Destination : Any
Service : Any
Action : Interdire
Journalisation : Oui
Notez que toutes les règles qui concernent le pare-feu lui-même sont journalisées, cela parce qu'il s'agit d'un composant critique et central du SI. Vous constaterez dans le guide que ce n'est pas le cas des règles d'autorisation des flux métiers, ce qui serait bien trop verbeux.
- Les règles d'autorisation des flux métiers et des services (Section 4) : c'est ici que les composants de la "liste blanche" seront définis. Elle contient donc tous les flux métier identifiés comme nécessaires au bon fonctionnement du SI, d'un système, service ou composant.
# Exemple de règle Section 4
Source : Zone Utilisateur
Destination : Active Directory
Service : Kerberos (TPC/88)
Action : Autoriser
Journalisation : Non
Même s'il s'agit d'une règle d'autorisation, il convient toujours de contrôler et de limiter au strict minimum la source, la destination et le type (protocole, ports) des flux. Ainsi, il sera très rare d'utiliser la directive "Any" sur des règles d'autorisation qui ciblent un ensemble d'objets sans restrictions.
- Les règles antiparasite (Section 5) qui visent à bloquer les flux non autorisés dont la trace (log) n'est volontairement pas conservée. L’objectif est de maintenir des journaux exploitables, de supprimer le bruit, sans perdre la visibilité sur les incidents de sécurité.
# Exemple de règle Section 5
Source : Réseau de test
Destination : 255.255.255.255
Service : UDP/137, UDP/138
Action : Interdire
Journalisation : Non
- La règle d'interdiction finale : à ne surtout pas oublier, c'est elle qui indique que tout ce qui n'a pas été explicitement autorisé par les règles précédentes est forcément bloqué, et journalisé.
# Exemple de règle Section 6
Source : Any
Destination : Any
Service : Any
Action : Interdire
Journalisation : Oui
