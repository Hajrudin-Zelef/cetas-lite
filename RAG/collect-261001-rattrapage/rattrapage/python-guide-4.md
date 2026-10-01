---
id: collect-261001-rattrapage/rattrapage/python-guide-4
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [964, 1288]
sha256: 0b8a74803652cec823d5e5a2b4ba440b4846751c96e450f52df8cab8d1a270d0
---

# Guide Python complet — De l'installation aux scripts sysadmin

```python
# Trier des IPs correctement (tri numérique, pas alphabétique !)
ips = ["10.0.0.11", "10.0.0.2"]
sorted(ips)  # ['10.0.0.11', '10.0.0.2']  ← FAUX ordre réseau
sorted(ips, key=lambda ip: tuple(int(p) for p in ip.split(".")))
# ['10.0.0.2', '10.0.0.11']  ← correct
```

> `sorted()` retourne une **nouvelle** liste ; `list.sort()` trie **en place** et retourne `None`.

---

## 20. Portée des variables (scope)

```python
COMPTEUR = 0          # globale (module)

def incrementer():
    global COMPTEUR   # déclare qu'on modifie la globale
    COMPTEUR += 1

def exemple():
    x = 10            # locale : invisible dehors
    print(x)
```

### Règles LEGB

1. **L**ocal : dans la fonction
2. **E**nclosing : fonction englobante (closures)
3. **G**lobal : module
4. **B**uiltins : `print`, `len`...

```python
def externe():
    total = 0
    def interne():
        nonlocal total   # modifie la variable de 'externe'
        total += 1
    interne()
    return total
```

### Bonnes pratiques

