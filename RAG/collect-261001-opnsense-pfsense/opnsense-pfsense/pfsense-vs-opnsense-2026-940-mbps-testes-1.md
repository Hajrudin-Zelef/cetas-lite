---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/pfsense-vs-opnsense-2026-940-mbps-testes-1
title: "pfsense-vs-opnsense-2026-940-mbps-testes"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["apache", "benchmarks", "intel", "open source"]
source: docs/RAG/collect-261001-opnsense-pfsense/pfsense-vs-opnsense-2026-940-mbps-testes.md
source_anchor: ""
source_lines: [1, 37]
sha256: 4059bfacd560629ebbad95fb4ca8d4bd96c3da466c1ff13219972de743db9206
---

# pfsense-vs-opnsense-2026-940-mbps-testes

Mis à jour le 15 août 2026 – Le duel entre **pfSense** et **OPNsense** est devenu en 2026 la question la plus posée par les administrateurs réseau et les RSSI européens cherchant un pare-feu open source robuste. Depuis la sortie d’**OPNsense 26.7 « Xenial Xenops »** le 15 juillet 2026, puis de son dernier correctif **26.7.3** le 27 août 2026 confirmant le socle FreeBSD 15.1 – désormais l’unique branche officiellement supportée par le projet, la précédente série 26.1 étant passée en fin de vie le 15 juillet 2026 selon eosl.date –, le débat a encore gagné en intensité – d’autant que Netgate a répliqué de son côté avec **pfSense Plus 26.07**, dérivée de la révision FreeBSD 24.6 et référencée par la documentation officielle le 13 août 2026. Sur du matériel identique (Intel i5-8500, 16 Go de RAM, carte Intel i350-T4), les deux plateformes atteignent **940 Mbps en routage 1 Gbit/s** – un écart inférieur à 0,3 %. Et pourtant, le choix entre les deux a des conséquences profondes : cycle de mises à jour **bi-hebdomadaire** chez OPNsense contre publications annuelles chez pfSense Community Edition, licence **BSD-2-Clause** intégralement open source contre un mélange Apache 2.0/propriétaire, interface **MVC Phalcon** moderne contre Bootstrap 3 hérité. Ce comparatif détaillé tranche le débat sur la base de bancs d’essai vérifiés, d’une analyse de licences, d’un tableau tarifaire complet et de cinq scénarios d’usage concrets pour 2026.

## pfSense vs OPNsense en bref : le verdict en 2026

Avant de plonger dans les benchmarks, voici la synthèse pour les lecteurs pressés. **OPNsense** s’impose en 2026 comme le choix par défaut pour la grande majorité des nouveaux déploiements : son cycle de mises à jour bi-hebdomadaire, son interface utilisateur moderne, ses 80+ plugins officiels et sa licence BSD-2-Clause intégrale collent parfaitement aux exigences européennes de souveraineté numérique et aux contraintes NIS 2 entrées en vigueur en 2024. **pfSense Plus**, dans sa version commerciale 24.x maintenue par Netgate, conserve néanmoins l’avantage sur des charges spécifiques : meilleure documentation technique, support des appliances Netgate (SG-1100 à SG-8300), et un écosystème de packages historique avec des références incontournables comme **pfBlockerNG**.

Côté performance brute, les deux plateformes font jeu égal sur du matériel identique. La différence se joue ailleurs : philosophie de développement, gouvernance, transparence du code, fréquence des correctifs de sécurité et coût total de possession sur trois ans. Pour un homelab, OPNsense remporte 7 critères sur 12. Pour une entreprise sous contrat de support, pfSense Plus reste compétitif grâce à son TAC Lite à 129 $/an. Pour une administration publique européenne soumise à la directive NIS 2, OPNsense – édité par Deciso aux Pays-Bas – coche toutes les cases.

## Origines et gouvernance : pourquoi le fork de 2015 a tout changé

**pfSense** est né en 2004 d’un fork de m0n0wall mené par Chris Buechler et Scott Ullrich. Le projet a été progressivement repris par **Netgate**, société texane fondée en 2003, qui en assure aujourd’hui le développement, la commercialisation des appliances et la maintenance des deux variantes : *Community Edition* (CE), gratuite, et *Plus*, version commerciale lancée en 2021 sous licence propriétaire. Cette bascule de Plus vers un modèle propriétaire en 2021 – initialement motivée par la lutte contre le piratage des appliances Netgate – a profondément divisé la communauté et précipité la migration de nombreux utilisateurs vers OPNsense.

