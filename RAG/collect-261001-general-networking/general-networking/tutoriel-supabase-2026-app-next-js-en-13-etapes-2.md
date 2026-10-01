---
id: collect-261001-general-networking/general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes-2
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-supabase-2026-app-next-js-en-13-etapes.md
source_anchor: ""
source_lines: [84, 279]
sha256: c59dc716d5b662c1eb4404003c3f4bd543162fa71ee0dd073625a6330fefaedb
---

# macOS (Homebrew)

`supabase migration new create_tickets_schema`
Éditez le fichier généré dans `supabase/migrations/` avec le SQL suivant :

```
-- supabase/migrations/20260406120000_create_tickets_schema.sql
-- Table tickets
create table public.tickets (
  id uuid primary key default gen_random_uuid(),
  user_id uuid references auth.users(id) on delete cascade not null,
  title text not null check (length(title) between 3 and 200),
  description text,
  status text not null default 'open' check (status in ('open','pending','closed')),
  priority text not null default 'medium' check (priority in ('low','medium','high','urgent')),
  attachment_path text,
  created_at timestamptz default now() not null,
  updated_at timestamptz default now() not null
);
-- Table messages
create table public.messages (
  id uuid primary key default gen_random_uuid(),
  ticket_id uuid references public.tickets(id) on delete cascade not null,
  user_id uuid references auth.users(id) on delete cascade not null,
  body text not null,
  created_at timestamptz default now() not null
);
-- Indexes pour la performance
create index idx_tickets_user_id on public.tickets(user_id);
create index idx_tickets_status on public.tickets(status);
create index idx_messages_ticket_id on public.messages(ticket_id);
-- Activer Row Level Security
alter table public.tickets enable row level security;
alter table public.messages enable row level security;
-- Policies tickets : un utilisateur ne voit que ses propres tickets
create policy "Users can view own tickets" on public.tickets
  for select using (auth.uid() = user_id);
create policy "Users can insert own tickets" on public.tickets
  for insert with check (auth.uid() = user_id);
create policy "Users can update own tickets" on public.tickets
  for update using (auth.uid() = user_id);
create policy "Users can delete own tickets" on public.tickets
  for delete using (auth.uid() = user_id);
-- Policies messages : seulement messages de ses propres tickets
create policy "Users view messages of own tickets" on public.messages
  for select using (
    exists (
      select 1 from public.tickets
      where tickets.id = messages.ticket_id and tickets.user_id = auth.uid()
    )
  );
create policy "Users insert messages on own tickets" on public.messages
  for insert with check (
    auth.uid() = user_id and
    exists (
      select 1 from public.tickets
      where tickets.id = ticket_id and tickets.user_id = auth.uid()
    )
  );
-- Trigger updated_at automatique
create or replace function public.update_updated_at()
returns trigger as $$
begin
  new.updated_at = now();
  return new;
end;
$$ language plpgsql;
create trigger set_updated_at
  before update on public.tickets
  for each row execute function public.update_updated_at();
```
Appliquez la migration :

```
supabase db reset
# Sortie attendue
Resetting local database...
Recreating database...
Initialising schema...
Applying migration 20260406120000_create_tickets_schema.sql...
Seeding data supabase/seed.sql...
Finished supabase db reset on branch main.
```
La commande `supabase db reset` détruit et recrée toute la base locale, applique chaque migration dans l’ordre alphabétique, puis exécute le fichier `seed.sql` pour insérer des données de test. Cette commande est destructrice : ne l’utilisez jamais sur la production.

## Étape 4 : Créer le projet Next.js 16 avec App Router et TypeScript

Next.js 16 introduit le compilateur Turbopack stable et un App Router optimisé pour les Server Components React 19. Cette combinaison réduit drastiquement la quantité de JavaScript envoyée au navigateur – un avantage décisif pour une application Supabase qui peut faire la majorité de ses requêtes côté serveur.

```
cd tickets-support
# Créer l'app Next.js dans un sous-dossier
npx create-next-app@latest web \
  --typescript \
  --tailwind \
  --eslint \
  --app \
  --src-dir \
  --turbopack \
  --import-alias "@/*"
cd web
# Installer les SDK Supabase
npm install @supabase/[email protected] @supabase/[email protected]
# Vérifier
npm list @supabase/supabase-js @supabase/ssr
```
Le package `@supabase/ssr` remplace l’ancien `@supabase/auth-helpers-nextjs` (déprécié depuis 2024) et gère correctement les cookies de session entre Server Components, Route Handlers, Server Actions et Client Components – un piège classique qui causait des fuites de session avant cette refonte.

## Étape 5 : Configurer les variables d’environnement et le client Supabase

Créez le fichier `web/.env.local` avec les clés issues de `supabase status` (ou de votre projet cloud) :

```
# web/.env.local
NEXT_PUBLIC_SUPABASE_URL=http://127.0.0.1:54321
NEXT_PUBLIC_SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
SUPABASE_SERVICE_ROLE_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
Ajoutez immédiatement `.env.local` à `.gitignore` – Next.js le fait par défaut, mais vérifiez. Créez ensuite trois utilitaires de client Supabase, un par contexte d’exécution :

```
// src/lib/supabase/client.ts
import { createBrowserClient } from '@supabase/ssr'
export function createClient() {
  return createBrowserClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!
  )
}
```
```
// src/lib/supabase/server.ts
import { createServerClient } from '@supabase/ssr'
import { cookies } from 'next/headers'
export async function createClient() {
  const cookieStore = await cookies()
  return createServerClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!,
    {
      cookies: {
        getAll() {
          return cookieStore.getAll()
        },
        setAll(cookiesToSet) {
          try {
            cookiesToSet.forEach(({ name, value, options }) =>
              cookieStore.set(name, value, options)
            )
          } catch {
            // Ignoré dans Server Component, géré par middleware
          }
        },
      },
    }
  )
}
```
```
// src/lib/supabase/middleware.ts
import { createServerClient } from '@supabase/ssr'
import { NextResponse, type NextRequest } from 'next/server'
export async function updateSession(request: NextRequest) {
  let response = NextResponse.next({ request })
  const supabase = createServerClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!,
    {
      cookies: {
        getAll() {
          return request.cookies.getAll()
        },
        setAll(cookiesToSet) {
          cookiesToSet.forEach(({ name, value }) =>
            request.cookies.set(name, value)
          )
          response = NextResponse.next({ request })
          cookiesToSet.forEach(({ name, value, options }) =>
            response.cookies.set(name, value, options)
          )
        },
      },
    }
  )
  const { data: { user } } = await supabase.auth.getUser()
  // Redirection si non authentifié sur routes protégées
  if (!user && request.nextUrl.pathname.startsWith('/app')) {
    const url = request.nextUrl.clone()
    url.pathname = '/login'
    return NextResponse.redirect(url)
  }
  return response
}
```
Créez ensuite `src/middleware.ts` à la racine de `src/` qui appelle `updateSession()` sur chaque requête. Ce middleware rafraîchit automatiquement le JWT toutes les heures sans intervention manuelle – un point critique pour éviter les déconnexions intempestives.

## Étape 6 : Implémenter l’authentification avec Magic Link et OAuth

