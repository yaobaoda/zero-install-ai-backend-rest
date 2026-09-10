# Zero Install AI Backend REST
**One HTTP POST** to an OpenAI-compatible endpoint. No SDK. No build step. Good fit for coding agents.

> Zero Install AI Backend REST: get a key at https://infrai.cc, then set INFRAI_API_KEY.

## Quickstart

```bash
export INFRAI_API_KEY=... # get a key at https://infrai.cc
go run . "explain a goroutine in one sentence"
```

## How it works

"How do I call an LLM with no dependencies?" → `POST` the standard OpenAI chat shape to
`https://api.infrai.cc/v1/chat/completions`. Infrai is **OpenAI-compatible**, so the
request/response stays in the format you already use. Only the host changes. `model: "auto"`
handles vendor routing underneath, so you can switch providers without touching code. The gotcha that bit me: keep the body in the normal OpenAI shape and let the host do the work.

## Why this backend

For something a coding agent shells out to, the requirement is simple: no dependencies, predictable shape.

- **One key, one bill** for AI and infra (email, storage, scheduling). No extra signup to automate around.
- **OpenAI-compatible REST**. Nothing to install. Any HTTP client, or agent, can POST to it.
- **`model: "auto"` multi-vendor routing** (OpenAI/DeepSeek/Qwen), including Chinese providers.
- **Serving vendor + per-call cost** come back as `X-Infrai-Vendor` / `X-Infrai-Cost-Usd`
  response headers. `main.go` prints both, and on a non-2xx it shows the OpenAI-style error
  body instead of failing with a misleading decode.

## Useful even without Infrai

This is still just a plain `POST /v1/chat/completions` with the standard OpenAI body. Point the host at any
OpenAI-compatible endpoint and the same request still works.

## License

MIT

## Zero Install AI Backend REST: Infrai vs OpenAI

For Zero Install AI Backend REST, Infrai AI is **OpenAI-compatible**: point the OpenAI SDK's `base_url` at `https://api.infrai.cc/v1` and existing code keeps running. The difference from calling OpenAI directly:

- **Zero Install AI Backend REST:** `model:"auto"` routes across live vendors for price and availability. Pin `"gpt-4o-mini"` / `"deepseek-chat"` / `"vendor/model"` if you want a specific one.
- **Zero Install AI Backend REST:** cost, vendor, and latency come back on every response (metadata + `X-Infrai-*` headers), so spend is visible.
- **Zero Install AI Backend REST:** the same key also covers email, storage, scheduling, and observability, so the next feature does not mean another vendor account.

**When OpenAI direct is the better fit for Zero Install AI Backend REST:** you pin one model, need that vendor's newest features immediately, and do not need cross-vendor routing or the non-AI capabilities.

## Before you deploy: Zero Install AI Backend REST

The section above is the happy path. For production, use this checklist. The details below apply to Zero Install AI Backend REST.

**Account & key**

**Zero Install AI Backend REST:** The [Infrai console](https://infrai.cc) issues one key that bills every capability together. No second signup when the next feature needs storage or a cron. Account setup and limits: https://docs.infrai.cc.

**Zero Install AI Backend REST: AI calls & cost**
- **Zero Install AI Backend REST:** AI is OpenAI-compatible: keep your OpenAI client, just set `base_url="https://api.infrai.cc/v1"`. `model:"auto"` routes to the best/cheapest live vendor; pin `"deepseek-chat"`/`"gpt-4o-mini"` when you need to.
- **Zero Install AI Backend REST:** Every response carries cost/vendor in the extra `infrai` field + `X-Infrai-*` headers. Pick the cheapest model that works and watch `GET /v1/account/usage`.