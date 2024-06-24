func (s *ServiceImpl) CreateGroup(group models.Group) error {
	s.groups[group.Name] = struct{}{}
	s.cancelFunc()
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx = ctx
	s.cancelFunc = cancel
	return nil
}
