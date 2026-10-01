---
id: collect-261001-cisco/cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026-4
title: "Commande observée par Mandiant (avril 2026)"
domain: cisco
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["cyber", "mai"]
source: docs/RAG/collect-261001-cisco/faille-cisco-sd-wan-7-zero-days-a-10-10-2026.md
source_anchor: ""
source_lines: [151, 189]
sha256: 42bb89f474d28a6ae6c6c0987afcb042b9e0371aa9e5ac87112ad88127a9f3ca
---

# Commande observée par Mandiant (avril 2026)

1. **D’autres zero-days SD-WAN émergeront.** La découverte de deux contournements d’authentification quasi identiques suggère une classe de défauts encore incomplètement explorée ; la prolifération des PoC accélérera les découvertes.
2. **Les régulateurs européens durciront les exigences.** Sous l’impulsion de NIS 2 et du Cyber Resilience Act, l’exposition d’un plan de gestion réseau deviendra un manquement de conformité explicitement sanctionnable.
3. **Les acteurs étatiques délaisseront la périphérie pour le plan de contrôle.** Le succès d’UAT-8616 sur le « cerveau » du réseau incitera d’autres groupes ORB à viser les contrôleurs plutôt que les seuls pare-feu.
4. **Cisco accélérera son virage vers une gestion cloud et zero-trust.** La pression poussera à retirer les plans de gestion d’Internet par défaut et à généraliser l’authentification forte des pairs.
5. **La cyberassurance intégrera l’exposition réseau comme critère.** L’ouverture de NETCONF ou SSH sur Internet deviendra un signal d’alerte à la souscription, avec surprimes ou exclusions à la clé.

### Couverture associée

## FAQ : la faille Cisco SD-WAN de 2026

### Quelles sont les failles Cisco SD-WAN les plus graves de 2026 ?

Les deux vulnérabilités les plus critiques sont CVE-2026-20127 et CVE-2026-20182, toutes deux notées 10,0 sur l’échelle CVSS. Il s’agit de contournements d’authentification affectant le Cisco Catalyst SD-WAN Controller et le Manager, exploitables à distance et sans authentification préalable. Elles permettent d’obtenir un accès hautement privilégié au plan de gestion du réseau.

### Qui est UAT-8616 ?

UAT-8616 est un acteur de menace suivi par Cisco Talos, qualifié de « hautement sophistiqué ». Il exploite des infrastructures Cisco SD-WAN depuis au moins 2023, cible les secteurs d’infrastructures critiques, et son infrastructure recoupe des réseaux de relais opérationnels (ORB), ce qui oriente l’analyse vers un acteur aligné sur un État et motivé par l’espionnage.

### Mon équipement Cisco SD-WAN est-il concerné ?

Les produits affectés sont principalement le Catalyst SD-WAN Manager (ex-vManage) et le Controller, ainsi que certains composants comme vBond, vEdge et vSmart pour la faille CVE-2022-20775. Consultez l’avis de sécurité officiel de Cisco pour vérifier votre branche logicielle et déployer la première version corrigée correspondante. Toute version antérieure à la 20.9 nécessite une migration.

### La faille Cisco SD-WAN a-t-elle touché la France ou l’Europe ?

Si la majorité des instances exposées se concentrent aux États-Unis, l’Europe héberge une part significative des équipements vulnérables. En France, le CERT Santé a publié une alerte dès le 15 mai 2026, et l’ANSSI suit les avis Cisco dans le cadre de la protection des opérateurs d’importance vitale et de la directive NIS 2.

### Que faire pour se protéger immédiatement ?

Trois actions prioritaires : appliquer sans délai les correctifs Cisco correspondant à votre version ; retirer le plan de gestion (Manager/Controller) de toute exposition sur Internet et restreindre l’accès à NETCONF (port 830) et SSH (port 22) à un réseau d’administration dédié ; auditer les comptes récents, les clés SSH ajoutées et les journaux effacés pour détecter une éventuelle compromission déjà en cours.

### Pourquoi la CISA a-t-elle émis une directive d’urgence ?

La CISA a publié la directive d’urgence 26-03 le 14 mai 2026 pour CVE-2026-20182, imposant une remédiation rapide aux agences fédérales américaines. Ces directives sont exceptionnelles et réservées aux menaces jugées « inacceptables ». Elles témoignent du niveau de gravité et de l’exploitation active de la faille au moment de sa divulgation.

### Combien d’équipements Cisco SD-WAN sont exposés sur Internet ?

Les estimations varient selon les moteurs de recherche d’exposition : entre 450 et 550 instances de SD-WAN Manager selon Censys et Shodan, environ 275 pour ZoomEye, plus de 1 000 pour FOFA, et jusqu’à près de 2 000 équipements Catalyst SD-WAN au total selon un relevé Censys de mai 2026. Chaque instance exposée représente une organisation entière potentiellement à risque.

*Article publié le 07 juillet 2026. Les scores CVSS, dates et versions correctives proviennent des avis officiels de Cisco, du catalogue KEV de la CISA et des analyses publiées par Mandiant, Tenable, VulnCheck et Censys. Les informations reflètent l’état des connaissances à la date de publication et sont susceptibles d’évoluer à mesure que l’enquête se poursuit.*
