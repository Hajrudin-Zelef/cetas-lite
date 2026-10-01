---
id: collect-261001-rattrapage/rattrapage/python-guide-2
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [278, 615]
sha256: 4f589eefbf181d6c0572c7f2c4f652aa931e8ca1d8feb55052cfd70bcf8d3882
---

# Guide Python complet — De l'installation aux scripts sysadmin

```python
nom = "srv-web-01"
print("Serveur :", nom)              # virgule = espace auto
print(f"Serveur : {nom}")            # f-string (Python 3.6+)
print(f"Port + 1 = {22 + 1}")        # expressions autorisées
print(f"Charge : {0.756:.1%}")       # formatage : 75,6%
```

### input() : lire au clavier

```python
reponse = input("Nom du serveur ? ")   # retourne toujours une str
port = int(input("Port ? "))           # conversion manuelle
```

---

## 5. Types fondamentaux

| Type | Exemple | Description |
|---|---|---|
| `int` | `42`, `-7` | Entiers, précision illimitée |
| `float` | `3.14`, `1e-3` | Nombres décimaux |
| `str` | `"texte"` | Chaînes (immutables) |
| `bool` | `True`, `False` | Booléens (attention à la majuscule) |
| `None` | `None` | Absence de valeur (`NoneType`) |
| `list` | `[1, 2]` | Liste modifiable |
| `tuple` | `(1, 2)` | Tuple immutable |
| `dict` | `{"a": 1}` | Dictionnaire clé→valeur |
| `set` | `{1, 2}` | Ensemble sans doublons |

```python
type(42)          # <class 'int'>
isinstance("x", str)   # True — préférez isinstance() à type() ==
```

### Valeurs "fausses" (falsy)

```python
bool(0), bool(0.0), bool(""), bool([]), bool({}), bool(None), bool(False)
# → tous False
```

Tout le reste est "vrai" (truthy). Utile pour tester rapidement :

```python
resultats = []
if not resultats:          # liste vide = falsy
    print("Aucun résultat")
```

---

## 6. Nombres et opérations

```python
7 + 3    # 10  addition
7 - 3    # 4   soustraction
7 * 3    # 21  multiplication
7 / 3    # 2.333... division (toujours float)
7 // 3   # 2   division entière
7 % 3    # 1   modulo (reste)
7 ** 3   # 343 puissance
-7       # négation
abs(-7)  # 7   valeur absolue
```

### Opérateurs d'affectation combinés

```python
compteur = 0
compteur += 1   # compteur = compteur + 1
compteur *= 2
```

### Fonctions numériques utiles

```python
round(3.14159, 2)   # 3.14
min(4, 1, 9)        # 1
max(4, 1, 9)        # 9
sum([1, 2, 3])      # 6
divmod(7, 3)        # (2, 1) → (quotient, reste)
pow(2, 10)          # 1024
```

### Précision des flottants (piège classique)

```python
0.1 + 0.2           # 0.30000000000000004 (!)
round(0.1 + 0.2, 2) # 0.3
```

> Pour de la monnaie ou des calculs exacts : module `decimal`.

```python
from decimal import Decimal
Decimal("0.1") + Decimal("0.2")   # Decimal('0.3')
```

### Conversions

```python
int("42")      # 42
int(3.9)       # 3 (tronque, n'arrondit pas)
float("3.14")  # 3.14
str(42)        # "42"
hex(255)       # '0xff'
bin(10)        # '0b1010'
```

---

## 7. Chaînes de caractères

### Bases

```python
s = "Bonjour"
s2 = 'Bonjour'          # ' ou " : équivalent
s3 = """Texte
sur plusieurs
lignes"""
```

### Indexation et slicing

```python
s = "srv-web-01"
s[0]      # 's'   premier caractère
s[-1]     # '1'   dernier caractère
s[0:3]    # 'srv' slice [début:fin[ (fin exclue)
s[4:]     # 'web-01'
s[:3]     # 'srv'
s[::-1]   # '10-bew-vrs' inversion
```

### Méthodes essentielles

```python
"  srv-01  ".strip()          # 'srv-01' (retire espaces)
"srv-01".upper()              # 'SRV-01'
"srv-01".lower()              # 'srv-01'
"srv-web-01".split("-")       # ['srv', 'web', '01']
"-".join(["srv", "01"])       # 'srv-01'
"srv-01".replace("srv", "db") # 'db-01'
"srv-01".startswith("srv")    # True
"log.txt".endswith(".txt")    # True
"ip=10.0.0.1".find("=")       # 2 (position, -1 si absent)
"a,b,c".split(",")            # ['a', 'b', 'c']
```

### f-strings : le formatage moderne

