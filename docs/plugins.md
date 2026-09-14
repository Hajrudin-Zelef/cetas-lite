# Plugins externes

Les **plugins** sont la voie d'extension générique de Cetas : n'importe quel
programme (Python, Node, PowerShell, binaire...) peut exposer des outils à
l'agent. Ils complètent les deux mécanismes existants :

| Mécanisme | Usage | Outils |
|---|---|---|
| MCP (`mcp.json`) | Protocole standard, serveurs externes (stdio/HTTP) | `mcp_<serveur>_<outil>` |
| Custom tools (`tools.json`) | Appels HTTP simples déclarés par l'opérateur | `custom_<outil>` |
| **Plugins** (`plugins/`) | **Scripts locaux ou HTTP**, n'importe quel langage | `plugin_<plugin>_<outil>` |

## Installation

1. Créez un dossier `$CETAS_LITE_HOME/plugins/<nom>/` (le dossier `plugins`
   est créé automatiquement au premier démarrage).
2. Ajoutez un manifeste `plugin.json` (voir ci-dessous).
3. Redémarrez le serveur, ou rechargez sans redémarrer via
   `POST /api/plugins/reload` (bouton « Recharger » dans Réglages → Plugins).
4. Les outils apparaissent sous `plugin_<nom>_<outil>` dans les appels
   d'outils de l'agent. **Chaque exécution demande l'approbation de
   l'utilisateur** (comme `mcp_*` et `custom_*`).

Un exemple complet est fourni dans `examples/plugins/horloge/`
(date/heure + compteur de mots en Python) : copiez le dossier dans
`$CETAS_LITE_HOME/plugins/horloge/` pour l'essayer.

## Manifeste `plugin.json`

```json
{
  "name": "horloge",
  "version": "1.0.0",
  "description": "Date/heure et utilitaires texte.",
  "tools": [
    {
      "name": "maintenant",
      "description": "Renvoie la date et l'heure actuelles.",
      "parameters": {
        "type": "object",
        "properties": { "format": { "type": "string" } }
      },
      "exec": { "command": ["python3", "horloge.py", "maintenant"], "timeout_sec": 10 }
    },
    {
      "name": "appel",
      "description": "Appel HTTP vers un service local.",
      "parameters": { "type": "object" },
      "http": { "url": "http://127.0.0.1:9000/outil", "method": "POST", "timeout_sec": 20 }
    }
  ]
}
```

Règles de validation :
- `name` (plugin et outil) : lettres, chiffres, `_`, `-` ; le nom complet
  `plugin_<plugin>_<outil>` est limité à 64 caractères ;
- le nom du dossier **doit** correspondre au `name` du manifeste ;
- chaque outil déclare **exactement un** de `exec` ou `http` ;
- `timeout_sec` : 0 = 60 s par défaut, maximum 300 s.

## Protocole `exec` (commande locale)

- **Pas de shell** : `command` est un tableau argv exact
  (`["python3", "script.py"]`). Sur Windows, utilisez `["powershell", "-File", "outil.ps1"]`
  ou un `.exe` / `.bat`.
- Les **arguments JSON** de l'appel sont écrits sur **stdin**.
- Le programme répond sur **stdout** :
  - soit du JSON `{"result": ...}` (`result` : texte, nombre, objet...)
    ou `{"error": "..."}` en cas d'échec — **recommandé** ;
  - soit du **texte brut**, utilisé tel quel comme résultat.
- Sortie limitée à 256 Ko (tronquée au-delà) ; `stderr` (8 Ko max) sert
  uniquement aux messages d'erreur si le programme échoue.
- **Répertoire de travail confiné** au dossier du plugin (`dir` optionnel,
  les `..` sont refusés). Variables d'environnement ajoutées :
  `CETAS_PLUGIN_NAME`, `CETAS_PLUGIN_DIR`.
- **Timeout strict** : à expiration, tout l'arbre de processus est tué
  (aucun processus orphelin) et l'appel échoue — via un groupe de processus
  dédié (Unix) ou un Job Object Windows (`KILL_ON_JOB_CLOSE`, démarrage
  suspendu puis assignation avant reprise : aucun petit-enfant ne peut
  échapper).

## Protocole `http`

- `POST` (ou `method` choisi) avec le JSON des arguments en corps,
  `Content-Type: application/json`.
- Réponse 2xx : même protocole que `exec` — JSON `{"result": ...}` /
  `{"error": "..."}` ou texte brut.

## Sécurité

- Les plugins sont installés **manuellement** par l'opérateur (aucun
  téléchargement automatique) ; ne copiez que du code de confiance.
- La commande tourne avec les droits du serveur : ne donnez pas au plugin
  plus de pouvoir que nécessaire.
- Chaque appel d'outil `plugin_*` exige une **approbation explicite**
  dans l'interface avant exécution.

## Diagnostic

- `GET /api/plugins` : liste des plugins chargés et erreurs par plugin
  (un plugin invalide n'empêche pas les autres de charger).
- Logs serveur au démarrage : `plugins charges count=N` / `plugin ignore`.
