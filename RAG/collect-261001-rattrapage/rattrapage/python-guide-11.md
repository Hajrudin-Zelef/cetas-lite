---
id: collect-261001-rattrapage/rattrapage/python-guide-11
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["decode"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [3071, 3361]
sha256: 10d176cd271092f6fce0619f45d9007233bfb13ac6b48741652b5e3e36c3bc4a
---

# Guide Python complet — De l'installation aux scripts sysadmin

## 60. Scripts sysadmin n°7 : scan de ports et services

```python
#!/usr/bin/env python3
"""scan.py — scan rapide des ports courants d'un hôte + bannière.

Usage : python3 scan.py srv-01 [--ports 22,80,443] [--timeout 1.5]
"""
import argparse
import socket

PORTS_COURANTS = [21, 22, 23, 25, 53, 80, 110, 143, 443, 445,
                  3306, 3389, 5432, 5900, 6379, 8080, 8443]

def parse_ports(texte):
    return [int(p) for p in texte.split(",")]

def scanner_port(host, port, timeout):
    try:
        with socket.create_connection((host, port), timeout=timeout) as s:
            s.settimeout(1.0)
            try:
                banniere = s.recv(1024).decode("utf-8", errors="replace").strip()
            except OSError:
                banniere = ""
            return True, banniere.split("\n")[0][:60]
    except OSError:
        return False, ""

def main():
    p = argparse.ArgumentParser(description="Scan de ports simple")
    p.add_argument("host")
    p.add_argument("--ports", default=",".join(map(str, PORTS_COURANTS)),
                   help="liste séparée par des virgules")
    p.add_argument("--timeout", type=float, default=1.0)
    args = p.parse_args()

    print(f"Scan de {args.host}...\n{'PORT':<8}{'ÉTAT':<10}BANNIÈRE")
    for port in parse_ports(args.ports):
        ouvert, banniere = scanner_port(args.host, port, args.timeout)
        etat = "OUVERT" if ouvert else "fermé"
        print(f"{port:<8}{etat:<10}{banniere}")

if __name__ == "__main__":
    main()
```

---

## 61. Erreurs classiques des débutants (20 pièges)

### 1. `=` vs `==`

```python
if x = 5:    # SyntaxError : affectation dans un test
if x == 5:   # comparaison
```

### 2. Oublier les deux-points

```python
if x > 0     # SyntaxError
    print(x)
```

### 3. Mélanger tabulations et espaces → `TabError`

Configurez l'éditeur : "insérer des espaces quand on appuie sur Tab".

### 4. Comparer avec `is` au lieu de `==`

```python
a == b   # égalité de contenu ✓ (sauf None : utilisez 'is None')
```

### 5. Argument mutable par défaut

```python
# BUG : la même liste est partagée entre tous les appels !
def ajouter(x, liste=[]):
    liste.append(x)
    return liste

# CORRECT :
def ajouter(x, liste=None):
    if liste is None:
        liste = []
    liste.append(x)
    return liste
```

### 6. Modifier une liste pendant qu'on l'itère

```python
# BUG : éléments sautés
for srv in serveurs:
    if not ping(srv):
        serveurs.remove(srv)

# CORRECT : itérer sur une copie
for srv in serveurs[:]:
    ...
# ou construire une nouvelle liste :
serveurs = [s for s in serveurs if ping(s)]
```

### 7. `except:` nu qui masque tout

```python
try:
    ...
except Exception as e:      # au minimum, pas de 'except:' nu
    logger.exception("contexte : %s", e)
    raise
```

### 8. Oublier `return` (la fonction retourne `None`)

```python
def double(x):
    x * 2          # calculé puis jeté !
double(21)         # None
```

### 9. `bool("False")` vaut `True`

Voir section 12 : comparez explicitement les chaînes.

### 10. Division entière surprise (Python 2 vs 3)

En Python 3, `7 / 2 = 3.5` (float). Si vous voyez `3`, c'est `//` ou un vieux Python 2.

### 11. Encodage : oublier `encoding="utf-8"`

Accents cassés sur Windows, `UnicodeDecodeError` sur des logs exotiques → `encoding="utf-8", errors="replace"`.

### 12. Chemins Windows en dur

```python
"C:\nouveau\dossier"   # \n = saut de ligne ! BUG
r"C:\nouveau\dossier"  # raw string ✓
Path("C:/nouveau/dossier")  # ou pathlib ✓
```

### 13. `sorted()` ne trie pas en place

```python
serveurs = sorted(serveurs)   # ✓ (nouvelle liste)
serveurs.sort()               # ✓ (en place)
sorted(serveurs)              # ✗ trié puis jeté !
```

### 14. Copie de liste par référence

```python
b = a       # même objet ! → b = a.copy()
```

### 15. Variable de boucle qui fuit

```python
for i in range(5):
    ...
print(i)   # 4 : 'i' existe encore après la boucle (pas de scope de bloc)
```

### 16. `range()` : fin exclue

`range(1, 10)` = 1..9, pas 1..10. Source classique d'erreurs off-by-one.

### 17. Comparer des types incompatibles

```python
"10" > 9    # TypeError en Python 3 (pas de comparaison implicite)
int("10") > 9  # ✓
```

### 18. Import circulaire

`a.py` importe `b.py` qui importe `a.py` → `ImportError`. Solution : restructurer, ou import local dans la fonction.

### 19. Ne pas fermer les ressources (sans `with`)

Fichier, connexion, socket : utilisez toujours un context manager.

### 20. `time.sleep()` dans du code asyncio

```python
async def tache():
    await asyncio.sleep(1)   # ✓ non bloquant (dans une coroutine)
    # time.sleep(1)          # ✗ bloque TOUTE la boucle événementielle
```

---

## 62. Cas pratiques commentés

### Cas n°1 : nettoyer des noms d'hôtes

```python
def normaliser_hostname(brut: str) -> str:
    """' SRV-Web_01 ' → 'srv-web-01'."""
    propre = brut.strip().lower()          # retire espaces, minuscules
    propre = propre.replace("_", "-")     # uniformise le séparateur
    # garde lettres, chiffres, tirets uniquement
    propre = "".join(c for c in propre if c.isalnum() or c == "-")
    return propre

for h in [" SRV-Web_01 ", "db#02", "cache_03"]:
    print(f"{h!r:18} → {normaliser_hostname(h)}")
# ' SRV-Web_01 '     → srv-web-01
# 'db#02'            → db02
# 'cache_03'         → cache-03
```

### Cas n°2 : calculer un taux de disponibilité depuis un log

```python
from datetime import datetime

def disponibilite(chemin, service="nginx"):
    """% de lignes 'started' vs 'failed' pour un service dans un journal."""
    ok = ko = 0
    with open(chemin, encoding="utf-8", errors="replace") as f:
        for ligne in f:
            if service not in ligne:
                continue
            if "started" in ligne:
                ok += 1
            elif "failed" in ligne:
                ko += 1
    total = ok + ko
    return ok / total * 100 if total else 0.0

print(f"Disponibilité nginx : {disponibilite('journal.log'):.2f}%")
```

### Cas n°3 : renommer des fichiers en masse

```python
from pathlib import Path

dossier = Path("/srv/exports")
for f in dossier.glob("rapport_*.csv"):
    # rapport_2026-01-05.csv → 2026-01-05_rapport.csv
    nouveau_nom = f"{f.stem[8:]}_{f.stem[:8]}{f.suffix}"
    f.rename(f.with_name(nouveau_nom))
    print(f"{f.name} → {nouveau_nom}")
```

### Cas n°4 : vérifier l'expiration de certificats (idée d'implémentation)

```python
import socket, ssl
from datetime import datetime, timezone

def expiration_certificat(hostname, port=443):
    """Retourne la date d'expiration du certificat TLS."""
    ctx = ssl.create_default_context()
    with socket.create_connection((hostname, port), timeout=5) as sock:
        with ctx.wrap_socket(sock, server_hostname=hostname) as tls:
            cert = tls.getpeercert()
    fin = datetime.strptime(cert["notAfter"], "%b %d %H:%M:%S %Y %Z")
    return fin.replace(tzinfo=timezone.utc)

exp = expiration_certificat("exemple.com")
jours = (exp - datetime.now(timezone.utc)).days
print(f"Expire dans {jours} jours ({exp.date()})")
```

### Cas n°5 : mini tableau de bord texte

```python
import shutil

def barre(pct, largeur=20):
    remplis = int(pct / 100 * largeur)
    return "█" * remplis + "░" * (largeur - remplis)

for point in ["/", "/var", "/srv"]:
    u = shutil.disk_usage(point)
    pct = u.used / u.total * 100
    print(f"{point:<6} {barre(pct)} {pct:5.1f}%  ({u.free/1024**3:.1f} Go libres)")
```

---

## 63. Checklist "script de production"

Avant de mettre un script en cron ou sur un serveur :

