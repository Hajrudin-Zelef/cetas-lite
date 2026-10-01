---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-1
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [1, 51]
sha256: 852f4a838130864eb1c7cc7142f5369234e33fe37102eb79ff92b200eb7d9988
---

# server.py - Serveur WebSocket basique

Les WebSockets transforment la façon dont les applications web communiquent en temps réel. Contrairement aux requêtes HTTP classiques où le client doit interroger le serveur en permanence (polling), le protocole WebSocket établit une connexion bidirectionnelle persistante entre le navigateur et le serveur. La bibliothèque `websockets` pour Python, désormais en version 17.1 depuis le 26 août 2026, exige toujours Python 3.11 ou supérieur et ne prend officiellement en charge que la dernière version corrective 3.x.y de chaque branche, selon la documentation mise à jour en août 2026 ; elle offre une API asyncio élégante et conforme à 100 % au standard RFC 6455. Dans ce tutoriel, vous allez construire un système de chat en temps réel complet, du serveur WebSocket au client interactif, en passant par l’authentification, les salons multiples et le déploiement en production.

Ce guide étape par étape couvre les 12 phases de construction d’une application WebSocket fonctionnelle avec Python 3.12+ et la bibliothèque `websockets`. Chaque étape inclut du code testé, des exemples de sortie et des solutions aux erreurs courantes. Que vous construisiez un tableau de bord temps réel, un outil collaboratif ou un système de notifications push, ce tutoriel vous donne les fondations solides pour maîtriser les WebSockets en Python.

## Prérequis et Environnement de Développement

Avant de commencer, assurez-vous que votre environnement de développement est correctement configuré. Le prérequis Python a évolué rapidement en 2026 : la version 16.0 de `websockets`, publiée sur PyPI le 10 janvier 2026 avec le statut « No Known Issues », a d’abord relevé le minimum à Python 3.10, avant que la branche 17.x ne le porte définitivement à Python 3.11 — la documentation officielle, mise à jour en août 2026, précise même que seule la dernière version corrective 3.x.y de chaque branche Python est désormais officiellement supportée. Nous recommandons donc Python 3.12 ou supérieur pour bénéficier des dernières optimisations asyncio. La bibliothèque `websockets`, désormais en version 17.1 depuis le 26 août 2026, reste construite sur l’implémentation Sans-I/O introduite depuis la version 13.0, offrant une architecture plus modulaire et performante que les versions précédentes.

| Outil | Version Requise | Commande de Vérification | 
|---|---|---|
| Python | 3.12+ | `python3 --version` | 
| pip | 24.0+ | `pip --version` | 
| websockets | 16.0 | `pip show websockets` | 
| asyncio | intégré | `python3 -c "import asyncio"` | 
| aiohttp (optionnel) | 3.10+ | `pip show aiohttp` | 
| uvicorn (production) | 0.32+ | `uvicorn --version` | 
| Navigateur | Chrome 120+ / Firefox 120+ | Support WebSocket natif | 

Créez un répertoire de projet et un environnement virtuel :

```
mkdir websocket-chat && cd websocket-chat
python3 -m venv venv
source venv/bin/activate
pip install websockets==16.0
pip install aiohttp  # pour le client HTTP optionnel
pip install python-dotenv  # pour les variables d'environnement
```
Vérifiez que l’installation est correcte en exécutant `python3 -c "import websockets; print(websockets.__version__)"`. Vous devez voir `17.1` s’afficher — la version stable la plus récente depuis le 26 août 2026, dont les wheels ARM ne pèsent que 176 à 177 Ko selon le dépôt **piwheels**, et qui reste également disponible via le paquet Gentoo `dev-python/websockets` déployé sur 11 architectures. Si vous obtenez une erreur `ModuleNotFoundError`, vérifiez que votre environnement virtuel est bien activé.

## Étape 1 : Comprendre le Protocole WebSocket

Le protocole WebSocket, défini par la RFC 6455, fonctionne au-dessus de TCP et commence par un handshake HTTP classique. Le client envoie une requête HTTP avec l’en-tête `Upgrade: websocket`, et le serveur répond avec un code 101 (Switching Protocols). À partir de ce moment, la connexion TCP reste ouverte et les deux parties peuvent envoyer des messages à tout moment, sans avoir à réétablir une connexion.

La différence fondamentale avec le HTTP traditionnel est majeure en termes de performances. Avec le polling HTTP, chaque requête génère un overhead de 800 octets à 2 Ko d’en-têtes. En comparaison, un frame WebSocket n’ajoute que 2 à 14 octets d’overhead. Pour une application de chat envoyant 50 messages par seconde, cela représente une réduction de bande passante de plus de 90 % par rapport au polling HTTP. Le protocole WebSocket supporte aussi bien les messages texte (UTF-8) que les messages binaires, ce qui le rend adapté aussi bien au chat qu’au streaming de données.

| Caractéristique | HTTP Polling | WebSocket | 
|---|---|---|
| Direction | Client → Serveur uniquement | Bidirectionnel | 
| Overhead par message | 800+ octets | 2-14 octets | 
| Latence | Intervalle de polling (100ms-5s) | Temps réel (<10ms) | 
| Connexions persistantes | Non | Oui | 
| Support navigateur | Universel | Tous navigateurs modernes | 
| Port par défaut | 80 (HTTP) / 443 (HTTPS) | 80 (WS) / 443 (WSS) | 

La bibliothèque Python `websockets` maintient une couverture de branche de 100 % pour la conformité RFC 6455, selon sa documentation officielle. C’est la bibliothèque WebSocket la plus utilisée en Python : après la version 14.2 taguée le 19 janvier 2025 puis packagée par openSUSE sous le nom `python313-websockets-14.2-1.1` le 6 mai 2025, le projet a publié la 15.0 le 16 février 2025 avec plusieurs wheels ARMv6/ARMv7 chez **piwheels**, suivie du correctif 15.0.1 le 5 mars 2025 pour Python 3.11 et 3.13. Ce rythme soutenu s’est poursuivi jusqu’à la 16.1.1, documentée comme stable dans le changelog officiel le 31 juillet 2026 — preuve d’une base de tests automatisés solide et d’une API construite sur `asyncio` qui s’intègre naturellement avec l’écosystème Python asynchrone. Face à des alternatives plus anciennes comme `ws4py`, dont la dernière version 0.6.0 ne date que du 2 août 2025 avec 54 contributeurs au total sur GitHub, ou à des distributions embarquées comme Apertis v2025 qui n’embarque encore que la version 10.4 selon Repology, `websockets` conserve une longueur d’avance nette en termes de rythme de publication.

## Étape 2 : Créer un Serveur WebSocket Basique

Commençons par un serveur WebSocket minimal. L’API de `websockets` 17.0.1 utilise le module `websockets.asyncio`, introduit depuis la version 13.0, qui repose sur une implémentation Sans-I/O ; la documentation de la version 16.1, datée du 9 juillet 2026, détaille d’ailleurs les changements incrémentaux qui ont préparé cette transition vers la branche 17.x. Ce serveur écoute sur le port 8765 et renvoie chaque message reçu en écho (echo server). C’est le point de départ classique pour comprendre le cycle de vie d’une connexion WebSocket.

