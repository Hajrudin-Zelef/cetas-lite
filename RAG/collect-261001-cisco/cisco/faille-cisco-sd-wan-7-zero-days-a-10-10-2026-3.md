---
id: collect-261001-cisco/cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026-3
title: "Commande observée par Mandiant (avril 2026)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "cyber", "exploit", "mai", "open source"]
source: docs/RAG/collect-261001-cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026.md
source_anchor: ""
source_lines: [101, 150]
sha256: 43524b88e8f8cc22dc57047e570b2bf750ef04e54215c389dac5691ab53021ce
---

# Commande observée par Mandiant (avril 2026)

| Versions corrigées pour CVE-2026-20182, d’après l’avis de sécurité Cisco (publié le 14 mai, mis à jour le 16 juin 2026). |  | 
|---|---|
| Branche logicielle | Première version corrigée (CVE-2026-20182) | 
|---|---|
| 20.9 | 20.9.9.1 | 
| 20.10 / 20.11 | 20.12.7.1 | 
| 20.12 | 20.12.5.4 / 20.12.6.2 / 20.12.7.1 | 
| 20.13 / 20.14 | 20.15.5.2 | 
| 20.15 | 20.15.4.4 / 20.15.5.2 | 
| 20.16 / 20.18 | 20.18.2.2 | 
| 26.1 | 26.1.1.1 | 

Au-delà du correctif, les recommandations défensives convergent : retirer immédiatement les plans de gestion de toute exposition Internet, restreindre l’accès à NETCONF (port 830) et SSH (port 22) à des réseaux d’administration dédiés, auditer les comptes et les clés SSH ajoutés récemment, et rechercher les traces d’effacement de journaux. Une plateforme de supervision comme un SIEM open source facilite la détection de ces anomalies à l’échelle d’un parc distribué.

## La France et l’Europe en première ligne

Contrairement à ce que la concentration américaine des instances exposées pourrait laisser croire, l’Europe n’est pas spectatrice. Dès le 15 mai 2026, le CERT Santé français (Cyberveille, rattaché au ministère de la Santé) a publié une alerte sur CVE-2026-20182 à destination des établissements de santé – un secteur particulièrement dépendant du SD-WAN pour relier hôpitaux, cliniques et plateaux techniques. L’ANSSI, via son CERT-FR, suit de près ces avis Cisco dans le cadre de sa mission de protection des opérateurs d’importance vitale.

Le contexte réglementaire européen amplifie l’enjeu. Avec la directive NIS 2, des dizaines de milliers d’entités deviennent responsables de la sécurité de leur chaîne d’approvisionnement numérique et de leurs équipements réseau. Une instance SD-WAN non corrigée n’est plus seulement un risque technique : c’est un potentiel manquement de conformité, passible de sanctions. Le Cyber Resilience Act renforce encore cette logique en imposant aux fabricants et aux utilisateurs des obligations de notification et de correction dans des délais contraints.

Le tout s’inscrit dans une année noire pour la cybersécurité française. La France figure parmi les cinq pays les plus ciblés au monde, avec une hausse d’environ 29 % des incidents de rançongiciel au premier semestre 2026, et une série de compromissions retentissantes déjà couvertes dans nos colonnes, de la cyberattaque contre l’ANTS aux statistiques faisant de la France l’un des pays les plus piratés. Dans ce climat, la sécurisation de la bordure réseau devient une priorité stratégique nationale.

## ArcaneDoor, Salt Typhoon : la bordure réseau, cible des États

La campagne UAT-8616 ne surgit pas de nulle part. Elle prolonge une tendance de fond documentée depuis plusieurs années : les équipements de bordure réseau – pare-feu, VPN, routeurs, contrôleurs SD-WAN – sont devenus la cible privilégiée des acteurs étatiques. Ces appareils présentent une combinaison idéale pour l’espionnage : ils sont exposés à Internet, ils voient passer tout le trafic, ils sont rarement surveillés aussi finement que les serveurs, et ils n’acceptent généralement pas d’agent de détection.

