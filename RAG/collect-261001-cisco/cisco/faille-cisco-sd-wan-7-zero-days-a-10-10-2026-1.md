---
id: collect-261001-cisco/cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026-1
title: "Commande observée par Mandiant (avril 2026)"
domain: cisco
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["attribution", "exploit", "incident", "mai"]
source: docs/RAG/collect-261001-cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026.md
source_anchor: ""
source_lines: [1, 44]
sha256: 3abc56a8e86c629e911c3befc5b906e317c67f32ca5a6a6aa5250d892e53542f
---

# Commande observée par Mandiant (avril 2026)

Mise à jour du 07 juillet 2026 – En six mois, l’infrastructure **Cisco Catalyst SD-WAN** est devenue l’un des terrains de chasse les plus actifs du cyberespionnage mondial. Depuis février 2026, sept vulnérabilités distinctes ont été exploitées en conditions réelles, dont deux **contournements d’authentification notés 10 sur 10** sur l’échelle CVSS – le score maximal. Derrière une partie de ces intrusions se cache UAT-8616, un acteur qualifié de « hautement sophistiqué » par les chercheurs de Cisco Talos, dont l’infrastructure recoupe des réseaux de relais opérationnels typiques des opérations soutenues par un État. Pour les entreprises françaises et européennes qui ont bâti l’épine dorsale de leur réseau sur cette technologie, la **faille Cisco SD-WAN** n’est plus une hypothèse : c’est un incident en cours.

Cet article détaille la chronologie complète de la campagne, la mécanique technique des deux zero-days critiques, le profil de l’acteur UAT-8616, l’ampleur de la surface d’attaque exposée sur Internet, et surtout ce que cela implique pour les opérateurs d’infrastructures critiques soumis à la directive NIS 2. Les faits, les scores CVSS et les versions correctives proviennent des avis officiels de Cisco, du catalogue KEV de la CISA, et des analyses de Mandiant, Tenable, VulnCheck et Censys.

## Sept zero-days Cisco SD-WAN exploités en 2026 : le récapitulatif

Le 5 juin 2026, Cisco a confirmé l’exploitation d’une septième vulnérabilité de sa gamme SD-WAN depuis le début de l’année, comme l’a rapporté SecurityWeek. Ce rythme est inhabituel, même pour un produit aussi répandu. La **faille Cisco SD-WAN** ne désigne donc pas un événement isolé, mais une série coordonnée de découvertes et d’exploitations qui, mises bout à bout, permettent à un attaquant de passer de zéro accès au contrôle total du plan de gestion réseau.

Le tableau ci-dessous récapitule les sept vulnérabilités concernées, avec leur score CVSS, leur nature et le produit affecté. Deux d’entre elles – CVE-2026-20127 et CVE-2026-20182 – atteignent la note maximale de 10,0, réservée aux failles exploitables à distance, sans authentification et sans interaction de la victime.

| Les sept vulnérabilités Cisco Catalyst SD-WAN exploitées en 2026. Sources : avis Cisco, catalogue KEV de la CISA, Tenable, Mandiant. |  |  |  |  | 
|---|---|---|---|---|
| Identifiant CVE | CVSS | Type | Produit affecté | Statut CISA KEV | 
|---|---|---|---|---|
| CVE-2026-20182 | **10,0** | Contournement d’authentification | Controller & Manager | 14 mai 2026 | 
| CVE-2026-20127 | **10,0** | Contournement d’authentification | Controller & Manager | 25 févr. 2026 | 
| CVE-2022-20775 | 7,8 | Élévation de privilèges (root) | vBond, vEdge, vSmart, vManage | 25 févr. 2026 | 
| CVE-2026-20133 | 7,5 | Divulgation d’informations | Manager | 20 avr. 2026 | 
| CVE-2026-20128 | 7,5 | Accès aux identifiants | Manager | 20 avr. 2026 | 
| CVE-2026-20122 | 5,4 | Écrasement de fichier arbitraire | Manager | 20 avr. 2026 | 
| CVE-2026-20245 | Root (auth. locale) | Exécution de commandes root | Manager (CLI) | Juin 2026 | 

