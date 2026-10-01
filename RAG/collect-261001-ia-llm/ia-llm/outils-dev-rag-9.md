---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-9
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "Huawei", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [1625, 1835]
sha256: 451e1a3cc4cb929bceb6aa960bf34675bf02bf1bac40b888cdd1737b5c574d02
---

# Trace complète (le meilleur outil de debug Playwright) :
context = browser.new_context()
context.tracing.start(screenshots=True, snapshots=True)
# ... tes actions ...
context.tracing.stop(path="debug/trace.zip")
# puis : playwright show-trace debug/trace.zip
```

**Codegen** — génère du code en te regardant naviguer :

```bash
playwright codegen https://example.com
# Clique dans le navigateur ouvert : le code Python s'écrit tout seul.
# Parfait pour découvrir les sélecteurs d'un site inconnu.
```

## 52. Télécharger des fichiers (tes PDF de CLI !)

```python
with page.expect_download() as dl_info:
    page.get_by_role("link", name="Download PDF").click()
download = dl_info.value
chemin = f"data/pdf/{download.suggested_filename}"
download.save_as(chemin)
print("Sauvé :", chemin)
```

Pour ton cas (14 PDF de Command Reference Huawei via la visionneuse EDOC) :
ce pattern `expect_download` remplace les bidouilles `requests` quand le
site exige une session navigateur.

## 53. Intégration à ton pipeline collect_*

```text
collect_failed.py   (requests) ──▶ failed.txt ──┐
collect_pw.py       (requests) ──▶ failed_pc.txt ─┤
collect_still.py    (4 workers) ─▶ failed_pc_still.txt ─┤
collect_final.py    (APIs) ──────▶ failed_final.txt ─────┤
                                                ▼
                                    PLAYWRIGHT (collect_pw2.py)
                                    dernier recours avant l'étape cloud
```

Squelette `scripts/collect_pw2.py` (réutilise TON format de listes) :

```python
"""Passe Playwright : lit failed_final.txt, écrit ok + still_failed."""
import argparse, time, random
from pathlib import Path
from playwright.sync_api import sync_playwright

def scrape_one(page, url: str, outdir: Path) -> bool:
    try:
        page.goto(url, wait_until="domcontentloaded", timeout=45000)
        page.get_by_role("main").first.wait_for(timeout=15000)
        texte = page.locator("main").inner_text()
        if len(texte.strip()) < 200:
            return False  # page vide / challenge : on laisse tomber proprement
        (outdir / (hash_url(url) + ".md")).write_text(f"# {page.title()}\n\n{texte}")
        return True
    except Exception as e:
        print(f"ÉCHEC {url} : {type(e).__name__}")
        return False

def hash_url(url: str) -> str:
    import hashlib
    return hashlib.sha256(url.encode()).hexdigest()[:16]

def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument("--input", type=Path, required=True)   # failed_final.txt
    ap.add_argument("--output", type=Path, required=True)   # dossier .md
    ap.add_argument("--still", type=Path, required=True)   # still_failed.txt
    args = ap.parse_args()
    urls = [l.strip() for l in args.input.read_text().splitlines() if l.strip()]
    args.output.mkdir(parents=True, exist_ok=True)
    rates, still = [], []
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        ctx = browser.new_context(locale="fr-FR")
        page = ctx.new_page()
        for url in urls:
            ok = scrape_one(page, url, args.output)
            (rates if ok else still).append(url)
            time.sleep(random.uniform(1.5, 3.5))  # politesse
        browser.close()
    args.still.write_text("\n".join(still) + "\n")
    print(f"OK : {len(rates)}/{len(urls)} — restants : {args.still}")

if __name__ == "__main__":
    main()
