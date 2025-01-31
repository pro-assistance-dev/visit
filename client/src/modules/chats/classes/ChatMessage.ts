import ChatMessageTypes from '../types/ChatMessageTypes';

import ChatUserT from '../types/ChatUserT';

export default class ChatMessage<UserT extends ChatUserT> {
  id?: string;
  userId?: string;

  user: ChatUserT = {} as UserT;
  chatId?: string;
  message = '';
  fileInfo = new PF.C.FileInfo();
  createdAt = new Date();
  type: ChatMessageTypes = ChatMessageTypes.Message;

  constructor(i?: ChatMessage<UserT>, userConstructor?: UserT) {
    PF.H.Classes.BuildClass(this, i);
    if (!userConstructor) {
      return;
    }
    this.user = new userConstructor(this.user) as UserT;
  }

  static Create<UserT extends ChatUserT>(message: string, chatId: string, userId: string, userConstructor?: UserT): ChatMessage<UserT> {
    const item = new ChatMessage();
    item.id = PF.H.Id.UUID();
    item.chatId = chatId;
    item.userId = userId;
    item.message = message;
    return item;
  }

  static CreatePingMessage<UserT extends ChatUserT>(): ChatMessage<UserT> {
    const item = new ChatMessage<UserT>();
    item.type = ChatMessageTypes.Ping;
    item.id = PF.H.Id.UUID();
    return item;
  }
  isMessage(): boolean {
    return this.type === ChatMessageTypes.Message || this.type === ChatMessageTypes.Empty;
  }

  getUserName(): string {
    return this.user?.getName() ?? '';
  }
}
