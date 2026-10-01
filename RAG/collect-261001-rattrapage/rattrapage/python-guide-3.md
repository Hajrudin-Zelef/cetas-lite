---
id: collect-261001-rattrapage/rattrapage/python-guide-3
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [616, 963]
sha256: 03b99da427561c48436356136415f961f365bd1e2428e66dcfd1619edb2d9a2f
---

# Guide Python complet — De l'installation aux scripts sysadmin

## 11. Sets (ensembles)

Ensembles non ordonnés, **sans doublons**. Parfaits pour dédupliquer et comparer.

```python
ips = {"10.0.0.1", "10.0.0.2", "10.0.0.1"}  # {'10.0.0.1', '10.0.0.2'}
vide = set()          # {} crée un DICT vide, pas un set !
```

### Opérations ensemblistes

```python
a = {"web", "db", "cache"}
b = {"db", "lb"}

a | b   # union : {'web', 'db', 'cache', 'lb'}
a & b   # intersection : {'db'}
a - b   # différence : {'web', 'cache'}
a ^ b   # différence symétrique : {'web', 'cache', 'lb'}
```

### Méthodes

```python
ips.add("10.0.0.3")
ips.discard("10.0.0.9")   # pas d'erreur si absent
ips.remove("10.0.0.1")    # KeyError si absent
```

### Cas d'usage sysadmin : dédupliquer des IPs de logs

```python
with open("/var/log/auth.log") as f:
    ips = {ligne.split()[10] for ligne in f if "Failed" in ligne}
print(f"{len(ips)} IP(s) distinctes en échec d'authentification")
```

---

## 12. Conversions de types

| Depuis → Vers | Code | Note |
|---|---|---|
| str → int | `int("42")` | `ValueError` si non numérique |
| str → float | `float("3.14")` | |
| nombre → str | `str(42)` | |
| list → set | `set([1,1,2])` | déduplique |
| set → list | `list({1,2})` | ordre non garanti |
| list → tuple | `tuple([1,2])` | |
| liste de paires → dict | `dict([("a",1)])` | |
| str → list de lignes | `"a\nb".splitlines()` | |
| str → bool | `s.lower() == "true"` | pas de `bool("False")` ! |

> **Piège :** `bool("False")` vaut `True` (chaîne non vide = truthy). Pour parser du booléen depuis un fichier de config, comparez explicitement.

```python
def str_to_bool(s):
    return s.strip().lower() in ("1", "true", "yes", "on")
```

---

## 13. Conditions (if/elif/else)

```python
charge_cpu = 85

if charge_cpu > 90:
    print("CRITIQUE")
elif charge_cpu > 75:
    print("Attention")
else:
    print("OK")
```

### Opérateurs de comparaison

| Opérateur | Sens |
|---|---|
| `==` | égal |
| `!=` | différent |
| `>`, `<`, `>=`, `<=` | ordre |
| `is` | même objet (identité, pas égalité) |
| `is not` | objets différents |
| `in` | appartenance |

```python
# Comparaisons chaînées (lisible !)
if 0 <= charge_cpu <= 100:
    print("valeur plausible")
```

### Opérateurs logiques

```python
if disque_plein and not maintenance:
    alerter()
if port == 22 or port == 2222:
    print("SSH")
if not reponse:
    print("vide")
```

> `and`/`or` sont **court-circuités** : `or` s'arrête au premier truthy. Pratique pour des valeurs par défaut : `timeout = timeout or 30`.

### Ternaire (expression conditionnelle)

```python
etat = "CRITIQUE" if charge > 90 else "OK"
```

### `is` vs `==`

```python
a = [1, 2]; b = [1, 2]
a == b   # True  (même contenu)
a is b   # False (objets différents)

x = None
if x is None:   # toujours 'is' pour None, jamais '=='
    ...
```

---

## 14. Boucles (for, while)

### for : parcours d'itérables

```python
for srv in ["web-01", "web-02"]:
    print(f"Ping {srv}...")

for i in range(5):          # 0,1,2,3,4
    print(i)

for i in range(2, 10, 2):   # 2,4,6,8 (début, fin exclue, pas)
    print(i)

for cle, val in config.items():
    print(cle, val)
```

### while : tant qu'une condition est vraie

```python
tentatives = 0
while tentatives < 3:
    if ping(host):
        break
    tentatives += 1
    time.sleep(5)
```

> **Danger :** boucle infinie si la condition ne devient jamais fausse. Prévoyez toujours une porte de sortie (`break`, compteur, timeout).

```python
# Boucle infinie contrôlée (supervision)
while True:
    verifier_services()
    time.sleep(60)
```

