---
id: collect-261001-general-networking/general-networking/attaque-supply-chain-axios-npm-2026-rat-nord-coreen-1
title: "attaque-supply-chain-axios-npm-2026-rat-nord-coreen"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: []
keywords: ["agent", "attribution", "aws", "distribution", "open source"]
source: docs/RAG/collect-261001-general-networking/attaque-supply-chain-axios-npm-2026-rat-nord-coreen.md
source_anchor: ""
source_lines: [1, 32]
sha256: f0fedc9933906ebb045b418d2dbabebe5d333a45f05e9aaab13cd18b016ca0a8
---

# attaque-supply-chain-axios-npm-2026-rat-nord-coreen

Le 31 mars 2026, l’écosystème JavaScript a subi l’une des attaques de supply chain les plus sophistiquées de son histoire. Deux versions malveillantes de la bibliothèque Axios – le client HTTP le plus utilisé au monde avec plus de 100 millions de téléchargements hebdomadaires sur npm – ont été publiées après le piratage du compte du mainteneur principal. En seulement 3 heures d’exposition, un cheval de Troie d’accès à distance (RAT) multiplateforme a été déployé sur des centaines de systèmes Windows, macOS et Linux, compromettant potentiellement des pipelines CI/CD, des clés cloud et des tokens d’authentification dans des milliers d’entreprises à travers le monde. Voici l’analyse complète de l’attaque supply chain Axios npm qui ébranle la confiance dans l’open source.

## Chronologie de l’attaque supply chain Axios npm : 3 heures qui ont ébranlé JavaScript

L’attaque supply chain Axios npm a été orchestrée avec une précision chirurgicale. Selon l’analyse de Palo Alto Networks Unit 42 et de Huntress, la séquence des événements révèle une opération minutieusement préparée par un acteur étatique nord-coréen. Le 30 mars 2026 à 05h57 UTC, un paquet « leurre » propre a été publié sur npm pour tester les mécanismes de détection. À 16h03 UTC le même jour, le domaine de commande et contrôle (C2) **sfrclak[.]com** a été enregistré via Namecheap Inc, hébergé chez Hostwinds LLC (AS54290) aux États-Unis.

Le véritable assaut a commencé le 30 mars à 23h59 UTC avec la publication du paquet malveillant **[email protected]**, un typosquat de la bibliothèque légitime crypto-js, contenant un script postinstall déployant le RAT. Seulement 6 minutes plus tard, à 00h05 UTC le 31 mars, le scanner automatisé de Socket a détecté le malware. Mais il était déjà trop tard pour empêcher la suite : à 00h21 UTC, **[email protected]** a été publié et tagué « latest », rendant l’attaque active pour tous les projets utilisant des versions flottantes. À 00h23 UTC, soit 89 secondes après la publication, la première infection a été observée sur un endpoint macOS. À 01h00 UTC, **[email protected]** a été publié et tagué « legacy », ciblant ainsi les deux branches principales de la bibliothèque.

npm a finalement retiré les versions malveillantes entre 03h15 et 03h29 UTC, plaçant un verrou de sécurité à 03h25 UTC et remplaçant la dépendance malveillante par un stub à 04h26 UTC. La fenêtre d’exposition totale : environ **3 heures**, durant lesquelles tout projet exécutant `npm install` avec une dépendance vers `axios@^1.14.0` ou `axios@^0.30.0` téléchargeait automatiquement le RAT.

## Le mécanisme technique : comment plain-crypto-js a déployé un RAT multiplateforme

L’attaque supply chain Axios npm se distingue par sa sophistication technique. Les versions compromises **[email protected]** et **[email protected]** introduisaient une dépendance cachée vers **[email protected]**, qui ne correspondait à aucun tag de release officiel sur GitHub – un signal d’alerte que les équipes de sécurité de Snyk et ArmorCode ont rapidement identifié. Ce paquet malveillant contenait un script postinstall (`node setup.js`) qui déclenchait le déploiement d’un cheval de Troie d’accès à distance (RAT) entièrement fonctionnel.

