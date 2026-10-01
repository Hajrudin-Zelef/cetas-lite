---
id: collect-261001-general-networking/general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes-4
title: "server.py - Serveur WebSocket basique"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/tutoriel-websocket-python-2026-chat-en-12-etapes.md
source_anchor: ""
source_lines: [349, 508]
sha256: 57c231d4828c82d3f934ba7199380b5a61e612c6ef410b4e4884c5ba39761b99
---

# server.py - Serveur WebSocket basique

```
# auth_server.py - Serveur WebSocket avec authentification
import asyncio
import websockets
import json
import hashlib
import secrets
from datetime import datetime, timedelta
# Base de données simplifiée d'utilisateurs
USERS_DB = {
    "alice": hashlib.sha256("motdepasse123".encode()).hexdigest(),
    "bob": hashlib.sha256("secret456".encode()).hexdigest(),
}
# Tokens actifs : {token: {"username": str, "expires": datetime}}
ACTIVE_TOKENS = {}
def generate_token(username):
    """Génère un token d'authentification."""
    token = secrets.token_urlsafe(32)
    ACTIVE_TOKENS[token] = {
        "username": username,
        "expires": datetime.now() + timedelta(hours=24)
    }
    return token
def validate_token(token):
    """Vérifie la validité d'un token."""
    if token not in ACTIVE_TOKENS:
        return None
    data = ACTIVE_TOKENS[token]
    if datetime.now() > data["expires"]:
        del ACTIVE_TOKENS[token]
        return None
    return data["username"]
async def auth_handler(websocket):
    """Gère l'authentification avant d'autoriser le chat."""
    try:
        # Attendre le message d'authentification (timeout 10s)
        raw = await asyncio.wait_for(websocket.recv(), timeout=10.0)
        data = json.loads(raw)
        if data.get("action") == "login":
            username = data.get("username", "")
            password_hash = hashlib.sha256(
                data.get("password", "").encode()
            ).hexdigest()
            if username in USERS_DB and USERS_DB[username] == password_hash:
                token = generate_token(username)
                await websocket.send(json.dumps({
                    "type": "auth",
                    "status": "success",
                    "token": token,
                    "username": username
                }))
                # Authentifié : passer au handler de chat
                await chat_handler(websocket, username)
            else:
                await websocket.send(json.dumps({
                    "type": "auth",
                    "status": "error",
                    "message": "Identifiants invalides"
                }))
                await websocket.close(4001, "Authentification échouée")
        elif data.get("action") == "token_auth":
            username = validate_token(data.get("token", ""))
            if username:
                await websocket.send(json.dumps({
                    "type": "auth",
                    "status": "success",
                    "username": username
                }))
                await chat_handler(websocket, username)
            else:
                await websocket.close(4001, "Token invalide ou expiré")
    except asyncio.TimeoutError:
        await websocket.close(4002, "Timeout d'authentification")
async def chat_handler(websocket, username):
    """Handler de chat post-authentification."""
    print(f"{username} authentifié et connecté")
    async for message in websocket:
        data = json.loads(message)
        data["username"] = username  # Forcer le nom authentifié
        data["timestamp"] = datetime.now().isoformat()
        await websocket.send(json.dumps(data))
async def main():
    async with websockets.serve(auth_handler, "0.0.0.0", 8765) as server:
        print("Serveur sécurisé démarré sur ws://0.0.0.0:8765")
        await server.serve_forever()
if __name__ == "__main__":
    asyncio.run(main())
```
Plusieurs points de sécurité importants dans ce code. Le timeout de 10 secondes sur le premier message empêche les connexions fantômes qui consommeraient des ressources. Le code de fermeture personnalisé `4001` (dans la plage 4000-4999 réservée aux applications) permet au client de distinguer une erreur d’authentification d’une fermeture normale. Le nom d’utilisateur est forcé côté serveur à chaque message pour empêcher l’usurpation d’identité.

### Bonnes Pratiques de Sécurité WebSocket

En production, plusieurs mesures additionnelles sont essentielles. Utilisez toujours `wss://` (WebSocket Secure) avec TLS/SSL, jamais `ws://` en clair. Validez systématiquement l’origine de la connexion via l’en-tête `Origin` dans le handshake pour prévenir les attaques CSRF. Limitez le taux de messages par client (rate limiting) pour empêcher les attaques par saturation. Enfin, définissez une taille maximale de message avec le paramètre `max_size` de la bibliothèque `websockets` pour éviter les attaques par déni de service basées sur la mémoire.

## Étape 8 : Heartbeat et Reconnexion Automatique

Les connexions WebSocket peuvent se fermer silencieusement, notamment à travers les proxys, les pare-feux ou les équilibreurs de charge qui coupent les connexions inactives. Le mécanisme de heartbeat (ping/pong) permet de détecter ces déconnexions et de maintenir la connexion active. La bibliothèque `websockets` gère automatiquement les ping/pong au niveau du protocole, mais il est recommandé d’ajouter un heartbeat applicatif pour les cas où le ping protocole est insuffisant.

```
# heartbeat_client.py - Client avec reconnexion automatique
import asyncio
import websockets
import json
class ResilientWebSocketClient:
    """Client WebSocket avec heartbeat et reconnexion."""
    def __init__(self, uri, reconnect_delay=3, max_retries=10):
        self.uri = uri
        self.reconnect_delay = reconnect_delay
        self.max_retries = max_retries
        self.websocket = None
        self.connected = False
    async def connect(self):
        """Connexion avec reconnexion automatique."""
        retries = 0
        while retries < self.max_retries:
            try:
                self.websocket = await websockets.connect(
                    self.uri,
                    ping_interval=20,    # Ping toutes les 20s
                    ping_timeout=10,     # Timeout ping de 10s
                    close_timeout=5,     # Timeout de fermeture
                    max_size=1_048_576,  # 1 Mo max par message
                )
                self.connected = True
                retries = 0  # Réinitialiser le compteur
                print(f"Connecté à {self.uri}")
                await self.listen()
            except websockets.exceptions.ConnectionClosed as e:
                self.connected = False
                print(f"Connexion fermée : {e.code} - Reconnexion...")
            except (ConnectionRefusedError, OSError) as e:
                self.connected = False
                retries += 1
                delay = min(self.reconnect_delay * (2 ** retries), 60)
                print(f"Erreur de connexion ({retries}/{self.max_retries}). "
                      f"Nouvelle tentative dans {delay}s...")
                await asyncio.sleep(delay)
        print("Nombre maximum de tentatives atteint. Arrêt.")
    async def listen(self):
        """Écoute les messages entrants."""
        async for message in self.websocket:
            data = json.loads(message)
            await self.on_message(data)
    async def on_message(self, data):
        """Callback pour les messages reçus."""
        print(f"Message : {data}")
    async def send(self, data):
        """Envoie un message si connecté."""
        if self.connected and self.websocket:
            await self.websocket.send(json.dumps(data))
async def main():
    client = ResilientWebSocketClient("ws://localhost:8765")
    await client.connect()
if __name__ == "__main__":
    asyncio.run(main())
```
Ce client implémente le backoff exponentiel pour les reconnexions : le délai double à chaque tentative (3s, 6s, 12s, 24s...) jusqu'à un maximum de 60 secondes. Les paramètres `ping_interval=20` et `ping_timeout=10` de `websockets.connect()` activent le heartbeat automatique du protocole. Si le serveur ne répond pas au ping dans les 10 secondes, la connexion est considérée comme morte et le mécanisme de reconnexion se déclenche.

## Étape 9 : Intégration avec FastAPI