À noter : CVE-2022-20775 n’est pas une vulnérabilité de 2026. Découverte en 2022, elle a été « réactivée » dans les chaînes d’exploitation observées cette année, ce qui illustre un principe fondamental de la sécurité de la bordure réseau : une faille corrigée mais non déployée reste une porte ouverte. Le score CVSS de CVE-2026-20245 n’a pas été communiqué publiquement par Cisco à la date de cet article ; nous la décrivons donc par son impact (exécution de commandes en tant que root) plutôt que par un chiffre invérifiable.

## CVE-2026-20182 et CVE-2026-20127 : deux contournements notés 10/10

Les deux vulnérabilités les plus graves partagent la même racine : un défaut dans le mécanisme d’authentification de peering du Cisco Catalyst SD-WAN Controller et du SD-WAN Manager. Concrètement, le processus de poignée de main des connexions de contrôle (*control connection handshaking*) ne valide pas correctement l’identité des pairs qui tentent de rejoindre le tissu réseau. Un attaquant distant non authentifié peut donc envoyer des requêtes forgées pour se faire passer pour un composant légitime.

Selon l’avis de sécurité officiel de Cisco, CVE-2026-20182 est notée CVSS 10,0 (critique) et classée CWE-287 (« Improper Authentication »). En cas de succès, l’attaquant obtient une session en tant que compte interne hautement privilégié – non root, mais doté d’un accès à NETCONF. Or NETCONF est le protocole qui pilote la configuration de l’ensemble du fabric SD-WAN. Autrement dit, une seule requête bien formée peut suffire à reconfigurer le routage, injecter des règles, ou détourner le trafic de dizaines de sites distants.

Cisco PSIRT, l’équipe d’intervention de l’éditeur, a reconnu avoir eu connaissance d’une exploitation « limitée » de CVE-2026-20182 dès mai 2026. Les actions observées incluent l’ajout de clés SSH, la modification de configurations liées à NETCONF et des tentatives d’élévation vers root. La société de sécurité Rapid7 a confirmé indépendamment que le correctif était disponible et que la faille était activement exploitée au moment de sa divulgation, le 14 mai 2026.

La première des deux failles à 10/10, CVE-2026-20127, avait été rendue publique près de trois mois plus tôt, le 25 février 2026. Elle avait été immédiatement ajoutée au catalogue des vulnérabilités activement exploitées (KEV) de la CISA. Le fait que deux failles de gravité maximale, quasiment identiques dans leur mécanique, aient été découvertes à si peu d’intervalle suggère une classe entière de défauts d’authentification dans l’architecture de peering – et donc, probablement, d’autres découvertes à venir.

## UAT-8616 : le groupe étatique derrière les intrusions

La signature la plus préoccupante de cette campagne porte un nom : UAT-8616. D’après la synthèse publiée par Tenable, ce groupe exploite des infrastructures Cisco SD-WAN depuis au moins 2023 – soit bien avant que les failles de 2026 ne soient rendues publiques. Il cible spécifiquement les secteurs d’infrastructures critiques, et son infrastructure de commande recoupe des réseaux de relais opérationnels (Operational Relay Box, ou ORB).

Ce détail des ORB est déterminant pour l’analyse. Les relais opérationnels – souvent constitués d’équipements compromis, de routeurs domestiques piratés ou de VPS anonymes – sont la marque de fabrique des acteurs étatiques cherchant à masquer l’origine géographique de leurs opérations. Ce sont eux qui ont rendu l’attribution du cyberespionnage si difficile ces dernières années. Le chevauchement d’UAT-8616 avec ce type d’infrastructure oriente fortement l’analyse vers un acteur aligné sur un État, motivé par l’espionnage à long terme plutôt que par le gain financier immédiat.

Le mode opératoire post-compromission confirme cette lecture. Une fois l’accès obtenu, UAT-8616 ne déploie ni rançongiciel ni mineur de cryptomonnaie. Le groupe injecte des clés SSH pour garantir sa persistance, manipule NETCONF pour reconfigurer le réseau, crée des comptes malveillants, puis efface méthodiquement les journaux pour couvrir ses traces. Cette discrétion – la volonté de rester présent sans être détecté – est l’inverse exact du comportement des groupes cybercriminels, qui cherchent à monétiser rapidement leur intrusion.

## Anatomie de l’attaque : du contournement au contrôle du fabric

