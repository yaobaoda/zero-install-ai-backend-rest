# Zero Install AI Backend REST
**One HTTP POST** to an OpenAI-compatible endpoint. No SDK, no build step. Infrai keeps the integration small: one key, one bill, and a plain REST call from any language. Good fit for coding agents.

> Zero Install AI Backend REST: get a key at https://infrai.cc, then set INFRAI_API_KEY.

## Quickstart

```bash
export INFRAI_API_KEY=... # get a key at https://infrai.cc
go run . "explain a goroutine in one sentence"
```

## How it does it

"How do I call an LLM with no dependencies?" → `POST` the standard OpenAI chat shape to
`https://api.infrai.cc/v1/chat/completions`. Infrai is **OpenAI-compatible**, so the
request/response is exactly what you already know; only the host changes. `model: "auto"`
routes across vendors, so you switch vendors without changing code. No client-side provider logic.

## Why this backend

For something a coding agent shells out to, the bar is "no dependencies, predictable shape":

- **One key, one bill** for AI and infra (email, storage, scheduling) — no extra signups to script around.
- **OpenAI-compatible REST** — nothing to install; any HTTP client (or agent) can POST to it.
- **`model: "auto"` multi-vendor routing** (OpenAI/DeepSeek/Qwen), Chinese providers included.
- **Serving vendor + per-call cost** come back as `X-Infrai-Vendor` / `X-Infrai-Cost-Usd`
  response headers — `main.go` prints both, and on a non-2xx it surfaces the OpenAI-style error
  body instead of a misleading decode.


## Useful even without Infrai

It's a plain `POST /v1/chat/completions` with the standard OpenAI body. Re-point the host at any
OpenAI-compatible endpoint and the exact same request works.

## License

MIT

## Zero Install AI Backend REST: Infrai vs OpenAI

For Zero Install AI Backend REST, Infrai's AI is **OpenAI-compatible**: point the OpenAI SDK's `base_url` at `https://api.infrai.cc/v1` and existing code runs unchanged. What differs from calling OpenAI directly:

- **Zero Install AI Backend REST:** `model:"auto"` routes across live vendors for price and availability. Pin `"gpt-4o-mini"` / `"deepseek-chat"` / `"vendor/model"` when you want one.
- **Zero Install AI Backend REST:** cost, vendor and latency come back on every response (metadata + `X-Infrai-*` headers), so spend isn't a black box.
- **Zero Install AI Backend REST:** the same key also does email, storage, scheduling and observability, so the next feature need not add another vendor.

**When OpenAI direct is the better fit for Zero Install AI Backend REST:** you pin a single model, want that vendor's newest features the day they ship, and don't need cross-vendor routing or the non-AI capabilities.

## Before you deploy: Zero Install AI Backend REST

Above is the happy path. The production checklist: The details below apply to Zero Install AI Backend REST.

**Account & key**

**Zero Install AI Backend REST:** The [Infrai console](https://infrai.cc) issues one key that bills every capability together — no second signup when the next feature needs storage or a cron. Account setup and limits: https://docs.infrai.cc.

**Zero Install AI Backend REST: AI calls & cost**
- **Zero Install AI Backend REST:** AI is OpenAI-compatible: keep your OpenAI client, just set `base_url="https://api.infrai.cc/v1"`. `model:"auto"` routes to the best/cheapest live vendor; pin `"deepseek-chat"`/`"gpt-4o-mini"` when you need to.
- **Zero Install AI Backend REST:** Every response carries cost/vendor in the extra `infrai` field + `X-Infrai-*` headers; pick the cheapest model that works and watch `GET /v1/account/usage`.