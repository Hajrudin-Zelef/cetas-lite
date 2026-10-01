---
id: collect-261001-rattrapage/rattrapage/python-guide-8
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "benchmark"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [2215, 2522]
sha256: d47edb5e0950dcee035fb5e25d0fca01c142a25e38a2d0f5eff8975f2b1c1ef0
---

# Insertion SÉCURISÉE (paramètres ? : anti-injection SQL)
cur.execute(
    "INSERT OR IGNORE INTO equipements (hostname, ip, type, site) VALUES (?, ?, ?, ?)",
    ("srv-web-01", "10.0.0.11", "serveur", "paris"),
)
conn.commit()

# Lecture
for row in cur.execute("SELECT hostname, ip FROM equipements WHERE site = ?", ("paris",)):
    print(row["hostname"], row["ip"])

conn.close()
```

### Avec context manager (commit/rollback auto)

```python
with sqlite3.connect("parc.db") as conn:
    conn.execute("INSERT INTO equipements ... ", (...))
    # commit automatique à la sortie ; rollback si exception
```

### Bonnes pratiques sqlite3

- **Toujours** des requêtes paramétrées (`?`), jamais de f-string dans le SQL.
- `conn.commit()` après les écritures (ou context manager).
- Un seul writer à la fois : SQLite verrouille le fichier en écriture.
- Sauvegarde = copier le fichier `.db` (ou `VACUUM INTO 'backup.db'`).

---

## 45. Tests unitaires : unittest et pytest

### Pourquoi tester ?

Un script sysadmin qui tourne à 3h du matin doit être fiable. Les tests vérifient que vos fonctions font ce qu'elles annoncent, **avant** la production.

### pytest (recommandé)

```bash
pip install pytest
```

```python
# test_reseau.py
from reseau import est_ip_privee, normaliser_mac

def test_ip_privee():
    assert est_ip_privee("10.0.0.1") is True
    assert est_ip_privee("8.8.8.8") is False

def test_mac():
    assert normaliser_mac("AA:BB:CC:DD:EE:FF") == "aa:bb:cc:dd:ee:ff"
```

```bash
pytest -v              # exécute les tests
pytest -x              # s'arrête à la 1ère erreur
pytest -k "mac"        # filtre par nom
```

### Fixtures et paramétrisation

```python
import pytest

@pytest.mark.parametrize("ip,attendu", [
    ("10.0.0.1", True),
    ("172.16.5.4", True),
    ("192.168.1.1", True),
    ("8.8.8.8", False),
])
def test_ip_privee_param(ip, attendu):
    assert est_ip_privee(ip) is attendu

@pytest.fixture
def base_test(tmp_path):
    db = tmp_path / "test.db"
    initialiser_base(str(db))
    return str(db)

def test_insertion(base_test):
    inserer_equipement(base_test, "srv-test", "10.9.9.9")
    assert compter(base_test) == 1
```

### unittest (stdlib, sans dépendance)

```python
import unittest

class TestReseau(unittest.TestCase):
    def test_ip(self):
        self.assertTrue(est_ip_privee("10.0.0.1"))
        self.assertEqual(normaliser_mac("AA"), "aa")

if __name__ == "__main__":
    unittest.main()
```

> Convention : fichiers `test_*.py` ou `*_test.py`, fonctions `test_*`. Visez 80 % de couverture sur le code critique (`pytest --cov` avec `pytest-cov`).

---

## 46. Débogage : pdb et techniques

### print() ciblé (le classique)

```python
print(f"DEBUG: host={host!r} port={port}")   # !r = repr, voit les espaces
```

### pdb : le débogueur intégré

```python
# Point d'arrêt dans le code :
breakpoint()   # Python 3.7+ : équivaut à import pdb; pdb.set_trace()
```

```bash
python3 -m pdb script.py     # lance sous pdb
```

| Commande pdb | Effet |
|---|---|
| `n` (next) | ligne suivante (sans entrer dans les fonctions) |
| `s` (step) | entre dans la fonction |
| `c` (continue) | reprend jusqu'au prochain breakpoint |
| `p var` | affiche une variable |
| `l` | montre le code autour |
| `b 42` | breakpoint ligne 42 |
| `q` | quitte |

### Techniques rapides

```python
import traceback

try:
    operation_risquee()
except Exception:
    traceback.print_exc()      # affiche la pile complète dans les logs
    # ou : logger.exception("Échec")  ← avec logging
