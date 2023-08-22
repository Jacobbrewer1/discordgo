package discordgo

import "fmt"

type Operation int

const (
	// OperationDispatch - An event was dispatched.
	OperationDispatch Operation = iota

	// OperationHeartbeat - Fired periodically by the client to keep the connection alive.
	OperationHeartbeat

	// OperationIdentify - Starts a new session during the initial handshake.
	OperationIdentify

	// OperationPresenceUpdate - Update the client's presence.
	OperationPresenceUpdate

	// OperationVoiceStateUpdate - Used to join/leave or move between voice channels.
	OperationVoiceStateUpdate

	// OperationResume - Resume a previous session that was disconnected.
	OperationResume

	// OperationReconnect - You should attempt to reconnect and resume immediately.
	OperationReconnect

	// OperationRequestGuildMembers - Request information about offline guild members in a large guild.
	OperationRequestGuildMembers

	// OperationInvalidSession - The session has been invalidated. You should reconnect and identify/resume accordingly.
	OperationInvalidSession

	// OperationHello - Sent immediately after connecting, contains the heartbeat_interval to use.
	OperationHello

	// OperationHeartbeatACK - Sent in response to receiving a heartbeat to acknowledge that it has been received.
	OperationHeartbeatACK
)

// String returns a string representation of the Operation
func (o Operation) String() string {
	switch o {
	case OperationDispatch:
		return "Dispatch"
	case OperationHeartbeat:
		return "Heartbeat"
	case OperationIdentify:
		return "Identify"
	case OperationPresenceUpdate:
		return "Presence Update"
	case OperationVoiceStateUpdate:
		return "Voice State Update"
	case OperationResume:
		return "Resume"
	case OperationReconnect:
		return "Reconnect"
	case OperationRequestGuildMembers:
		return "Request Guild Members"
	case OperationInvalidSession:
		return "Invalid Session"
	case OperationHello:
		return "Hello"
	case OperationHeartbeatACK:
		return "Heartbeat ACK"
	}
	return fmt.Sprintf("Unknown_Operation_(%d)", o)
}

// MarshalJSON returns the int code for the Operation
func (o *Operation) MarshalJSON() ([]byte, error) {
	i, err := Marshal(*o)
	if err != nil {
		return nil, fmt.Errorf("error marshaling operation: %w", err)
	}
	return i, nil
}

// UnmarshalJSON receives the int code from discord and converts it to the Operation
func (o *Operation) UnmarshalJSON(b []byte) error {
	var i int
	if err := Unmarshal(b, &i); err != nil {
		return err
	}
	*o = Operation(i)
	return nil
}
