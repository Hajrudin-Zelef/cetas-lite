---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-3
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["diffusion"]
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [191, 348]
sha256: 435016b65e1353403e03412fecec717c542083e9fa490f4b932b16260a43ca9a
---

# server.py - Serveur WebSocket basique

```
<!-- index.html - Client de chat WebSocket -->
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>Chat WebSocket</title>
    <style>
        body { font-family: system-ui, sans-serif; max-width: 600px; margin: 2rem auto; }
        #messages { height: 400px; overflow-y: auto; border: 1px solid #ccc; padding: 1rem; }
        .system { color: #666; font-style: italic; }
        .message { margin: 0.5rem 0; }
        input, button { padding: 0.5rem; margin-top: 0.5rem; }
        input[type="text"] { width: 70%; }
    </style>
</head>
<body>
    <h1>Chat WebSocket en Temps Réel</h1>
    <div id="messages"></div>
    <input type="text" id="username" placeholder="Votre pseudo" />
    <br/>
    <input type="text" id="input" placeholder="Tapez votre message..." />
    <button onclick="sendMessage()">Envoyer</button>
    <script>
        const ws = new WebSocket("ws://localhost:8765");
        const messages = document.getElementById("messages");
        ws.onopen = () => {
            addMessage("Connecté au serveur", "system");
        };
        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);
            if (data.type === "system") {
                addMessage(data.message, "system");
            } else {
                addMessage(`${data.username}: ${data.message}`, "message");
            }
        };
        ws.onclose = () => {
            addMessage("Déconnecté du serveur", "system");
        };
        function addMessage(text, className) {
            const div = document.createElement("div");
            div.className = className;
            div.textContent = text;
            messages.appendChild(div);
            messages.scrollTop = messages.scrollHeight;
        }
        function sendMessage() {
            const username = document.getElementById("username").value || "Anonyme";
            const input = document.getElementById("input");
            if (input.value) {
                ws.send(JSON.stringify({ username, message: input.value }));
                input.value = "";
            }
        }
        document.getElementById("input").addEventListener("keypress", (e) => {
            if (e.key === "Enter") sendMessage();
        });
    </script>
</body>
</html>
```
Ouvrez ce fichier dans deux onglets de navigateur différents. Chaque message envoyé depuis un onglet apparaît instantanément dans l’autre. Le serveur gère automatiquement la sérialisation JSON et la diffusion. Vous pouvez tester avec trois, quatre ou dix onglets : le mécanisme de broadcast reste identique et les performances restent stables grâce à l’architecture asynchrone.

## Étape 6 : Salons de Chat et Routage des Messages

Une application de chat professionnelle nécessite des salons (rooms) pour organiser les conversations par sujet. Au lieu d’un seul ensemble global de clients, nous allons utiliser un dictionnaire de salons où chaque clé est le nom du salon et chaque valeur est l’ensemble des connexions WebSocket qui y participent. Ce pattern est utilisé par toutes les plateformes de messagerie, de Slack à Discord.

```
# rooms_server.py - Serveur WebSocket avec salons
import asyncio
import websockets
import json
from datetime import datetime
from collections import defaultdict
# Dictionnaire des salons : {room_name: set(websockets)}
ROOMS = defaultdict(set)
# Mapping inverse : {websocket: room_name}
USER_ROOMS = {}
# Pseudos des utilisateurs
USER_NAMES = {}
async def broadcast_to_room(room, message):
    """Diffuse un message à tous les clients d'un salon."""
    clients = ROOMS.get(room, set())
    if clients:
        await asyncio.gather(
            *[client.send(message) for client in clients],
            return_exceptions=True
        )
async def handler(websocket):
    """Gère une connexion avec routage par salon."""
    current_room = None
    username = "Anonyme"
    try:
        async for raw_message in websocket:
            data = json.loads(raw_message)
            action = data.get("action")
            if action == "join":
                # Quitter l'ancien salon si nécessaire
                if current_room and websocket in ROOMS[current_room]:
                    ROOMS[current_room].discard(websocket)
                    await broadcast_to_room(current_room, json.dumps({
                        "type": "system",
                        "message": f"{username} a quitté le salon",
                        "room": current_room
                    }))
                # Rejoindre le nouveau salon
                current_room = data.get("room", "général")
                username = data.get("username", "Anonyme")
                ROOMS[current_room].add(websocket)
                USER_ROOMS[websocket] = current_room
                USER_NAMES[websocket] = username
                # Notifier le salon
                count = len(ROOMS[current_room])
                await broadcast_to_room(current_room, json.dumps({
                    "type": "system",
                    "message": f"{username} a rejoint #{current_room} ({count} membre{'s' if count > 1 else ''})",
                    "room": current_room
                }))
                # Envoyer la liste des salons disponibles
                room_list = {name: len(members) for name, members in ROOMS.items() if members}
                await websocket.send(json.dumps({
                    "type": "room_list",
                    "rooms": room_list
                }))
            elif action == "message" and current_room:
                await broadcast_to_room(current_room, json.dumps({
                    "type": "message",
                    "username": username,
                    "message": data.get("message", ""),
                    "room": current_room,
                    "timestamp": datetime.now().isoformat()
                }))
    except websockets.exceptions.ConnectionClosed:
        pass
    finally:
        if current_room:
            ROOMS[current_room].discard(websocket)
            await broadcast_to_room(current_room, json.dumps({
                "type": "system",
                "message": f"{username} a quitté le salon",
                "room": current_room
            }))
        USER_ROOMS.pop(websocket, None)
        USER_NAMES.pop(websocket, None)
async def main():
    async with websockets.serve(handler, "0.0.0.0", 8765) as server:
        print("Serveur avec salons démarré sur ws://0.0.0.0:8765")
        await server.serve_forever()
if __name__ == "__main__":
    asyncio.run(main())
```
Ce serveur gère proprement les transitions entre salons : quand un utilisateur rejoint un nouveau salon, il est automatiquement retiré de l’ancien. Le dictionnaire `defaultdict(set)` simplifie la gestion des salons en créant automatiquement un nouvel ensemble pour chaque nouveau nom de salon. La liste des salons actifs est envoyée au client à chaque connexion, permettant une interface dynamique.

## Étape 7 : Authentification et Sécurité WebSocket

La sécurité est critique pour toute application WebSocket en production. Contrairement aux requêtes HTTP classiques où chaque requête peut porter un cookie ou un token, la connexion WebSocket ne dispose que du handshake initial pour l’authentification. Deux approches sont couramment utilisées : les tokens dans les paramètres de requête URL et les en-têtes personnalisés lors du handshake. La première est plus simple, la seconde est plus sécurisée.

