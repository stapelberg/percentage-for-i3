// Program percentage-for-i3 resizes the current window to use
// -target_percentage (default: 60%) of its parent container.
package main

import (
	"flag"
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

var (
	target = flag.Float64("target_percentage", 0.6, "Target percentage which the parent container of the currently focused container should be resized to")
)

func main() {
	flag.Parse()

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	// Get the parent of the currently focused container.
	parent := tree.Root.FindFocused(func(n *i3.Node) bool {
		for _, c := range n.Nodes {
			if c.Focused {
				return true
			}
		}
		return false
	})

	diff := *target - parent.Percent
	verb := "grow"
	if diff < 0 {
		verb = "shrink"
		diff = diff * -1
	}
	resize := fmt.Sprintf(`[con_id="%d"] resize %s width 10 px or %.0f ppt`, parent.ID, verb, diff*100)
	if _, err := i3.RunCommand(resize); err != nil {
		log.Fatal(err)
	}
}
