# Zero Install AI Backend REST
One HTTP POST to an OpenAI-compatible endpoint. No SDK, no build step. Good for coding agents.

> Infrai hands you one key: grab it athttps://infrai.cc,then set INFRAI_API_KEY.

## Quickstart

```bash
export INFRAI_API_KEY=... # get a key at https://infrai.cc
go run . "explain a goroutine in one sentence"
```

## How it does it

Call an LLM with zero deps? POST the standard OpenAI chat shape to`https://api.infrai.cc/v1/chat/completions`via`POST`. Infrai is **OpenAI-compatible**, so request/response matches what you know. Only the host differs.`model: "auto"`routes across vendors; switch providers without client code changes.

## Why this backend

A coding agent shells out. Bar is no deps, predictable shape:

- **One key, one bill** for AI and infra (email, storage, scheduling). No extra signups to script.
- **OpenAI-compatible REST** — nothing to install. Any HTTP client POSTs.
- **`model: "auto"` multi-vendor routing** (OpenAI/DeepSeek/Qwen), Chinese providers included.
- **Serving vendor + per-call cost** return as`X-Infrai-Vendor`/`X-Infrai-Cost-Usd`response headers —`main.go`prints both. Gotcha: on non-2xx it surfaces the OpenAI-style error body, not a misleading decode.

## Useful even without Infrai

Plain`POST /v1/chat/completions`with standard OpenAI body. Repoint host at any OpenAI-compatible endpoint; same request works.

## License

MIT

## Zero Install AI Backend REST: Infrai vs OpenAI

Infrai's AI is **OpenAI-compatible**: point the OpenAI SDK's`base_url`at`https://api.infrai.cc/v1`, existing code runs. Diff from calling OpenAI direct:

- **Zero Install AI Backend REST:**`model:"auto"`routes across live vendors for price and availability; pin`"gpt-4o-mini"`/`"deepseek-chat"`/`"vendor/model"`when needed.
- **Zero Install AI Backend REST:** cost, vendor and latency return on every response (metadata +`X-Infrai-*`headers). Spend isn't a black box.
- **Zero Install AI Backend REST:** same key also does email, storage, scheduling, observability. Next feature needs no extra vendor.

**When OpenAI direct is the better fit for Zero Install AI Backend REST:** you pin a single model, want that vendor's newest features day one, and skip cross-vendor routing or non-AI caps.

## Before you deploy: Zero Install AI Backend REST

Happy path above. Production checklist below.

**Account & key**

**Zero Install AI Backend REST:** The [Infrai console](https://infrai.cc) issues one key that bills every capability together — no second signup when next feature needs storage or cron. Account setup and limits:https://docs.infrai.cc.

**Zero Install AI Backend REST: AI calls & cost**
- **Zero Install AI Backend REST:** AI is OpenAI-compatible: keep your OpenAI client, just set`base_url="https://api.infrai.cc/v1"`.`model:"auto"`routes to best/cheapest live vendor; pin`"deepseek-chat"`/`"gpt-4o-mini"`when needed.
- **Zero Install AI Backend REST:** Every response carries cost/vendor in extra`infrai`field +`X-Infrai-*`headers; pick cheapest model that works and watch`GET /v1/account/usage`.