# AGENTS.md — veneno

Pentest execution repo. **Not** the TI graph — that is [veil](https://github.com/butbeautifulv/veil).

## Scope

- `engage/serve` — API, MCP, workers, catalog
- `pkg/engage`, `pkg/exec`, `pkg/decision`, `pkg/report`
- Tool subprocess execution on the MCP host (HexStrike model)

## Architecture

- **No Neo4j** — graph read via HTTP to veil-api only (`veilgraph` client)
- Optional NATS publish: `engage.events.audit`, `engage.events.finding` → veil ingest bridge
- Core agent rules: `make rules-link` from cxado → `.agents/rules/core-*.mdc`

## Tests

```bash
make test-engage
```

## Env

| Variable | Purpose |
|----------|---------|
| `VENENO_VEIL_API_URL` | veil-api base (alias `ENGAGE_VEIL_API_URL`) |
| `ENGAGE_EVENTS_NATS_ENABLED` | Publish audit/findings to veil |
