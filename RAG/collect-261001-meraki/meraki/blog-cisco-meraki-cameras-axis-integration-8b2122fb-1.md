---
id: collect-261001-meraki/meraki/blog-cisco-meraki-cameras-axis-integration-8b2122fb-1
title: "blog-cisco-meraki-cameras-axis-integration-8b2122fb"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-meraki/blog-cisco-meraki-cameras-axis-integration-8b2122fb.md
source_anchor: ""
source_lines: [1, 42]
sha256: d7459810390e42bb0776b90072fb657332b25e307c7ba3d40e9e61a4ba9652e7
---

# blog-cisco-meraki-cameras-axis-integration-8b2122fb

Une caméra encore en bon état ne devrait pas devoir partir à la benne parce que l’entreprise change de console d’administration. L’ouverture de Meraki aux équipements Axis répond à ce problème : conserver une partie du parc vidéo tout en rapprochant son exploitation de celle du réseau.
Mais « visible dans Meraki » et « exploitable comme une caméra Meraki » ne signifient pas la même chose. L’intégration est réelle, mais son périmètre dépend de la licence, du modèle et des services Axis. Et un choix mérite d’être posé avant toute commande : garder son logiciel vidéo actuel, ou confier cette fonction à Meraki ?
Le billet Cisco du 16 septembre 2026 met cette convergence en avant. Ce n’est toutefois pas la date de première disponibilité : le guide d’intégration situe Essentials au 14 avril 2026 et Advantage au 23 juin. Voici ce qu’il faut examiner pour un projet, à partir des informations vérifiées le 18 septembre 2026.
Ce qu’il faut retenir
- Essentials ajoute une couche d’administration et de vidéo en direct. Le VMS local, c’est-à-dire le logiciel de gestion vidéo existant, peut rester en service.
- Advantage change le système d’exploitation vidéo du parc : Meraki devient le VMS exclusif. Le matériel peut être réemployé, mais pas nécessairement toute l’architecture qui l’entoure.
- 90 références sont répertoriées dans la matrice officielle, dont 87 pour les deux niveaux de licence. Les trois exceptions sont détaillées plus bas.
- Axis Cloud Connect reste indispensable à l’intégration. Une console commune ne signifie pas un fournisseur technique unique.
- Les prix, la disponibilité dans le pays concerné et la couverture fonctionnelle de chaque caméra doivent être validés pour le projet. La compatibilité d’un modèle n’est pas une garantie de parité avec toute la gamme MV.
Dashboard pour administrer, Vision pour exploiter les images
Le premier intérêt est organisationnel. L’équipe réseau et l’équipe sûreté peuvent travailler dans un environnement commun sans confondre leurs métiers. La première cherche pourquoi une caméra est hors ligne ; la seconde cherche ce qui s’est passé devant une porte.
Cisco positionne Meraki Dashboard comme l’interface de gestion des équipements et Meraki Vision comme celle de consultation et d’investigation vidéo. Le billet montre un mur vidéo mêlant des vues Axis et Cisco. L’accès à l’historique et aux exports relève du niveau Advantage, pas de la seule présence d’une caméra dans l’inventaire.
Ce qui est effectivement administrable
Le guide d’enrôlement Meraki détaille des actions allant au-delà d’un simple lien vers une autre console :
| Opération | Essentials | Advantage | 
|---|---|---|
| Vidéo en direct, état et connectivité | Oui | Oui | 
| Configuration et gestion du firmware | Oui | Oui | 
| Diagnostic : ping, redémarrage, traceroute, capture réseau | Oui | Oui | 
| Zoom, ouverture, mise au point ; orientation PTZ | Selon le matériel | Selon le matériel | 
| Historique, exports et enregistrement SD pilotés par Meraki | Non | Oui | 
| Stockage vidéo cloud de 30 jours | Non | Inclus | 
| Qualité/résolution vidéo, zones de mouvement et analyses | Non dans ce périmètre | Oui, selon les capacités prises en charge | 
| Conservation du VMS local en parallèle | Oui | Non : Meraki exclusif | 
La ligne sur l’enregistrement décrit ce que gère cette intégration, pas une interdiction pour le VMS existant d’enregistrer en mode Essentials.
Les analyses annoncées pour Advantage comprennent la recherche d’événements avec filtres personnes/véhicules, les alertes de mouvement et les cartes de chaleur. Pour l’historique, le guide demande une carte SD Axis ou l’activation du stockage cloud inclus. Vérifiez la capacité de stockage et les fonctions réellement disponibles sur chaque référence du pilote.
En pratique, il faut rédiger la recette à partir des gestes des utilisateurs : régler l’image d’une entrée, retrouver un événement, extraire une séquence, vérifier une panne. Un bouton présent dans une interface n’est utile que si l’opération aboutit sur les caméras du site, avec les bons droits.
Ne transposons pas non plus automatiquement les fonctions des caméras Cisco Meraki MV à un modèle Axis. Notre article sur le suivi inter-caméras dans Meraki Vision traite d’une fonction spécifique : la compatibilité Axis générale ne suffit pas à en garantir le bénéfice.
Essentials ou Advantage : conserver les caméras n’est pas conserver le VMS
Prenons un cas théorique : plusieurs bureaux disposent déjà de caméras Axis et d’un logiciel vidéo utilisé quotidiennement par les équipes. La DSI souhaite mieux superviser ces équipements depuis Meraki.
Avec Essentials, le projet porte d’abord sur l’exploitation technique. Il peut être pertinent de conserver les habitudes de consultation, les processus d’export et le stockage existants, tout en rapprochant le suivi des caméras de celui du réseau. Il faut néanmoins tester que les changements de configuration ou de firmware ne perturbent pas ces usages.
Avec Advantage, le projet devient une migration vidéo. L’exclusivité du VMS Meraki impose de vérifier ce qui est abandonné : intégrations métier, postes opérateurs, mécanismes d’alerte, procédures d’investigation et accès aux archives antérieures. Ne retirez pas un NVR simplement parce que le direct apparaît dans Vision : faites d’abord valider la continuité de service et le devenir des enregistrements existants.
Axis présente lui-même cette offre comme une approche hybride avec deux niveaux de licence. Le bon critère de choix n’est donc pas « quelle formule a le plus de fonctions ? », mais « quel fonctionnement voulons-nous conserver ou remplacer ? ».
Le coût doit être comparé sur la durée : abonnements, éventuelles cartes SD, capacité Internet, maintenance, formation et temps de migration. Réemployer les caméras peut éviter un remplacement matériel ; cela ne prouve pas à lui seul que le coût global baisse.
Axis Cloud Connect ne disparaît pas derrière Meraki
L’intégration relie les organisations Cisco et Axis par OAuth 2.0. Cisco décrit cette liaison cloud-to-cloud comme nécessaire à l’émission des commandes et à la fourniture des mises à jour. L’application embarquée sur la caméra s’appuie également sur l’écosystème ACAP d’Axis.
Autrement dit, Meraki fournit le point d’entrée commun, pas l’effacement de la plateforme Axis. Axis Cloud Connect demeure la plateforme de services permettant de connecter et de gérer les équipements avec des solutions partenaires.
Cette architecture appelle trois vérifications concrètes. Ce sont des points de qualification du projet, pas des pannes constatées sur cette solution :
- Identités et délégations. Qui possède les deux organisations ? Qui peut autoriser ou révoquer leur liaison ? Évitez que le départ d’un administrateur ou d’un prestataire rende les accès difficiles à reprendre.
- Continuité. Que reste-t-il accessible pendant une coupure Internet ? Que deviennent l’enregistrement, la consultation et les alertes ? Testez séparément le stockage local et le stockage cloud ; un écran d’inventaire ne répond pas à ces questions.
- Responsabilités. Qui prend en charge le matériel, le réseau, AXIS OS, l’intégration et le service vidéo ? Inscrivez le chemin d’escalade dans le dossier d’exploitation, avec un responsable de suivi côté client.
La FAQ prévoit une collaboration Cisco/Axis sur un même dossier de support, mais affiche des plages différentes : Cisco 24 h/24, 7 j/7 ; Axis 24 h/24, 5 j/7. Ce dispositif ne dispense pas de convenir de la prise en charge d’un incident vidéo critique hors heures ouvrées.
