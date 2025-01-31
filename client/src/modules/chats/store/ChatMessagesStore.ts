// import BaseStore from '@/core/helpers/BaseStore';
import ChatMessage from '@/modules/chats/classes/ChatMessage';

class S<UserT> extends PF.H.BaseStore<ChatMessage<UserT>> {
  constructor() {
    super(ChatMessage, 'chat-messages');
  }
}

const store: S<unknown> = new S();
export default store;
