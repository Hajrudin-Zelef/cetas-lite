---
id: collect-261001-general-networking/general-networking/dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026-2
title: "dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "benchmarks", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-general-networking/dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026.md
source_anchor: ""
source_lines: [40, 86]
sha256: 08bddcae6da088869a48ecd7d67b2589f88e1201fcb7295468b1b1757229a567
---

# dlss-4-5-vs-fsr-4-1-128-fps-sur-geforce-now-2026

Ce tableau illustre un point souvent négligé dans les comparatifs génériques : sur GeForce NOW, la question “FSR ou XeSS” ne se pose pas vraiment, puisque le service cloud de NVIDIA n’exécute que sa propre pile logicielle DLSS. FSR 4.1 et XeSS 2 restent pertinents pour les joueurs qui possèdent leur propre PC avec une carte AMD ou Intel, mais ils sont absents de l’infrastructure GeForce NOW. Pour un abonné GeForce NOW, la vraie comparaison est donc “DLSS 4.5 en streaming cloud” contre “FSR 4.1 ou XeSS 2 sur un PC local équivalent”.

## GeForce NOW RTX 5080 : ce que le tarif inclut réellement

Le passage aux GPU RTX 5080 dans les data centers NVIDIA n’a pas fait grimper le prix de l’abonnement Ultimate. NVIDIA maintient sa grille tarifaire officielle à 19,99 dollars par mois, 99,99 dollars pour six mois ou 199,99 dollars pour un an. En euros, les guides d’abonnement 2026 pour l’Europe évoquent un tarif mensuel autour de 10,99 € pour le palier le plus élevé, avec des sessions pouvant atteindre 8 heures, la 4K à 120 im/s, le HDR et le support natif de Reflex, la technologie anti-latence de NVIDIA.

| Palier GeForce NOW | Prix mensuel (zone euro) | Résolution / FPS | Durée de session | DLSS / Reflex | 
|---|---|---|---|---|
| Free | 0 € | 1080p, 60 im/s | 1 heure | Non | 
| Performance | ~5,99 € | 1440p, 60 im/s | 6 heures | Non | 
| Ultimate (RTX 5080) | ~10,99 € | 4K/5K, jusqu’à 120-360 im/s selon résolution | 8 heures | Oui, DLSS 4.5 + Reflex | 
| Ultimate (USD, référence NVIDIA) | 19,99 $/mois | 4K/5K | 8 heures | Oui | 
| Ultimate (6 mois, USD) | 99,99 $ (soit ~16,67 $/mois) | 4K/5K | 8 heures | Oui | 
| Ultimate (annuel, USD) | 199,99 $ (soit ~16,67 $/mois) | 4K/5K | 8 heures | Oui | 

Les tarifs annuels et semestriels restent la meilleure affaire pour qui compte utiliser GeForce NOW toute l’année. À l’inverse, un joueur qui teste ponctuellement le service pendant les vacances ou un lancement de jeu peut se limiter au mois à mois, quitte à payer un peu plus par mois en échange de la flexibilité. Notez que le palier Free ne donne accès ni au matériel RTX 5080 ni à DLSS : ce tier reste cantonné à du matériel mutualisé standard, ce qui exclut d’office toute comparaison d’upscaling IA pour les utilisateurs gratuits.

## Benchmarks : trois sources, trois résultats convergents

Les gains de performance annoncés par chaque éditeur ne se testent jamais dans les mêmes conditions, ce qui rend les comparaisons brutes trompeuses. Voici ce que montrent trois bancs d’essai indépendants publiés en 2026.

### Cyberpunk 2077, 4K, ray tracing activé

Sur ce titre particulièrement gourmand en calcul de trajectoires lumineuses, DLSS 4.5 en mode Qualité délivre un gain de +128 % de FPS par rapport au rendu natif avec anti-aliasing temporel classique (TAA). Les modes plus agressifs progressent logiquement : le mode Performance grimpe jusqu’à +130 % avec des pertes de détail plus visibles, et le mode Ultra Performance dépasse les +180 % au prix d’une image nettement dégradée. Ce test, réalisé fin juin 2026 par des testeurs spécialisés en configurations gaming, confirme que le nouveau modèle Transformer 2 tire le meilleur parti du ray tracing intensif.

### Marvel’s Spider-Man 2, 1440p

