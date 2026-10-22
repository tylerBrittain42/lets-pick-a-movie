# Claude Feedback: Frontend ↔ Backend Architecture

Notes from a design discussion (2026-10-21).

## Decision so far

Leaning toward a **traditional Multi-Page Application (MPA)**: server-rendered
Go templates and plain HTML forms, "old school" style. Business logic lives in
internal service packages.

---

## 1. Ways a site can talk to a backend

| Approach | One-liner | Pros | Cons | Search terms |
|---|---|---|---|---|
| REST (JSON/HTTP) | Resources as URLs, verbs as actions | Universal, easy to debug, `net/http` is enough | No contract, over/under-fetching | "REST resource design", "Richardson maturity model", "Go 1.22 ServeMux patterns" |
| REST + OpenAPI | Spec file, then generated server/client | Single source of truth, typed clients | Tooling weight, clunky codegen | "OpenAPI 3", "oapi-codegen" |
| GraphQL | One endpoint, client asks for exact shape | No over-fetching, typed schema | Overkill for small apps, caching, N+1 | "gqlgen", "dataloader", "when not to use GraphQL" |
| gRPC / Connect | Protobuf contracts plus codegen | Strict, fast, very Go-friendly | Browsers need gRPC-Web/Connect | "connectrpc Go", "buf CLI" |
| Server-rendered HTML | Backend sends finished pages | Simplest, one codebase | Less interactive by default | "Go html/template", "templ" |
| HTMX / hypermedia | Server returns HTML fragments | Interactivity without a JS framework | Hard for complex client state | "HTMX", "Hypermedia Systems", "Datastar" |
| WebSockets | Persistent two-way connection | Real-time both directions | Connection state, scaling | "coder/websocket" |
| Server-Sent Events | One-way server→browser stream | Simple, auto-reconnect | Server-to-client only | "SSE Go", "EventSource" |
| tRPC | Typed calls end to end | — | **TypeScript only, so not an option with a Go backend** | — |

**Takeaway:** my workload (CRUD plus a CSV upload) doesn't need GraphQL, gRPC, or
WebSockets. Plain HTTP covers it.

---

## 2. CSV upload

- **`multipart/form-data`**: how browsers send files. See `Request.FormFile`
  and `ParseMultipartForm`.
- **`http.MaxBytesReader`**: cap the upload size.
- **Streaming vs. buffering**: does the parser need the whole file in memory?
- **Parser design question:** what does `letterboxdParser` accept as input?
  A file path ties it to disk. What could it accept to work equally with a file,
  an HTTP upload, and a string in a test? Research `io.Reader` and "accept
  interfaces, return structs".
- **Open questions:**
  - Malformed row 300 of 500: reject everything, or skip it and report?
  - Save the parsed rows inside one DB transaction?

---

## 3. One site with an `/api/` route?

Not a bad idea. It's a very common pattern.

- **Pros:** same origin (mostly no CORS), cookie auth is easy, one deploy, can split
  later.
- **Cons:** frontend and backend deploy and scale together (doesn't matter at this
  scale).
- **Research:** `embed`, `http.FileServer` / `http.FS`, SPA fallback routing,
  `/api/v1/` prefix.

---

## 4. Critique of my layering plan (internal packages + controllers + api)

1. **Two rendering paths:** if page load renders HTML in Go and interactions
   return JSON, the same markup ends up written in both a Go template and JS, and
   the two drift apart. Decide who owns the HTML: the server (HTMX), the client
   (SPA), or both deliberately (SSR + hydration, which is painful by hand).
2. **Dependency direction:** services should not import `net/http` or know
   about request/response types. Search "Go dependency direction".
3. **Thin handlers:** controller and API handlers should stay a few lines. If
   they grow, logic is leaking out of the services.
4. **Small points:**
   - "Controller" is an MVC term. Go usually says *handlers* (a `web`/`ui` package).
   - `pkg/` vs. `internal/`: will anything outside this repo import the parser?
     If not, it may not belong in `pkg/`.
   - Package names: lowercase, no camelCase (`letterboxdparser`).

**Exercise:** trace one interaction ("mark movie as watched") through the whole
design: what the browser sends, which package handles it, what comes back, and how
the page updates.

---

## 5. SPA vs. initial page load

- An SPA is valid, but much of the effort shifts to JS/TS instead of Go.
- Initial-load concerns usually mean one of:
  1. SEO / link previews (likely irrelevant here)
  2. A loading flash (UX, solved with a loading state)
  3. An extra round trip (milliseconds)
- **Middle ground:** the Go template injects the initial data as JSON into the shell,
  and the client renders it. That keeps one rendering path. Search "bootstrapping
  initial state SPA", "data island", "html/template contexts JSON".
- Start simple and add this only if the flash actually bothers me.

---

## 6. Traditional MPA (current direction)

