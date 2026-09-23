---
id: etape8-phasec-backend-frameworks/00-backend-frameworks/22-minimal-code-sketches-idiomatic-2026-style
title: "22. Minimal code sketches (idiomatic 2026 style)"
domain: step-8-phase-c-backend-frameworks-apis
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape8_phaseC_backend_frameworks.md
source_anchor: ""
source_lines: [551, 637]
section: "Step 8 — Phase C: Backend Frameworks & APIs"
sha256: af5357d3a7a2bd69a4432e25edbb1f6caf982a90b8bf9e5122fc34bb6d94275c
---

# 22. Minimal code sketches (idiomatic 2026 style)

## 22. Minimal code sketches (idiomatic 2026 style)

FastAPI — typed handler with automatic OpenAPI:

```python
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class Item(BaseModel):
    name: str
    price: float

@app.post("/items", response_model=Item, status_code=201)
async def create_item(item: Item) -> Item:
    return item  # validation + OpenAPI schema generated from type hints
```

Hono — edge-runnable TypeScript API with Zod validation:

```ts
import { Hono } from 'hono'
import { zValidator } from '@hono/zod-validator'
import { z } from 'zod'

const app = new Hono()
const schema = z.object({ name: z.string(), price: z.number() })

app.post('/items', zValidator('json', schema), (c) => {
  const item = c.req.valid('json') // fully typed
  return c.json(item, 201)
})
export default app // deploys to Workers, Bun, Deno, Node unchanged
```

Fiber v3 — Go handler with native context.Context:

```go
package main

import "github.com/gofiber/fiber/v3"

func main() {
    app := fiber.New()
    app.Post("/items", func(c fiber.Ctx) error {
        var item Item
        if err := c.Bind().JSON(&item); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(201).JSON(item) // c IS a context.Context in v3
    })
    app.Listen(":8080")
}
```

ASP.NET Core Minimal API — validation + OpenAPI 3.1 in a few lines:

```csharp
var builder = WebApplication.CreateBuilder(args);
builder.Services.AddOpenApi(); // OpenAPI 3.1 document generation
var app = builder.Build();
app.MapOpenApi();

app.MapPost("/items", (Item item) => TypedResults.Created($"/items/{item.Name}", item))
   .WithName("CreateItem")
   .WithOpenApi(); // typed results, validation, OpenAPI annotations

app.Run();
record Item(string Name, decimal Price);
```

NestJS — modular controller with DI and validation pipe:

```ts
import { Body, Controller, Post, UsePipes, ValidationPipe } from '@nestjs/common';
import { IsNumber, IsString } from 'class-validator';

class CreateItemDto {
  @IsString() name: string;
  @IsNumber() price: number;
}

@Controller('items')
export class ItemsController {
  constructor(private readonly itemsService: ItemsService) {}

