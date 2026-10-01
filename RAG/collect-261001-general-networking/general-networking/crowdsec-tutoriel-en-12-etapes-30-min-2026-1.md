---
id: collect-261001-general-networking/general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026-1
title: "Mettre à jour la liste des paquets et le système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026.md
source_anchor: ""
source_lines: [1, 55]
sha256: 22564b57dc28b71f8d7d0d2e7d46f00488f44604abd3b0ab118be545c54501b6
---

# Mettre à jour la liste des paquets et le système

Les serveurs exposés sur Internet subissent un déluge permanent de scans, de tentatives de brute-force SSH et d’injections web. Selon le bilan annuel *2025 Wrap* publié par CrowdSec, la communauté de l’éditeur bloque désormais près de **26 millions d’attaques par jour**, émanant de **1,5 million d’attaquants uniques quotidiens**, et a recensé environ **8 millions d’adresses IP malveillantes**. Face à cette pression, `fail2ban` montre ses limites : il agit en silo, ne partage rien et ne connaît que les attaques qu’il voit passer sur votre seule machine.

Ce tutoriel CrowdSec vous guide pas à pas, en 12 étapes et environ 30 minutes, pour installer, configurer et exploiter ce moteur de détection collaboratif sur un serveur Linux. À la fin, vous disposerez d’un projet complet et fonctionnel : SSH et Nginx protégés, un pare-feu applicatif (WAF) actif, une Console de supervision branchée, et votre serveur connecté à la plus grande source de renseignement sur les menaces (CTI) communautaire et open source du marché. Le tout avec un logiciel **français**, sous licence MIT, pensé pour la souveraineté numérique européenne.

## CrowdSec en 2026 : un IPS collaboratif made in France

CrowdSec est un moteur de sécurité open source qui détecte les comportements malveillants en analysant vos journaux (logs), puis applique des décisions de blocage via des composants appelés *bouncers*. Sa particularité : chaque IP malveillante détectée par un utilisateur, une fois validée, alimente une **liste de blocage communautaire** redistribuée à l’ensemble du réseau. Autrement dit, si un serveur en Allemagne se fait attaquer par une adresse IP, le vôtre en France peut la bloquer *avant* même qu’elle ne vous cible. C’est le principe du renseignement sur les menaces mutualisé, ou CTI.

Le projet est porté par une société française fondée en 2020 et basée à Montrouge, en région parisienne, par Philippe Humeau (CEO), Thibault Koechlin et Laurent Soubrevilla. CrowdSec a levé **14 millions d’euros** en série A en octobre 2022 (annonce officielle), menée par Supernova Invest aux côtés de l’investisseur historique Breega, pour un total d’environ 20,6 M$ sur trois tours. Le cœur du logiciel, écrit en Go 1.24, est publié sous **licence MIT** sur GitHub où il cumule plus de **14 000 étoiles** et totalise 221 releases au compteur du dépôt officiel, le tag v1.7.8 ayant été publié le 11 mai 2026 comme correctif de sécurité pour **CVE-2026-44982** et **CVE-2026-44981**. Le rythme des sorties s’est encore accéléré depuis : la branche **1.8.0**, publiée en août 2026, a ajouté une détection de bots directement dans le WAF ainsi qu’une source de logs native pour **Kubernetes**, avant que la **1.8.1** ne devienne, en septembre 2026, la dernière version disponible sur le dépôt GitHub officiel ; l’installeur Windows (MSI) de la 1.7.8, référencée comme la version la plus récente sur SourceForge, sortie le 11 mai 2026 et buildée le 27 mai 2026 sous le nom de code « alphaga » selon les relevés de Tech Insider, pèse quant à lui 90,6 Mo.

L’ampleur du réseau parle d’elle-même. Publié en janvier 2026, ce bilan 2025 révèle que les moteurs de sécurité CrowdSec ont analysé près de **320 milliards de lignes de logs** réparties sur **23 000 organisations**, avec une croissance de 65 % du nombre de moteurs connectés sur l’année ; le rapport souligne au passage que la branche Security Engine 1.7 a introduit de nouvelles métriques de parseurs facilitant le débogage sur le terrain. Environ 70 % des attaques détectées sont de type HTTP, et les États-Unis arrivent en tête des pays émetteurs avec 4,8 millions d’adresses IP malveillantes uniques. Ce volume fait de la blocklist CrowdSec un atout défensif difficile à égaler pour un outil local isolé.

