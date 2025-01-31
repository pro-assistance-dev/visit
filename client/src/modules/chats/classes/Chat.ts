import ChatMessage from './ChatMessage';
import ChatUserT from '../types/ChatUserT';

// export default class Chat<UserT extends NameGetter<UserT>> {
export default class Chat<UserT extends ChatUserT> {
  id?: string;
  name = '';
  createdAt = new Date();
  newMessage = '';
  // userId = '';
  @PF.H.Classes.GetClassConstructor(ChatMessage)
  chatMessages: ChatMessage<UserT>[] = [];
  userConstructor?: UserT;

  constructor(i?: Chat<UserT>, userConstructor?: UserT) {
    PF.H.Classes.BuildClass(this, i);
    this.userConstructor = userConstructor;
    if (!userConstructor) {
      return;
    }
    this.chatMessages = this.chatMessages.map((c) => new ChatMessage(c, userConstructor));
  }

  static Create<UserT extends ChatUserT>(): Chat<UserT> {
    const item = new Chat<UserT>();
    item.id = PF.H.Id.UUID();
    return item;
  }

  addMessage(message: ChatMessage<UserT>) {
    const item = ChatMessage.Create<UserT>(message.message, this.id as string, message.userId as string);
    if (this.userConstructor) {
      item.user = new this.userConstructor(message.user) as UserT;
    }
    this.chatMessages.push(item);
  }

  createMessage(user: UserT): ChatMessage<UserT> | undefined {
    if (this.newMessage === '') {
      return;
    }
    const item = ChatMessage.Create<UserT>(this.newMessage, this.id as string, user.id as string);
    if (this.userConstructor) {
      item.user = new this.userConstructor(user) as UserT;
    }
    this.chatMessages.push(item);
    this.newMessage = '';
    return item;
  }
}
