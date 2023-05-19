package main

import (
	"testing"
)

func TestFindCommandPath(t *testing.T) {
	// このテストでは、"go" コマンドのパスが見つかることを確認します。
	// "go" コマンドは、このテストを実行する環境にインストールされていると想定しています。
	path, err := findCommandPath("go")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if path == "" {
		t.Fatalf("expected a path, got empty string")
	}

	// 存在しないコマンド名を指定した場合、エラーが返ることを確認します。
	_, err = findCommandPath("nonexistentcommand")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}
