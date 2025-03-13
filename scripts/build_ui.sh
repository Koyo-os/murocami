# !/bin/bash

npm --prefix ./web run build:css &&
npm --prefix ./web run build:js &&
templ generate internal/view