#!/usr/bin/env bash
set -euo pipefail

pattern=${GO_MODEL_TEST_PATTERN:-'Test.*(Marshal|Unmarshal|Snapshot|Presence|Union|Equals|Validate)'}

# Run the targeted model/unit tests without counting cache
exec go test ./zulip -run "${pattern}" -count=1 -timeout 5m
