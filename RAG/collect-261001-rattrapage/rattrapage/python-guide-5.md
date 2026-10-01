---
id: collect-261001-rattrapage/rattrapage/python-guide-5
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [1289, 1587]
sha256: 0df0f2449174b8196df183ceb280c95cd624139056023d32d66f28b9e071df39
---

# Guide Python complet — De l'installation aux scripts sysadmin

```python
import os
os.getcwd()                 # dossier courant
os.listdir("/tmp")          # contenu (noms simples)
os.environ.get("HOME")      # variable d'environnement
os.environ.get("TIMEOUT", "30")
os.cpu_count()              # nombre de CPU
os.uname()                  # infos système (Linux/macOS)
```

### sys : interpréteur

```python
import sys
sys.argv                    # arguments CLI : ["script.py", "arg1", ...]
sys.exit(1)                 # quitte avec code de retour 1
sys.version_info            # (3, 11, ...)
sys.platform                # 'linux', 'win32', 'darwin'
sys.stderr.write("erreur\n")
```

### shutil : opérations fichiers de haut niveau

```python
import shutil
shutil.copy("a.txt", "b.txt")          # copie fichier
shutil.copytree("src", "dst")          # copie arborescence
shutil.move("a.txt", "archives/")      # déplace/renomme
shutil.rmtree("vieux_dossier")         # supprime récursivement (!)
shutil.disk_usage("/").free            # espace libre en octets
shutil.which("docker")                 # trouve un exécutable dans le PATH
```

### subprocess : lancer des commandes externes

```python
import subprocess

# Méthode moderne et sûre :
res = subprocess.run(
    ["ping", "-c", "2", "10.0.0.1"],
    capture_output=True, text=True, timeout=10,
)
print(res.returncode)   # 0 = succès
print(res.stdout)       # sortie standard

# Vérifier le code de retour automatiquement :
subprocess.run(["systemctl", "is-active", "nginx"], check=True)
# lève CalledProcessError si returncode != 0
```

> **Sécurité :** passez la commande en **liste**, jamais en chaîne avec `shell=True` (injection de commandes — voir section 53).

```python
# DANGEREUX si 'host' vient d'un utilisateur :
subprocess.run(f"ping -c1 {host}", shell=True)

# SÛR :
subprocess.run(["ping", "-c1", host])
```

---

## 28. Entrées-sorties avancées (stdin/stdout/stderr)

```python
import sys

# Lire depuis un pipe : cat log.txt | python3 script.py
for ligne in sys.stdin:
    traiter(ligne)

# Écrire sur stderr (les erreurs ne polluent pas le pipe)
print("Échec du ping", file=sys.stderr)

# Codes de retour : 0 = OK, autre = erreur
sys.exit(0)
sys.exit(2)   # convention : 2 = usage incorrect
```

### Script "filtre" Unix typique

```python
#!/usr/bin/env python3
"""Filtre les lignes contenant ERREUR (usage : cat f.log | python3 filtre.py)."""
import sys

def main():
    for ligne in sys.stdin:
        if "ERREUR" in ligne.upper():
            sys.stdout.write(ligne)

if __name__ == "__main__":
    main()
```

```bash
cat /var/log/syslog | python3 filtre.py > erreurs.txt
echo "code retour : $?"
```

---

## 29. Programmation orientée objet : classes

```python
class Serveur:
    """Représente un serveur supervisé."""

    def __init__(self, hostname, ip, port=22):
        self.hostname = hostname
        self.ip = ip
        self.port = port
        self.en_ligne = None   # inconnu au départ

    def ping(self, timeout=2):
        """Simule un ping. Retourne True/False."""
        # ... vrai code ici ...
        self.en_ligne = True
        return self.en_ligne

    def __str__(self):
        return f"{self.hostname} ({self.ip})"
```

```python
srv = Serveur("web-01", "10.0.0.11")
print(srv)            # web-01 (10.0.0.11)  → via __str__
srv.ping()
```

### Vocabulaire

