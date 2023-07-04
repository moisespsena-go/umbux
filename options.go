package umbux

// Options defines template output behavior.
type Options struct {
	// Setting if pretty printing is enabled.
	// Pretty printing ensures that the output html is properly indented and in human readable form.
	// If disabled, produced HTML is compact. This might be more suitable in production environments.
	// Default: true
	PrettyPrint bool
	// Setting if line number emitting is enabled
	// In this form, Umbux emits line number comments in the output template. It is usable in debugging environments.
	// Default: false
	LineNumbers bool
	// Setting Builtin funcs names
	// If set, when identifier matches key, disable to DIT convertion.
	// Default: nil
	BuiltinFuncsNames map[string]any
}

// DirOptions is used to provide options to directory compilation.
type DirOptions struct {
	// File extension to match for compilation
	Ext string
	// Whether or not to walk subdirectories
	Recursive bool
}

// DefaultOptions sets pretty-printing to true and line numbering to false.
var DefaultOptions = Options{true, false, nil}

// DefaultDirOptions sets expected file extension to ".pug" and recursive search for templates within a directory to true.
var DefaultDirOptions = DirOptions{".pug", true}
