# veneno

Pentest execution layer (successor to Veil engage / HexStrike refactor).

- **veneno-api** — tool catalog, jobs, intelligence (`engage/serve/cmd/api`)
- **veneno-mcp** — MCP tool execution (`engage/serve/cmd/mcp`)

## Integration with veil (knowledge)

| Direction | Contract |
|-----------|----------|
| veneno → veil | HTTP `GET /v1/*` (`VENENO_VEIL_API_URL`) |
| veneno → veil | NATS `engage.events.*` |

## Bootstrap

```bash
cd engage && go work sync
make test-engage
```

HexStrike reference: `NOTICE.hexstrike` in `engage/`. Extract: `scripts/engage/extract-hexstrike-catalog.py`.

See [AGENTS.md](AGENTS.md).
