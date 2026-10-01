---
id: collect-261001-rattrapage/rattrapage/python-guide-6
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [1588, 1899]
sha256: 939da6b013939ade73f83c9ffec4934fc6be1ea5ae2d4a6cf967080029e0bbca
---

# Guide Python complet — De l'installation aux scripts sysadmin

```python
def reessayer(tentatives=3, delai=2):
    def decorateur(fonction):
        @functools.wraps(fonction)
        def enveloppe(*args, **kwargs):
            for i in range(tentatives):
                try:
                    return fonction(*args, **kwargs)
                except Exception as e:
                    if i == tentatives - 1:
                        raise
                    print(f"Tentative {i+1} échouée ({e}), nouvel essai...")
                    time.sleep(delai)
        return enveloppe
    return decorateur

@reessayer(tentatives=3, delai=5)
def requete_api(url):
    ...
```

### Décorateurs standards utiles

| Décorateur | Usage |
|---|---|
| `@functools.lru_cache` | Mémorise les résultats (cache) |
| `@property` | Attribut calculé (voir ci-dessous) |
| `@staticmethod` / `@classmethod` | Méthodes sans instance |

```python
from functools import lru_cache

@lru_cache(maxsize=128)
def resolution_dns(hostname):
    # coûteux → mis en cache automatiquement
    ...
```

```python
class Serveur:
    def __init__(self, ip):
        self.ip = ip

    @property
    def reseau(self):
        return ".".join(self.ip.split(".")[:3]) + ".0/24"

srv = Serveur("10.0.5.11")
print(srv.reseau)   # 10.0.5.0/24 — appelé comme un attribut !
```

---

## 34. Générateurs et itérateurs

Un générateur produit les valeurs **à la demande**, sans tout stocker en mémoire.

```python
def lignes_erreurs(chemin):
    with open(chemin, encoding="utf-8", errors="replace") as f:
        for ligne in f:
            if "error" in ligne.lower():
                yield ligne        # pause ici, reprend au prochain next()

# Usage : ne charge JAMAIS tout le fichier en RAM
for ligne in lignes_erreurs("/var/log/syslog"):
    print(ligne, end="")
```

### Générateur en compréhension

```python
total = sum(int(l.split()[3]) for l in open("tailles.txt"))
# ^ parenthèses = générateur : 1 ligne en mémoire à la fois
```

### itertools : la boîte à outils

```python
import itertools

itertools.islice(range(10**9), 5)        # les 5 premiers, sans tout générer
itertools.chain([1, 2], [3, 4])          # chaîner des itérables
itertools.groupby(donnees_triees, key=...)  # regrouper (données triées !)
itertools.batched(range(10), 3)          # par paquets de 3 (Python 3.12+)
```

> Pour traiter des logs de plusieurs Go : générateurs + `itertools` = mémoire constante.

---

## 35. Context managers (`with`)

Le `with` garantit le nettoyage (fermeture, verrou, connexion...).

```python
with open("f.txt") as f:      # ferme le fichier à la sortie
    ...

import threading
verrou = threading.Lock()
with verrou:                   # libère le verrou à la sortie
    ...
```

### Créer son context manager (2 façons)

```python
# Façon 1 : classe avec __enter__ / __exit__
class Chrono:
    def __enter__(self):
        self.debut = time.perf_counter()
        return self
    def __exit__(self, exc_type, exc, tb):
        print(f"Durée : {time.perf_counter() - self.debut:.2f}s")
        return False   # False = propage l'exception éventuelle

with Chrono():
    traitement_lourd()

# Façon 2 : décorateur @contextmanager (plus simple)
from contextlib import contextmanager

@contextmanager
def connexion_bd(chemin):
    conn = ouvrir(chemin)
    try:
        yield conn          # le 'as conn' reçoit ceci
    finally:
        conn.close()        # toujours exécuté

with connexion_bd("parc.db") as conn:
    ...
```

---

## 36. Annotations de type (typing)

Les annotations documentent les types attendus. Python ne les vérifie **pas** à l'exécution (sauf outils externes comme `mypy`).

```python
def ping(host: str, timeout: float = 5.0) -> bool:
    ...

# Variables
compteur: int = 0
serveurs: list[str] = []            # Python 3.9+ (sinon List[str] de typing)
config: dict[str, str] = {}
```

### Types courants

```python
from typing import Optional, Union

def trouver(nom: str) -> Optional[str]:
    # retourne str ou None
    ...

# Python 3.10+ : syntaxe moderne avec |
def trouver(nom: str) -> str | None:
    ...

def traiter(valeur: int | str) -> None: ...
```