```python
host, port = "srv-01", 8080
f"{host}:{port}"              # 'srv-01:8080'
f"{port:05d}"                 # '08080' (zéro-padding)
f"{0.8567:.1%}"               # '85.7%'
f"{1234567:,}"                # '1,234,567'
f"{host!r}"                   # repr() : 'srv-01' avec guillemets
```

### Concaténation et répétition

```python
"srv" + "-" + "01"   # 'srv-01'
"ab" * 3             # 'ababab'
```

> Les `str` sont **immutables** : chaque opération crée une nouvelle chaîne. Pour construire un gros texte en boucle, accumulez dans une liste puis `""join(...)`.

```python
# Lent (crée N chaînes) :
texte = ""
for i in range(10000):
    texte += str(i)

# Rapide :
morceaux = [str(i) for i in range(10000)]
texte = "".join(morceaux)
```

---

## 8. Listes

```python
serveurs = ["srv-web-01", "srv-web-02", "srv-db-01"]
vide = []
mixte = [1, "deux", 3.0]      # possible mais déconseillé
```

### Accès et modification

```python
serveurs[0]          # 'srv-web-01'
serveurs[-1]         # dernier
serveurs[1] = "srv-web-03"   # modification (mutable !)
serveurs[0:2]        # ['srv-web-01', 'srv-web-03']
```

### Méthodes principales

```python
serveurs.append("srv-cache-01")   # ajoute à la fin
serveurs.insert(0, "srv-lb-01")   # insère à la position 0
serveurs.extend(["a", "b"])       # fusionne une autre liste
serveurs.remove("srv-lb-01")      # retire par valeur (1ère occurrence)
dernier = serveurs.pop()          # retire et retourne le dernier
premier = serveurs.pop(0)         # retire et retourne l'index 0
serveurs.sort()                   # tri en place
serveurs.reverse()                # inverse en place
serveurs.count("a")               # occurrences
serveurs.index("srv-db-01")       # position (ValueError si absent)
serveurs.clear()                  # vide la liste
```

### Fonctions utiles sur les listes

```python
len(serveurs)                  # taille
"srv-db-01" in serveurs        # True (test d'appartenance)
sorted(serveurs)               # trié, sans modifier l'original
reversed(serveurs)             # itérateur inversé
```

### Copier une liste (piège !)

```python
a = [1, 2, 3]
b = a          # b POINTE vers la même liste ! Modifier b modifie a.
b = a.copy()   # vraie copie (superficielle)
b = a[:]       # équivalent
import copy
b = copy.deepcopy(a)   # copie profonde (listes imbriquées)
```

### Parcours avec index

```python
for i, srv in enumerate(serveurs):
    print(f"{i}: {srv}")

for srv in serveurs:
    print(srv)
```

---

## 9. Tuples

Tuples = listes **immutables**. Idéals pour des données fixes (coordonnées, paires host/port).

```python
paire = ("srv-01", 22)
paire[0]          # 'srv-01'
paire[0] = "x"    # TypeError : immutable

# Tuple à 1 élément : la virgule est obligatoire
t = (42,)         # tuple
pas_tuple = (42)  # int !
```

### Déballage (unpacking)

```python
host, port = ("srv-01", 22)
a, b, *reste = (1, 2, 3, 4, 5)   # a=1, b=2, reste=[3, 4, 5]

# Échange de variables sans temporaire
a, b = b, a
```

### Pourquoi des tuples ?

- Clés de dictionnaire possibles (les listes ne le sont pas).
- Retour multiple de fonction : `return host, port` retourne un tuple.
- Protection contre la modification accidentelle.

---

## 10. Dictionnaires

```python
srv = {"hostname": "srv-web-01", "ip": "10.0.0.11", "port": 22}
vide = {}
```

### Accès

```python
srv["hostname"]              # 'srv-web-01'
srv["os"]                    # KeyError !
srv.get("os")                # None (pas d'erreur)
srv.get("os", "linux")       # 'linux' (valeur par défaut)
```

### Modification

```python
srv["os"] = "debian"         # ajout / modification
srv["port"] = 2222
del srv["port"]              # suppression
val = srv.pop("ip")          # retire et retourne
srv.update({"ram": "16G"})   # fusion
srv.setdefault("os", "linux")# n'écrase pas si existe déjà
```

### Parcours

```python
for cle in srv:                    # clés
    print(cle)
for cle, valeur in srv.items():    # paires
    print(f"{cle} = {valeur}")
for valeur in srv.values():        # valeurs
    print(valeur)
```

### Méthodes et tests

```python
"hostname" in srv      # True
len(srv)               # nombre de paires
list(srv.keys())       # ['hostname', 'ip', 'port']
```

### dict en compréhension (Python 3.x)

```python
carres = {x: x**2 for x in range(5)}   # {0: 0, 1: 1, 2: 4, ...}
```

> Depuis Python 3.7, les dicts **conservent l'ordre d'insertion**. On peut s'y fier.

---

