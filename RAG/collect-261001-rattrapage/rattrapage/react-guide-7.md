---
id: collect-261001-rattrapage/rattrapage/react-guide-7
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [1553, 1829]
sha256: 9b338718b13d791b6b8fd21fef9291a47a53b33c0c83d07fcf6da96063f04cbf
---

# Guide React — Le manuel complet

## 29. Error Boundaries — attraper les plantages

En React, une erreur JS **non rattrapée** pendant le rendu démonte **toute** l'application (écran blanc). Les **Error Boundaries** sont des composants qui capturent les erreurs de leurs enfants et affichent une UI de secours.

> ⚠️ Il n'existe **pas** de hook `useErrorBoundary` en natif : les Error Boundaries s'écrivent encore en **classe** (la seule exception légitime aux composants fonctionnels).

```jsx
import { Component } from 'react';

class ErrorBoundary extends Component {
  constructor(props) {
    super(props);
    this.state = { aErreur: false, erreur: null };
  }

  static getDerivedStateFromError(erreur) {
    return { aErreur: true, erreur }; // affiche le fallback au prochain rendu
  }

  componentDidCatch(erreur, infos) {
    // Ici : log vers Sentry / console / API interne
    console.error('Erreur capturée :', erreur, infos.componentStack);
  }

  render() {
    if (this.state.aErreur) {
      return this.props.fallback ?? (
        <div className="erreur">
          <h2>⚠️ Un module a planté</h2>
          <p>{this.state.erreur?.message}</p>
          <button onClick={() => this.setState({ aErreur: false, erreur: null })}>
            Réessayer
          </button>
        </div>
      );
    }
    return this.props.children;
  }
}

// Utilisation : isoler les zones à risque
function Dashboard() {
  return (
    <main>
      <ErrorBoundary>
        <GraphiqueTemperatures />  {/* si ça plante, le reste survit */}
      </ErrorBoundary>
      <ErrorBoundary>
        <CarteOnduleurs />
      </ErrorBoundary>
    </main>
  );
}
```

**Ce que ça capture :** erreurs de rendu, des méthodes de cycle de vie, des constructeurs des enfants.

**Ce que ça NE capture PAS :** erreurs dans les gestionnaires d'événements (try/catch manuel), le code asynchrone (`setTimeout`, promesses), l'Error Boundary elle-même, le rendu SSR partiel.

**Stratégie :** placer des boundaries **par zone fonctionnelle** (un widget qui plante ne doit pas tuer le dashboard), pas une seule au sommet.

---

## 30. Modules — organiser son code

### Imports / exports ES

```jsx
// utils/format.js
export function formatKva(v) { return `${v} kVA`; }
export function formatPct(v) { return `${v.toFixed(1)} %`; }
export default function formatDate(d) { /* ... */ }

// Composant.jsx
import formatDate, { formatKva, formatPct } from './utils/format.js';
```

|  | Export nommé | Export default |
|---|---|---|
| Syntaxe | `export function …` / `export const …` | `export default …` |
| Import | `import { nom } from '…'` (nom exact) | `import NimporteQuelNom from '…'` |
| Par fichier | Autant qu'on veut | Un seul |
| Convention React | Utilitaires, hooks | **Composants** (un par fichier) |

### Arborescence conseillée pour un projet qui grandit

```
src/
├── main.jsx
├── App.jsx
├── pages/            # écrans liés au routeur (Dashboard, Tickets…)
├── components/       # composants réutilisables (Bouton, Carte, Tableau…)
│   └── ui/
├── hooks/            # usePolling, useLocalStorage…
├── contexts/         # AuthContext, ThemeContext…
├── services/         # appels API (api.js : fetch centralisé)
├── utils/            # formatage, constantes
└── styles/
```

