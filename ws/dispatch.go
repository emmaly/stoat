package ws

import "context"

// EventHandler is the interface for handling server-to-client WebSocket events.
// Embed DefaultEventHandler and override only the methods you need.
type EventHandler interface {
	OnAuthenticated(AuthenticatedEvent)
	OnReady(ReadyEvent)
	OnError(ErrorEvent)
	OnLogout(LogoutEvent)
	OnPong(PongEvent)
	OnMessage(MessageEvent)
	OnMessageUpdate(MessageUpdateEvent)
	OnMessageAppend(MessageAppendEvent)
	OnMessageDelete(MessageDeleteEvent)
	OnMessageReact(MessageReactEvent)
	OnMessageUnreact(MessageUnreactEvent)
	OnMessageRemoveReaction(MessageRemoveReactionEvent)
	OnChannelCreate(ChannelCreateEvent)
	OnChannelUpdate(ChannelUpdateEvent)
	OnChannelDelete(ChannelDeleteEvent)
	OnChannelGroupJoin(ChannelGroupJoinEvent)
	OnChannelGroupLeave(ChannelGroupLeaveEvent)
	OnChannelStartTyping(ChannelStartTypingEvent)
	OnChannelStopTyping(ChannelStopTypingEvent)
	OnChannelAck(ChannelAckEvent)
	OnServerCreate(ServerCreateEvent)
	OnServerUpdate(ServerUpdateEvent)
	OnServerDelete(ServerDeleteEvent)
	OnServerMemberUpdate(ServerMemberUpdateEvent)
	OnServerMemberJoin(ServerMemberJoinEvent)
	OnServerMemberLeave(ServerMemberLeaveEvent)
	OnServerRoleUpdate(ServerRoleUpdateEvent)
	OnServerRoleDelete(ServerRoleDeleteEvent)
	OnUserUpdate(UserUpdateEvent)
	OnUserRelationship(UserRelationshipEvent)
	OnUserPlatformWipe(UserPlatformWipeEvent)
	OnEmojiCreate(EmojiCreateEvent)
	OnEmojiDelete(EmojiDeleteEvent)
	OnAuth(AuthEvent)
}

// DefaultEventHandler implements EventHandler with no-op methods.
// Embed it in your handler and override only what you need.
type DefaultEventHandler struct{}

func (DefaultEventHandler) OnAuthenticated(AuthenticatedEvent)               {}
func (DefaultEventHandler) OnReady(ReadyEvent)                               {}
func (DefaultEventHandler) OnError(ErrorEvent)                               {}
func (DefaultEventHandler) OnLogout(LogoutEvent)                             {}
func (DefaultEventHandler) OnPong(PongEvent)                                 {}
func (DefaultEventHandler) OnMessage(MessageEvent)                           {}
func (DefaultEventHandler) OnMessageUpdate(MessageUpdateEvent)               {}
func (DefaultEventHandler) OnMessageAppend(MessageAppendEvent)               {}
func (DefaultEventHandler) OnMessageDelete(MessageDeleteEvent)               {}
func (DefaultEventHandler) OnMessageReact(MessageReactEvent)                 {}
func (DefaultEventHandler) OnMessageUnreact(MessageUnreactEvent)             {}
func (DefaultEventHandler) OnMessageRemoveReaction(MessageRemoveReactionEvent) {}
func (DefaultEventHandler) OnChannelCreate(ChannelCreateEvent)               {}
func (DefaultEventHandler) OnChannelUpdate(ChannelUpdateEvent)               {}
func (DefaultEventHandler) OnChannelDelete(ChannelDeleteEvent)               {}
func (DefaultEventHandler) OnChannelGroupJoin(ChannelGroupJoinEvent)         {}
func (DefaultEventHandler) OnChannelGroupLeave(ChannelGroupLeaveEvent)       {}
func (DefaultEventHandler) OnChannelStartTyping(ChannelStartTypingEvent)     {}
func (DefaultEventHandler) OnChannelStopTyping(ChannelStopTypingEvent)       {}
func (DefaultEventHandler) OnChannelAck(ChannelAckEvent)                     {}
func (DefaultEventHandler) OnServerCreate(ServerCreateEvent)                 {}
func (DefaultEventHandler) OnServerUpdate(ServerUpdateEvent)                 {}
func (DefaultEventHandler) OnServerDelete(ServerDeleteEvent)                 {}
func (DefaultEventHandler) OnServerMemberUpdate(ServerMemberUpdateEvent)     {}
func (DefaultEventHandler) OnServerMemberJoin(ServerMemberJoinEvent)         {}
func (DefaultEventHandler) OnServerMemberLeave(ServerMemberLeaveEvent)       {}
func (DefaultEventHandler) OnServerRoleUpdate(ServerRoleUpdateEvent)         {}
func (DefaultEventHandler) OnServerRoleDelete(ServerRoleDeleteEvent)         {}
func (DefaultEventHandler) OnUserUpdate(UserUpdateEvent)                     {}
func (DefaultEventHandler) OnUserRelationship(UserRelationshipEvent)         {}
func (DefaultEventHandler) OnUserPlatformWipe(UserPlatformWipeEvent)         {}
func (DefaultEventHandler) OnEmojiCreate(EmojiCreateEvent)                   {}
func (DefaultEventHandler) OnEmojiDelete(EmojiDeleteEvent)                   {}
func (DefaultEventHandler) OnAuth(AuthEvent)                                 {}

