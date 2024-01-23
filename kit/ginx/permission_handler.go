package ginx

type Permission struct {
	NeedToken  bool // 是否需要token
	NeedCasbin bool // 是否进行权限  api路径权限验证
}

func (p *Permission) WithNeedToken(needToken bool) *Permission {
	p.NeedToken = needToken
	return p
}

func (p *Permission) WithNeedCasBin(needCasBin bool) *Permission {
	p.NeedCasbin = needCasBin
	return p
}
