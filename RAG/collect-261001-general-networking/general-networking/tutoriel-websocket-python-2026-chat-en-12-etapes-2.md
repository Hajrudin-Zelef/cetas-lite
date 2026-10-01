---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-2
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [52, 190]
sha256: 88e1d23ae4028b68ad1512a3fca333ea4d5b36abddca381fdbd7ddfdf2080378
---

# server.py - Serveur WebSocket basique

```
# server.py - Serveur WebSocket basique
import asyncio
import websockets
async def handler(websocket):
    """Gère chaque connexion WebSocket entrante."""
    print(f"Nouvelle connexion : {websocket.remote_address}")
    try:
        async for message in websocket:
            print(f"Reçu : {message}")
            # Renvoyer le message en écho
            await websocket.send(f"Écho : {message}")
    except websockets.exceptions.ConnectionClosed as e:
        print(f"Connexion fermée : {e.code} {e.reason}")
async def main():
    # Démarrer le serveur sur localhost:8765
    async with websockets.serve(handler, "localhost", 8765) as server:
        print("Serveur WebSocket démarré sur ws://localhost:8765")
        await server.serve_forever()
if __name__ == "__main__":
    asyncio.run(main())
```
Lancez le serveur avec `python3 server.py`. Vous verrez le message `Serveur WebSocket démarré sur ws://localhost:8765`. Le serveur reste actif et attend des connexions. La boucle `async for message in websocket` itère automatiquement sur les messages entrants et gère proprement la fermeture de connexion lorsque le client se déconnecte.

**Sortie attendue du serveur :**

```
$ python3 server.py
Serveur WebSocket démarré sur ws://localhost:8765
Nouvelle connexion : ('127.0.0.1', 54321)
Reçu : Bonjour le monde
Connexion fermée : 1000 
```
## Étape 3 : Construire un Client WebSocket Python

Maintenant créons un client Python qui se connecte au serveur. Le client utilise la même bibliothèque `websockets` et le pattern `async with` qui gère automatiquement l’ouverture et la fermeture de la connexion. La connexion est établie en moins de 10 millisecondes en local, démontrant l’avantage des WebSockets pour les communications en temps réel.

```
# client.py - Client WebSocket Python
import asyncio
import websockets
async def client():
    uri = "ws://localhost:8765"
    async with websockets.connect(uri) as websocket:
        # Envoyer un message
        message = "Bonjour le monde"
        await websocket.send(message)
        print(f"Envoyé : {message}")
        # Recevoir la réponse
        response = await websocket.recv()
        print(f"Reçu : {response}")
        # Envoyer plusieurs messages
        for i in range(3):
            msg = f"Message numéro {i + 1}"
            await websocket.send(msg)
            resp = await websocket.recv()
            print(f"Envoyé : {msg} → Reçu : {resp}")
if __name__ == "__main__":
    asyncio.run(client())
```
**Sortie attendue du client :**

```
$ python3 client.py
Envoyé : Bonjour le monde
Reçu : Écho : Bonjour le monde
Envoyé : Message numéro 1 → Reçu : Écho : Message numéro 1
Envoyé : Message numéro 2 → Reçu : Écho : Message numéro 2
Envoyé : Message numéro 3 → Reçu : Écho : Message numéro 3
```
Le pattern `async with websockets.connect(uri)` est fondamental : il ouvre la connexion, exécute le bloc de code, puis ferme proprement la connexion avec le code de fermeture 1000 (Normal Closure). Si une erreur survient, la connexion est quand même fermée grâce au gestionnaire de contexte.

## Étape 4 : Système de Broadcast Multi-Clients

Un echo server est utile pour apprendre, mais une application de chat réelle doit diffuser les messages à tous les clients connectés. C’est le pattern de broadcast. Nous allons maintenir un ensemble (`set`) de toutes les connexions actives et envoyer chaque message reçu à tous les autres participants. Cette architecture est la base de toute application temps réel multi-utilisateurs.

```
# chat_server.py - Serveur de chat avec broadcast
import asyncio
import websockets
import json
from datetime import datetime
# Ensemble de toutes les connexions actives
CLIENTS = set()
async def register(websocket):
    """Enregistre un nouveau client."""
    CLIENTS.add(websocket)
    count = len(CLIENTS)
    print(f"Client connecté. Total : {count}")
    # Notifier tous les clients du nombre de connectés
    await broadcast(json.dumps({
        "type": "system",
        "message": f"Un utilisateur a rejoint le chat ({count} connecté{'s' if count > 1 else ''})",
        "timestamp": datetime.now().isoformat()
    }))
async def unregister(websocket):
    """Désenregistre un client déconnecté."""
    CLIENTS.discard(websocket)
    count = len(CLIENTS)
    print(f"Client déconnecté. Total : {count}")
    await broadcast(json.dumps({
        "type": "system",
        "message": f"Un utilisateur a quitté le chat ({count} connecté{'s' if count > 1 else ''})",
        "timestamp": datetime.now().isoformat()
    }))
async def broadcast(message):
    """Diffuse un message à tous les clients connectés."""
    if CLIENTS:
        await asyncio.gather(
            *[client.send(message) for client in CLIENTS],
            return_exceptions=True
        )
async def handler(websocket):
    """Gère une connexion client."""
    await register(websocket)
    try:
        async for raw_message in websocket:
            data = json.loads(raw_message)
            # Ajouter un timestamp serveur
            data["timestamp"] = datetime.now().isoformat()
            data["type"] = "message"
            await broadcast(json.dumps(data))
    except websockets.exceptions.ConnectionClosed:
        pass
    finally:
        await unregister(websocket)
async def main():
    async with websockets.serve(handler, "0.0.0.0", 8765) as server:
        print("Serveur de chat démarré sur ws://0.0.0.0:8765")
        await server.serve_forever()
if __name__ == "__main__":
    asyncio.run(main())
```
Le serveur utilise `asyncio.gather` avec `return_exceptions=True` pour envoyer le message à tous les clients simultanément sans bloquer si l’un d’eux est lent. L’utilisation de `CLIENTS.discard()` plutôt que `CLIENTS.remove()` évite une `KeyError` si le client a déjà été retiré. L’écoute sur `0.0.0.0` au lieu de `localhost` permet d’accepter les connexions depuis d’autres machines sur le réseau.

## Étape 5 : Client HTML/JavaScript pour le Chat

Pour tester le serveur de chat, créons une interface web simple en HTML et JavaScript. L’API WebSocket est intégrée nativement dans tous les navigateurs modernes depuis plus de dix ans. Le constructeur `new WebSocket(url)` établit la connexion, et les événements `onopen`, `onmessage`, `onclose` et `onerror` permettent de réagir aux différents états de la connexion.

