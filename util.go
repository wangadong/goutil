package ugo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	//	The string format used in LogError().
	LogErrorFormat = "%v"
)

func init() {
}

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

//	Returns all paths listed in the GOPATH environment variable.
func GoPaths() []string {
	return strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
}

//	Returns the path/filepath.Join()ed full directory path for a specified $GOPATH/src sub-directory.
//	Example: util.GopathSrc("tools", "importers", "sql") = "c:\gd\src\tools\importers\sql" if $GOPATH is c:\gd.
func GopathSrc(subDirNames ...string) (gps string) {
	gp := []string{"", "src"}
	for _, goPath := range GoPaths() { // in 99% of setups there's only 1 GOPATH, but hey..
		gp[0] = goPath
		if gps = filepath.Join(append(gp, subDirNames...)...); dirExists(gps) {
			break
		}
	}
	return
}

//	Returns the path/filepath.Join()ed full directory path for a specified $GOPATH/src/github.com sub-directory.
//	Example: util.GopathSrcGithub("metaleap", "go-util", "num") = "c:\gd\src\github.com\metaleap\go-util\num" if $GOPATH is c:\gd.
func GopathSrcGithub(gitHubName string, subDirNames ...string) string {
	return GopathSrc(append([]string{"github.com", gitHubName}, subDirNames...)...)
}

//	Returns the result of os.Hostname() if any, else "localhost".
func HostName() (hostName string) {
	if hostName, _ = os.Hostname(); len(hostName) == 0 {
		hostName = "localhost"
	}
	return
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifb(cond, ifTrue, ifFalse bool) bool {
	return (cond && ifTrue) || ((!cond) && ifFalse)
	// if cond {
	// 	return ifTrue
	// }
	// return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifd(cond bool, ifTrue, ifFalse float64) float64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi(cond bool, ifTrue, ifFalse int) int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi16(cond bool, ifTrue, ifFalse int16) int16 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi32(cond bool, ifTrue, ifFalse int32) int32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifi64(cond bool, ifTrue, ifFalse int64) int64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifs(cond bool, ifTrue string, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifu32(cond bool, ifTrue, ifFalse uint32) uint32 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	Returns ifTrue if cond is true, otherwise returns ifFalse.
func Ifu64(cond bool, ifTrue, ifFalse uint64) uint64 {
	if cond {
		return ifTrue
	}
	return ifFalse
}

//	A convenience short-hand for log.Println(fmt.Sprintf(LogErrorFormat, err)) if err isn't nil.
func LogError(err error) {
	if err != nil {
		log.Println(fmt.Sprintf(LogErrorFormat, err))
	}
}

//	Returns the human-readable representation associated with the specified GOOS name in OSNames.
func OSName(goOS string) (name string) {
	switch goOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "Mac OS X"
	case "linux":
		return "Linux"
	case "freebsd":
		return "FreeBSD"
	default:
		return "OS"
	}
	return
}

//	Attempts to extract major and minor version components from a string that begins with a version number.
//	Example: returns []int{3, 2} and float64(3.2) for a verstr that is "3.2.0 - Build 8.15.10.2761"
func ParseVersion(verstr string) (majorMinor [2]int, both float64) {
	var (
		pos, j int
		i      uint64
		err    error
	)
	for _, p := range strings.Split(verstr, ".") {
		if pos = strings.Index(p, " "); pos > 0 {
			p = p[:pos]
		}
		if i, err = strconv.ParseUint(p, 10, 8); err == nil {
			majorMinor[j] = int(i)
			if j++; j >= len(majorMinor) {
				break
			}
		} else {
			break
		}
	}
	if len(majorMinor) > 0 {
		both = float64(majorMinor[0])
	}
	if len(majorMinor) > 1 {
		both += (float64(majorMinor[1]) * 0.1)
	}
	return
}

//	Returns the path to the current user's home directory.
//	Might be C:\Users\Kitty under Windows, /home/Kitty under Linux or /Users/Kitty under OSX.
//	Specifically, returns the value of either the $userprofile (Windows) or the $HOME (others) environment variable, whichever one is set.
func UserHomeDirPath() (dirPath string) {
	if dirPath = os.Getenv("userprofile"); len(dirPath) == 0 {
		dirPath = os.Getenv("HOME")
	}
	return
}
