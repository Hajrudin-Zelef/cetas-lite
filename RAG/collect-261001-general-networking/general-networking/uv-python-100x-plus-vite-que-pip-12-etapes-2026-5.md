---
id: collect-261001-general-networking/general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026-5
title: "macOS et Linux"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: ["2026-06-01"]
keywords: ["apache", "claude", "copilot", "open source", "packaging"]
source: docs/RAG/collect-261001-general-networking/uv-python-100x-plus-vite-que-pip-12-etapes-2026.md
source_anchor: ""
source_lines: [480, 576]
sha256: 2684ac859579a9113cac946cc7a78a80ea0ecc890719624b535fff08312b72be
---

# macOS et Linux

| Commande | Action | 
|---|---|
| `uv init --package nom` | Créer un nouveau projet packagé | 
| `uv add paquet` | Ajouter une dépendance (met à jour le lock) | 
| `uv add --dev paquet` | Ajouter une dépendance de développement | 
| `uv remove paquet` | Retirer une dépendance | 
| `uv run commande` | Exécuter dans l’environnement du projet | 
| `uv lock` | Générer / mettre à jour `uv.lock` | 
| `uv sync --locked` | Installer depuis le lock (CI) | 
| `uv python install 3.13` | Installer une version de Python | 
| `uv python pin 3.13` | Épingler la version du projet | 
| `uvx outil` | Exécuter un outil de façon éphémère | 
| `uv build` | Construire sdist + wheel | 
| `uv publish` | Publier sur PyPI | 
| `uv cache clean` | Vider le cache global | 

## Erreurs fréquentes et pièges à éviter

Voici les pièges les plus courants observés lors de l’adoption de uv, et comment les éviter.

- **Confondre `uv add` et `uv pip install`.** Seul`uv add` met à jour`pyproject.toml` et le lock.`uv pip install` se comporte comme pip et n’assure aucune reproductibilité pour un projet géré.
- **Versionner `.venv` dans Git.** N’ajoutez jamais l’environnement virtuel au dépôt. Versionnez`uv.lock` et`.python-version` , et mettez`.venv` dans`.gitignore` .
- **Oublier `--locked` ou `--frozen` en CI.** Sans ces options, uv peut re-résoudre les dépendances et introduire des versions différentes de celles testées en local.
- **Lancer `python script.py` au lieu de `uv run`.** Vous risquez d’utiliser le mauvais interpréteur ou un environnement global pollué.
- **Mélanger dépendances de production et de développement.** Utilisez`--dev` pour pytest, Ruff et consorts, afin de ne pas les livrer aux utilisateurs.
- **Installer des paquets avec un pip global dans un projet uv.** Cela contourne le lock et brise la cohérence de l’environnement géré.

## Dépannage : 8 problèmes courants et leurs solutions

Cette section rassemble huit situations fréquentes rencontrées en production, avec la marche à suivre.

- **1. `uv: command not found` après installation.** Le binaire est dans`~/.local/bin` ; redémarrez le terminal ou rechargez votre shell (`source ~/.bashrc` ) pour rafraîchir le PATH.
- **2. « No solution found » lors de la résolution.** Deux contraintes de version sont incompatibles. Assouplissez une borne dans`pyproject.toml` ou inspectez le conflit affiché par uv.
- **3. Conflit Git sur `uv.lock`.** Ne fusionnez pas le fichier à la main : relancez`uv lock` pour le régénérer proprement.
- **4. Erreur de certificat SSL en entreprise.** Ajoutez l’option`--native-tls` , ou définissez`SSL_CERT_FILE` pour pointer vers le certificat racine de votre organisation.
- **5. Mauvaise version de Python utilisée.** Vérifiez le fichier`.python-version` et exécutez`uv python pin` pour forcer la bonne version.
- **6. `uv sync` supprime un paquet installé à la main.** C’est normal : l’environnement reflète strictement le lock. Ajoutez la dépendance avec`uv add` au lieu d’un pip manuel.
- **7. Publication qui échoue en 403.** Jeton invalide ou nom de paquet déjà pris. Régénérez`UV_PUBLISH_TOKEN` et testez d’abord sur TestPyPI.
- **8. Build Docker lent.** Copiez`pyproject.toml` et`uv.lock` avant le code source pour profiter du cache de couches, et activez`UV_LINK_MODE=copy` .

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, ces fonctionnalités avancées font passer votre usage de uv à la vitesse supérieure.

