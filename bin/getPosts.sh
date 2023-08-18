#!/bin/bash

echo "🌟 postsディレクトリを取得します 🚀"

# 既存の「posts」ディレクトリを削除して新しく作成
echo "🧹 古い 'posts' ディレクトリを削除中..."
rm -rf posts
echo "📂 新しい 'posts' ディレクトリを作成中..."
mkdir posts

# 'posts'ディレクトリに移動
cd posts

# Gitリポジトリの初期化
echo "🌀 新しいGitリポジトリを初期化中..."
git init

# .git/config ファイルにsparsecheckoutの設定を追加
echo "🔧 sparse checkoutの設定を追加中..."
git config core.sparsecheckout true

# リモートリポジトリの追加
echo "🌌 'origin' リモートリポジトリを追加中..."
git remote add origin https://github.com/HidemaruOwO/VirtualSlime.git

# sparse-checkoutの設定
echo "🎨 sparse-checkoutのパターンを設定中..."
echo "posts" > .git/info/sparse-checkout

# リモートリポジトリからデータを取得
echo "🚚 'origin' リポジトリからデータを取得中..."
git pull origin main

echo "🎉 スクリプトの実行が完了しました。エンハンスされたコードをお楽しみください！ 🌈"

