package bot

type Controller struct {
	MsgCh chan string
	Done  chan bool
}

// func (c *Controller) Done() {
// 	c.done <- true
// }