### Workspaces pour les monorepos

À la manière de Cargo (l’outil de Rust), uv gère des **workspaces** : plusieurs paquets dans un même dépôt, partageant un unique `uv.lock`. Déclarez-les dans le `pyproject.toml` racine :

```
[tool.uv.workspace]
members = ["packages/*"]
```
### Cache, mode hors ligne et builds figés dans le temps

Le cache global vit dans `~/.cache/uv`. Nettoyez-le avec `uv cache clean` ou élaguez les entrées obsolètes avec `uv cache prune`. Pour des builds strictement reproductibles, l’option `--exclude-newer 2026-06-01` ignore toute version publiée après une date donnée. Enfin, `--offline` force uv à n’utiliser que le cache, utile en environnement isolé.

### Index privés et registres d’entreprise

Pour tirer des paquets depuis un miroir interne, définissez la variable `UV_INDEX_URL` ou déclarez un index dans le `pyproject.toml`. uv gère l’authentification par jetons et prend en charge plusieurs index avec priorités, ce qui couvre les besoins des grandes organisations. Ces mécanismes, combinés à l’assistance d’outils comme Claude Code ou GitHub Copilot, forment une chaîne d’outillage Python particulièrement productive en 2026.

## Foire aux questions (FAQ)

### uv est-il gratuit et open source ?

Oui. uv est distribué sous licence permissive (Apache 2.0 ou MIT) et reste gratuit. Le rachat d’Astral par OpenAI en mars 2026 s’est accompagné d’un engagement public à maintenir cette ouverture, avec la possibilité de forker le projet si besoin.

### Dois-je désinstaller pip ou Poetry pour utiliser uv ?

Non. uv coexiste sans problème avec pip et Poetry. Vous pouvez migrer progressivement, projet par projet, sans big bang. uv respecte les standards de packaging Python, ce qui facilite les allers-retours.

### uv fonctionne-t-il sous Windows ?

Oui, uv est multiplateforme : macOS, Linux et Windows 10 et supérieur sont pris en charge. L’installation se fait via un script PowerShell, et toutes les commandes présentées ici fonctionnent à l’identique.

### uv remplace-t-il conda ?

En partie. uv couvre l’installation de paquets, la gestion des versions de Python et l’isolation des environnements. En revanche, pour les piles scientifiques lourdes qui reposent sur des binaires non-Python complexes, conda reste parfois pertinent, même si uv gère la grande majorité des cas via les wheels de PyPI.

### Où est stocké le cache de uv et comment le vider ?

Le cache global se trouve dans `~/.cache/uv` (personnalisable via `UV_CACHE_DIR`). Videz-le avec `uv cache clean`. C’est ce cache partagé, avec sa déduplication par liens durs, qui explique les gains de vitesse à chaud pouvant atteindre 80 à 115×.

### Le fichier uv.lock est-il vraiment multiplateforme ?

Oui. Le lockfile universel de uv encode les résolutions pour différents systèmes, architectures et versions de Python. Un même `uv.lock` garantit des installations identiques sur macOS, Linux et Windows, ce qu’un `requirements.txt` classique ne permet pas.

### uv gère-t-il les monorepos et projets multi-paquets ?

Oui, grâce aux workspaces de style Cargo. Vous déclarez les membres dans le `pyproject.toml` racine, et uv résout l’ensemble avec un seul lockfile partagé – idéal pour les grandes bases de code.

### Puis-je utiliser uv avec un projet requirements.txt existant ?

Absolument. Importez les dépendances avec `uv add -r requirements.txt`, ou restez en mode pip avec `uv pip install -r requirements.txt`. La migration est progressive et sans rupture.

### Pour aller plus loin : articles liés

*Sources et documentation : documentation officielle de uv, dépôt GitHub astral-sh/uv et page PyPI de uv. Toutes les données de version et de performance reflètent l’état de uv 0.12.7 au 27 août 2026.*