**Règles :**
- Un composant = un fichier = nom du fichier.
- Les appels `fetch` **jamais** en dur dans les composants : les centraliser dans `services/api.js` (base URL, headers, gestion d'erreurs commune).
- Imports absolus via alias `@` → config Vite (section 33) pour éviter `../../../../`.

```js
// services/api.js — client HTTP centralisé
const BASE = import.meta.env.VITE_API_URL ?? 'http://localhost:3000';

async function request(path, options = {}) {
  const res = await fetch(`${BASE}${path}`, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  });
  if (!res.ok) throw new Error(`API ${res.status} sur ${path}`);
  return res.status === 204 ? null : res.json();
}

export const api = {
  listerOnduleurs: () => request('/api/onduleurs'),
  detailOnduleur: (id) => request(`/api/onduleurs/${id}`),
  creerTicket: (t) => request('/api/tickets', { method: 'POST', body: JSON.stringify(t) }),
};
```

---

## 31. React Router — routes de base

**React Router** est le routeur standard : il associe des URL à des composants (navigation sans rechargement de page = SPA).

```bash
npm install react-router-dom
```

```jsx
// main.jsx
import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import App from './App.jsx';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </StrictMode>
);
```

```jsx
// App.jsx
import { Routes, Route, Link } from 'react-router-dom';
import Dashboard from './pages/Dashboard.jsx';
import Tickets from './pages/Tickets.jsx';
import NonTrouve from './pages/NonTrouve.jsx';

export default function App() {
  return (
    <div>
      <nav>
        {/* Link = navigation côté client (pas de rechargement) */}
        <Link to="/">Dashboard</Link>
        <Link to="/tickets">Tickets</Link>
      </nav>

      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/tickets" element={<Tickets />} />
        {/* Route 404 : toujours en dernier */}
        <Route path="*" element={<NonTrouve />} />
      </Routes>
    </div>
  );
}
```

**`<Link>` vs `<a>` :** `<a href>` recharge toute la page (on perd l'état React). `<Link to>` navigue côté client. **Toujours `<Link>`** pour la navigation interne.

> **Note version :** React Router v6/v7 : API `createBrowserRouter` (data routers) recommandée pour les nouveaux projets. La syntaxe `<Routes>` ci-dessus reste valable et plus simple pour débuter.

---

## 32. React Router — params, navigation, layout

### Routes avec paramètres

```jsx
<Routes>
  <Route path="/onduleurs" element={<ListeOnduleurs />} />
  <Route path="/onduleurs/:id" element={<DetailOnduleur />} />
</Routes>;

// pages/DetailOnduleur.jsx
import { useParams } from 'react-router-dom';

function DetailOnduleur() {
  const { id } = useParams(); // { id: 'ups-01' } pour /onduleurs/ups-01
  // … charger les données avec id (section 28)
}
```

### Navigation programmatique

```jsx
import { useNavigate } from 'react-router-dom';

function FormulaireTicket() {
  const navigate = useNavigate();

  const envoyer = async (e) => {
    e.preventDefault();
    const ticket = await api.creerTicket(donnees);
    navigate(`/tickets/${ticket.id}`); // redirige vers le ticket créé
  };
}
```

### Layout partagé (routes imbriquées)

```jsx
import { Outlet } from 'react-router-dom';

function Layout() {
  return (
    <div className="app">
      <BarreLaterale />
      <main>
        <Outlet /> {/* la page enfant s'affiche ici */}
      </main>
    </div>
  );
}

<Routes>
  <Route element={<Layout />}>
    <Route path="/" element={<Dashboard />} />
    <Route path="/tickets" element={<Tickets />} />
  </Route>
  <Route path="/login" element={<Login />} /> {/* sans layout */}
</Routes>;
```

**Autres hooks utiles :** `useLocation()` (URL actuelle), `useSearchParams()` (query string `?tri=nom`).

```jsx
const [params, setParams] = useSearchParams();
const tri = params.get('tri') ?? 'nom';
// changer : setParams({ tri: 'puissance' })
```

---

## 33. Vite — configurer et builder

### Fichier `vite.config.js`

```js
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: { '@': '/src' }, // imports absolus : import X from '@/components/X'
  },
  server: {
    port: 5173,
    proxy: {
      // Redirige /api vers le backend en dev → évite les problèmes CORS
      '/api': { target: 'http://localhost:3000', changeOrigin: true },
    },
  },
});
```