## CrowdSec vs Fail2ban : pourquoi migrer en 2026

Beaucoup d’administrateurs arrivent à CrowdSec depuis `fail2ban`, l’outil historique de bannissement d’IP. Les deux partagent une idée commune – lire des logs, repérer des abus, bannir des adresses – mais leur philosophie diffère radicalement. Fail2ban est monolithique : il lit, décide et bannit dans le même processus, à l’aide d’expressions régulières (regex) propres à chaque *jail*. CrowdSec sépare la **détection** (le moteur) de la **remédiation** (les bouncers), ce qui le rend bien plus modulaire et adapté aux infrastructures distribuées.

| Critère | Fail2ban | CrowdSec | 
|---|---|---|
| Langage | Python | Go (binaire compilé) | 
| Détection | Regex locale par jail | Scénarios comportementaux + parseurs | 
| Renseignement (CTI) | Aucun | Blocklist communautaire mondiale | 
| Remédiation | iptables intégré | Bouncers découplés (pare-feu, Nginx, Cloudflare, etc.) | 
| Architecture | Monolithique | Moteur + LAPI + bouncers | 
| Supervision | Ligne de commande seule | Console SaaS gratuite (tableau de bord) | 
| Multi-serveur | Manuel | LAPI centralisée native | 
| Pare-feu applicatif (WAF) | Non | Oui (composant AppSec) | 
| Licence | GPLv2 | MIT | 

Le verdict n’est pas que CrowdSec « tue » fail2ban : pour un serveur unique et minimaliste, fail2ban reste léger et suffisant. Mais dès que vous gérez plusieurs machines, que vous voulez un tableau de bord, une protection web applicative ou bénéficier d’une intelligence collective, le rapport bénéfice/effort penche nettement en faveur de CrowdSec. C’est ce que confirme la communauté francophone : le guide de référence d’IT-Connect présente d’ailleurs CrowdSec comme « une alternative collaborative à Fail2ban ».

### L’architecture en clair : moteur, LAPI, bouncers, Hub et Console

Avant de taper la moindre commande, il faut comprendre les six briques qui composent un déploiement CrowdSec. Cette vision d’ensemble vous évitera 80 % des erreurs de débutant.

- **Le moteur de sécurité (Security Engine)** : l’agent qui lit les logs, les normalise via des*parseurs* , évalue le comportement contre des*scénarios* d’attaque, puis génère des décisions selon des*profils* . C’est la couche de détection (IDS), avec une option de pare-feu applicatif (WAF).
- **La LAPI (Local API)** : l’API locale qui stocke les décisions et fait le lien entre le moteur et les bouncers. Elle écoute par défaut sur`127.0.0.1:8080` .
- **Les bouncers (composants de remédiation)** : ils appliquent les décisions sans jamais les prendre eux-mêmes. Bouncer pare-feu (iptables/nftables), bouncer Nginx, bouncer Cloudflare, etc. C’est la couche de prévention (IPS).
- **Le Hub** : le dépôt communautaire de collections, parseurs et scénarios prêts à l’emploi (SSH, Nginx, WordPress, etc.).
- **La Console** : le tableau de bord SaaS gratuit hébergé sur app.crowdsec.net pour superviser vos moteurs, vos alertes et vos décisions.
- **La CAPI (Central API)** : l’API centrale qui mutualise le renseignement. C’est elle qui vous envoie la blocklist communautaire et reçoit, si vous y consentez, vos signaux anonymisés.

| Composant | Rôle | Port / Emplacement | 
|---|---|---|
| Security Engine | Détection (IDS) | service `crowdsec` | 
| LAPI | API locale des décisions | 8080/tcp (localhost) | 
| Firewall Bouncer | Remédiation réseau (IPS) | iptables / nftables | 
| Nginx Bouncer | Remédiation web | module Lua Nginx | 
| AppSec | Pare-feu applicatif (WAF) | 127.0.0.1:7422 | 
| Console | Tableau de bord SaaS | app.crowdsec.net (443) | 
| CAPI | CTI communautaire | api.crowdsec.net (443) | 

## Prérequis : versions et matériel requis

CrowdSec est remarquablement frugal : il tourne sur un VPS d’entrée de gamme, un Raspberry Pi ou un gros serveur de production sans configuration différente. Voici ce dont vous avez besoin avant de commencer l’installation de CrowdSec.

