#!/usr/bin/env sh
exec /api --dbAddr "${MONGODB_HOST}" --dbUser "${MONGODB_USER}" --dbPass "${MONGODB_PASS}"
