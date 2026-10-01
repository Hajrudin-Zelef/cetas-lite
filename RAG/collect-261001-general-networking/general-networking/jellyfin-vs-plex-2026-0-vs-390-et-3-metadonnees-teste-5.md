---
id: collect-261001-general-networking/general-networking/jellyfin-vs-plex-2026-0-vs-390-et-3-metadonnees-teste-5
title: "Exemple de configuration Nginx pour Jellyfin"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "gpu", "intel", "nvidia", "open source"]
source: docs/RAG/collect-261001-general-networking/jellyfin-vs-plex-2026-0-vs-390-et-3-metadonnees-teste.md
source_anchor: ""
source_lines: [290, 335]
sha256: 8906092fc6d107f35d33bd106f8bab5ab5368cde1f719e805a2f901f4a20ce7c
---

# Exemple de configuration Nginx pour Jellyfin

| Cas d’usage | Recommandation | Raison | 
|---|---|---|
| Famille non technique | Plex (Pass à vie) | Interface intuitive, contrôle parental, accès distant en 1 clic | 
| Développeur / sysadmin | Jellyfin | Open source, zéro télémétrie, intégration homelab | 
| Cinéphile 4K HDR | Jellyfin ou Plex | Performances similaires, Jellyfin gratuit, Plex plus de polish | 
| Conformité RGPD | Jellyfin | Aucune donnée collectée, aucun transfert hors UE | 
| Grande bibliothèque musicale | Plex (Plexamp) | Sonic Analysis, mix automatiques, interface audio dédiée | 
| Partage avec amis/famille éloignée | Plex | Partage en un clic vs reverse proxy manuel | 
| Budget zéro | Jellyfin | 100 % gratuit sans aucune limitation | 
| NAS Synology / QNAP | Plex | Package natif, meilleure intégration DSM 7 | 

## Couverture associée

Pour approfondir les sujets connexes, consultez nos guides et comparatifs liés :

## Questions fréquentes (FAQ)

### Jellyfin est-il vraiment 100 % gratuit ?

Oui. Jellyfin est un logiciel libre sous licence GPL v2, développé et maintenu par la communauté. Toutes les fonctionnalités – transcodage matériel, Live TV, DVR, applications mobiles – sont gratuites et le resteront. Le projet est financé par des dons volontaires, pas par des abonnements ou de la publicité.

### Plex collecte-t-il mes données personnelles ?

Plex requiert un compte cloud obligatoire et collecte des métadonnées de lecture (titres regardés, durée, appareil utilisé). Ces données alimentent les recommandations et la plateforme de streaming gratuite avec publicités. Pour les utilisateurs européens soumis au RGPD, cette collecte peut poser des questions de conformité si les données sont transférées vers des serveurs américains.

### Puis-je utiliser Jellyfin sur un Raspberry Pi ?

Oui, Jellyfin fonctionne sur Raspberry Pi 4 et 5. Le Pi 5 offre de meilleures performances pour le transcodage logiciel, mais le transcodage matériel reste limité sur cette plateforme. Pour un usage en Direct Play (sans transcodage), un Raspberry Pi 4 avec 4 Go de RAM est suffisant pour 1-2 streams simultanés.

### Combien coûte Plex Pass en 2026 ?

Au 15 juillet 2026, le Plex Pass coûte 6,99 $ par mois (environ 6,50 €) et 69,99 $ par an (environ 65 €), des tarifs inchangés selon Thurrott mais déjà en hausse de 40 % par rapport aux 4,99 $/mois pratiqués avant mars 2026, selon Bytesized Hosting. Le pass à vie, lui, a connu deux flambées successives : d’abord de 119,99 $ à 249,99 $ (+108 %, documenté par Bytesized Hosting en avril 2026), puis de 249,99 $ à **749,99 $** (environ 695 €) le 1er juillet 2026, une hausse de 200 % rapportée par TechGeeks, AppleInsider, Tweaktown, The FPS Review et le comparateur suisse Digitec. Un nouveau **Plex Pass 5 ans à 249,99 $** (environ 232 €) a été introduit le 2 juillet 2026 comme alternative, la lecture locale de base restant gratuite. Un Remote Watch Pass supplémentaire est disponible à 1,99 $/mois pour accéder aux serveurs d’autres utilisateurs à distance.

### Jellyfin supporte-t-il le transcodage 4K HDR ?

Oui. Jellyfin supporte le transcodage 4K HEVC avec tone mapping HDR→SDR via OpenCL et VAAPI. L’accélération matérielle est gratuite avec les GPU Intel (Quick Sync), NVIDIA (NVENC) et AMD (VAAPI). La qualité du tone mapping s’est améliorée en 2026.

### Est-il difficile de migrer de Plex vers Jellyfin ?

La migration prend environ 1 à 2 heures pour un utilisateur expérimenté. Les fichiers médias restent identiques – vous pointez simplement Jellyfin vers les mêmes dossiers. L’historique de lecture peut être transféré via des scripts communautaires. La partie la plus complexe est la configuration de l’accès distant, qui nécessite un reverse proxy.

### Quel serveur matériel recommandez-vous pour Jellyfin ou Plex ?

Pour 1-3 streams simultanés en transcodage 1080p, un Intel Core i5-12400 (ou supérieur) avec son GPU intégré UHD 730 et 8 Go de RAM est suffisant. Pour le Direct Play sans transcodage, un NAS Synology DS224+ ou un mini PC Intel N100 suffit. Le stockage dépend de votre bibliothèque : comptez environ 50 Go par film 4K et 10 Go par film 1080p.

*Dernière mise à jour : 18 avril 2026*
