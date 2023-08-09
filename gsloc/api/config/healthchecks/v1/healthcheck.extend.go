package hcconf

func (x *HealthCheckPayload) GetData() []byte {
	if x == nil {
		return nil
	}
	if len(x.GetBinary()) > 0 {
		return x.GetBinary()
	}
	return []byte(x.GetText())
}
