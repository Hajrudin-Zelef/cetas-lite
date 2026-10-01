---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-11
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [1520, 1708]
sha256: 5f15a66e3b7a3383ba5ea2fb9f2c6eda03815b7ccf346bf130a353fb89af305d
---

# Concepts : agents IA, agentic, autonomie

Le point le plus important du guide. Version stricte, sans `shell=True` :

```python
import shlex, subprocess

ALLOWED_BINARIES = {
    "df", "free", "uptime", "ls", "cat", "head", "tail", "grep", "wc",
    "ping", "ip", "ss", "systemctl", "journalctl", "whoami", "hostname",
    "uname", "du", "ps", "lsblk",
}
FORBIDDEN_SUBSTR = (";", "|", "&", ">", "<", "`", "$(", "${", "\n", "\r")

def run_shell(command: str, timeout: int = 15) -> dict:
    """Exécute une commande simple. Refuse tout le reste."""
    if any(f in command for f in FORBIDDEN_SUBSTR):
        return {"ok": False, "error": "opérateurs shell interdits (pas de pipe/redirection)"}
    parts = shlex.split(command)
    if not parts:
        return {"ok": False, "error": "commande vide"}
    if parts[0] not in ALLOWED_BINARIES:
        return {"ok": False, "error": f"binaire '{parts[0]}' hors liste blanche"}
    try:
        p = subprocess.run(parts, capture_output=True, text=True, timeout=timeout)
    except subprocess.TimeoutExpired:
        return {"ok": False, "error": f"timeout après {timeout}s"}
    except OSError as e:
        return {"ok": False, "error": f"échec d'exécution : {e}"}
    out = (p.stdout + p.stderr)[-4000:]  # borne la taille réinjectée
    return {"ok": p.returncode == 0, "code": p.returncode, "output": out}
```

Pourquoi c'est sûr (relativement) :
- pas de `shell=True` → pas d'injection via `;`, `$()`, etc. (doublement refusés) ;
- binaire en liste blanche → l'agent ne peut pas lancer `rm`, `dd`, `curl`… ;
- timeout + sortie bornée → pas de pendaison, pas d'explosion de contexte.

Limites assumées : pas de pipes (volontaire dans cette version) ; `systemctl`
sans `sudo` ne fera que du `status` — c'est le but. Pour aller plus loin :
autorise les pipes en découpant toi-même sur `|` et en chaînant les `subprocess`
(sans jamais passer par un shell), ou ajoute des outils dédiés (`restart_service`
avec approbation, voir section 85).

## 82. Étape 3 : les outils fichiers (confinés à WORKDIR)

```python
import fnmatch

SKIP_DIRS = {".git", "node_modules", "__pycache__", ".venv"}

def _inside_workdir(path: str) -> str | None:
    p = os.path.abspath(path)
    if p == WORKDIR or p.startswith(WORKDIR + os.sep):
        return p
    return None

def find_files(root: str, pattern: str = "*", max_results: int = 50) -> dict:
    base = _inside_workdir(root or WORKDIR)
    if base is None:
        return {"ok": False, "error": "hors périmètre autorisé"}
    hits = []
    for dirpath, dirnames, filenames in os.walk(base):
        dirnames[:] = [d for d in dirnames if d not in SKIP_DIRS]
        for f in filenames:
            if fnmatch.fnmatch(f, pattern):
                hits.append(os.path.join(dirpath, f))
                if len(hits) >= max_results:
                    return {"ok": True, "files": hits, "truncated": True}
    return {"ok": True, "files": hits}

def read_file(path: str, max_chars: int = 8000) -> dict:
    p = _inside_workdir(path)
    if p is None:
        return {"ok": False, "error": "hors périmètre autorisé"}
    try:
        with open(p, "r", encoding="utf-8", errors="replace") as fh:
            data = fh.read(max_chars + 1)
    except OSError as e:
        return {"ok": False, "error": str(e)}
    cut = len(data) > max_chars
    return {"ok": True, "content": data[:max_chars], "truncated": cut}

TOOL_IMPLS = {
    "run_shell": lambda a: run_shell(a["command"]),
    "find_files": lambda a: find_files(a.get("root", WORKDIR), a.get("pattern", "*")),
    "read_file": lambda a: read_file(a["path"]),
}
```

Le confinement `WORKDIR` est non négociable : sans lui, `read_file("/etc/shadow")`
est à un appel de fonction du modèle. Ici l'outil refuse avant même d'essayer.

## 83. Étape 4 : la boucle ReAct

```python
class BudgetExceeded(Exception):
    pass

class Budgets:
    def __init__(self, max_steps=25, max_tokens=200_000, timeout_s=600):
        self.max_steps, self.max_tokens = max_steps, max_tokens
        self.deadline = __import__("time").time() + timeout_s
        self.steps = 0
        self.tokens_used = 0

    def check(self):
        import time
        if self.steps >= self.max_steps:
            raise BudgetExceeded(f"étapes {self.steps}/{self.max_steps}")
        if self.tokens_used >= self.max_tokens:
            raise BudgetExceeded(f"tokens {self.tokens_used}/{self.max_tokens}")
        if time.time() >= self.deadline:
            raise BudgetExceeded("timeout global")

def run_agent(goal: str, budgets: Budgets) -> dict:
    messages = [
        {"role": "system", "content": SYSTEM},
        {"role": "user", "content": goal},
    ]
    log_event({"type": "start", "goal": goal})
    try:
        while True:
            budgets.check()
            resp = client.chat.completions.create(
                model=MODEL, messages=messages,
                tools=TOOL_SCHEMAS, temperature=0)
            usage = resp.usage
            if usage:
                budgets.tokens_used += usage.total_tokens
            msg = resp.choices[0].message

            # 1. on fige le message assistant en dict sérialisable
            assistant = {"role": "assistant", "content": msg.content or ""}
            tool_calls = []
            if msg.tool_calls:
                for tc in msg.tool_calls:
                    tool_calls.append({
                        "id": tc.id, "type": "function",
                        "function": {"name": tc.function.name,
                                     "arguments": tc.function.arguments or "{}"}})
                assistant["tool_calls"] = tool_calls
            messages.append(assistant)

            # 2. pas d'appel d'outil -> réponse finale
            if not tool_calls:
                final = msg.content or ""
                log_event({"type": "end", "status": "done"})
                return {"status": "done", "answer": final}

            # 3. exécution des outils, un par un
            for tc in tool_calls:
                name = tc["function"]["name"]
                try:
                    args = json.loads(tc["function"]["arguments"])
                except json.JSONDecodeError:
                    args = {}
                log_event({"type": "tool_call", "tool": name, "args": args})
                if name not in TOOL_IMPLS:
                    result = {"ok": False, "error": f"outil inconnu : {name}"}
                elif name in SENSITIVE and not ask_approval(name, args):
                    result = {"ok": False, "error": "refusé par l'opérateur"}
                else:
                    try:
                        result = TOOL_IMPLS[name](args)
                    except Exception as e:  # noqa: BLE001 — jamais de crash de boucle
                        result = {"ok": False, "error": f"exception outil : {e}"}
                log_event({"type": "tool_result", "tool": name,
                           "ok": bool(result.get("ok"))})
                print(f"[outil] {name} -> ok={result.get('ok')}")
                messages.append({
                    "role": "tool", "tool_call_id": tc["id"],
                    "content": json.dumps(result, ensure_ascii=False)[:6000],
                })
            budgets.steps += 1
    except BudgetExceeded as e:
        log_event({"type": "end", "status": "budget", "reason": str(e)})
        return {"status": "budget_exceeded", "reason": str(e),
                "partial": messages[-1].get("content", "")[:2000]}
```

Points à noter :
- le `try/except` autour de chaque outil : **un outil qui plante ne doit jamais
  tuer la boucle** — l'échec devient une observation comme une autre ;
- `SENSITIVE` (section 85) : ici vide par défaut (`SENSITIVE = set()`), à remplir
  quand tu ajoutes des outils d'écriture ;
- le compteur `steps` s'incrémente par **tour** (pas par appel d'outil) ;
- en cas de budget épuisé : rapport partiel, pas de crash silencieux.

## 84. Étape 5 : journalisation JSONL (à écrire AVANT de tester)

