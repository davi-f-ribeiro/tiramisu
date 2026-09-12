// Package dashboard provides the stub management UI page.
package dashboard

import "embed"

// StubManagementHTML is the stub management page embedded from stub_management.html.
//
//go:embed stub_management.html
var StubManagementHTML embed.FS

// StubManagementContent returns the embedded HTML content.
func StubManagementContent() ([]byte, error) {
	return StubManagementHTML.ReadFile("stub_management.html")
}
