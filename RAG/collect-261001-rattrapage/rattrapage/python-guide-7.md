---
id: collect-261001-rattrapage/rattrapage/python-guide-7
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["sol"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [1900, 2214]
sha256: d0130b526dedacfd046b3c956ba5f6d4ad5b842fbfd7f5e942b0b12b4e2ee343
---

# Guide Python complet — De l'installation aux scripts sysadmin

```python
# Chaîne → datetime
dt = datetime.strptime("2026-09-26 10:30:00", "%Y-%m-%d %H:%M:%S")

# datetime → chaîne
dt.strftime("%d/%m/%Y à %Hh%M")     # '26/09/2026 à 10h30'
dt.isoformat()                       # '2026-09-26T10:30:00'

# Python 3.11+ : parse l'ISO directement
datetime.fromisoformat("2026-09-26T10:30:00+00:00")
```

### Codes de format courants

| Code | Sens | Exemple |
|---|---|---|
| `%Y` | année 4 chiffres | 2026 |
| `%m` | mois | 09 |
| `%d` | jour | 26 |
| `%H` | heure (24h) | 10 |
| `%M` | minute | 30 |
| `%S` | seconde | 00 |

### Calculs avec timedelta

```python
expiration = datetime.now(timezone.utc) + timedelta(days=90)
il_y_a_1h = datetime.now(timezone.utc) - timedelta(hours=1)

duree = expiration - il_y_a_1h
print(duree.days, duree.seconds)

# Timestamp Unix
ts = datetime.now(timezone.utc).timestamp()     # float
dt = datetime.fromtimestamp(ts, tz=timezone.utc)
```

### Exemple : rotation de logs par date

```python
from datetime import date
nom = f"sauvegarde-{date.today().isoformat()}.tar.gz"  # sauvegarde-2026-09-26.tar.gz
```

> **Règle d'or :** stockez/manipulez en UTC (`timezone.utc`), convertissez en heure locale uniquement pour l'affichage.

---

## 40. argparse : scripts en ligne de commande

```python
#!/usr/bin/env python3
"""Ping une liste d'hôtes."""
import argparse

def main():
    p = argparse.ArgumentParser(
        description="Ping des hôtes et affiche un résumé.",
        epilog="Exemple : %(prog)s -t 2 srv-01 srv-02",
    )
    p.add_argument("hosts", nargs="+", help="hôtes à pinger")
    p.add_argument("-t", "--timeout", type=float, default=2.0,
                   help="timeout en secondes (défaut : %(default)s)")
    p.add_argument("-v", "--verbose", action="store_true",
                   help="mode verbeux")
    p.add_argument("--fichier", "-f", type=argparse.FileType("r"),
                   help="fichier contenant des hôtes (un par ligne)")
    args = p.parse_args()

    if args.verbose:
        print(f"Timeout : {args.timeout}s")
    for h in args.hosts:
        print(f"Ping {h}...")

if __name__ == "__main__":
    main()
```

```bash
python3 ping.py srv-01 srv-02 -t 1 -v
python3 ping.py --help     # aide générée automatiquement
```

### Types d'arguments

| Paramètre | Effet |
|---|---|
| `nargs="+"` | 1 ou + valeurs (liste) |
| `nargs="*"` | 0 ou + valeurs |
| `action="store_true"` | flag booléen |
| `type=int` | conversion auto (+ erreur claire) |
| `choices=["a","b"]` | valeurs autorisées |
| `required=True` | option obligatoire |
| `default=...` | valeur par défaut |

---

## 41. logging : journaliser proprement

Oubliez `print()` pour les scripts sérieux : `logging` gère niveaux, fichiers, formats.

```python
import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
    handlers=[
        logging.FileHandler("script.log", encoding="utf-8"),
        logging.StreamHandler(),          # console aussi
    ],
)

logging.debug("détail technique")
logging.info("Sauvegarde démarrée")
logging.warning("Espace disque < 20%")
logging.error("Échec de connexion à srv-01")
logging.critical("Base de données inaccessible")
```

### Niveaux

| Niveau | Valeur | Usage |
|---|---|---|
| DEBUG | 10 | détails de développement |
| INFO | 20 | déroulement normal |
| WARNING | 30 | anomalie non bloquante |
| ERROR | 40 | échec d'une opération |
| CRITICAL | 50 | panne grave |

### Logger nommé (recommandé dans les modules)

```python
logger = logging.getLogger(__name__)

def traiter():
    logger.info("Traitement de %s", nom)   # lazy % : pas de formatage si niveau filtré
```

### Rotation automatique des logs

```python
from logging.handlers import RotatingFileHandler

handler = RotatingFileHandler(
    "app.log", maxBytes=5_000_000, backupCount=5, encoding="utf-8"
)  # app.log, app.log.1 ... app.log.5
```

> **Anti-piège :** ne jamais logger de secrets (mots de passe, tokens). Si une URL contient un token, masquez-le avant de logger.

---

## 42. requests : HTTP(S) simplifié

```bash
pip install requests
```

```python
import requests

# GET simple
r = requests.get("https://api.exemple.com/equipements", timeout=10)
r.status_code          # 200
r.json()               # dict/list si la réponse est du JSON
r.text                 # corps en texte

# Avec paramètres et en-têtes
r = requests.get(
    "https://api.exemple.com/equipements",
    params={"site": "paris"},
    headers={"Authorization": "Bearer VOTRE_TOKEN_ICI"},
    timeout=10,
)

# POST JSON
r = requests.post(
    "https://api.exemple.com/equipements",
    json={"hostname": "srv-99", "ip": "10.0.0.99"},
    timeout=10,
)
r.raise_for_status()   # lève une exception si code 4xx/5xx
```

### Toujours `timeout=` et `raise_for_status()`

Sans timeout, une API qui ne répond pas bloque votre script **indéfiniment**.

### Sessions : connexions réutilisées + auth par défaut

```python
s = requests.Session()
s.headers.update({"Authorization": "Bearer VOTRE_TOKEN_ICI"})
s.get("https://api.exemple.com/a", timeout=10)
s.get("https://api.exemple.com/b", timeout=10)   # connexion réutilisée
```

### Gestion d'erreurs réseau

```python
try:
    r = requests.get(url, timeout=10)
    r.raise_for_status()
except requests.exceptions.Timeout:
    print("Délai dépassé")
except requests.exceptions.ConnectionError:
    print("Connexion impossible")
except requests.exceptions.HTTPError as e:
    print(f"Erreur HTTP : {e}")
```

### Télécharger un fichier

```python
with requests.get(url, stream=True, timeout=30) as r:
    r.raise_for_status()
    with open("fichier.iso", "wb") as f:
        for chunk in r.iter_content(chunk_size=8192):
            f.write(chunk)
```

---

## 43. Réseau : sockets, ping, inventaire

### Tester un port TCP (sans dépendance externe)

```python
import socket

def port_ouvert(host, port, timeout=2.0):
    """True si le port TCP accepte une connexion."""
    try:
        with socket.create_connection((host, port), timeout=timeout):
            return True
    except OSError:
        return False

port_ouvert("10.0.0.11", 22)    # True/False
port_ouvert("10.0.0.11", 3389)
```

### Résolution DNS

```python
import socket
socket.gethostbyname("srv-web-01")            # '10.0.0.11'
socket.gethostbyaddr("10.0.0.11")             # (nom, alias, ips)
socket.getaddrinfo("srv-01", 22)              # IPv4 + IPv6
```

### Petit serveur TCP (exemple pédagogique)

```python
import socket

with socket.socket() as srv:
    srv.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    srv.bind(("127.0.0.1", 9999))
    srv.listen(1)
    print("En écoute sur 9999...")
    conn, addr = srv.accept()
    with conn:
        print("Client :", addr)
        conn.sendall(b"OK\n")
```

### Inventaire réseau : balayer un /24 sur le port 22

```python
import ipaddress
from concurrent.futures import ThreadPoolExecutor
import socket

def ssh_ouvert(ip, timeout=1.0):
    try:
        with socket.create_connection((str(ip), 22), timeout=timeout):
            return str(ip)
    except OSError:
        return None

reseau = ipaddress.ip_network("10.0.0.0/24")
with ThreadPoolExecutor(max_workers=50) as pool:
    trouves = [ip for ip in pool.map(ssh_ouvert, reseau.hosts()) if ip]

print(f"{len(trouves)} hôtes avec SSH ouvert :")
print("\n".join(trouves))
```

> `ipaddress` (stdlib) : manipule IPv4/IPv6 proprement — `ip_network`, `ip_address`, tests d'appartenance (`ip in reseau`).

---

## 44. sqlite3 : base de données embarquée

Zéro serveur, un simple fichier. Parfait pour un inventaire local.

```python
import sqlite3

conn = sqlite3.connect("parc.db")
conn.row_factory = sqlite3.Row      # accès par nom de colonne
cur = conn.cursor()

cur.execute("""
    CREATE TABLE IF NOT EXISTS equipements (
        id INTEGER PRIMARY KEY,
        hostname TEXT UNIQUE NOT NULL,
        ip TEXT NOT NULL,
        type TEXT,
        site TEXT
    )
""")

