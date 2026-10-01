---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-7
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "distribution", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [809, 926]
sha256: 3df2f8e9b5e836fb26f19c7b6e1e20a2df1fc4375fc200447fec0da54906b939
---

# server.py - Serveur WebSocket basique

**Erreur 2 : Ne pas gérer les déconnexions dans le broadcast.** Si un client se déconnecte pendant un `asyncio.gather` de broadcast, une exception `ConnectionClosed` est levée et interrompt l'envoi aux autres clients. La solution est d'utiliser `return_exceptions=True` dans `asyncio.gather()`.

**Erreur 3 : Utiliser des variables globales mutables sans protection.** Modifier un `set` ou un `dict` global depuis plusieurs coroutines concurrentes peut sembler fonctionner en développement, mais provoque des comportements imprévisibles sous charge. Pour les structures de données partagées complexes, utilisez un `asyncio.Lock()`.

**Erreur 4 : Ignorer la limite de taille des messages.** Sans le paramètre `max_size`, un client malveillant peut envoyer un message de plusieurs gigaoctets et faire crasher le serveur par épuisement mémoire. Définissez toujours `max_size` (1 Mo est un bon défaut pour le chat).

**Erreur 5 : Tester uniquement avec un seul client.** Les bugs de concurrence n'apparaissent qu'avec plusieurs connexions simultanées. Écrivez des tests avec au moins 5 clients connectés en parallèle pour détecter les conditions de concurrence (race conditions) et les problèmes de broadcast.

## Dépannage : 10 Problèmes Fréquents et Solutions

Voici les problèmes les plus fréquemment rencontrés lors du développement et du déploiement d'applications WebSocket en Python, avec leurs solutions détaillées.

| Problème | Cause | Solution | 
|---|---|---|
| `ConnectionRefusedError` | Serveur non démarré ou mauvais port | Vérifier que le serveur tourne : `lsof -i :8765` | 
| `InvalidStatusCode: 403` | Pare-feu ou proxy bloque les WebSockets | Vérifier les en-têtes `Upgrade` dans la config Nginx | 
| `ConnectionClosedError: 1006` | Déconnexion anormale (réseau coupé) | Implémenter le heartbeat avec `ping_interval` | 
| `InvalidMessage` | Données non-JSON envoyées au serveur | Entourer `json.loads()` d'un`try/except` | 
| Messages non reçus | `await` manquant sur`send()` | Vérifier que chaque `send()` est précédé de`await` | 
| Mémoire qui augmente | Clients déconnectés non retirés du `set` | Utiliser `try/finally` pour toujours appeler`discard()` | 
| Timeout Nginx `504` | `proxy_read_timeout` trop bas | Augmenter à 86400s pour les connexions longues | 
| `ssl.SSLError` en production | Certificat TLS invalide ou expiré | Renouveler avec `certbot renew` et redémarrer | 
| Erreur CORS dans le navigateur | Origine non autorisée | Configurer `origins` dans`websockets.serve()` | 
| Performance dégradée à 500+ clients | Broadcast séquentiel au lieu de concurrent | Utiliser `asyncio.gather()` pour le broadcast parallèle | 

Pour le problème de mémoire, un outil de diagnostic utile est `tracemalloc` intégré à Python. Ajoutez `tracemalloc.start()` au début de votre serveur et affichez les statistiques avec `tracemalloc.get_traced_memory()` pour identifier les fuites. Pour les problèmes de performance sous charge, utilisez `websockets.broadcast()` (disponible depuis la version 10.0) qui est optimisé pour envoyer le même message à de nombreux clients simultanément, sans créer une tâche par client.

## Astuces Avancées pour les WebSockets en Python

Pour aller au-delà des bases, voici des techniques avancées qui distinguent les applications WebSocket professionnelles des projets débutants.

### Compression des Messages avec Per-Message Deflate

Le protocole WebSocket supporte la compression via l'extension Per-Message Deflate (RFC 7692). La bibliothèque `websockets` l'active par défaut. Pour les messages JSON répétitifs (comme les mises à jour de cours de bourse), la compression réduit la bande passante de 60 à 80 %. Si votre serveur est limité en CPU plutôt qu'en bande passante, désactivez-la avec `compression=None`.

