---
id: collect-261001-ia-llm/ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026-3
title: "La sortie doit afficher Python 3.10.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "claude", "mcp"]
source: docs/RAG/collect-261001-ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026.md
source_anchor: ""
source_lines: [73, 201]
sha256: 7f10e1e3b37ec85a72370fb679f8a7903ed01998bc7c00ec8430f76ce54050ff
---

# La sortie doit afficher Python 3.10.x ou une version supérieure

```
python3 --version
# La sortie doit afficher Python 3.10.x ou une version supérieure
# Exemple de sortie attendue : Python 3.12.4
```
Si la commande renvoie une version antérieure à 3.10 ou une erreur, installez une version récente depuis le site officiel de Python avant de continuer. Une fois la version confirmée, créez le dossier du projet et un environnement virtuel pour isoler les dépendances du serveur du reste de votre système.

```
mkdir serveur-mcp-git-journal && cd serveur-mcp-git-journal
python3 -m venv .venv
source .venv/bin/activate
# Sous Windows (PowerShell) :
# .venv\Scripts\Activate.ps1
```
Terminez par la vérification de Node.js et de Claude Code, tous deux utilisés plus loin, respectivement pour l’Inspecteur MCP et pour la connexion finale du serveur.

```
node --version
# Exemple de sortie attendue : v20.11.0
claude --version
# Exemple de sortie attendue : claude-code/1.x.x darwin-arm64 node-v20.11.0
```
## Étapes 4 à 6 : installer le kit de développement MCP et créer le squelette du serveur

**Étape 4** : installez le kit de développement officiel MCP pour Python, en incluant les outils en ligne de commande utiles au débogage.

```
pip install "mcp[cli]"
# Avec uv, en alternative plus rapide :
# uv add "mcp[cli]"
```
**Étapes 5 et 6** : créez un fichier `server.py` et posez le squelette minimal du serveur, avant d’y ajouter les outils proprement dits à l’étape suivante. Ce squelette suffit déjà à démarrer un serveur MCP valide, même s’il n’expose encore rien d’utile.

```
# server.py
from mcp.server.fastmcp import FastMCP
mcp = FastMCP("git-journal")
if __name__ == "__main__":
    mcp.run()
```
La classe `FastMCP` gère pour vous la partie protocolaire : sérialisation JSON-RPC, annonce des capacités du serveur au client, boucle d’écoute sur l’entrée standard. Le nom passé en argument, ici `git-journal`, sert d’identifiant unique lorsque plusieurs serveurs sont déclarés en parallèle dans la configuration de Claude Code, un point qui revient dans la section dépannage.

## Étapes 7 à 9 : écrire les outils et la ressource du serveur

**Étape 7** : ajoutez l’outil `derniers_commits`, qui interroge l’historique Git du dépôt courant via `subprocess`. Notez l’usage d’une liste d’arguments plutôt que d’une chaîne de commande unique : cela évite tout risque d’injection shell, même si les paramètres d’entrée venaient à contenir des caractères inattendus.

**Étape 8** : ajoutez l’outil `chercher_todo`, qui parcourt les fichiers sources à la recherche de commentaires `TODO` et `FIXME`.

**Étape 9** : ajoutez la ressource `git://status`, qui expose l’état courant du dépôt. Contrairement à un outil, une ressource n’attend pas nécessairement de paramètres : elle représente une donnée que l’agent peut consulter à la demande.

