---
id: collect-261001-general-networking/general-networking/nginx-rift-faille-de-18-ans-cvss-9-2-2026-1
title: "Vérifier la version de NGINX installée"
domain: general-networking
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["apache", "exploit", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/nginx-rift-faille-de-18-ans-cvss-9-2-2026.md
source_anchor: ""
source_lines: [1, 53]
sha256: ca7592ac23e3b2cdf48cbd764d9bfbd76178e371b975fb9fc6d2b1f115653b27
---

# Vérifier la version de NGINX installée

Une faille NGINX vieille de dix-huit ans, logée dans le code source depuis 2008, a été activement exploitée quelques jours à peine après sa divulgation publique, le 13 mai 2026. Baptisée **NGINX Rift** et référencée **CVE-2026-42945**, elle affiche un score CVSS de 9,2 sur 10 en version 4.0. Un mois plus tard, le 18 juin 2026, l’éditeur F5 a dû publier un correctif d’urgence hors cycle pour colmater deux autres failles critiques, **CVE-2026-42530** et **CVE-2026-42055**, elles aussi notées 9,2. Puis, en juillet 2026, une quatrième vulnérabilité critique, **CVE-2026-42533** (CVSS 9,2 en version 4.0, 8,1 en version 3.1 selon Orca Security), a nécessité un nouveau correctif via les versions **NGINX 1.31.3** et **1.30.4**, comme le confirme la documentation officielle de nginx.org. En l’espace d’un peu plus de deux mois, le serveur web le plus utilisé au monde a donc cumulé quatre vulnérabilités critiques presque consécutives.

Le sujet dépasse largement le cercle des administrateurs systèmes. Selon les statistiques de W3Techs publiées en juillet 2026, NGINX équipe 31,8 % des sites dont le serveur web est identifiable, loin devant Apache (24,4 % en janvier 2026) et le serveur de Cloudflare (25,8 %). Une faille NGINX critique touche donc, potentiellement, plus d’un site sur trois dans le monde : infrastructures cloud, passerelles Kubernetes, reverse proxies d’entreprise et API gateways compris.

Ce dossier revient sur la chronologie de NGINX Rift, détaille les correctifs disponibles, et analyse pourquoi cette succession de failles critiques relance le débat sur la dette technique accumulée dans les briques logicielles les plus déployées de la planète.

## Qu’est-ce que NGINX Rift, la faille CVE-2026-42945 ?

Cette faille NGINX désigne un débordement de tampon en zone de tas (*heap-based buffer overflow*, référencé CWE-122) situé dans le module `ngx_http_rewrite_module`, la brique chargée de réécrire les URL à la volée. Le bug a été rendu public le 13 mai 2026 par l’équipe d’AlmaLinux, en parallèle de la fiche officielle publiée sur la base de données nationale des vulnérabilités américaine (NVD). Son score CVSS atteint 9,2 sur l’échelle version 4.0 (8,1 en version 3.1), ce qui la classe parmi les vulnérabilités jugées critiques.

Concrètement, un attaquant non authentifié peut déclencher la faille en envoyant une simple requête HTTP forgée vers un serveur mal configuré. Dans le pire des cas, cela peut mener à une exécution de code à distance. Dans le meilleur des cas pour la victime, cela provoque tout de même le plantage du processus de travail (*worker process*) de NGINX, donc une interruption de service. Le nom « Rift » fait référence à la fracture entre la taille de tampon calculée par le serveur et celle réellement écrite en mémoire.

## Chronologie : de la divulgation à l’exploitation active en quelques jours

La rapidité avec laquelle cette faille NGINX est passée du statut de découverte académique à celui de menace active illustre l’état actuel de la course entre attaquants et défenseurs.

- **13 mai 2026** : divulgation publique de CVE-2026-42945, publication simultanée des correctifs par NGINX/F5.
- **15 mai 2026** : premiers dépôts de preuve de concept (PoC) sur GitHub, dont une variante permettant l’exécution de code via une technique de « heap feng shui » inter-requêtes.
- **17 mai 2026** : la société de renseignement sur les vulnérabilités VulnCheck confirme, via des*honeypots* , une exploitation active en conditions réelles à peine quelques jours après la divulgation, comme le rapporte The Hacker News.
- **18 juin 2026** : F5 publie un correctif hors cycle pour deux nouvelles failles critiques distinctes, CVE-2026-42530 et CVE-2026-42055.
- **27 juin 2026** : le NVD met à jour la sévérité de CVE-2026-42945 vers le niveau critique dans le référentiel CVSS v4.0.

Cette chronologie serrée laisse peu de marge de manœuvre aux équipes qui gèrent des parcs de serveurs importants, d’autant que deux jeux de correctifs distincts sont sortis à moins de cinq semaines d’intervalle.

## Anatomie technique de la faille

### Le module de réécriture, cœur du problème

La vulnérabilité apparaît lorsqu’une directive `rewrite` est suivie d’une directive `rewrite`, `if` ou `set`, combinée à une capture PCRE non nommée (du type `$1` ou `$2`) dont la chaîne de remplacement contient un point d’interrogation. NGINX calcule la taille du tampon de destination avec une méthode d’échappement, puis écrit réellement les données avec une méthode différente. Des caractères comme `+`, `%` ou `&` s’étendent différemment selon la méthode utilisée, ce qui provoque un dépassement du tampon alloué. Les octets qui débordent proviennent directement de l’URI envoyée par l’attaquant, ce qui rend la corruption mémoire contrôlable, et non un simple plantage aléatoire.

### ASLR, la différence entre plantage et prise de contrôle

Sur un système où la randomisation de l’espace d’adressage (ASLR) est active et correctement configurée, l’exploitation se traduit « seulement » par un plantage du processus, donc un déni de service. Ce n’est que sur des hôtes où l’ASLR est désactivée, ou contournée par une chaîne d’exploitation plus élaborée, que l’exécution de code à distance devient réellement possible. C’est la variante avec exécution de code, publiée par le chercheur connu sous le pseudonyme jelasin, qui a le plus inquiété la communauté de la sécurité en mai 2026.

## Le 18 juin, rebelote : F5 corrige deux nouvelles failles critiques

Un mois après NGINX Rift, F5 a publié le 18 juin 2026 une mise à jour hors cycle corrigeant quatre vulnérabilités au total : deux critiques et deux de sévérité élevée, selon The Hacker News. Les deux failles critiques sont :

- **CVE-2026-42530** (CVSS v4.0 : 9,2) : une vulnérabilité de type*use-after-free* dans le module`ngx_http_v3_module` , qui gère le protocole HTTP/3 et QUIC.
- **CVE-2026-42055** (CVSS v4.0 : 9,2) : un débordement de tampon en zone de tas affectant les modules`ngx_http_proxy_v2_module` et`ngx_http_grpc_module` .

Contrairement à NGINX Rift, ces deux failles n’étaient, à la date de publication de cet article, pas répertoriées comme activement exploitées et ne figuraient pas dans le catalogue des vulnérabilités connues exploitées (KEV) de la CISA. F5 précise que les produits BIG-IP, BIG-IQ, F5 AI Gateway, F5 Distributed Cloud, F5OS, F5 Silverline, NGINX One Console et Traffix SDC ne sont pas concernés : seuls NGINX Open Source, NGINX Plus, NGINX Gateway Fabric, NGINX Instance Manager et NGINX Ingress Controller le sont. Le correctif publié en juillet 2026 pour CVE-2026-42533 a, lui, élargi le périmètre des composants touchés : selon Orca Security, étaient vulnérables NGINX Open Source de la version 0.9.6 à 1.31.2, NGINX Plus antérieur aux versions 37.0.3.1 et R36 P7, NGINX Ingress Controller (3.5.0–3.7.2, 4.0.0–4.0.1 et 5.0.0–5.4.2), NGINX App Protect WAF (4.9.0–4.16.0 et 5.1.0–5.8.0) ainsi que NGINX Instance Manager (2.16.0–2.22.0).

## Versions affectées et correctifs disponibles

Le tableau ci-dessous résume les trois failles critiques recensées sur l’écosystème NGINX entre mai et juin 2026, auxquelles est venue s’ajouter une quatrième, CVE-2026-42533, corrigée en juillet 2026 par les versions NGINX 1.31.3 et 1.30.4.

| CVE | Score CVSS (v4.0) | Composant touché | Date de divulgation | Statut d’exploitation | 
|---|---|---|---|---|
| CVE-2026-42945 (NGINX Rift) | 9,2 | ngx_http_rewrite_module | 13 mai 2026 | Activement exploitée | 
| CVE-2026-42530 | 9,2 | ngx_http_v3_module (HTTP/3) | 18 juin 2026 | Non exploitée à ce jour | 
| CVE-2026-42055 | 9,2 | ngx_http_proxy_v2_module / ngx_http_grpc_module | 18 juin 2026 | Non exploitée à ce jour | 

