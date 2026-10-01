---
id: collect-261001-general-networking/general-networking/nginx-rift-faille-de-18-ans-cvss-9-2-2026-2
title: "Vérifier la version de NGINX installée"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["agent", "apache", "cybersecurity", "exploit", "incident", "mai", "open source", "zero-day"]
source: docs/RAG/collect-261001-general-networking/nginx-rift-faille-de-18-ans-cvss-9-2-2026.md
source_anchor: ""
source_lines: [54, 106]
sha256: 4a2451f480bab1d1f59e57f126d926c5bf148c5f9a5530abc9ff62b660f8b7ab
---

# Vérifier la version de NGINX installée

Pour identifier précisément si un parc est concerné, le second tableau détaille les versions vulnérables et les versions corrigées, produit par produit.

| Produit | Versions vulnérables | Version corrigée | 
|---|---|---|
| NGINX Open Source (Rift) | 0.6.27 à 1.30.0 | 1.30.1 ou 1.31.0+ | 
| NGINX Plus (Rift) | R32 à R36 | R32 P6, R35 P2 ou R36 P4 | 
| NGINX Open Source (42530/42055) | 1.31.0 à 1.31.1 | 1.31.2 | 
| NGINX Plus (42530/42055) | R32 à R36 et 37.0.0 à 37.0.1 | R32 P6 / R36 P4 ou 37.0.2.1 | 
| NGINX Gateway Fabric | 2.0.0 à 2.6.3 et 1.3.0 à 1.6.2 | 2.6.4 (ou version ultérieure) | 
| NGINX Ingress Controller | 3.5.0–3.7.2, 4.0.0–4.0.1, 5.0.0–5.5.0 | Versions ultérieures publiées par F5 | 

## NGINX, colonne vertébrale du web mondial : quel impact réel ?

NGINX n’est pas un simple serveur web parmi d’autres. Selon W3Techs, il occupait encore 33,3 % du marché en janvier 2026 avant de refluer légèrement à 31,8 % en juillet, un tassement à mettre en lien avec la progression du serveur proxy de Cloudflare (25,8 %). Il reste malgré tout la référence numéro un, devant Apache, tombé à 24,4 % en début d’année. Au-delà des sites web classiques, NGINX sert aussi de reverse proxy, de répartiteur de charge, de passerelle API et de contrôleur d’entrée (*ingress controller*) pour une écrasante majorité de clusters Kubernetes dans le monde.

C’est cette omniprésence qui complique la réponse à incident. De nombreuses équipes ne savent pas précisément où NGINX est déployé dans leur organisation : il se cache aussi bien dans des conteneurs applicatifs que dans des appliances réseau, des CDN internes ou des passerelles d’API tierces. Une faille NGINX critique dans son module de réécriture d’URL, fonctionnalité utilisée par une part significative des configurations de production, transforme donc un correctif technique isolé en chantier d’inventaire à l’échelle de toute l’infrastructure.

## Une dette technique vieille de 18 ans : comment est-ce possible ?

Le code fautif du module de réécriture existe depuis 2008, année où NGINX ne pesait encore qu’une fraction de son poids actuel sur le web mondial. Pendant dix-huit ans, cette logique de calcul de tampon a survécu à des centaines de versions, d’audits de sécurité et de revues de code, sans être détectée. Ce n’est pas un cas isolé dans l’histoire des logiciels d’infrastructure : les composants les plus anciens et les plus stables sont aussi, paradoxalement, ceux que l’on relit le moins souvent avec un œil neuf, précisément parce qu’ils sont considérés comme éprouvés.

NGINX Rift rappelle une leçon récurrente de la sécurité des infrastructures critiques : la maturité d’un logiciel ne garantit pas l’absence de failles, elle garantit seulement que les failles restantes sont plus difficiles à débusquer. Pour un projet aussi central que NGINX, chaque ligne de code écrite en 2008 reste, en 2026, une surface d’attaque potentielle tant qu’elle n’a pas été formellement revérifiée.

## NGINX Rift face aux autres failles critiques de 2026

2026 n’aura pas été une année calme pour les équipes de sécurité. Cette faille NGINX s’ajoute à une liste déjà longue de vulnérabilités critiques touchant des briques d’infrastructure massivement déployées, que Tech Insider a suivies au fil de l’eau.

