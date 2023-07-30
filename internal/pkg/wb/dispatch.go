package wb

import (
	"fmt"
	"github.com/sysatom/linkit/internal/pkg/types"
)

func (s *Session) dispatch(msg *types.ServerComMessage) {
	fmt.Println(msg)
}
