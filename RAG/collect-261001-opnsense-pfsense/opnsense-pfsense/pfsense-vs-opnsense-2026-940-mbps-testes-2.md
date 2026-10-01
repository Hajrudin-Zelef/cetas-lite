---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/pfsense-vs-opnsense-2026-940-mbps-testes-2
title: "pfsense-vs-opnsense-2026-940-mbps-testes"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmarks", "cyber", "datacenter", "intel"]
source: docs/RAG/collect-261001-opnsense-pfsense/pfsense-vs-opnsense-2026-940-mbps-testes.md
source_anchor: ""
source_lines: [38, 91]
sha256: 1216e46bf714f8499bb1c88900601025cf5c70d403e120b955f4d26a7a9d31d9
---

# pfsense-vs-opnsense-2026-940-mbps-testes

Ensuite, le **cycle de mises à jour**. OPNsense publie en moyenne 26 versions de patch par an (une toutes les deux semaines), ce qui permet d’absorber les correctifs FreeBSD upstream en quelques jours – la branche 26.7 « Xenial Xenops », lancée le 15 juillet 2026, a ainsi enchaîné trois builds de maintenance en six semaines : **26.7.1** le 21 juillet 2026, **26.7.2** le 12 août 2026, puis **26.7.3** le 27 août 2026, tandis que l’édition commerciale **26.4.2 Business Edition**, publiée le 14 août 2026 mais toujours adossée à la branche communautaire 26.1.11, illustre le léger décalage assumé entre Community et Business – preuve de la cadence incrémentale du projet. pfSense CE en publie 1 à 2 par an, et les correctifs de sécurité majeurs arrivent fréquemment dans Plus avant CE. Pour les environnements soumis à NIS 2 ou à la doctrine cyber publiée par l’ANSSI dans la stratégie nationale de cybersécurité 2026-2030, ce différentiel est décisif.

Enfin, l’**API REST**. OPNsense propose depuis 2018 une API REST officielle exhaustive couvrant 100 % des fonctionnalités exposées dans l’interface web. pfSense Plus a rattrapé son retard à partir de la version 23.09, mais reste limité côté CE où il faut recourir au package communautaire *pfsense-api*. Cette API native est cruciale pour les pipelines DevOps, l’intégration Ansible et les déploiements infrastructure-as-code.

## Benchmarks de performance : 1 Gbit/s, 10 Gbit/s, WireGuard et IDS

Les bancs d’essai 2026 publiés par DIYMediaServer, Home Network Guy et FreeBSDsoftware.org convergent : sur du matériel identique, l’écart de performance entre pfSense et OPNsense est négligeable. La machine de référence est un Mini PC Intel Core i5-8500 (6 cœurs, 4,1 GHz turbo), 16 Go de DDR4, carte réseau Intel i350-T4 quad-port, SSD NVMe 256 Go.

| Banc d’essai | pfSense CE 2.7.2 | OPNsense 25.7 | Écart | Source | 
|---|---|---|---|---|
| Routage 1 Gbit/s WAN-LAN | 938 Mbps | 940 Mbps | 0,2 % | DIYMediaServer 2026 | 
| WireGuard 1 Gbit/s | 710 Mbps | 720 Mbps | 1,4 % | FreeBSDsoftware 2026 | 
| OpenVPN AES-256-GCM | 375 Mbps | 380 Mbps | 1,3 % | HomeNetworkGuy 2026 | 
| IPsec IKEv2 AES-NI | 895 Mbps | 905 Mbps | 1,1 % | DIYMediaServer 2026 | 
| Suricata IDS 3 rulesets | 670 Mbps (-28 %) | 680 Mbps (-27 %) | 1,5 % | FreeBSDsoftware 2026 | 
| RAM idle 4 interfaces | 312 Mo | 328 Mo | +5 % | HomeNetworkGuy 2026 | 
| Boot to login prompt | 42 secondes | 38 secondes | -9,5 % | DIYMediaServer 2026 | 
| Sessions PF concurrentes | ~250 000 | ~250 000 | 0 % | Forums Reddit r/PFSENSE 2026 | 

Ces chiffres confirment que le hardware reste le facteur limitant n°1. Sur une carte réseau gigabit, les deux plateformes saturent le lien physique : aucune ne laisse plus de 2 % de marge à l’autre. La différence se creuse en revanche sur les charges 10 Gbit/s. Sur une appliance Netgate 8300 équipée d’un Xeon D-1700 et de cartes Mellanox ConnectX-6, pfSense Plus 24.03 atteint 9,4 Gbit/s en routage pur grâce à l’optimisation *netmap* activée par défaut. OPNsense 25.7 plafonne à 8,8 Gbit/s sur le même matériel sans réglage manuel des *tx/rx queues* via `sysctl`. L’écart se résorbe après tuning, mais pfSense Plus garde l’avantage en sortie d’usine pour les charges 10 Gbit/s.

