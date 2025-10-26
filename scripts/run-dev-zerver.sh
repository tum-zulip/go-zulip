#!/usr/bin/env bash

set -euo pipefail
set -m  # Enable job control

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

log "Starting Zulip dev server"

cd $REPO_DIR
source /srv/zulip-py3-venv/bin/activate || true # legacy Zulip uses a global venv
source .venv/bin/activate || true # modern Zulip uses a local venv

./tools/run-dev
