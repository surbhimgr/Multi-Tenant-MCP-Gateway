from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from datetime import datetime
from collections import defaultdict
import time

app = FastAPI(title="Multi-Tenant MCP Gateway")

# -------------------------
# Fake tenant API keys
# -------------------------
TENANTS = {
    "tenant_123": "CompanyA",
    "tenant_456": "CompanyB"
}

# -------------------------
# Rate limiting store
# -------------------------
rate_limit_store = defaultdict(list)
RATE_LIMIT = 5  # requests
WINDOW = 60     # seconds


# -------------------------
# Request schema
# -------------------------
class ToolRequest(BaseModel):
    api_key: str
    tool: str
    payload: dict | None = None


# -------------------------
# Rate limiting logic
# -------------------------
def check_rate_limit(api_key):
    now = time.time()

    requests = rate_limit_store[api_key]

    # remove old requests
    rate_limit_store[api_key] = [
        r for r in requests if now - r < WINDOW
    ]

    if len(rate_limit_store[api_key]) >= RATE_LIMIT:
        raise HTTPException(status_code=429, detail="Rate limit exceeded")

    rate_limit_store[api_key].append(now)


# -------------------------
# Tool Handlers
# -------------------------
def get_time():
    return str(datetime.utcnow())


def echo(payload):
    return payload


TOOLS = {
    "get_time": get_time,
    "echo": echo
}


# -------------------------
# MCP Gateway Endpoint
# -------------------------
@app.post("/tool")
def call_tool(request: ToolRequest):

    # Authenticate tenant
    if request.api_key not in TENANTS:
        raise HTTPException(status_code=401, detail="Invalid API key")

    tenant_name = TENANTS[request.api_key]

    # Rate limiting
    check_rate_limit(request.api_key)

    # Tool routing
    if request.tool not in TOOLS:
        raise HTTPException(status_code=404, detail="Tool not found")

    tool_function = TOOLS[request.tool]

    if request.payload:
        result = tool_function(request.payload)
    else:
        result = tool_function()

    return {
        "tenant": tenant_name,
        "tool": request.tool,
        "result": result
    }