### Scaling Horizontal avec Redis Pub/Sub

Un seul processus Python peut gérer des milliers de connexions WebSocket grâce à asyncio. Mais quand vous déployez plusieurs instances derrière un load balancer, les messages d'un serveur ne sont pas envoyés aux clients connectés à un autre serveur. La solution standard est d'utiliser Redis Pub/Sub comme bus de messages inter-processus. Chaque instance Python s'abonne à un canal Redis et publie les messages dessus, garantissant que tous les clients reçoivent tous les messages, quelle que soit l'instance à laquelle ils sont connectés.

```
# scaled_server.py - WebSocket avec Redis Pub/Sub (extrait)
import asyncio
import websockets
import json
import redis.asyncio as redis
REDIS_URL = "redis://localhost:6379"
CHANNEL = "chat:messages"
CLIENTS = set()
async def redis_listener():
    """Écoute les messages Redis et les diffuse aux clients locaux."""
    r = redis.from_url(REDIS_URL)
    pubsub = r.pubsub()
    await pubsub.subscribe(CHANNEL)
    async for msg in pubsub.listen():
        if msg["type"] == "message":
            data = msg["data"].decode()
            await asyncio.gather(
                *[c.send(data) for c in CLIENTS],
                return_exceptions=True
            )
async def handler(websocket):
    """Publie les messages sur Redis au lieu du broadcast local."""
    r = redis.from_url(REDIS_URL)
    CLIENTS.add(websocket)
    try:
        async for message in websocket:
            # Publier sur Redis (toutes les instances recevront)
            await r.publish(CHANNEL, message)
    except websockets.exceptions.ConnectionClosed:
        pass
    finally:
        CLIENTS.discard(websocket)
    await r.aclose()
async def main():
    # Lancer l'écoute Redis en parallèle du serveur
    listener_task = asyncio.create_task(redis_listener())
    async with websockets.serve(handler, "0.0.0.0", 8765):
        await asyncio.Future()  # Tourner indéfiniment
if __name__ == "__main__":
    asyncio.run(main())
```
Cette architecture permet de scaler horizontalement à des dizaines de milliers de connexions simultanées en ajoutant simplement des instances Python derrière un load balancer. Redis gère la distribution des messages entre les instances avec une latence de l'ordre de la milliseconde.

## Projet Complet : Structure et Fichiers

Voici la structure complète du projet que nous avons construit au fil de ce tutoriel. Chaque fichier est fonctionnel et prêt à être utilisé. Le serveur principal combine les salons, l'authentification basique et le broadcast optimisé.

```
# Structure du projet
websocket-chat/
├── venv/
├── server.py              # Étape 2 - Serveur écho basique
├── client.py              # Étape 3 - Client Python
├── chat_server.py         # Étape 4 - Broadcast multi-clients
├── index.html             # Étape 5 - Interface web
├── rooms_server.py        # Étape 6 - Salons de chat
├── auth_server.py         # Étape 7 - Authentification
├── heartbeat_client.py    # Étape 8 - Reconnexion auto
├── fastapi_ws.py          # Étape 9 - Intégration FastAPI
├── binary_server.py       # Étape 10 - Données binaires
├── test_chat.py           # Étape 11 - Tests
├── production_server.py   # Étape 12 - Production
├── scaled_server.py       # Avancé - Redis Pub/Sub
├── requirements.txt
└── nginx.conf             # Config reverse proxy
# requirements.txt
websockets==16.0
fastapi>=0.115.0
uvicorn>=0.32.0
redis>=5.2.0
pytest>=8.3.0
pytest-asyncio>=0.24.0
python-dotenv>=1.0.0
```
Pour lancer le projet complet, clonez ou créez le répertoire, installez les dépendances avec `pip install -r requirements.txt`, puis démarrez le serveur de votre choix. Le `chat_server.py` est le meilleur point de départ pour un chat fonctionnel. Ouvrez `index.html` dans plusieurs onglets de navigateur pour tester la communication en temps réel.

## WebSocket vs Alternatives : Quand Utiliser Quoi