```

- `python3 -i script.py` : console interactive après exécution (inspectez les variables).
- Lisez la traceback **de bas en haut** : la dernière ligne = l'erreur réelle.

---

## 47. Asyncio : programmation asynchrone

`asyncio` = faire des milliers d'opérations d'attente (réseau) **en parallèle** dans un seul thread. Idéal pour interroger 500 équipements.

```python
import asyncio

async def interroger(host):
    await asyncio.sleep(0.5)          # simule une attente réseau
    return f"{host}: OK"

async def main():
    hosts = [f"srv-{i:02d}" for i in range(10)]
    resultats = await asyncio.gather(*(interroger(h) for h in hosts))
    print(resultats)

asyncio.run(main())   # point d'entrée
```

### Concepts clés

| Notion | Sens |
|---|---|
| `async def` | définit une coroutine |
| `await` | "attends ici, fais autre chose en attendant" |
| `asyncio.gather(...)` | lance plusieurs coroutines en parallèle |
| `asyncio.run(...)` | exécute le programme async |

### Exemple réaliste : requêtes HTTP parallèles (avec aiohttp)

```bash
pip install aiohttp
```

```python
import asyncio, aiohttp

async def fetch(session, url):
    async with session.get(url, timeout=10) as r:
        return url, r.status

async def main(urls):
    async with aiohttp.ClientSession() as session:
        resultats = await asyncio.gather(*(fetch(session, u) for u in urls))
    return resultats

urls = [f"https://srv-{i:02d}.lan/health" for i in range(50)]
print(asyncio.run(main(urls)))
```

### Limiter la concurrence (ne pas DoS ses propres serveurs !)

```python
sem = asyncio.Semaphore(20)   # max 20 requêtes simultanées

async def fetch_limite(session, url):
    async with sem:
        return await fetch(session, url)
```

> `asyncio` ≠ threads : un seul thread, pas de parallélisme CPU. Pour du calcul lourd, voir multiprocessing (section 48).

---

## 48. Threading et multiprocessing

### Choisir le bon outil

| Situation | Outil |
|---|---|
| Beaucoup d'attentes réseau/disque (I/O bound) | `threading` ou `asyncio` |
| Calculs lourds (CPU bound) | `multiprocessing` |
| Code simple, parallélisme modéré | `concurrent.futures` |

### concurrent.futures : l'API simple (recommandée)

```python
from concurrent.futures import ThreadPoolExecutor, as_completed
import subprocess

def ping(host):
    r = subprocess.run(["ping", "-c", "1", "-W", "1", host],
                       capture_output=True)
    return host, r.returncode == 0

hosts = [f"10.0.0.{i}" for i in range(1, 51)]

with ThreadPoolExecutor(max_workers=20) as pool:
    futurs = {pool.submit(ping, h): h for h in hosts}
    for fut in as_completed(futurs):
        host, ok = fut.result()
        print(f"{host}: {'OK' if ok else 'KO'}")
```

### multiprocessing : contourner le GIL

```python
from concurrent.futures import ProcessPoolExecutor

def calcul_lourd(n):
    return sum(i * i for i in range(n))

if __name__ == "__main__":      # OBLIGATOIRE sous Windows/macOS
    with ProcessPoolExecutor() as pool:
        print(list(pool.map(calcul_lourd, [10**7, 10**7, 10**7, 10**7])))
```

> **Le GIL** (Global Interpreter Lock) : un seul thread Python exécute du bytecode à la fois. Les threads n'accélèrent donc **pas** le calcul pur — mais ils sont parfaits pour les I/O.

### threading bas niveau (verrous)

```python
import threading

verrou = threading.Lock()
compteur = 0

def incrementer():
    global compteur
    with verrou:          # protège la section critique
        compteur += 1
```

---

## 49. Performance : mesurer et optimiser

### Mesurer d'abord, optimiser ensuite

```python
import time

debut = time.perf_counter()   # horloge précise (pas time.time !)
traitement()
print(f"{time.perf_counter() - debut:.3f} s")
```

```bash
# timeit : micro-benchmark fiable
python3 -m timeit "'-'.join(str(i) for i in range(1000))"
python3 -m timeit --setup "import mon_module" "mon_module.f()"
```

### cProfile : où passe le temps ?

```bash
python3 -m cProfile -s cumulative script.py | head -30
```

```python
import cProfile, pstats
cProfile.run("traitement()", "stats.prof")
p = pstats.Stats("stats.prof")
p.sort_stats("cumulative").print_stats(15)
```

### Règles d'optimisation (par ordre d'impact)