| Terme | Exemple |
|---|---|
| Classe | `Serveur` (le moule) |
| Instance / objet | `srv` (l'exemplaire) |
| Attribut | `srv.hostname` (donnée) |
| Méthode | `srv.ping()` (fonction de l'objet) |
| `self` | l'instance elle-même (1er paramètre, obligatoire) |
| `__init__` | constructeur, appelé à la création |

### Attributs de classe vs d'instance

```python
class Compteur:
    total = 0              # attribut de CLASSE : partagé
    def __init__(self):
        self.n = 0         # attribut d'INSTANCE : propre à chacune
        Compteur.total += 1
```

---

## 30. Héritage et polymorphisme

```python
class Equipement:
    def __init__(self, nom, ip):
        self.nom = nom
        self.ip = ip

    def decrire(self):
        return f"{self.nom} @ {self.ip}"

class Switch(Equipement):          # Switch hérite d'Equipement
    def __init__(self, nom, ip, nb_ports):
        super().__init__(nom, ip)  # appelle le constructeur parent
        self.nb_ports = nb_ports

    def decrire(self):             # surcharge (override)
        return f"Switch {super().decrire()} - {self.nb_ports} ports"

class Onduleur(Equipement):
    def __init__(self, nom, ip, puissance_kva):
        super().__init__(nom, ip)
        self.puissance_kva = puissance_kva
```

```python
parc = [Switch("sw-01", "10.0.0.2", 48), Onduleur("ups-01", "10.0.0.3", 20)]
for eq in parc:
    print(eq.decrire())   # polymorphisme : même appel, comportements différents
```

### Quand hériter, quand composer ?

- Héritage = relation **"est un"** (un Switch *est un* Équipement).
- Composition = relation **"a un"** (un Serveur *a une* CarteReseau) → souvent préférable, plus souple.

> En pratique sysadmin, on utilise l'OOP avec modération : des classes pour modéliser le parc, des fonctions pour les traitements.

---

## 31. Méthodes spéciales (dunder)

Les méthodes `__x__` permettent d'intégrer vos objets au langage.

```python
class PlageIP:
    def __init__(self, debut, fin):
        self.debut, self.fin = debut, fin

    def __len__(self):                 # len(obj)
        return self.fin - self.debut + 1

    def __contains__(self, ip):        # ip in obj
        return self.debut <= ip <= self.fin

    def __repr__(self):                # représentation "officielle"
        return f"PlageIP({self.debut}, {self.fin})"

    def __eq__(self, autre):           # obj1 == obj2
        return (self.debut, self.fin) == (autre.debut, autre.fin)
```

| Méthode | Déclenchée par |
|---|---|
| `__init__` | création `C(...)` |
| `__str__` | `print(obj)`, `str(obj)` |
| `__repr__` | `repr(obj)`, affichage en console |
| `__len__` | `len(obj)` |
| `__eq__` | `==` |
| `__lt__` | `<` (permet `sorted()`) |
| `__contains__` | `in` |
| `__iter__` | `for x in obj` |

> Définissez toujours `__repr__` pour vos classes métier : le débogage devient 10x plus lisible.

---

## 32. Dataclasses

Pour les classes qui sont surtout des conteneurs de données (Python 3.7+).

```python
from dataclasses import dataclass, field

@dataclass
class Interface:
    nom: str
    ip: str
    masque: str = "255.255.255.0"
    active: bool = True
    tags: list = field(default_factory=list)  # obligatoire pour les mutables !
```

Équivaut à écrire `__init__`, `__repr__`, `__eq__` à la main :

```python
eth0 = Interface("eth0", "10.0.0.11")
print(eth0)
# Interface(nom='eth0', ip='10.0.0.11', masque='255.255.255.0', active=True, tags=[])
```

### Options utiles

```python
@dataclass(frozen=True)     # immutable (comme un namedtuple amélioré)
class Config:
    host: str
    port: int = 22

@dataclass(order=True)      # génère <, <=, >, >= (tri possible)
class Alarme:
    niveau: int
    message: str
```

> **Piège :** ne jamais mettre `tags: list = []` en défaut mutable — toutes les instances partageraient la même liste. Utilisez `field(default_factory=list)`.

---

## 33. Décorateurs

Un décorateur enveloppe une fonction pour ajouter un comportement (log, cache, retry...).

```python
import functools
import time

def chronometrer(fonction):
    @functools.wraps(fonction)   # conserve nom et docstring
    def enveloppe(*args, **kwargs):
        debut = time.perf_counter()
        resultat = fonction(*args, **kwargs)
        duree = time.perf_counter() - debut
        print(f"{fonction.__name__} : {duree:.3f}s")
        return resultat
    return enveloppe

@chronometrer
def sauvegarde(dossier):
    # ... code long ...
    return "ok"
```

### Décorateur avec paramètres : retry automatique

