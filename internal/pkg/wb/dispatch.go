package wb

import (
	"fmt"
	"github.com/sysatom/linkit/internal/pkg/types"
)

func (s *Session) dispatch(msg *types.ServerComMessage) {
	fmt.Println(msg)
	// todo msg.data to instruct
	// todo instruct.RunInstruct(app, window, cache, item)
}
