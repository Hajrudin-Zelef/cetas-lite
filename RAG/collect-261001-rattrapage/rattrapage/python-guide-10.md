---
id: collect-261001-rattrapage/rattrapage/python-guide-10
title: "Guide Python complet — De l'installation aux scripts sysadmin"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/python_guide.md
source_anchor: ""
source_lines: [2804, 3070]
sha256: 6c32e58854ebccfd2b6bb31fc958b9104a3386fa2808542cb5e4f3ccc7b3d11b
---

# Guide Python complet — De l'installation aux scripts sysadmin

    actifs = [r for r in resultats if r["ssh"] == "ouvert"]
    with open(args.output, "w", newline="", encoding="utf-8") as f:
        w = csv.DictWriter(f, fieldnames=["ip", "hostname", "ssh", "http", "https"])
        w.writeheader()
        w.writerows(resultats)

    print(f"{len(actifs)} hôtes avec SSH ouvert → {args.output}")

if __name__ == "__main__":
    main()
```

---

## 55. Scripts sysadmin n°2 : parsing de logs

```python
#!/usr/bin/env python3
"""top_ips.py — top 10 des IPs en échec d'authentification SSH.

Usage : python3 top_ips.py /var/log/auth.log [-n 20] [--seuil 5]
"""
import argparse
import re
from collections import Counter

MOTIF = re.compile(r"Failed password .* from (?P<ip>\S+) port")

def analyser(chemin):
    compteur = Counter()
    with open(chemin, encoding="utf-8", errors="replace") as f:
        for ligne in f:
            m = MOTIF.search(ligne)
            if m:
                compteur[m.group("ip")] += 1
    return compteur

def main():
    p = argparse.ArgumentParser()
    p.add_argument("fichier")
    p.add_argument("-n", type=int, default=10, help="nombre d'IPs à afficher")
    p.add_argument("--seuil", type=int, default=5,
                   help="seuil d'alerte (tentatives)")
    args = p.parse_args()

    compteur = analyser(args.fichier)
    print(f"{'IP':<16} {'Tentatives':>10}  Alerte")
    print("-" * 40)
    for ip, nb in compteur.most_common(args.n):
        alerte = " !!!" if nb >= args.seuil else ""
        print(f"{ip:<16} {nb:>10}{alerte}")

    total = sum(compteur.values())
    print(f"\nTotal : {total} échecs, {len(compteur)} IP(s) distinctes")

if __name__ == "__main__":
    main()
```

---

## 56. Scripts sysadmin n°3 : sauvegarde et rotation

```python
#!/usr/bin/env python3
"""backup.py — archive un dossier en tar.gz daté + rotation (garde N dernières).

Usage : python3 backup.py /srv/data /srv/backups --keep 7
"""
import argparse
import logging
import shutil
import tarfile
from datetime import date
from pathlib import Path

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(message)s")
log = logging.getLogger(__name__)

def archiver(source: Path, dest_dir: Path) -> Path:
    nom = f"{source.name}-{date.today().isoformat()}.tar.gz"
    cible = dest_dir / nom
    log.info("Archivage de %s → %s", source, cible)
    with tarfile.open(cible, "w:gz") as tar:
        tar.add(source, arcname=source.name)
    taille_mo = cible.stat().st_size / 1024**2
    log.info("Archive créée : %.1f Mo", taille_mo)
    return cible

def rotation(dest_dir: Path, prefixe: str, garder: int):
    archives = sorted(dest_dir.glob(f"{prefixe}-*.tar.gz"))
    for vieille in archives[:-garder] if len(archives) > garder else []:
        log.info("Suppression ancienne archive : %s", vieille.name)
        vieille.unlink()

def main():
    p = argparse.ArgumentParser(description="Sauvegarde avec rotation")
    p.add_argument("source", type=Path)
    p.add_argument("destination", type=Path)
    p.add_argument("--keep", type=int, default=7)
    args = p.parse_args()

    if not args.source.is_dir():
        raise SystemExit(f"Source introuvable : {args.source}")
    args.destination.mkdir(parents=True, exist_ok=True)

    archiver(args.source, args.destination)
    rotation(args.destination, args.source.name, args.keep)

    libre_go = shutil.disk_usage(args.destination).free / 1024**3
    log.info("Espace libre sur la destination : %.1f Go", libre_go)

if __name__ == "__main__":
    main()
```

---

## 57. Scripts sysadmin n°4 : surveillance disque et alertes

```python
#!/usr/bin/env python3
"""disk_watch.py — alerte si un point de montage dépasse un seuil.

Usage : API_TOKEN=xxx python3 disk_watch.py --seuil 85 --webhook https://...
"""
import argparse
import logging
import os
import shutil
import sys
import requests

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(message)s")
log = logging.getLogger(__name__)

