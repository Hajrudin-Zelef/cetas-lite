---
id: collect-261001-rattrapage/rattrapage/python-guide-9
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["packaging"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [2523, 2803]
sha256: 7748145951c314e13fe31cba097e2ca651eac768897e5ef2b421cced5631a747
---

# Guide Python complet — De l'installation aux scripts sysadmin

1. **Algorithme** : une recherche en `set`/`dict` (O(1)) bat une liste (O(n)).
2. **I/O** : parallélisez (threads/asyncio) — c'est là que dorment 90 % des scripts lents.
3. **Éviter le travail inutile** : cache (`lru_cache`), streaming (générateurs).
4. **Micro-optimisations** en dernier : f-strings > `%` > `+`, variables locales en boucle chaude.

```python
# Lent : 'in' sur une liste (O(n) par test)
autorises = ["a", "b", ...]      # 10 000 éléments
if x in autorises: ...

# Rapide : set (O(1))
autorises = {"a", "b", ...}
if x in autorises: ...
```

---

## 50. Packaging : pyproject.toml

Pour distribuer un outil (ou juste l'installer proprement), décrivez le projet.

```
mon_outil/
├── pyproject.toml
├── README.md
└── src/
    └── mon_outil/
        ├── __init__.py
        └── cli.py
```

```toml
[build-system]
requires = ["setuptools>=61"]
build-backend = "setuptools.build_meta"

[project]
name = "mon-outil"
version = "1.0.0"
description = "Inventaire réseau pour sysadmins"
requires-python = ">=3.10"
dependencies = ["requests>=2.28", "pyyaml"]
authors = [{name = "Zelef"}]

[project.scripts]
inventaire = "mon_outil.cli:main"   # crée la commande 'inventaire'

[project.optional-dependencies]
dev = ["pytest", "ruff"]
```

```bash
pip install -e .          # installe en mode éditable (dev)
pip install ".[dev]"      # avec les extras dev
python3 -m build         # construit sdist + wheel (pip install build)
```

> Après `pip install -e .`, la commande `inventaire` est disponible dans le venv. C'est la façon propre de "publier" un script interne.

---

## 51. Intro à FastAPI

FastAPI = framework web moderne pour créer des **API REST** rapidement, avec validation automatique.

```bash
pip install fastapi uvicorn
```

```python
# api.py
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI(title="API Parc informatique")

class Equipement(BaseModel):
    hostname: str
    ip: str
    type: str = "serveur"

PARC: dict[str, Equipement] = {}

@app.get("/equipements")
def lister():
    return list(PARC.values())

@app.post("/equipements", status_code=201)
def ajouter(eq: Equipement):
    PARC[eq.hostname] = eq
    return eq

@app.get("/equipements/{hostname}")
def detail(hostname: str):
    if hostname not in PARC:
        raise HTTPException(status_code=404, detail="Introuvable")
    return PARC[hostname]
```

```bash
uvicorn api:app --reload      # serveur de dev
# Docs auto : http://127.0.0.1:8000/docs
```

### Pourquoi c'est intéressant pour un sysadmin

- Exposer l'inventaire/la supervision via une API propre en 30 lignes.
- Validation Pydantic : les données invalides sont rejetées avec un message clair.
- Documentation OpenAPI générée **automatiquement** (`/docs`).

---

## 52. Bonnes pratiques & PEP 8

### PEP 8 : le style officiel (résumé)

| Règle | Exemple |
|---|---|
| 4 espaces d'indentation | pas de tabulations |
| Lignes ≤ 79-88 caractères | (88 = défaut de `black`) |
| `snake_case` fonctions/variables | `def ping_host():` |
| `PascalCase` classes | `class Serveur:` |
| `MAJUSCULES` constantes | `TIMEOUT = 30` |
| 2 lignes vides entre fonctions top-level | |
| Espaces autour des opérateurs | `x = 1`, pas `x=1` |
| Imports en haut, groupés | stdlib → tiers → locaux |

### Outils automatiques

```bash
pip install ruff black
ruff check .        # lint : erreurs + style
ruff check --fix .  # corrige automatiquement
black .             # reformate le code
```

> `ruff` remplace flake8/pylint en 10x plus rapide. `black` tranche les débats de formatage : on ne discute plus, on lance `black`.

### Principes généraux

1. **Lisibilité > astuce.** Code relu à 3h du matin pendant une astreinte.
2. **Une fonction = une chose.** Si le nom contient "et", découpez.
3. **Noms explicites** : `delai_reconnexion_s` > `d`.
4. **Pas de nombres magiques** : `TIMEOUT_SSH = 10` en constante nommée.
5. **Échouez vite, échouez fort** : validez les entrées en début de fonction.
6. **Ne répétez pas** (DRY) : 3e copier-coller = fonction.
7. **Docstrings** sur tout ce qui est réutilisé.
8. **Gestion d'erreurs explicite** : pas de `except: pass` silencieux.

---

## 53. Sécurité : écrire du code sûr

### 1. Injection de commandes (subprocess)

```python
# DANGEREUX — si host = "x; rm -rf /" :
subprocess.run(f"ping -c1 {host}", shell=True)

# SÛR — liste d'arguments, pas de shell :
subprocess.run(["ping", "-c1", host])
```

### 2. Injection SQL

```python
# DANGEREUX :
cur.execute(f"SELECT * FROM eq WHERE hostname = '{nom}'")

# SÛR — requête paramétrée :
cur.execute("SELECT * FROM eq WHERE hostname = ?", (nom,))
```

### 3. Secrets : jamais en dur dans le code

```python
import os

# Bien : variable d'environnement
token = os.environ.get("API_TOKEN")
if not token:
    raise SystemExit("API_TOKEN non défini")

# Bien : fichier protégé (chmod 600)
token = Path.home().joinpath(".secrets/api_token").read_text().strip()

# MAL : token = "ghp_xxxx" dans le script (finira sur git un jour)
```

> Utilisez un gestionnaire de secrets (Vault, gestionnaire d'équipe) en production. Au minimum : variables d'environnement + `.gitignore`.

### 4. YAML, pickle, eval : l'exécution de code cachée

| À éviter | Alternative sûre |
|---|---|
| `yaml.load(f)` | `yaml.safe_load(f)` |
| `pickle.loads(donnees_inconnues)` | `json` |
| `eval(input_utilisateur)` | `ast.literal_eval()` (littéraux uniquement) |

```python
import ast
ast.literal_eval("{'a': 1}")   # OK : dict
ast.literal_eval("__import__('os')")  # ValueError : refusé
```

### 5. Chemins : path traversal

```python
from pathlib import Path

BASE = Path("/srv/rapports")

def lire_rapport(nom):
    chemin = (BASE / nom).resolve()
    if not chemin.is_relative_to(BASE):   # Python 3.9+
        raise ValueError("Chemin interdit")
    return chemin.read_text()
```

### 6. TLS : ne pas désactiver la vérification

```python
requests.get(url, verify=False)   # À ÉVITER (sauf test local conscient)
```

### Checklist sécurité

- [ ] Aucun secret en dur (ni dans le code, ni dans git)
- [ ] Commandes shell en liste, pas de `shell=True` avec données variables
- [ ] SQL paramétré partout
- [ ] `yaml.safe_load`, pas de `pickle` sur données externes
- [ ] Validation des chemins de fichiers
- [ ] TLS vérifié (`verify=True` par défaut)

---

## 54. Scripts sysadmin n°1 : inventaire réseau

```python
#!/usr/bin/env python3
"""inventaire.py — balaye un réseau, teste SSH/HTTP et exporte en CSV.

Usage : python3 inventaire.py 10.0.0.0/24 -o inventaire.csv
"""
import argparse
import csv
import ipaddress
import socket
from concurrent.futures import ThreadPoolExecutor

PORTS = {"ssh": 22, "http": 80, "https": 443}

def sonder(ip):
    """Retourne un dict avec l'état des ports pour une IP."""
    resultat = {"ip": str(ip)}
    for nom, port in PORTS.items():
        try:
            with socket.create_connection((str(ip), port), timeout=1.0):
                resultat[nom] = "ouvert"
        except OSError:
            resultat[nom] = "fermé"
    try:
        resultat["hostname"] = socket.gethostbyaddr(str(ip))[0]
    except OSError:
        resultat["hostname"] = ""
    return resultat

def main():
    p = argparse.ArgumentParser(description="Inventaire réseau rapide")
    p.add_argument("reseau", help="ex : 10.0.0.0/24")
    p.add_argument("-o", "--output", default="inventaire.csv")
    p.add_argument("-w", "--workers", type=int, default=50)
    args = p.parse_args()

    reseau = ipaddress.ip_network(args.reseau, strict=False)
    hotes = list(reseau.hosts())
    print(f"Balayage de {len(hotes)} hôtes...")

    with ThreadPoolExecutor(max_workers=args.workers) as pool:
        resultats = list(pool.map(sonder, hotes))

