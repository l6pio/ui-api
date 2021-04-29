#!/usr/bin/env sh
exec /ui-api --dbAddr "${MONGODB_HOST}" --dbUser "${MONGODB_USER}" --dbPass "${MONGODB_PASS}"
