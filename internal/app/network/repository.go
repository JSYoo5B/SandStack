package network

type NetworkRepository interface {
	Create(network Network) Network
	List() []Network
	Get(id string) (Network, error)
	Update(network Network) (Network, error)
	Delete(id string) error
	Reset()
}

type SubnetRepository interface {
	Create(subnet Subnet) Subnet
	List() []Subnet
	Get(id string) (Subnet, error)
	Delete(id string) error
	Reset()
}

type PortRepository interface {
	Create(port Port) Port
	List() []Port
	Get(id string) (Port, error)
	Delete(id string) error
	Reset()
}

type SecurityGroupRepository interface {
	Create(securityGroup SecurityGroup) SecurityGroup
	List() []SecurityGroup
	Get(id string) (SecurityGroup, error)
	Delete(id string) error
	Reset()
}

type SecurityGroupRuleRepository interface {
	Create(rule SecurityGroupRule) SecurityGroupRule
	List() []SecurityGroupRule
	Get(id string) (SecurityGroupRule, error)
	Delete(id string) error
	Reset()
}

type RouterRepository interface {
	Create(router Router) Router
	List() []Router
	Get(id string) (Router, error)
	Delete(id string) error
	Reset()
}

type FloatingIPRepository interface {
	Create(floatingIP FloatingIP) FloatingIP
	List() []FloatingIP
	Get(id string) (FloatingIP, error)
	Delete(id string) error
	Reset()
}
