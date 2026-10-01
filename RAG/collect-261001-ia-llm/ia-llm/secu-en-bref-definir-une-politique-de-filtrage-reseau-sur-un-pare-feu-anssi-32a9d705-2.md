---
id: collect-261001-ia-llm/ia-llm/secu-en-bref-definir-une-politique-de-filtrage-reseau-sur-un-pare-feu-anssi-32a9d705-2
title: "Exemple de règle Section 1"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-ia-llm/secu-en-bref-definir-une-politique-de-filtrage-reseau-sur-un-pare-feu-anssi-32a9d705.md
source_anchor: ""
source_lines: [73, 117]
sha256: 05b4d5143cfeb9692f7f4631e41d3f989262cc6abf3afa324d8b436b6a79b84e
---

# Exemple de règle Section 1

Dans ce contexte, la journalisation des flux bloqués présente un intérêt en termes de sécurité. Les flux bloqués résultent forcément d'une activité non prévue, ou non nécessaire. Il est donc généralement utile de regarder fréquemment ces journaux afin de constater soit un incident, soit un composant qui ne devrait pas échanger depuis telle source vers telle destination et sur tel protocole (oubli, ou composant inutile ?).
Il est intéressant de relever cette remarque du guide concernant ce flux explicite de blocage et de journalisation :
Certaines solutions techniques appliquent automatiquement une règle d’interdiction à la fin de la politique de filtrage, mais celle-ci n’est généralement pas journalisée ou n’apparaît pas explicitement à la fin de la politique ; c’est la raison pour laquelle une règle finale explicite est ajoutée dans tous les cas. - Source - ANSSI
Pensez donc bien à ne pas utiliser la règle implicite de blocage, mais à déclarer votre propre règle d’interdiction finale, en activant également sa journalisation.
Comment mettre en forme sa politique de filtrage
Au-delà des règles de filtrage elles-mêmes, il est important d'avoir des règles de gestion commune de celles-ci. L'objectif ? S'assurer que n'importe quel membre d'une équipe puisse intervenir sur n'importe quel pare-feu, et qu'il puisse comprendre leur organisation, leur définition et leur raison d'être.
Ainsi, le guide de bonnes pratiques de l'ANSSI souligne l'importance d'adopter des règles de nomenclature, mais aussi de coloration et de description de ces dernières.
Les quelques règles décrites permettront par exemple de faciliter la recherche, la manipulation et la revue (audit) des règles.
Concernant la coloration, lorsque le pare-feu le permet, on peut par exemple adopter une convention telle que :
- Rouge : concerne le réseau externe
- Orange : concerne la DMZ publique
- Jaune : concerne la DMZ privée
- Vert : réseau interne
Il ne s'agit bien sûr que d'un exemple, la réalité et la conception de votre réseau pouvant vous amener à faire des choix différents.
La nomenclature des règles, c'est-à-dire la structure qui régit le nom qui leur est donné, a aussi une grande importance. Elle permet notamment d'identifier rapidement les caractéristiques techniques ou fonctionnelles de la règle (composants ou protocoles concernés). L'important étant bien sûr d'avoir des règles de nommage communes au sein d'une même équipe.
Enfin, le champ commentaire de chaque règle ne doit dans l'absolu jamais être vide. On peut par exemple y indiquer la date d'implémentation/mise à jour de la règle, la raison de sa mise en place, son éventuelle date d'expiration, le nom de la personne l'ayant mise en place, etc.
Il faut absolument éviter les règles de filtrage dont l’utilité n’est plus claire, mais qui sont conservées par crainte de provoquer un dysfonctionnement quelque part dans le SI. C'est généralement ce qui mène à l'accumulation de règles d’autorisation, ce qui casse complètement la segmentation du SI.
L'idéal est donc de mettre en place une convention commune de création, nommage et description des règles qui facilite leur compréhension par chaque personne susceptible d'intervenir sur vos pare-feu. Cela facilitera au passage la revue des règles par des tiers, par exemple lors d’un audit de sécurité, de l’intervention d’un prestataire sur vos pare-feu, etc.
Les bonnes pratiques et la documentation
Pour terminer, ce guide contient un certain nombre de bonnes pratiques globales qu'il vaut mieux avoir en tête lorsque l'on met en place une politique de filtrage.
Par exemple, la désactivation des flux implicites : il s'agit de "groupements" de flux préconçus qui embarquent souvent plus de protocoles ou de services que réellement nécessaires. Lorsqu'ils sont utilisés, on se retrouve alors à ajouter dans notre liste blanche des protocoles non utilisés, ce qui pourrait ouvrir des chemins de compromission plus tard.
Également, un chapitre entier est dédié à la documentation et au maintien dans le temps de notre politique de filtrage. Cette documentation vise notamment à formaliser les choix de nomenclature, d'organisation et de structure des règles décrites précédemment. Bien que des bonnes pratiques soient proposées, il reste nécessaire de les adapter à chaque contexte, et donc d'avoir une documentation.
Deux éléments sont mentionnés dans le guide de l'ANSSI concernant cette documentation : sa fréquence de revue et sa validation pratique (tests, outils d'analyse réseau).
Pour illustrer l'ensemble des bonnes pratiques de ce document, l'ANSSI nous propose la capture d'écran suivante :
On identifie notamment l'organisation par catégories (les 6 sections), la description des règles avec une date, etc.
Conclusion
Je vous laisse le soin d'aller consulter l'entièreté du document pour avoir des informations plus complètes sur ces bonnes pratiques. Il s’agit d’un guide pratique très intéressant et qui sera à coup sûr utile dans toutes les équipes informatiques et entreprises.
Pour aller plus loin et notamment procéder au nettoyage d'une politique de pare-feu existante, ce qui est un exercice encore plus périlleux, sachez qu’il existe aussi un guide de l’ANSSI pour vous accompagner dans cet exercice :
N’hésitez pas à utiliser les commentaires pour nous faire un retour sur l’utilisation et l’application de ces guides de l’ANSSI dans votre contexte, ou à nous indiquer si vous souhaitez plus de résumés de ce type sur IT-Connect !
Bonjour,
Lien au début de l’article down
Bonjour Tib,
Merci pour l’info, je vais corriger ça. L’ANSSI vient de déplacer ses guides sur un nouveau sous-domaine et visiblement ce lien n’a pas suivi.
Bonjour,
L’article de l’ANSSI date de 2016, n’y a-t-il pas eu une évolution sur les recommandations en 10 ans ? Surtout avec l’évolution des techniques d’attaques.
Mais cela reste un article très intéressant et un sujet au cœur de l’actualité !
Merci
Bonjour
L’exemple de la section 3 empèche l’éxécution de la section 4 et 5. Si source et destination sont à ANY, tous les flux seront bloqués quoi qu’on mette ensuite dans les autres sections. Si on se réfère à l’exemple Figure1 en fin d’article, la destination serait plutot: Destination : Pare-feu voir Pare-feu-any pour prendre en compte toutes les interfaces du pare-feu.
# Exemple de règle Section 3
Source : Any
Destination : Any
Service : Any
Action : Interdire
Journalisation : Oui
