package pathutil

import (
	"testing"
)

func TestGetLastPathComponent(t *testing.T) {
	t.Logf("Testing last path component")
	lastPathComponent := GetLastPathComponent("./Something/Hello.csproj")
	if lastPathComponent != "Hello.csproj" {
		t.Errorf("%s != Hello.csproj", lastPathComponent)
	} else {
		t.Logf("%s == Hello.csproj", lastPathComponent)
	}
}

func TestNormalizePath(t *testing.T) {
	t.Logf("Testing normalize path")
	normalizedPath := NormalizePath("./Something\\Hello.csproj")
	if normalizedPath != "./Something/Hello.csproj" {
		t.Errorf("%s not normalized", normalizedPath)
	} else {
		t.Logf("%s normalized", normalizedPath)
	}
}

func TestPathParser_GetPathString(t *testing.T) {
	t.Logf("Testing removal of last n-actions")
	parser := NewPathParserFromString("./Something/Hello.csproj")
	pathString := parser.GetPathString(true)
	if pathString != "./Something/Hello.csproj" {
		t.Errorf("%s != ./Something/Hello.csproj", pathString)
	} else {
		t.Logf("%s == ./Something/Hello.csproj", pathString)
	}
}
