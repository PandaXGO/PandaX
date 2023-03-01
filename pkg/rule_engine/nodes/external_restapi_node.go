package nodes

type externalRestapiNode struct {
	bareNode
	RestEndpointUrlPattern         string
	RequestMethod                  string
	headers                        map[string]string
	UseSimpleClientHttpFactory     bool
	ReadTimeoutMs                  int
	MaxParallelRequestsCount       int
	UseRedisQueueForMsgPersistence bool
	trimQueue                      bool
	MaxQueueSize                   int
}

type externalRestapiNodeFactory struct{}

func (f externalRestapiNodeFactory) Name() string     { return "ExternalRestapiNode" }
func (f externalRestapiNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalRestapiNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &externalRestapiNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}
