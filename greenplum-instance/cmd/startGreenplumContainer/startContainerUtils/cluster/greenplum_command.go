package cluster

import (
	"os/exec"

	"github.com/pivotal/greenplum-for-kubernetes/pkg/commandable"
)

type GreenplumCommand struct {
	command commandable.CommandFn
}

func NewGreenplumCommand(cmd commandable.CommandFn) *GreenplumCommand {
	return &GreenplumCommand{command: cmd}
}

func (g *GreenplumCommand) Command(command string, args ...string) *exec.Cmd {
	cmd := g.command(command, args...)
	cmd.Env = append(cmd.Env,
		"HOME=/home/gpadmin",
		"USER=gpadmin",
		"LOGNAME=gpadmin",
		"GPHOME=/usr/local/greenplum-db",
		"PATH=/usr/local/greenplum-db/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"LD_LIBRARY_PATH=/usr/local/greenplum-db/lib:/usr/local/greenplum-db/ext/python/lib",
		"MASTER_DATA_DIRECTORY=/greenplum/data-1",
		"PYTHONHOME=/usr/local/greenplum-db/ext/python",
		"PYTHONPATH=/usr/lib/python2.7:/usr/lib/python2.7/plat-x86_64-linux-gnu:/usr/lib/python2.7/lib-tk:/usr/lib/python2.7/lib-old:/usr/lib/python2.7/lib-dynload:/usr/local/lib/python2.7/dist-packages:/usr/lib/python2.7/dist-packages:/usr/local/greenplum-db/lib/python")
	return cmd
}
