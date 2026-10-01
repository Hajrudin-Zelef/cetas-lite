---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-5
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [509, 626]
sha256: 2d5ec9288aef6b9ffbeec57b30ee73c4e25d12e860848f0450b35d5cb55bf9ed
---

# server.py - Serveur WebSocket basique

Pour les applications web complètes, FastAPI offre un support WebSocket natif intégré directement dans son framework — un guide publié le 23 mars 2026 par **Websocket.org** confirme que ce support repose sur la couche ASGI de **Starlette**, le framework sous-jacent de FastAPI. L'avantage de FastAPI est de combiner les endpoints HTTP REST classiques avec les endpoints WebSocket dans la même application, partageant les mêmes modèles Pydantic, le même système de dépendances et la même documentation OpenAPI.

```
# fastapi_ws.py - WebSocket avec FastAPI
from fastapi import FastAPI, WebSocket, WebSocketDisconnect
from fastapi.responses import HTMLResponse
import json
from datetime import datetime
app = FastAPI(title="Chat WebSocket FastAPI")
class ConnectionManager:
    """Gestionnaire centralisé des connexions WebSocket."""
    def __init__(self):
        self.active_connections: dict[str, list[WebSocket]] = {}
    async def connect(self, websocket: WebSocket, room: str):
        await websocket.accept()
        if room not in self.active_connections:
            self.active_connections[room] = []
        self.active_connections[room].append(websocket)
    def disconnect(self, websocket: WebSocket, room: str):
        if room in self.active_connections:
            self.active_connections[room].remove(websocket)
            if not self.active_connections[room]:
                del self.active_connections[room]
    async def broadcast(self, message: dict, room: str):
        for connection in self.active_connections.get(room, []):
            await connection.send_json(message)
manager = ConnectionManager()
@app.websocket("/ws/{room}/{username}")
async def websocket_endpoint(websocket: WebSocket, room: str, username: str):
    await manager.connect(websocket, room)
    await manager.broadcast(
        {"type": "system", "message": f"{username} a rejoint #{room}"},
        room
    )
    try:
        while True:
            data = await websocket.receive_text()
            message = {
                "type": "message",
                "username": username,
                "message": data,
                "room": room,
                "timestamp": datetime.now().isoformat()
            }
            await manager.broadcast(message, room)
    except WebSocketDisconnect:
        manager.disconnect(websocket, room)
        await manager.broadcast(
            {"type": "system", "message": f"{username} a quitté #{room}"},
            room
        )
@app.get("/api/rooms")
async def get_rooms():
    """Endpoint REST pour lister les salons actifs."""
    return {
        room: len(clients)
        for room, clients in manager.active_connections.items()
    }
# Lancer avec : uvicorn fastapi_ws:app --host 0.0.0.0 --port 8000
```
Lancez ce serveur avec `uvicorn fastapi_ws:app --host 0.0.0.0 --port 8000`. La connexion WebSocket se fait sur `ws://localhost:8000/ws/général/alice`. L'endpoint REST `/api/rooms` retourne la liste des salons actifs au format JSON, démontrant comment les deux protocoles cohabitent. FastAPI gère automatiquement la validation des paramètres de chemin et la documentation interactive à `/docs`.

## Étape 10 : Envoi de Fichiers et Données Binaires

Les WebSockets ne se limitent pas au texte. Le protocole supporte nativement les messages binaires, ce qui permet d'envoyer des images, des fichiers audio ou tout autre type de données. En Python, la distinction entre message texte et binaire se fait automatiquement : `str` est envoyé comme texte, `bytes` comme binaire. Cette fonctionnalité est essentielle pour les tableaux de bord qui affichent des graphiques en temps réel ou les outils collaboratifs qui partagent des fichiers.

```
# binary_server.py - Gestion des messages binaires
import asyncio
import websockets
import json
import base64
CLIENTS = set()
async def handler(websocket):
    CLIENTS.add(websocket)
    try:
        async for message in websocket:
            if isinstance(message, bytes):
                # Message binaire : fichier reçu
                size_kb = len(message) / 1024
                print(f"Fichier reçu : {size_kb:.1f} Ko")
                # Métadonnées en JSON, contenu en binaire
                # Convention : les 4 premiers octets = taille des métadonnées
                meta_size = int.from_bytes(message[:4], byteorder="big")
                meta_json = json.loads(message[4:4 + meta_size].decode())
                file_data = message[4 + meta_size:]
                print(f"Fichier : {meta_json['filename']} "
                      f"({meta_json['mime_type']})")
                # Diffuser à tous les clients
                for client in CLIENTS:
                    if client != websocket:
                        await client.send(message)
            else:
                # Message texte : chat normal
                data = json.loads(message)
                for client in CLIENTS:
                    if client != websocket:
                        await client.send(json.dumps(data))
    except websockets.exceptions.ConnectionClosed:
        pass
    finally:
        CLIENTS.discard(websocket)
async def main():
    async with websockets.serve(
        handler, "0.0.0.0", 8765,
        max_size=10_485_760  # 10 Mo max par message
    ) as server:
        print("Serveur binaire démarré (max 10 Mo/message)")
        await server.serve_forever()
if __name__ == "__main__":
    asyncio.run(main())
```
Le paramètre `max_size=10_485_760` limite la taille des messages à 10 Mo pour éviter les abus mémoire. Notre convention de protocole utilise les 4 premiers octets pour encoder la taille des métadonnées JSON, suivies du contenu binaire du fichier. Cette approche est plus efficace que l'encodage base64 qui augmente la taille des données de 33 %.

## Étape 11 : Tests Unitaires et d'Intégration

Tester des applications WebSocket nécessite des outils spécifiques. La bibliothèque `websockets` fournit un serveur de test intégré, et `pytest-asyncio` permet d'écrire des tests asynchrones. Chaque test démarre un serveur éphémère, exécute les assertions, puis ferme proprement toutes les connexions. Cette approche garantit l'isolation entre les tests et évite les effets de bord.