// Listen reads events from the WebSocket connection and dispatches them to the
// handler. It blocks until the context is cancelled or the connection closes.
func (c *Conn) Listen(ctx context.Context, handler EventHandler) error {
	for {
		ev, err := c.ReadEvent(ctx)
		if err != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return err
			}
		}
		dispatchEvent(handler, ev)
	}
}

func dispatchEvent(handler EventHandler, ev Event) {
	switch e := ev.(type) {
	case *AuthenticatedEvent:
		handler.OnAuthenticated(*e)
	case *ReadyEvent:
		handler.OnReady(*e)
	case *ErrorEvent:
		handler.OnError(*e)
	case *LogoutEvent:
		handler.OnLogout(*e)
	case *PongEvent:
		handler.OnPong(*e)
	case *BulkEvent:
		for _, sub := range e.V {
			if sub.Value != nil {
				dispatchEvent(handler, sub.Value)
			}
		}
	case *MessageEvent:
		handler.OnMessage(*e)
	case *MessageUpdateEvent:
		handler.OnMessageUpdate(*e)
	case *MessageAppendEvent:
		handler.OnMessageAppend(*e)
	case *MessageDeleteEvent:
		handler.OnMessageDelete(*e)
	case *MessageReactEvent:
		handler.OnMessageReact(*e)
	case *MessageUnreactEvent:
		handler.OnMessageUnreact(*e)
	case *MessageRemoveReactionEvent:
		handler.OnMessageRemoveReaction(*e)
	case *ChannelCreateEvent:
		handler.OnChannelCreate(*e)
	case *ChannelUpdateEvent:
		handler.OnChannelUpdate(*e)
	case *ChannelDeleteEvent:
		handler.OnChannelDelete(*e)
	case *ChannelGroupJoinEvent:
		handler.OnChannelGroupJoin(*e)
	case *ChannelGroupLeaveEvent:
		handler.OnChannelGroupLeave(*e)
	case *ChannelStartTypingEvent:
		handler.OnChannelStartTyping(*e)
	case *ChannelStopTypingEvent:
		handler.OnChannelStopTyping(*e)
	case *ChannelAckEvent:
		handler.OnChannelAck(*e)
	case *ServerCreateEvent:
		handler.OnServerCreate(*e)
	case *ServerUpdateEvent:
		handler.OnServerUpdate(*e)
	case *ServerDeleteEvent:
		handler.OnServerDelete(*e)
	case *ServerMemberUpdateEvent:
		handler.OnServerMemberUpdate(*e)
	case *ServerMemberJoinEvent:
		handler.OnServerMemberJoin(*e)
	case *ServerMemberLeaveEvent:
		handler.OnServerMemberLeave(*e)
	case *ServerRoleUpdateEvent:
		handler.OnServerRoleUpdate(*e)
	case *ServerRoleDeleteEvent:
		handler.OnServerRoleDelete(*e)
	case *UserUpdateEvent:
		handler.OnUserUpdate(*e)
	case *UserRelationshipEvent:
		handler.OnUserRelationship(*e)
	case *UserPlatformWipeEvent:
		handler.OnUserPlatformWipe(*e)
	case *EmojiCreateEvent:
		handler.OnEmojiCreate(*e)
	case *EmojiDeleteEvent:
		handler.OnEmojiDelete(*e)
	case *AuthEvent:
		handler.OnAuth(*e)
	}
}
