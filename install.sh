#!/usr/bin/env bash

set -e

APP_NAME="nanta"
INSTALL_DIR="$HOME/.local/bin"
CONFIG_DIR="$HOME/.config/$APP_NAME"
NOTES_DIR="$HOME/notes"
CONFIG_FILE="$CONFIG_DIR/config.yaml"
TEMPLATE_DIR="$CONFIG_DIR/templates"

echo "🔧 Installing $APP_NAME..."

mkdir -p "$INSTALL_DIR"
cp "./$APP_NAME" "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/$APP_NAME"

mkdir -p "$CONFIG_DIR"
if [ ! -f "$CONFIG_FILE" ]; then
    cat <<EOF > "$CONFIG_FILE"
notes_dir: "$NOTES_DIR"
savemode: "daily"
editor: "nvim"
previewer: "glow"
EOF
    echo "✅ Created config at $CONFIG_FILE"
fi

mkdir -p "$NOTES_DIR"

mkdir -p "$TEMPLATE_DIR"
cp -n ./templates/*.md "$TEMPLATE_DIR/"
echo "✅ Copied templates to $TEMPLATE_DIR"

echo "🎉 $APP_NAME installed successfully to $INSTALL_DIR"