Le RAT fonctionnait de manière identique sur Windows, macOS et Linux, avec des chemins de communication C2 spécifiques à chaque plateforme pour imiter le trafic légitime du registre npm : `packages.npm[.]org/product0` pour macOS, `/product1` pour Windows et `/product2` pour Linux. Une fois installé, le malware effectuait immédiatement une reconnaissance du système : énumération des répertoires utilisateur, des racines de volumes, des processus en cours d’exécution, puis transmettait ces données au serveur C2 à l’adresse **142.11.206.73** sur le port 8000.

Un détail révélateur et potentiellement utile pour la détection : le RAT utilisait un **user-agent Internet Explorer 8 sous Windows XP** codé en dur de manière identique sur les trois plateformes. Cette anomalie – un user-agent IE8/XP sur un Mac ou un serveur Linux – est triviale à détecter sur les réseaux modernes. Le chemin de campagne `/6202033` était également significatif : inversé, il donne « 3-30-2026 », soit la date de préparation de l’attaque. Le RAT établissait une boucle de beacon toutes les 60 secondes, prêt à accepter des commandes distantes, exécuter des scripts arbitraires et injecter des binaires en mémoire. Sur Windows spécifiquement, il établissait une **persistance au redémarrage**, relançant le téléchargement du payload à chaque connexion utilisateur.

## Attribution : Sapphire Sleet et UNC1069, les acteurs nord-coréens derrière l’attaque

Microsoft Threat Intelligence a attribué l’attaque supply chain Axios npm au groupe **Sapphire Sleet**, un acteur étatique nord-coréen connu pour cibler les chaînes d’approvisionnement logicielles. Selon le rapport publié le 1er avril 2026, l’infrastructure C2 utilisée – hébergée chez Hostwinds – correspond aux méthodes opérationnelles habituelles de ce groupe. « Cette activité suit le schéma des récentes attaques supply chain de haut profil, où des adversaires empoisonnent des frameworks open source largement adoptés et leurs canaux de distribution pour obtenir un impact en aval étendu », a déclaré l’équipe Microsoft Threat Intelligence dans son analyse technique.

De son côté, le Google Threat Intelligence Group a attribué l’attaque à **UNC1069**, un acteur à motivation financière lié à la Corée du Nord et actif depuis 2018. L’infrastructure C2 montrait des connexions depuis un nœud AstrillVPN précédemment utilisé par UNC1069, avec une infrastructure adjacente sur le même ASN historiquement liée aux opérations de ce groupe. Cette double attribution – Sapphire Sleet par Microsoft, UNC1069 par Google – pointe vers le même appareil de cyberespionnage et de vol financier du régime de Pyongyang, qui utilise les attaques supply chain comme vecteur de financement de son programme d’armement.

L’attaque s’inscrit dans une campagne plus large baptisée **TeamPCP** par certains analystes, qui avait déjà compromis plusieurs projets open source populaires dans les semaines précédentes : Trivy (19 mars), KICS (23 mars), LiteLLM (24 mars) et Telnyx (27 mars 2026). Le ciblage d’Axios, avec ses 100 millions de téléchargements hebdomadaires, représentait le point culminant de cette campagne d’escalade progressive.

## L’impact mesuré : 135 endpoints compromis et des milliers de secrets exposés

Malgré une fenêtre d’exposition relativement courte de 3 heures, l’attaque supply chain Axios npm a eu un impact significatif. Huntress, la société de cybersécurité spécialisée dans la détection et la réponse, a observé **plus de 135 endpoints** contactant l’infrastructure C2 à travers les trois systèmes d’exploitation. Ce chiffre ne représente que les infections détectées par les capteurs de Huntress – le nombre réel d’installations malveillantes durant la fenêtre d’exposition reste inconnu.

Pour mettre ce chiffre en perspective, Axios cumule entre 70 et 100 millions de téléchargements hebdomadaires selon les sources. Même en supposant que seulement 0,1 % des installations se sont produites durant la fenêtre de 3 heures (qui couvrait la nuit en Europe et la soirée aux États-Unis), cela représenterait potentiellement **des dizaines de milliers de systèmes** ayant téléchargé et exécuté le payload malveillant. Chaque système compromis exposait potentiellement : des tokens npm, des clés AWS et cloud, des clés SSH, des secrets CI/CD, des variables d’environnement (.env), des credentials OAuth et des clés API stockées localement.

