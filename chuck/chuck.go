package chuck

import (
	"fmt"

	"github.com/fgalilee/bob/bob"
)

func Chuck() string {
	return fmt.Sprintf("Chuck 1.1.0: <%s>", bob.Bob())
}