**What it looks like:** each page is a Go handler that loads data from the
services and renders an `html/template`. Changes go through `<form>` POSTs. The
CSV upload is a plain form with a file input and needs no JS.

**Concepts to research:**
- **Post/Redirect/Get (PRG)**: why redirect after a POST (hint: hit refresh).
- **Template layouts**: `{{define}}`, `{{block}}`, `ParseFS` + `embed`.
- **Forms**: `r.ParseForm`, `r.FormValue`, `enctype="multipart/form-data"`.
- **Validation errors**: re-show the user's input along with the errors.
- **Flash messages**: carry "Imported 312 movies!" across a redirect.
- **CSRF**: search "CSRF Go"; see `http.CrossOriginProtection`.

**Impact on the plan:**
- `/api/` becomes optional unless something other than the browser needs it.
- Internal service packages stay as planned.
- HTMX can be added later without a rewrite ("progressive enhancement").

**To do:** review `plan_api.md` and decide how much was really about JSON/HTTP
and how much was service logic in disguise. Move the latter to
`plan_service.md`.

---

## 7. What to call it when searching

- **Multi-Page Application (MPA)**: the main term.
- **Server-side rendering (SSR)**: add "Go" to filter out Next.js results.
- **Hypermedia-driven application**: HTMX community term.
- **Progressive enhancement**: plain HTML first, JS only where it helps.
- **"Choose Boring Technology"**: essay making the case for this style.
- **Go-specific:** "Go html/template web app", *Let's Go* by Alex Edwards,
  "Go server rendered CRUD app".

---

## 8. Streaming Follow-Up

### What to look up

**Foundations**
- **`io.Reader` / `io.Writer`**: nearly everything that streams goes through these.
- **`bufio`**: buffered readers and writers, and why buffering matters.
- **`io.Pipe`**: connects a writer to a reader in memory (producer/consumer).
- **`io.Copy`**: moves data between them without loading all of it.

**JSON**
- **`json.Decoder` vs. `json.Unmarshal`**: stream vs. all bytes up front.
- **`Decoder.Token` / `Decoder.More`**: walk a huge array one element at a time.
- **`json.Encoder`**: write straight to an `io.Writer` (e.g. an HTTP response).
- **NDJSON / JSON Lines**: one object per line. Search "NDJSON vs JSON array streaming".
- **`encoding/json/v2` / `jsontext`**: newer package with better streaming.
  Started as an experiment in Go 1.25, so check its status in 1.27.

**HTTP**
- **Chunked transfer encoding**: sending a response before knowing its size.
- **`http.Flusher` / `http.ResponseController`**: push partial output now.
- **Server-Sent Events (SSE)**: simple browser streaming format.
- **`context` cancellation**: stop cleanly when the client disconnects
  (`r.Context().Done()`).

**Concepts**
- **Backpressure**: producer faster than consumer.
- **`iter.Seq`**: range-over-func iterators for "one item at a time".

**Tie-in:** `csv.Reader.Read` vs. `ReadAll` is the same streaming vs. buffering
choice. Could the upload stream from request → parser → DB without holding the
whole file in memory?

### Exercises to try now

1. **Measure memory:** parse a 1GB CSV with `ReadAll` vs. `Read`. Compare with
   `runtime.MemStats` or `/usr/bin/time -v`.
2. **Write your own `io.Reader`:** e.g. one that uppercases text or counts bytes.
   Search "Go wrapping io.Reader".
3. **Stream an HTTP response slowly:** one line per second, flush each, then watch
   with `curl -N`. Then remove the flush and compare.
4. **Handle a disconnect:** Ctrl+C the `curl` mid-stream; make the server notice
   and stop.
5. **Pipe goroutines together:** read → transform → write, with each stage in its
   own goroutine. Search "Go pipelines and cancellation".
6. **Read source:** `io.Copy`, `bufio.Scanner`, and `csv.Reader` in the stdlib.

**Start here:** exercise 1 with the Letterboxd parser and a fake 1GB CSV.

### Follow-up project ideas (roughly easy → hard)

- **Log tailer:** `tail -f` clone with filtering. Readers, polling, signals.
- **Live dashboard with SSE:** stream stats/events to a browser. Flushing,
  connection lifetime, fan-out. *(Suggested next project.)*
- **Large-file converter:** CSV ↔ NDJSON for files larger than RAM. Constant memory.
- **Chat server:** TCP (`net`) first, then WebSockets. Concurrency, broadcasting.
- **Reverse proxy:** stream bodies through. `net/http/httputil`, headers.
- **Mini `grep` / `jq`:** stdin → stdout CLI. Unix-style composition.
- **Job queue / worker pool:** backpressure, bounded channels, graceful shutdown.

### Resources

- *Let's Go* / *Let's Go Further* by Alex Edwards
- Go blog: "Go Concurrency Patterns: Pipelines" and "Context"
- *Concurrency in Go* by Katherine Cox-Buday
