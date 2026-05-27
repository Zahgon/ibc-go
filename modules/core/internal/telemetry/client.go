package telemetry

func ReportCreateClient(clientType string) { _ = "STUB: not implemented"; return }

func ReportUpdateClient(foundMisbehaviour bool, clientType, clientID string) {
	_ = "STUB: not implemented"
	return
}

func ReportUpgradeClient(clientType, clientID string) { _ = "STUB: not implemented"; return }

func ReportRecoverClient(clientType, subjectClientID string) { _ = "STUB: not implemented"; return }