- Évitez `global` : préférez passer des arguments et retourner des valeurs.
- Les fonctions pures (mêmes entrées → mêmes sorties, pas d'effet de bord) sont plus faciles à tester.

---

## 21. Gestion d'erreurs : try/except

```python
try:
    with open("config.ini") as f:
        contenu = f.read()
except FileNotFoundError:
    print("Fichier introuvable, utilisation des défauts")
    contenu = ""
except PermissionError:
    print("Droits insuffisants")
    raise                       # re-lève l'exception après log
except OSError as e:            # attrape les autres erreurs OS
    print(f"Erreur OS : {e}")
else:
    print("Lecture OK")         # exécuté si PAS d'exception
finally:
    print("Toujours exécuté")   # nettoyage garanti
```

### Hiérarchie d'exceptions courantes

```
BaseException
 └─ Exception
     ├─ ValueError        (int("abc"))
     ├─ TypeError         ("a" + 1)
     ├─ KeyError          (dict["absent"])
     ├─ IndexError        (liste[99])
     ├─ FileNotFoundError
     ├─ PermissionError
     ├─ TimeoutError
     └─ RuntimeError
```

### Bonnes pratiques

| À faire | À éviter |
|---|---|
| Attraper l'exception la plus précise | `except:` nu (attrape tout, même Ctrl+C) |
| Logger le contexte | `pass` silencieux |
| `raise` pour re-propager | écraser l'info d'origine |

```python
# À éviter absolument :
try:
    risky()
except:        # attrape même KeyboardInterrupt et SystemExit !
    pass
```

---

## 22. Lever ses propres exceptions

```python
class ConfigError(Exception):
    """Erreur de configuration de l'application."""
    pass

def charger_config(chemin):
    if not chemin.endswith((".ini", ".yaml")):
        raise ConfigError(f"Format non supporté : {chemin}")
```

### raise ... from (chaînage)

```python
try:
    port = int(valeur)
except ValueError as e:
    raise ConfigError(f"Port invalide : {valeur!r}") from e
```

Le `from` conserve la cause d'origine dans la traceback — indispensable pour déboguer.

### assert : vérifications internes

```python
def pourcentage(valeur, total):
    assert total > 0, "total doit être > 0"
    return valeur / total * 100
```

> `assert` = garde-fou pendant le développement. Ne pas s'en servir pour valider des entrées utilisateur en production (désactivable avec `python -O`).

---

## 23. Modules et imports

Un module = un fichier `.py`. L'importer exécute le fichier une fois, puis le met en cache (`sys.modules`).

```python
import os                    # import simple
import os.path as osp        # alias
from pathlib import Path     # import ciblé
from collections import defaultdict, Counter
```

### Ce qu'il faut éviter

```python
from module import *   # À ÉVITER : pollue le namespace, masque les noms
```

### Où Python cherche les modules (`sys.path`)

1. Dossier du script lancé
2. `PYTHONPATH` (variable d'environnement)
3. Dossiers d'installation (site-packages du venv...)

```python
import sys
print("\n".join(sys.path))
```

### Imports relatifs (dans un package)

```python
from . import utils        # même package
from ..config import PARAM # package parent
```

> Les imports relatifs ne marchent qu'à l'intérieur d'un package importé, pas dans un script lancé directement.

---

## 24. Packages et `__init__.py`

```
mon_projet/
├── pyproject.toml
├── main.py
└── outils/
    ├── __init__.py
    ├── reseau.py
    └── logs.py
```

```python
# main.py
from outils.reseau import ping_host
from outils import logs
```

- `__init__.py` (même vide) marque le dossier comme package importable.
- Depuis Python 3.3, les *namespace packages* marchent sans `__init__.py`, mais le garder reste la pratique claire.

### `__all__` : contrôler `from x import *`

```python
# outils/reseau.py
__all__ = ["ping_host", "scan_port"]
```

### Exécuter un module comme script

```bash
python3 -m outils.reseau   # exécute outils/reseau.py avec le bon sys.path
```

> Préférez `python -m package.module` à `python package/module.py` : les imports relatifs et `sys.path` sont corrects.

---

## 25. Fichiers : lecture/écriture

### La règle d'or : `with open(...)`

```python
# Lecture
with open("serveurs.txt", encoding="utf-8") as f:
    contenu = f.read()            # tout le fichier (str)

with open("serveurs.txt", encoding="utf-8") as f:
    lignes = f.readlines()        # liste des lignes

with open("serveurs.txt", encoding="utf-8") as f:
    for ligne in f:               # itération ligne à ligne (mémoire OK)
        print(ligne.rstrip("\n"))
```

Le `with` **ferme le fichier automatiquement**, même en cas d'exception.

### Écriture

```python
with open("rapport.txt", "w", encoding="utf-8") as f:   # "w" écrase !
    f.write("Rapport du 2026-09-26\n")

with open("journal.log", "a", encoding="utf-8") as f:   # "a" ajoute
    f.write("nouvelle entrée\n")
```

### Modes d'ouverture

| Mode | Effet |
|---|---|
| `"r"` | Lecture (défaut), erreur si absent |
| `"w"` | Écriture, **écrase** le fichier |
| `"a"` | Ajout en fin de fichier |
| `"x"` | Création exclusive (erreur si existe) |
| `"r+"` | Lecture + écriture |
| `"b"` | Binaire (`"rb"`, `"wb"`) : images, zip... |

### Toujours préciser `encoding="utf-8"`

Sans `encoding`, Python utilise l'encodage locale du système (sur Windows : cp1252 → caractères accentués cassés).

### Gros fichiers : ne pas tout charger

```python
# Bien : streaming ligne par ligne
with open("gros.log", encoding="utf-8", errors="replace") as f:
    for ligne in f:
        traiter(ligne)

# errors="replace" : remplace les octets invalides par � au lieu de planter
```

---

## 26. pathlib : manipuler les chemins

`pathlib` = l'approche moderne (orientée objet). Préférez-la à `os.path`.

```python
from pathlib import Path

p = Path("/var/log/syslog")
p.name        # 'syslog'
p.parent      # Path('/var/log')
p.suffix      # ''
p.stem        # 'syslog'
p.exists()    # True/False
p.is_file()
p.is_dir()

# Construction multi-OS (gère / vs \)
dossier = Path("/srv") / "backups" / "2026-09-26"
```

### Opérations courantes

```python
Path("rapport.txt").read_text(encoding="utf-8")     # lit tout
Path("rapport.txt").write_text("hello", encoding="utf-8")

# Lister un dossier
for f in Path("/var/log").glob("*.log"):
    print(f)

for f in Path("/srv").rglob("*.conf"):   # récursif
    print(f)

Path("nouveau/dossier").mkdir(parents=True, exist_ok=True)
```

### Tableau os.path vs pathlib

| Besoin | os.path (ancien) | pathlib (moderne) |
|---|---|---|
| Joindre | `os.path.join(a, b)` | `Path(a) / b` |
| Nom du fichier | `os.path.basename(p)` | `Path(p).name` |
| Existe ? | `os.path.exists(p)` | `Path(p).exists()` |
| Lister | `os.listdir(d)` | `Path(d).iterdir()` |

> `pathlib` accepte les `~` via `Path.home()` et `Path("~/x").expanduser()`.

---

## 27. os, sys, shutil, subprocess

### os : système d'exploitation

