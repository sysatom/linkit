package bot

import (
	"github.com/sysatom/linkit/internal/pkg/client"
	"github.com/sysatom/linkit/internal/pkg/types"
	"log"
	"time"
)

const (
	DevAgentVersion = 1
	ImportAgentId   = "import_agent"
)

func DevImport(c *client.Tinode) {
	_, err := c.Agent(types.AgentContent{
		Id:      ImportAgentId,
		Version: DevAgentVersion,
		Content: types.KV{
			"time": time.Now().String(),
		},
	})
	if err != nil {
		log.Println(err)
	}
}