Côté **traffic shaping**, OPNsense bénéficie depuis 2018 du portage du système *pipes/queues* issu de dummynet, plus moderne et performant qu’ALTQ utilisé par pfSense. Sur un test de QoS avec 200 utilisateurs simulés et trois classes de trafic prioritaire (VoIP, vidéoconférence, web), OPNsense maintient 940 Mbps avec une gigue inférieure à 1 ms, contre 920 Mbps et 2,3 ms de gigue chez pfSense. Cet écart de 100 % sur la latence devient critique pour la téléphonie SIP en entreprise.

### Performance WireGuard : le verdict 2026

Le grand chantier des deux plateformes en 2024-2025 a été l’intégration native de WireGuard en mode kernel. OPNsense a intégré WireGuard officiellement avec la version 22.7 en juillet 2022 via le module FreeBSD `if_wg`. pfSense Plus a embarqué WireGuard kernel-mode dès la 22.05, et pfSense CE 2.7.0 (publié en juin 2023) l’a finalement reçu en stable. Sur du matériel équivalent, les deux atteignent désormais ~700 Mbps avec un seul tunnel ChaCha20-Poly1305, soit 13 % de plus qu’OpenVPN AES-256-GCM. Pour comparer plus en détail les deux protocoles VPN, notre comparatif WireGuard vs OpenVPN détaille chaque scénario d’usage.

## Tarification complète : licences, appliances et coût total de possession

La gratuité est un mythe partiel des deux côtés. pfSense CE et OPNsense sont effectivement téléchargeables sans bourse délier, mais le coût total de possession sur trois ans dépend des appliances, du support et des plugins commerciaux comme Zenarmor (ex-Sensei) ou pfBlockerNG Devel.

| Offre | Tarif 2026 | Inclus | Cible | 
|---|---|---|---|
| pfSense CE | 0 € | Logiciel + forum | Homelab, PME bricoleuse | 
| pfSense Plus Home+Lab (TAC Lite) | 129 $/an | 1 instance, support email 8×5 | Homelab avancé | 
| pfSense Plus Standard | 799 $/an + appliance | Support 24×7, mises à jour Plus | PME 50-200 postes | 
| Netgate 1100 | 189 $ | ARM Marvell, 1 Gbps WAN | Homelab d’entrée | 
| Netgate 4100 | 699 $ | Atom C3558, 4x 1 GbE + 2x 2,5 GbE | PME | 
| Netgate 8300 | 3 499 $ | Xeon D-1700, 4x 10 GbE SFP+ | Datacenter, ETI | 
| OPNsense Community | 0 € | Logiciel + forum + plugins | Tous | 
| OPNsense Business Edition | 159 €/an | Branche stable, support email | PME, administrations | 
| Deciso DEC600 | 459 € | Atom C3338, 4x 1 GbE | PME, agences | 
| Deciso DEC2700 | 2 199 € | Atom C3758, 4x 10 GbE SFP+ | Datacenter | 
| Zenarmor SME | 9 $/mois/appliance | NGFW DPI, AppControl | Add-on OPNsense | 
| pfBlockerNG (CE/Plus) | 0 € | DNS blocklists, GeoIP | Add-on pfSense | 

Sur un budget homelab à 300 €, OPNsense + un mini-PC Protectli Vault FW6E (399 $ avec 8 Go de RAM) reste la meilleure option pour qui sait flasher un firmware. Sur un budget PME 50 utilisateurs, pfSense Plus + Netgate 4100 + TAC Standard revient à **1 498 $** la première année (699 + 799), contre **2 358 €** pour Deciso DEC2700 + OPNsense Business Edition (2 199 + 159). Le matériel européen reste plus cher de 30 à 40 % à puissance équivalente, en partie à cause des droits de douane de la Section 232 sur les semi-conducteurs et de l’absence d’effet d’échelle face à Netgate.

À l’horizon trois ans, le calcul change. OPNsense Community + Protectli FW6E coûte 399 $ d’investissement initial sans abonnement récurrent, soit **9,1 fois moins** que pfSense Plus Standard sur Netgate 4100 (1 498 $ la première année + 799 $/an x 2 = 3 096 $). Pour les administrations soumises à des règles de marchés publics restrictives, OPNsense Community est de loin l’option la plus budget-friendly.

## Interface utilisateur : MVC Phalcon vs Bootstrap 3

L’écart le plus visible entre les deux plateformes apparaît dès la première connexion à l’interface web. **pfSense** conserve l’architecture héritée de m0n0wall : un menu de navigation horizontal en haut de page, des sous-menus déroulants et un thème Bootstrap 3 daté esthétiquement, mais redoutablement efficace pour les power users qui connaissent les chemins d’accès par cœur. La densité d’information par écran est élevée, ce qui plaît aux administrateurs expérimentés.

