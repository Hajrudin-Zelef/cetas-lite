---
id: collect-261001-rattrapage/rattrapage/python-guide-1
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["packaging"]
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [1, 277]
sha256: c9f389134407a282d2a4a8f97077a61982823f9066f701f648840dd419cb49aa
---

# Guide Python complet — De l'installation aux scripts sysadmin

> **Public :** sysadmins, techniciens systèmes & réseaux, curieux de l'automatisation.
> **Angle :** pratique. Moins de théorie abstraite, plus de scripts qui marchent.
> **Version de référence :** Python 3.10+ (3.11/3.12/3.13 recommandés). Les notes "Python 3.10+" indiquent les syntaxes dépendantes de la version.

---

## Sommaire

1. Installation & tooling
2. pip et venv : l'environnement propre
3. Votre premier script
4. Syntaxe de base : indentation, commentaires, variables
5. Types fondamentaux
6. Nombres et opérations
7. Chaînes de caractères
8. Listes
9. Tuples
10. Dictionnaires
11. Sets (ensembles)
12. Conversions de types
13. Conditions (if/elif/else)
14. Boucles (for, while)
15. break, continue, else de boucle
16. Comprehensions
17. Fonctions : bases
18. Fonctions : *args, **kwargs, paramètres avancés
19. lambda, map, filter, sorted
20. Portée des variables (scope)
21. Gestion d'erreurs : try/except
22. Lever ses propres exceptions
23. Modules et imports
24. Packages et `__init__.py`
25. Fichiers : lecture/écriture
26. pathlib : manipuler les chemins
27. os, sys, shutil, subprocess
28. Entrées-sorties avancées (stdin/stdout/stderr)
29. Programmation orientée objet : classes
30. Héritage et polymorphisme
31. Méthodes spéciales (dunder)
32. Dataclasses
33. Décorateurs
34. Générateurs et itérateurs
35. Context managers (`with`)
36. Annotations de type (typing)
37. Expressions régulières (regex)
38. JSON, YAML, CSV
39. datetime : dates et heures
40. argparse : scripts en ligne de commande
41. logging : journaliser proprement
42. requests : HTTP(S) simplifié
43. Réseau : sockets, ping, inventaire
44. sqlite3 : base de données embarquée
45. Tests unitaires : unittest et pytest
46. Débogage : pdb et techniques
47. Asyncio : programmation asynchrone
48. Threading et multiprocessing
49. Performance : mesurer et optimiser
50. Packaging : pyproject.toml
51. Intro à FastAPI
52. Bonnes pratiques & PEP 8
53. Sécurité : écrire du code sûr
54. Scripts sysadmin n°1 : inventaire réseau
55. Scripts sysadmin n°2 : parsing de logs
56. Scripts sysadmin n°3 : sauvegarde et rotation
57. Scripts sysadmin n°4 : surveillance disque et alertes
58. Scripts sysadmin n°5 : requêtes API automatisées
59. Scripts sysadmin n°6 : traitement CSV/Excel
60. Scripts sysadmin n°7 : scan de ports et services
61. Erreurs classiques des débutants (20 pièges)
62. Cas pratiques commentés
63. Checklist "script de production"
64. Pense-bête de poche
65. Glossaire
66. Quiz (10 questions + réponses)
67. Pour aller plus loin

---

## 1. Installation & tooling

### Vérifier l'installation

```bash
python3 --version      # doit afficher 3.10 ou plus
python3 -m pip --version
```

### Installer Python

| OS | Méthode recommandée |
|---|---|
| Debian/Ubuntu | `sudo apt update && sudo apt install python3 python3-pip python3-venv` |
| RHEL/CentOS | `sudo dnf install python3 python3-pip` |
| Windows | Installeur python.org — **cocher "Add python.exe to PATH"** |
| macOS | `brew install python` |

> Sur Debian/Ubuntu récents, le système protège son Python : `pip install` hors venv refuse avec `externally-managed-environment`. C'est normal, utilisez un venv (section 2).

### Éditeurs et IDE

| Outil | Usage |
|---|---|
| VS Code + extension Python | Le plus courant, gratuit, débogueur intégré |
| PyCharm Community | IDE complet, excellent pour les gros projets |
| Vim/Neovim, nano | Édition rapide sur serveur |

### Vérifications rapides après installation