def usage_points(points):
    resultats = []
    for pt in points:
        u = shutil.disk_usage(pt)
        pct = u.used / u.total * 100
        resultats.append((pt, pct, u.free / 1024**3))
    return resultats

def alerter(webhook, message):
    try:
        requests.post(webhook, json={"text": message}, timeout=10).raise_for_status()
    except requests.RequestException as e:
        log.error("Échec envoi alerte : %s", e)

def main():
    p = argparse.ArgumentParser()
    p.add_argument("--seuil", type=float, default=85.0)
    p.add_argument("--webhook", default=os.environ.get("ALERT_WEBHOOK"))
    p.add_argument("points", nargs="*", default=["/"])
    args = p.parse_args()

    probleme = False
    for pt, pct, libre_go in usage_points(args.points):
        etat = "ALERTE" if pct >= args.seuil else "ok"
        log.info("%s : %.1f%% utilisé (%.1f Go libres) [%s]", pt, pct, libre_go, etat)
        if pct >= args.seuil:
            probleme = True
            if args.webhook:
                alerter(args.webhook,
                        f"⚠ Disque {pt} à {pct:.1f}% ({libre_go:.1f} Go libres)")

    sys.exit(1 if probleme else 0)   # code retour exploitable par cron/nagios

if __name__ == "__main__":
    main()
```

> Astuce : `sys.exit(1)` permet à cron, Nagios, Zabbix ou systemd de détecter l'alerte via le code de retour.

---

## 58. Scripts sysadmin n°5 : requêtes API automatisées

```python
#!/usr/bin/env python3
"""glpi_tickets.py — récupère les tickets ouverts via l'API (exemple générique).

Usage : API_URL=https://glpi.lan/apirest.php API_TOKEN=xxx python3 glpi_tickets.py
"""
import os
import sys
import requests

API_URL = os.environ.get("API_URL")
TOKEN = os.environ.get("API_TOKEN")

def session_api():
    if not API_URL or not TOKEN:
        raise SystemExit("API_URL et API_TOKEN requis en variables d'environnement")
    s = requests.Session()
    s.headers.update({"Authorization": f"Bearer {TOKEN}",
                      "Content-Type": "application/json"})
    return s

def tickets_ouverts(session):
    r = session.get(f"{API_URL}/tickets", params={"status": "open"}, timeout=15)
    r.raise_for_status()
    return r.json()

def main():
    s = session_api()
    try:
        tickets = tickets_ouverts(s)
    except requests.RequestException as e:
        raise SystemExit(f"Erreur API : {e}")
    print(f"{len(tickets)} ticket(s) ouvert(s)")
    for t in sorted(tickets, key=lambda x: x.get("priority", 0), reverse=True):
        print(f"  #{t['id']} [P{t.get('priority')}] {t.get('title')}")

if __name__ == "__main__":
    main()
```

---

## 59. Scripts sysadmin n°6 : traitement CSV/Excel

```python
#!/usr/bin/env python3
"""fusion_parc.py — fusionne deux exports CSV d'inventaire (déduplique par hostname).

Usage : python3 fusion_parc.py export1.csv export2.csv -o parc_fusionne.csv
"""
import argparse
import csv

CHAMPS = ["hostname", "ip", "type", "site", "source"]

def lire(chemin, source):
    with open(chemin, newline="", encoding="utf-8-sig") as f:  # utf-8-sig : BOM Excel
        for ligne in csv.DictReader(f):
            ligne["source"] = source
            yield ligne

def main():
    p = argparse.ArgumentParser()
    p.add_argument("fichiers", nargs=2)
    p.add_argument("-o", "--output", required=True)
    args = p.parse_args()

    parc = {}
    for chemin in args.fichiers:
        for ligne in lire(chemin, source=chemin):
            cle = ligne["hostname"].strip().lower()
            parc.setdefault(cle, ligne)   # le 1er fichier gagne en cas de doublon

    with open(args.output, "w", newline="", encoding="utf-8") as f:
        w = csv.DictWriter(f, fieldnames=CHAMPS, extrasaction="ignore")
        w.writeheader()
        w.writerows(parc.values())

    print(f"{len(parc)} équipements uniques → {args.output}")

if __name__ == "__main__":
    main()
```

> Pour lire/écrire du vrai Excel (`.xlsx`) : `pip install openpyxl` puis `openpyxl.load_workbook(...)`.

---

