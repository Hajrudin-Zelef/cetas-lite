---
id: collect-261001-ia-llm/ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026-4
title: "La sortie doit afficher Python 3.10.x ou une version supérieure"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "claude", "mcp"]
source: docs/RAG/collect-261001-ia-llm/serveur-mcp-pour-claude-code-tutoriel-2026.md
source_anchor: ""
source_lines: [202, 283]
sha256: c93aac2a3d294a77860eefc514b9f81250278d82a95ecb0f7d744a84cc5b1912
---

# La sortie doit afficher Python 3.10.x ou une version supérieure

```
# server.py
# Serveur MCP "git-journal" : expose l'activité récente
# d'un dépôt Git à un client compatible MCP (Claude Code, Cursor, etc.)
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
if __name__ == "__main__":
    mcp.run()
```
Ce fichier tient volontairement en une quarantaine de lignes, mais la structure se généralise sans effort : chaque nouvel outil est une fonction supplémentaire décorée avec `@mcp.tool()`, chaque nouvelle ressource suit le même principe avec `@mcp.resource()`. Une évolution naturelle de ce serveur consisterait à ajouter un outil `derniers_contributeurs` agrégeant les auteurs de commits sur une période donnée, ou une ressource exposant directement le contenu du fichier `CLAUDE.md` du projet.

## À quoi ressemble une session Claude Code avec MCP actif : exemples de sortie

Voici un exemple simplifié de ce qu’affiche le terminal une fois le serveur `git-journal` connecté, utile pour vérifier que votre propre installation se comporte normalement.

```
$ claude
Claude Code v1.x — connecté en tant que [email protected]
Serveurs MCP actifs : git-journal
> Résume les 5 derniers commits et signale s'il reste des TODO non traités
● Appel de l'outil derniers_commits(nombre=5)
● Appel de l'outil chercher_todo(dossier=".")
● Résultat : 5 commits résumés, 3 occurrences de TODO trouvées
  dans src/api/users.py (lignes 12, 47 et 88)
```
Si votre terminal n’affiche aucune mention de serveur MCP actif au démarrage, ou si l’agent répond sans jamais appeler les outils déclarés alors que la question s’y prête clairement, reportez-vous à la section dépannage : cela indique le plus souvent un chemin de configuration incorrect ou une session lancée depuis le mauvais dossier.

## Les erreurs courantes à éviter avec un serveur MCP

Ces pièges reviennent le plus souvent chez les développeurs qui construisent leur premier serveur MCP. La plupart se corrigent en quelques secondes une fois identifiés, mais ils suffisent à décourager un premier essai si personne ne les a signalés à l’avance, un constat déjà observé sur d’autres outils de codage IA et qui se vérifie tout autant sur la construction d’un serveur MCP maison.

- **Oublier d’activer l’environnement virtuel avant de lancer le serveur.** Claude Code lance la commande déclarée dans`.mcp.json` telle quelle, sans activer automatiquement un environnement virtuel Python à votre place.
- **Utiliser un chemin relatif dans la configuration.** Le dossier de travail au démarrage de l’agent ne correspond pas toujours à celui attendu ; un chemin absolu évite ce piège une fois pour toutes.
- **Négliger les docstrings des outils.** Sans description claire, l’agent choisit moins bien quand appeler un outil, ou l’ignore complètement.
- **Donner un accès trop large à un outil.** Exposer une fonction du type “exécute n’importe quelle commande shell” transforme un serveur pratique en risque de sécurité majeur, en particulier sur un dépôt partagé.
- **Ignorer la gestion des erreurs de `subprocess`.** Sans capture propre des erreurs, une commande Git absente ou un dépôt invalide fait planter le serveur au lieu de renvoyer un message exploitable.
- **Confondre outil et ressource.** Un outil doit modéliser une action paramétrable, une ressource une donnée consultable ; les mélanger complique inutilement la conception du serveur.
- **Ne jamais tester avec l’Inspecteur avant de connecter le serveur à Claude Code.** Sauter cette étape revient à déboguer à l’aveugle, à travers une couche supplémentaire qui masque l’origine réelle d’un problème.

## Dépannage : les problèmes les plus fréquents et leurs solutions

Voici les incidents les plus courants lors de la construction et de la connexion d’un serveur MCP, avec la cause probable et la marche à suivre. Avant de chercher plus loin dans cette liste, un réflexe simple règle une bonne partie des cas : fermer complètement la session Claude Code et le terminal, puis les rouvrir tous les deux, pour éviter de travailler avec une configuration ou un chemin resté en cache depuis avant la dernière modification du serveur.

- **« ModuleNotFoundError: No module named ‘mcp’ ».** L’environnement virtuel n’est pas activé au moment de lancer le serveur, ou la commande déclarée dans`.mcp.json` pointe vers un interpréteur Python différent de celui où le paquet a été installé.
- **Le serveur ne répond pas dans l’Inspecteur.** Vérifiez que`mcp.run()` reste bien la dernière instruction du fichier et qu’aucune erreur silencieuse n’interrompt le script avant d’atteindre cette ligne.
- **Claude Code ne détecte pas le serveur configuré.** Vérifiez le chemin absolu déclaré, puis redémarrez complètement la session : la configuration n’est relue qu’au lancement de l’agent.
- **« FileNotFoundError » pour git ou grep.** Le binaire n’est pas accessible depuis l’environnement d’exécution, un cas fréquent sous Windows sans Git Bash ni WSL installé.
- **Les outils apparaissent mais échouent silencieusement.** Ajoutez une gestion explicite des erreurs autour de chaque appel`subprocess` plutôt que de laisser une exception remonter sans contexte.
- **Conflit de nom entre deux serveurs MCP déclarés.** Deux entrées portant le même identifiant dans`mcpServers` entrent en conflit ; renommez l’une des deux pour lever l’ambiguïté.
- **Permission refusée lors de l’exécution du script.** Vérifiez les droits d’exécution du fichier ou invoquez-le explicitement via l’interpréteur Python plutôt que directement.
- **Le serveur fonctionne en local mais pas une fois déployé à distance.** Le transport HTTP/SSE nécessite une configuration réseau et une authentification distinctes du mode stdio local, détaillées dans la section suivante.
- **Lenteur perçue lors de l’appel d’un outil.** Un outil qui parcourt un dépôt volumineux sans limite peut ralentir sensiblement une session ; bornez le périmètre analysé, par exemple en limitant le nombre de commits ou de fichiers traités.
- **La sortie d’un outil dépasse ce que l’agent peut exploiter utilement.** Tronquez ou paginez la sortie renvoyée par vos fonctions plutôt que de renvoyer un historique Git entier sans limite.

## Astuces avancées : serveurs distants, sécurité et écosystème de serveurs prêts à l’emploi