```bash
python3 -c "import sys; print(sys.version)"
python3 -m venv /tmp/testvenv && echo "venv OK"
```

### Checklist installation

- [ ] `python3 --version` >= 3.10
- [ ] `pip` fonctionnel (`python3 -m pip --version`)
- [ ] `venv` disponible (`python3 -m venv --help`)
- [ ] Éditeur configuré avec l'interpréteur correct

---

## 2. pip et venv : l'environnement propre

### Pourquoi un environnement virtuel ?

Chaque projet a ses dépendances (ex. `requests==2.31`). Sans venv, tout s'installe dans le Python système : conflits de versions garantis entre projets.

### Créer et activer un venv

```bash
python3 -m venv ~/venvs/projet1
source ~/venvs/projet1/bin/activate   # Linux/macOS
# Windows : ~\venvs\projet1\Scripts\activate
```

Une fois activé, le prompt affiche `(projet1)`. `pip` installe alors **dans le venv uniquement**.

### Commandes pip essentielles

```bash
pip install requests              # installer un paquet
pip install requests==2.31.0      # version précise
pip install -r requirements.txt   # depuis un fichier
pip freeze > requirements.txt     # figer les versions installées
pip list                          # paquets installés
pip show requests                 # détails d'un paquet
pip install --upgrade requests    # mettre à jour
pip uninstall requests            # désinstaller
```

### requirements.txt : exemple

```text
requests==2.31.0
pyyaml==6.0.1
rich==13.7.0
```

### Quitter / supprimer un venv

```bash
deactivate                        # quitter le venv
rm -rf ~/venvs/projet1            # supprimer = effacer le dossier
```

### Tableau récapitulatif venv

| Action | Commande |
|---|---|
| Créer | `python3 -m venv <nom>` |
| Activer (Linux) | `source <nom>/bin/activate` |
| Activer (Windows) | `<nom>\Scripts\activate` |
| Désactiver | `deactivate` |
| Lister les paquets | `pip list` |
| Figer | `pip freeze > requirements.txt` |

> **Réflexe pro :** un projet = un venv = un `requirements.txt` (ou `pyproject.toml`, section 50). Ne jamais `pip install` en root sur le Python système.

---

## 3. Votre premier script

Créez `hello.py` :

```python
#!/usr/bin/env python3
"""Mon premier script : affiche un message et la version de Python."""

import sys

def main():
    print("Bonjour, sysadmin !")
    print(f"Python {sys.version.split()[0]}")

if __name__ == "__main__":
    main()
```

Exécution :

```bash
chmod +x hello.py
./hello.py
# ou
python3 hello.py
```

### Anatomie d'un script

| Élément | Rôle |
|---|---|
| `#!/usr/bin/env python3` | Shebang : rend le script exécutable directement |
| `"""..."""` | Docstring : documentation du module |
| `def main():` | Point d'entrée logique |
| `if __name__ == "__main__":` | N'exécute `main()` que si le fichier est lancé directement (pas importé) |

> Le bloc `if __name__ == "__main__":` est une convention quasi obligatoire : il permet d'importer vos fonctions dans un autre script sans exécuter tout le fichier.

---

## 4. Syntaxe de base : indentation, commentaires, variables

### L'indentation remplace les accolades

```python
if True:
    print("indenté de 4 espaces")
    print("toujours 4 espaces")
print("hors du if")
```

- **4 espaces** par niveau (convention PEP 8). Pas de tabulations.
- Un mélange espaces/tabulations = `TabError`.

### Commentaires

```python
# Ceci est un commentaire sur une ligne

"""
Ceci est une docstring :
commentaire multi-lignes, utilisé pour documenter
modules, classes et fonctions.
"""
```

### Variables : pas de déclaration de type

```python
hostname = "srv-web-01"   # str
port = 22                 # int
actif = True              # bool
charge = 0.75             # float
```

- Le nom doit commencer par une lettre ou `_`, puis lettres/chiffres/`_`.
- Convention : `snake_case` pour variables et fonctions.
- Constantes : `MAJUSCULES` (convention, pas d'obligation technique).

```python
TIMEOUT = 30          # constante par convention
MAX_RETRIES = 3
```

### Affectations multiples

```python
a, b = 1, 2            # a=1, b=2
x = y = 0              # x=0 et y=0
host, port = "srv", 22 # déballage (unpacking)
```

### print() et f-strings

