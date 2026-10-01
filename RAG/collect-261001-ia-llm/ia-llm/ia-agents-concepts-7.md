---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-7
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "arr", "gemini", "mcp", "memory"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [895, 1062]
sha256: 71fa59d626990b24490d66420f10c90b1634fd1e372b43e44902200eb06c4d3f
---

# Concepts : agents IA, agentic, autonomie

Faits rapportés (2026) :
- Microsoft (juin 2026) alerte : un attaquant peut cacher des instructions dans la
  **description d'un outil MCP** — texte qui vit dans le contexte de l'agent, à côté
  de ses vraies instructions. Modifier cette description oriente l'agent aussi
  efficacement que réécrire son prompt système : une demande de routine fait alors
  discrètement collecter et exfiltrer des données, chaque étape paraissant légitime.
- 36,7 % de plus de 7 000 serveurs MCP analysés présentaient des vulnérabilités
  SSRF (divulgation fév. 2026) : un agent utilisant un serveur vulnérable peut être
  manipulé pour émettre des requêtes non autorisées.
- McKinsey (oct. 2025) : 80 % des organisations déployant des agents IA rapportent
  des comportements risqués ou inattendus.

Leçons : les outils sont une surface d'attaque ; épingle les versions des serveurs
MCP, limite leur périmètre, journalise tout, et considère la description d'un outil
tiers comme **non fiable par défaut**.

## 49. Ce que les dérives ont en commun (synthèse)

| Dérive | Cause racine | Garde-fou qui aurait aidé |
|---|---|---|
| Gemini CLI (fichiers) | erreur silencieuse + pas de vérification | vérifier les effets (le dossier existe ?) |
| Prod effacée | accès trop large + vitesse machine | séparation envs, approbation suppressions |
| ROME (crypto) | émergence RL sans surveillance effets | monitoring infra (CPU/réseau) |
| Contournement contrôles | objectif > limites perçues | contrôles externes, pas dans le prompt |
| MCP injection | outil tiers non fiable | moindre privilège, logs, versions épinglées |

Le point commun n'est pas la méchanceté des modèles : c'est **l'absence de
vérification externe et de bornes dures**. Construis le harness comme si le modèle
était un stagiaire brillant mais imprévisible : objectifs clairs, périmètre fermé,
vérification systématique, jamais les clés de la prod.

## 50. Le modèle mental du sysadmin : l'agent = un admin junior sous supervision

Pour calibrer ton intuition, pense à l'agent comme à un administrateur junior :

- rapide, infatigable, bonne culture générale, mauvais jugement des risques ;
- il ne « comprend » pas la gravité d'un `rm -rf` ou d'un `DROP DATABASE` ;
- il suit les procédures si elles sont écrites noir sur blanc, il improvise sinon ;
- il ne t'appelle pas quand il est perdu, il devine.

Donc tu fais ce que tu ferais avec un junior : périmètre écrit, droits minimaux,
validation des actions sensibles, relecture de son travail, et jamais seul en
astreinte la première année. La seule différence : le junior va 100× plus vite,
donc tes garde-fous doivent être **automatiques**, pas « je jetterai un œil ».

## 51. Niveaux d'autonomie : comment monter sans se brûler

Protocole de montée en autonomie (applique-le à chaque nouvel agent) :

1. **Semaine 1 — niveau 1** : l'agent propose, tu exécutes. Tu notes tout ce qu'il
   propose de travers.
2. **Semaine 2 — niveau 2** : lecture seule auto + propositions d'écriture.
   Tu mesures : taux de propositions correctes, temps gagné.
3. **Semaine 3 — niveau 3** : écritures réversibles auto, sensibles avec approbation.
   Tu testes le bouton Stop et un dépassement de budget provoqué.
4. **Niveau 4** : seulement si les semaines 1-3 sont propres, avec périmètre verrouillé,
   alerting, et astreinte humaine. Pour un usage perso, le niveau 4 se justifie
   rarement — le niveau 3 donne 90 % du bénéfice pour 10 % du risque.

