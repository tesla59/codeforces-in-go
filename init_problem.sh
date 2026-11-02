#!/usr/bin/env bash

set -e

# Resolve script root directory
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SNIPPET_FILE="${REPO_ROOT}/snippet.go"

if [ ! -f "$SNIPPET_FILE" ]; then
    echo "Error: snippet.go not found at ${SNIPPET_FILE}" >&2
    exit 1
fi

# Usage message
usage() {
    echo "Usage: $0 <codeforces-url-or-problem-id> [problem-letter]"
    echo ""
    echo "Examples:"
    echo "  $0 https://codeforces.com/problemset/problem/69/A"
    echo "  $0 https://codeforces.com/contest/2276/problem/A"
    echo "  $0 69/A"
    echo "  $0 69A"
    echo "  $0 69 A"
    exit 1
}

INPUT="$1"

if [ -z "$INPUT" ]; then
    read -rp "Enter Codeforces URL or problem (e.g., 69A or URL): " INPUT
fi

if [ -z "$INPUT" ]; then
    usage
fi

CONTEST=""
PROBLEM=""

# If two arguments are provided (e.g. ./new.sh 69 A)
if [ -n "$2" ]; then
    CONTEST="$1"
    PROBLEM="$2"
else
    # Parse URL formats
    if [[ "$INPUT" =~ codeforces\.com/(problemset/problem|contest)/([0-9]+)/?(problem/)?([A-Za-z0-9]+) ]]; then
        CONTEST="${BASH_REMATCH[2]}"
        PROBLEM="${BASH_REMATCH[4]}"
    # Parse "69/A"
    elif [[ "$INPUT" =~ ^([0-9]+)/([A-Za-z0-9]+)$ ]]; then
        CONTEST="${BASH_REMATCH[1]}"
        PROBLEM="${BASH_REMATCH[2]}"
    # Parse "69A" or "1791G1"
    elif [[ "$INPUT" =~ ^([0-9]+)([A-Za-z][0-9]?)$ ]]; then
        CONTEST="${BASH_REMATCH[1]}"
        PROBLEM="${BASH_REMATCH[2]}"
    fi
fi

if [ -z "$CONTEST" ] || [ -z "$PROBLEM" ]; then
    echo "Error: Could not parse contest and problem from input: '$INPUT'" >&2
    usage
fi

# Normalize problem letter to uppercase
PROBLEM=$(echo "$PROBLEM" | tr '[:lower:]' '[:upper:]')

# Calculate range directory (e.g. 1-99, 100-199, 2200-2299)
if [ "$CONTEST" -lt 100 ]; then
    RANGE="1-99"
else
    START=$(( (CONTEST / 100) * 100 ))
    END=$(( START + 99 ))
    RANGE="${START}-${END}"
fi

TARGET_DIR="${REPO_ROOT}/${RANGE}/${CONTEST}/${PROBLEM}"
TARGET_FILE="${TARGET_DIR}/main.go"

# Check if solution already exists to avoid accidental overwrites
if [ -f "$TARGET_FILE" ]; then
    echo "Warning: File already exists at: ${TARGET_FILE}"
    exit 0
fi

# Create target directories
mkdir -p "$TARGET_DIR"

# Determine problem URL for header comment
if [[ "$INPUT" =~ ^https?:// ]]; then
    PROBLEM_URL="$INPUT"
else
    PROBLEM_URL="https://codeforces.com/problemset/problem/${CONTEST}/${PROBLEM}"
fi

# Write file with header comment followed by snippet.go
{
    echo "// ${PROBLEM_URL}"
    echo ""
    cat "$SNIPPET_FILE"
} > "$TARGET_FILE"

echo "Initialized: ${RANGE}/${CONTEST}/${PROBLEM}/main.go"