En 2024, la campagne ArcaneDoor avait visé les pare-feu Cisco ASA et Firepower via des zero-days, dans le cadre d’opérations attribuées à un acteur étatique. En 2024-2025, le groupe Salt Typhoon, lié à la Chine, avait compromis des opérateurs télécoms occidentaux en s’appuyant précisément sur des équipements réseau pour établir un espionnage de longue durée. La campagne SD-WAN de 2026 s’inscrit dans cette continuité, avec un raffinement supplémentaire : au lieu de viser la périphérie, elle vise le plan de contrôle – le point d’orchestration depuis lequel tout le réseau peut être manipulé.

Ce glissement de la périphérie vers le cœur du réseau change la nature du risque. Compromettre un pare-feu, c’est ouvrir une brèche ; compromettre un contrôleur SD-WAN, c’est prendre le contrôle de la carte routière elle-même. Pour un adversaire cherchant à observer discrètement les communications d’une administration ou d’un opérateur d’infrastructure critique, il n’existe pas de meilleur poste d’écoute.

## Impact sur le marché : le SD-WAN, colonne vertébrale des entreprises

Cisco est le leader historique du marché du SD-WAN, une technologie devenue incontournable pour connecter les sites distants à moindre coût sans dépendre des liaisons MPLS traditionnelles. Cette domination fait de son plan de gestion une cible à très fort effet de levier : un défaut dans le produit Cisco se répercute sur des milliers d’organisations à travers le monde, des banques aux distributeurs, en passant par la santé et les services technologiques – précisément les secteurs que Mandiant identifie comme dépendants du SD-WAN.

Pour les directions informatiques, l’épisode de 2026 impose une remise en question de la gouvernance des équipements réseau. Trop d’organisations traitent encore le SD-WAN comme une « boîte » que l’on installe et que l’on oublie, avec des cycles de correctifs lents et une visibilité limitée sur les versions déployées. Le retour arrière logiciel exploité par les attaquants démontre qu’un parc « à jour » sur le papier peut rester vulnérable si le mécanisme de mise à jour lui-même n’est pas verrouillé.

L’impact se mesurera aussi sur le marché de la cyberassurance. L’exposition d’un plan de gestion réseau sur Internet, ou l’ouverture de ports NETCONF et SSH, deviendront très probablement des critères d’exclusion ou de surprime lors de la souscription. La faille Cisco SD-WAN de 2026 servira de cas d’école pour justifier des exigences de durcissement plus strictes, à l’image de ce qui s’est produit après les grandes vagues de rançongiciels.

## Comparaison : Cisco face à Fortinet et aux autres

Cisco n’est évidemment pas le seul fournisseur de SD-WAN, et il serait injuste d’en faire un cas isolé. Ses concurrents directs – Fortinet, VMware (VeloCloud), Palo Alto Networks – ont tous connu leur lot de vulnérabilités critiques sur les équipements de bordure. Fortinet, en particulier, a été frappé par plusieurs failles majeures ces dernières années sur ses passerelles, ce qui a nourri un débat de fond sur la sécurité intrinsèque des architectures d’accès distant.

Le point commun à tous ces incidents n’est pas une négligence propre à un éditeur, mais une réalité structurelle : les équipements qui concentrent le contrôle du réseau sont, par nature, des cibles de très grande valeur. Ce qui distingue la campagne Cisco de 2026, c’est le nombre de failles exploitées sur une courte période, la présence confirmée d’un acteur étatique persistant, et surtout le fait que la vulnérabilité touche le plan de contrôle plutôt qu’une simple passerelle terminale. La leçon vaut pour toute l’industrie : la comparaison entre fournisseurs compte moins que la discipline opérationnelle appliquée à leur exploitation.

Pour les équipes qui souhaitent renforcer leur posture défensive au-delà du seul correctif, des outils libres comme Wazuh (SIEM open source) ou CrowdSec permettent de détecter les comportements anormaux – création de comptes, injections de clés, effacements de journaux – qui caractérisent ce type d’intrusion furtive.

## Cinq prévisions pour la sécurité de la bordure réseau

À partir de la trajectoire observée en 2026, plusieurs évolutions apparaissent probables pour le second semestre et au-delà :