```

> Ce script respecte TON contrat : 1 URL/ligne en entrée, `.md` propres
> en sortie, fichier des restants pour l'étape suivante.

## 54. Pièges Playwright (7)

1. **RAM** : ~200-600 Mo par instance (vérifié 2026). 10 scrapes parallèles
   = 2-6 Go. Surveille avec `htop`, limite à 2-4 workers.
2. **`networkidle` qui n'arrive jamais** : les SPA qui pollent en continu
   ne deviennent jamais « idle ». Timeout garanti → préfère
   `wait_until="domcontentloaded"` + `wait_for()` sur un sélecteur.
3. **Headless détecté** : certains WAF bloquent le mode headless. Teste
   `headless=False` avec `xvfb-run` sur serveur
   (`xvfb-run python scripts/collect_pw2.py`) avant de conclure.
4. **Oublier `browser.close()`** : en cas d'exception, le navigateur reste
   en RAM. Utilise `with sync_playwright()` + `try/finally`.
5. **Téléchargements sans `expect_download`** : le fichier part dans le
   vide. Toujours le pattern du §52.
6. **Sélecteurs trop fragiles** : `div.css-1a2b3c` change à chaque déploiement.
   Préfère `get_by_role` / `get_by_text`.
7. **Ne pas isoler les contextes** : réutiliser le même contexte pour 500
   pages = cookies/cache qui s'accumulent et fuites d'état. Un contexte
   neuf par lot (ou par site).

## 55. Pense-bête Playwright

```bash
pip install playwright && playwright install chromium
playwright codegen <url>                 # découvrir les sélecteurs
playwright show-trace debug/trace.zip    # rejouer un bug
```

```python
with sync_playwright() as p:
    b = p.chromium.launch(headless=True)
    pg = b.new_page()
    pg.goto(url, wait_until="domcontentloaded")
    pg.locator("main").inner_text()
    b.close()
```

---

## 56. Puppeteer : panorama

**Puppeteer** (équipe Google Chrome) = l'équivalent Node.js de Playwright,
mais **limité à Chromium** (support Firefox via BiDi, expérimental) et
**uniquement JavaScript/TypeScript**. Historiquement le premier (2017),
il reste très utilisé et bien documenté.

```bash
npm init -y
npm install puppeteer     # télécharge un Chromium compatible
```

```javascript
// scrape.js — exemple commenté
const puppeteer = require('puppeteer');

async function scrape(url) {
  // launch() : démarre Chromium. headless: true par défaut depuis les
  // versions récentes ("new" headless = vrai Chrome sans fenêtre).
  const browser = await puppeteer.launch({ headless: true });
  const page = await browser.newPage();

  // Bloquer les images = même optimisation que Playwright (§50)
  await page.setRequestInterception(true);
  page.on('request', (req) => {
    if (['image', 'font', 'stylesheet'].includes(req.resourceType())) req.abort();
    else req.continue();
  });

  // waitUntil: 'networkidle2' ≈ "plus de 2 connexions réseau pendant 500 ms"
  await page.goto(url, { waitUntil: 'networkidle2', timeout: 45000 });

  const data = await page.evaluate(() => ({
    title: document.title,
    // page.evaluate exécute du JS DANS la page : extraction sur mesure
    texte: document.querySelector('main')?.innerText.slice(0, 5000) ?? '',
  }));

  await browser.close();
  return data;
}

scrape('https://example.com').then(console.log).catch(console.error);
```

```bash
node scrape.js
```

## 57. Playwright vs Puppeteer : quand choisir quoi (vérifié 2026)

| Critère | Playwright | Puppeteer |
|---|---|---|
| Langages | JS/TS, **Python**, Java, .NET | JS/TS uniquement |
| Navigateurs | Chromium, Firefox, **WebKit** | Chromium (+ Firefox expérimental) |
| Auto-waiting | ✅ natif, strict | partiel (API Locator récente, opt-in) |
| Contextes isolés | ✅ first-class (proxy par contexte) | contextes incognito |
| Codegen / trace viewer | ✅ | ❌ |
| Mainteneur | Microsoft | Google (équipe Chrome) |
| Intégration Chrome DevTools | bonne | **excellente** (CDP natif) |
| Poids | ~330 Mo (npm + navigateurs) | ~280 Mo |

**Verdict 2026 (consensus des comparatifs) :**
- **Nouveau projet** → Playwright (ton cas : tu es en Python).
- **Service Node.js existant, 100 % Chrome** → Puppeteer reste légitime,
  plus léger, CDP au plus près.
- **Tu maintiens du Puppeteer** → ça marche toujours, pas de raison
  impérieuse de migrer.

## 58. Cohabitation dans ton pipeline

Tu n'as pas à choisir « pour toujours » : `collect_pw.py` peut rester en
Python/Playwright pendant qu'un micro-service Node/Puppeteer génère des PDF
côté API. Le seul contrat qui compte, c'est le **format d'échange**
(`failed*.txt` en entrée, `.md` + manifestes en sortie).