| Faille | Produit | Donnée clé | 
|---|---|---|
| NGINX Rift (CVE-2026-42945) | NGINX (F5) | CVSS 9,2, faille vieille de 18 ans, exploitée en quelques jours | 
| Faille Chrome zero-day | Google Chrome | CVSS 8,8, 65,7 % des instances exposées | 
| SharePoint RCE | Microsoft SharePoint | CVSS 8,8, échéance de correction au 4 juillet | 
| Zero-day Adobe ColdFusion | Adobe ColdFusion | Exploitée en moins de 2 heures après divulgation | 
| Cisco SD-WAN | Cisco | 7 zero-days recensés, dont 2 notés CVSS 10 | 

Ce tableau illustre une tendance de fond : les délais entre divulgation et exploitation active se comptent désormais en heures ou en jours, rarement en semaines. Sur ce point précis, cette faille NGINX se situe dans la moyenne haute de 2026, entre les quelques heures qui ont suffi pour la faille Adobe ColdFusion et les fenêtres plus longues observées sur d’autres produits.

## La réponse encore timide des autorités françaises et européennes

Le CERT-FR a toutefois publié, le 25 mai 2026, un avis consacré à NGINX mais portant sur une référence distincte, **CVE-2026-9256**, qui recensait comme vulnérables les versions 0.1.17 à 1.31.0 de NGINX Open Source et la version 37.0.0 de NGINX Plus. À la date de publication de cet article, l’organisme n’a en revanche pas émis d’avis dédié et numéroté spécifiquement à NGINX Rift ni aux deux CVE corrigées le 18 juin, même s’il continue de surveiller en continu les vulnérabilités affectant les briques web les plus déployées sur le territoire français. Au Canada, le Centre canadien pour la cybersécurité a, de son côté, publié en juillet 2026 l’avis AV26-704, qui recense NGINX Agent dans ses versions 2.37.0 à 2.46.5 parmi les composants concernés. Cette communication en ordre dispersé, faille par faille et pays par pays, illustre une difficulté structurelle : les autorités nationales ne peuvent matériellement pas publier un avis dédié pour chacune des vulnérabilités critiques touchant les logiciels open source les plus répandus.

Le dossier NGINX Rift s’inscrit dans un contexte réglementaire européen déjà mouvant. Tech Insider a récemment détaillé le renvoi de la France devant la Cour de justice de l’Union européenne pour retard de transposition de la directive NIS2, ainsi que la montée en puissance annoncée du Cybersecurity Act 2, qui doit renforcer les moyens de l’ENISA. Les deux textes visent justement à réduire les délais de correction des vulnérabilités critiques dans les infrastructures numériques essentielles, un objectif que la succession de failles NGINX rend d’autant plus concret.

## Quelles conséquences pour les entreprises et les fournisseurs cloud ?

Pour les hébergeurs, les fournisseurs de CDN et les éditeurs SaaS qui embarquent NGINX ou des dérivés comme OpenResty, la succession des trois CVE de 2026 se traduit par un cycle de correctifs en cascade : chaque nouvelle version corrigée doit être re-testée, re-packagée puis redéployée sur des flottes parfois composées de dizaines de milliers d’instances. Le rapport annuel sur la réponse à incident publié par l’unité de recherche Unit 42 de Palo Alto Networks souligne, dans son édition 2026, la place croissante que prennent les vulnérabilités dans les composants d’infrastructure partagée parmi les vecteurs d’intrusion initiaux traités par les équipes de réponse à incident.

Pour une entreprise française de taille intermédiaire qui héberge son propre reverse proxy NGINX, le coût n’est pas seulement technique. Il faut mobiliser des équipes DevOps ou SRE en urgence, parfois interrompre des déploiements planifiés, et documenter la remédiation pour répondre aux obligations de traçabilité qui accompagnent désormais la directive NIS2. Multiplié par trois correctifs critiques en cinq semaines, l’effort cumulé pèse sur des équipes techniques déjà sollicitées par ailleurs.

## Ce que les équipes techniques doivent faire dès maintenant

La priorité immédiate consiste à identifier précisément les versions de NGINX déployées, puis à planifier la montée de version en fonction du niveau d’exposition de chaque instance (accessible depuis Internet ou non). Les commandes suivantes permettent de vérifier rapidement la version installée sur un serveur Linux :

