---
id: collect-261001-general-networking/general-networking/fr-review-synology-rt6600ax-router-review-4297dcac-2
title: "fr-review-synology-rt6600ax-router-review-4297dcac"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "exploit"]
source: docs/RAG/collect-261001-general-networking/fr-review-synology-rt6600ax-router-review-4297dcac.md
source_anchor: ""
source_lines: [40, 49]
sha256: b1c5d4703db571ef487492daf54e2ec9dff484105e5869571886601e7fbfeb87
---

# fr-review-synology-rt6600ax-router-review-4297dcac

Accompagnant la dernière version de RSM, l'application mobile pour Andriod et IOS : routeur DS. Le routeur DS permet la configuration initiale des réseaux avec le RT6600ax et des réseaux maillés si nécessaire. Le routeur DS permet aux utilisateurs de créer des réseaux WiFi, de gérer la redirection de port et le contrôle du trafic. Les utilisateurs peuvent également configurer Safe Access et gérer les connexions VPN s'ils le souhaitent.
Performances du Synology RT660ax
Pour les tests du Synology RT6600ax, nous nous sommes concentrés sur le seul port 2.5 GbE sur le routeur qui peut être exploité à la fois pour une connexion WAN plus rapide (Internet multi-gig) ou un seul périphérique rapide sur le réseau (NAS). Nous avons opté pour l'angle NAS, en utilisant un modèle QNAP à 4 baies entièrement flash plus petit avec une interface 2.5 GbE (oui, nous savons… mais nous avions besoin d'un appareil 2.5 GbE). Nous avons ensuite utilisé les trois ports LAN 1GbE restants vers trois ordinateurs portables câblés pour mesurer les performances globales sur le routeur.
Le plan de test était assez simple. Ayez trois dossiers partagés sur le NAS avec une empreinte de 25 Go chacun et attribuez un dossier par ordinateur portable. Nous avons ensuite mesuré les vitesses de lecture et d'écriture séquentielles sur chaque système fonctionnant simultanément et avons combiné les chiffres pour obtenir un score final.
Sur le câble dans son ensemble, nous avons mesuré les performances de lecture sur le port 2.5 GbE à 157 Mo/s et les performances d'écriture à 285.6 Mo/s. Les performances d'écriture ont pu mieux utiliser la vitesse de 2.5 GbE dans ce cas par rapport à la vitesse de lecture, mais cela montre également une certaine faiblesse dans la configuration.
La meilleure configuration de port LAN serait de 4 ports 2.5 GbE avec 2 ports 2.5 GbE comme bon compromis. Le port unique de 2.5 GbE n'aide vraiment que si plusieurs utilisateurs utilisent un périphérique de stockage rapide ou prennent plus que probablement en charge une connexion WAN plus rapide.
Conclusion
Dans l'ensemble, le Synology RT6600ax est un solide routeur Wi-Fi tri-bande optimisé pour distribuer une connexion et des performances fiables tout en évitant les goulots d'étranglement. Le RT6600ax dispose également d'un support d'application solide, ce qui est quelque peu rare dans l'espace des routeurs. Cela offre aux utilisateurs un accès aux serveurs VPN, un package de téléchargement pour gérer les protocoles FTP, torrent et autres, et même un hôte multimédia intégré. Bon nombre de ces fonctionnalités chevauchent ce qui serait considéré comme une fonction NAS, mais elles peuvent être exécutées à bord du RT6600ax si nécessaire.
Les performances du Synology RT6600ax nous semblent un peu insuffisantes en ce qui concerne la connectivité Ethernet. Le RT6600ax comprend un seul port 2.5 GbE, qui peut être utilisé pour les besoins côté WAN ou LAN. Pour le WAN, cela vous permet de gérer le trafic à partir d'une connexion Internet multi-gig, ou du côté LAN quelque chose de plus rapide comme un NAS. L'inconvénient d'avoir un seul port signifie que si vous avez un poste de travail dans votre réseau domestique, vous ne pouvez pas avoir un chemin de données complètement rapide de l'ordinateur au NAS. Il aurait été bien ici de voir deux ports 2.5 GbE à ce sujet pour ouvrir des cas d'utilisation supplémentaires. Lors de nos tests de performances, nous avons mesuré un total de 157 Mo/s en lecture et 286 Mo/s à partir de trois ordinateurs portables câblés accédant à un seul partage de fichiers flash sur un NAS connecté.
Le Synology RT6600ax est un excellent routeur/AP/appareil maillé qui offre une large gamme de flexibilité avec une fonctionnalité multi-gig et un équilibrage de charge. Avec Synology Router Management, le RT6600ax permet aux administrateurs d'avoir un contrôle total sur n'importe quelle partie du réseau. Le RT6600ax est maintenant disponible pour 299.99 $