```
import subprocess
from mcp.server.fastmcp import FastMCP
mcp = FastMCP("git-journal")
@mcp.tool()
def derniers_commits(nombre: int = 5) -> str:
    """Retourne les N derniers commits du dépôt Git courant."""
    resultat = subprocess.run(
        ["git", "log", f"-{nombre}", "--pretty=format:%h | %an | %ad | %s", "--date=short"],
        capture_output=True, text=True, check=True
    )
    return resultat.stdout or "Aucun commit trouvé."
@mcp.tool()
def chercher_todo(dossier: str = ".") -> str:
    """Recherche les commentaires TODO et FIXME dans le dossier donné."""
    resultat = subprocess.run(
        ["grep", "-rn", "-E", "TODO|FIXME", dossier,
         "--include=*.py", "--include=*.js", "--include=*.ts"],
        capture_output=True, text=True
    )
    return resultat.stdout or "Aucun TODO ni FIXME trouvé."
@mcp.resource("git://status")
def statut_git() -> str:
    """Expose l'état courant du dépôt (git status --short)."""
    resultat = subprocess.run(
        ["git", "status", "--short"],
        capture_output=True, text=True
    )
    return resultat.stdout or "Aucune modification en attente."
```
La docstring placée sous chaque fonction n’a rien de décoratif : c’est elle que Claude Code lit pour décider quand appeler tel ou tel outil, en complément du nom de la fonction et des types déclarés sur ses paramètres. Un outil mal documenté reste techniquement fonctionnel, mais l’agent peine à deviner dans quel contexte l’utiliser, ce qui se traduit par des appels inutiles ou, à l’inverse, par un outil jamais sollicité alors qu’il aurait été pertinent.

## Étapes 10 et 11 : tester le serveur avec l’Inspecteur MCP

**Étape 10** : avant de connecter quoi que ce soit à Claude Code, testez le serveur isolément avec l’Inspecteur MCP, un outil officiel distribué en paquet npm qui ouvre une interface web locale pour appeler chaque outil manuellement.

```
npx @modelcontextprotocol/inspector python server.py
# L'Inspecteur ouvre une interface web locale, généralement sur
# http://localhost:6274, pour tester chaque outil manuellement
```
Dans cette interface, les deux outils et la ressource déclarés à l’étape précédente doivent apparaître dans une liste, avec la possibilité de renseigner leurs paramètres et de lancer un appel de test. C’est le moment de vérifier que `derniers_commits` renvoie bien un historique lisible et que `git://status` reflète l’état réel du dépôt, avant d’introduire la couche supplémentaire que représente Claude Code.

**Étape 11** : corrigez les erreurs signalées par l’Inspecteur avant d’aller plus loin. L’erreur la plus fréquente à ce stade concerne un binaire manquant dans l’environnement d’exécution.

```
# Erreur fréquente si git n'est pas accessible depuis l'environnement :
# FileNotFoundError: [Errno 2] No such file or directory: 'git'
# Vérifiez que git est bien installé au niveau du système
which git
# Sous Windows : where git
```
## Étapes 12 et 13 : connecter le serveur à Claude Code

**Étape 12** : une fois le serveur validé dans l’Inspecteur, déclarez-le auprès de Claude Code. Deux approches équivalentes existent : la commande dédiée, ou l’édition directe du fichier de configuration du projet.

`claude mcp add git-journal -- python /chemin/complet/vers/server.py`
Cette commande ajoute automatiquement l’entrée correspondante dans le fichier `.mcp.json` à la racine du projet. Vous pouvez aussi l’écrire vous-même directement :

```
{
  "mcpServers": {
    "git-journal": {
      "command": "python",
      "args": ["/chemin/complet/vers/server.py"]
    }
  }
}
```
Utilisez systématiquement un chemin absolu plutôt que relatif dans cette configuration : Claude Code peut être lancé depuis un dossier de travail différent de celui où vit réellement le fichier `server.py`, ce qui casse silencieusement la connexion si le chemin est relatif.

**Étape 13** : ouvrez une nouvelle session Claude Code dans le dossier du projet et vérifiez que le serveur apparaît bien comme actif, avant de lui poser une première question qui mobilise les outils que vous venez d’écrire.

À ce stade, Claude Code doit lister `git-journal` parmi ses serveurs MCP actifs au démarrage. Si ce n’est pas le cas, fermez complètement la session et relancez-la : la configuration n’est lue qu’au démarrage de l’agent, pas en cours de session.

## Exemple de projet complet : le code source intégral du serveur

Voici le fichier `server.py` complet, tel qu’il doit se présenter une fois les treize étapes suivies. Vous pouvez le copier tel quel dans un nouveau projet pour reproduire le tutoriel de bout en bout.

