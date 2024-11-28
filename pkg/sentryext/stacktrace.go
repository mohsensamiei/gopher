package sentryext

import (
	"github.com/getsentry/sentry-go"
	"go/build"
	"runtime"
	"strings"
)

const unknown string = "unknown"

var goRoot = strings.ReplaceAll(build.Default.GOROOT, "\\", "/")

func isCompilerGeneratedSymbol(name string) bool {
	// In versions of Go 1.20 and above a prefix of "type:" and "go:" is a
	// compiler-generated symbol that doesn't belong to any package.
	// See variable reservedimports in cmd/compile/internal/gc/subr.go
	if strings.HasPrefix(name, "go:") || strings.HasPrefix(name, "type:") {
		return true
	}
	return false
}

func packageName(name string) string {
	if isCompilerGeneratedSymbol(name) {
		return ""
	}

	pathend := strings.LastIndex(name, "/")
	if pathend < 0 {
		pathend = 0
	}

	if i := strings.Index(name[pathend:], "."); i != -1 {
		return name[:pathend+i]
	}
	return ""
}

func splitQualifiedFunctionName(name string) (pkg string, fun string) {
	pkg = packageName(name)
	if len(pkg) > 0 {
		fun = name[len(pkg)+1:]
	}
	return
}

func extractFrames(pcs []uintptr) []runtime.Frame {
	var frames = make([]runtime.Frame, 0, len(pcs))
	callersFrames := runtime.CallersFrames(pcs)

	for {
		callerFrame, more := callersFrames.Next()

		frames = append(frames, callerFrame)

		if !more {
			break
		}
	}

	// TODO don't append and reverse, put in the right place from the start.
	// reverse
	for i, j := 0, len(frames)-1; i < j; i, j = i+1, j-1 {
		frames[i], frames[j] = frames[j], frames[i]
	}

	return frames
}

func createFrames(frames []runtime.Frame) []sentry.Frame {
	if len(frames) == 0 {
		return nil
	}

	result := make([]sentry.Frame, 0, len(frames))

	for _, frame := range frames {
		function := frame.Function
		var pkg string
		if function != "" {
			pkg, function = splitQualifiedFunctionName(function)
		}

		if !shouldSkipFrame(pkg) {
			result = append(result, newFrame(pkg, function, frame.File, frame.Line))
		}
	}

	return result
}

func isAbsPath(path string) bool {
	if len(path) == 0 {
		return false
	}

	// If the volume name starts with a double slash, this is an absolute path.
	if len(path) >= 1 && (path[0] == '/' || path[0] == '\\') {
		return true
	}

	// Windows absolute path, see https://learn.microsoft.com/en-us/dotnet/standard/io/file-path-formats
	if len(path) >= 3 && path[1] == ':' && (path[2] == '/' || path[2] == '\\') {
		return true
	}

	return false
}

func newFrame(module string, function string, file string, line int) sentry.Frame {
	frame := sentry.Frame{
		Lineno:   line,
		Module:   module,
		Function: function,
	}

	switch {
	case len(file) == 0:
		frame.Filename = unknown
		// Leave abspath as the empty string to be omitted when serializing event as JSON.
	case isAbsPath(file):
		frame.AbsPath = file
		// TODO: in the general case, it is not trivial to come up with a
		// "project relative" path with the data we have in run time.
		// We shall not use filepath.Base because it creates ambiguous paths and
		// affects the "Suspect Commits" feature.
		// For now, leave relpath empty to be omitted when serializing the event
		// as JSON. Improve this later.
	default:
		// f.File is a relative path. This may happen when the binary is built
		// with the -trimpath flag.
		frame.Filename = file
		// Omit abspath when serializing the event as JSON.
	}

	setInAppFrame(&frame)

	return frame
}

func setInAppFrame(frame *sentry.Frame) {
	if strings.HasPrefix(frame.AbsPath, goRoot) ||
		strings.Contains(frame.Module, "vendor") ||
		strings.Contains(frame.Module, "third_party") {
		frame.InApp = false
	} else {
		frame.InApp = true
	}
}

func shouldSkipFrame(module string) bool {
	// Skip Go internal frames.
	if module == "runtime" || module == "testing" {
		return true
	}

	// Skip Sentry internal frames, except for frames in _test packages (for testing).
	if strings.HasPrefix(module, "github.com/getsentry/sentry-go") &&
		!strings.HasSuffix(module, "_test") {
		return true
	}

	return false
}

func NewStacktrace(skip int) *sentry.Stacktrace {
	pcs := make([]uintptr, 100)
	n := runtime.Callers(skip, pcs)

	if n == 0 {
		return nil
	}
	println(n)

	runtimeFrames := extractFrames(pcs[:n])
	frames := createFrames(runtimeFrames)

	stacktrace := sentry.Stacktrace{
		Frames: frames,
	}

	return &stacktrace
}
