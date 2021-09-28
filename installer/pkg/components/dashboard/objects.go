package dashboard

import "github.com/gitpod-io/gitpod/installer/pkg/common"

var Objects = common.CompositeRenderFunc(
	deployment,
	networkpolicy,
	rolebinding,
	common.GenerateService(Component),
	common.DefaultServiceAccount(Component),
)
