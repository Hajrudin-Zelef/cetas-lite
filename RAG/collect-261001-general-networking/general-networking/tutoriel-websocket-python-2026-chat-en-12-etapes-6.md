---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-6
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [627, 808]
sha256: b975b508eccaf04fa2b82cc4c8c5635ca00b0f1fd0897d6247b35eaec3b70bd2
---

# server.py - Serveur WebSocket basique

```
# test_chat.py - Tests d'intégration WebSocket
import asyncio
import pytest
import websockets
import json
# Installation : pip install pytest pytest-asyncio
SERVER_PORT = 8799  # Port dédié aux tests
@pytest.fixture
async def chat_server():
    """Lance un serveur de test éphémère."""
    clients = set()
    async def handler(ws):
        clients.add(ws)
        try:
            async for msg in ws:
                for c in clients:
                    if c != ws:
                        await c.send(msg)
        except websockets.exceptions.ConnectionClosed:
            pass
        finally:
            clients.discard(ws)
    server = await websockets.serve(handler, "localhost", SERVER_PORT)
    yield server
    server.close()
    await server.wait_closed()
@pytest.mark.asyncio
async def test_echo_message(chat_server):
    """Vérifie qu'un message est bien diffusé."""
    async with websockets.connect(f"ws://localhost:{SERVER_PORT}") as ws1, \
               websockets.connect(f"ws://localhost:{SERVER_PORT}") as ws2:
        test_message = json.dumps({"user": "test", "message": "hello"})
        await ws1.send(test_message)
        response = await asyncio.wait_for(ws2.recv(), timeout=5.0)
        data = json.loads(response)
        assert data["message"] == "hello"
        assert data["user"] == "test"
@pytest.mark.asyncio
async def test_multiple_clients(chat_server):
    """Vérifie le broadcast vers plusieurs clients."""
    clients = []
    for _ in range(5):
        ws = await websockets.connect(f"ws://localhost:{SERVER_PORT}")
        clients.append(ws)
    # Le premier client envoie un message
    await clients[0].send(json.dumps({"message": "broadcast"}))
    # Tous les autres doivent le recevoir
    for client in clients[1:]:
        msg = await asyncio.wait_for(client.recv(), timeout=5.0)
        assert json.loads(msg)["message"] == "broadcast"
    for client in clients:
        await client.close()
@pytest.mark.asyncio
async def test_disconnection_handling(chat_server):
    """Vérifie la gestion propre des déconnexions."""
    async with websockets.connect(f"ws://localhost:{SERVER_PORT}") as ws1:
        ws2 = await websockets.connect(f"ws://localhost:{SERVER_PORT}")
        await ws2.close()
        # ws1 doit pouvoir continuer à fonctionner
        await ws1.send(json.dumps({"message": "still alive"}))
        # Pas d'exception = test réussi
```
**Exécution des tests :**

```
$ pytest test_chat.py -v
========================= test session starts ==========================
collected 3 items
test_chat.py::test_echo_message PASSED                           [ 33%]
test_chat.py::test_multiple_clients PASSED                       [ 66%]
test_chat.py::test_disconnection_handling PASSED                 [100%]
========================= 3 passed in 1.24s ===========================
```
Utilisez `asyncio.wait_for()` avec un timeout dans chaque test pour éviter que les tests ne bloquent indéfiniment si un message n'est jamais reçu. Le port dédié `8799` empêche les conflits avec un serveur de développement qui tournerait sur le port par défaut. La fixture `chat_server` assure le nettoyage complet avec `server.close()` et `await server.wait_closed()`.

## Étape 12 : Déploiement en Production avec TLS

Le déploiement en production d'une application WebSocket diffère significativement du développement local. Trois éléments sont essentiels : le chiffrement TLS (wss://), un reverse proxy comme Nginx pour gérer les connexions, et un processus supervisé (systemd ou Docker) pour la fiabilité. Le protocole `wss://` est obligatoire sur les sites HTTPS — les navigateurs bloquent les connexions `ws://` non chiffrées depuis une page HTTPS.

