# Archer

Archer is a Go-based Agent platform with a React administration console. It supports conversational agents, tool calling, knowledge bases, workflows, and integrations.

## Features

- Streaming chat and configurable agents
- Knowledge-base ingestion and retrieval
- Workflow execution with human approval
- MCP and A2A integrations
- Model configuration, memory, tracing, and metrics

## Stack

- Backend: Go, Hertz, PostgreSQL, Redis, Milvus, MinIO
- Agent: Eino and OpenAI-compatible model APIs
- Frontend: React, TypeScript, Tailwind CSS, Vite
- Observability: OpenTelemetry, Prometheus, Grafana

## Run locally

Requires Go 1.26+, Node.js 18+, and Docker Compose.

```bash
cd backend
make compose-core
make migrate
make run
```

In another terminal:

```bash
cd frontend
npm install
npm run dev
```

Open `http://localhost:5173`. The default development account is `admin / admin123`; change it before exposing the service. Configure a model provider in the administration console to enable conversations. The worker and optional services can be started with `make run-worker` and `make compose-up` from `backend`.

Configuration lives in `backend/configs/config.yaml`. Environment variables with the `ARCHER_` prefix can override settings. Development credentials in the example configuration are not suitable for production.

## Attribution

This project is adapted from an open-source Agent platform by [chengpeng-cp](https://github.com/chengpeng-cp). Archer is a derivative project, not the original author's course or an official release. Review the upstream project's license and attribution requirements before redistribution.
