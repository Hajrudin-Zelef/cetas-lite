---
id: collect-261001-general-networking/general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026-5
title: "Vérifier l'espace disque disponible (minimum 20 Go recommandé)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "open source"]
source: docs/RAG/collect-261001-general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026.md
source_anchor: ""
source_lines: [342, 388]
sha256: dab15aaf1abf5f93b29fa470ea2475b3e1ea7ba7eb484c55707f906cc92748ce
---

# Vérifier l'espace disque disponible (minimum 20 Go recommandé)

Pour une collectivité ou une PME de moins de 200 postes, ce compromis est presque toujours favorable à OpenVAS : le coût d'un serveur virtuel dédié (souvent inférieur à 50 euros par mois chez un hébergeur européen) reste très inférieur à une licence commerciale équivalente. Pour un grand groupe avec plusieurs milliers d'actifs répartis sur de multiples clouds, l'équation change : le temps d'administration et le manque d'automatisation avancée (tableau de bord multi-tenant, intégration SSO poussée, support commercial avec SLA) peuvent justifier le passage à une solution payante comme Greenbone Enterprise ou un concurrent commercial.

Un point souvent négligé dans ce calcul : le coût d'une absence de scan est presque toujours supérieur au coût d'un outil, gratuit ou payant. Les rapports d'incident analysés année après année montrent régulièrement que la majorité des compromissions exploitent des vulnérabilités connues depuis plusieurs mois, jamais corrigées faute de visibilité. Un outil gratuit correctement configuré et suivi vaut largement mieux qu'un outil payant acheté puis laissé à l'abandon après la démonstration initiale.

## Foire aux questions sur OpenVAS et Greenbone

### OpenVAS est-il vraiment gratuit ?

Oui, le moteur de scan, l'interface GSA et le feed communautaire de tests de vulnérabilités sont entièrement gratuits sous licence open source. Seul le feed « Enterprise », avec une couverture élargie et un support commercial, est payant et vendu par Greenbone AG.

### Combien de temps dure un scan complet avec OpenVAS ?

Cela dépend du nombre de cibles, du profil choisi et de la présence de scans authentifiés. Un scan « Full and fast » sur un seul serveur prend généralement entre 20 minutes et 2 heures. Sur un sous-réseau de 50 machines, comptez plusieurs heures.

### OpenVAS peut-il faire planter un serveur pendant un scan ?

Certains tests actifs, notamment sur des équipements réseau ou IoT fragiles, peuvent provoquer des instabilités. C'est pourquoi il est recommandé d'utiliser des profils de scan moins agressifs sur les systèmes critiques et de toujours tester d'abord hors production.

### La directive NIS2 impose-t-elle explicitement OpenVAS ?

Non, la directive NIS2 n'impose aucun outil précis, mais elle exige une gestion des risques incluant l'identification régulière des vulnérabilités. OpenVAS répond à cette exigence sans coût de licence, ce qui explique son adoption croissante chez les entités essentielles et importantes de taille modeste.

### Quelle est la différence entre OpenVAS et Greenbone Community Edition ?

OpenVAS désigne historiquement le moteur de scan lui-même. Greenbone Community Edition est la suite complète qui inclut ce moteur, la base de données, l'interface web GSA et le service de synchronisation du feed. Dans l'usage courant, les deux noms sont souvent utilisés l'un pour l'autre.

### Peut-on utiliser OpenVAS sur un cluster Proxmox en production ?

Oui, à condition d'exclure les interfaces de migration à chaud de la portée du scan et de privilégier un scan authentifié plutôt qu'un scan réseau agressif, afin de limiter le risque d'interruption sur les nœuds actifs.

### Faut-il un serveur dédié pour faire tourner Greenbone Community Edition ?

Ce n'est pas obligatoire pour un petit environnement de test, mais fortement recommandé en production. Le scanner consomme beaucoup de CPU et de bande passante pendant les scans actifs, ce qui peut ralentir d'autres services hébergés sur la même machine.

### Comment savoir si mon feed OpenVAS est à jour ?

Rendez-vous dans Administration → Feed Status sur l'interface GSA. Le statut doit afficher « Current » avec une date de synchronisation datant de moins de 24 heures, puisque le feed communautaire est mis à jour quotidiennement.

### Peut-on scanner un site web ou une API avec OpenVAS ?

OpenVAS détecte les vulnérabilités classiques d'un serveur web (versions obsolètes, mauvaises configurations TLS, en-têtes de sécurité manquants) mais reste limité pour l'analyse applicative en profondeur d'une API REST ou d'une logique métier spécifique. Pour ce périmètre, un scanner applicatif dédié (DAST) ou un test d'intrusion manuel complète utilement les résultats d'OpenVAS.

### Que faire si un scan OpenVAS remonte trop de faux positifs ?

Passez d'abord à un scan authentifié plutôt qu'un scan réseau seul : la majorité des faux positifs viennent d'une détection de version approximative basée sur une bannière de service, que le scan authentifié corrige en lisant directement l'état réel du système. Utilisez ensuite les Overrides pour documenter les exceptions confirmées, plutôt que d'ignorer silencieusement les résultats douteux.

Pour aller plus loin dans la sécurisation de votre infrastructure après ce scan initial, pensez à croiser vos résultats avec les recommandations de l'ENISA, à consulter la documentation officielle sur openvas.org et le site de l'éditeur Greenbone, puis à documenter chaque correctif appliqué : c'est ce dossier de preuve qui sera demandé lors d'un audit NIS2 ou d'une certification ISO 27001.
