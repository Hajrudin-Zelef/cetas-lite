---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-8
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [927, 960]
sha256: 436e0073208fa205aa91ed44102d53e12e97f559bad0c2ce2c03d18724e68c63
---

# server.py - Serveur WebSocket basique

Les WebSockets ne sont pas la seule technologie pour la communication en temps réel. Selon votre cas d'usage, d'autres protocoles peuvent être plus adaptés. Voici un comparatif pour vous aider à choisir la bonne technologie.

| Technologie | Direction | Cas d'Usage Idéal | Latence | Complexité | 
|---|---|---|---|---|
| WebSocket | Bidirectionnel | Chat, jeux, collaboration | <10ms | Moyenne | 
| Server-Sent Events (SSE) | Serveur → Client | Notifications, flux d'actualités | <50ms | Faible | 
| HTTP Long Polling | Simulé bidirectionnel | Compatibilité anciens navigateurs | 100ms-5s | Faible | 
| WebRTC | Pair à pair | Vidéo, audio, partage d'écran | <5ms | Élevée | 
| gRPC Streaming | Bidirectionnel | Microservices, IoT | <10ms | Élevée | 
| MQTT | Pub/Sub | IoT, capteurs, domotique | <50ms | Moyenne | 

Les WebSockets sont le choix optimal quand vous avez besoin d'une communication bidirectionnelle en temps réel avec une latence minimale. Si votre application n'a besoin que de recevoir des mises à jour du serveur (notifications, cours de bourse), les Server-Sent Events (SSE) sont plus simples à mettre en œuvre et fonctionnent nativement avec HTTP/2. Pour la communication pair-à-pair (vidéo, audio), WebRTC est la technologie appropriée, bien que les WebSockets soient souvent utilisés comme canal de signalisation pour WebRTC.

## FAQ : Questions Fréquentes sur les WebSockets Python

**Combien de connexions WebSocket un serveur Python peut-il gérer ?** Un seul processus Python avec asyncio peut gérer entre 10 000 et 50 000 connexions simultanées, selon la charge de travail par connexion et les ressources du serveur. Le facteur limitant est généralement la mémoire (chaque connexion consomme environ 10 à 50 Ko) et le CPU pour le traitement des messages. Avec le scaling horizontal via Redis Pub/Sub, vous pouvez atteindre des centaines de milliers de connexions.

**Les WebSockets fonctionnent-ils derrière un CDN comme Cloudflare ?** Oui, Cloudflare supporte les WebSockets sur tous ses plans. Cependant, des limitations s'appliquent : le timeout d'inactivité est de 100 secondes par défaut. Assurez-vous d'implémenter un heartbeat avec un intervalle inférieur à 100 secondes pour maintenir la connexion active à travers le CDN.

**Quelle est la différence entre `websockets` et `socket.io` en Python ?** La bibliothèque `websockets` implémente le protocole WebSocket pur (RFC 6455). Socket.IO est un protocole de niveau supérieur qui utilise WebSocket comme transport mais ajoute des fonctionnalités comme la reconnexion automatique, les rooms, et le fallback vers le long polling. Si vous contrôlez le client et le serveur, `websockets` est plus léger et plus performant. Si vous avez besoin de compatibilité avec des clients JavaScript Socket.IO existants, utilisez `python-socketio`.

**Comment déboguer les messages WebSocket dans le navigateur ?** Ouvrez les outils de développement (F12), allez dans l'onglet Réseau (Network), filtrez par "WS". Cliquez sur la connexion WebSocket pour voir tous les messages échangés en temps réel, avec leur taille et leur direction (envoyé/reçu). Chrome et Firefox affichent aussi les frames ping/pong du protocole.

**Les WebSockets consomment-ils beaucoup de batterie sur mobile ?** Une connexion WebSocket idle consomme très peu de batterie car elle n'échange aucune donnée entre les heartbeats. Cependant, les heartbeats trop fréquents (intervalles inférieurs à 30 secondes) empêchent le modem radio du téléphone de passer en mode veille. Pour les applications mobiles, utilisez un intervalle de ping de 60 secondes ou plus.

**Peut-on utiliser les WebSockets avec Django ?** Django ne supporte pas nativement les WebSockets car il est basé sur WSGI, un protocole synchrone. Pour ajouter les WebSockets à Django, utilisez Django Channels qui remplace WSGI par ASGI (Asynchronous Server Gateway Interface). Alternativement, vous pouvez exécuter un serveur `websockets` séparé à côté de votre application Django, communiquant via Redis.

**Quelle est la taille maximale d'un message WebSocket ?** Le protocole WebSocket n'impose pas de limite de taille — les messages peuvent théoriquement atteindre 2^63 octets. En pratique, la bibliothèque `websockets` impose un défaut de 1 Mo (`max_size=2**20`). Augmentez cette valeur uniquement si nécessaire (transfert de fichiers) et toujours avec une validation côté serveur pour éviter les attaques par épuisement mémoire.

**Comment monitorer un serveur WebSocket en production ?** Exportez des métriques Prometheus depuis votre serveur : nombre de connexions actives, messages par seconde, latence de broadcast, taille des messages. La bibliothèque `websockets` fournit des hooks pour intercepter les connexions et les messages. Combinez avec Grafana pour visualiser les tableaux de bord en temps réel et configurer des alertes sur les seuils critiques.

### Couverture Connexe

Pour approfondir vos connaissances en développement temps réel et Python, consultez ces articles connexes :
