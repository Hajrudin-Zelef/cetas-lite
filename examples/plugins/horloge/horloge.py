#!/usr/bin/env python3
"""Exemple de plugin Cetas : protocole d'echange.

Cetas ecrit les arguments de l'appel (JSON) sur stdin.
Le script repond sur stdout, au choix :
  - du JSON {"result": ...}  (recommande : "result" peut etre texte, nombre,
    objet... ; {"error": "..."} signale un echec) ;
  - du texte brut (utilise tel quel comme resultat).

Le 1er argument de la ligne de commande designe l'outil a executer
(ici on reutilise un seul script pour les deux outils du plugin).
"""
import json
import sys
from datetime import datetime


def main() -> int:
    outil = sys.argv[1] if len(sys.argv) > 1 else ""
    try:
        args = json.load(sys.stdin)
    except Exception:
        args = {}
    if not isinstance(args, dict):
        args = {}

    if outil == "maintenant":
        fmt = args.get("format", "court")
        now = datetime.now()
        if fmt == "complet":
            texte = now.strftime("%A %d %B %Y, %H:%M:%S")
        else:
            texte = now.strftime("%Y-%m-%d %H:%M")
        print(json.dumps({"result": texte}, ensure_ascii=False))
    elif outil == "compte_mots":
        texte = str(args.get("texte", ""))
        mots = len(texte.split())
        print(json.dumps({"result": f"{mots} mot(s), {len(texte)} caractere(s)"}))
    else:
        print(json.dumps({"error": f"outil inconnu: {outil}"}))
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
