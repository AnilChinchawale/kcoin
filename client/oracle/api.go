package oracle

// PrivateOracleAPI provides private RPC methods to control the oracle service
type PrivateOracleAPI struct {
	oracleServ *Service
}

// NewPrivateOracleAPI creates a new RPC service which controls the oracle service
func NewPrivateOracleAPI(oracleServ *Service) *PrivateOracleAPI {
	return &PrivateOracleAPI{oracleServ: oracleServ}
}

// Start the oracle service
func (api *PrivateOracleAPI) Start() error {
	if !api.oracleServ.IsReporting() {
		return api.oracleServ.StartReporting()
	}
	return nil
}

// Stop the oracle service
func (api *PrivateOracleAPI) Stop() bool {
	api.oracleServ.StopReporting()
	return true
}
