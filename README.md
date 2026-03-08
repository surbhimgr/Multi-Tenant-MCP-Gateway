# Multi-Tenant MCP Gateway

A lightweight backend gateway that routes AI tool requests for multiple tenants using the Model Context Protocol (MCP).  
This project demonstrates core backend concepts such as tenant isolation, API authentication, rate limiting, and tool routing.

## Overview

AI agents often need to call external tools such as databases, APIs, or services.  
In a multi-tenant system, these requests must be securely isolated between different tenants.

This project implements a simple MCP gateway that:

- authenticates tenants using API keys
- applies rate limiting per tenant
- routes tool calls
- returns structured responses to AI agents

## Architecture

Client (AI Agent)
      |
      v
MCP Gateway (FastAPI)
      |
      +---- Authentication (API Key)
      +---- Rate Limiting
      +---- Tool Routing
      |
      v
Tool Handlers (example tools)

## Features

- Multi-tenant API key authentication
- Basic per-tenant rate limiting
- Tool invocation routing
- Simple JSON-based MCP style request/response
- Lightweight implementation using FastAPI

## Tech Stack

- Python
- FastAPI
- Redis (optional for production rate limiting)

## Installation

Clone the repository