**OPNsense** est né en janvier 2015 d’un fork de pfSense par **Deciso BV**, société néerlandaise basée à Middelharnis. Le déclencheur : un différend juridique en 2014 entre Netgate et Deciso autour de la marque pfSense, suivi d’une volonté de revenir aux fondamentaux open source. Le projet a fêté ses dix ans le 29 janvier 2025 avec la sortie de la version **25.1 « Ultimate Unicorn »**, bâtie sur FreeBSD 14.2 et PHP 8.3 ; dix-huit mois plus tard, la branche **26.7 « Xenial Xenops »** a pris le relais, poussée jusqu’à la version **26.7.3** le 27 août 2026 avec au passage la bascule vers PHP 8.5. OPNsense est porté par une équipe core composée notamment de Franco Fichtner, ancien committeur FreeBSD reconnu, et publie une nouvelle version mineure tous les quinze jours en moyenne, rythmée par deux versions majeures par an en janvier et en juillet – un calendrier confirmé par le suivi indépendant d’endoflife.date, qui ne recensait plus au 27 août 2026 qu’une seule branche activement supportée, 26.7.x, contre 23 versions antérieures désormais marquées end-of-life, dont la série 26.1 elle-même passée en fin de vie le 15 juillet 2026. Le code, hébergé sur GitHub, est intégralement sous licence BSD-2-Clause – la plus permissive et transparente du monde des pare-feu open source.

Cette divergence d’origine façonne l’ensemble du comparatif : pfSense reste un produit commercial avec une édition communautaire dérivée, tandis qu’OPNsense est un projet communautaire avec une édition business optionnelle (Business Edition à 159 €/an, sans fonctionnalités exclusives). Pour les organisations européennes attachées à la souveraineté numérique, l’origine néerlandaise d’OPNsense et son hébergement git en Europe sont des arguments décisifs face aux pressions américaines exercées par la Section 232 et le CLOUD Act.

## Tableau de spécifications : 14 critères techniques comparés

| Critère | pfSense CE 2.7.x | pfSense Plus 24.x | OPNsense 25.7 | 
|---|---|---|---|
| Première version | 2006 | 2021 | 2015 | 
| Éditeur | Netgate (États-Unis) | Netgate (États-Unis) | Deciso (Pays-Bas) | 
| Licence | Apache 2.0 | Propriétaire | BSD-2-Clause | 
| Système d’exploitation | FreeBSD 14 | FreeBSD 14 | FreeBSD 14 (HardenedBSD historiquement) | 
| Cycle de mises à jour | Annuel | 2 à 3 par an | Bi-hebdomadaire + 2 majeures/an | 
| Interface utilisateur | Bootstrap 3 (menu haut) | Bootstrap 3 + thèmes Plus | MVC Phalcon (sidebar) | 
| WireGuard natif | Package (kernel-mode depuis 2.7) | Kernel-mode natif (Plus 22.05+) | Natif intégré | 
| IDS/IPS | Snort, Suricata (packages) | Snort, Suricata (packages) | Snort, Suricata (intégré) | 
| Plugins/Packages | ~80 packages communauté | ~80 packages + Plus exclusives | 80+ officiels + 250 communauté | 
| Traffic shaping | ALTQ + Limiters | ALTQ + Limiters | Pipes/Queues (dummynet moderne) | 
| Haute disponibilité (HA) | CARP + pfsync | CARP + pfsync | CARP + pfsync | 
| API | Tiers (pfsense-api) | REST API native (Plus 23.09+) | API REST officielle complète | 
| Support officiel | Forum communautaire | TAC Lite 129 $/an | Business Edition 159 €/an | 
| Tarif d’entrée matériel | Auto-hébergement | Netgate 1100 à 189 $ | Deciso DEC600 à 459 € | 

Trois lignes méritent d’être détaillées. D’abord la **licence** : Apache 2.0 et BSD-2-Clause sont toutes deux permissives, mais BSD-2-Clause est nettement plus courte (deux clauses contre dix-huit pages chez Apache 2.0) et n’impose aucune obligation de mention des modifications. C’est la raison pour laquelle de nombreux dérivés commerciaux choisissent désormais OPNsense plutôt que pfSense CE comme base.

