---
id: collect-261001-general-networking/general-networking/fr-review-ultimate-edge-eatons-infrastructure-for-modern-it-deployments-1734aa20-2
title: "fr-review-ultimate-edge-eatons-infrastructure-for-modern-it-deployments-1734aa20"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution", "ethernet"]
source: docs/RAG/collect-261001-general-networking/fr-review-ultimate-edge-eatons-infrastructure-for-modern-it-deployments-1734aa20.md
source_anchor: ""
source_lines: [26, 50]
sha256: e591e8ff71b310c07a052f5a33c248c8d04c70e49a01c31d962932c8fd268fab
---

# fr-review-ultimate-edge-eatons-infrastructure-for-modern-it-deployments-1734aa20

| Sortie 120 V, ensembles réseau 50/60 Hz – Contient une carte réseau Gigabit préinstallée (Network-M2) |  |  |  | 
| 5PX1000RTNG2 | 1000/1000 | 5-15P, 10 pieds | (8) 5-15R | 
| 5PX1500RTNG2 | 1440/1440 | 5-15P, 10 pieds | (8) 5-15R | 
| 5PX2000RTNG2 | 1950/1950 | 5-20P, 10 pieds. | (6) 5-20R, (1) L5-20R | 
| 5PX2000RT3UNG2 | 1950/1950 | 5-20P, 10 pieds. | (6) 5-20R, (1) L5-20R | 
| 5PX3000RTNG2 | 3000/3000 | L5-30P, 10 pieds | (6) 5-20R, (1) L5-20R | 
| 5PX3000RT3UNG2 | 3000/3000 | L5-30P, 10 pieds | (6) 5-20R, (1) L5-20R | 
| Modèles globaux à sortie 208 V/230 V, 50/60 Hz |  |  |  | 
| 5PX1500HRTG2 | 1500/1500 | C14 | (8) C13 | 
| 5PX2200HRTG2 | 2200/2200 | C20 / L6-20P1 | (8) C13, (2) C19 | 
| 5PX3000HRTG2 | 3000/3000 | C20 / L6-20P1 | (8) C13, (2) C19 | 
| 5PX3000HRTNG2 | 3000/3000 | C20 / L6-20P1 | (8) C13, (2) C19 | 
Le PDU – Eaton EMAT09-10
Pour la distribution électrique, nous avons opté pour le PDU rack administrable EMAT09-10 . Ce PDU 1U de 12 A est équipé de 8 prises 5-20R à l'arrière, permettant d'étendre les sorties de votre onduleur. Son entrée d'alimentation est un connecteur L5-20P ; toutefois, si vous ne disposez que de sorties 5-20P standard, Eaton fournit un adaptateur 5-20P vers L5-20R. Pour des fonctionnalités similaires dans un format vertical 0U, Eaton propose également le modèle EMA113-10.
L'EMAT09-10 étant géré et pas seulement mesuré, cela signifie que les administrateurs peuvent visualiser la consommation électrique de chaque port et peuvent également activer et désactiver les ports individuellement. Ces fonctionnalités peuvent ne pas sembler attrayantes à tout moment, mais peuvent être extrêmement utiles dans certains scénarios. Ces ports peuvent être utilisés pour redémarrer facilement les appareils à distance pour les sites qui ne disposent pas de personnel informatique sur place. Ci-dessous, nous avons une image de l'écran de prise où chaque prise individuelle peut être surveillée pour l'alimentation et l'état, ainsi que mise sous tension et hors tension.
Pour les appareils tels que les serveurs qui peuvent avoir 2 alimentations ou plus, vous pouvez les ajouter en tant qu'appareil dans l'interface Web de la PDU, afin que chaque appareil puisse être surveillé pour la consommation électrique combinée des prises liées. La fonction des appareils peut également vous permettre d'allumer et d'éteindre simultanément le groupe de prises pour redémarrer ou éteindre un appareil, au lieu de parcourir les ports individuellement.
Grâce aux ports Ethernet intégrés, ces PDU peuvent également être connectées en série pour partager une seule adresse IP pour un maximum de 8 PDU. Lors de la connexion en série de PDU, toutes les unités liées peuvent être gérées sous la même interface. La fonctionnalité des appareils prend également en charge la répartition des appareils sur plusieurs PDU.
Pour des informations plus détaillées sur les PDU gérés d'Eaton, consultez notre test du PDU universel G3. Ce test porte sur un PDU vertical 42U, mais en ce qui concerne la connectivité réseau, l'interface web et la surveillance environnementale, l'expérience est quasiment identique pour les PDU gérés G3 et G4.
Conclusion
La combinaison du SR25UB, du 5PX2000RTNG2 et de l'EMAT09-10 constitue une combinaison parfaite pour une variété d'environnements tout en s'adaptant aux contraintes d'alimentation et d'espace. La possibilité de préconfigurer un boîtier comme le SR25UB peut être essentielle pour les environnements périphériques pour un déploiement rapide sans avoir besoin d'un technicien sur site pendant de longues périodes. L'ajout d'un onduleur peut s'avérer essentiel pour protéger les équipements de votre environnement et garantir la disponibilité. Pour les environnements où l'alimentation électrique n'est pas idéale, un UPS peut protéger votre équipement coûteux contre les anomalies de puissance, le maintenir suffisamment longtemps pour qu'un générateur se déclenche en cas de panne ou lui signaler d'arrêter l'équipement en toute sécurité jusqu'à ce que l'alimentation soit rétablie. L'ajout de l'EPDU EMAT09-10 peut être extrêmement bénéfique pour redémarrer les appareils à distance afin de résoudre rapidement les problèmes sans avoir recours à une personne sur place.
Déployer et gérer des ressources informatiques à la périphérie n’est pas chose facile. Les sites de déploiement ne sont généralement pas uniformes, il y a peu ou pas de personnel informatique sur place et il est important de planifier les défis liés à la protection physique et à l'accès. Les produits Eaton contribuent grandement à atténuer bon nombre de ces problèmes, en offrant aux administrateurs informatiques un équipement fiable qui peut être géré à distance. Et grâce à la profondeur du catalogue Eaton, il est assez facile de copier nos sélections ultimes ou d'ajuster une version pour un cas d'utilisation spécifique.
Les liens vers le matériel que nous avons reçu se trouvent ci-dessous :
Onduleur Eaton 5PX2000RTNG2
Eaton / Tripp Lite SR25UB
Eaton EMAT09-10 PDU
