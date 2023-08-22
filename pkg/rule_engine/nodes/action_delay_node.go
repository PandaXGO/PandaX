package nodes

import (
	"fmt"
	"pandax/pkg/rule_engine/message"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const DelayNodeName = "DelayNode"

type delayNode struct {
	bareNode
	PeriodTs           int               `json:"periodTs" yaml:"periodTs" jpath:"periodTs"`                               //周期时间
	MaxPendingMessages int               `json:"maxPendingMessages" yaml:"maxPendingMessages" jpath:"maxPendingMessages"` //最大等待消息数
	messageQueue       []message.Message `jpath:"-"`
	delayTimer         *time.Timer       `jpath:"-"`
	lock               sync.Mutex        `jpath:"-"`
}

type delayNodeFactory struct{}

func (f delayNodeFactory) Name() string     { return DelayNodeName }
func (f delayNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f delayNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f delayNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &delayNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
		lock:     sync.Mutex{},
	}
	_, err := decodePath(meta, node)
	node.messageQueue = make([]message.Message, node.MaxPendingMessages)
	return node, err
}

func (n *delayNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if successLabelNode == nil || failureLabelNode == nil {
		return fmt.Errorf("no valid label linked node in %s", n.Name())
	}

	// check wethere the time had already been started, queue message if started
	if n.delayTimer == nil {
		n.messageQueue = append(n.messageQueue, msg)
		n.delayTimer = time.NewTimer(time.Duration(n.PeriodTs) * time.Second)

		go func(n *delayNode) error {
			defer n.delayTimer.Stop()
			for {
				<-n.delayTimer.C
				n.lock.Lock()
				defer n.lock.Unlock()
				if len(n.messageQueue) > 0 {
					msg := n.messageQueue[0]
					n.messageQueue = n.messageQueue[0:]
					return successLabelNode.Handle(msg)
				}
			}
		}(n)
		return nil
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	if len(n.messageQueue) == n.MaxPendingMessages {
		return failureLabelNode.Handle(msg)
	}
	n.messageQueue = append(n.messageQueue, msg)
	return nil
}
