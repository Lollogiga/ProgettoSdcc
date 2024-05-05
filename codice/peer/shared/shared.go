package shared

import "codice/peer/registry"

// Variabili di configurazione:

var PeerList []registry.PeerInfo
var Port int
var MyId, LeaderId int32

// variabili per algoritmo di Dolev
var Token, MyToken = 0, -1
var NumNode = 0

// variabili per algoritmo DKR
var StimateLeader int32
var State = "active"