```
# production_server.py - Serveur WebSocket prêt pour la production
import asyncio
import websockets
import ssl
import json
import logging
import signal
# Configuration du logging
logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s"
)
logger = logging.getLogger("ws-server")
CLIENTS = set()
MAX_CLIENTS = 1000
MAX_MESSAGE_SIZE = 1_048_576  # 1 Mo
async def handler(websocket):
    if len(CLIENTS) >= MAX_CLIENTS:
        await websocket.close(4003, "Serveur saturé")
        logger.warning("Connexion refusée : limite atteinte")
        return
    CLIENTS.add(websocket)
    remote = websocket.remote_address
    logger.info(f"Connexion : {remote} (total: {len(CLIENTS)})")
    try:
        async for message in websocket:
            if len(message) > MAX_MESSAGE_SIZE:
                await websocket.close(4004, "Message trop volumineux")
                return
            data = json.loads(message)
            # Sanitisation basique
            data["message"] = data.get("message", "")[:2000]
            broadcast_msg = json.dumps(data)
            await asyncio.gather(
                *[c.send(broadcast_msg) for c in CLIENTS if c != websocket],
                return_exceptions=True
            )
    except websockets.exceptions.ConnectionClosed as e:
        logger.info(f"Déconnexion {remote}: {e.code}")
    except json.JSONDecodeError:
        await websocket.close(4005, "JSON invalide")
    finally:
        CLIENTS.discard(websocket)
        logger.info(f"Nettoyage {remote} (total: {len(CLIENTS)})")
async def main():
    # Configuration TLS (décommenter en production)
    # ssl_context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
    # ssl_context.load_cert_chain("cert.pem", "key.pem")
    stop = asyncio.get_event_loop().create_future()
    # Gestion propre de l'arrêt
    loop = asyncio.get_event_loop()
    for sig in (signal.SIGINT, signal.SIGTERM):
        loop.add_signal_handler(sig, stop.set_result, None)
    async with websockets.serve(
        handler,
        "0.0.0.0",
        8765,
        # ssl=ssl_context,  # Décommenter en production
        max_size=MAX_MESSAGE_SIZE,
        max_queue=32,
        ping_interval=30,
        ping_timeout=10,
    ) as server:
        logger.info("Serveur production démarré sur :8765")
        await stop  # Attendre signal d'arrêt
    logger.info("Serveur arrêté proprement")
if __name__ == "__main__":
    asyncio.run(main())
```
La configuration Nginx suivante gère le reverse proxy WebSocket avec TLS :

```
# /etc/nginx/sites-available/websocket
upstream websocket_backend {
    server 127.0.0.1:8765;
}
server {
    listen 443 ssl;
    server_name chat.example.com;
    ssl_certificate /etc/letsencrypt/live/chat.example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/chat.example.com/privkey.pem;
    location /ws/ {
        proxy_pass http://websocket_backend;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_read_timeout 86400s;  # 24h pour les connexions longues
        proxy_send_timeout 86400s;
    }
}
```
Les en-têtes `Upgrade` et `Connection` sont obligatoires pour que Nginx transmette correctement le handshake WebSocket. Le `proxy_read_timeout` de 24 heures empêche Nginx de couper les connexions WebSocket inactives. En production, ajoutez aussi un `limit_conn` Nginx pour limiter les connexions par IP.

## 5 Erreurs Courantes et Comment les Éviter

Après avoir construit le projet complet, voici les cinq erreurs les plus fréquentes que rencontrent les développeurs lorsqu'ils travaillent avec les WebSockets en Python, ainsi que leurs solutions.

**Erreur 1 : Oublier `await` sur les appels WebSocket.** Toutes les opérations WebSocket (`send`, `recv`, `close`) sont des coroutines. Écrire `websocket.send(msg)` sans `await` ne génère pas d'erreur visible en Python 3.12, mais le message ne sera jamais envoyé. Ajoutez toujours `await` devant chaque appel.