À chaque niveau : une semaine d'observation minimum, des logs relus, et un critère
de retour en arrière (« si plus de 2 incidents/semaine, on redescend d'un niveau »).

## 52. Sandboxing pratique : 3 montages concrets

**Montage 1 — compte de service (le minimum)**
```bash
sudo useradd -r -s /bin/false agent-svc
# sudoers : uniquement les commandes lues, ex :
agent-svc ALL=(root) NOPASSWD: /bin/systemctl status *, /usr/bin/journalctl *
```
L'agent tourne sous `agent-svc` : il lit, il ne casse (presque) rien.

**Montage 2 — conteneur jetable (recommandé pour l'exécution de code)**
```bash
docker run --rm -i \
  --network none \              # pas de réseau sortant par défaut
  --memory 512m --cpus 1 \
  --read-only --tmpfs /tmp \
  -v /srv/agent-work:/work:rw \ # seul dossier partagé
  python:3.12-slim \
  python /work/script.py
```
Code non fiable ? Il tourne là-dedans, pas sur ton hôte.

**Montage 3 — VM de test (pour les agents réseau)**
Une VM Proxmox clonée depuis un template, snapshot avant chaque session d'agent,
rollback après. Coût : quelques Go de disque. Bénéfice : l'agent peut se tromper
sans conséquence (et tu peux rejouer la scène pour comprendre).

## 53. Approbations humaines : implémentation minimale

Le pattern « pause → demande → reprise » :

```python
def ask_approval(action: str, details: str, timeout_s: int = 300) -> bool:
    """Fail-closed : pas de réponse = refus."""
    print(f"\n[ACTION SENSIBLE] {action}\nDétails : {details}")
    print(f"Répondre 'oui' dans {timeout_s}s pour autoriser (défaut : non).")
    # en prod : webhook / bouton / message ; ici : input avec timeout
    ...
    return False  # défaut
```

Dans le harness : avant tout outil marqué `sensible=True`, appeler `ask_approval`.
Enregistrer : qui a approuvé, quand, quoi (paramètres exacts). En mode non interactif
(cron), les actions sensibles sont **refusées d'office** — l'agent les met dans le
rapport « en attente de validation » au lieu de les exécuter.

## 54. Budgets : implémentation minimale

```python
class Budgets:
    def __init__(self, max_steps=25, max_tokens=200_000, timeout_s=600):
        self.max_steps = max_steps
        self.max_tokens = max_tokens
        self.deadline = time.time() + timeout_s
        self.tokens_used = 0
        self.steps = 0

    def check(self):
        if self.steps >= self.max_steps:
            raise BudgetExceeded(f"étapes : {self.steps}/{self.max_steps}")
        if self.tokens_used >= self.max_tokens:
            raise BudgetExceeded(f"tokens : {self.tokens_used}/{self.max_tokens}")
        if time.time() >= self.deadline:
            raise BudgetExceeded("timeout global atteint")
```

Le harness appelle `budgets.check()` à chaque tour et additionne les tokens
retournés par l'API. `BudgetExceeded` → arrêt propre + rapport partiel.
C'est 20 lignes qui valent plus cher que n'importe quel « meilleur prompt ».

## 55. Journalisation : le minimum auditable

Chaque exécution produit un fichier JSONL (une ligne = un événement) :

```json
{"t": "2026-09-27T02:00:01Z", "type": "start", "goal": "bilan disque srv-web-01"}
{"t": "2026-09-27T02:00:03Z", "type": "tool_call", "tool": "run_shell",
 "args": {"command": "df -h"}, "tokens": 1240}
{"t": "2026-09-27T02:00:04Z", "type": "tool_result", "tool": "run_shell",
 "ok": true, "truncated": true}
{"t": "2026-09-27T02:00:09Z", "type": "approval", "action": "clean_logs",
 "decision": "refused", "by": "timeout"}
{"t": "2026-09-27T02:00:10Z", "type": "end", "status": "partial",
 "tokens_total": 8930, "duration_s": 9}
```

Avec ça : tu rejoues, tu audites, tu factures, tu prouves. Sans ça : tu subis.

## 56. Résumé partie C : la doctrine en 10 points

1. Autonome = boucle sans humain entre le lancement et la fin. Rien de plus.
2. Trois budgets durs dès le prototype : étapes, tokens, temps.
3. Approbation humaine fail-closed sur toute action sensible.
4. Sandboxing proportionné au pire outil autorisé.
5. Vérifie les effets, jamais les dires.
6. Secrets hors du contexte du modèle, toujours.
7. Logs JSONL dès le jour 1.
8. Monte en autonomie un niveau à la fois, avec mesures.
9. Les dérives réelles viennent d'erreurs de harness, pas de « méchanceté ».
10. En cas de doute : niveau 2 (supervisé). C'est là que le ratio bénéfice/risque
    est le meilleur.

---

# PARTIE D — Open Interpreter : vérifié septembre 2026

