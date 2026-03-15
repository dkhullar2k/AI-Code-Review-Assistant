# AI Code Review Assistant

An automated AI-powered pull request review system built with Go.

## Architecture

GitHub Webhooks → Redis Queue → Review Worker → AI Engine → GitHub PR Comments

## Tech Stack

Backend: Go (Gin)
Queue: Redis
Database: PostgreSQL
AI: Groq LLM
Integration: GitHub API

## Features

• Automated pull request code reviews
• AI detection of bugs, performance issues, and security risks
• Asynchronous job processing using Redis
• Worker-based architecture
• Review history stored in PostgreSQL