---

## 15. break, continue, else de boucle

```python
for srv in serveurs:
    if srv == "srv-hs":
        continue        # saute ce serveur, passe au suivant
    if ping(srv):
        print(f"{srv} OK")
        break           # sort de la boucle au premier succès
else:
    # Exécuté SEULEMENT si la boucle n'a pas rencontré de break
    print("Aucun serveur joignable")
```

| Instruction | Effet |
|---|---|
| `break` | Sort immédiatement de la boucle |
| `continue` | Saute à l'itération suivante |
| `else` (de boucle) | Exécuté si pas de `break` |

> Le `else` de boucle surprend : pensez "pas de break → else". Très utile pour les recherches ("trouvé ? sinon...").

---

## 16. Comprehensions

Façon concise et rapide de construire des collections.

### Listes en compréhension

```python
carres = [x**2 for x in range(10)]

# Avec filtre
pairs = [x for x in range(20) if x % 2 == 0]

# Équivalent verbeux :
pairs = []
for x in range(20):
    if x % 2 == 0:
        pairs.append(x)
```

### Dict et set en compréhension

```python
longueurs = {srv: len(srv) for srv in serveurs}
ports_uniques = {s["port"] for s in serveurs_list}
```

### Exemple sysadmin : filtrer des lignes de log

```python
erreurs = [l for l in open("/var/log/syslog") if "error" in l.lower()]
ips_suspectes = {l.split()[0] for l in open("access.log") if " 404 " in l}
```

### Règles de bon usage

- 1 compréhension = 1 transformation simple. Au-delà, écrivez une boucle normale.
- Pas de double `for` imbriqué illisible : la lisibilité prime.
- Les générateurs `(x for x in ...)` (parenthèses) ne construisent rien en mémoire — voir section 34.

---

## 17. Fonctions : bases

```python
def ping_host(host, timeout=5):
    """Ping un hôte. Retourne True si joignable.

    Args:
        host: nom ou IP de l'hôte.
        timeout: délai en secondes.
    """
    # ... code ...
    return True
```

### Appels

```python
ping_host("srv-01")                 # timeout=5 par défaut
ping_host("srv-01", timeout=2)       # argument nommé
ping_host(host="srv-01", timeout=2)
```

### return

```python
def diviser(a, b):
    if b == 0:
        return None        # sortie anticipée
    return a / b

resultat = diviser(10, 2)  # 5.0
```

> Une fonction sans `return` retourne `None`. `return` seul = `return None`.

### Docstrings : documentez toujours

Convention : 1ère ligne = résumé impératif, puis détails. Outils comme `help()` et les IDE les affichent.

```python
help(ping_host)   # affiche la docstring
```

---

## 18. Fonctions : *args, **kwargs, paramètres avancés

```python
def journaliser(niveau, *messages, **metadonnees):
    print(f"[{niveau}]", *messages, metadonnees)

journaliser("INFO", "démarrage", "ok", host="srv-01")
# messages = ("démarrage", "ok"), metadonnees = {"host": "srv-01"}
```

### Ordre des paramètres

```python
def f(positionnel, *args, defaut=10, **kwargs):
    ...
```

1. Paramètres positionnels
2. `*args` (tuple des positionnels supplémentaires)
3. Paramètres nommés avec défaut
4. `**kwargs` (dict des nommés supplémentaires)

### Déballage à l'appel

```python
args = ["srv-01", 22]
ping_host(*args)                 # ping_host("srv-01", 22)

opts = {"timeout": 2}
ping_host("srv-01", **opts)     # ping_host("srv-01", timeout=2)
```

### Paramètres positionnels-only et nommés-only (Python 3.8+)

```python
def connecter(host, port, /, *, timeout=5):
    # host, port : positionnels uniquement ; timeout : nommé uniquement
    ...
connecter("srv-01", 22, timeout=2)   # OK
```

---

## 19. lambda, map, filter, sorted

### lambda : petite fonction anonyme

```python
double = lambda x: x * 2
double(21)   # 42
```

Utile comme argument court :

```python
serveurs = [{"nom": "b", "charge": 80}, {"nom": "a", "charge": 30}]
tries = sorted(serveurs, key=lambda s: s["charge"])
# [{'nom': 'a', ...}, {'nom': 'b', ...}]
```

### map / filter

```python
list(map(str.upper, ["a", "b"]))        # ['A', 'B']
list(filter(lambda x: x > 2, [1, 2, 3]))  # [3]

# En général, préférez les comprehensions (plus lisibles) :
[x.upper() for x in ["a", "b"]]
[x for x in [1, 2, 3] if x > 2]
```

### sorted avec key

