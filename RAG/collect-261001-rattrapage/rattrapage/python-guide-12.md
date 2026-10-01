---
id: collect-261001-rattrapage/rattrapage/python-guide-12
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: ["Lambda"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [3362, 3544]
sha256: d31bbf8733ed41ad151c5479e1296608a3cd51d31912bf7eee9b2a0c0d76b848
---

# Guide Python complet — De l'installation aux scripts sysadmin

- [ ] Shebang `#!/usr/bin/env python3` + `chmod +x`
- [ ] Bloc `if __name__ == "__main__":` présent
- [ ] `argparse` pour les paramètres (pas de valeurs en dur)
- [ ] `logging` au lieu de `print()` (fichier + niveau configurable)
- [ ] Gestion d'erreurs : `try/except` ciblés, pas de `except: pass`
- [ ] Timeouts sur tout ce qui est réseau (`requests`, `socket`)
- [ ] Codes de retour explicites (`sys.exit(0/1/2)`)
- [ ] Secrets via variables d'environnement, jamais en dur
- [ ] `encoding="utf-8"` sur les ouvertures de fichiers texte
- [ ] Testé avec `python3 -m py_compile` (syntaxe) + un jeu de test réel
- [ ] `requirements.txt` ou `pyproject.toml` à jour
- [ ] Testé depuis un venv propre (pas seulement votre poste)
- [ ] Documenté : docstring + `--help` clair + exemple d'usage
- [ ] Idempotent si possible (relançable sans dégâts)
- [ ] Journalisé : on sait ce qu'il a fait (et quand il a échoué)

---

## 64. Pense-bête de poche

### Lancement & environnement

```bash
python3 script.py arg1 --opt      # exécuter
python3 -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
python3 -m py_compile script.py  # vérif syntaxe rapide
python3 -i script.py             # console après exécution
```

### Structures

```python
[x for x in l if cond]           # compréhension filtrante
{x: f(x) for x in l}             # dict en compréhension
a, b = b, a                      # échange
for i, x in enumerate(l): ...    # index + valeur
for a, b in zip(l1, l2): ...     # parcours parallèle
```

### Fichiers & chemins

```python
from pathlib import Path
Path("f.txt").read_text(encoding="utf-8")
Path("f.txt").write_text("...", encoding="utf-8")
list(Path(".").glob("*.log"))
```

### Réseau & système

```python
subprocess.run(["cmd", "arg"], capture_output=True, text=True, check=True)
socket.create_connection((host, port), timeout=2)
requests.get(url, timeout=10).raise_for_status()
```

### Debug express

```text
breakpoint()                     # stoppe ici (pdb)
f"{var=}"                        # f-string debug : affiche nom=valeur (3.8+)
python3 -m pdb script.py
python3 -m cProfile -s cumulative script.py
```

### Datetime express

```python
from datetime import datetime, timezone, timedelta
datetime.now(timezone.utc).isoformat()
datetime.strptime(s, "%Y-%m-%d %H:%M:%S")
```

---

## 65. Glossaire

| Terme | Définition |
|---|---|
| Argument | Valeur passée à une fonction à l'appel |
| Attribut | Variable attachée à un objet/classe |
| Bytecode | Code intermédiaire exécuté par la VM Python (`.pyc`) |
| Classe | Moule définissant objets (attributs + méthodes) |
| Closure | Fonction qui capture des variables de son contexte |
| Compréhension | Construction concise de collection `[x for x in ...]` |
| Coroutine | Fonction `async def` exécutable par `asyncio` |
| Décorateur | Fonction qui enveloppe une autre (`@deco`) |
| Dict | Dictionnaire : mapping clé → valeur |
| Docstring | Chaîne de documentation (`"""..."""`) |
| Dunder | Méthode spéciale `__nom__` (double underscore) |
| Exception | Erreur levée à l'exécution, attrapable via `try` |
| f-string | Chaîne formatée `f"..."` (Python 3.6+) |
| Générateur | Fonction avec `yield` : produit à la demande |
| GIL | Verrou global : 1 seul thread exécute du bytecode à la fois |
| Idempotent | Relançable sans effet de bord supplémentaire |
| Instance | Objet créé à partir d'une classe |
| Itérable | Objet parcourable par `for` / `in` |
| Kwargs | Arguments nommés supplémentaires (`**kwargs`) |
| Lambda | Petite fonction anonyme |
| Linter | Outil d'analyse statique (ruff, flake8) |
| Méthode | Fonction définie dans une classe |
| Module | Fichier `.py` importable |
| Mutable / Immutable | Modifiable (`list`) / non (`tuple`, `str`) |
| Namespace | Espace de noms (module, fonction...) |
| Package | Dossier de modules (`__init__.py`) |
| Paramètre | Variable déclarée dans la signature d'une fonction |
| PEP 8 | Guide de style officiel Python |
| Pickle | Sérialisation Python native (dangereuse sur données externes) |
| Polymorphisme | Même interface, comportements différents selon l'objet |
| REPL | Console interactive (`python3` sans argument) |
| Scope | Portée d'une variable (local/global...) |
| Shebang | `#!/usr/bin/env python3` en 1ère ligne |
| Slice | Sous-séquence `seq[début:fin:pas]` |
| Thread | Fil d'exécution (partage la mémoire) |
| Traceback | Pile d'appels affichée lors d'une exception |
| Tuple | Séquence immutable `(a, b)` |
| Type hint | Annotation de type (`def f(x: int) -> bool`) |
| venv | Environnement virtuel isolé |
| Wheel | Format de distribution binaire (`.whl`) |

---

## 66. Quiz (10 questions + réponses)

**Q1.** Que vaut `bool("")`, `bool([])`, `bool(0)` ?
> `False` pour les trois : ce sont des valeurs "falsy".

**Q2.** Quelle est la différence entre `a = b` et `a = b.copy()` pour une liste ?
> `a = b` : `a` pointe vers le **même** objet (modifier l'un modifie l'autre). `.copy()` crée une vraie copie indépendante (superficielle).

**Q3.** Pourquoi ce code est-il bogué ?
> ```text
> def ajouter(x, historique=[]):
>     historique.append(x)
>     return historique
> ```
> Le défaut `[]` est créé **une seule fois** à la définition : tous les appels partagent la même liste. Corriger avec `historique=None` puis création dans le corps.

**Q4.** `sorted(maliste)` vs `maliste.sort()` : différence ?
> `sorted()` retourne une **nouvelle** liste triée (original intact). `.sort()` trie **en place** et retourne `None`.

**Q5.** Dans quel ordre Python cherche-t-il un nom de variable (règle LEGB) ?
> **L**ocal → **E**nclosing → **G**lobal → **B**uiltins.

**Q6.** Pourquoi faut-il préférer `subprocess.run(["ping", host])` à `subprocess.run(f"ping {host}", shell=True)` ?
> Avec `shell=True` et une chaîne, un `host` malveillant (`"x; rm -rf /"`) provoque une **injection de commandes**. La liste d'arguments évite le shell.

**Q7.** Que fait le `else` d'une boucle `for` ?
> Il s'exécute **si la boucle s'est terminée sans `break`**. Utile pour "recherché mais non trouvé".

**Q8.** Pourquoi `0.1 + 0.2 != 0.3` en Python ?
> Les `float` sont en base 2 (IEEE 754) : `0.1` n'est pas représentable exactement. Utilisez `round()` ou `decimal.Decimal` pour des calculs exacts.

**Q9.** Threads ou multiprocessing pour : (a) pinger 200 hôtes, (b) chiffrer 4 gros fichiers ?
> (a) **Threads** (ou asyncio) : tâche I/O-bound, on attend le réseau. (b) **Multiprocessing** : tâche CPU-bound, le GIL empêche les threads de paralléliser le calcul.

**Q10.** Que manque-t-il à ce script pour être "production-ready" ?
> ```text
> import requests
> r = requests.get("https://api.lan/data")
> print(r.json())
> ```
> Au minimum : un `timeout=` (sinon blocage infini), `raise_for_status()` ou test du code HTTP, et un `try/except requests.RequestException`.

---

## 67. Pour aller plus loin

### Bibliothèques à connaître (sysadmin & réseau)

| Domaine | Bibliothèque | Usage |
|---|---|---|
| HTTP | `requests`, `httpx` (sync+async) | API, web |
| SSH | `paramiko`, `netmiko`, `nornir` | automatisation d'équipements réseau |
| Parsing | `pyyaml`, `openpyxl`, `pandas` | configs, Excel, data |
| CLI jolies | `rich`, `typer` | tableaux, barres de progression |
| Scheduling | `schedule`, `APScheduler` | tâches périodiques en Python |
| Chiffrement | `cryptography` | TLS, certificats, hash |
| Tests | `pytest`, `pytest-cov` | qualité |

### Sujets d'approfondissement