Sur ce jeu optimisé pour plusieurs technologies d’upscaling, FSR 4.1 en mode Performance combiné à la génération d’images atteint 105,2 images par seconde, soit un gain de +236 % par rapport au rendu natif. C’est légèrement supérieur au chiffre brut de DLSS 4 Performance + Frame Generation sur la même scène, même si les testeurs notent une image sensiblement plus propre côté NVIDIA. Ce résultat illustre une réalité qui dérange parfois les partisans d’une technologie unique : sur certains titres, l’écart de FPS pur favorise AMD, mais la qualité perçue reste souvent en faveur de DLSS.

### F1 24, 1440p, GPU Arc

Sur les GPU Intel Arc, XeSS 2 combiné à la génération d’images produit le gain le plus spectaculaire des trois technologies sur ce titre : de 48 images par seconde en natif à 186 images par seconde avec XeSS 2 et Frame Generation activés, soit +287 %. Les testeurs soulignent toutefois que cette agressivité se paie en homogénéité d’image, avec des artefacts visibles sur certaines textures fines en mouvement rapide, un problème que DLSS 4.5 a largement corrigé grâce à son modèle Transformer 2.

Sur la latence, un test réalisé sur Battlefield 6 apporte un éclairage complémentaire : à 1080p, la génération d’images de DLSS ajoute seulement 6,4 millisecondes de latence, contre jusqu’à 17,5 millisecondes pour la génération d’images FSR. À 4K, l’écart se resserre étonnamment, DLSS FG ajoutant 17 ms contre 13,4 ms pour FSR FG, un résultat qui montre que l’avantage latence de NVIDIA n’est pas garanti à toutes les résolutions.

## DLSS 4.5 sur GeForce NOW : le cas particulier du streaming cloud

Le cloud gaming ajoute une variable que les bancs d’essai en local ne mesurent pas : la latence réseau entre le PC ou la Steam Deck du joueur et le data center NVIDIA. Sur GeForce NOW RTX 5080, DLSS 4.5 fonctionne en local sur le serveur, donc son gain de FPS interne (les fameux +128 % à +180 % observés sur Cyberpunk 2077) profite directement au flux vidéo envoyé au joueur. Mais la latence perçue par l’utilisateur additionne trois éléments : le rendu GPU assisté par DLSS, l’encodage vidéo côté serveur, et le trajet réseau jusqu’à l’écran du joueur.

C’est là que Reflex, la technologie anti-latence de NVIDIA intégrée nativement à DLSS 4.5, joue un rôle disproportionné en cloud gaming. En réduisant la latence de rendu côté serveur au minimum, Reflex compense en partie l’aller-retour réseau inévitable. Pour un joueur français connecté à un data center européen (les principaux points de présence GeForce NOW en Europe se trouvent en Allemagne, en France et désormais dans les pays nordiques après la migration de Stockholm), la latence totale ajoutée par le trajet réseau reste généralement sous les 15-20 ms sur une connexion fibre stable, ce qui laisse une marge confortable même en ajoutant les quelques millisecondes de Frame Generation.

FSR et XeSS n’ont, à ce jour, aucune présence officielle sur l’infrastructure GeForce NOW : le service reste une vitrine exclusive pour la pile logicielle NVIDIA, ce qui est cohérent avec son modèle économique (vendre du temps de calcul sur ses propres GPU RTX). Les joueurs qui préfèrent l’écosystème AMD ou Intel doivent se tourner vers d’autres services de cloud gaming ou rester sur une configuration locale.

## GeForce NOW face aux autres services de cloud gaming en Europe

GeForce NOW n’est pas seul sur le marché européen du cloud gaming, et la question de l’upscaling par IA se pose différemment selon le service. Shadow PC, l’alternative française du secteur, propose des machines virtuelles complètes avec un GPU dédié par abonné plutôt qu’une mutualisation à la demande, mais ne communique pas sur un équivalent direct de DLSS 4.5 : la technologie d’upscaling dépend alors du jeu installé et du pilote GPU utilisé sur la machine virtuelle, généralement une carte NVIDIA de génération antérieure au RTX 5080. Xbox Cloud Gaming, propulsé par les serveurs Xbox Series X de Microsoft, ne repose pas sur les technologies NVIDIA et n’offre donc ni DLSS ni FSR au sens où l’entend ce comparatif, la mise à l’échelle de l’image étant gérée différemment côté consoles.

