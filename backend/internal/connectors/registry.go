package connectors

// Registry holds all registered connectors by name.
type Registry struct {
	connectors map[string]Connector
}

func NewRegistry() *Registry {
	return &Registry{connectors: make(map[string]Connector)}
}

func (r *Registry) Register(name string, c Connector) {
	r.connectors[name] = c
}

func (r *Registry) Get(name string) (Connector, bool) {
	c, ok := r.connectors[name]
	return c, ok
}

func (r *Registry) All() map[string]Connector {
	return r.connectors
}

func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.connectors))
	for n := range r.connectors {
		names = append(names, n)
	}
	return names
}
