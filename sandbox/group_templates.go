// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2018 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sandbox

// Templates for group replication

var (
	init_nodes_template string = `#!/bin/sh
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
multi_sb={{.SandboxDir}}
# workaround for Bug#89959
{{range .Nodes}}
{{.SandboxDir}}/{{.NodeLabel}}{{.Node}}/use -h {{.MasterIp}} -u {{.RplUser}} -p{{.RplPassword}} -e 'set @a=1'
{{end}}
[ -z "$SLEEP_TIME" ] && SLEEP_TIME=1
{{range .Nodes}}
    user_cmd='reset master;'
    user_cmd="$user_cmd CHANGE MASTER TO MASTER_USER='rsandbox', MASTER_PASSWORD='rsandbox' {{.ChangeMasterExtra}} FOR CHANNEL 'group_replication_recovery';"
	echo "# Node {{.Node}} # $user_cmd"
    $multi_sb/{{.NodeLabel}}{{.Node}}/use -u root -e "$user_cmd"
{{end}}
echo ""

BEFORE_START_CMD="SET GLOBAL group_replication_bootstrap_group=ON;"
START_CMD="START GROUP_REPLICATION;"
AFTER_START_CMD="SET GLOBAL group_replication_bootstrap_group=OFF;"
echo "# Node 1 # $BEFORE_START_CMD"
$multi_sb/n1 -e "$BEFORE_START_CMD"
{{ range .Nodes}}
	echo "# Node {{.Node}} # $START_CMD"
	$multi_sb/n{{.Node}} -e "$START_CMD"
	sleep $SLEEP_TIME
{{end}}
echo "# Node 1 # $AFTER_START_CMD"
$multi_sb/n1 -e "$AFTER_START_CMD"
$multi_sb/check_nodes
`
	check_nodes_template string = `#!/bin/sh
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
multi_sb={{.SandboxDir}}
[ -z "$SLEEP_TIME" ] && SLEEP_TIME=1

CHECK_NODE="select * from performance_schema.replication_group_members"
{{ range .Nodes}}
	echo "# Node {{.Node}} # $CHECK_NODE"
	$multi_sb/{{.NodeLabel}}{{.Node}}/use -t -e "$CHECK_NODE"
	sleep $SLEEP_TIME
{{end}}
`
	GroupTemplates = TemplateCollection{
		"init_nodes_template": TemplateDesc{
			Description: "Initialize group replication after deployment",
			Notes:       "",
			Contents:    init_nodes_template,
		},
		"check_nodes_template": TemplateDesc{
			Description: "Checks the status of group replication",
			Notes:       "",
			Contents:    check_nodes_template,
		},
	}
)
