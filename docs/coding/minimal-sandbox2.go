// This is a sample source file that shows how
// to create two MySQL sandboxes using dbdeployer code
// from another Go program.
package main

import (
	"github.com/datacharmer/dbdeployer/common"
	"github.com/datacharmer/dbdeployer/defaults"
	"github.com/datacharmer/dbdeployer/sandbox"
	"os"
)

func main() {
	// Searches for expanded sandboxes in $HOME/opt/mysql
	sandbox_binary := os.Getenv("HOME") + "/opt/mysql"

	// Creates sandboxes in $HOME/sandboxes
	sandbox_home := os.Getenv("HOME") + "/sandboxes"

	// For this to work, we need to have
	// a MySQL tarball expanded in $HOME/opt/mysql/5.7.22
	version1 := "5.7.22"
	version2 := "5.6.25"

	sandbox_name1 := "msb_5_7_22"
	sandbox_name2 := "msb_5_6_25"

	// The unique ports for these sandboxes
	port1 := 5722
	port2 := 5625

	// MySQL will look for binaries in $HOME/opt/mysql/5.7.22
	basedir1 := sandbox_binary + "/" + version1 // This is what dbdeployer expects
	// i.e. a name containing the full version

	// MySQL will look for binaries in $HOME/opt/mysql/my-5.6
	basedir2 := sandbox_binary + "/my-5.6" // This is a deviation from dbdeployer
	// paradigm, using a non-standard name
	// for the base directory

	// Username and password for this sandbox
	user := "msandbox"
	password := "msandbox"

	// Creates the base target directory if it doesn't exist
	if !common.DirExists(sandbox_home) {
		common.Mkdir(sandbox_home)
	}

	// Minimum data to be filled for a simple sandbox.
	// See sandbox/sandbox.go for the full description
	// of this data structure
	var sdef = sandbox.SandboxDef{
		Version:    version1,
		Basedir:    basedir1,
		SandboxDir: sandbox_home,
		DirName:    sandbox_name1,
		LoadGrants: true,
		// This is the list of ports to ignore
		// when checking if the designated port is
		// used or not.
		// Try changing the Port item to 3306.
		// You will see that the sandbox will install using 3307
		InstalledPorts: []int{1186, 3306, 33060},
		Port:           port1,
		DbUser:         user,
		DbPassword:     password,
		RplUser:        "r" + user,
		RplPassword:    "r" + password,
		RemoteAccess:   "127.%",
		BindAddress:    "127.0.0.1",
	}

	// Calls the sandbox creation
	sandbox.CreateSingleSandbox(sdef)

	sdef.Version = version2
	sdef.Basedir = basedir2
	sdef.DirName = sandbox_name2
	sdef.Port = port2

	// Calls the sandbox creation for the second sandbox
	sandbox.CreateSingleSandbox(sdef)

	// Invokes the sandbox self-testing script
	common.Run_cmd(sandbox_home + "/" + sandbox_name1 + "/test_sb")
	common.Run_cmd(sandbox_home + "/" + sandbox_name2 + "/test_sb")

	// Removes the sandbox from disk
	sandbox.RemoveSandbox(sandbox_home, sandbox_name1, false)
	sandbox.RemoveSandbox(sandbox_home, sandbox_name2, false)

	// Removes the sandbox from dbdeployer catalog
	defaults.DeleteFromCatalog(sandbox_home + "/" + sandbox_name1)
	defaults.DeleteFromCatalog(sandbox_home + "/" + sandbox_name2)
}
