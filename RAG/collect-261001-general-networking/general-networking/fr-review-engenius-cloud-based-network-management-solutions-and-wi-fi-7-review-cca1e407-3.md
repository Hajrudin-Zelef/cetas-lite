---
id: collect-261001-general-networking/general-networking/fr-review-engenius-cloud-based-network-management-solutions-and-wi-fi-7-review-cca1e407-3
title: "fr-review-engenius-cloud-based-network-management-solutions-and-wi-fi-7-review-cca1e407"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-engenius-cloud-based-network-management-solutions-and-wi-fi-7-review-cca1e407.md
source_anchor: ""
source_lines: [33, 68]
sha256: 9bc4e860011b2bcf97e12343ea394e67c0ac6619390943f2fa4747a40f9691ef
---

# fr-review-engenius-cloud-based-network-management-solutions-and-wi-fi-7-review-cca1e407

Le nombre de combinaisons possibles est limité en pratique, car plus la variation d'amplitude et d'angle de phase est faible, plus il est difficile de distinguer les différentes valeurs parmi le bruit et les interférences lors d'une utilisation quotidienne. C'est aussi pourquoi les hautes fréquences ont une portée réduite. La création de ces différentes valeurs et la compensation des erreurs reposent sur des calculs mathématiques et une ingénierie complexes que je ne maîtrise pas suffisamment pour les expliquer. Toutefois, il suffit de dire que les améliorations apportées par le Wi-Fi 7 représentent un bond en avant considérable en termes de débit. Si vous souhaitez une représentation plus visuelle du fonctionnement de la modulation QAM, je vous recommande vivement cette excellente vidéo YouTube sur le sujet.
Opération multi-liens (MLO)
L'opération multi-liens, ou MLO, est un concept beaucoup plus simple à comprendre que QAM. Cela signifie simplement qu'un appareil compatible peut se connecter simultanément aux fréquences 6 GHz, 5 GHz et/ou 2.4 GHz pour un débit accru. Vous pouvez y penser un peu comme l’agrégation de liens pour le Wi-Fi.
16 flux MU-MIMO
MU-MIMO est une entrée multiple multi-utilisateur, une sortie multiple. Cela permet à un point d'accès de transmettre et de recevoir des données de plusieurs clients simultanément. Les anciennes itérations du Wi-Fi étaient à flux unique, ce qui signifiait que plusieurs utilisateurs sur un seul point d'accès devaient attendre dans une file d'attente d'un seul fichier pour envoyer ou recevoir leurs données. Le Wi-Fi 6 a activé 8 flux de MU-MIMO, et le Wi-Fi 7 a amélioré ce chiffre à 16. C'est également à cela que font référence les fabricants lorsque vous voyez des points d'accès et des routeurs annoncés comme 4×4:4 ou 8×8:8 ( émetteurs x récepteurs x flux spatiaux). L'EnGenius ECW536 est un AP 4×4:4. Il est rare de trouver du 8×8:8, et il est peu probable qu’un 16×16:16 soit à vendre de sitôt, voire jamais. Le facteur limitant ici est côté client, car intégrer autant d'émetteurs et de récepteurs dans un appareil sans fil est pour la plupart peu pratique, et la grande majorité des appareils mobiles n'ont aujourd'hui que des configurations 2 × 2: 2.
Plusieurs RU vers un seul client et perforation MRU
Le multi-RU est un autre concept assez simple. Une RU, ou Resource Unit, est une petite tranche de la fréquence Wi-Fi attribuée à un client pour transférer des données. Chaque canal Wi-Fi peut accueillir de nombreuses RU différentes, augmentant avec la largeur du canal. Le Wi-Fi 6 disposait de mécanismes assez intelligents pour optimiser le nombre d'utilisateurs par canal, mais ce n'était pas parfait. Le principal domaine d'amélioration apporté par le Wi-Fi 7 est de permettre d'attribuer plusieurs RU inutilisées à un seul client, car il est beaucoup plus courant d'avoir plus de RU que d'utilisateurs, en particulier sur les nouveaux canaux ultra-larges de 7 MHz du Wi-Fi 320.
Cette amélioration est combinée à une autre innovation, qui est le RU Puncturing. Bien que cela semble violent, en réalité, cela permet aux points d'accès Wi-Fi 7 de diviser les canaux plus grands en canaux plus petits pour optimiser l'utilisation des canaux et permettre aux clients d'accéder à plus de canaux en cas d'interférences. Autrement dit, si un seul canal de 80 MHz subit des interférences provenant d'une source extérieure, le Wi-Fi 7 permet de diviser ce canal en canaux de 20 MHz ou de 40 MHz, empêchant ainsi la perturbation d'un canal entier de 80 MHz et évitant la section du canal qui a été interrompue. ingérence.
Grâce à la combinaison de canaux ultra-larges de 320 MHz permettant d'augmenter le nombre d'unités de ressources, de la modulation 4096-QAM pour optimiser le nombre de bits par longueur d'onde et de l'agrégation de fréquences MLO, le Wi-Fi 7 offre une norme qui a permis une augmentation impressionnante du débit. De nombreuses autres fonctionnalités, comme les améliorations apportées à la sécurité et à la qualité de service (QoS), n'ont pas été abordées ici. Si vous souhaitez en savoir plus sur le Wi-Fi 7, je vous recommande vivement ce livre blanc très accessible (lien vers un PDF) de RUCKUS Networks.
Gestion du cloud
Assez parlé du Wi-Fi, revenons à l'examen en question et parlons de tout l'intérêt de ces appareils et du principal argument de vente d'EnGenius ; leur gestion cloud. Ils le présentent comme exploitant un modèle FaaS (Function as a Service), par opposition au SaaS (Software as a Service) et « sans serveur ». Bien sûr, à un moment donné, des machines physiques effectuent un travail, mais cela signifie que chaque fonction du système de gestion est indépendante, l'idée étant qu'il est hautement évolutif, plus résilient aux points de défaillance uniques et moins coûteux.
Comme mentionné au début de cette revue, EnGenius Cloud est une solution de gestion sans licence intégrée à leur matériel, avec la possibilité d'un niveau Pro payant avec des fonctionnalités supplémentaires. Au moment de la rédaction, la licence Pro coûte 50 $ par an et par appareil pour les commutateurs et les points d'accès et 100 $ par an par appareil pour les passerelles.
| Fonctionnalité | Sans licence | Cloud Pro | 
| Historique des statistiques | 3 jours | 30 jours | 
| Comptes administrateur | 10 | Illimité | 
| Service de bons d'achat (alias laissez-passer invités) | Entrées 100 | Entrées 10,000 | 
| Rapports planifiés | 3 jours | 30 jours | 
| Réseaux par organisation | 50 | 500 | 
| Notifications d'alerte | Appareil activé/hors ligne | Notifications détaillées | 
| Exportation de la liste des clients | N/D | Oui | 
| Clonage de réseau | N/D | Oui | 
| Sauvegarde et restauration du réseau | N/D | Oui | 
| VPN automatique | Oui | Oui | 
| Traversée automatique du NAT VPN | N/D | Oui | 
| Outils de diagnostic en direct | Fonction Plug & Play | Avancé | 
| Carte thermique Wi-Fi | Oui | Oui | 
| Liste des clients en direct (Wi-Fi) | N/D | Oui | 
| VLAN dynamique | N/D | Oui | 
| Regroupement de VLAN | N/D | Oui | 
| Vue de la topologie du réseau | Fourni par EnGenius | EnGenius et assistance tierce | 
| Statistiques des ports de commutation | Oui | Oui | 
| Capture de paquets | N/D | Oui | 
(Pour la liste complète, consultez le site web d'EnGenius )
EnGenius Cloud Pro offre définitivement de nombreuses fonctionnalités, certaines sont grandement facilitées par la gestion centralisée du cloud par rapport aux réseaux traditionnels avec des appareils indépendants. La version sans licence offre beaucoup de choses, et même si elle manque de confort et d'avantages, comme les statistiques à long terme, vous pouvez déployer un réseau robuste et ne pas avoir l'impression de manquer quelque chose d'essentiel au quotidien. -les opérations quotidiennes. Il est également important de mentionner que vous pouvez acheter leur licence Pro uniquement pour certains appareils, ce qui est un modèle que j'aime beaucoup. Si vous n'avez besoin que des fonctionnalités avancées pour les points d'accès, vous n'avez pas besoin de payer pour les licences Pro sur les commutateurs ou les passerelles.
Configuration du réseau
J'avais plusieurs appareils à déployer, probablement pas très différents, par exemple, d'un petit café doté d'une connexion Wi-Fi publique, de systèmes PoS connectés à Internet et de caméras de sécurité alimentées par PoE. Une fois que vous avez créé un compte EnGenius Cloud et téléchargé leur application (disponible pour iOS ou Android), l'ajout d'appareils est aussi simple que de scanner le code QR sur l'appareil. C'est vraiment aussi simple que cela. Scannez, appuyez sur « S'inscrire » et vous avez terminé. Il m'a fallu plus de temps pour monter les commutateurs dans le rack que pour les ajouter au réseau.
