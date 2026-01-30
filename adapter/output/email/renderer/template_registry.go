package renderer

// pathTemplates maps email template identifiers to their corresponding
// HTML template file paths.
//
// This centralizes the relationship between a logical template ID
// (used by the application layer) and the physical HTML file
// used by the renderer. The renderer relies on this map to resolve
// which template should be loaded and executed for a given email type.
//
// Adding a new email template requires only:
// 1. Creating the HTML file under the templates directory.
// 2. Registering its template ID and path in this map.
var pathTemplates = map[string]string{
	"activation_code":   "./../templates/activation_code.html",
	"notify_activation": "./../templates/notify_activation.html",

	"change_email_code":   "./../templates/change_email_code.html",
	"notify_change_email": "./../templates/notify_change_email.html",

	"change_password_code":   "./../templates/change_password_code.html",
	"notify_change_password": "./../templates/notify_change_password.html",

	"reset_password_code":   "./../templates/reset_password_code.html",
	"notify_reset_password": "./../templates/notify_reset_password.html",

	"deletion_code":   "./../templates/deletion_code.html",
	"notify_deletion": "./../templates/notify_deletion.html",
}

