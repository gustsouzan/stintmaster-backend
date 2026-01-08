# Stint Master Backend — Roadmap

Data base deste roadmap: 2026-01-08
Branch atual: `refacto/structure`

## 1) Visão do produto (resumo)
O backend deve permitir organizar eventos de endurance (pista, duração, classes), cadastrar pilotos com carros e restrições de horário, montar um grid/roster e gerar (e opcionalmente persistir) um plano de stints consistente e explicável.

## 2) Estado atual (o que já existe)
### API
- `GET /api/health`
- `GET /api/v1/cars/` + auxiliares (`/classes`, `/class-map`, `/car-suggestions`)
- `GET /api/v1/tracks/`
- `POST /api/v1/pilots/`, `GET /api/v1/pilots/`, `DELETE /api/v1/pilots/:id`
- `POST /api/v1/events/`
- `GET /api/v1/events/calculate-Event` (placeholder; domínio retorna `nil`)

### Persistência
- Postgres + GORM com `AutoMigrate` em runtime.
- Seed antigo em `init.sql` e uma migration Goose (`integrations/postgres/migrations/...`) que não bate com os models.

## 3) Objetivos (3-6 semanas)
### P0 (MVP funcional para o front)
- Reprodutibilidade do banco (schema definido e consistente).
- CRUD mínimo de Pilotos e Eventos + roster (inscrições no evento).
- Endpoint de “calcular stints” retornando um plano simples e válido.

### P1 (robustez e operação)
- Persistir corrida/stints (se necessário) e permitir re-calcular.
- Testes do domínio/algoritmo + CI.
- Padronização de erros/contratos de API + documentação.

## 4) Princípios de engenharia
- 1 “fonte de verdade” para o schema do banco.
- Separar regras de negócio (domínio) de IO (API e repos).
- Contrato de API estável (sem breaking changes sem versionamento).
- Entregas pequenas, com critérios de aceite claros.

## 5) Milestones por fase

### Milestone 0 — Alinhamento e base (Sprint 0 | 3–5 dias)
**Entrega:** base pronta para desenvolver sem retrabalho.
- Definir MVP e escopo (o que entra e o que não entra).
- Definir entidades e campos (Pilot, Car, Track, Event, Race/Corrida, Stint, Availability).
- Decidir estratégia de migrations: **Goose como fonte de verdade** (recomendação) e política para `AutoMigrate` (ex.: apenas dev ou removido).
- Padronizar contrato de API (erros, validação, naming JSON).

**Critérios de aceite**
- Existe um documento curto com o modelo de domínio e o contrato de API mínimo.
- Subir ambiente local resulta no mesmo schema sempre.


### Milestone 1 — Master Data + CRUD mínimo (Sprint 1 | 1 semana)
**Entrega:** front consegue operar dados base.
- Cars/Tracks: confirmar se serão “seed-only” (sem CRUD) ou editáveis.
- Pilots: completar CRUD mínimo (Create/List/Get/Update/Delete) e garantir persistência correta de carros e restrições.
- Events: além de Create, implementar ao menos List/Get (Update/Delete opcional no MVP).

**Critérios de aceite**
- Todos endpoints retornam erros em formato consistente.
- Validação de request body consistente (incluindo campos obrigatórios).


### Milestone 2 — Roster (inscrição de pilotos no evento) (Sprint 2 | 1 semana)
**Entrega:** evento tem uma lista de pilotos inscritos.
- Endpoints de roster: add/remove/list pilotos no evento.
- Regras: impedir duplicidade, validar existência, respeitar `min/max drivers`.
- (Opcional MVP) Checar compatibilidade de carro/classe.

**Critérios de aceite**
- Dá para montar um evento com N pilotos inscritos e consultar o estado atual.


### Milestone 3 — Motor de cálculo de stints (Sprint 3–4 | 2 semanas)
**Entrega:** `calculate` retorna plano de stints válido.
- Definir entradas necessárias (lap time, consumo, stint max, pit time, janelas de disponibilidade).
- Implementar algoritmo MVP determinístico:
  - respeita restrições de horário
  - distribui tempo de pilotagem
  - não excede stint max
  - retorna métricas (tempo por piloto, nº de trocas)
- (Opcional) Persistir plano em `Corrida/Stint`.

**Critérios de aceite**
- Para um evento de teste com 3–6 pilotos e restrições simples, o resultado é sempre válido e reprodutível.


### Milestone 4 — Operação e qualidade (Sprint 5 | 1 semana)
**Entrega:** confiável para uso contínuo.
- Testes focados do domínio/algoritmo.
- CI rodando `go test`.
- Logs estruturados e health melhor (idealmente com checagem de DB).
- Documentação de API (OpenAPI ou markdown detalhado).

## 6) Backlog priorizado (Epics → histórias)

### Epic A — Banco e consistência de schema (P0)
- A1: Definir schema final (tabelas/colunas/relacionamentos).
- A2: Unificar migrations (Goose) e remover duplicidade (init.sql / AutoMigrate).
- A3: Seed controlado para dev (cars/tracks de exemplo).

### Epic B — Pilots (P0)
- B1: CRUD completo de pilotos.
- B2: Persistir e consultar carros disponíveis do piloto.
- B3: Persistir e consultar restrições de horário.

### Epic C — Events (P0)
- C1: List/Get de eventos.
- C2: Validar min/max pilotos e classes.

### Epic D — Roster (P0)
- D1: Adicionar piloto ao evento.
- D2: Remover piloto do evento.
- D3: Listar pilotos inscritos no evento.

### Epic E — Cálculo de stints (P0/P1)
- E1: Definir request/response do calculate.
- E2: Implementar algoritmo MVP.
- E3: Persistência da corrida/plano (se necessário).

### Epic F — Qualidade e DX (P1)
- F1: Padronização de erros e validação.
- F2: Testes de domínio.
- F3: CI.
- F4: Docs.

## 7) Métricas de sucesso (para guiar decisões)
- Tempo para configurar ambiente local (target: < 10 min).
- % de requests validados/normalizados com erro consistente.
- Reprodutibilidade do cálculo (mesma entrada → mesma saída).
- Cobertura de testes do domínio do cálculo (não precisa ser alta, mas tem que cobrir invariantes).

## 8) Riscos e dívidas técnicas atuais (para tratar cedo)
- Inconsistência entre schema (init.sql / Goose / GORM models) → alto risco de retrabalho.
- Endpoint `calculate-Event` e domínio de corrida/stint incompletos.
- README e documentação de API quase inexistentes.

## 9) Decisões em aberto (responder para fechar o escopo)
1) MVP precisa **persistir** corrida/stints no banco ou é “calcular e retornar”?
2) É multiusuário (precisa auth/tenant) ou single-admin?
3) Regras de carro/classe são obrigatórias na lógica do roster/cálculo?
