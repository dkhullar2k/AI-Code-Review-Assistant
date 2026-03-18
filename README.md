# AI Code Review Assistant

An AI-powered system that automatically reviews GitHub Pull Requests using LLMs.

## Features

- GitHub webhook integration
- Automated PR review using AI
- Distributed architecture with Redis queue
- Worker-based async processing
- PostgreSQL storage
- React dashboard for visualization

## Tech Stack

- Backend: Go (Gin)
- Frontend: React + TypeScript + TailwindCSS
- Queue: Redis
- Database: PostgreSQL
- AI: Groq (LLM)

## Architecture

GitHub PR → Webhook Service (Go) → Redis Queue → Worker Service (Go) → AI Review Engine (LLM) → PostgreSQL Storage → API Server (Go) → React Dashboard