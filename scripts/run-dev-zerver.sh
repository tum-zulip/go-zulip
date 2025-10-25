#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'EOF'
Usage: run-dev-zerver.sh [--ref GIT_REF] [--repo URL]

Spin up an isolated Zulip dev server instance cloned from the specified git ref.
The server runs inside its own network namespace and is reachable via the
provided loopback IP (e.g. 127.0.123.123).

Options:
  --ref   Git ref (branch, tag, or commit) to check out. Defaults to "main".
  --repo  Git repository URL. Defaults to https://github.com/zulip/zulip.git.
  -h, --help  Show this help.

Example:
  sudo ./scripts/run-dev-zerver.sh --ref main
EOF
}

die() {
  echo "[run-dev-zerver] $*" >&2
  exit 1
}

log() {
  echo "[run-dev-zerver] $*" >&2
}

REF="main"
REPO_URL="https://github.com/zulip/zulip.git"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --ref)
      [[ $# -ge 2 ]] || die "Missing value for --ref"
      REF="$2"
      shift 2
      ;;
    --repo)
      [[ $# -ge 2 ]] || die "Missing value for --repo"
      REPO_URL="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      die "Unknown argument: $1"
      ;;
  esac
done

BASE_DIR="${TMPDIR:-/tmp}/zulip-zerver"
INSTANCE_DIR="${BASE_DIR}/${REF}"
REPO_DIR="${INSTANCE_DIR}/zulip"
mkdir -p "$INSTANCE_DIR"
log "Instance files will be stored in $INSTANCE_DIR"

log "Cloning Zulip repo ($REPO_URL) into $REPO_DIR"
if [[ ! -d "$REPO_DIR/.git" ]]; then
  git clone --filter=blob:none "$REPO_URL" "$REPO_DIR"
else
  log "Repository already exists at $REPO_DIR; reusing"
fi

log "Checking out $REF"

git -C "$REPO_DIR" fetch origin "$REF" --depth=1 || true
git -C "$REPO_DIR" checkout "$REF"

log "Provisioning Zulip (this may take several minutes)"

cd $REPO_DIR && ./tools/provision
cd $REPO_DIR && ./tools/rebuild-dev-database

clients=(desdemona@zulip.com iago@zulip.com shiva@zulip.com polonius@zulip.com AARON@zulip.com hamlet@zulip.com cordelia@zulip.com)

# Function to continuously unblock clients
unblock_clients_loop() {
  cd "$REPO_DIR"
  while true; do
    for client in "${clients[@]}"; do
      python3 manage.py rate_limit unblock -e "$client" 2>/dev/null || true
    done
    sleep 1
  done
}

# Start unblocking loop in background
unblock_clients_loop &
UNBLOCK_PID=$!
log "Started rate limit unblocking loop (PID: $UNBLOCK_PID)"

# Cleanup function to kill unblock loop when script exits
cleanup() {
  log "Stopping rate limit unblocking loop"
  kill $UNBLOCK_PID 2>/dev/null || true
}
trap cleanup EXIT

log "Starting Zulip dev server"

PATH="$REPO_DIR/.venv/bin:$PATH" exec "$REPO_DIR/tools/run-dev"
