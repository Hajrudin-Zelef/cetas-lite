---
id: collect-261001-general-networking/general-networking/nftables-pare-feu-linux-en-13-etapes-2026-5
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/nftables-pare-feu-linux-en-13-etapes-2026.md
source_anchor: ""
source_lines: [334, 384]
sha256: 9217f5731f77f46ebd619ec4c5163160b0da716c9a3b72ac559b1f38589294e5
---

# Debian / Ubuntu

| Critère | nftables (natif) | ufw | firewalld | 
|---|---|---|---|
| Distribution typique | Toutes, usage direct | Ubuntu, Debian | Fedora, RHEL, Rocky Linux | 
| Courbe d’apprentissage | Élevée | Faible | Moyenne | 
| Granularité des règles | Totale | Limitée | Bonne, via zones | 
| Gestion dynamique (runtime) | Oui, sans coupure | Rechargement complet | Oui, via zones actives | 
| Cas d’usage recommandé | Serveurs de production, routeurs, règles complexes | Postes de travail, petits serveurs perso | Environnements d’entreprise avec zones réseau multiples | 

Pour un serveur de production avec des besoins de filtrage précis, le passage direct par nftables offre le contrôle le plus fin. Pour un poste de travail personnel ou un petit VPS où la simplicité prime, ufw reste un choix pertinent puisqu’il génère lui-même des règles nftables optimisées en coulisses. firewalld convient particulièrement aux environnements d’entreprise qui séparent le trafic par zones de confiance (interne, DMZ, externe) sans vouloir écrire de règles nftables brutes.

La documentation officielle du wiki nftables et le site du projet Netfilter restent les références techniques les plus fiables pour approfondir la syntaxe et suivre les évolutions du projet. Pour les administrateurs sous Ubuntu qui préfèrent une approche progressive, la documentation officielle Ubuntu sur les pare-feu détaille aussi bien ufw que l’usage direct de nftables.

Sur le plan réglementaire, le site de l’ANSSI publie des recommandations de durcissement système régulièrement mises à jour, et le CERT-FR diffuse des bulletins de vulnérabilités qui aident à prioriser quels services filtrer en urgence sur un serveur exposé.

## Questions fréquentes

**nftables remplace-t-il complètement iptables en 2026 ?**

Sur la majorité des distributions modernes, oui : iptables reste disponible en façade pour des raisons de compatibilité, mais les commandes passent par la couche iptables-nft qui traduit tout vers le moteur nftables en interne. Écrire directement en syntaxe nft reste toutefois recommandé pour profiter des sets, verdict maps et flowtables.

**Peut-on utiliser ufw et nftables en même temps sur le même serveur ?**

Non, il faut choisir l’un ou l’autre. ufw génère lui-même des règles nftables en coulisses, donc faire tourner les deux en parallèle avec des règles manuelles crée des conflits et des comportements imprévisibles.

**Comment revenir en arrière si une règle nftables bloque tout l’accès réseau ?**

Depuis une console physique ou un accès KVM, exécutez `nft flush ruleset` pour vider toutes les règles actives, puis redémarrez avec une configuration corrigée. C’est pourquoi garder un accès console de secours est indispensable avant tout test en production.

**nftables suffit-il seul pour sécuriser un serveur, ou faut-il d’autres outils ?**

Un pare-feu filtre le trafic réseau mais ne remplace pas les autres couches de sécurité : mises à jour régulières, authentification par clé SSH plutôt que mot de passe, supervision des logs, et éventuellement un outil de détection d’intrusion comme Suricata pour analyser le contenu du trafic autorisé.

**Quelle est la différence entre reject et drop dans une règle nftables ?**

Drop abandonne silencieusement le paquet sans réponse, ce qui rend le port indistinguable d’un port inexistant pour un attaquant. Reject renvoie un message d’erreur explicite (par exemple ICMP port unreachable), plus poli pour du trafic légitime mais qui révèle davantage d’informations à un scanner.

**Comment tester une configuration nftables sans risquer de couper l’accès SSH ?**

Utilisez toujours `nft -c -f` pour valider la syntaxe avant application, testez depuis une VM locale non critique en premier, et programmez un job de rollback automatique via `at` ou `cron` qui restaure l’ancienne configuration après quelques minutes si vous ne confirmez pas manuellement que tout fonctionne.

**Les sets nftables ralentissent-ils les performances par rapport à des règles individuelles ?**

C’est l’inverse : les sets utilisent des structures de données optimisées (tables de hachage ou arbres) évaluées en temps quasi constant, alors qu’une longue liste de règles individuelles s’évalue séquentiellement. Sur un serveur avec des centaines d’IP à filtrer, les sets sont nettement plus rapides.

**Faut-il désactiver le pare-feu de l’hébergeur si nftables est déjà configuré sur le serveur ?**

Non, les deux couches sont complémentaires plutôt que redondantes. Le pare-feu réseau de l’hébergeur filtre le trafic avant qu’il n’atteigne le serveur, ce qui protège contre des volumes de trafic qu’une configuration locale ne pourrait pas absorber seule, tandis que nftables offre une granularité plus fine, propre à chaque service et à chaque conteneur hébergé sur la machine.

**Pourquoi un port publié par Docker reste-t-il accessible malgré une politique nftables restrictive ?**

Docker insère ses propres chaînes et règles NAT pour exposer les ports de ses conteneurs, souvent avec une priorité qui les place avant les règles écrites manuellement dans la chaîne forward. Publier un conteneur en le liant à une adresse IP précise plutôt qu’à toutes les interfaces, puis vérifier avec nmap depuis l’extérieur, reste la méthode la plus fiable pour éviter les mauvaises surprises.