### Pourquoi typer ?

- L'IDE détecte les erreurs avant l'exécution (autocomplétion, refactoring).
- La signature devient une documentation vivante.
- `mypy` en CI attrape des bugs : `mypy mon_script.py`.

> Pour un sysadmin : typez au moins les signatures des fonctions partagées/réutilisées. Pas besoin de tout typer dans un script jetable.

---

## 37. Expressions régulières (regex)

```python
import re

log = "2026-09-26 10:15:32 FAILED login from 192.168.1.25 port 22"

# Recherche
m = re.search(r"from (\d+\.\d+\.\d+\.\d+)", log)
if m:
    print(m.group(1))   # 192.168.1.25

# Toutes les occurrences
re.findall(r"\d+\.\d+\.\d+\.\d+", log)

# Remplacement
re.sub(r"port \d+", "port XXX", log)

# Découpage
re.split(r"\s+", "a  b   c")   # ['a', 'b', 'c']
```

### Toujours des raw strings `r"..."`

Sans `r`, `\d` devient un problème d'échappement. **Réflexe : regex = `r"..."`.**

### Méta-caractères essentiels

| Motif | Sens |
|---|---|
| `\d` | chiffre |
| `\w` | lettre/chiffre/`_` |
| `\s` | espace |
| `.` | n'importe quel caractère |
| `*` / `+` / `?` | 0+, 1+, 0-1 répétitions |
| `{2,4}` | entre 2 et 4 répétitions |
| `^` / `$` | début / fin de chaîne |
| `(...)` | groupe capturant |
| `[...]` | classe : `[0-9a-f]` |
| `\|` | alternative |

### Exemple sysadmin : parser des logs SSH

```python
import re

MOTIF = re.compile(
    r"(?P<date>\S+ \S+) \S+ sshd\[\d+\]: Failed password for (invalid user )?(?P<user>\S+) from (?P<ip>\S+)"
)

with open("/var/log/auth.log", encoding="utf-8", errors="replace") as f:
    for ligne in f:
        m = MOTIF.search(ligne)
        if m:
            print(m.group("date"), m.group("user"), m.group("ip"))
```

> `re.compile()` : pré-compile le motif quand on l'utilise en boucle (performance + lisibilité). Groupes nommés `(?P<nom>...)` = code auto-documenté.

### Pièges regex

- `.*` est **glouton** : `"<a>1</a><a>2</a>"` avec `<.*>` capture tout. Utilisez `<.*?>` (non-glouton).
- Valider un email "parfaitement" en regex est un mythe ; contentez-vous d'un motif raisonnable.

---

## 38. JSON, YAML, CSV

### JSON (bibliothèque standard)

```python
import json

# Python → JSON
config = {"host": "srv-01", "ports": [22, 80]}
texte = json.dumps(config, indent=2, ensure_ascii=False)
Path("config.json").write_text(texte, encoding="utf-8")

# JSON → Python
config = json.loads(Path("config.json").read_text(encoding="utf-8"))
```

### YAML (paquet externe `pyyaml`)

```bash
pip install pyyaml
```

```python
import yaml

# Lecture sûre : TOUJOURS safe_load (pas load !)
with open("inventaire.yaml", encoding="utf-8") as f:
    data = yaml.safe_load(f)

texte = yaml.safe_dump(data, allow_unicode=True, sort_keys=False)
```

> `yaml.load` sans loader = exécution de code arbitraire possible. `safe_load` systématique.

### CSV (bibliothèque standard)

```python
import csv

# Lecture
with open("parc.csv", newline="", encoding="utf-8") as f:
    for ligne in csv.DictReader(f):     # dict par ligne (via l'en-tête)
        print(ligne["hostname"], ligne["ip"])

# Écriture
with open("export.csv", "w", newline="", encoding="utf-8") as f:
    w = csv.DictWriter(f, fieldnames=["hostname", "ip"])
    w.writeheader()
    w.writerow({"hostname": "srv-01", "ip": "10.0.0.11"})
```

> `newline=""` à l'ouverture : évite les lignes vides parasites sous Windows.

---

## 39. datetime : dates et heures

```python
from datetime import datetime, date, time, timedelta, timezone

maintenant = datetime.now()                    # heure locale (naïve)
utc = datetime.now(timezone.utc)               # heure UTC (aware) ← préférez
aujourdhui = date.today()                      # 2026-09-26

datetime(2026, 9, 26, 10, 30)                  # construction
```

### Parsing et formatage

