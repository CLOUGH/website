#!/bin/bash

# Seed the database with content
# Usage: ./scripts/seed.sh [--clear]

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

CLEAR_FIRST=false
if [[ "$1" == "--clear" ]]; then
    CLEAR_FIRST=true
fi

echo "🌱 Seeding database with content..."
echo ""

cd "$PROJECT_ROOT"

# Check if credentials exist
if [ ! -f "credentials/firebase-service-account.json" ]; then
    echo "❌ Error: Firebase credentials not found!"
    echo "   Please ensure credentials/firebase-service-account.json exists"
    exit 1
fi

# Check if .env exists
if [ ! -f ".env" ]; then
    echo "⚠️  Warning: .env file not found"
    echo "   Copy env.example to .env and configure it"
fi

export GOOGLE_APPLICATION_CREDENTIALS="credentials/firebase-service-account.json"

# Clear existing sections if requested
if [ "$CLEAR_FIRST" = true ]; then
    echo "🗑️  Clearing existing sections..."
    go run scripts/clear-sections.go
    echo ""
fi

# Run the seed script
echo "📦 Creating new sections..."
go run scripts/seed-full-content.go

echo ""
echo "✅ Database seeded successfully!"
echo ""
echo "Next steps:"
echo "  • Start the server: make run"
echo "  • Visit admin dashboard: http://localhost:8080/admin/dashboard"
echo "  • Edit content directly in the database"
